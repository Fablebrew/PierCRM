# Этап сборки - используем Node.js 22 (LTS)
FROM node:22-alpine AS builder

WORKDIR /app

# Копируем файлы зависимостей
COPY package*.json ./

# Устанавливаем зависимости
RUN npm install

# Копируем исходный код
COPY . .

# Сборка приложения
RUN npm run build

# Финальный образ
FROM node:22-alpine

WORKDIR /app

# Устанавливаем serve для раздачи статических файлов
RUN npm install -g serve

# Копируем собранное приложение
COPY --from=builder /app/dist ./dist

EXPOSE 5173

# Запускаем serve
CMD ["serve", "-s", "dist", "-l", "5173"]