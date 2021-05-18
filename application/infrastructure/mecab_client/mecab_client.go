package mecab_client

import (
	"github.com/bluele/mecab-golang"
	"log"
)

type Client interface {
	GetMecabClient() *mecab.MeCab
}

type client struct {
}

func ClientImpl() Client {
	return &client{}
}

func (c *client)GetMecabClient () *mecab.MeCab {
	m, err := mecab.New("-Owakati")

	if err != nil {
		log.Fatalln(err)
	}
	return m
}
