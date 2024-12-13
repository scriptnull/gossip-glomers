package main

import (
	"encoding/json"
	"log"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()

	var messages []int
	var mu sync.RWMutex

	n.Handle("broadcast", func(msg maelstrom.Message) error {
		type request struct {
			Message int
		}
		var req request
		if err := json.Unmarshal(msg.Body, &req); err != nil {
			return err
		}

		mu.Lock()
		messages = append(messages, req.Message)
		mu.Unlock()

		return n.Reply(msg, map[string]string{"type": "broadcast_ok"})
	})

	n.Handle("read", func(msg maelstrom.Message) error {
		mu.RLock()
		defer mu.RUnlock()

		return n.Reply(msg, map[string]any{
			"type":     "read_ok",
			"messages": messages,
		})
	})

	n.Handle("topology", func(msg maelstrom.Message) error {
		return n.Reply(msg, map[string]string{"type": "topology_ok"})
	})

	if err := n.Run(); err != nil {
		log.Fatalln(err)
	}
}
