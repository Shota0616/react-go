# ビルドステージ
FROM node:20 AS builder
WORKDIR /app
COPY react/vite-project/package*.json ./
RUN npm install
COPY react/vite-project .
RUN npm run build

# 本番ステージ
FROM node:20-alpine AS production
WORKDIR /app
ENV NODE_ENV=production
COPY --from=builder /app/dist ./dist
COPY --from=builder /app/node_modules ./node_modules
RUN npm install -g serve
USER node
CMD ["serve", "-s", "dist", "-l", "5173"]

# 開発ステージ
FROM node:20 AS development
WORKDIR /app
COPY react/vite-project/package*.json ./
RUN npm install -g vite && npm install
COPY react/vite-project .
USER node
CMD ["vite", "dev", "--host", "0.0.0.0"]
