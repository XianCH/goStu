package errortest

import (
	"fmt"
	"os"
)

func readFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func processFile2(filename string) error {
	data, err := readFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %s", err)
	}
	fmt.Println(string(data))
	return nil
}
func processFile(filename string) error {
	data, err := readFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	fmt.Println(string(data))
	return nil
}

//
// func main() {
// 	err := processFile("1.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Println(errors.Is(err, os.ErrNotExist))
// 		err = errors.Unwrap(err)
// 		fmt.Println(err)
// 		err = errors.Unwrap(err)
// 		fmt.Println(err)
// 		return
// 	}
//
// }
