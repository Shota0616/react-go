# Node.js公式イメージ
FROM node:20

WORKDIR /app

# 依存関係のインストール
COPY package*.json ./
# RUN npm install

# ソースコードのコピー
COPY . .

# 開発用サーバー起動
CMD ["npm", "run", "dev"]

# 本番用ビルド
# RUN npm run build
# CMD ["npm", "run", "preview"]
