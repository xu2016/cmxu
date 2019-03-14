package main

import (
	"cmxu/pcapi/xtsz"
	"fmt"
)

func main() {
	fmt.Println("*********1111111111111********")
	fmt.Println(xtsz.Menus[0])
	fmt.Println("*********2222222222222********")
	fmt.Println(xtsz.Menus[1])
	fmt.Println("*********3333333333333********")
	fmt.Println(xtsz.Menus[2])
	fmt.Println("*********4444444444444********")
	fmt.Println(xtsz.Menus[3])
	fmt.Println(xtsz.MenuSort)
	t := map[string]int{
		`ddgl`: 31, `sjfx`: 63, `zygl`: 63, `qxsz`: 1,
	}
	fmt.Println(xtsz.Getmenu(t))

}
