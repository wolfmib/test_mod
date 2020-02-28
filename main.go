package main


import "github.com/wolfmib/test_mod/add"

func main() {

	add.Show()
	
	var num_1 int = 1
	var num_2 int = 2
	_ = add.MyNum(num_1, num_2)
	
}
