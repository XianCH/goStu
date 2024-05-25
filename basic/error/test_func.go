package errortest

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func readFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", Wrap(err, fmt.Sprintf("failed to read file: %v", err))
	}
	return string(data), nil
}

// processFile processes the content of a file.
func ProcessFile(filename string) error {
	content, err := readFile(filename)
	if err != nil {
		return Wrap(err, "failed to process file")
	}

	fmt.Println("File content:", content)
	return nil
}

func err1() error {
	return errors.New("err1")
}

func Sum(a, b int) error {
	if a+b != 3 {
		err := err1()
		return Wrap(err, "sum add error")
	}
	return nil
}

var DB *gorm.DB

func InitDB() error {
	dsn := "root:3954@tcp(127.0.0.1:3306)/g{otest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return Wrap(err, "gorm init error")
	}
	DB = db
	return nil
}

func causePanic() {
	panic(New("something went wrong"))
}

func TestPanicHelper() {
	defer PanicHelper()
	causePanic()
}
