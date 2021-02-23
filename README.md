# go-docker-sample

## Overview

Dockerを使ってGoの開発環境を構築し、ソースをGitHubにpushすることで、cloneしてVSCodeで編集できるようにする。

docker-composeをインストールして動作させるか、
cloneして、VSCodeのRemote Development拡張機能で開くとコンテナが作成される。

## Usage

### Preparation

Docker、docker-composeを使用するので、インストールしておいてください。

### Clone Repository

```
$ git clone https://github.com/kazuhiro-f/go-docker-sample #このリポジトリをクローン
$ cd go-docker-sample #ワーキングディレクトリに移動
```

### Exec with docker-compose

```
$ docker-compose build #Dockerイメージをビルド
$ docker-compose up -d #Docker Composeでコンテナ作成
$ docker-compose exec app go run main.go #作成したコンテナ内でgo run main.goコマンドを実行
(ブラウザでlocalhost:8000にアクセス)

$ docker-compose exec app /bin/bash #作成したコンテナ内でシェルを起動
$ go run main.go #作成したコンテナ内でgo run main.goを実行(上と同じ結果)
(ブラウザでlocalhost:8000にアクセス)
$ exit #シェルを終了

$ docker-compose stop #コンテナを停止
$ docker-compose rm #イメージを削除
```

### Edit with VSCode

VSCodeの拡張機能「Remote Development」または「Remote - Containers」のインストールが必須。
(厳密には、「Remote Development」は「Remote - WSL」「Remote - Containers」「Remote - SSH」をまとめた拡張機能パッケージなので、必要に応じて選択)

VSCodeを起動後、ウインドウ左下端の緑の部分をクリック。
メニューから「Remote-Containers:Open Folder in Container...」をクリック。
git cloneしたdocker-composeディレクトリを選択。
(初回はコンテナをビルドするため時間がかかるが)コンテナが起動し、編集可能になる。

## Deploy to Heroku

```
$ heroku login # CLIからログイン
$ heroku create <appname> # herokuアプリを作成
$ heroku stack:set container # スタックタイプをコンテナに変更
$ git push heroku <branchname> # プッシュ
$ heroku open # アプリを起動
```

HerokuのAutomatic deploysで、特定のブランチが更新される度に自動デプロイすることも可能(?)
