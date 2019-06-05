package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

  // connect to this socket
  conn, _ := net.Dial("tcp", "0.0.0.0:8081")
  for index:=0;index<10; index++ { 
    // read in input from stdin
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text to send of client-new: ")
    text, _ := reader.ReadString('\n')
    // send to socket
    fmt.Fprintf(conn, text + "\n")
    // listen for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message from server-new: "+message)
  }
}
