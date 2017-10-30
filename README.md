# Timebooster

Stable [![CircleCI](https://circleci.com/gh/topgate/timebooster-oss/tree/master.svg?style=svg&circle-token=df62020d029fce5c29848a01b958a89883b749e5)](https://circleci.com/gh/topgate/timebooster-oss/tree/master)

`Timebooster` はCircleCI等のCIとGoogle Compute Engine(GCE)を連携させてビルドを高速に行うためのシステムです。
CircleCIに比べて非常に強力なマシンを使用できるため、ビルド速度を飛躍的に加速させることができます。

Androidプロジェクトの場合、30～40分のビルドを5～10分程度にできます。

導入支援、デモ環境についてのお問い合わせは [株式会社トップゲート](https://www.topgate.co.jp/) にお願いします。

ビルドにはdockerを使用するため、 `Dockerfile` が必要になります。
Dockerfileをビルド、もしくはdocker pullを行なうため、初回のみビルド時間が長くなります。

また、何らかの理由でdocker上でビルドが行えない場合はTimeboosterを使用できません。
例として、iOSアプリ開発ではmacOSが必須のため使用できません。

ビルドはデフォルトでプリエンプティブVMを使用し、ビルドタスクが無い時間は自動的にシャットダウンするため、マシンスペックに対して運用コストが非常に安価です。
ただしプリエンプティブVMの仕様上、稀にビルド中にVMがシャットダウンされます。その場合、ビルドはタイムアウトします。

Timebooster自体も、Timebooster上でビルドされています。

## システム概要

Timeboosterは3つの要素で構成されています。

### timeboosterコマンド

CircleCIからGCEへビルドをリクエストするために使用するコマンドです。
ビルドリクエストは一旦GAE/Goに保存され、GCEに対してビルドリクエストを投げます。

GCEでビルド後の成果物はFirebase Storageに保存され、自動的にダウンロードされます。

### tbrunコマンド

GCE内部で使用されます。
startup-scriptで自動的にダウンロード＆実行されます。

timeboosterコマンドでリクエストされたタスクを処理し、成果物をFirebase Storageにアップロードします。

## GAE/Go

timeboosterコマンドで発行されたビルドリクエストを管理します。
ビルドマシンの生成や起動等、GCEとの連携を行ないます。

# デプロイ手順

初期セットアップや手動デプロイは下記の手順で行ないます。

 1. Google Cloud Platformのコンソールから、Timeboosterをデプロイするためのプロジェクトを作成する
 1. Compute Engineを利用するため、Compute Engine有効化と課金設定を行なう
 1. 作成したGCPとFirebaseとの連携を行なう
 1. Firebase AdminのService Accountを発行する
  * `server/gae/assets/firebase-admin.json` に配置する
 1. Google App Engineデプロイ用Service Accountを発行する
  * デフォルト状態の場合、Firebase AdminのServiceAccountがそのまま利用可能
  * `.timebooster/service-account.json` に配置する
 1. Dockerをインストールする
  * デプロイ作業は基本的にDockerコンテナ内で実行されます
 1. `/deploy.sh` を実行する

# TimeboosterのCircle CIへの導入方法

 1. `GCPのコンソール > APIとサービス > 認証情報` からAPIキーを発行
 1. `POST /api/v1/machine` APIでビルドマシンを生成する
  * ビルドマシンは自動的にセットアップされる
  * デフォルトは `us-central1-a | vCPU24 | RAM 64GB` で生成される
 1. `timebooster` をリポジトリに配置する
 1. `timebooster.yml` を記述する
 1. `timebooster` コマンドからビルドを実行ようにCIを変更する

## timebooster.ymlの記述例

```
env:
  # 環境変数設定
  # デフォルト値-> TIMEBOOSTER_ARTIFACTS=/artifacts
  # デフォルト値-> TIMEBOOSTER_BUILD_ID=一意に特定されるビルドID
  variable:
    FOO: bar
  # キャッシュディレクトリの登録
  # キャッシュディレクトリはビルドマシン（APIキー）単位で共有される
  # dockerの指定Pathに直接マウントされるため、同時書込み耐性については保証されない。
  cache:
    - /root/.m2
    - /root/.gradle/caches/modules-2/files-2.1
task:
    # execは複数個登録できる
    # 現状では逐次実行を前提とする
    exec:
      # `docker build -f 指定dockerfile .` を実行する
      # リポジトリRootからの相対パスで記述する
      - dockerfile: "dockerfiles/android-build.dockerfile"
       # Dockerコンテナ内で実行されるスクリプト
       # リポジトリルートがカレントディレクトリとして実行される
       # /work/{リポジトリ名} にマウントされ、"cd /work/{リポジトリ名}"された状態で実行される
       # 実行時の標準出力は"exec-xxx.txt"に保存され、アーティファクトの一つとして回収される
        cmd:
          - chmod 755 .timebooster/script/timebooster-build.sh
          - .timebooster/script/timebooster-build.sh
# EOF
```

## timebooster実行

下記のコマンドを実行します。

ビルドが成功した場合、ステータスコード0が返却されます。

  * グローバルオプション
    * config
      * timebooster.yamlのパスを指定
    * api-key
      * timeboosterのAPIKeyを指定
    * endpoint
      * APIのエンドポイントを指定
      * 通常 `https://your-gcp-id.appspot.com` を指定する
  * runコマンドオプション
    * repository
      * githubリポジトリURLを指定
    * github-access-token
      * cloneのためのアクセストークンを指定(プロジェクトメンバーの誰かが発行する）
    * revision
      * ビルドリビジョンを指定（optional)
      * 指定がない場合、cloneしたリポジトリをそのままビルドする
    * artifact
      * ビルド成果物ダウンロードディレクトリ
    * env-from-circleci
      * CircleCIが定義している環境変数をビルド時に引き継ぐ
    * env
      * 指定した環境変数をロードし、ビルド時に引き継ぐ

```
chmod 755 path/to/timebooster
path/to/timebooster \
  -config path/to/timebooster.yaml \
  -api-key your-api-key \
  -endpoint https://your-gcp-id.appspot.com
  run \
  -repository ${CIRCLE_REPOSITORY_URL} \
  -github-access-token ${GITHUB_API_KEY} \
  -revision ${CIRCLE_SHA1} \
  -artifact ${CIRCLE_ARTIFACTS} \
  -env-from-circleci \
  -env DEPLOYGATE_USER_NAME \
  -env DEPLOYGATE_API_KEY
```

## 未実装機能(TODO)

 * 指定した公開鍵を利用したclone
    * 現時点ではgithubからのcloneをAPIキーを使用して行います。
