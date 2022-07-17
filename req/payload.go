package req

type Payload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Payloads struct {
	Payloads []Payload `json:"data"`
}
