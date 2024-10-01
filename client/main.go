package main

import (
	"fmt"
	"log"
	"net"

	"github.com/myselfBZ/chat/pkg"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")    
    if err != nil{
        log.Fatal(err)
    }
    buf := make([]byte, 128)
    n, err := conn.Read(buf)
    log.Println("Message from the server: ", string(buf[:n]))
    go func(){
        for{
            input := pkg.ReadLine() 
            if input == ""{
                fmt.Println("Enter something")
            } else{
                conn.Write([]byte(input))
            }
        }
    }()
    for{
        n, err := conn.Read(buf)
        if err != nil{
            log.Fatal("Error reading from the server: ", err)
        }
        log.Println(string(buf[:n]))
    }
}
