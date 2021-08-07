package easyJSONPackage

type Example struct {
	Mock int
}

//easyjson:json
type JSONData struct {
	Data []string
}

//easyjson:json
type EasyJSONStruct struct {
	a int    `json:"a"`
	b string `json:"b"`
	c bool   `json:"c"`
	d interface{}
	e []interface{}
	f []int
	g []string
	h Example
	i []Example
}
