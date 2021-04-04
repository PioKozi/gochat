package entity

import (
	"bufio"
	"fmt"
	"net"

	"github.com/piokozi/gochat/pkg/common"
)

// A node
type Entity struct {
	Socket        string   // the socket on which the entity resides
	KnownEntities []string // the entities which the entity knows
}

// Constructs a new node/entity and returns it
func NewEntity(socket string) Entity {
	return Entity{
		Socket:        socket,
		KnownEntities: make([]string, 0),
	}
}

// Adds socket to KnownEntities with value pubkey
// This can also be used to overwrite the known public key of a known entity
func (N *Entity) AddKnown(socket string) {
	N.KnownEntities = append(N.KnownEntities, socket)
}

// Sets up a listener on socket
// Contains an infinite loop!!!
func (N Entity) Listen() {
	listener, err := net.Listen("tcp", N.Socket)
	common.Errcheck(err)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		common.Errcheck(err)
		go func() {
			for {
				msg, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					break
				}
				fmt.Println(msg)
			}
		}()
	}
}

// Send msg to every socket in KnownEntities
func (N Entity) Send(msg string) {
	for _, socket := range N.KnownEntities {
		conn, err := net.Dial("tcp", socket)
		if err == nil {
			fmt.Fprintf(conn, "%s> %s", N.Socket, msg)
			conn.Close()
		}
	}
}
