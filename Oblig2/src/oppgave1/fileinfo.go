package main

import (
	"fmt"
	"os"
	"log"
	"flag"
	"path/filepath"
)

func main() {
// først  go build fileinfo.go

	filePeker := flag.String("f", "", "filnavn")
	flag.Parse()
 // fileinfo -f [filnavn] ofr å finne fil
	path, err := filepath.Abs(*filePeker)
	if err !=nil {
		log.Fatal(err)
	}

	fmt.Println("Programmet skal returnere informasjon om en fil")
	file, err := os.Lstat(*filePeker) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	x := float64(file.Size())
	kb := x/1024
	mb := kb/1024
	gb := mb/1024
	fmt.Printf("Info om '%s' : \n", path)
	fmt.Printf("Størrelse : %.0fB, %fKB, %fMB, %.9fGB \n", x, kb, mb, gb)


	mode :=file.Mode()
	dir := mode.IsDir()
	reg := mode.IsRegular()
	append := mode&os.ModeAppend !=0
	dev := mode%os.ModeDevice !=0
	sym := mode%os.ModeSymlink !=0
	if dir == true{
		fmt.Println("is directory")
	} else {
		fmt.Println("is not directory")
	}
	if reg == true{
		fmt.Println("is regular file")
	} else {
		fmt.Println("is not regular file")
	}
	fmt.Println("Permissions i Unix ",mode)
	if append == true{
		fmt.Println("is append-only")
	} else {
		fmt.Println("is not append-only")
	}
	if dev == true {
		fmt.Println("is device file")
	} else {
		fmt.Println("is not device file")
	}
	if sym == true{
		fmt.Println("is symbolic link")

	} else {
		fmt.Println("is not symbolic link")
	}






}
