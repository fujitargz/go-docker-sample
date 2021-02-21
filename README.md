# go-docker-sample

## Overview

Dockerを使ってGoの開発環境を構築し、ソースをGitHubにpushすることで、cloneしてvscodeで編集できるようにする。

docker-composeをインストールして動作させるか、
cloneして、vscodeのRemote Development拡張機能で開くとコンテナが作成される。

## Usage

#### 事前準備

Docker、Docker Composeを使用するので、インストールしておいてください。
Goは入っていなくても問題ないはず。

#### リポジトリのクローン

```
$ git clone https://github.com/kazuhiro-f/go-docker-sample #このリポジトリをクローン
$ cd go-docker-sample #ワーキングディレクトリに移動
```
#### docker-composeで起動・実行する場合

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

#### VSCodeで編集する場合

VSCodeの拡張機能「Remote Development」または「Remote - Containers」のインストールが必須。
(厳密には、「Remote Development」は「Remote - WSL」「Remote - Containers」「Remote - SSH」をまとめた拡張機能パッケージなので、必要に応じて選択)

VSCodeを起動後、ウインドウ左下端の緑の部分をクリック。
メニューから「Remote-Containers:Open Folder in Container...」をクリック。
git cloneしたdocker-composeディレクトリを選択。
(初回はコンテナをビルドするため時間がかかるが)コンテナが起動し、編集可能になる。
