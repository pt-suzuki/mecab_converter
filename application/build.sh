#!/bin/sh

export CGO_LDFLAGS="-L/usr/local/Cellar/mecab/0.996/lib -lmecab -lstdc++"
export CGO_CFLAGS="-I/usr/local/Cellar/mecab/0.996/include"
go build cmd/main.go