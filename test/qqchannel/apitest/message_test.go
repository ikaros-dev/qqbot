package apitest

import (
	"github.com/tencent-connect/botgo/dto"
	"os"
	"testing"
)

func TestMessage(t *testing.T) {
	t.Run(
		"SendHelloIkaros", func(t *testing.T) {
			api, ctx := getBotApiCtx()
			message, err := api.PostMessage(
				ctx, os.Getenv("QQ_BOT_CHAT_CHANNEL_ID"), &dto.MessageToCreate{
					Content: "Hello Ikaros",
				},
			)
			if err != nil {
				t.Error(err)
			}
			t.Logf("message : %v", message)
		},
	)
}
