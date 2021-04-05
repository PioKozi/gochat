package entity

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"github.com/piokozi/gochat/pkg/common"
)

// A node
type Entity struct {
	Nick          string   // the nick/name given to this node/user
	Socket        string   // the socket on which the entity resides
	KnownEntities []string // the entities which the entity knows
}

// Constructs a new node/entity and returns it
func NewEntity(socket string) Entity {
	return Entity{
		Nick:          socket,
		Socket:        socket,
		KnownEntities: make([]string, 0),
	}
}

// Adds socket to N.KnownEntities
func (N *Entity) Introduce(socket string) {
	conn, err := net.DialTimeout("tcp", socket, time.Millisecond*5)
	if err != nil {
		fmt.Println("*** Socket is not listening ***")
	} else {
		N.KnownEntities = append(N.KnownEntities, socket)
		conn.Close()
	}
}

// Changes the nick of the entity and informs all other known entities
func (N *Entity) ChangeNick(newnick string) {
	N.SendAll(fmt.Sprintf("*** %s is now known as %s ***\n", N.Nick, newnick))
	N.Nick = newnick
}

func (N *Entity) Forget(socket string) {
	for i, v := range N.KnownEntities {
		if v == socket {
			N.KnownEntities[i] = N.KnownEntities[len(N.KnownEntities)-1]
			N.KnownEntities = N.KnownEntities[:len(N.KnownEntities)-1]
			return
		}
	}
	fmt.Println("*** No such known socket ***")
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
func (N Entity) SendAll(msg string) {
	for _, socket := range N.KnownEntities {
		conn, err := net.Dial("tcp", socket)
		if err == nil {
			fmt.Fprintf(conn, "%s> %s", N.Nick, msg)
			conn.Close()
		}
	}
}
