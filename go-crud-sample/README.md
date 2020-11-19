# go-crud-sample

サーバサイド講義用のサンプルコード

## Requirement

### Versions

- Go v1.15  
- macOS 64 bit  
- Windows10 64 bit  
- Docker Engine v19.03  

### Directories

参考: golang-standards/project-layout

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

`hoge`

TODO: Windows 用のインストール方法を追加  
TODO: 今後 air でライヴリロードする

### Env

- configs 配下の .env で管理
- docker-compose.yml の `env_file` に Path を指定
- YAML からの取得がスタンダードだが Go コードが増えるため .env で簡易化

## Execution

### docker-compose を使う場合

`$ docker-compose up --build -d`

### docker-compose を使わない場合

`hoge`

TODO: docker build/run/exec のコマンド追加
