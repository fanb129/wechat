package conf

import (
	"github.com/yaotian/gowechat"
	"github.com/yaotian/gowechat/wxcontext"
)

var config = wxcontext.Config{
	//微信公众平台，商户平台，需要填写的信息
	AppID:          "wx168be488480d1c62",
	AppSecret:      "5fa816b6e786e16eda72a636b958940e",
	Token:          "wenyi",
	EncodingAESKey: "e1UkjzTIx24iOMm1xbspUcforcdbwRK69SA312CbW6p",

	//以下是 mch商户平台需要的变量
	//SslCertFilePath string //证书公钥文件的路径
	//SslKeyFilePath  string //证书私钥文件的路径
	//SslCertContent  string //公钥证书的内容
	//SslKeyContent   string //私钥证书的内容
	//MchID           string //商户ID
	//MchAPIKey       string //商户平台设置的api key
}
var Mp *gowechat.MpMgr

func InitConf() (err error) {
	wc := gowechat.NewWechat(config)
	Mp, err = wc.MpMgr()
	if err != nil {
		return
	}

	return nil
}
