FROM golang:alpine3.20

# 作業ディレクトリを設定
WORKDIR /go/src/app

# ローカルのsrcディレクトリをコンテナ内の/appにコピー
COPY . .

# メインアプリケーションを実行
CMD ["go", "run", "main.go"]