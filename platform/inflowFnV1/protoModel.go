package inflowV1

type ProtocolHeaderV1[T any] struct {
	Headers T `json:"_headers" bson:"headers"`
}

type ProtocolBodyV1[T any] struct {
	InlineParams map[string]interface{} `json:"_params"`
	Body         T                      `json:"_data" bson:"data"`
}

type Describe[T any] struct {
	Commands  []DescribeCommand `json:"commands"`
}
type DescribeCommand struct {
	Name   string `json:"name"`
	Short  string `json:"string"`
	Readme string `json:"readme"`
}
