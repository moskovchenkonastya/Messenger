
package main

import (
	"fmt"
	//"io/ioutil"
	//"log"
	"net"
	//"time"
)

/*
//  Structure for profile of user
type profile struct {

	id		   int
	name	   string
	password   string
}

*/
func init() {

	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')

}


func main(){

	/*
		1. open a telnet connection with hostname and port only
		2. send a command inside this session or connection
		3. retrieve the output or http status (OPTIONAL)
		4. exit/quit session connection
	 */

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}


}

func handleConnection (conn *TCPConn) {

	var i int
	for {
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Forgive password")
		fmt.Println("4. Exit")

		fmt.Scanf("%d", &i)

		switch i {
		case 1:
			login()
		case 2:
			register()
		case 3:
			fogivePassword()
		case 4:
			return
		}
	}
}

func login() {

}

func  register(){

}

func fogivePassword() {

}
