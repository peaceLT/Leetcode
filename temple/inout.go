package temple

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

func test() {
	// 多行输入， 每行两个整数
	var a, b int
	for {
		if _, err := Scan(&a, &b); err != nil {
			break
		}
		// ...
	}

	// 多组数据， 每组第一行为n，之后输入n行两个整数
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

	// 若干行输入，每行输入两个整数，遇到特定条件终止
	for {
		if _, err := Scan(&a, &b); err != nil {
			break
		}
		if a == 0 && b == 0 {
			break
		}
		// ...
	}

	// 若干行输入，遇到0终⽌，每⾏第⼀个数为N，表示本⾏后⾯有n个数
	for {
		if _, err := Scan(&n); err != nil {
			break
		}
		if n == 0 {
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

	// 若干行输入，每⾏包括两个整数a和b，由空格分隔，每⾏输出后接⼀个空⾏
	for {
		if _, err := Scan(&a, &b); err != nil {
			break
		}
		if a == 0 && b == 0 {
			break
		}
		Printf("%d\n\n", a+b)
	}

	// 多组n⾏数据，每⾏先输⼊⼀个整数N，然后在同⼀⾏内输⼊M个整数,每组输出之间输出⼀个空⾏。
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

	// 多组测试样例，每组输⼊数据为字符串，字符⽤空格分隔,输出为⼩数点后两位
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
}
