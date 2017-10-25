package main

import "fmt"
import "net"
import "bufio"
import "time"


// TODO: Add docstrings


type user struct {
    userId int
    name string
    conn net.Conn
    address net.Addr
    joined Time
}

var clients []user

func main() {

    fmt.Println("Server has started")
    ln, _ := net.Listen("tcp", ":5555")

    acceptClients(ln)
}

func manageClient(c net.Conn) {
    send(c,"Welcome to this echo server, Whats your name:\n")

    name := receive(c)

    fmt.Println(name, "just connected")

    client := createClient(c, id, name)

    clients = append(clients, client)

    for {
        msg := receive(c)
        if msg == "-q\n" {
            send(c, "-q")
            fmt.Println(c.RemoteAddr(), " left")
            removeClient(c)
            c.Close()
            return
        }
        fmt.Println("Message received from client: ", msg)
        send(c, msg)
    }
}

func acceptClients(ln net.Listener) {
    for {
        c, _ := ln.Accept()
        fmt.Println(c.RemoteAddr(), " just connected")
        go manageClient(c)
    }
}

func createClient(c net.Conn, id int, name string) user {
    return []user{
        userId: id,
        name: name,
        conn: c,
        address: c.RemoteAddr(),
        joined: time.Now()
    }
}

func removeClient(c net.Conn) {
    for i, client := range clients {
        if client.address == c.RemoteAddr() {
            clients = append(clients[:i], clients[i+1:]...)
            return
        }
    }
}

// TODO: Create func
func sendToAll() {

}

// TODO: Create func
func getClientsInfo() {

}

func send(c net.Conn, msg string) {
    c.Write([]byte(msg))
}

func receive(c net.Conn) string {
    msg, _ := bufio.NewReader(c).ReadString('\n')

    return msg
}
