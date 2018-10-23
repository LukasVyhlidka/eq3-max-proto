package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"regexp"
	"time"

	"github.com/LukasVyhlidka/eq3-max-proto/model"
	"github.com/LukasVyhlidka/eq3-max-proto/parser"
)

type MaxClient struct {
	host string
	port int
}

func ObtainInitialMessages(url string) ([]model.Message, error) {
	conn, err := net.Dial("tcp", url)

	if err != nil {
		log.Panic("Connection error")
		return nil, err
	}

	time.Sleep(1000 * time.Millisecond)

	fmt.Fprintf(conn, "q:\r\n")

	messages := []model.Message{}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		log.Println("cube line: " + line)

		message, err := processLine(line)
		if err != nil {
			log.Fatal("Cube line is not valid.")
		} else {
			messages = append(messages, message)
		}
	}

	conn.Close()

	return messages, nil
}

var lineRegexp *regexp.Regexp = regexp.MustCompile(`^(\w):.+$`)

func processLine(line string) (model.Message, error) {
	msgParts := lineRegexp.FindStringSubmatch(line)
	if msgParts == nil {
		return model.MMessage{}, parser.Error("Wrong max cube line.")
	}

	msgType := msgParts[1]
	switch msgType {
	case "L":
		return parser.ParseLMessage(line)
	case "M":
		return parser.ParseMMessage(line)
	default:
		return model.GenericMessage{Message: line, MessageType: msgType}, nil
	}

}
