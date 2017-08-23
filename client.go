package main

import (
	"fmt"
	//"io/ioutil"
	"log"
	"net"
	//"bufio"
	"flag"
	"time"
	"encoding/json"
	"bytes"
)

type tParamsLoginStruct struct {
	name     string
	password string
}

type CommandData struct{

	method string
}

func main(){

	log.Fatal(listenTCP())
}

func listenTCP() error {

	var listenAddr string
	// in command line: -listen-addr = 8080
	flag.StringVar(&listenAddr, "listen-addr", ":8080", "address to listen")
	flag.Parse()

	l, err := net.Listen("tcp", listenAddr)

	if err != nil {
		return fmt.Errorf("error listening on %q: %s", listenAddr, err)
	}
	defer func() {
		closeErr := l.Close()
		if closeErr != nil {
			log.Printf("can't close listen socket: %s", err)
		}
	}()

	fmt.Printf("Listening on %q\n", listenAddr)

	for {
		conn, err := l.Accept()
		//defer conn.Close()

		if err != nil {
			fmt.Printf("Error accepting connection %q: %s", listenAddr, err)
			time.Sleep(100 * time.Millisecond)
			continue
		}

		go handleConnection(conn)

	}
}

func handleConnection (conn net.Conn) {

	var i int
	for {
		fmt.Println("1 - Login")
		fmt.Println("2 - Register")
		fmt.Println("3 - Forgive password")
		fmt.Println("4 - Exit")

		fmt.Scanf("%d", &i)

		switch i {
		case 1:
			login()
		case 2:
			register()
		case 3:
			fogivePassword()
		case 4: break

		}
	}
}

func ParseJsonIntoStruct(c []byte, v interface{}) error {

	c = bytes.Trim(c, "")

	err := json.Unmarshal(c, v)
	if err != nil {
		return fmt.Errorf("Can't parse json (parseJsonIntoStruct): %s", err)

	}

	return err
}


func login() {

	fmt.Println("login")

	

	var params = &tParamsLoginStruct{}
	err := ParseJsonIntoStruct(CommandData.method, params)

	userName := params.name
	userPassword := params.password


}




func  register(){
	fmt.Println("register")
}

func fogivePassword() {
	fmt.Println("forgive Password")
}
