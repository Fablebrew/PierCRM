# Этап сборки - используем конкретную версию на Alpine
FROM golang:1.25-alpine AS builder

# Установка git (apk работает в Alpine)
RUN apk add --no-cache git

WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Копируем исполняемый файл
COPY --from=builder /app/main .

# Копируем миграции
COPY --from=builder /app/database ./database

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]