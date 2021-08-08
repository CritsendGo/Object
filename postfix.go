package csObject

import (
	"strconv"
	"time"
)

type Postfix struct {
	Id        int       `json:"postfix_id" oType:"int" oPrimary:"true" oOptional:"true"`
	Status    string    `json:"postfix_status" oType:"string" oPrimary:"false" oOptional:"true"`
	Size      int64     `json:"postfix_size" oType:"int" oPrimary:"false" oOptional:"true"`
	Position  int64     `json:"postfix_position" oType:"int" oPrimary:"false" oOptional:"true"`
	Updated   time.Time `json:"postfix_updated" oType:"time" oPrimary:"false" oOptional:"true"`
	Signature string    `json:"postfix_signature" oType:"string" oPrimary:"false" oOptional:"false"`
	Server    *Server   `json:"server_id" oType:"struct" oPrimary:"false" oOptional:"false"`
	Path      string    `json:"postfix_path" oType:"string" oPrimary:"false" oOptional:"true"`
}

func (o *Postfix) Load() error { return Get(o) }
func (o *Postfix) Save() error { return Save(o) }
func (o *Postfix) LoadOrCreate() error {
	if Get(o) != nil {
		return Save(o)
	} else {
		return nil
	}
}
func GetLastToRead(serverId int) (*Postfix, error) {
	o := &Postfix{}
	data, _ := apiClient.Get("postfix/?postfix_status=Present&server_id" + strconv.Itoa(serverId))
	e := MapToObject(data[0], o)
	return o, e
}
