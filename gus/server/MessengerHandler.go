package server

import (
"context"

"github.com/Dariusrussellkish/gus/gus/message"
)

type MessageHandler struct {}


func (mh MessageHandler) SendMessage(ctx context.Context, msg *message.Message) (*message.MessageAck, error) {
	return &message.MessageAck{
		ID: msg.ID,
		Frm: "Client",
	}, nil
}

