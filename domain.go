package csObject

import (
	"time"
)

//Domain    *Domain  `json:"message_id" oType:"struct" oPrimary:"false" oOptional:"false"`

type Domain struct {
	//Key
	Id    int    `json:"domain_id" oType:"int" oPrimary:"true" oOptional:"true"`
	Value string `json:"domain_name" oType:"int" oPrimary:"false" oOptional:"false"`
	//Data
	Key   string `json:"certificate_key" oType:"string" oPrimary:"false" oOptional:"true"`
	Pem   string `json:"certificate_pem" oType:"string" oPrimary:"false" oOptional:"true"`
	Chain string `json:"certificate_chain" oType:"string" oPrimary:"false" oOptional:"true"`
	// Timer
	Create time.Time `json:"certificate_create" oType:"time" oPrimary:"false" oOptional:"true"`
	Update time.Time `json:"certificate_update" oType:"time" oPrimary:"false" oOptional:"true"`
}

func (o *Domain) Load() error { return Get(o) }
func (o *Domain) Save() error { return Save(o) }
func (o *Domain) LoadOrCreate() error {
	if Get(o) != nil {
		return Save(o)
	} else {
		return nil
	}
}
