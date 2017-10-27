package main

import (
    "fmt"
    "net"
    "bufio"
    "time"
    "strings"
    "os"
    "encoding/gob"
)


// TODO: Add docstrings


type User struct {
    userId int
    name string
    conn net.Conn
    address net.Addr
    joined string
}

type Message struct {
    Sender string
    Text string
}


var clients []User
var uid int   // id

func main() {

    fmt.Println("Server has started")
    ln, _ := net.Listen("tcp", ":5555")

    go acceptClients(ln)
    manageServer()

}

func manageServer() {
    var cmd string

    for {
        cmd, _ = bufio.NewReader(os.Stdin).ReadString('\n')
        cmd = strings.Replace(cmd, "\n", "", -1)

        // TODO: Send -q to all clients
        switch cmd {
            case "-q": return
            case "-help": fmt.Println(getCommands())
            case "-users": fmt.Println(getUserInfo())
            default: fmt.Println("Unknown command")
        }

    }
}

func manageClient(c net.Conn, id int) {
    send(c, Message{"server", "Welcome to this echo server, Whats your name:"})

    name := receive(c)

    client := createClient(c, id, name.Sender)

    fmt.Println(client.name, "connected")

    clients = append(clients, client)

    for {
        msg := receive(c)
        if msg.Text == "-q" {
            send(c, Message{"server", "-q"})
            fmt.Println(client.name, " left")
            removeClient(c)
            c.Close()
            return
        }
        fmt.Println(msg.Sender + ": " + msg.Text)
        send(c, Message{"server", msg.Text})
    }
}

func acceptClients(ln net.Listener) {
    for {
        c, _ := ln.Accept()
        uid++
        go manageClient(c, uid)
    }
}

func createClient(c net.Conn, id int, name string) User {
    time := strings.Split(time.Now().String(), ".")[0]
    return User{
        userId: id,
        name: name,
        conn: c,
        address: c.RemoteAddr(),
        joined: time,
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

func send(c net.Conn, msg Message) {
    _ = gob.NewEncoder(c).Encode(msg)

    //c.Write([]byte(msg + "\n"))
}

func receive(c net.Conn) Message {
    var msg Message
    _ = gob.NewDecoder(c).Decode(&msg)

    //msg, _ := bufio.NewReader(c).ReadString('\n')

    //msg.Text = strings.Replace(msg.Text, "\n", "", -1)

    return msg
}

// TODO: Create func
func sendToAll(msg string, user User) {

}

func getUserInfo() string {
    str := "<List of current users>\n"
    for _, user := range clients {
        str += fmt.Sprintf("%d: %s, %s, (%s)\n", user.userId, user.name, user.address, user.joined)
    }

    return str
}

func getCommands() string {
    return "<List of commands>\n-q: Kill server\n-users: Get user info\n"
}
