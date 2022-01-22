package main

import ("fmt"
        "os"
        "bufio"
        "time")

func main() {

  ctime := time.Now().Format("01-02-2006 15:04:05")


  textScanner := bufio.NewScanner(os.Stdin)

  YNQ := true

  openfile, _ := os.OpenFile("go-log.txt", os.O_RDWR|os.O_APPEND, 0644)

  for YNQ == true {

    fmt.Println("What do you want to add to the log?")

    textScanner.Scan()

    openfile.WriteString(ctime + "\n")
    openfile.WriteString(textScanner.Text() + "\n")

    //add date and time here

    fmt.Println("Would you like to add anything else?")

    textScanner.Scan()

    if textScanner.Text() == "no" {
      YNQ = false
    }
  }
  openfile.Close()
}
