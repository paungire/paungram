import { defineStore } from "pinia";

export const useAuthStore = defineStore("auth", {
	state: () => ({
		accessToken: null,
		userId: null,
		isLoggedIn: false,
		refreshPromise: null,
	}),
	actions: {
		setToken(token, userId) {
			this.accessToken = token;
			this.userId = userId;
			this.isLoggedIn = !!token;
		},
		clearToken() {
			this.accessToken = null;
			this.userId = null;
			this.isLoggedIn = false;
		},

		async refresh() {
			if (this.refreshPromise) {
				return this.refreshPromise;
			}

			let resolveAll, rejectAll;
			const promise = new Promise((resolve, reject) => {
				resolveAll = resolve;
				rejectAll = reject;
			});
			this.refreshPromise = promise;

			try {
				const response = await fetch("/api/v1/auth/refresh", {
					method: "POST",
					credentials: "include",
					headers: {
						"Content-Type": "application/json",
					},
				});
				if (!response.ok) {
					throw new Error("Не удалось обновить токен");
				}
				const data = await response.json();
				if (data?.access_token && data?.user_id) {
					this.setToken(data.access_token, data.user_id);
					resolveAll(data.access_token); // ✅ разруливаем всех ждунов
				} else {
					this.clearToken();
					rejectAll(new Error("Некорректный ответ refresh"));
				}
			} catch (error) {
				console.error("Ошибка refresh:", error);
				this.clearToken();
				rejectAll(error); // ❌ сообщаем всем ждунам об ошибке
			} finally {
				this.refreshPromise = null;
			}

			return promise;
		},

		getAccessToken() {
			return this.accessToken;
		},
		getUserId() {
			return this.userId;
		},
		getIsLoggedIn() {
			return this.isLoggedIn;
		},
	},
	persist: true,
});
