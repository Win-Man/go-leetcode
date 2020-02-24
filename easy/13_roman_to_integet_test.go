/*
 * Created: 2019-11-18 21:16:35
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

func romanToInt(s string) int {
	var res int = 0
	symbolList := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	lastCharVal := 0
	for _, val := range s {
		if symbolList[val] > lastCharVal {
			res += symbolList[val]
			res -= lastCharVal * 2
			lastCharVal = symbolList[val]
		} else {
			res += symbolList[val]
			lastCharVal = symbolList[val]
		}
	}
	return res
}

func TestRomanToIntDriven(t *testing.T) {
	var tests = []struct {
		s    string
		want int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := romanToInt(tt.s)
			if ans != tt.want {
				t.Errorf("romanToInt(%s) got:%d want:%d", tt.s, ans, tt.want)
			}
		})
	}
}