package main

import (
	"log"
	"sync"

	"github.com/myselfBZ/chat/internal"
)


func main(){
    wg := sync.WaitGroup{}
    s := internal.NewServer()
    log.Println("starting the server")
    // listen for incoming requests 
    wg.Add(1)
    go func(){
        s.HandleConn()
        wg.Done()
    }()
    // Handle messages
    wg.Add(1)
    go func(){
        s.HandleMessage()
        wg.Done()
    }()
    // broadcast 
    wg.Add(1)
    go func(){
        s.Broadcast()
        wg.Done()
    }()
    wg.Wait()
}

