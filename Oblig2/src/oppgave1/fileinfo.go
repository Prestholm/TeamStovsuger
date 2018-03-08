package oblig2

import (
	"fmt"

)

func main() {

	fmt.Println("Programmet skal returnere informasjon om en fil")
type FileInfo interface{
	Size() int64
	isDir() bool
	Sys() interface{}
}
}
