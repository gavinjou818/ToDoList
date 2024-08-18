package ctl

import (
	"github.com/gin-gonic/gin"
	"grpc_todolist/pkg/e"
)

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

func RespSuccess(ctx *gin.Context, data interface{}, code ...int) *Response {
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}

	// 打开这里看不到 token 数据
	//if data != nil {
	//	data = "操作成功"
	//}

	r := &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
	}
	return r
}

func RespError(ctx *gin.Context, err error, data string, code ...int) *Response {
	status := e.ERROR
	if code != nil {
		status = code[0]
	}

	r := &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
		Error:  err.Error(),
	}
	return r
}
