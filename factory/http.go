package factory

import (
	"net/http"
	"net/http/httptest"
)
//多个工厂函数返回一个接口的不同实现
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}
func NewHTTPClient () Doer {
	return &http.Client{}
}

type mockHTTPClient struct{}

func(*mockHTTPClient) Do(req *http.Request) (*http.Response, error){
	res := httptest.NewRecorder()
	return res.Result(), nil
}

func NewMockHTTPClient() Doer{
	return &mockHTTPClient{}
}