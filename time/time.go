package time

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func TestTime()  {
	//date := time.Date(1,2,3,5,0,0,0,nil)
	////10 000 000 000
	//fmt.Printf("%+v",date)

	now := time.Now()
	loc, _ := time.LoadLocation("UTC")
	fmt.Printf("UTC Time: %s\n", now.In(loc))
	loc, _ = time.LoadLocation("Asia/Shanghai")
	fmt.Printf("Aisa Time: %s\n", now.In(loc))
	//t:= time.Now()
	//t := time.Now().UnixNano() * int64(time.Nanosecond) / int64(time.Millisecond) / int64(time.Microsecond)


}
func downloadFile(url, filename string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.New("Received non 200 response code")
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}