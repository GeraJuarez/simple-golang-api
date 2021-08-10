package model

type KeyVal struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Key struct {
	Key string `json:"key"`
}

type Values struct {
	Array []string `json:"values"`
}
