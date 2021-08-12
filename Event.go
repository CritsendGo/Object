package csObject

import (
	"time"
)

type Event struct {
	Id      int       `json:"event_id" oType:"int" oPrimary:"true" oOptional:"true"`
	Message *Message  `json:"message_id" oType:"struct" oPrimary:"false" oOptional:"false"`
	Code    int       `json:"event_type_id" oType:"int" oPrimary:"false" oOptional:"false"`
	Create  time.Time `json:"event_create" oType:"time" oPrimary:"false" oOptional:"true"`
	Detail  string    `json:"event_detail" oType:"string" oPrimary:"false" oOptional:"true"`
}

func (o *Event) Load() error { return Get(o) }
func (o *Event) Save() error { return Save(o) }
func (o *Event) LoadOrCreate() error {
	if Get(o) != nil {
		return Save(o)
	} else {
		return nil
	}
}
