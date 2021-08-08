package csObject

import "time"

type IpStatus struct {
	Id       int       `json:"ip_status_id?omitempty" oType:"int" oOptional:"true"`
	Name     string    `json:"ip_status_name" oType:"string" oOptional:"false"`
	Reverse  string    `json:"ip_reverse?omitempty" oType:"struct" oOptional:"true"`
	Create   time.Time `json:"ip_status_create,omitempty" oType:"time" oOptional:"true"`
	Update   time.Time `json:"ip_status_update" oType:"time" oOptional:"true"`
	Bookable bool      `json:"ip_status_bookable" oType:"bool" oOptional:"true"`
}

func (o *IpStatus) Load() error { return Get(o) }
func (o *IpStatus) Save() error { return Save(o) }
func (o *IpStatus) LoadOrCreate() error {
	if Get(o) != nil {
		return Save(o)
	} else {
		return nil
	}
}
