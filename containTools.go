package main

import (
    "fmt"
    "time"
    "math/rand"
    "os"
)

// Os Functions
func fOSinfos(){

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
    } else {
      fmt.Println("OS ", osWin)
      fmt.Printf("User %s lives in %s.\n", os.Getenv("USERNAME"), os.Getenv("HOMEPATH"))
  }
 fmt.Println(" ")
}

// Fonction sleep
func fSleep(i int){

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

  fmt.Println(" ")
  fOSinfos()

  // infinite loop
  for {
      sleepTime := fWaiTime()
      if sleepTime >= 90 {
        fmt.Printf("Siesta finished after %v Seconds, with %v micro sleeps generated ", siesta, loop)
        os.Exit(1)
      }

    siesta = siesta + sleepTime
    loop++
    fSleep(sleepTime)

  }
}
