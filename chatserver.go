package main

import "fmt"
import "net"
import "bufio"
import "time"
import "strings"


// TODO: Add docstrings
// TODO: Fix Problems with newline


type user struct {
    userId int
    name string
    conn net.Conn
    address net.Addr
    joined time.Time
}

var clients []user
var uid int   // id

func main() {

    fmt.Println("Server has started")
    ln, _ := net.Listen("tcp", ":5555")

    acceptClients(ln)
}

func manageClient(c net.Conn, id int) {
    send(c,"Welcome to this echo server, Whats your name:\n")

    name := receive(c)

    client := createClient(c, id, name)

    fmt.Println(client.name, "connected")

    clients = append(clients, client)

    for {
        msg := receive(c)
        if msg == "-q" {
            send(c, "-q")
            fmt.Println(client.name, " left")
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
        uid++
        go manageClient(c, uid)
    }
}

func createClient(c net.Conn, id int, name string) user {
    return user{
        userId: id,
        name: name,
        conn: c,
        address: c.RemoteAddr(),
        joined: time.Now(),
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
    c.Write([]byte(msg + "\n"))
}

func receive(c net.Conn) string {
    msg, _ := bufio.NewReader(c).ReadString('\n')

    msg = strings.Replace(msg, "\n", "", -1)

    return msg
}
