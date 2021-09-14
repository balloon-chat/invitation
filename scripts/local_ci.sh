#!/bin/zsh

# 概要
#   CircleCiの動作をローカルで確認するためのスクリプト
# 注意点
#   必ずプロジェクトのルートディレクトリ上で実行する。
#   秘密鍵のパス(GOOGLE_APPLICATION_CREDENTIALS)を予めセットする

if [ "$GOOGLE_APPLICATION_CREDENTIALS" = "" ]; then
    echo "GOOGLE_APPLICATION_CREDENTIALS is empty"
    exit 1
fi
echo "$GOOGLE_APPLICATION_CREDENTIALS"

# 秘密鍵をbase64エンコーディングし、環境変数にセット
export GCLOUD_SERVICE_KEY=$(base64 "$GOOGLE_APPLICATION_CREDENTIALS")

circleci local execute --job deploy \
  --env GCLOUD_SERVICE_KEY="$GCLOUD_SERVICE_KEY" \
  --env GOOGLE_PROJECT_ID=balloon-6bad2 \
  --env VERSION=CF \
  --env CLIENT_ENTRY_POINT=https://omochat.app
