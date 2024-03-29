package sqlite

import (
	dbutils "github.com/byzk-worker/go-db-utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Client struct {
	url string
	*dbutils.GormCommon
}

func (c *Client) Init() error {
	return c.GormCommon.Init(sqlite.Open(c.url), &gorm.Config{QueryFields: true})
}
