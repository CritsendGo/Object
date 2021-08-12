package csObject

import (
	"time"
)

type Route struct {
	Id     int       `json:"route_id" oType:"int" oPrimary:"true" oOptional:"true"`
	Name   string    `json:"route_name" oType:"string" oPrimary:"false" oOptional:"false"`
	Create time.Time `json:"message_create" oType:"time" oPrimary:"false" oOptional:"true"`
}

func (o *Route) Load() error { return Get(o) }
func (o *Route) Save() error { return Save(o) }
func (o *Route) LoadOrCreate() error {
	if Get(o) != nil {
		return Save(o)
	} else {
		return nil
	}
}
