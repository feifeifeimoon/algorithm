// ## 请在下方进行输入 ( 支持Markdown、插入公式及 LaTex 数学公式，点击上方按钮“环境说明”查看详情 )

//两个协程交替打印1-100的奇数偶数
package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello, world")

	gCh := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)
	ch1, ch2 := make(chan struct{}, 1), make(chan struct{}, 1)
	print := func(ch1 <-chan struct{}, ch2 chan struct{}) {
		defer wg.Done()
		for {
			<-ch1
			v := <-gCh
			fmt.Println(v)
			ch2 <- struct{}{}
			if v == 99 || v == 100 {
				break
			}
		}
	}

	go print(ch1, ch2)
	go print(ch2, ch1)
	ch1 <- struct{}{}
	for i := 1; i <= 100; i++ {
		gCh <- i
	}

	wg.Wait()
}
