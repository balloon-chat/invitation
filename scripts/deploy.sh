if [ "$GOOGLE_SERVICE_ACCOUNT_CREDENTIALS" = "" ]; then
    echo "GOOGLE_SERVICE_ACCOUNT_CREDENTIALS is empty"
    exit 1
fi

circleci local execute --job=deploy --env GCLOUD_SERVICE_KEY=`base64 $GOOGLE_SERVICE_ACCOUNT_CREDENTIALS` --env GOOGLE_PROJECT_ID=balloon-6bad2
