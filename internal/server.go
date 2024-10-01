package internal

import (
	"log"
	"net"
	"sync"
	"time"
)

func NewServer() *Server {
    lis, err := net.Listen("tcp", ":8080")
    if err != nil{
        log.Fatal("Error starting the server: ",err)
    }
    return &Server{
        Conns: make([]net.Conn, 0),
        msgs: make(chan string),
        mux: sync.Mutex{},
        lis: lis,
    }
}


type Server struct{
    Conns  []net.Conn 
    msgs   chan string
    mux    sync.Mutex
    lis     net.Listener
}

func(s *Server) Broadcast(){
    log.Println("Waiting for messages to broadcast...")
    for{
        msg := <- s.msgs
        s.mux.Lock()
        for _, c := range s.Conns{
            c.Write([]byte(msg))
        }
        s.mux.Unlock()
    }
}

func(s *Server) HandleMessage(){
    activeConns := make([]net.Conn, 0)
    for{
        s.mux.Lock()
        activeConns = s.Conns
        s.mux.Unlock()
        for _, c := range activeConns{
            go s.readMessages(c)
        }
        time.Sleep(100 * time.Millisecond) // Add a small sleep to avoid CPU overuse
    }
}

func(s *Server) HandleConn(){
    for{
        conn, err := s.lis.Accept() 
        if err != nil{
            log.Println("Error connecting to a client: ", err)
            continue
        }
        s.mux.Lock()
        s.Conns = append(s.Conns, conn)
        log.Println("Number of clients: ", len(s.Conns))
        s.mux.Unlock()
        log.Println("New client connected!")
        conn.Write([]byte("Hello new guy\n"))
    }
}

func (s *Server) readMessages(c net.Conn)  {
    buf := make([]byte, 128)
    n, err := c.Read(buf)
    if err != nil{
        log.Println("Error reading from a client: ",err)
        c.Close()
    }
    log.Println("Client said something: ", string(buf[:n]))
    s.msgs <- string(buf[:n])
}


