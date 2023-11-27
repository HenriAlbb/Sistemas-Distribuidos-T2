package sensores

type Luminosidade struct {
	Nome         string  `json:"Nome"`
	Tipo         string  `json:"Tipo"`
	Power        string  `json:"Power"`
	Ip           string  `json:"Ip"`
	Port         string  `json:"Port"`
	Leitura float64 `json:"Luminosidade"`
}
