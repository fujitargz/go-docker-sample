# 2021/02/08最新version
FROM golang:1.15.8

# ディレクトリ作成
RUN mkdir /go/src/go-docker-sample

#コンテナログイン時のディレクトリを設定
WORKDIR /go/src/go-docker-sample

#ホストのファイルをコンテナにコピー
ADD . /go/src/go-docker-sample
