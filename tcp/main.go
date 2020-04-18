package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
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

//消息通道
var msgCenter chan ChanMsg
//连接
var connList  map[string]net.Conn
var userChan map[string]string  //name->constring
var chanName map[string]string  //constring->name

func main()  {
	//初始化
	msgCenter = make(chan ChanMsg)
	connList = make(map[string]net.Conn)
	userChan = make(map[string]string)
	chanName = make(map[string]string)
	listener , err := net.Listen("tcp" ,"127.0.0.1:8888")
	if err != nil {
		log.Panic("Listen err " , err)
	}
	go msgCenterHandle()
	defer listener.Close()
	for {
		conn , err := listener.Accept()
		if err != nil {
			fmt.Println("connet error" , err)
			continue
		}
		//启动协程处理该链接的读写
		go handel_conn(conn)
	}
}
//处理链接的的消息
//1.记录connet放入map
//2.处理读写行为
func handel_conn(conn net.Conn)  {
	connString := conn.RemoteAddr().String()
	//最后断开连接
	defer loginout(conn)
	fmt.Println("connect :" +  connString)
	msg := ChanMsg{"system" , "all" ,"Login " + connString}
	msgCenter <- msg
	connList[connString] = conn
	msgContent := make([]byte , 255)
	for{
		line , err := conn.Read(msgContent)
		if err != nil || line < 0 {
			fmt.Println("conn err" , err)
			break
		}
		clientMsg := &ClientMsg{}
		fmt.Println("accpet msg" ,string(msgContent[:line]))
		err = json.Unmarshal(msgContent[:line] , clientMsg)
		if err != nil {
			fmt.Println("msg type err" , err)
			continue
		}
		//对比长度
		if clientMsg.DataLen  != unsafe.Sizeof(*clientMsg) {
			fmt.Println("msg len err" , err)
			continue
		}
		//发送
		msg := ChanMsg{connString , "all" ,""}
		switch clientMsg.MsgType {
		case SETNAME:
			msg.From = "system"
			if name, ok := chanName[connString] ; ok {
				msg.Msg  = connString + " has set name :" + name
			}else{
				userChan[clientMsg.Msg] = connString
				chanName[connString] = clientMsg.Msg
				msg.Msg  = connString + "  set name " + clientMsg.Msg +" success"
			}
		case BROADCAST:
			if userName ,ok := chanName[connString] ; ok{
				msg.From = userName
				msg.Msg  = clientMsg.Msg
			}else{
				//请先设置用户名
				msg.From = "system"
				msg.To = connString
				msg.Msg  = "please set name first"
			}
		case PRATIVE:
			//检查接收方
			constring, ok := userChan[clientMsg.To]
			if !ok {
				//接收方不存在 告知对方错误
				if msg.From != "system" {
					msg := ChanMsg{ "system", msg.From ,clientMsg.To + " no exists " }
					msgCenter <- msg
					continue
				}
			}else{
				msg.From  = chanName[connString]
				msg.To  = constring
				msg.Msg = clientMsg.Msg
			}
		case USERLIST:
			msg.From  = "system"
			msg.To = "me"
		}
		msgCenter <- msg
	}
}

func loginout(conn net.Conn)  {
	defer conn.Close()
	connString := conn.RemoteAddr().String()
	delete(connList ,connString)
	msg := ChanMsg{connString , "all" ,"LoginOut " + connString}
	msgCenter <- msg
}

func msgCenterHandle(){
	for {
		msg :=  <- msgCenter
		go sendMsgHandle(msg)
	}
}

func sendMsgHandle(msg ChanMsg){
	data , err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Fail to Marshal" ,err)
		return
	}
	if msg.To  == "all"{
		for k , v := range connList {
			if msg.From != k {
				v.Write(data)
			}
		}
	}else if msg.To  == "me"{
		con, ok := connList[msg.From]
		if !ok {
			return
		}
		userdata, err  := json.Marshal(userChan)
		if err != nil {
			fmt.Println("Fail to Marshal" ,err)
			return
		}
		msg.Msg = string(userdata)
		data , _ := json.Marshal(msg)
		con.Write(data)
	}else{
		con ,ok := connList[msg.To]
		if !ok {
			//接收方不存在 告知对方错误
			if msg.From != "system" {
				msg := ChanMsg{ "system", msg.From ,msg.To + " down " }
				msgCenter <- msg
			}
			return
		}
		con.Write(data)
	}
}

