package main

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	//这里是一个递归,comma会调用[:n-3]
	return comma(s[:n-3] + "," + s[n-3:])
}
