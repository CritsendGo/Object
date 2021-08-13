package csObject

import (
	"time"
)

//Domain    *Domain  `json:"message_id" oType:"struct" oPrimary:"false" oOptional:"false"`

type Domain struct {
	//Key
	Id   int    `json:"domain_id" oType:"int" oPrimary:"true" oOptional:"true"`
	Name string `json:"domain_name" oType:"string" oPrimary:"false" oOptional:"false"`
	//Data
	// Timer
	Create time.Time `json:"domain_create" oType:"time" oPrimary:"false" oOptional:"true"`
	Update time.Time `json:"domain_update" oType:"time" oPrimary:"false" oOptional:"true"`
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
func (o *Domain) FromMap(a map[string]string) (n *Domain) { MapToObject(a, n); return n }
func (o *Domain) GetAll() (out []*Domain) {
	for _, row := range GetAll(o) {
		a := &Domain{}
		MapToObject(row, a)
		out = append(out, a)
	}
	return out
}
func (o *Domain) Clone() *Domain { out := &Domain{}; return out }
