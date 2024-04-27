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

type PolicyCron struct {
	Minutes uint `json:"period"`
	Url Url `json:"url"`
	DocId string `json:"docId"`
	PolicyId string `json:"policyId"`
}
type Url struct {
	Url       string `json:"url"`
	AccessKey string `json:"access_key"`
	Token     string `json:"token"`
}
type HTTPRequest struct {
	URL    string      `json:"call_url"`
	Body   interface{} `json:"with_body"`
	Method string      `json:"method"`
}
