//containtools.go
package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"containtools/osinfos"
)

// Fonction sleep
func fSleep(i int) {

	fmt.Println("So tired... i will sleep for ", i, " seconds...")
	for j := 1; j <= i; j++ {
		time.Sleep(1 * time.Second)
	}
}

func fWaiTime() int {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var w = (r1.Intn(100))

	return w
}

func main() {
	
	siesta := 0
	loop := 0
	avg := 0
	var tab [100]int
	
	fmt.Println("")
	osinfos.OSinfos()

	// infinite loop
	for {
		sleepTime := fWaiTime()
		tab[loop] = sleepTime
		loop++
		
		if ( sleepTime >= 90 || loop == 100 ) {
			fmt.Printf("Siesta finished after %v Seconds, with %v micro sleeps generated \n", (siesta), (loop) )
			avg=((siesta)/(loop-1))
			fmt.Printf(" %v seconds micro sleeps average (%v loops) \n", avg, (loop-1) )
			for h := 0; h < loop; h++ { fmt.Printf(" %d ", tab[h]) }
			fmt.Println("")
			os.Exit(1)
		}

		siesta = siesta + sleepTime
		fSleep(sleepTime)

	}
}
