# go-docker-sample

Dockerを使ってGoの開発環境を構築し、ソースをGitHubにpushすることで、cloneしてvscodeで編集できるようにする。

cloneして、vscodeのRemote Development拡張機能で開くとコンテナが作成される感じ？

Docker、Docker Composeを使用するので、インストールしておいてください。
Goは入っていなくても問題ないはず。

```
$ cd [WORKSPACE] #お好きなディレクトリに移動
$ git clone https://github.com/kazuhiro-f/go-docker-sample #このリポジトリをクローン
$ docker-compose build #Dockerイメージをビルド
$ docker-compose up -d #Docker Composeでコンテナ作成
$ docker-compose exec app go run main.go #作成したコンテナ内でgo run main.goコマンドを実行

$ docker-compose exec app /bin/bash #作成したコンテナ内でシェルを起動
$ go run main.go #作成したコンテナ内でgo run main.goを実行(上と同じ結果)
$ exit #シェルを終了

$ docker-compose stop app #コンテナを停止
$ docker-compose rm app #イメージを削除
```
