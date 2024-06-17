package helper

import (
	"fmt"
	"io"
	"net/http"
)

func LoadJson(link string) ([]byte, error) {
	data, err := http.Get(link)
	if err != nil {
		fmt.Println("failed to get json fata")
		return nil, err
	}

	defer data.Body.Close()
	res, err := io.ReadAll(data.Body)
	if err != nil {
		fmt.Println("failed to get json fata")
		return nil, err
	}

	return res, nil
}
