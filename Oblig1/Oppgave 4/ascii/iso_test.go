package ascii

import (



	"testing"


)

func testExtendedASCII(t * testing.B){
	g := ExtendedASCIIText()
	if tests(g) == false {
		t.Errorf("fail")
	}
}

func tests(z2 string) bool{
for _, i := range z2{
		if i<0x80 ||i>0xff{
			return false
		}
		}
	return true
}
