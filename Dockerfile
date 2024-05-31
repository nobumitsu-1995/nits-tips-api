FROM golang:alpine3.20

# 作業ディレクトリを設定
WORKDIR /go/src/app

# ローカルのsrcディレクトリをコンテナ内の/appにコピー
COPY src/ .

# メインアプリケーションを実行
CMD ["go", "run", "main.go"]