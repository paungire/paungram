FROM golang:1.25

WORKDIR /webrtc

COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./

# Для разработки можно ставить Air, чтобы перезапускать при изменении кода
RUN go install github.com/air-verse/air@v1.62.0

CMD ["air", "-c", "air.toml"]
