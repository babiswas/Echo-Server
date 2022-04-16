package main
import "os"
import "fmt"
import "app/server"


func main(){
  fmt.Println("Starting a server")
  server.Create_server(os.Args[1],server.Handler)
}