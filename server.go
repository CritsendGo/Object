package csObject

type Server struct {
	Id        	int 		`json:"server_id,omitempty" oType:"int" oPrimary:"true" oOptional:"true"`
	Name   	 	string 		`json:"server_name" oType:"string" oPrimary:"false" oOptional:"false"`
	Datacenter	*Datacenter  `json:"datacenter_id" oType:"struct" oPrimary:"false" oOptional:"true"`
}
func (o *Server) Load() error{return Get(o)}
func (o *Server) Save() error{return Save(o)}
func (o *Server) LoadOrCreate() error{if Get(o)!=nil{return Save(o)}else{return nil}}