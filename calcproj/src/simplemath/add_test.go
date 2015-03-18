/**
* @Author:LCore
* @Model:addTest
 */

package simplemath

import (
	"testing"
)

func TestAddl(t *testing.T) {
	r := Add(1, 2)
	if r != 3 {
		t.Errorf("Add(1,2) failed. Got %d,EXPECTED 3.", r)
	}
}
