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

## CirceCI

| パラメータ                       | 説明                                                                                    |
| -------------------------------- | --------------------------------------------------------------------------------------- |
| `GOOGLE_APPLICATION_CREDENTIALS` | (**必須**) Googleサービスアカウントの秘密鍵となるJSONファイルをbase64でエンコードした値 |
| `GOOGLE_PROJECT_ID`              | (**必須**) デプロイ先のGCPプロジェクトID                                                |
