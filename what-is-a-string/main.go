package main

import "strings"

func main() {
	str := "alpha alpha alpha alpha alpha"
	str = replaceNth(str, "alpha", "beta", 3)
	println(str)

}

func replaceNth(s, old, new string, n int) string {
	// index
	i := 0

	for j := 1; j <= n; j++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			// we did not find it
			break
		}

		i = i + x
		if j == n {
			return s[:i] + new + s[i+len(old):]
		}

		i += len(old)
	}

	return s
}
