package filesop

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func WordByWordScan()  {
	file, err := os.Open("./scanner/sample.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
func LineByLineScan(){
	file, err := os.Open("./scanner/sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
func ReadEntryFile() {
	fileBytes, err := ioutil.ReadFile("./scanner/sample.txt")
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	fmt.Println(fileString)
}
func ReadEntryFile1(){
	file, err := os.Open("./scanner/sample.txt")
	if err != nil {
		log.Fatalf("Error while opening file, Err:%s", err)
	}
	defer file.Close()
	fileBuffer := new(bytes.Buffer)
	fileBuffer.ReadFrom(file)
	fileString := fileBuffer.String()
	fmt.Print(fileString)
}
func ReadEntryFile2(){
	file, err := os.Open("./scanner/sample.txt")
	if err != nil {
		log.Fatalf("Error while opening file, Err:%s",err)
	}
	defer file.Close()
	sb := new (strings.Builder)
	io.Copy(sb,file)
	fmt.Println(sb.String())
}
func WriteFile(){
	file, err := os.Create("./temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriter(file)
	linesToWrite := []string{"This is an example", "to show how", "to write to a file", "line by line"}
	for _, line := range linesToWrite {
		bytesWritten, err := writer.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("Got error while writing to a file. Err:%s",err)
		}
		fmt.Printf("Bytes Written: %d\n", bytesWritten)
		fmt.Printf("Available: %d\n", writer.Available())
		fmt.Printf("Buffered: %d\n", writer.Buffered())
	}
}
func WriteFileWithSize(){
	file, err := os.Create("./temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	writer := bufio.NewWriterSize(file, 10)
	linesToWrite := []string{"This is an example", "to show how", "to write to a file", "line by line"}
	for _, line := range linesToWrite {
		bytesWritten, err := writer.WriteString(line +"\n")
		if err != nil {
			log.Fatalf("Got error while writting to a file. Err:%s\n",err)
		}
		fmt.Printf("Bytes written: %d\n", bytesWritten)
		fmt.Printf("Available: %d\n",writer.Available())
		fmt.Printf("Buffered: %d\n", writer.Buffered())
	}
	writer.Flush()
}
func WriteFileImmediately(){
	file, err := os.Create("./temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	linesToWrite := []string{"This is an example", "to show how", "to write to a file", "line by line"}

	for _, line := range linesToWrite {
		file.WriteString(line+"\n")
	}
}

func WriteFileImmediately1(){
	//linesToWrite := []string{"This is an example", "to show how", "to write to a file", "line by line"}
	linesToWrite := "This is an example\", \"to show how\", \"to write to a file\", \"line by line"
	err := ioutil.WriteFile("./temp.txt", []byte(linesToWrite), 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func AppendToFile(){
	err := ioutil.WriteFile("./temp.txt", []byte("first line\n"), 0644)
	if err != nil {
		log.Fatal(err)
	}
	file,err := os.OpenFile("./temp.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if _, err := file.WriteString("second line\n"); err != nil {
		log.Fatal(err)
	}
	data,err := ioutil.ReadFile("./temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
func RemoveFolder(){
	err:= os.Remove("sample_folder")
	if err != nil {
		log.Fatal(err)
	}
}
func Chtimes(){
	fileName :="sample.txt"
	currentTime := time.Now().Local()
	err := os.Chtimes(fileName,currentTime, currentTime)
	if err != nil {
		log.Fatal(err)
	}
}
func Rename(){
	file, err := os.Create("./temp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = os.Chmod("./temp.txt", 0777)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Rename("./temp.txt", "./newtemp.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir("temp", 0755)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Rename("temp", "newTemp")
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir("/Users")
	newDir, err := os.Getwd()
	fmt.Printf("Current Working Directory: %s\n", newDir)
}