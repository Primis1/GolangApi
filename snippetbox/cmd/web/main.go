// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"flag"
// 	"log"
// 	"net"
// 	"os"
// )

// func main() {

// 	const name string = "writeShit"

// 	log.SetPrefix(name + "\t")

// 	port := flag.Int("p", 8080, "port to connect to")

// 	flag.Parse()

// 	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{Port: *port})

// 	if err != nil {
// 		log.Fatalf("error bla bla bla, to %d %v",*port, err)
// 	}
// 	log.Printf("connected to %d", conn.RemoteAddr())

// 	defer conn.Close()

// 	go func() {
// 		for connScanner := bufio.NewScanner(conn); connScanner.Scan(); {
// 			fmt.Printf("%s\n", connScanner.Text())

// 			if err := connScanner.Err(); err != nil {
// 				log.Fatalf("Error reading from %s:  %v ", conn.RemoteAddr(), err)
// 			}
// 			if  connScanner.Err() != nil {
// 				log.Fatalf("Error reading from %s:  %v ", conn.RemoteAddr(), err)
// 			}
// 		}
// 	}()

// 	for stdinScanner := bufio.NewScanner(os.Stdin); stdinScanner.Scan(); {
// 		log.Printf("Error: in for loop, %s\n", stdinScanner.Text())
// 		if _, err := conn.Write(stdinScanner.Bytes()); err != nil {
// 			log.Fatalf("Error writing to %s: %v", conn.RemoteAddr(), err)
// 		}
// 		if _, err = conn.Write([]byte("\n")); err != nil {
// 			log.Fatalf("Error writing to %s: %v", conn.RemoteAddr(), err)
// 		}

// 		if stdinScanner.Err() != nil {
// 			log.Fatalf("Error scanned of %s: %v", conn.RemoteAddr(), err)
// 		}

// 	}
// }

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(data int) {
    newNode := &Node{data: data}

    if list.head == nil {
        list.head = newNode
    } else {
        current := list.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
}

func (list *LinkedList) Display() {
    current := list.head

    if current == nil {
        fmt.Println("Linked list is empty")
        return
    }

    fmt.Print("Linked list: ")
    for current != nil {
        fmt.Printf("%d ", current.data)
        current = current.next
    }
    fmt.Println()
}

func main() {
    list := LinkedList{}

    list.Insert(10)
    list.Insert(20)
    list.Insert(30)
    list.Insert(40)

    list.Display()
}
