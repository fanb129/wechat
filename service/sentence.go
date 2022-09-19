package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Sentence struct {
	Content string `json:"content"`
	Note    string `json:"note"`
	Picture string `json:"picture"`
}

func GetSentence() string {
	resp, err := http.Get("http://open.iciba.com/dsapi/")
	//txt := message.NewText("获取每日一句失败")
	//badReply := &message.Reply{MsgType: message.MsgTypeText,MsgData: txt}
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	var s Sentence
	json.Unmarshal([]byte(body), &s)
	//	txt = message.NewText(string(body))
	//	return &message.Reply{MsgType: message.MsgTypeText,MsgData: txt}
	return s.Content + "\n" + s.Note
}
