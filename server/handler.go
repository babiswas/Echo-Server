package server
import "net"
import "fmt"


func Handler(conn net.Conn){
  fmt.Println("Connected to the server")
  defer conn.Close()
  var buf [512]byte
  for{
     n,err1:=conn.Read(buf[0:])
     fmt.Printf("Read: %s \n",string(buf[0:n]))
     if err1!=nil{
        return
     }
     n,err2:=conn.Write(buf[0:n])
     fmt.Printf("Written %d bytes to the client \n",n)
     if err2!=nil{
        return
     }
  }
}