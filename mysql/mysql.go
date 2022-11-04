package mysql

import (
	dbutils "github.com/byzk-worker/go-db-utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Client struct {
	url string
	*dbutils.GormCommon
}

func (c *Client) Init() error {
	return c.GormCommon.Init(mysql.Open(c.url), &gorm.Config{QueryFields: true})
}
