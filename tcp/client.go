package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"unsafe"
)

//1. 支持多端
//2. 支持广播和私信
//3. 上线通知
//4. 下线通知
//5. 设置昵称
const (
	LOGIN  =  iota
	LOGINOUT
	SETNAME
	BROADCAST
	PRATIVE
	USERLIST
)

type ChanMsg struct {
	From ,To ,Msg string
}

//消息类型
type ClientMsg struct {
	To string `json:"to"`
	MsgType int `json:"msgType"`
	Msg string `json:"msg"`
	DataLen uintptr `json:"dataLen"`
}

func Help()  {
	fmt.Println("0 : user:list")
	fmt.Println( "1 : set your name")
	fmt.Println("2 : all your msg")
	fmt.Println("3 : anyone:your msg")
}
func main()  {
	Help()
	conn , err := net.Dial("tcp" , "127.0.0.1:8888")
	if err != nil {
		log.Panic("faild to dial " , err)
	}
	defer conn.Close()
	go handle_client_conn(conn)
	fmt.Println("welcome to chat")
	reader := bufio.NewReader(os.Stdin)
	for {
		msg , err := reader.ReadString('\n')
		if err != nil {
			log.Panic("read err" ,err)
		}
		msg  = strings.Trim(msg,"\r\n")
		if msg == "quit" {
			fmt.Println("bye bye")
			break
		}
		msgs := strings.Split(msg ,":")
		if len(msgs) == 2 {
			var clientMsg ClientMsg
			clientMsg.Msg = msgs[1]
			if msgs[0] == "set"{
				clientMsg.To = msgs[0]
				clientMsg.MsgType = SETNAME
				clientMsg.DataLen = unsafe.Sizeof(clientMsg)
			}else if msgs[0] == "all"{
				clientMsg.MsgType = BROADCAST
				clientMsg.To = "all"
				clientMsg.DataLen = unsafe.Sizeof(clientMsg)
			}else if msgs[0] == "user"{
				clientMsg.To = msgs[0]
				clientMsg.MsgType = USERLIST
				clientMsg.DataLen = unsafe.Sizeof(clientMsg)
			}else{
				clientMsg.To = msgs[0]
				clientMsg.MsgType = PRATIVE
				clientMsg.DataLen = unsafe.Sizeof(clientMsg)
			}
			data , _ := json.Marshal(clientMsg)
			_ , err := conn.Write(data)
			if err != nil {
				fmt.Println("send err" , err)
			}
		}else{
			fmt.Println("please stdin like  xx:xx")
		}
	}
}

func handle_client_conn(conn net.Conn){
	buf := make([]byte  ,256)
	for {
		n ,err := conn.Read(buf)
		if err != nil {
			log.Panic("read err" , err)
		}
		chanMsg := &ChanMsg{}
		err = json.Unmarshal(buf[:n] , chanMsg)
		if err != nil {
			fmt.Println("err json data")
			return
		}
		fmt.Println(chanMsg.From + "->" + chanMsg.Msg)
	}
}