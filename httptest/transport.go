package httpptest

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

func TestTransport(){
	cachedTransport := newTransport()
	client := &http.Client{
		Transport: cachedTransport,
		Timeout: time.Second * 5,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	cacheClearTicker := time.NewTicker(time.Second * 5)
	reqTicker := time.NewTicker(time.Second * 1)
	terminateChannel := make(chan os.Signal, 1)
	signal.Notify(terminateChannel, syscall.SIGTERM, syscall.SIGHUP)
	req , err := http.NewRequest(http.MethodGet, "http://localhost:8000", strings.NewReader(""))
	if err != nil {}
	for{
		select{
		case <-cacheClearTicker.C:
			cachedTransport.Clear()
		case <-terminateChannel:
			cacheClearTicker.Stop()
			reqTicker.Stop()
			return
		case <-reqTicker.C:
			resp, err := client.Do(req)
			if err != nil {continue}
			buf,err := ioutil.ReadAll(resp.Body)
			if err != nil {continue}
			fmt.Printf("response is %s",string(buf))
		}
	}

}


type cacheTransport struct {
	data map[string]string
	mu sync.RWMutex
	originalTransport http.RoundTripper
}
func newTransport() *cacheTransport {
	return &cacheTransport{
		data:              make(map[string]string),
		originalTransport: http.DefaultTransport,
	}
}
func cachekey(r *http.Request) string {
	return r.URL.String()
}
func (c *cacheTransport)Set(r *http.Request, value string)  {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[cachekey(r)]=value
}
func (c *cacheTransport)Get(r *http.Request) (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if val, ok := c.data[cachekey(r)]; ok {
		return val,nil
	}
	return "", errors.New("key not found int cache")
}
func (c *cacheTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	if val, err := c.Get(r);err == nil {
		return cacheResponse([]byte(val), r)
	}
	resp, err := c.originalTransport.RoundTrip(r)
	if err != nil {}
	buf, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}
	c.Set(r, string(buf))
	fmt.Println("Fetching the data form the real source")
	return resp, nil
}
func (c *cacheTransport) Clear (){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]string)
	return
}
func cacheResponse(b []byte, r *http.Request)(*http.Response, error)  {
	buf := bytes.NewBuffer(b)
	return http.ReadResponse(bufio.NewReader(buf),r)
}