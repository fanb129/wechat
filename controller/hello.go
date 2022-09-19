package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yaotian/gowechat/mp/message"
	"wechat/conf"
	"wechat/service"
)

func Hello(c *gin.Context) {
	// 传入request和responseWriter
	msgHandler := conf.Mp.GetMsgHandler(c.Request, c.Writer)
	//设置接收消息的处理方法
	msgHandler.SetHandleMessageFunc(func(msg message.MixMessage) *message.Reply {

		switch msg.MsgType {
		case message.MsgTypeText:
			switch msg.Content {
			case "时间":
				return &message.Reply{message.MsgTypeText, message.NewText(service.GetTime())}
			case "天气":
				return &message.Reply{message.MsgTypeText, message.NewText(service.MyWeather())}
			}

		case message.MsgTypeImage:
			pic := message.NewImage(msg.MediaID)
			return &message.Reply{message.MsgTypeImage, pic}
		}

		text := message.NewText("范范最爱文艺\n" +
			"输入\"时间\"获取时间\n" +
			"输入\"天气\"获取天气" +
			"\n\n" + service.GetSentence())
		return &message.Reply{message.MsgTypeText, text}
	})

	//处理消息接收以及回复
	err := msgHandler.Handle()
	if err != nil {
		fmt.Println(err)
		return
	}
}
