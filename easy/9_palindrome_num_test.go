/*
 * Created: 2019-11-17 21:51:14
 * Author : Win-Man
 * Email : gang.shen0423@gmail.com
 * -----
 * Last Modified:
 * Modified By:
 * -----
 * Description:
 */

package easy

import (
	"fmt"
	"testing"
)

func isPalindrome(x int) bool {
	var res bool = true
	if x < 0 {
		res = false
	} else if x == 0 {
		res = true
	} else {
		num_list := make([]int, 0)
		var tmp int = x
		for {
			num_list = append(num_list, tmp%10)
			tmp = tmp / 10
			if tmp == 0 {
				break
			}
		}
		for i := 0; i < len(num_list); i++ {
			j := len(num_list) - 1 - i
			if j < i {
				res = true
				break
			}
			if num_list[i] != num_list[j] {
				res = false
				break
			}
		}
	}
	return res
}

func isPalindrome1(x int) bool {
	var res bool = true
	if x < 0 {
		res = false
	} else {
		var y, tmp int = 0, x
		for tmp > 0 {
			y = y*10 + tmp%10
			tmp = tmp / 10
		}
		if x != y {
			res = false
		}
	}
	return res
}

func TestIsPalindromeDriven(t *testing.T) {
	var tests = []struct {
		x    int
		want bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{1, true},
		{100, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.x)
		t.Run(testname, func(t *testing.T) {
			ans := isPalindrome(tt.x)
			if ans != tt.want {
				t.Errorf("isPalindrome(%d) got %t,want %t", tt.x, ans, tt.want)
			}
		})

	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.x)
		t.Run(testname, func(t *testing.T) {
			ans := isPalindrome1(tt.x)
			if ans != tt.want {
				t.Errorf("isPalindrome1(%d) got %t,want %t", tt.x, ans, tt.want)
			}
		})

	}
}
