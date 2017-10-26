package main

import "fmt"
import "net"
import "bufio"
import "time"
import "strings"
import "os"


// TODO: Add docstrings


type user struct {
    userId int
    name string
    conn net.Conn
    address net.Addr
    joined string
}

var clients []user
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

        //fmt.Println("You entered: ", cmd)

        switch cmd {
            case "-q": return
            case "-help": fmt.Println(getCommands())
            case "-users": fmt.Println(getUserInfo())
            default: fmt.Println("Unknown command")
        }

    }
}

func manageClient(c net.Conn, id int) {
    send(c,"Welcome to this echo server, Whats your name:")

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
    time := strings.Split(time.Now().String(), ".")[0]
    return user{
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

func send(c net.Conn, msg string) {
    c.Write([]byte(msg + "\n"))
}

func receive(c net.Conn) string {
    msg, _ := bufio.NewReader(c).ReadString('\n')

    msg = strings.Replace(msg, "\n", "", -1)

    return msg
}

// TODO: Create func
func sendToAll() {

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
