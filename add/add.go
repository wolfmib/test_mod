package add

import "github.com/sirupsen/logrus"

import "fmt"

func Show(){
	fmt.Println("haha")
}


func MyNum(num_1 int, num_2 int) int {
	var result int = num_1 + num_2
	logrus.Info(fmt.Sprintf("running %v + %v = % v", num_1, num_2, result))
	return result

}

func MyStr(string_1 string, string_2 string) string {
	//logrus.info("information")
	result := string_1 + string_2
	return result

}

/*
package numbers

import "math"

// Checks if a number is prime or not
func IsPrime(num int) bool {
    for i := 2; i <= int(math.Floor(math.Sqrt(float64(num)))); i++ {
        if num%i == 0 {
            return false
        }
	}
    return num > 1
}
*/
