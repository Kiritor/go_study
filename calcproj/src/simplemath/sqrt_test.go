/**
* @Author:LCore
* @Model:sqrtTest
 */

package simplemath

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	r := Sqrt(9)
	if r != 3 {
		t.Error("开方运算出错")
	}
}
