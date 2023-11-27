package sensores

type Sensor interface {
}

type SensorP struct{
	Nome                string  `json:"Nome"`
	Tipo                string  `json:"Tipo"`
	Ip                  string  `json:"Ip"`
	Port                string  `json:"Port"`
	Power               bool    `json:"Power"`
	Leitura float64 `json:"Luminosidade"`
}
