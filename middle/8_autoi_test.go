/*
 * Created: 2020-02-27 15:53:36
 * Author : Win-Man
 * Email : gang.shen0423@gmail.com
 * -----
 * Last Modified:
 * Modified By:
 * -----
 * Description:
 */

package middle

import (
	"fmt"
	"testing"
)

const (
	INT_MAX  = (1 << 31) - 1
	INT_MIN = -(1 << 31)
)

func myAtoi(str string) int {
	var num int = 0
	for i := 0; i < len(str); i++ {
		if str[i] == '-' || str[i] == '+' {
			for j := i + 1; j < len(str); j++ {
				if str[j] >= '0' && str[j] <= '9' {
					num = num*10 + int(str[j]-'0')
					if num > INT_MAX && str[i] == '-'{
						return INT_MIN
					}else if num > INT_MAX{
						return INT_MAX
					}
				} else {
					break
				}
			}
			if str[i] == '-' {
				num = -num
			}
			break
		}else if str[i] >= '0' && str[i] <= '9'{
			
			num = int(str[i] - '0')
			for j := i + 1; j < len(str); j++ {
				if str[j] >= '0' && str[j] <= '9' {
					num = num*10 + int(str[j]-'0')
					if num > INT_MAX{
						return INT_MAX
					}
				} else {
					break
				}
			}
			break
		}else if str[i] == ' '{
			continue
		}else{
			break
		}
	}
		return num
}

func TestMyAtoi(t *testing.T) {
	var tests = []struct {
		a    string
		want int
	}{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", INT_MIN},
		{"9223372036854775808",2147483647},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := myAtoi(tt.a)
			if ans != tt.want {
				t.Errorf("myAtoi(%v) got:%d want:%d", tt.a, ans, tt.want)
			}
		})
	}
}
