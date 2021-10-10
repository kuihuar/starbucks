package os

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	user "os/user"
	_ "github.com/joho/godotenv"
)

func TestOS()  {
	//exec.Command("/bin/sh", "sample.sh").Output()
	//
	//_ = os.Setenv("a", "b")
	//fmt.Println(os.Environ())
	//os.Clearenv()
	//fmt.Println(os.Environ())
	//os.Getenv("a")
	//os.Unsetenv("a")
	//val, present := os.LookupEnv("a")
	//fmt.Printf("a=%t,and b=%t",val,present)
	//fmt.Print(runtime.GOOS)
	//fmt.Println(os.Hostname())
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Username: %s\n", user.Username)
	fmt.Printf("HomeDirectory: %s\n", user.HomeDir)
	err = godotenv.Load("env/local.env")
	if err != nil {
		log.Fatalf("Some error ocured, Err: %s\n", err)
	}
	fmt.Println(os.Getenv("STACK"))
	os.Exit(0)
}
