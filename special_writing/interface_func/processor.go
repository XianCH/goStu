package main

import (
	"fmt"
	"strings"
)

// Processor 定义了处理数据的接口
type Processor func(data string) (string, error)

// Compose 组合多个处理器成一个
func Compose(processors ...Processor) Processor {
	return func(data string) (string, error) {
		for _, processor := range processors {
			var err error
			if data, err = processor(data); err != nil {
				return "", err
			}
		}
		return data, nil
	}
}

// TrimProcessor 修剪空白字符的处理器
func TrimProcessor(data string) (string, error) {
	return strings.TrimSpace(data), nil
}

// UppercaseProcessor 将数据转换为大写的处理器
func UppercaseProcessor(data string) (string, error) {
	return strings.ToUpper(data), nil
}

// AddPrefixProcessor 添加前缀的处理器
func AddPrefixProcessor(prefix string) Processor {
	return func(data string) (string, error) {
		return prefix + data, nil
	}
}

func processorTest() {
	// 创建数据处理管道：修剪空白字符 -> 转换为大写 -> 添加前缀
	processor := Compose(
		TrimProcessor,
		UppercaseProcessor,
		AddPrefixProcessor("PREFIX_"),
	)

	// 处理数据
	result, err := processor("  hello world  ")
	if err != nil {
		fmt.Println("Error processing data:", err)
		return
	}

	// 打印结果
	fmt.Println("Processed data:", result)
}
