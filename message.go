package main

import (
	"fmt"
	"log"
	"time"

	"go.nanomsg.org/mangos"
	"go.nanomsg.org/mangos/protocol/pair"
	"go.nanomsg.org/mangos/transport/all"
)

var (
	node string
)

func sendMsg(socket mangos.Socket, msg string) {
	log.Printf("Node %s sends %s\n", node, msg)
	err := socket.Send([]byte(msg))
	if err != nil {
		errLog(err.Error())
	}
}

func recMsg(socket mangos.Socket) string {
	bytes, err := socket.Recv()
	if err != nil {
		fmt.Println(err)
		errLog(err.Error())
	}
	msg := string(bytes)
	log.Printf("Node %s received %s\n", node, msg)
	return msg
}

func runNode(url string) {
	socket, err := pair.NewSocket()
	if err != nil {
		errLog(err.Error())
	}
	all.AddTransports(socket)
	socket.SetOption(mangos.OptionRecvDeadline, 10*time.Second)
	err = socket.Listen(url)
	if err != nil {
		errLog(err.Error())
		err = socket.Dial(url)
		if err != nil {
			errLog(err.Error())
		}
	}
	defer socket.Close()
	for i := 0; i < 3; i++ {
		sendMsg(socket, fmt.Sprintf("message %d from node %s.", i, node))
		_ = recMsg(socket)
		time.Sleep(1 * time.Second)
	}
	log.Printf("Node %s: Done.\n", node)
}
