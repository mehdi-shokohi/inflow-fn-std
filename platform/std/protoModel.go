package std

type ProtocolHeaderV1[T any] struct {
	Headers T `json:"_headers" bson:"headers"`
}

type ProtocolBodyV1[T any] struct {
	Body T `json:"_data" bson:"data"`
}

type Describe struct {
	Commands []DescribeCommand `json:"commands"`
}
type DescribeCommand struct {
	Name   string `json:"name"`
	Short  string `json:"string"`
	Readme string `json:"readme"`
}
