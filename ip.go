package csObject

import "time"

type Ip struct {
	Id        	int 			`json:"ip_id?omitempty" oType:"int" oOptional:"true"`
	Value  	 	string 			`json:"ip_value" oType:"string" oOptional:"false"`
	Reverse		string			`json:"ip_reverse?omitempty" oType:"struct" oOptional:"true"`
	Create 		time.Time		`json:"ip_create,omitempty" oType:"time" oOptional:"true"`
	Update      time.Time 		`json:"ip_update" oType:"time" oOptional:"true"`
	Server 		*Server			`json:"server_id" oType:"int" oOptional:"true"`
}


func (o *Ip) Load() error{return Get(o)}
func (o *Ip) Save() error{return Save(o)}
func (o *Ip) LoadOrCreate() error{if Get(o)!=nil{return Save(o)}else{return nil}}