package wraperrors

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

func ReadFile(path string) ([]byte, error) {
	f ,err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open file failed")
	}
	defer f.Close()
	return nil, err
}

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ReadFile(filepath.Join(home, ".settings.xml"))
	return config, errors.WithMessage(err, "could not read config")

}
func Test()  {
	_, err := ReadConfig()
	if err != nil {
		// Cause 根因
		fmt.Printf("original error : %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Println("stack trace :%+v", err)
		fmt.Println(err)
		os.Exit(1)
	}
}