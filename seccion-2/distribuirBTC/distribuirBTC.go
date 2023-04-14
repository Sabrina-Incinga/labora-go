package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

/* func includesString(s string, value string) bool{
	var result bool = false
	for _, v := range s {
		if string(v) == value{
			result = true
		}
	}

	return result
} */

func countString(s string, value string) int{
	var result int = 0
	for _, v := range s {
		if strings.ToLower(string(v)) == value{
			result += 1
		}
	}

	return result
}


func distributeBTC(total *int, users []string, distr *map[string]int){
	assignBTCAmount := func(s string) int{
		a, e, i, o, u := 0, 0, 0, 0, 0
	
		a = countString(s, "a")
		e = countString(s, "e")
		i = countString(s, "i")*2
		o = countString(s, "o")*3
		u = countString(s, "u")*4
	
		return a+e+i+o+u
	}
	
	for _, v := range users {
		var count int = 0
		if *total >0 {
			if assignBTCAmount(v) <=10 {
				count = assignBTCAmount(v)	
			} else{ count = 10}

			(*distr)[v] =  count	
		}
		*total = *total-count
	}
}

func main(){
	distributeBTC(&coins, users, &distribution)
	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)

}
