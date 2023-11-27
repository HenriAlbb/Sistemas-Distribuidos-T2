package main

import (
	"encoding/json"
	"example/m/entities/clientes"
	"example/m/entities/sensores"
	//"example/m/entities/atuadores"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

var luminosidade sensores.Luminosidade
var mensagem string
var msg clientes.Mensagem
var conn net.Conn

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
		case "GET":		
			 //http.ServeFile(w, r, "form.html")
		case "POST":
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			btn := r.FormValue("button")

			fmt.Println(btn)

			var msgOp clientes.Mensagem

			msgOp.MsgOperar.CMD = "/OP"

			res := strings.Split(btn, ";")
			k, _ := strconv.Atoi(res[1])

			fmt.Println(res)
			fmt.Println(k)
			msgOp.MsgOperar.TipoOp = res[0]
			msgOp.MsgOperar.K = k

			msgOp.MsgOperar.Nome = msg.Atuadores[k].GetNome()
			msgOp.MsgOperar.Tipo = msg.Atuadores[k].GetTipo()

			s, _ := json.Marshal(msgOp)
			send(conn, []byte(s))
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	
	fmt.Fprintf(w, "<table>")
	fmt.Fprintf(w, "<tr>")
	fmt.Fprintf(w, "<th>Nome</th>")
	fmt.Fprintf(w, "<th>Tipo</th>")
	fmt.Fprintf(w, "<th>Power</th>")
	fmt.Fprintf(w, "<th>t3</th>")
	fmt.Fprintf(w, "</tr>")
	for i,a := range msg.Atuadores{
		fmt.Fprintf(w, "<tr>")
		fmt.Fprintf(w, "<td>%s</td>", a.GetNome())
		fmt.Fprintf(w, "<td>%s</td>", a.GetTipo())
		fmt.Fprintf(w, "<td>%s</td>", a.Power)
		fmt.Fprintf(w, "<form method='POST'>")
		fmt.Fprintf(w, "<td><button type='submit' name='button' value='Ligar;"+strconv.Itoa(i)+"'>Ligar</button></td>")
		fmt.Fprintf(w, "<td><button type='submit' name='button' value='Desligar;"+strconv.Itoa(i)+"'>Desligar</button></td>")
		fmt.Fprintf(w, "</form>")
		fmt.Fprintf(w, "</tr>")
	}
	fmt.Fprintf(w, "</table>")

	fmt.Fprintf(w, "<table>")
	fmt.Fprintf(w, "<tr>")
	fmt.Fprintf(w, "<th>Nome</th>")
	fmt.Fprintf(w, "<th>Tipo</th>")
	fmt.Fprintf(w, "<th>Leitura</th>")
	fmt.Fprintf(w, "</tr>")
	for _,s := range msg.Sensores{
		fmt.Fprintf(w, "<tr>")
		fmt.Fprintf(w, "<td>%s</td>", s.Nome)
		fmt.Fprintf(w, "<td>%s</td>", s.Tipo)
		fmt.Fprintf(w, "<td>%s</td>", s.Leitura)
		fmt.Fprintf(w, "</tr>")
	}
	fmt.Fprintf(w, "</table>")

}

func send(connection net.Conn, msg []byte) {
	fmt.Println("send")
	_, err := connection.Write(msg)
	if err != nil {
		fmt.Println("Error write:", err.Error())
	}

}

func read(connection net.Conn) {
	buffer := make([]byte, 1024)
	defer connection.Close()
	for {
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}

		json.Unmarshal(buffer[:mLen], &msg)
	
		fmt.Println(msg.Sensores)
	}
}

func main() {
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error connection:", err.Error())
	}
	conn = connection
	defer connection.Close()

	go read(connection)
	
	http.HandleFunc("/", helloHandler)

	fmt.Printf("Starting server at por 7000\n")
	if err := http.ListenAndServe(":7000", nil); err != nil {
		log.Fatal(err)
	}
}
