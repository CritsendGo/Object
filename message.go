package csObject

import (
	"time"
)

type Message struct {
	Id    int    `json:"message_id" oType:"int" oPrimary:"true" oOptional:"true"`
	Token string `json:"message_key" oType:"string" oPrimary:"false" oOptional:"false"`
	//PostfixId string    `json:"message_pid" oType:"string" oPrimary:"false" oOptional:"true"`
	Update    time.Time  `json:"message_create" oType:"time" oPrimary:"false" oOptional:"true"`
	Create    time.Time  `json:"message_update" oType:"time" oPrimary:"false" oOptional:"true"`
	Contact   *Contact   `json:"contact_id" oType:"struct" oPrimary:"false" oOptional:"true"`
	MailFrom  *MailFrom  `json:"mail_from_id" oType:"struct" oPrimary:"false" oOptional:"true"`
	Transport *Transport `json:"transport_id" oType:"struct" oPrimary:"false" oOptional:"true"`
}

func (o *Message) Load() error { return Get(o) }
func (o *Message) Save() error { return Save(o) }
func (o *Message) LoadOrCreate() error {
	if Get(o) != nil {
		return Save(o)
	} else {
		return nil
	}
}
