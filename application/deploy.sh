#!/bin/sh

gcloud functions deploy MecabConverterFunction \
  --project pig-allowance-book-core \
  --region=asia-northeast1 \
  --runtime go113 \
  --trigger-http \
  --allow-unauthenticated \
  --set-env-vars=CGO_LDFLAGS=./mecab_lib/mecab_lib,CGO_CFLAGS=./mecab_lib/include
