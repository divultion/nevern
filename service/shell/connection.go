package shell

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Connection struct {
	conn                   net.Conn
	disconnected           bool
	mutex                  sync.RWMutex
	disconnectChan         <-chan struct{}
	output                 <-chan string
	outputReaderKillswitch chan<- struct{}
}

func NewConnection(conn net.Conn, chanLength int, bufferLength int) Connection {
	output := make(chan string, chanLength)
	outputReaderKillswitch := make(chan struct{})
	disconnectChan := make(chan struct{})

	go OutputReader(conn, bufferLength, disconnectChan, output, outputReaderKillswitch)

	return Connection{conn: conn, disconnected: false, mutex: sync.RWMutex{}, output: output, outputReaderKillswitch: outputReaderKillswitch, disconnectChan: disconnectChan}
}

func (conn *Connection) Write(data []byte) (int, error) {
	conn.mutex.Lock()
	defer conn.mutex.Unlock()

	return conn.conn.Write(data)
}

func (conn *Connection) TryRead() (string, bool) {
	conn.mutex.Lock()
	defer conn.mutex.Unlock()

	select {
	case val, ok := <-conn.output:
		return val, ok
	default:
		return "", false
	}
}

func (conn *Connection) GetRemoteAddress() string {
	conn.mutex.RLock()
	defer conn.mutex.RUnlock()

	return conn.conn.RemoteAddr().String()
}

func (conn *Connection) IsDisconnected() bool {
	conn.mutex.Lock()
	defer conn.mutex.Unlock()

	select {
	case <-conn.disconnectChan:
		conn.disconnected = true
		break
	default:
		break
	}

	return conn.disconnected
}

func (conn *Connection) IsConnected() bool {
	return !conn.IsDisconnected()
}

func (conn *Connection) Close() {
	conn.mutex.Lock()
	defer conn.mutex.Unlock()
	if !conn.disconnected {
		close(conn.outputReaderKillswitch)

		conn.conn.Close()
		conn.disconnected = true
	}
}

func OutputReader(reader io.Reader, bufferLength int, disconnectChan chan<- struct{}, sender chan<- string, killswitch <-chan struct{}) {
	defer close(sender)
	buffer := make([]byte, bufferLength)

	for {
		if IsKsDone(killswitch) {
			return
		}

		for i := range buffer {
			buffer[i] = 0
		}
		_, err := reader.Read(buffer)
		if IsKsDone(killswitch) {
			return
		}

		if err != nil {
			if err != io.EOF {
				fmt.Printf("Read error %s, stopping\n", err)
			}
			disconnectChan <- struct{}{}
			return
		}

		sender <- string(buffer)
	}
}

func IsKsDone(killswitch <-chan struct{}) bool {
	select {
	case <-killswitch:
		return true
	default:
		return false
	}
}
