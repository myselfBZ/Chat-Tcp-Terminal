package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
)



func ReadLine() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a message: ")
    s, err := reader.ReadString('\n')
    if err != nil{
        log.Fatal(err)
    }
    if s == ""{
        return s
    }
    return s
}
