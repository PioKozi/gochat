package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/piokozi/gochat/internal/entity"
)

func main() {
	port := os.Args[1]
	ent := entity.NewEntity("localhost:" + port)

	go ent.Listen()

	reader := bufio.NewReader(os.Stdin)
	for {
		msg, _ := reader.ReadString('\n')
		fields := strings.Fields(msg)
		switch fields[0] {
		case "/help":
			fmt.Println(helpDisplay)
		case "/quit":
			return
		case "/introduce":
			ent.Introduce(fields[1])
		case "/forget":
			ent.Forget(fields[1])
		default:
			ent.Send(msg)
		}
	}
}

const helpDisplay = `
---
/help                   displays this help
/introduce <socket>     allows to send messages on a socket, if it is present
/forget <socket>        opposite of /introduce
/quit                   exits
---
`
