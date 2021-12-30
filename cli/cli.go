package main

import (
	"bufio"
	"flag"
	"fmt"
	c "kv/client"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

// commands
const (
	Insert string = "insert"
	Lookup        = "lookup"
	Delete        = "delete"
	Help          = "help"
	Clear         = "clear"
	Exit          = "exit"
)

var port *int = flag.Int("port", 20020, "port")

// interactive repl functions
func prompt() {
	fmt.Print("kv> ")
}

func help() {
	fmt.Println("available commands:")
	fmt.Println("insert  - Insert key and value into kv")
	fmt.Println("lookup  - Lookup key in kv")
	fmt.Println("delete  - Delete key in kv")
	fmt.Println("help    - Show available commands")
	fmt.Println("clear   - Clear the terminal screen")
	fmt.Println("exit    - Close your kv client")
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func insert(args []string) {
	if len(args) != 3 {
		fmt.Printf("insert command requires 3 arguments, found %d\n", len(args))
		return
	}
	key, value := args[1], args[2]
	c.Insert(key, value)
}

func lookup(args []string) {
	if len(args) != 2 {
		fmt.Printf("lookup command requires 2 arguments, found %d\n", len(args))
		return
	}
	key := args[1]
	value := c.Lookup(key)
	if value != "" {
		fmt.Printf("%s\n", value)
	}
}

func delete(args []string) {
	if len(args) != 2 {
		fmt.Printf("delete command requires 2 arguments, found %d\n", len(args))
		return
	}
	key := args[1]
	c.Delete(key)
}

func tokenize(text string) []string {
	output := strings.Split(strings.TrimSpace(text), " ")
	return output
}

func processCommand(args []string) {
	cmd := args[0]
	if cmd == "" {
		return
	}
	switch cmd {
	case Insert:
		insert(args)
	case Lookup:
		lookup(args)
	case Delete:
		delete(args)
	case Help:
		help()
	case Clear:
		clear()
	case Exit:
		fmt.Println()
		fmt.Println("exiting")
		os.Exit(0)
	default:
		fmt.Printf("%s: command not found\n", cmd)
	}
}

func repl() {
	reader := bufio.NewScanner(os.Stdin)
	prompt()
	for reader.Scan() {
		processCommand(tokenize(reader.Text()))
		prompt()
	}
	fmt.Println()
}

func setupGracefulExits() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println()
		fmt.Println("gracefully shutting down")
		os.Exit(0)
	}()
}

func main() {
	setupGracefulExits()
	flag.Parse()
	log.SetOutput(os.Stdout)
	c.Connect(port)
	repl()

}
