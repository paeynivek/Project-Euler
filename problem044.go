//Kevin Yeap
/*========== Project Euler: Problem 44 ==========
Pentagonal numbers are generated by the formula, Pn=n(3n−1)/2.
The first ten pentagonal numbers are:

1, 5, 12, 22, 35, 51, 70, 92, 117, 145, ...

It can be seen that P_4 + P_7 = 22 + 70 = 92 = P_8.
However, their difference, 70 − 22 = 48, is not pentagonal.

Find the pair of pentagonal numbers, P_j and P_k, for which their sum and
difference are pentagonal and D = |P_k − P_j| is minimised;
what is the value of D?
=========================*/

/*========== program design notes ==========
program should stop searching when the value of D can not possibly get any
lower based on the numbers generated by the formula

so when the pentagon(i+1)-pentagon(i) > lowest_diff. is a solid uppoer bound.

once you find the bound compute diffs back from j-k.
update bound as you find a lower one.


forum notes - interesting potential optimizations read from the forum
=========================*/

/*========== program output ==========
        == Project Euler: Problem 44 ==

answer: 5482660
=========================*/
//this takes 2.1 seconds to run can it be faster?


package main

import (
	"fmt"
	"math"
	"os"
)

func handle_err(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

func main() {
	var result = minimal_diff()

	fmt.Println("\t== Project Euler: Problem 44 ==\n")
	fmt.Printf("answer: %d", result)
}

func minimal_diff() int {

	//find a bound first
	j := 1;
	ld := -1;
	for ; ld == -1; j++ {
		for k := 1; k < j; k++ {
			p_j := pentagon(j)
			p_k := pentagon(k)
			sum := p_j + p_k
			diff := int(math.Abs(float64(p_k) - float64(p_j)))
			if is_pentagon(sum) && is_pentagon(diff) && diff > 1 {
				ld = diff
				break
			}
		}
	}

	//compute diffs until up to the bound.
	for ; pentagon(j+1) - pentagon(j) < ld; j++ {
		for k := j-1; pentagon(j) - pentagon(k) < ld; k-- {
			p_j := pentagon(j)
			p_k := pentagon(k)
			sum := p_j + p_k
			diff := int(math.Abs(float64(p_k) - float64(p_j)))
			if is_pentagon(sum) && is_pentagon(diff) && diff < ld {
				ld = diff //update the bound if a lower bound should be found
			}
		}
	}

	return ld

}

//determines pentagon number based on position
func pentagon(position int) int {
	return position * (3*position - 1) / 2
}

//determines a number is pentagonal
func is_pentagon(num int) bool {
	//triangle number equation reverse is
	//3n^2 - n - 2(num) = 0

	//quadratic formula variables
	a := 3
	b := -1
	c := -2 * num

	//x is the opposite of b +- the rooooot of b^2 - 4*ac all over 2a,
	//and thats the quadratic formula. da da da da da da da
	n := (float64(-b) + math.Sqrt(float64(math.Pow(float64(b), 2)-float64(4*a*c)))) / float64(2*a)

	if float64(int(n)) == n {
		return true
	}

	return false
}
