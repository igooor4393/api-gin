# Используем официальный образ Go для компиляции нашего приложения.
FROM golang:latest as builder

# Устанавливаем рабочую директорию внутри контейнера.
WORKDIR /app

# Копируем файлы go.mod и go.sum для кэширования зависимостей.
COPY go.mod go.sum ./

# Загружаем зависимости. Это будет кэшироваться, если файлы go.mod и go.sum не изменятся.
RUN go mod download

# Копируем все файлы из локального контекста в рабочую директорию контейнера.
COPY . .

# Собираем наше приложение.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

# Используем образ alpine для минималистичного финального образа.
FROM alpine:latest

# Устанавливаем нужные пакеты, такие как ca-certificates.
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем скомпилированный исполняемый файл из стадии builder в финальный образ.
COPY --from=builder /app/myapp .
# Копируем файл конфигурации в корень финального образа
COPY app.env /root/app.env
# Открываем порт, который используется нашим приложением.
EXPOSE 8080

# Запускаем приложение.
CMD ["./myapp"]