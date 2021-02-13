/*
 * Created: 2021-02-13 18:11:41
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

func reverseLeftWords(s string, n int) string {
	str1 := s[:n]
	str2 := s[n:]
	res := str2 + str1
	return res
}

func TestReverseLeftWords(t *testing.T) {
	var tests = []struct {
		s    string
		n    int
		want string
	}{
		{"abcdefg", 2, "cdefgab"},
		{"lrloseumgh", 6, "umghlrlose"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := reverseLeftWords(tt.s, tt.n)
			if ans != tt.want {
				t.Errorf("reverseLeftWords(%s,%d) got %s,want %s", tt.s, tt.n, ans, tt.want)
			}
		})
	}
}
