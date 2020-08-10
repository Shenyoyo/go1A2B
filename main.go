package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func gennum(length int) []int {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b := make([]int, length, length)                     //創建length長度的陣列，儲存答案
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //放入亂數環境隨機生成
	for i := 0; i < length; i++ {
		tmp := r.Intn(len(a) - i) //亂數讀取10-i內的數字
		b[i] = a[tmp]             //將a[tmp]值放入答案
		for b[0] == 0 {
			tmp = r.Intn(len(a) - i) //第一位數不能為0
			b[0] = a[tmp]
		}
		a[tmp], a[len(a)-1-i] = a[len(a)-1-i], a[tmp] //使答案數字不重複
	}
	return b
}

func checknum(a, b []int, count int) bool {
	var aa, bb int
	if len(a) != len(b) {
		return false
	}
	dict := make(map[int]int)
	for i := 0; i < len(a); i++ {
		dict[a[i]] = 0
	}
	for i := 0; i < len(a); i++ {
		if _, ok := dict[b[i]]; ok {
			bb++
		}
	}
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			aa++
		}
	}
	bb = bb - aa
	fmt.Printf("%dA%dB，第%d次的猜測\n", aa, bb, count)
	if aa == len(a) {
		fmt.Printf("恭喜猜對了，遊戲結束，總共猜了 %d 次\n", count)
		return true
	} else {
		return false
	}
}

func removeDuplicateElement(addrs []string) []string {
	result := make([]string, 0, len(addrs))
	temp := map[string]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func main() {

	var ipulen string
	var i int
	var count int
	var isString bool
	status := false
	length := 4
	fmt.Println("1A2B 遊戲開始")
	src := gennum(length)
	//答案
	// fmt.Println(src)
	req := make([]int, length, length)
	for !status {
		fmt.Printf("請輸入 %d 數字\n", length)
		fmt.Scanf("%s", &ipulen)
		//檢查輸入是否有字串
		for _, r := range ipulen {
			// fmt.Printf("%c = %v\n", r, unicode.IsLetter(r))
			if unicode.IsLetter(r) {
				isString = true
				break
			} else {
				isString = false
			}
		}
		if isString {
			fmt.Println("輸入有誤，請輸入數字")
			continue
		}
		if len(ipulen) != length {
			fmt.Println("輸入有誤，請輸入長度4的數字")
			continue
		}
		ipulens := strings.Split(ipulen, "")
		if ipulens[0] == "0" {
			fmt.Println("第一位不能為0，請重新輸入")
			continue
		}
		if len(removeDuplicateElement(ipulens)) != length {
			fmt.Println("輸入答案重複數字，請重新輸入")
			continue
		}
		ipu, _ := strconv.Atoi(ipulen)
		for i = length - 1; ipu > 0; i-- {
			req[i] = ipu % 10
			ipu = ipu / 10
		}
		count++
		status = checknum(src, req, count)
	}
	fmt.Println("程式將於10秒後關閉")
	time.Sleep(time.Duration(10) * time.Second)
}
