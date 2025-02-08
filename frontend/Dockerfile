FROM node:18 AS builder

WORKDIR /app

COPY frontend/package.json frontend/yarn.lock ./

RUN yarn install --frozen-lockfile

COPY frontend ./

FROM node:18

WORKDIR /app

# Копируем файлы из builder-контейнера
COPY --from=builder /app ./

EXPOSE 80

CMD ["yarn", "serve"]
