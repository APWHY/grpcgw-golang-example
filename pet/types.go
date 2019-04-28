package pet

import (
	"database/sql"
	"time"
)

type Pet struct {
	Id      int            `xorm:"pk"`
	Name    string         `xorm:"varchar(20)"`
	Owner   string         `xorm:"varchar(20)"`
	Species string         `xorm:"varchar(20)"`
	Sex     string         `xorm:"char(1)"`
	Birth   string         `xorm:"date"`
	Death   sql.NullString `xorm:"date default '2019-01-01'"` //TODO figure out why this doesn't work
	Created time.Time      `xorm:"created"`
	Updated time.Time      `xorm:"updated"`
}

// for more info on xorm column mappings go to the below site:
// http://gobook.io/read/github.com/go-xorm/manual-en-US/chapter-02/4.columns.html
