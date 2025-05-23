# React + Go 開発環境テンプレート

このプロジェクトはReact (Vite)とGoを使ったWebアプリケーションの開発テンプレートです。Dockerを使用して開発環境を構築できます。

## 必要な環境
- Docker
- Docker Compose
- Git

## セットアップ手順

1. リポジトリをクローン:
```bash
git clone git@github.com:Shota0616/react-go.git
cd react-go
```

2. 環境変数ファイルの作成:
```bash
cp .env.example .env
```

## viteプロジェクト作成

```bash
cd react
docker run --rm -it -v "$PWD":/app -w /app node:20 bash
```

```bash
npm create vite@latest
cd vite-project/
npm install
exit
```

## 開発環境の起動

```bash
docker-compose up --build
```

- Go API: http://localhost:8080/api/hello
- Go API: http://localhost:8080/api/db-test
- Reactアプリ: http://localhost:3000

## 本番環境のビルドと起動

```bash
ENV_MODE=production docker-compose up --build
```

### ENV_MODEについて
- `development` (デフォルト): ホットリロード有効の開発用設定
- `production`: 最適化された本番用ビルド
- docker-compose.ymlと各Dockerfileでこの値を使用してビルドターゲットを切り替え

## ディレクトリ構成

```
.
├── go/                # Goバックエンド
│   ├── main.go        # メインアプリケーション
│   ├── go.mod         # Goモジュール定義
│   └── Dockerfile     # Goコンテナ設定
├── react/             # Reactフロントエンド
│   └── vite-project/  # Viteプロジェクト
└── docker-compose.yml # 全体のDocker設定
```
