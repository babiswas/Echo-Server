package server
import "fmt"
import "os"

func check_error(err error){
   if err!=nil{
     fmt.Fprintf(os.Stderr,"Fatal error:%s",err.Error())
     os.Exit(1)
   }
}