FROM golang:alpine3.20

# 作業ディレクトリを設定
WORKDIR /go/src/app

# ローカルのsrcディレクトリをコンテナ内の/appにコピー
COPY . .

# アプリケーションを実行用ファイルの設定
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# アプリケーションを実行
ENTRYPOINT ["/entrypoint.sh"]