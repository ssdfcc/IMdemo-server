package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"go-zero-chat/ent"
	"go-zero-chat/ent/group"
	"gopkg.in/fatih/set.v0"
	"net"
	"net/http"
	"sort"
	"sync"
	"time"
)

const (
	HeartBeatMaxTime = 1 * 30000 // 最大心跳时间
)

type Node struct {
	Conn          *websocket.Conn // 连接
	Addr          string          // 客户端地址
	FirstTime     int64           // 首次连接时间
	HeartBeatTime int64           // 心跳时间
	LoginTime     int64           // 登录时间
	DataQueue     chan []byte     // 消息
	GroupSets     set.Interface   //好友 群
}

type Msg struct {
	UserID   int    `json:"userId"`
	TargetID int    `json:"targetId"`
	Type     int    `json:"type"`
	Media    int    `json:"media"`
	Content  string `json:"content"`
}

// 映射关系
var clientMap = make(map[int]*Node)
var db *ent.Client
var red *redis.Client

// 读写锁
var rwLocker sync.RWMutex

func Chat(w http.ResponseWriter, r *http.Request, entClient *ent.Client, reds *redis.Client) {
	userId := r.Context().Value("userId").(int)
	db = entClient
	red = reds
	isValida := true
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isValida
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取conn
	currentTime := time.Now().Unix()
	node := &Node{
		Conn:          conn,
		Addr:          conn.RemoteAddr().String(),
		FirstTime:     currentTime,
		HeartBeatTime: currentTime,
		LoginTime:     currentTime,
		DataQueue:     make(chan []byte, 50),
		GroupSets:     set.New(set.ThreadSafe),
	}
	// 用户关系
	// userId跟node绑定并加锁
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()
	// 完成发送逻辑
	go sendProc(node)
	// 完成接收逻辑
	go recvProc(node)
	SetUserOnlineInfo(red, fmt.Sprintf("online_%d", userId), []byte(node.Addr))
	sendMsg(userId, []byte("欢迎进入聊天系统"))
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			fmt.Println("[ws] sendMsg >>>> msg:", string(data))
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		var msg Msg
		err = json.Unmarshal(data, &msg)
		if err != nil {
			fmt.Println(err)
		}
		// 心跳检查
		if msg.Type == -1 {
			node.HeartBeat(time.Now().Unix())
		} else {
			dispatch(data)
			broadMsg(data) // 广播消息
			fmt.Println("[ws] recvProc <<<< ", string(data))
		}
	}
}

var udpSendChan = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpSendChan <- data
}

func init() {
	go udpSendProc()
	go updRecvProc()
	fmt.Println("init goroutine")
}

// 完成udp数据发送协程
func udpSendProc() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 0, 255),
		Port: 3000,
	})
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case data := <-udpSendChan:
			fmt.Println("udpSendProc:", string(data))
			_, err := conn.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

// 完成udp数据接收协程
func updRecvProc() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
	}
	for {
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("updRecvProc:", string(buf[:n]))
		dispatch(buf[:n])
	}
}

// 后端调度逻辑处理
func dispatch(data []byte) {
	msg := Msg{}
	fmt.Println("dispatch:", string(data))
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case 1: // 私信
		fmt.Println("dispatch:", string(data))
		sendMsg(msg.TargetID, data)
	case 2: // 群发
		sendGroupMsg(msg.TargetID, data) // 发送的群ID，消息内容
		//case 3: // 广播
		//	sendAllMsg()
	}
}

func sendMsg(userId int, msg []byte) {
	rwLocker.RLock()
	node, ok := clientMap[userId]
	rwLocker.RUnlock()
	var jsonMsg Msg
	json.Unmarshal(msg, &jsonMsg)
	ctx := context.Background()
	r, err := red.Get(ctx, fmt.Sprintf("online_%d", userId)).Result()
	if err != nil {
		fmt.Println(err)
	}
	if r != "" {
		if ok {
			fmt.Println("sendMsg >>> userId:", userId, " msg:", string(msg))
			node.DataQueue <- msg
		}
	}
	if jsonMsg.Type == 1 {
		ids := []int{jsonMsg.UserID, jsonMsg.TargetID}
		sort.Ints(ids)
		red.RPush(ctx, fmt.Sprintf("msg_%v_%v", ids[0], ids[1]), string(msg))
	}
}

func sendGroupMsg(groupId int, msg []byte) {
	fmt.Println("sendMsg >>> groupId:", groupId, " msg:", string(msg))
	groupRelations, _ := db.Group.Query().Where(group.ID(groupId)).QueryGroupRelationGroup().All(context.Background())
	var userIds []int
	for i, _ := range groupRelations {
		userIds = append(userIds, groupRelations[i].UserID)
	}
	for i, _ := range userIds {
		sendMsg(userIds[i], msg)
	}
	ctx := context.Background()
	red.RPush(ctx, fmt.Sprintf("group_msg_%v", groupId), string(msg))
}

// 更新用户心跳
func (node *Node) HeartBeat(currentTime int64) {
	node.HeartBeatTime = currentTime
}

// 用户心跳是否超时
func (node *Node) IsHeartBeatTimeOut(currentTime int64) bool {
	if node.HeartBeatTime+HeartBeatMaxTime <= currentTime {
		return true
	}
	return false
}

// 清理超时连接
func ClearConnection(param interface{}) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("定时任务，清理超时连接，出现异常", r)
		}
	}()
	fmt.Println("定时任务，清理超时连接", param)
	var delNodeList []int
	for i, node := range clientMap {
		if node.IsHeartBeatTimeOut(time.Now().Unix()) {
			fmt.Println("清理超时连接", node.Addr)
			// 释放资源
			node.Conn.Close()
			// 删除映射关系
			delNodeList = append(delNodeList, i)
		}
	}
	// 加锁
	rwLocker.Lock()
	for _, v := range delNodeList {
		delete(clientMap, v)
	}
	rwLocker.Unlock()
	return true
}
