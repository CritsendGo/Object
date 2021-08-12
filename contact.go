package csObject

import (
	"time"
)

type Contact struct {
	Id     int       `json:"contact_id"  oType:"int" oPrimary:"true" oOptional:"true"`
	Name   string    `json:"contact_name" oType:"string" oPrimary:"false" oOptional:"false"`
	Create time.Time `json:"contact_create" oType:"time" oPrimary:"false" oOptional:"true"`
}

func (o *Contact) Load() error { return Get(o) }
func (o *Contact) Save() error { return Save(o) }
func (o *Contact) LoadOrCreate() error {
	if Get(o) != nil {
		return Save(o)
	} else {
		return nil
	}
}
