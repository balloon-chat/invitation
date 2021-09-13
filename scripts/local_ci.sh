#!/bin/zsh

# 秘密鍵のパスを予めセットする
echo "$GOOGLE_APPLICATION_CREDENTIALS"

# 秘密鍵をbase64エンコーディングし、環境変数にセット
export GCLOUD_SERVICE_KEY=$(base64 "$GOOGLE_APPLICATION_CREDENTIALS")

# .envファイルを追加
export DOT_ENV=$(base64 ./.env)

circleci local execute --job deploy \
  --env GCLOUD_SERVICE_KEY="$GCLOUD_SERVICE_KEY" \
  --env GOOGLE_PROJECT_ID=balloon-6bad2 \
  --env DOT_ENV="$DOT_ENV"
