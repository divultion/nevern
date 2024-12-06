package shell

import (
	"fmt"

	"github.com/divultion/nevern/service/id"
)

type ConnectionNotFound struct {
	id id.Id
}

func NewConnectionNotFoundError(id id.Id) error {
	return &ConnectionNotFound{id: id}
}

func (error *ConnectionNotFound) Error() string {
	return fmt.Sprintf("Connection not found, invalid id: %s.", error.id.ToHex())
}

// Disconnected & removed from the Server
var PermanentDisconnect error = &PermanentDisconnectError{}

type PermanentDisconnectError struct{}

func (error *PermanentDisconnectError) Error() string {
	return "Peer Disconnected (0)"
}

// Disconnected but there's still messages to read
var TemporaryDisconnect error = &TemporaryDisconnectError{}

type TemporaryDisconnectError struct{}

func (error *TemporaryDisconnectError) Error() string {
	return "Peer Disconnected (Temporary)"
}
