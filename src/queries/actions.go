package queries

import (
	"MINIDB/src/objects"
	"strings"
)

var QueryReturn = new(objects.ActionReturn)

const (
	success         = 0
	undefined       = -1
	syntaxError     = 4
	systemError     = 5
	unKnownCommand  = 1
	objectNotFound  = 2
	incompleteQuery = 3
	unknown         = 6
)

// query handlers
func HandleShow(args []string) *objects.ActionReturn {
	object := args[1]
	switch object {
	case Collections.Type:
		// note:use the return value and print it using the repl's printer
		_, err := ShowCollections()
		if err != nil {
			QueryReturn.SetReturnValue("error", unknown, nil, true, err.Error())
		}
	case Databases.Type, DBs.Type:
		// note:use the return value and print it using the repl's printer
		ShowDBs()
	default:
		HandleInvalidArgs()
	}
	QueryReturn.SetReturnValue("success", success, nil, false, "")
	return QueryReturn
}
func HandleCreate(args []string) *objects.ActionReturn {
	object := args[1]
	identifier := args[2]
	switch object {
	case Collection.Type:
		CreateCollection(identifier)
	case Database.Type:
		CreateDB(identifier)
	default:
		HandleInvalidArgs()
	}
	return QueryReturn
}
func HandleDelete(args []string) *objects.ActionReturn {
	if len(args) < 3 {
		HandleInvalidArgs()
		return nil
	}
	object := args[1]
	identifier := args[2]
	switch object {
	case Collection.Type:
		DropCollection(identifier)
	case Database.Type:
		DropDB(identifier)
	default:
		HandleInvalidArgs()
	}
	return QueryReturn
}
func HandleRead(args []string) *objects.ActionReturn {
	return QueryReturn
}
func HandleReadOne(args []string) *objects.ActionReturn {
	return QueryReturn
}
func HandleUpdate(args []string) *objects.ActionReturn {
	return QueryReturn
}
func HandleUSe(args []string) *objects.ActionReturn {
	dbName := args[1]
	// note:use the return value to print to repl
	UseDB(dbName)
	return QueryReturn
}
func HandleInsert(args []string) *objects.ActionReturn {
	if len(args) < 4 {
		println("syntaxError invalid length")
		QueryReturn.SetReturnValue("error", incompleteQuery, nil, true, "incomplete query")
		return QueryReturn
	}
	document := args[1]
	pointer := args[2]
	collection := args[3]
	hasPref := strings.HasPrefix(document, Document.Type+"(")
	hasPost := strings.HasSuffix(document, ")")
	document = strings.TrimSpace(document)

	if !hasPost || !hasPref {
		println("syntax error prefix suffix")
		QueryReturn.SetReturnValue("error", syntaxError, nil, true, "syntax error check document(obj) syntax ")
		return QueryReturn
	}
	if pointer != Into.Value {
		println("syntax error of pointer to collection")
		QueryReturn.SetReturnValue("error", syntaxError, nil, true,
			"syntax error, check pointer \"into\" syntax ")
		return QueryReturn
	}
	if err := InsertDocument(collection, document); err != nil {
		println("system error")
		QueryReturn.SetReturnValue("error", 500, nil, true, "system error")
		return QueryReturn
	}
	QueryReturn.SetReturnValue("success", success, nil, false, "")
	return QueryReturn
}
func HandleExit(args []string) *objects.ActionReturn {
	ClearConsole()
	Exit()
	return QueryReturn
}
func HandleClear(args []string) *objects.ActionReturn {
	ClearConsole()
	return QueryReturn
}
func HandleUndefined() *objects.ActionReturn {
	println("command not found")
	return QueryReturn
}
func HandleInvalidArgs() *objects.ActionReturn {
	println("invalid arguments passed to command")
	QueryReturn.SetReturnValue("error", incompleteQuery, nil, true, "incomplete query")
	return QueryReturn
}
func HandleSystemError(ErrorMessage string) *objects.ActionReturn {
	QueryReturn.SetReturnValue("error", 500, nil, true, "system error: "+ErrorMessage)
	return QueryReturn
}
func ValidCommand(args []string) bool {
	command := args[0]
	if command == "clear" || command == "exit" {
		return true
	}
	if len(args) < 2 {
		println("command and argument needed")
		return false
	}
	return true
}
func EvaluateQuery(args []string) {
	if len(args) < 1 {
		return
	}
	function := args[0]
	for _, Token := range Functions {
		if function == Token.Name {
			if ValidCommand(args) {
				Token.Action(args)
			}
			QueryReturn.SetReturnValue("", undefined, nil, false, "")
			return
		}
	}
	HandleUndefined()
}
