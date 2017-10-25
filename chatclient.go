package main


import "fmt"
import "net"
import "bufio"
import "os"

func main() {

    conn, _ := net.Dial("tcp", "127.0.0.1:5555")

    msg := receive(conn)
    fmt.Println(msg)

    name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    send(conn, name + "\n")

    for {
        fmt.Println("Write a message:")
        text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

        send(conn, text + "\n")

        msg := receive(conn)

        if msg == "-q\n" {
            break
        }

        fmt.Println("Message from server: ", msg)
    }

    conn.Close()

}

func send(c net.Conn, msg string) {
    c.Write([]byte(msg))
}

func receive(c net.Conn) string {
    msg, _ := bufio.NewReader(c).ReadString('\n')
    return msg
}
