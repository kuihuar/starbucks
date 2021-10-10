package httpptest

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"
)

func GetProducts1(w http.ResponseWriter, r *http.Request)  {
	query := r.URL.Query()
	filters, present := query["filters"]
	if !present || len(filters) == 0 {
		fmt.Println("filters no present")
	}
	w.WriteHeader(200)
	w.Write([]byte(strings.Join(filters, ",")))
}
func GetProducts(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	filters, present := r.Form["filters"]
	//query := r.URL.Query()
	//filters := query.Get("filters")
	if !present || len(filters) == 0 {
		fmt.Println("filters no present")
	}
	r.FormValue("filters")
	w.WriteHeader(200)
	w.Write([]byte(strings.Join(filters, ",")))
}

func DetectTimeout()  {
	client := &http.Client{
		Timeout:time.Nanosecond * 1,
	}
	_, err := client.Get("https://google.com")
	fmt.Println(os.IsTimeout(err))
}
func BasicAuth() error{
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", "https://localhost:8080/example", nil)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	req.SetBasicAuth("username", "password")
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()
	return nil
}
func HandBasicAuth(w http.ResponseWriter, r *http.Request){
	user,pass, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(401)
		return
	}
	if user != "username" {
		w.WriteHeader(401)
		return
	}
	if pass != "passwordd" {
		w.WriteHeader(401)
		return
	}
	w.WriteHeader(200)
	return
}
func UrlEncode(response http.ResponseWriter, request *http.Request){
	request.ParseForm()
	request.Form.Get("aaa")
	request.PostForm.Get("bbb")

}
func call(urlpath, method string) error {
	client := &http.Client{Timeout:time.Second * 10}
	data := url.Values{}
	data.Set("name", "john")
	data.Set("age", "18")
	data.Add("hobbies", "sport")
	data.Add("hobbies", "music")
	encodeData := data.Encode()
	req, err := http.NewRequest(method, urlpath, strings.NewReader(encodeData))
	if err != nil {
		fmt.Printf("got erro %s", err.Error())
	}
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("got error %s",err.Error())
	}
	defer response.Body.Close()
	return nil
}
func JpegServer (response http.ResponseWriter, request *http.Request){
	err := request.ParseMultipartForm(32 << 20) // 32左移10 就是 32K // 32左移20,就是 32M
	http.Redirect(response, request, "Http://www.baidu.com", http.StatusFound)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	file, fileHeader, err := request.FormFile("photo")
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	tempfile,err := os.Create("./"+ fileHeader.Filename)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(tempfile, file)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusOK)
	return
}
func JpegClient(urlPath, method string) error {
	client := &http.Client{Timeout:time.Second * 10}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormField("name")
	io.Copy(fw, strings.NewReader("jon"))
	fileWriter, err := writer.CreateFormFile("photo", "test.jpg")
	if err != nil {
		return err
	}
	file, err := os.Open("test.jpg")
	if err != nil {
		return err
	}
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return err
	}
	writer.Close()
	req, err := http.NewRequest(method, urlPath, bytes.NewReader(body.Bytes()))
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rep, _ := client.Do(req)
	if rep.StatusCode != http.StatusOK {
		log.Printf("Request failed with reponse code: %d", rep.StatusCode)
	}
	return nil
}
func Octet(urlPath, method string)error{
	client := &http.Client{Timeout:time.Second * 10}

	b, err := ioutil.ReadFile("photo.png")
	if err != nil {

	}
	req, err := http.NewRequest(method, urlPath, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/octet-stream")

	rep, err := client.Do(req)
	if rep.StatusCode != http.StatusOK {

	}
	return nil
}
func OctetServer(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {}
	tempFile, err := os.Create("./"+ "photo.png")
	defer tempFile.Close()
	tempFile.Write(data)
	w.WriteHeader(http.StatusOK)
	return
}
func Cookie(w http.ResponseWriter, r http.Request){
	cookie := &http.Cookie{
		Name:       "id",
		Value:      "abcd",
		Path:       "",
		Domain:     "",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     300,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	}
	http.SetCookie(w,cookie)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Doc get successful"))
	return
}
func CookieClient(w http.ResponseWriter, r *http.Request){
	c,_ := r.Cookie("id")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("get doc id successful"))
	fmt.Println(c)
	return
}

func CookieJarServer(w http.ResponseWriter, r *http.Request){

}
func CookieJarClient(urlPath, method string){
	jar, err := cookiejar.New(nil)
	if err != nil {}
	client := &http.Client{Jar: jar}

	req, err := http.NewRequest(method, urlPath, nil)
	if err != nil {}
	cookie := &http.Cookie{
		Name:"token",
		Value:"someToken",
		MaxAge: 300,
	}
	urlObj, _ := url.Parse(urlPath)
	client.Jar.SetCookies(urlObj, []*http.Cookie{cookie})
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	req,err = http.NewRequest("GET", "http://localhost:8080/doc/id", nil)
	if err != nil {}
	resp,_ = client.Do(req)
	defer resp.Body.Close()
}
