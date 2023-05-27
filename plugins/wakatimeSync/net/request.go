package net

import (
	"encoding/json"
	"fmt"
	"wakatimeSync/entity"

	"github.com/valyala/fasthttp"
)

type ApiInterface interface{}
type Api struct {
	ApiInterface
}

func (a *Api) Summary(req entity.SummariesRequest) (entity.SummariesResponse, error) {
	return client[entity.SummariesResponse]("POST", SUMMARY, req)
}

func (a *Api) Heartbeat(req entity.SummariesRequest) (entity.SummariesResponse, error) {
	return client[entity.SummariesResponse]("POST", SUMMARY, req)
}

func client[T interface{}, P interface{}](method string, path string, params P) (T, error) {
	var r T

	args := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(args)

	// 创建请求对象
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	// 发起GET请求
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(path)

	body, err := json.Marshal(params)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		return r, err
	}

	req.SetBody(body)

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod(method)

	resp := &fasthttp.Response{}

	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		return r, err
	}

	if err := json.Unmarshal(resp.Body(), &r); err != nil {
		return r, err
	}

	return r, nil
}
