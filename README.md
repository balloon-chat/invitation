# Invitation API

## 概要

ある話題に対して、招待コードを作成する。  
また、招待コードから、話題のIDを取得することが可能。

## テストを実行

テスト時にはワーキングディレクトリが随時変化していくので、  
以下のようにタグを指定して実行することで、.envを読み込まないようにする。

```shell
go test -v ./... -tags=test
```

## Circle CI

### 準備
1. Firebase RealtimeDatabaseを作成
2. Firebase > 設定 > プロジェクトの設定 > サービスアカウント から秘密鍵を作成
3. 秘密鍵のパスを指定

### 環境変数
| 環境変数名                       | 説明                                                                         |
| -------------------------------- | ---------------------------------------------------------------------------- |
| `GOOGLE_APPLICATION_CREDENTIALS` | Googleサービスアカウントの秘密鍵となるJSONファイルをbase64でエンコードした値 |
| `GOOGLE_PROJECT_ID`              | デプロイ先のGCPプロジェクトID                                                |
| `VERSION`                        | `VERSION=CF`とする                                                           |
| `CLIENT_ENTRY_POINT`             | クライアントのサーバーを指すURL                                              |


## Cloud Functions
Cloud Functionsでは、`.env`ファイルを読み込ませるのが難しいため、直接環境変数を指定する必要がある。

### 環境変数
- `ENV=CF`を設定
- `CLIENT_ENTRY_POINT`を設定

### ローカルでデバッグ
```shell
export GOOGLE_APPLICATION_CREDENTIALS="/your/credential/file/path.json"
sh script/debug_cloud_functions.sh
```
  
### デプロイのテスト
```shell
sh script/deploy.sh
```

## 環境変数

| 環境変数名                     | 説明                                                                       | サンプル                                         |
| ------------------------------ | -------------------------------------------------------------------------- | ------------------------------------------------ |
| VERSION                        | アプリケーションの実行モード(この値に対応する`.env`ファイルが読み込まれる) | `development`, `production`, `CF`(CloudFunction) |
| CLIENT_ENTRY_POINT             | クライアントのサーバーを指すURL                                            | `http://localhost:3000`                          |
| GOOGLE_APPLICATION_CREDENTIALS | 秘密鍵ファイルのパス(ローカルでデバッグ時のみ指定する)                     | `/home/user/Downloads/service-account-file.json` |

## .envファイル

### `.env`ファイルのテンプレート(デバッグ)
```.env
GOOGLE_APPLICATION_CREDENTIALS=YOUR_CREDENTIAL_FILE_PATH
VERSION=development
```
- `.env.development`
```shell
CLIENT_ENTRY_POINT=http://localhost:8000
```

### `.env`ファイルのテンプレート(プロダクション)
- .env
```.env
VERSION=production
```
- `.env.production`
```shell
CLIENT_ENTRY_POINT=https://your_clinet_entry_point.com
```