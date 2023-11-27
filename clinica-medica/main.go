package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"example/m/entities/atuadores"
	"example/m/entities/clientes"
	"example/m/entities/sensores"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

var msgluminosidade []byte
var luminosidade sensores.Luminosidade
var msgWeb clientes.Mensagem

type clienteAtuador interface {
}

type HomeAssistant struct {
	Sock                net.Listener
	connections         net.Conn
	Atuadores           []atuadores.Atuador
	AtuadorLampada1     atuadores.Lampada
	SensorLuminosidade1 sensores.Luminosidade
	MsgWebAtual         clientes.Mensagem
}


func clientSensorLuminosidade() {
	fmt.Println("Cliente Sensor Luminosidade rodando...")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"fila-1",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
)

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"fila-1", // exchange
		false,
		nil,
)

failOnError(err, "Failed to declare a queue")


	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	var b sensores.SensorP
	msgWeb.Sensores = append(msgWeb.Sensores, b)
	k := len(msgWeb.Sensores) - 1

	go func() {
		for d := range msgs {
			json.Unmarshal(d.Body, &b)
			msgluminosidade = d.Body
			msgWeb.Sensores[k] = b

			log.Printf("Received  message: %s", &msgWeb.Sensores)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func clientSensorFumaca() {
	fmt.Println("Cliente Sensor Luminosidade rodando...")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"fila-2",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
)

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"fila-2", // exchange
		false,
		nil,
)

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	var b sensores.SensorP
	msgWeb.Sensores = append(msgWeb.Sensores, b)
	k := len(msgWeb.Sensores) - 1

	go func() {
		for d := range msgs {
			json.Unmarshal(d.Body, &b)
			msgluminosidade = d.Body
			msgWeb.Sensores[k] = b

			log.Printf("Received  message: %s", &msgWeb.Sensores)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func clientSensorTemperatura() {
	fmt.Println("Cliente Sensor Luminosidade rodando...")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"fila-3",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
)

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"fila-3", // exchange
		false,
		nil,
)

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	var b sensores.SensorP
	msgWeb.Sensores = append(msgWeb.Sensores, b)
	k := len(msgWeb.Sensores) - 1

	go func() {
		for d := range msgs {
			json.Unmarshal(d.Body, &b)
			msgluminosidade = d.Body
			msgWeb.Sensores[k] = b

			log.Printf("Received  message: %s", &msgWeb.Sensores)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}


func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (h *HomeAssistant) send(connection net.Conn, buffer []byte) {
	for {
		r, _ := json.Marshal(msgWeb)
		_, err := connection.Write(r)
		if err != nil {
			fmt.Println("Error write:", err.Error())
		}

		time.Sleep(time.Second)
	}

}

func (h *HomeAssistant) read(connection net.Conn) {
	buffer := make([]byte, 1024)
	defer connection.Close()
	for {
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}

		var m clientes.Mensagem
		json.Unmarshal(buffer[:mLen], &m)


		fmt.Println(m)

		switch m.MsgOperar.CMD{
		case "/OP":
			switch m.MsgOperar.TipoOp{
			case  "Ligar":
				r := h.Atuadores[m.MsgOperar.K].Ligar()
				msgWeb.Atuadores[m.MsgOperar.K] = h.Atuadores[m.MsgOperar.K].GetAtuador()
				fmt.Println(r)
			case  "Desligar":
				r := h.Atuadores[m.MsgOperar.K].Desligar()
				msgWeb.Atuadores[m.MsgOperar.K] = h.Atuadores[m.MsgOperar.K].GetAtuador()
				fmt.Println(r)
			default:
			}
		default:
		}
	}
}

func (h *HomeAssistant) serverSocket() {
	fmt.Println("Servidor Socket rodando...")

	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	h.Sock = server
	defer server.Close()

	go func() {

		for {
			connection, err := server.Accept()
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())
				os.Exit(1)
			}
			h.connections = connection
			
			go h.read(connection)

			h.send(connection, msgluminosidade)

		}
	}()

	for {

	}

}

func main() {

	ha := &HomeAssistant{}

	go clientSensorLuminosidade()
	go clientSensorFumaca()
	go clientSensorTemperatura()

	t1 := &atuadores.Lampada{}
	ha.Atuadores = append(ha.Atuadores, t1)
	ha.Atuadores[0].Conectar("localhost", "9000")
	var r = ha.Atuadores[0].Desligar()
	msgWeb.Atuadores = append(msgWeb.Atuadores, t1.La)

	t2 := &atuadores.ArCondicionado{}
	ha.Atuadores = append(ha.Atuadores, t2)
	ha.Atuadores[1].Conectar("localhost", "8000")
	var p = ha.Atuadores[1].Desligar()
	msgWeb.Atuadores = append(msgWeb.Atuadores, t2.AC)

	t3 := &atuadores.ControleDeIncendio{}
	ha.Atuadores = append(ha.Atuadores, t3)
	ha.Atuadores[2].Conectar("localhost", "6000")
	var e = ha.Atuadores[2].Desligar()
	msgWeb.Atuadores = append(msgWeb.Atuadores, t3.CDI)
	

	fmt.Println(r)
	fmt.Println(p)
	fmt.Println(e)
	fmt.Println(msgWeb)

	ha.serverSocket()

}
