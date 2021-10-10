package log

import(
	"fmt"
	ratatelogs "github.com/lestrrat/go-file-rotatelogs"
	"log"
	"time"
)
func init()  {
	path := "log/old.UTC."
	writer, err := ratatelogs.New(
		fmt.Sprintf("%s.%s", path, "%Y-%m-%d.%H:%M:%S"),
		ratatelogs.WithLinkName("log/current"),
		ratatelogs.WithMaxAge(time.Second * 10),
		ratatelogs.WithRotationTime(time.Second * 1),
		)
	log.SetOutput(writer)
	if err != nil {
		log.Fatalf("Failed to initialize log file %s", err)
	}
}

func TestLog()  {
	for i:=0; i<100; i++ {
	time.Sleep(time.Second * 1)
	log.Printf("Heool world")
	}
	fmt.Scanln()
}