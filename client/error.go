package client
import "os"
import "log"
import "net"


func check_error(err error){
   if err!=nil{
      log.Fatal(err)
      os.Exit(1)
   }
}

func check_error_close(err error,conn net.Conn){
   if err!=nil{
      log.Fatal(err)
      conn.Close()
      os.Exit(1)
   }
}