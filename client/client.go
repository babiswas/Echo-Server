package client
import "net"
import "fmt"
import "bufio"
import "os"

func Create_client(service string){
  var buf [512]byte
  tcpAddr,err:=net.ResolveTCPAddr("tcp",service)
  check_error(err)
  conn,err:=net.DialTCP("tcp",nil,tcpAddr)
  check_error(err)
  scanner:=bufio.NewScanner(os.Stdin)
  for{
      fmt.Println("Enter Message:")
      scanner.Scan()
      n,err:=conn.Write([]byte(scanner.Text()))
      fmt.Printf("written %d bytes:\n",n)
      check_error(err)
      fmt.Println("Waiting for the server\n")
      n,err=conn.Read(buf[0:])
      check_error_close(err,conn)
      fmt.Println(string(buf[0:n]))
  }
}