# Invitation API

## 概要
ある話題に対して、招待コードを作成する。  
また、招待コードから、話題のIDを取得することが可能。  

## ローカル(CI)環境で実行する場合

1. Firebase RealtimeDatabaseを作成
2. Firebase > 設定 > プロジェクトの設定 > サービスアカウント から秘密鍵を作成
3. 秘密鍵のパスを指定

```bash
export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/service-account-file.json"
```

## テストを実行
テスト時にはワーキングディレクトリが随時変化していくので、  
以下のようにタグを指定して実行することで、.envを読み込まないようにする。
```shell
go test -v ./... -tags=test
```

## 環境変数
| 環境変数名                     | 説明                                                                       | サンプル                                         |
| ------------------------------ | -------------------------------------------------------------------------- | ------------------------------------------------ |
| ENV                            | アプリケーションの実行モード(この値に対応する`.env`ファイルが読み込まれる) | `development`, `production`                      |
| CLIENT_ENTRY_POINT             | クライアントのサーバーを指すURL                                            | `http://localhost:3000`                          |
| GOOGLE_APPLICATION_CREDENTIALS | 秘密鍵ファイルのパス(ローカルでデバッグ時に必須)                           | `/home/user/Downloads/service-account-file.json` |


### `.env`ファイルのテンプレート
- デバッグ
```.env
GOOGLE_APPLICATION_CREDENTIALS=YOUR_CREDENTIAL_FILE_PATH
VERSION=development
```
- プロダクション
```.env
VERSION=production
```
- `.env.xxx`
```shell
CLIENT_ENTRY_POINT=https://your_clinet_entry_point.com
```

## CirceCI

| パラメータ                       | 説明                                                                                    |
| -------------------------------- | --------------------------------------------------------------------------------------- |
| `GOOGLE_APPLICATION_CREDENTIALS` | (**必須**) Googleサービスアカウントの秘密鍵となるJSONファイルをbase64でエンコードした値 |
| `GOOGLE_PROJECT_ID`              | (**必須**) デプロイ先のGCPプロジェクトID                                                |
