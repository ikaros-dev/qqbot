package apitest

import (
	"github.com/tencent-connect/botgo/dto"
	"log"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run(
		"get user info", func(t *testing.T) {
			api, ctx := getBotApiCtx()
			user, meError := api.Me(ctx)
			if meError != nil {
				log.Fatalln("调用 Me 接口失败, err = ", meError)
			}
			log.Println("user: ", user)
		},
	)
	t.Run("get channel list", func(t *testing.T) {
		api, ctx := getBotApiCtx()
		guilds, meGuildError := api.MeGuilds(ctx, &dto.GuildPager{})
		if meGuildError != nil {
			log.Fatalln("调用 MeGuild 接口失败, err = ", meGuildError)
		}
		for guild := range guilds {
			log.Println("Guild: ", guild)
		}
	})
}
