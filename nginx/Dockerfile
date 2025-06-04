FROM nginx:1.26.2

# logrotateのインストール
RUN apt-get update && apt-get install -y logrotate

# cronジョブの設定
RUN echo "0 0 * * * /usr/sbin/logrotate /etc/logrotate.d/nginx" > /etc/cron.d/logrotate

# 実行権限の付与
RUN chmod 0644 /etc/cron.d/logrotate