#!/bin/bash

# 概要
#   Cloud Functionsでは,.envファイルの読み込みができないため、
#   環境変数で直接指定する必要がある。
#   このため、ローカルで簡単に動作確認をする事ができるスクリプトである。
# 注意点
#   必ずプロジェクトのルートディレクトリ上で実行する。
#   秘密鍵のパス(GOOGLE_APPLICATION_CREDENTIALS)を予めセットする


if [ "$GOOGLE_APPLICATION_CREDENTIALS" = "" ]; then
    echo "GOOGLE_APPLICATION_CREDENTIALS is empty"
    exit 1
fi
echo "$GOOGLE_APPLICATION_CREDENTIALS"

export VERSION=CF
export CLIENT_ENTRY_POINT=https://omochat.app
go run cmd/server.go
