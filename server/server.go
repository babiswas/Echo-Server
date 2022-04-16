package server
import "net"
import "fmt"

func Create_server(service string,handler func(conn net.Conn)){
   fmt.Println("Server Started")
   tcpAddr,err:=net.ResolveTCPAddr("tcp",service)
   check_error(err)
   listener,err:=net.ListenTCP("tcp",tcpAddr)
   check_error(err)
   for{
       conn,err:=listener.Accept()
       if err!=nil{
          continue
       }
       go handler(conn)
   }
}