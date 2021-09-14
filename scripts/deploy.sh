#!/bin/bash

# 概要
#   Cloud Functionsへ、ローカルのファイルをデプロイする
# 注意点
#   必ずプロジェクトのルートディレクトリ上で実行する。

gcloud config set project balloon-6bad2
gcloud functions deploy create-invitation \
  --region asia-northeast1 \
  --entry-point CreateInvitation \
  --runtime go113 \
  --trigger-http \
  --set-env-vars VERSION=CF,CLIENT_ENTRY_POINT=https://omochat.app
