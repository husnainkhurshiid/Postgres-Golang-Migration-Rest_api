package shared

import "fmt"

func CheckError(err error) {
	if err != nil {
		fmt.Print(err)
	}
}
