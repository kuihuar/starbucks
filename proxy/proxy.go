package proxy

import "fmt"

type server interface {
	handleRequest(string, string)(int, string)
}

type nginx struct {
	application *application
	maxAllowedRequest int
	rateLimiter map[string]int
}
func newNginxServer() *nginx {
	return &nginx{
		application:       &application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}
func (n *nginx)handleRequest(url, method string)(int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.application.handleRequest(url, method)
}
func (n *nginx)checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] +1
	return true
}
type application struct {

}

func (a *application)handleRequest(url, method string) (int, string) {
	if url == "/apps/status" && method == "GET" {
		return 200, "Ok"
	}
	if url == "/create/user" && method == "POST" {
		return 201, "User Created"
	}
	return 404, "Not Found"
}

func TestProxy (){
	nginxServer := newNginxServer()
	appStatusURL := "/apps/status"
	createUserURL := "/create/user"
	httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("URL: %s\n HttpCode:%d\nBody:%s\n",appStatusURL,httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("URL: %s\n HttpCode:%d\nBody:%s\n",appStatusURL,httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("URL: %s\n HttpCode:%d\nBody:%s\n",appStatusURL,httpCode, body)

	httpCode, body = nginxServer.handleRequest(createUserURL, "POST")
	fmt.Printf("URL: %s\n HttpCode:%d\nBody:%s\n",createUserURL,httpCode, body)
	httpCode, body = nginxServer.handleRequest(createUserURL, "GET")
	fmt.Printf("URL: %s\n HttpCode:%d\nBody:%s\n",createUserURL,httpCode, body)
	httpCode, body = nginxServer.handleRequest(createUserURL, "POST")
	fmt.Printf("URL: %s\n HttpCode:%d\nBody:%s\n",createUserURL,httpCode, body)
}