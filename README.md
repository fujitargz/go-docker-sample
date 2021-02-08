# go-docker-sample

Dockerを使ってGoの開発環境を構築し、ソースをGitHubにpushすることで、cloneしてvscodeで編集できるようにする。

cloneして、vscodeのRemote Development拡張機能で開くとコンテナが作成される感じ？

```
$ git clone https://github.com/kazuhiro-f/go-docker-sample
$ docker-compose build
$ docker-compose up -d
$ docker-compose exec app go run main.go
```
