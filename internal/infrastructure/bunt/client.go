package bunt

import (
	"fmt"
	"github.com/tidwall/buntdb"
)

const (
	IndexWhen = "when"
)

type Client struct {
	db *buntdb.DB
}

func NewClient(options ClientOptions) *Client {
	path := options.Name
	if options.Path != "" {
		path = fmt.Sprintf("%s/%s", options.Path, options.Name)
	}

	db, err := buntdb.Open(path)
	if err != nil {
		panic(err)
	}

	err = db.CreateIndex(IndexWhen, "*", buntdb.IndexJSON("When"))
	if err != nil {
		panic(err)
	}

	return &Client{
		db: db,
	}
}

func (c *Client) Database() *buntdb.DB {
	return c.db
}
