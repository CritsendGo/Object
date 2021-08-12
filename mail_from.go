package csObject

type MailFrom struct {
	Id     int    `json:"mailfrom_id" oType:"int" oPrimary:"true" oOptional:"true"`
	Name   string `json:"mailfrom_name" oType:"string" oPrimary:"false" oOptional:"false"`
	Domain int    `json:"domain_id" oType:"int" oPrimary:"false" oOptional:"false"`
}

func (o *MailFrom) Load() error { return Get(o) }
func (o *MailFrom) Save() error { return Save(o) }
func (o *MailFrom) LoadOrCreate() error {
	if Get(o) != nil {
		return Save(o)
	} else {
		return nil
	}

}
