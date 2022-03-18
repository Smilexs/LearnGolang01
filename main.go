package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
)

var (
	//参考文档：https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN#d65d109d
	//机器人Z的webhook: https://open.feishu.cn/open-apis/bot/v2/hook/0ea34cfa-88c8-4816-81d4-b9844fd08bf9
	robotWebhook = flag.String("url", "https://open.feishu.cn/open-apis/bot/v2/hook/0ea34cfa-88c8-4816-81d4-b9844fd08bf9", "robot webhook")
	//MessageType  = flag.String("type", "text", "Lark message type")
)

func main(){
	//fmt.Println("hello")

	//
	flag.Parse()

	//
	msg := "你好 小机器人"
	if err := SendMessage(msg); err != nil {
		fmt.Printf("发送飞书失败：%s\n", err)
		return
	}
	fmt.Printf("发送飞书成功：哈哈\n")
}


//请求的结构（json)
type ReqData struct {
	MsgType string `json:"msg_type"`
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
}

//返回的结构
type RespData struct {
	StatusCode    int
	StatusMessage string
}

/*
	请求的格式
	-X POST -H "Content-Type: application/json" \
	-d '{"msg_type":"text","content":{"text":"request example"}}' \
  	https://open.feishu.cn/open-apis/bot/v2/hook/xxxxxxxxxxxxxxxxx
*/
func SendMessage(msg string) error {
	//构建请求结构体
	reqData := &ReqData{
		MsgType: "text",
		//MsgType: *MessageType,
		Content: struct {
			Text string `json:"text"`
		}{
			Text: msg,//设置发送的内容
		},
	}

	//解析请求结构体成json格式
	data, err := json.Marshal(reqData)
	if err != nil {
		return errors.New("json marshal failed: " + err.Error())
	}

	//构建http请求
	req, err := http.NewRequest(http.MethodPost, *robotWebhook, bytes.NewBuffer(data))
	if err != nil {
		return errors.New("http new request failed: " + err.Error())
	}
	//设置请求头
	req.Header.Set("Content-Type", "application/json")

	//开始请求
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.New("http do failed" + err.Error())
	}

	//
	defer response.Body.Close()
	//解析返回的结果
	var respData RespData
	err = json.NewDecoder(response.Body).Decode(&respData)
	if err != nil {
		return errors.New("json NewDecoder failed" + err.Error())
	}
	//
	if respData.StatusCode != 0 {
		return errors.New("json NewDecoder failed: " + respData.StatusMessage)
	}

	return nil
}

