package queries

import "MINIDB/src/objects"

// valid token types

// type Token = objects.Token

// var Operators = []*Token{&And, &Or, &All, &From}
// var Objects = []*Token{&Collection, &Collections, &Database, &DBs, &Databases}
var Functions = []*Function{&Show, &Create, &Delete, &Read, &Update, &USe, &Quit, &Clear, &Insert}

const (
	and         = "and"
	or          = "or"
	show        = "show"
	create      = "create"
	use         = "use"
	read        = "read"
	delete      = "delete"
	update      = "update"
	exit        = "exit"
	clear       = "clear"
	insert      = "insert"
	collections = "collections"
	collection  = "collection"
	database    = "database"
	databases   = "databases"
	document    = "document"
	all         = "all"
	from        = "from"
	user        = "user"
	dbs         = "dbs"
	documents   = "documents"
	into        = "into"
)

//  var keyWords  = []string {};

type Function objects.Function
type DBObject objects.DBObject
type Operator objects.Operator

type SomeInf interface {
	Class(...any) string
}

type Token objects.Token

// mapping tokens to respective actions
var And = Operator{
	Value: and,
}
var Or = Operator{
	Value: or,
}

// function Tokens
var Show = Function{
	Name:   show,
	Action: HandleShow,
}
var Create = Function{
	Name:   create,
	Action: HandleCreate,
}
var Delete = Function{
	Name:   delete,
	Action: HandleDelete,
}
var Read = Function{
	Name:   read,
	Action: HandleRead,
}
var Update = Function{
	Name:   update,
	Action: HandleUpdate,
}
var USe = Function{
	Name:   use,
	Action: HandleUSe,
}
var Quit = Function{
	Name:   exit,
	Action: HandleExit,
}
var Clear = Function{
	Name:   clear,
	Action: HandleClear,
}
var Insert = Function{
	Name:   insert,
	Action: HandleInsert,
}

// object Tokens
var Collection = DBObject{
	Type: collection,
}
var Database = DBObject{
	Type: database,
}
var Databases = DBObject{
	Type: databases,
}
var DBs = DBObject{
	Type: dbs,
}
var Collections = DBObject{
	Type: collections,
}
var Document = DBObject{
	Type: document,
}
var Documents = DBObject{
	Type: documents,
}
var Into = Operator{
	Value: into,
}

// operator Tokens
var From = Operator{
	Value: from,
}
var All = Operator{
	Value: all,
}
