package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

const quote = "The greatest glory in living lies not in never falling, but in rising every time we fall."

const difficulty = 4

func main() {
	fmt.Println("Starting Word of Wisdom TCP server...")
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Client connected.")
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	challenge := generateChallenge()
	fmt.Println("Challenge:", challenge)
	conn.Write([]byte(strconv.Itoa(challenge) + "\n"))
	var buf [1024]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("Error reading from client:", err)
		return
	}
	solution := string(buf[:n])
	fmt.Println("Solution:", solution)
	if verifyChallenge(challenge, solution) {
		conn.Write([]byte(getQuote() + "\n"))
	} else {
		conn.Write([]byte("Error: invalid solution.\n"))
	}
}

func generateChallenge() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(99999)
}

func verifyChallenge(challenge int, solution string) bool {
	hash := sha256.Sum256([]byte(strconv.Itoa(challenge) + solution))
	hexHash := fmt.Sprintf("%x", hash)
	for i := 0; i < difficulty; i++ {
		if hexHash[i] != '0' {
			return false
		}
	}
	return true
}

func getQuote() string {
	return quote
}
