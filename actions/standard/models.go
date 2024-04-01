package standardActions

type CastTo struct {
	CastTo string   `json:"cast_to"`
	Keys   []string `json:"fields"`
}

type DBSettings struct {
	URI  string `json:"uri"`
	User string `json:""`
}

type Keys struct {
	Keys []string `json:"keys"`
}

type DeleteKeys struct {
	Keys []string `json:"delete_key"`
}

type IsNull struct {
	Request struct {
		Key string `json:"key"`
	} `intent:"req"`
	Response struct {
		Key    string `json:"key"`
		IsNull bool   `json:"isNull"`
	} `intent:"res"`
}

type IsArray struct {
	Request struct {
		Key string `json:"key"`
	} `intent:"req"`
	Response struct {
		Key     string `json:"key"`
		IsArray bool   `json:"isArray"`
	} `intent:"res"`
}

type Scheduler struct {
	Period    uint   `json:"period"`
	Url       string `json:"url"`
	AccessKey string `json:"access_key"`
}

type HTTPRequest struct {
	URL    string `json:"call_url"`
	Body   interface{} `json:"with_body"`
	Method string `json:"method"`
}
