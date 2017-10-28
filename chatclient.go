package main

import (
    "fmt"
    "net"
    "bufio"
    "os"
    "strings"
    "encoding/gob"
)

// TODO: Add docstrings

type Message struct {
    Sender string
    Type string
    Text string
}


func main() {

    c, _ := net.Dial("tcp", "127.0.0.1:5555")

    receive(c)

    name, _ := bufio.NewReader(os.Stdin).ReadString('\n')

    send(c, Message{Sender: name, Type: "welcome", Text: name})


    go chat(c, name)
    receive(c)

    c.Close()
}

func chat(c net.Conn, name string) {

    for {
        fmt.Println("Write a message:")
        text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

        send(c, Message{name, "chat", text})
    }
}

func send(c net.Conn, msg Message) {
    msg.Sender = strings.Replace(msg.Sender, "\n", "", -1)
    msg.Text = strings.Replace(msg.Text, "\n", "", -1)

    _ = gob.NewEncoder(c).Encode(msg)
}

func receive(c net.Conn) {
    for {
        var msg Message
        _ = gob.NewDecoder(c).Decode(&msg)

        switch msg.Type {
        case "chat":
            fmt.Println(msg.Sender + ": " + msg.Text)
        case "welcome":
            fmt.Println(msg.Sender + ": " + msg.Text)
            return
        case "quit":
            return
        default:
            fmt.Println("subject:", msg.Type)
        }
    }
}
