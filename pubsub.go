package main

import (
	"fmt"
	"time"

	"go.nanomsg.org/mangos"
	"go.nanomsg.org/mangos/protocol/pub"
	"go.nanomsg.org/mangos/protocol/sub"
	"go.nanomsg.org/mangos/transport/all"
)

func server(url string) {
	var socket mangos.Socket
	var err error
	if socket, err = pub.NewSocket(); err != nil {
		errLog(err.Error())
	}
	all.AddTransports(socket)
	if err = socket.Listen(url); err != nil {
		errLog(err.Error())
	}
	for {
		d := time.Now().Format(time.ANSIC)
		fmt.Printf("SERVER: PUBLISHING DATE %s\n", d)
		if err = socket.Send([]byte(d)); err != nil {
			errLog(err.Error())
		}
		time.Sleep(time.Second)
	}
}

func client(url string, name string) {
	var socket mangos.Socket
	var err error
	var msg []byte

	if socket, err = sub.NewSocket(); err != nil {
		errLog(err.Error())
	}
	if err = socket.Dial(url); err != nil {
		errLog(err.Error())
	}
	err = socket.SetOption(mangos.OptionSubscribe, []byte(""))
	if err != nil {
		errLog(err.Error())
	}
	for {
		if msg, err = socket.Recv(); err != nil {
			errLog(err.Error())
		}
		fmt.Printf("CLIENT(%s): RECEIVED %s\n", name, string(msg))
	}

}
