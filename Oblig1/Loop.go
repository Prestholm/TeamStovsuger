package main


import "fmt"
import "os"
import "os/signal"
import "syscall"

//Chan står for "channel":
func main() {
	sign := make(chan os.Signal, 1)
	quitCh := make(chan bool, 1)

	//Valg av signal (SIGINT):
	signal.Notify(sign, syscall.SIGINT)

	//Kjører programmet:
	go func() {
		run := <-sign
		fmt.Println(run)
		quitCh <- true
	}()

	//Evig løkke:
	for i := 1;  ; i++ {

		select {
		default: fmt.Println("Venter på signal...")

			//Avslutningsmelding ved mottatt signal (Ctrl + C):
		case <- quitCh:
			fmt.Println("-------------------Signal mottatt!-------------------")
			os.Exit(1)
		}
	}
}
