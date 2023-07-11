package main

import (
	"fmt"
	"go_handcraft/interface/testing"
)

// func retrieve(url string) []byte {
//func retrieve(url string) string {
//	resp, err := http.Get(url)
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//	bytes, _ := io.ReadAll(resp.Body)
//	fmt.Printf("%s\n", bytes)
//	return string(bytes)
//}

type retriever interface {
	Get(string) string
}

func getRetriever() retriever {
	//return infra.Retriever{}
	return testing.Retriever{}
}
func main() {
	//resp, err := http.Get("https://www.imooc.com")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//bytes, _ := io.ReadAll(resp.Body)
	//fmt.Printf("%s\n", retrieve("https://www.imooc.com"))
	r := getRetriever()
	//fmt.Println(retrieve("https://www.imooc.com"))
	fmt.Println(r.Get("https://www.imooc.com"))
}
