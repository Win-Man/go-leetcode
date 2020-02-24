/*
 * Created: 2019-11-17 23:05:37
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

func reverse(x int) int {
	if x <= 0 {
		x = -x
		var res int = 0
		for x > 0 {
			res = res*10 + x%10
			x = x / 10
		}
		if res > 1 << 31{
			res = 0
		}
		return -res
	} else {
		var res int = 0
		for x > 0 {
			res = res*10 + x%10
			x = x / 10
		}
		if res > 1 << 31 - 1{
			res = 0
		}
		return res
	}
}

func TestReverseDriven(t *testing.T) {
	var tests = []struct {
		x    int
		want int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{1534236469,0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.x)
		t.Run(testname, func(t *testing.T) {
			ans := reverse(tt.x)
			if ans != tt.want {
				t.Errorf("reverse(%d) got %d,want %d", tt.x, ans, tt.want)
			}
		})
	}
}
