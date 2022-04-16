package main
import "fmt"
import "os"
import "app/client"

func main(){
   fmt.Println("Creating client")
   client.Create_client(os.Args[1])
}