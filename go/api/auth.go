package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret []byte

func init() {
	// Загружаем .env
	_ = godotenv.Load()

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("❌ JWT_SECRET не найден в .env")
	}
	jwtSecret = []byte(secret)
}

func writeJSON(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"message": message,
	})
}

// 🔹 Регистрация
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, "Неверный запрос")
		return
	}
	// Хэш пароля
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := User{Email: req.Email, Password: string(hash)}
	if err := h.DB.Create(&user).Error; err != nil {
		writeJSON(w, http.StatusInternalServerError, "Ошибка создания пользователя")
		return
	}

	writeJSON(w, http.StatusCreated, "Пользователь создан")
}

// 🔹 Логин
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, "Неверный запрос")
		return
	}

	var user User
	if err := h.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		writeJSON(w, http.StatusUnauthorized, "Неверный email или пароль")
		return
	}

	// Проверка пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		writeJSON(w, http.StatusUnauthorized, "Неверный email или пароль")
		return
	}

	// Генерация токенов
	access, refresh, err := generateTokens(user.ID, h)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Отправляем JSON с токенами
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refresh, // plain token, который мы хэшируем и сохраняем в БД
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		Path:     "/",                  // можно '/' или конкретный путь
		HttpOnly: true,                 // ключевой флаг — JS не увидит cookie
		SameSite: http.SameSiteLaxMode, // или Strict
		// Secure:   false,                // локально можно убрать
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(TokenResponse{
		AccessToken: access,
		UserID:      user.ID,
	})
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, "Refresh token missing", http.StatusUnauthorized)
		return
	}

	UserID, err := ValidateRefreshToken(cookie.Value)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	hash := hashToken(cookie.Value)

	var token RefreshToken
	if err := h.DB.Where("user_id = ? AND token_hash = ?", UserID, hash).First(&token).Error; err != nil {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	// проверка exp
	if token.ExpiresAt.Before(time.Now()) {
		h.DB.Delete(&token)
		http.Error(w, "Refresh token expired", http.StatusUnauthorized)
		return
	}

	access, refresh, err := generateTokens(UserID, h)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, "Ошибка генерации токена")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refresh, // plain token, который мы хэшируем и сохраняем в БД
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		Path:     "/",  // можно '/' или конкретный путь
		HttpOnly: true, // ключевой флаг — JS не увидит cookie
		// Secure:   true,                    // требует HTTPS (true в проде; локально mkcert)
		SameSite: http.SameSiteLaxMode, // защита от CSRF (варианты: Lax/Strict/None)
	})

	json.NewEncoder(w).Encode(TokenResponse{
		AccessToken: access,
		UserID:      UserID,
	})
}

func ValidateRefreshToken(refreshString string) (uint, error) {
	// Парсим токен
	token, err := jwt.Parse(refreshString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("неверный метод подписи")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return 0, err
	}

	// Проверяем claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("невалидный токен")
	}

	// Проверка срока жизни
	if exp, ok := claims["exp"].(float64); ok {
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			return 0, errors.New("токен истёк")
		}
	} else {
		return 0, errors.New("нет exp в токене")
	}

	// Получаем user_id
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("нет user_id в токене")
	}

	return uint(userIDFloat), nil
}

func ValidateAccessToken(r *http.Request) (uint, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, errors.New("нет заголовка Authorization")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return 0, errors.New("неверный формат Authorization")
	}

	tokenStr := parts[1]

	// Парсим JWT
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("неверный метод подписи")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("невалидный токен")
	}

	// Проверка срока жизни
	if exp, ok := claims["exp"].(float64); ok {
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			return 0, errors.New("токен истёк")
		}
	} else {
		return 0, errors.New("нет exp в токене")
	}

	// Получаем user_id
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("нет user_id в токене")
	}

	return uint(userIDFloat), nil
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := ValidateAccessToken(r)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Всё ок, userID известен, можно выполнять защищённую логику
	w.Write([]byte(fmt.Sprintf("Hello user: %d. \n Is a secret information -_-\n - %d", userID, time.Now().Unix())))
}

// --- JWT генерация ---
func hashToken(token string) string {
	h := sha256.New()
	h.Write([]byte(token))
	return hex.EncodeToString(h.Sum(nil))
}
func generateTokens(userID uint, h *AuthHandler) (string, string, error) {
	// Access (15 мин)
	accessTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(15 * time.Second).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	// Refresh (7 дней)
	refreshTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	hash := hashToken(refreshString)

	// Удаляем старые токены и создаем новый
	if err := h.DB.Where("user_id = ?", userID).Delete(&RefreshToken{}).Error; err != nil {
		return "", "", err
	}
	if err := h.DB.Create(&RefreshToken{
		UserID:    userID,
		TokenHash: string(hash),
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour), // 7 дней
	}).Error; err != nil {
		return "", "", err
	}

	return accessString, refreshString, nil
}
