package packageone

import "fmt"

var PackageVar = "PackageVar"

func PrintMe(str1, str2 string) {
	fmt.Println(str1, str2, PackageVar)
}
