FROM node:18 AS builder

WORKDIR /app

COPY package.json yarn.lock ./

RUN yarn install --frozen-lockfile

COPY . ./

FROM node:18

WORKDIR /app

# Копируем файлы из builder-контейнера
COPY --from=builder /app ./

EXPOSE 8080

CMD ["yarn", "serve"]
