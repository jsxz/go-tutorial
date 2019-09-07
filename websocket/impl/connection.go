package impl

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
)

type Connection struct {
	wsConn    *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan []byte
	mutex     sync.Mutex
	isClosed  bool
}

func InitConnection(wsConn *websocket.Conn) (*Connection, error) {
	conn := &Connection{
		wsConn:  wsConn,
		inChan:  make(chan []byte, 1000),
		outChan: make(chan []byte, 1000),
	}
	//启动 读协程
	go conn.readLoop()
	//启动 写协程
	go conn.writeLoop()
	return conn, nil
}
func (conn *Connection) ReadMessage() ([]byte, error) {
	select {
	case data := <-conn.inChan:
		{
			return data, nil
		}
	case <-conn.closeChan:
		err := errors.New("connection is closed")
		return nil, err
	}

}
func (conn *Connection) WriteMessage(data []byte) error {
	select {
	case conn.outChan <- data:
		{
		}
	case <-conn.closeChan:
		err := errors.New("connection is closed")
		return err
	}
	return nil
}
func (conn *Connection) Close() {
	//线程安全的close
	conn.mutex.Lock()
	conn.wsConn.Close()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()

}

//内部实现
func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			goto ERR
		}
		select {
		case conn.inChan <- data:
			{
			}
		case <-conn.closeChan:
			goto ERR
		}

	}
ERR:
	conn.Close()
}
func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-conn.outChan:
			{
			}
		case <-conn.closeChan:
			goto ERR
		}

		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
