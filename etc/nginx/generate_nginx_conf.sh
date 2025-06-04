# 設定ファイルのパス
AWS_CONFIG="nginx.conf.aws"
PROD_CONFIG="nginx.conf.production"
DEV_CONFIG="nginx.conf.development"
TARGET_CONFIG="default.conf"

# 環境に応じて適切なファイルをコピー
if [ "$AWS_MODE" = "true" ]; then
    cp -f "/etc/nginx/conf.d/$AWS_CONFIG" "/etc/nginx/conf.d/$TARGET_CONFIG"
    echo "AWS環境用の Nginx 設定 ($AWS_CONFIG) を $TARGET_CONFIG にコピーしました。"
elif [ "$NODE_ENV" = "production" ]; then
    cp -f "/etc/nginx/conf.d/$PROD_CONFIG" "/etc/nginx/conf.d/$TARGET_CONFIG"
    echo "本番環境用の Nginx 設定 ($PROD_CONFIG) を $TARGET_CONFIG にコピーしました。"
elif [ "$NODE_ENV" = "development" ]; then
    cp -f "/etc/nginx/conf.d/$DEV_CONFIG" "/etc/nginx/conf.d/$TARGET_CONFIG"
    echo "開発環境用の Nginx 設定 ($DEV_CONFIG) を $TARGET_CONFIG にコピーしました。"
else
    echo "エラー: NODE_ENV は 'production' または 'development' のいずれかである必要があります。"
    exit 1
fi
