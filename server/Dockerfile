FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем приложение
RUN go build -v -o build/lit-log cmd/main.go

# # Создаем финальный образ
# FROM alpine:latest

# # Устанавливаем рабочую директорию
# WORKDIR /app

# # Копируем собранное приложение из предыдущего этапа
# COPY --from=builder /app/build/lit-log ./build/lit-log

# # Создаем папку dist
# COPY dist ./dist
# RUN mkdir -p db

# Указываем команду по умолчанию для запуска приложения
CMD ["./build/lit-log", "--port=80"]