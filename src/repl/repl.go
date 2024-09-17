package repl

import (
	"MINIDB/src/queries"
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type EVALUATED struct {
	ExitCode     int
	Error        bool
	ErrorMessage string
	Message      string
	ReturnValue  []string
}

func READ(c chan<- string) {
	var READER = bufio.NewReader(os.Stdin)
	for {
		repl_name := ""
		db, err := queries.GetDBInUse()
		if err != nil {
			fmt.Println("no database in")
			panic(err)
		}
		if db != nil {
			repl_name = db.Name
		}
		fmt.Printf("%s >> ", repl_name)
		input, _ := READER.ReadString('\n')
		c <- input
		time.Sleep(time.Millisecond * 200)
	}
}

func ExitRepl(evaluated chan<- *EVALUATED) {
	println("exiting...")
	close(evaluated)
	os.Exit(0)
}

func EVALUATE(inputs <-chan string, evalFunc func(args []string), evaluated chan<- *EVALUATED) {
	// loop listening to the channel to send input
	for input := range inputs {
		input = strings.TrimSpace(input)
		unSanitizedTokens := strings.Split(input, " ")
		tokens := []string{}
		for _, token := range unSanitizedTokens {
			if strings.TrimSpace(token) != "" {
				tokens = append(tokens, token)
			}
		}
		if evalFunc != nil {
			evalFunc(tokens)
		}
		evaluated <- &EVALUATED{}
	}
}

func PRINT(c <-chan *EVALUATED) {
	//loop listening to the channel to send evaluated input and print it out
	for evaluated := range c {
		println(evaluated.Message)
		if evaluated.Error {
			println("\033[m31", evaluated.ErrorMessage, "\033[m0")
		} else {
			for _, v := range evaluated.ReturnValue {
				println(v)
			}
		}
	}
}

func REPL(evalFunc func(args []string)) {
	input := make(chan string)
	evaluated := make(chan *EVALUATED)

	var waiter sync.WaitGroup
	waiter.Add(1)

	go READ(input)
	go EVALUATE(input, evalFunc, evaluated)
	go PRINT(evaluated)

	waiter.Wait()
}
