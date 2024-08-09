package objects

import "reflect"

type Meta struct {
	CreatedAt string
	CreatedBy string
}
type DATABASE struct {
	Name        string
	Path        string
	Collections []*COLLECTION
	Meta
}

type COLLECTION struct {
	Name     string
	Path     string
	Database *DATABASE
	Meta
}

type ConfigFile struct {
	Storage struct {
		DataPath string `json:"dataPath"`
	} `json:"storage"`
	Logs struct {
		Path string `json:"path"`
	} `json:"logs"`
	Net struct {
		Ip   string `json:"ip"`
		Port string `json:"port"`
	} `json:"net"`
}

// type Token struct {
// 	TokenType string
// 	Value     string
// 	Action    func(args []string) *ActionReturn
// }

type Token interface {
	getValue() string
}

type Keyword interface {
	getType() string
}

type Function struct {
	Name   string                       // the name of the function
	Action func([]string) *ActionReturn // what the function does
}

// implement the Token interface
func (f Function) getValue() string {
	return f.Name
}

// implement the keyword interface
func (f Function) getType() string {
	return reflect.TypeOf(f).String()
}

type DBObject struct {
	Type string
}

// implement the keyWord interface
func (o DBObject) getType() string {
	return reflect.TypeOf(o).String()
}

// implement the Token interface
func (o DBObject) getValue() string {
	return o.Type
}

type Operator struct {
	Value string
}

// implement Token interface
func (o Operator) getValue() string {
	return o.Value
}

type UserError struct {
	Type    string
	Root    error
	Message string
}

func (e *UserError) Error() string {
	return e.Message
}

type ActionReturn struct {
	Error        bool
	ErrorMessage string
	Status       string
	ExitCode     int
	ReturnValue  *any
}

func (actionReturn *ActionReturn) SetReturnValue(status string, exitCode int, returnValue *any, Error bool, ErrorMessage string) {
	actionReturn.ExitCode = exitCode
	actionReturn.ReturnValue = returnValue
	actionReturn.Status = status
	actionReturn.Error = Error
	actionReturn.ErrorMessage = ErrorMessage
}
