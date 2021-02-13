/*
 * Created: 2021-02-13 17:06:28
 * Author : Win-Man
 * Email : gang.shen0423@gmail.com
 * -----
 * Last Modified:
 * Modified By:
 * -----
 * Description: 
 */

package pkg 

func compareArrayInt(a,b []int)bool{
	if len(a) != len(b){
		return false
	}
	if (a==nil) != (b==nil){
		return false
	}
	for i,v := range a{
		if v != b[i]{
			return false
		}
	}
	return true
}