package osinfos

import	(
	"fmt"
	"os"
	)

// Os Functions
func OSinfos() {

	fmt.Println("OS Host Detection and informations : ")
	hostname, hostnameError := os.Hostname()
	osWin, osBool := os.LookupEnv("OS")

	fmt.Printf("PID : %v, PPID : %v\n", os.Getpid(), os.Getppid())

	if ( hostnameError != nil ) {
		panic(hostnameError)
	}
	fmt.Println("Hostname : ", hostname)

	if ( osBool == false ) {
		fmt.Println("OS Linux")
		fmt.Printf("User %s lives in %s with UID %v and GID %v.\n", os.Getenv("USER"), os.Getenv("HOME"), os.Geteuid(), os.Getgid())
		fmt.Printf("Hostname %s.\n", os.Getenv("HOSTNAME"))
	} else {
		fmt.Println("OS ", osWin)
		fmt.Printf("User %s lives in %s.\n", os.Getenv("USERNAME"), os.Getenv("HOMEPATH"))
	}
	fmt.Println(" ")
}
