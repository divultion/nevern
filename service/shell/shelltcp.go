package shell

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"

	"github.com/divultion/nevern/service/id"
	"golang.org/x/exp/slices"
)

type ShellTCP struct {
	listener     net.Listener
	connections  map[id.Id]*Connection // MAX SIZE = chanLength * bufferLength
	chanLength   int
	bufferLength int
	mutex        sync.RWMutex
}

func ShellTCPListen(host string, port string, chanLength int, bufferLength int) (*ShellTCP, error) {

	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		return nil, err
	}

	shell := ShellTCP{listener: listener, connections: make(map[id.Id]*Connection), mutex: sync.RWMutex{}, chanLength: chanLength, bufferLength: bufferLength}
	return &shell, nil
}

func (shell *ShellTCP) ListenForConnections() {
	for {
		conn, err := shell.listener.Accept()
		if err != nil {
			log.Println("Failed accepting a connection request:", err)
			continue
		}

		go shell.handleConnection(conn)
		// TODO: Check for permanent disconnects
	}
}

func (shell *ShellTCP) checkPermanentDisconnect(id id.Id) bool {
	connected, err := shell.getConnectedByIdUnchecked(id) // GetConnectedById calls this function so it would be recursive call (
	if err != nil {
		return false
	}
	messagesAvailable, err := shell.getConnectionMessagesAvailableByIdUnchecked(id)
	if err != nil {
		return false
	}

	if !connected && messagesAvailable < 1 {
		shell.forgetById(id)
		return true
	}

	return false
}

func (shell *ShellTCP) handleConnection(conn net.Conn) {
	id, err := shell.newConnectionId()
	if err != nil {
		log.Fatal(err)
		return
	}

	shell.OpenNewConnection(id, conn)
}

func (shell *ShellTCP) newConnectionId() (id_ id.Id, err error) {
	for {
		id_, err = id.RandomId()
		if err != nil {
			fmt.Printf("Error generating id %s\n", err)
			continue
		}

		if !slices.Contains(Keys(shell.connections), id_) {
			break
		}
	}

	return id_, nil
}

func (shell *ShellTCP) GetAllConnectionIds() []id.Id {
	shell.mutex.RLock()
	defer shell.mutex.RUnlock()
	return Keys(shell.connections)
}

func (shell *ShellTCP) TryReadOutputById(id id.Id) (string, bool, error) {
	shell.mutex.Lock()
	defer shell.mutex.Unlock()
	conn := shell.getById(id)
	if conn == nil {
		return "", false, NewConnectionNotFoundError(id)
	}

	s, ok := conn.TryRead()
	if shell.checkPermanentDisconnect(id) {
		return s, ok, PermanentDisconnect
	}
	return s, ok, nil
}

// To connected only
func (shell *ShellTCP) WriteInputById(data string, id id.Id) (int, error) {
	connected, err := shell.GetConnectedById(id)
	if err != nil {
		return 0, err
	}

	if !connected {
		return 0, TemporaryDisconnect
	}

	shell.mutex.Lock()
	defer shell.mutex.Unlock()
	conn := shell.getById(id)
	if conn == nil {
		return 0, NewConnectionNotFoundError(id)
	}

	return conn.Write([]byte(strings.TrimSpace(data) + "\n"))
}

func (shell *ShellTCP) GetConnectionAddressById(id id.Id) (string, error) {
	shell.mutex.RLock()
	defer shell.mutex.RUnlock()
	conn := shell.getById(id)
	if conn == nil {
		return "", NewConnectionNotFoundError(id)
	}

	return conn.GetRemoteAddress(), nil
}

func (shell *ShellTCP) GetConnectionMessagesAvailableById(id id.Id) (int, error) {
	shell.mutex.RLock()
	defer shell.mutex.RUnlock()
	return shell.getConnectionMessagesAvailableByIdUnchecked(id)
}

func (shell *ShellTCP) getConnectionMessagesAvailableByIdUnchecked(id id.Id) (int, error) {
	conn := shell.getById(id)
	if conn == nil {
		return 0, NewConnectionNotFoundError(id)
	}

	return len(conn.output), nil
}

func (shell *ShellTCP) GetConnectedById(id id.Id) (bool, error) {
	shell.mutex.Lock()
	defer shell.mutex.Unlock()
	connected, err := shell.getConnectedByIdUnchecked(id)
	if shell.checkPermanentDisconnect(id) {
		return false, PermanentDisconnect
	}
	return connected, err
}

func (shell *ShellTCP) getConnectedByIdUnchecked(id id.Id) (bool, error) {
	conn := shell.getById(id)
	if conn == nil {
		return false, NewConnectionNotFoundError(id)
	}

	connected := conn.IsConnected()
	return connected, nil
}

func (shell *ShellTCP) OpenNewConnection(id id.Id, conn net.Conn) {
	shell.mutex.Lock()
	connection := NewConnection(conn, shell.chanLength, shell.bufferLength)
	shell.connections[id] = &connection
	shell.mutex.Unlock()
}

func (shell *ShellTCP) ForgetById(id id.Id) error {
	shell.mutex.Lock()
	defer shell.mutex.Unlock()
	return shell.forgetById(id)
}

func (shell *ShellTCP) DisconnectById(id id.Id) error {
	shell.mutex.Lock()
	defer shell.mutex.Unlock()
	return shell.disconnectById(id)
}

func (shell *ShellTCP) forgetById(id id.Id) error {
	conn := shell.getById(id)
	if conn == nil {
		return NewConnectionNotFoundError(id)
	}
	conn.Close()
	delete(shell.connections, id)
	return nil
}

func (shell *ShellTCP) disconnectById(id id.Id) error {
	conn := shell.getById(id)
	if conn == nil {
		return NewConnectionNotFoundError(id)
	}
	conn.Close()
	return nil
}

func (shell *ShellTCP) getById(id id.Id) *Connection {
	conn, ok := shell.connections[id]
	if !ok {
		return nil
	}

	return conn
}

func (shell *ShellTCP) Close() error {
	shell.mutex.Lock()
	defer shell.mutex.Unlock()
	for _, connection := range shell.connections {
		connection.Close()
	}

	return shell.listener.Close()
}

func Keys(m map[id.Id]*Connection) (keys []id.Id) {
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
