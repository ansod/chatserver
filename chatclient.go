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
    Text string
}


func main() {

    c, _ := net.Dial("tcp", "127.0.0.1:5555")

    msg := receive(c)
    fmt.Println(msg.Text)

    name, _ := bufio.NewReader(os.Stdin).ReadString('\n')

    send(c, Message{Sender: name, Text: name})

    for {
        fmt.Println("Write a message:")
        text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

        send(c, Message{name, text})

        msg := receive(c)

        if msg.Text == "-q" {
            break
        }

        fmt.Println("Message from server: ", msg.Text)
    }

    c.Close()
}

func send(c net.Conn, msg Message) {
    msg.Sender = strings.Replace(msg.Sender, "\n", "", -1)
    msg.Text = strings.Replace(msg.Text, "\n", "", -1)

    _ = gob.NewEncoder(c).Encode(msg)

    //c.Write([]byte(msg + "\n"))
}

func receive(c net.Conn) Message {
    var msg Message
    _ = gob.NewDecoder(c).Decode(&msg)


    //msg, _ := bufio.NewReader(c).ReadString('\n')

    //msg = strings.Replace(msg, "\n", "", -1)

    return msg
}
