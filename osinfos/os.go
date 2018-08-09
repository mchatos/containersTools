package osinfos

import	(
	"fmt"
	"os"
	)

	// Os Functions
	func GetOSinfos(){

	  fmt.Println("OS Host Detection and informations : ")
	  hostname, hostnameError := os.Hostname()
	  osWin, osBool := os.LookupEnv("OS")

	  fmt.Printf("PID : %v, PPID : %v\n", os.Getpid(), os.Getppid())

	  if hostnameError != nil {panic(hostnameError)}
	  fmt.Println("Hostname : ", hostname)

	  if osBool == false {
	      fmt.Println("OS Linux")
	      fmt.Printf("User %s lives in %s with UID %v and GID %v.\n", os.Getenv("USER"), os.Getenv("HOME"), os.Geteuid(), os.Getgid())
	      fmt.Printf("Hostname %s.\n", os.Getenv("HOSTNAME"))
	      // if _, err := os.Stat("/usr/sbin/ip"); err == nil {
	      //   cmd := exec.Command("/usr/sbin/ip", "addr")
	      //   os.Stdout
	      // }
	      // fmt.Printf
	    } else {
	      fmt.Println("OS Windows : ", osWin)
	      fmt.Printf("User %s lives in %s.\n", os.Getenv("USERNAME"), os.Getenv("HOMEPATH"))
	  }
	 fmt.Println(" ")
	}
