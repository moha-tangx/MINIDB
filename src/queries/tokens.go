package queries

import "MINIDB/src/objects"

// valid token types
const (
	function   = "function"
	object     = "object"
	aggregator = "aggregator"
	operator   = "operator"
	identifier = "identifier"
)

type Keyword interface {
}

// var operators = []*Token{&And, &Or}
// var aggregators = []*Token{&All, &From}
// var dbObjects = []*Token{&Collection, &Collections, &Database, &DBs, &Databases}

var functions = []*Token{&Show, &Create, &Delete, &Read, &Update, &USe, &Quit, &Clear, &Insert}

type SomeInf interface {
	Class(...any) string
}

type Token objects.Token

// mapping tokens to respective actions
var And = Token{
	TokenType: operator,
	Value:     "and",
}
var Or = Token{
	TokenType: operator,
	Value:     "or",
}

// function Tokens
var Show = Token{
	TokenType: function,
	Value:     "show",
	Action:    HandleShow,
}
var Create = Token{
	TokenType: function,
	Value:     "create",
	Action:    HandleCreate,
}
var Delete = Token{
	TokenType: function,
	Value:     "delete",
	Action:    HandleDelete,
}
var Read = Token{
	TokenType: function,
	Value:     "read",
	Action:    HandleRead,
}
var Update = Token{
	TokenType: function,
	Value:     "update",
	Action:    HandleUpdate,
}
var USe = Token{
	TokenType: function,
	Value:     "use",
	Action:    HandleUSe,
}
var Quit = Token{
	TokenType: function,
	Value:     "exit",
	Action:    HandleExit,
}
var Clear = Token{
	TokenType: function,
	Value:     "clear",
	Action:    HandleClear,
}
var Insert = Token{
	TokenType: function,
	Value:     "insert",
	Action:    HandleInsert,
}

// object Tokens
var Collection = Token{
	TokenType: object,
	Value:     "collection",
	Action:    nil,
}
var Database = Token{
	TokenType: object,
	Value:     "database",
	Action:    nil,
}
var Databases = Token{
	TokenType: object,
	Value:     "databases",
	Action:    nil,
}
var DBs = Token{
	TokenType: object,
	Value:     "dbs",
	Action:    nil,
}
var Collections = Token{
	TokenType: object,
	Value:     "collections",
	Action:    nil,
}
var Document = Token{
	TokenType: object,
	Value:     "document",
	Action:    nil,
}
var Documents = Token{
	TokenType: object,
	Value:     "documents",
	Action:    nil,
}
var Into = Token{
	TokenType: aggregator,
	Value:     "into",
	Action:    nil,
}

// operator Tokens
var From = Token{
	TokenType: aggregator,
	Value:     "from",
	Action:    nil,
}
var All = Token{
	TokenType: aggregator,
	Value:     "all",
	Action:    nil,
}
