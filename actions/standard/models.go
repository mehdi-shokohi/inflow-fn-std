package standardActions

type CastTo struct{
	CastTo string `json:"cast_to"`
	Key string `json:"field"`

}

type DBSettings struct{
	URI string `json:"uri"`
	User string `json:""`
}