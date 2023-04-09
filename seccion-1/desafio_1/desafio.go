package main
import "fmt"

func main(){
	v := 120
	x := 860
	s := 80

	fmt.Println(segmentarValorPorRangos((v)))
	fmt.Println(segmentarValorPorRangos((x)))
	fmt.Println(segmentarValorPorRangos((s)))

}

func segmentarValorPorRangos(num int) (int, int, int, int, int){
	var s1, s2, s3, s4, s5 int

	if num > 0{
		if num <= 50 {
			s1 = num
			s2, s3, s4, s5 = 0, 0, 0, 0
		} else if num <= 100{
			s1 = 50
			s2 = num - s1
			s3, s4, s5 = 0, 0, 0
		} else if num <= 700{
			s1, s2 = 50, 50
			s3 = num - s1 - s2
			s4, s5 = 0, 0
		} else if num <= 1500{
			s1, s2 = 50, 50
			s3 = 600
			s4 = num - s1 - s2 - s3
			s5 = 0
		} else {
			s1, s2 = 50, 50
			s3 = 600
			s4 = 800
			s5 = num - s1 - s2 - s3 - s4
		}
	}

	return s1, s2, s3, s4, s5
}