package apitest

import (
	"os"
	"testing"
)

func TestChannel(t *testing.T) {
	t.Run(
		"get child channel list", func(t *testing.T) {
			api, ctx := getBotApiCtx()
			list, err := api.Channels(ctx, os.Getenv("QQ_BOT_CHANNEL_ID"))
			if err != nil {
				t.Error(err)
			}
			for _, channel := range list {
				t.Logf("%+v", channel)
			}
			t.Logf(api.TraceID())
		},
	)
}
