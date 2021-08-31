package event

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/pusher/pusher-http-go"
	uuid "github.com/satori/go.uuid"
)

const chunkSize = 8000

type Client interface {
	Publish(channel string, eventName string, payload interface{}) error
	AuthenticatePrivateChannel(params []byte) (*[]byte, error)
}

type event struct {
	client pusher.Client
}

func NewClient(config Config) Client {
	return &event{
		pusher.Client{
			AppID:   config.AppID,
			Key:     config.Key,
			Secret:  config.Secret,
			Cluster: config.Cluster,
		},
	}
}

func (e *event) Publish(channel string, eventName string, payload interface{}) error {
	if channel == "" {
		return errors.New("channel cannot be empty")
	}

	if eventName == "" {
		return errors.New("event name cannot be empty")
	}

	bytePayload, err := json.Marshal(payload)

	if err != nil {
		return errors.WithStack(err)
	}

	if len(string(bytePayload)) > chunkSize {
		var events []pusher.Event

		id := uuid.NewV4().String()
		chunkedString := chunkString(string(bytePayload), chunkSize)
		lastElement := chunkedString[len(chunkedString)-1]

		for index, ch := range chunkedString {
			chunkPayload := map[string]interface{}{
				"id":          id,
				"index":       index,
				"chunkString": ch,
				"final":       lastElement == ch,
			}

			events = append(events, pusher.Event{
				Channel: channel,
				Name:    "chunked-" + eventName,
				Data:    chunkPayload,
			})
		}

		// we can send max 10 items in the batch so we send them per 10 by chunking the event slice
		for _, batch := range chunkEvents(events, 10) {
			if err := e.client.TriggerBatch(batch); err != nil {
				return errors.WithStack(err)
			}
		}

		return nil
	}

	return errors.WithStack(e.client.Trigger(channel, eventName, payload))
}

func (e *event) AuthenticatePrivateChannel(params []byte) (*[]byte, error) {
	response, err := e.client.AuthenticatePrivateChannel(params)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &response, nil
}

func chunkString(s string, chunkSize int) []string {
	if chunkSize >= len(s) {
		return []string{s}
	}

	var chunks []string
	chunk := make([]rune, chunkSize)
	l := 0

	for _, r := range s {
		chunk[l] = r
		l++
		if l == chunkSize {
			chunks = append(chunks, string(chunk))
			l = 0
		}
	}

	if l > 0 {
		chunks = append(chunks, string(chunk[:l]))
	}

	return chunks
}

func chunkEvents(items []pusher.Event, chunkSize int) (chunks [][]pusher.Event) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}

	return append(chunks, items)
}
