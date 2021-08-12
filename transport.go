package csObject

import (
	"time"
)

type Transport struct {
	Id     int       `json:"transport_id" oType:"int" oPrimary:"true" oOptional:"true"`
	Name   string    `json:"transport_name" oType:"string" oPrimary:"false" oOptional:"false"`
	Update time.Time `json:"transport_create" oType:"time" oPrimary:"false" oOptional:"true"`
	Create time.Time `json:"transport_update" oType:"time" oPrimary:"false" oOptional:"true"`
}

func (o *Transport) Load() error { return Get(o) }
func (o *Transport) Save() error { return Save(o) }
func (o *Transport) LoadOrCreate() error {
	if Get(o) != nil {
		return Save(o)
	} else {
		return nil
	}
}
