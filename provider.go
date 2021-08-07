package csObject

import "time"

type Provider struct {
	status		objectStatus
	Id        	int 			`json:"provider_id,omitempty" oType:"int"`
	Name   	 	string 			`json:"provider_name" oType:"string"`
	IsRemovable bool			`json:"is_removable" oType:"string"`
	Create 		time.Time		`json:"provider_create,omitempty" oType:"time"`
	Update      time.Time 		`json:"provider_update" oType:"time"`
}
func NewProvider (name string) (i Provider){
	return Provider{status:oStatusNew,Name: name}
}