# go-crud-sample

サーバサイド講義用のサンプルコード

## Requirement

### Versions

- Go v1.15  
- macOS 64 bit  
- Windows10 64 bit  
- Docker Engine v19.03  

### Directories

参考: [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

```shell
go-crud-sample
├── bin         ライヴリロード用のバイナリ出力
├── build       Docker 等のパッケージファイル
├── cmd         main 関数が定義されたファイル
├── configs     設定ファイル
├── db          ローカル DB
├── deployments デプロイファイル
├── docs        各種ドキュメント
├── internal    app 内部のコード
└── scripts     作業用スクリプト
```

## Setup

### Installation

Go

参考: [VSCodeでGo言語の開発環境を構築する](https://qiita.com/melty_go/items/c977ba594efcffc8b567)  
TODO: 今後 air でライヴリロードする

### Env

- configs 配下の .env で管理
- docker-compose.yml の `env_file` に Path を指定
- YAML からの取得がスタンダードだが Go コードが増えるため .env で簡易化

## Execution

### docker-compose を使う場合

`% docker-compose up --build -d`

### docker-compose を使わない場合

Docker Image をビルド

```shell
% docker build -f {path/to/dockerfile} -t {image-name:tag}
# dockerfile が多くなると大変
```

Docker Container を起動

```shell
% docker run --rm -it \
  -p {host-port:container-port} \
  -v {host-contents:container-content} \
  --name {container name}
# コマンドが長くなると大変
```

Docker Exec で実行

```shell
% docker exec {container-name} {command}
# Container 内で直接コマンドを実行したいとき
# Container 内で直接エラー確認したいとき
```

### Docker が起動できない場合

Windows の問題でエラーが発生する場合 Docker は使用しない

MySQL 環境構築

[MySQLの開発環境を用意しよう (windows)](https://prog-8.com/docs/mysql-env-win)