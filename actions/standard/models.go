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
	}
	Response struct {
		Key    string `json:"key"`
		IsNull bool   `json:"isNull"`
	}
}

type IsArray struct {
	Request struct {
		Key string `json:"key"`
	}
	Response struct {
		Key     string `json:"key"`
		IsArray bool   `json:"isArray"`
	}
}

type Scheduler struct {
	Period    uint `json:"period"`
	Url       string `json:"url"`
	AccessKey string `json:"access_key"`
}
