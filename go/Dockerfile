# ベースイメージ
FROM golang:1.22.4

WORKDIR /usr/src/app

# ソースコードのコピー
COPY . .

# コンテナbuild時に実行するコマンド
# 依存関係のインストール
RUN go mod download
# CompileDaemon のインストール
RUN go install github.com/githubnemo/CompileDaemon@latest

# 開発環境では CompileDaemon で自動再ビルドし、本番環境ではビルド済みのファイルを実行
# CMDはコンテナ起動時に実行されるコマンド
CMD if [ "$ENV_MODE" = "production" ]; then \
        ./main; \
    else \
        CompileDaemon -build="go build -o main cmd/api/main.go" -command="./main"; \
    fi