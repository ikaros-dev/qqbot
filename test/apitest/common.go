package apitest

import (
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
	"golang.org/x/net/context"
	"log"
	"os"
	"strconv"
	"time"
)

func getBotApiCtx() (openapi.OpenAPI, context.Context) {
	qqBotAppId := os.Getenv("QQ_BOT_APP_ID")
	accessToken := os.Getenv("QQ_BOT_ACCESS_TOKEN")

	appId, err := strconv.Atoi(qqBotAppId)
	if err != nil {
		log.Fatalln("转换QQ机器人ID到unit64格式失败, err = ", err)
	}

	token := token.BotToken(uint64(appId), accessToken)
	api := botgo.NewSandboxOpenAPI(token).WithTimeout(3 * time.Second)
	ctx := context.Background()

	return api, ctx
}
