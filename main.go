package main

import (
	"MINIDB/src/queries"
	"MINIDB/src/repl"
)

func main() {
	repl.REPL(queries.EvaluateQuery)
}
