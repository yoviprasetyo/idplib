package sekolahdesain

// IDP struct
type IDP struct {
	Raw string
}

// SetRaw json to IDP
func (idp IDP) SetRaw(json string) {
	idp.Raw = json
}
