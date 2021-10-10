package httpptest

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"mvdan.cc/xurls"
)

var client http.Client
func init(){
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("got error while create cookie jar:%s", err.Error())
	}
	client = http.Client{
		Jar: jar,
	}
}
func SetCookieAtRequest(){
	cookie := &http.Cookie{
		Name: "token",
		Value: "sometoken",
		MaxAge: 300,
	}
	cookie2 := &http.Cookie{
		Name: "clicked",
		Value: "true",
		MaxAge: 300,
	}
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {}
	req.AddCookie(cookie)
	req.AddCookie(cookie2)
	resp, err := client.Do(req)
	if err != nil {}
	defer resp.Body.Close()

}
func SetCookieAtResponse(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name: "id",
		Value: "abcd",
		MaxAge:200,
	}
	http.SetCookie(w,cookie)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Doc Get Successful"))
	w.Header().Set("Content-Type", "application/text")
	w.Header().Set("Content-Type", "application/text")
	w.Header().Set("Content-Type", "application/octet-stream")
	token, err := r.Cookie("token")
	if err != nil {

	}

	fmt.Println(token)

	return
}

func usrlparse()  {
	inputUrl:=" fafdf  fdfhttps://test:abcd@example.com:8000/path/user?type=advance&compact=false#historyf dfds"
	schema, _ := url.Parse(inputUrl)
	query := schema.Query()
	fmt.Println(query)
	fmt.Println(schema.Scheme,schema.User,schema.String())
	xurlsStrict,err :=xurls.StrictMatchingScheme("https")
	if err != nil {}
	//input = inputUrl
	//hostName := schema.Scheme + "://" + schema.Host
	output := xurlsStrict.FindString(inputUrl)
	fmt.Println(output)
}
