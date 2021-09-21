//exercise 8.12 - 8.15

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
    "flag"
    "time"
)

const usage = `Usage simple chat on Go:
    
    Build chat:
    $ git clone [url to github repo] && cd chat
    $ go build chat.go

    Run:
    $ ./chat &
    
    Connect to chat:
    $ nc localhost 8000

    Disconnect:
    [windows/linux] Ctrl + c
    [mac] control + c
    
    Stop:
    $ kill -9 [pid] 
    (you can see pid after ./chat & command or can use that command: 
        % ps aux | grep chat
    pid is specified in second column
    )
`

func main() {
    log.SetFlags(0)

    var (
        helpFlag = flag.Bool("help", false, "")
    )

    flag.Usage = func() {
        fmt.Fprint(flag.CommandLine.Output(), usage)
    }
    flag.Parse()
    if *helpFlag {
        fmt.Print(usage)
        return
    }

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

type client struct {
    ch chan<- string
    name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
            for c := range clients {
                cli.ch <- c.name + " is online"
            }
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
    var cli client
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
    cli.ch = ch
    cli.name = who

	ch <- "You are " + who
	messages <- who + " is connected"
	entering <- cli

    timer := time.NewTimer(15 * time.Second)
    go func() {
        <-timer.C
        conn.Close()
    }()

	input := bufio.NewScanner(conn)
	for input.Scan() {
        timer.Reset(15 * time.Second)
		messages <- who + ": " + input.Text()
	}

	leaving <- cli
	messages <- who + " is disconnected"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
