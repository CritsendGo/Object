package csObject

import (
	"time"
)

type Datacenter struct {
	Id        	int 			`json:"datacenter_id,omitempty" oType:"int" oPrimary:"true" oOptional:"true"`
	Name   	 	string 			`json:"datacenter_name" oType:"string" oPrimary:"false" oOptional:"false"`
	Provider	Provider		`json:"provider_id" oType:"struct" oPrimary:"false" oOptional:"true"`
	Create 		time.Time		`json:"datacenter_create,omitempty" oType:"time" oPrimary:"false" oOptional:"true"`
	Update      time.Time 		`json:"datacenter_update" oType:"time" oPrimary:"false" oOptional:"true"`
}
func (o *Datacenter) Load() error{return Get(o)}
func (o *Datacenter) Save() error{return Save(o)}

