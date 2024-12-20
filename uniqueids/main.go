package main

import (
	"encoding/json"
	"fmt"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("generate", func(msg maelstrom.Message) error {
		var body map[string]any
		err := json.Unmarshal(msg.Body, &body)
		if err != nil {
			return err
		}

		body["type"] = "generate_ok"
		body["id"] = fmt.Sprintf("%s-%f", msg.Dest, body["msg_id"])

		return n.Reply(msg, body)
	})
	if err := n.Run(); err != nil {
		log.Fatalln(err)
	}
}
