package main

import (
	"fmt"
	"go_handcraft/retriever/mock"
	"go_handcraft/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func post(p Poster) {
	p.Post("https://www.imooc.com", map[string]string{
		"name":   "daiyuehao",
		"course": "golang",
	})
}

type RetrieverPoster interface { //接口的组合
	Retriever
	Poster
}

func Session(s RetrieverPoster) {
	s.Get()
	s.Post()
}

func main() {
	//接口变量r内部封装了r本身（r的本身就是实现者类型）的动态类型以及指向r本身（实现者变量）的指针。
	//虽然你输出r没有被断言时候的类型时依然是之前的动态类型，但是实际上你不断言获得本身的实现类型变量的话
	//你是没有办法拿到里面的字段的
	var r Retriever
	r = &mock.Retriever{Contents: "this is fake imooc.com"}
	inspect(r)
	r = &real.Retriever{UserAgent: "Mozilla/5", TimeOut: time.Minute}
	inspect(r)
	//fmt.Println(download(r))

	//类型断言 type assertion
	//类型断言的时候要防止两种错误：1、动态类型不存在（没有实现接口）2、断言了不是该变量的动态类型
	//防止错误的方法就是加上ok，如果类型断言成功，ok 就会是 true，否则就会是 false，程序不会抛出错误。
	//r = &mock.Retriever{Contents: "this is fake imooc.com"}
	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Println("UserAgent:", realRetriever.UserAgent)
	} else {
		fmt.Println("not a *real retriever")
	}

}

// 如何知道接口变量的具体内容
func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) { //特殊的switch类型断言，将r这个接口变量打回原形。type switch
	// 变回他本身的动态类型从而取得内部字段
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents) //此处如果使用r.Contents就会报错
		// 因为r还没有被断言为它的动态类型。虽然你输出它的类型的确是它的动态类型
		fmt.Printf("%T, %T\n", v, r)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

//接口变量自带指针（所以接口也是切片、map类似的引用类型）
//接口变量同样使用值传递，几乎不需要使用接口的指针（类似于切片、map）
//如果接口由指针接收者来实现，那么具体的接口变量必须也是指针；由值接收者实现的话则两种都可以

//取得接口变量的方法上面展示了两种：type assertion以及type switch
//接下来介绍第三种：interface{}表示任何类型
