package osinfos

import	(
	"fmt"
	"os"
	)


func OSDetection() (string) {

	fmt.Println("OS Host Detection... ")

  var operatingsystem string

	_, osBool := os.LookupEnv("OS")

		  if osBool == false {
				operatingsystem = "unix"
				} else {
				operatingsystem = "win"
				}

	return operatingsystem

}

// Os Functions
func GetOSinfos(){

		hostname, hostnameError := os.Hostname()

	  fmt.Printf("PID : %v, PPID : %v\n", os.Getpid(), os.Getppid())

	  if hostnameError != nil {panic(hostnameError)}
	  fmt.Println("Hostname : ", hostname)

		  if OSDetection() == "unix" {
	      fmt.Printf("User %s lives in %s with UID %v and GID %v.\n", os.Getenv("USER"), os.Getenv("HOME"), os.Geteuid(), os.Getgid())
	      fmt.Printf("Hostname %s.\n", os.Getenv("HOSTNAME"))
	      // if _, err := os.Stat("/usr/sbin/ip"); err == nil {
	      //   cmd := exec.Command("/usr/sbin/ip", "addr")
	      //   os.Stdout
	      // }
	      // fmt.Printf
	    } else {
	      fmt.Printf("User %s lives in %s.\n", os.Getenv("USERNAME"), os.Getenv("HOMEPATH"))
	  }
	 fmt.Println(" ")
	}
