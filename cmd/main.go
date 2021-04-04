package main

import (
	"bufio"
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
		if fields[0] == "/introduce" {
			ent.AddKnown(fields[1])
		} else {
			ent.Send(msg)
		}
	}
}
