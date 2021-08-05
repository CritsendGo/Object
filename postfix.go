package csObject

type Postfix struct {
	Status    string `json:"postfix_status"`
	Size      string `json:"postfix_size"`
	Position  string `json:"postfix_position"`
	Updated   string `json:"postfix_updated"`
	Signature string `json:"postfix_signature"`
	Server    Server `json:"server_id"`
	Id        string `json:"postfix_id"`
	Path      string `json:"postfix_path"`
}
