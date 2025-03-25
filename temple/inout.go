package temple

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

func test() {
	// 1.多行输入， 每行两个整数
	var a, b int
	for {
		if _, err := Scan(&a, &b); err != nil {
			break
		}
		// ...
	}

	// 2.多组数据， 每组第一行为n，之后输入n行两个整数
	var n int
	for {
		if _, err := Scan(&n); err != nil {
			break
		}
		for n > 0 {
			if _, err := Scan(&a, &b); err != nil {
				break
			}
			// ...
			n--
		}
	}

	// 3.若干行输入，每行输入两个整数，遇到特定条件终止
	for {
		if _, err := Scan(&a, &b); err != nil || (a == 0 && b == 0) {
			break
		}
		// ...
	}

	// 4.若干行输入，遇到0终⽌，每⾏第⼀个数为N，表示本⾏后⾯有n个数
	for {
		if _, err := Scan(&n); err != nil || n == 0 {
			break
		}
		for n > 0 {
			if _, err := Scan(&a); err != nil {
				break
			}
			// ...
			n--
		}
	}

	// 5.若干行输入，每⾏包括两个整数a和b，由空格分隔，每⾏输出后接⼀个空⾏
	for {
		if _, err := Scan(&a, &b); err != nil || (a == 0 && b == 0) {
			break
		}
		Printf("%d\n\n", a+b)
	}

	// 6.多组n⾏数据，每⾏先输⼊⼀个整数N，然后在同⼀⾏内输⼊M个整数,每组输出之间输出⼀个空⾏。
	var m int
	for {
		if _, err := Scan(&n); err != nil {
			break
		}
		for i := 0; i < n; i++ {
			Scan(&m)

			for j := 0; j < m; j++ {
				Scan(&a)
			}
			// ...
			if i < n-1 {
				// ...
			}
		}
	}

	// 7.多组测试样例，每组输⼊数据为字符串，字符⽤空格分隔,输出为⼩数点后两位
	reader := bufio.NewReader(os.Stdin)
	for {
		s, _, err := reader.ReadLine()
		s_list := strings.Split(string(s), " ")
		if err != nil {
			break
		}
		for i := 0; i < len(s_list); i++ {

		}
	}
	Println(Sprintf("%.2f", 3.14152))

	// 8.多组测试⽤例，第⼀⾏为正整数n, 第⼆⾏为n个正整数，n=0时，结束输⼊，每组输出结果的下⾯都输出⼀个空⾏
	for {
		if _, err := Scan(&n); err != nil || n == 0 {
			break
		}
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			Scan(&nums[i])
		}
	}

	// 9.多组测试数据，每组数据只有⼀个整数，对于每组输⼊数据，输出⼀⾏，每组数据下⽅有⼀个空⾏。
	for {
		if _, err := Scan(&n); err != nil || n == 0 {
			break
		}
		for n != 0 {
			a = n % 10
			n = n / 10
		}
	}

	// 10.多组测试数据，每个测试实例包括2个整数M，K（2<=k<=M<=1000)。M=0，K=0代表输⼊结束。
	var k int
	for {
		if _, err := Scan(&m, &k); err != nil || (m == 0 && k == 0) {
			break
		}
	}
}
