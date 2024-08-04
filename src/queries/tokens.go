package queries

import "MINIDB/src/objects"

// valid token types
const (
	Function   = "function"
	Object     = "object"
	Aggregator = "aggregator"
	Operator   = "operator"
	Identifier = "identifier"
)

var Operators = []*Token{&And, &Or}
var Aggregators = []*Token{&All, &From}
var Objects = []*Token{&Collection, &Collections, &Database, &DBs, &Databases}
var Functions = []*Token{&Show, &Create, &Delete, &Read, &Update, &USe, &Quit, &Clear, &Insert}

type SomeInf interface {
	Class(...any) string
}

type Token objects.Token

// mapping tokens to respective actions
var And = Token{
	TokenType: Operator,
	Value:     "and",
}
var Or = Token{
	TokenType: Operator,
	Value:     "or",
}

// function Tokens
var Show = Token{
	TokenType: Function,
	Value:     "show",
	Action:    HandleShow,
}
var Create = Token{
	TokenType: Function,
	Value:     "create",
	Action:    HandleCreate,
}
var Delete = Token{
	TokenType: Function,
	Value:     "delete",
	Action:    HandleDelete,
}
var Read = Token{
	TokenType: Function,
	Value:     "read",
	Action:    HandleRead,
}
var Update = Token{
	TokenType: Function,
	Value:     "update",
	Action:    HandleUpdate,
}
var USe = Token{
	TokenType: Function,
	Value:     "use",
	Action:    HandleUSe,
}
var Quit = Token{
	TokenType: Function,
	Value:     "exit",
	Action:    HandleExit,
}
var Clear = Token{
	TokenType: Function,
	Value:     "clear",
	Action:    HandleClear,
}
var Insert = Token{
	TokenType: Function,
	Value:     "insert",
	Action:    HandleInsert,
}

// object Tokens
var Collection = Token{
	TokenType: Object,
	Value:     "collection",
	Action:    nil,
}
var Database = Token{
	TokenType: Object,
	Value:     "database",
	Action:    nil,
}
var Databases = Token{
	TokenType: Object,
	Value:     "databases",
	Action:    nil,
}
var DBs = Token{
	TokenType: Object,
	Value:     "dbs",
	Action:    nil,
}
var Collections = Token{
	TokenType: Object,
	Value:     "collections",
	Action:    nil,
}
var Document = Token{
	TokenType: Object,
	Value:     "document",
	Action:    nil,
}
var Documents = Token{
	TokenType: Object,
	Value:     "documents",
	Action:    nil,
}
var Into = Token{
	TokenType: Aggregator,
	Value:     "into",
	Action:    nil,
}

// operator Tokens
var From = Token{
	TokenType: Aggregator,
	Value:     "from",
	Action:    nil,
}
var All = Token{
	TokenType: Aggregator,
	Value:     "all",
	Action:    nil,
}
