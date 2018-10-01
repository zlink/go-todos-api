package helper

import "fmt"

func echo(data []string) {
	for _, item := range data {
		fmt.Println(item)
	}
}
