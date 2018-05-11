

package main

import (
	"testing"
	"os"

)


func findFiles(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}



func testFile1(t *testing.T){
	r := findFiles("form.html")
	if r == false {
		t.Error("Did not find ....")
	}
}


func testFile2(t *testing.T){
	r := findFiles("index.html")
	if r == false {
		t.Error("Did not find ....")
	}
}

func testFile3(t *testing.T){
	r := findFiles("info.html")
	if r == false {
		t.Error("Did not find ....")
	}
}


func testFile4(t *testing.T){
	r := findFiles("print.html")
	if r == false {
		t.Error("Did not find ....")
	}
}





