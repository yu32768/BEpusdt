package task

import (
	"context"
	"github.com/v03413/bepusdt/app/conf"
	"time"
)

func ethInit() {
	register(task{
		ctx: context.WithValue(context.Background(), "cfg", evmCfg{
			Endpoint: conf.GetEthereumRpcEndpoint(),
			Type:     conf.Ethereum,
			Decimals: decimals{
				Usdt:   conf.UsdtEthDecimals,
				Native: -18, // ethereum.eth 小数位数
			},
			Block: block{
				InitStartOffset: -100,
				ConfirmedOffset: 12,
			},
		}),
		duration: time.Second * 12,
		callback: evmBlockRoll,
	})
}
