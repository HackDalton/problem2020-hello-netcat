package main

import "net"

const message = `
 _   _            _    ____        _ _              
| | | | __ _  ___| | _|  _ \  __ _| | |_ ___  _ __  
| |_| |/ _' |/ __| |/ / | | |/ _' | | __/ _ \| '_ \ 
|  _  | (_| | (__|   <| |_| | (_| | | || (_) | | | |
|_| |_|\__,_|\___|_|\_\____/ \__,_|_|\__\___/|_| |_|

Welcome to HackDalton. Your first flag is:

hackDalton{h3ll0_w0rld_CBkJk20B-G}
`

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

func handleConnection(connection net.Conn) {
	connection.Write([]byte(message))
	connection.Close()
}
