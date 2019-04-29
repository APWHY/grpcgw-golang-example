package pet

import (
	"database/sql"
	"time"
)

// xorm tags: gobook.io/read/github.com/go-xorm/manual-en-US/chapter-02/4.columns.html
// valid tags: https://github.com/asaskevich/govalidator

// Pet -- an example object that corresponds to a 'pet' table in our database
type Pet struct {
	ID      int32          `xorm:"pk autoincr" valid:"-"`
	Name    string         `xorm:"varchar(20)" valid:"alpha"`
	Owner   string         `xorm:"varchar(20)" valid:"-"`
	Species string         `xorm:"varchar(20)" valid:"-"`
	Sex     string         `xorm:"char(1)" valid:"-"`
	Birth   string         `xorm:"date" valid:"-"`
	Death   sql.NullString `xorm:"date" valid:"-"` //TODO figure out why this doesn't work
	Created time.Time      `xorm:"created" valid:"-"`
	Updated time.Time      `xorm:"updated" valid:"-"`
}

// Clean -- Helper function to sanitise the struct before insertion
func (p *Pet) Clean() {
	p.ID = 0
	p.Created = time.Unix(0, 0)
	p.Updated = time.Unix(0, 0)
}
