package objects

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

type Token struct {
	TokenType string
	Value     string
	Action    func(args []string) *ActionReturn
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
