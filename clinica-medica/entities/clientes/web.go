package clientes

import (
	"example/m/entities/atuadores"
	"example/m/entities/sensores"
)

type Web struct {
	Nome string
	Ip   string
	Port string
}

type Mensagem struct {
	SensorLuminosidade1 sensores.Luminosidade `json:"SensorLuminosidade1"`
	Atuadores           []atuadores.AtuadorP  `json:"Atuadores"`
	Sensores            []sensores.SensorP     `json:"Sensores"`
	MsgOperar MensagemOperar
}

type MensagemOperar struct {
	CMD    string
	Nome   string
	Tipo   string
	TipoOp string
	K int
}

type MensagemAtuador struct{

}
