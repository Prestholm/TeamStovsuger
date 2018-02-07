package ascii

import (
	"fmt"

)




func IterateOverASCIIStringLiteral(sl string) {
	// Kode for Oppgave 2a

	for i := 0; i < len(sl); i++ {
		fmt.Printf("%X %c %b \n", sl[i], sl[i], sl[i])

	}
}

// Kode for Oppgave 2b
func ExtendedASCIIText()string {
	t1 := []byte("\x80\xf7\xbe\x64\x6f\x6c\x6c\x61\x72")
	for i := 0; i < len(t1); i++ {
		fmt.Printf("%c", t1[i])
	}
	t2 := "\x80\xf7\xbe\x64\x6f\x6c\x6c\x61\x72"
	return t2
}
