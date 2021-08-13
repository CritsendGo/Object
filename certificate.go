package csObject

import (
	"time"
)

//

type Certificate struct {
	//Key
	Id     int     `json:"certificate_id" oType:"int" oPrimary:"true" oOptional:"true"`
	Domain *Domain `json:"domain_id" oType:"struct" oPrimary:"false" oOptional:"false"`
	//Data
	Key   string `json:"certificate_key" oType:"string" oPrimary:"false" oOptional:"true"`
	Pem   string `json:"certificate_pem" oType:"string" oPrimary:"false" oOptional:"true"`
	Chain string `json:"certificate_chain" oType:"string" oPrimary:"false" oOptional:"true"`
	// Timer
	Create time.Time `json:"certificate_create" oType:"time" oPrimary:"false" oOptional:"true"`
	Update time.Time `json:"certificate_update" oType:"time" oPrimary:"false" oOptional:"true"`
}

func (o *Certificate) Load() error { return Get(o) }
func (o *Certificate) Save() error { return Save(o) }
func (o *Certificate) LoadOrCreate() error {
	if Get(o) != nil {
		return Save(o)
	} else {
		return nil
	}
}
func (o *Certificate) FromMap(a map[string]string) (n *Certificate) { MapToObject(a, n); return n }
func (o *Certificate) GetAll() (out []*Certificate) {
	for _, row := range GetAll(o) {
		a := &Certificate{}
		MapToObject(row, a)
		out = append(out, a)
	}
	return out
}
func (o *Certificate) Clone() *Certificate { out := &Certificate{}; return out }
