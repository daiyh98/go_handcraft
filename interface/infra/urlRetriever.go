package infra

import (
	"fmt"
	"io"
	"net/http"
)

type Retriever struct {
}

func (Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, _ := io.ReadAll(resp.Body)
	fmt.Printf("%s\n", bytes)
	return string(bytes)
}
