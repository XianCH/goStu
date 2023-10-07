package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ServerRequest struct {
	r *http.Request
}

type ServerResponse struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	ServerRequest := &ServerRequest{r: r}
	if !ServerRequest.CheckRequestMethod() {
		return
	}

	var user User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println("解析错误", err)
		return
	}

	fmt.Printf("Received JSON data: %+v\n", user)

	// 返回响应 以ServerRequest结构体为例，返回ServerResponse，生成json格式返回
	response := &ServerResponse{
		Code: "200",
		Msg:  "OK",
		Data: user, // 这里将 user 数据返回
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		fmt.Println("编码错误", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (Request *ServerRequest) CheckRequestMethod() bool {
	if Request.r.Method != http.MethodPost {
		fmt.Println("请求方法错误")
		return false
	}
	return true
}
