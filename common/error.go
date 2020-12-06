package common

type ScriptErrorType int

const (
	TypeNotFoundError ScriptErrorType = iota
	SymbolError
	IndexError
	MismatchError
	TypeError
	MathError
	InitializationError
	CompileError
)

type ScriptError interface {
	Error() string
	Type() ScriptErrorType
}

type TypeNotFoundErr struct {
	errorMessage string
}

func (t *TypeNotFoundErr) Error() string {
	return "Type not found error: " + t.errorMessage
}

func (t *TypeNotFoundErr) Type() ScriptErrorType {
	return TypeNotFoundError
}

func NewTypeNotFoundErr(message string) ScriptError {
	return &TypeNotFoundErr{errorMessage: message}
}

type SymbolErr struct {
	errorMessage string
}

func (t *SymbolErr) Error() string {
	return "Symbol error: " + t.errorMessage
}

func (t *SymbolErr) Type() ScriptErrorType {
	return SymbolError
}

func NewSymbolErr(message string) ScriptError {
	return &SymbolErr{errorMessage: message}
}

type IndexErr struct {
	errorMessage string
}

func (t *IndexErr) Error() string {
	return "Index error: " + t.errorMessage
}

func (t *IndexErr) Type() ScriptErrorType {
	return IndexError
}

func NewIndexErr(message string) ScriptError {
	return &IndexErr{errorMessage: message}
}

type MismatchErr struct {
	errorMessage string
}

func (t *MismatchErr) Error() string {
	return "Mismatch error: " + t.errorMessage
}

func (t *MismatchErr) Type() ScriptErrorType {
	return MismatchError
}

func NewMismatchErr(message string) ScriptError {
	return &MismatchErr{errorMessage: message}
}

type TypeErr struct {
	errorMessage string
}

func (t *TypeErr) Error() string {
	return "Type error: " + t.errorMessage
}

func (t *TypeErr) Type() ScriptErrorType {
	return TypeError
}

func NewTypeErr(message string) ScriptError {
	return &TypeErr{errorMessage: message}
}

type MathErr struct {
	errorMessage string
}

func (t *MathErr) Error() string {
	return "Math error: " + t.errorMessage
}

func (t *MathErr) Type() ScriptErrorType {
	return MathError
}

func NewMathErr(message string) ScriptError {
	return &MathErr{errorMessage: message}
}

type InitializationErr struct {
	errorMessage string
}

func (t *InitializationErr) Error() string {
	return "Initialization error: " + t.errorMessage
}

func (t *InitializationErr) Type() ScriptErrorType {
	return InitializationError
}

func NewInitializationErr(message string) ScriptError {
	return &InitializationErr{errorMessage: message}
}

type CompileErr struct {
	errorMessage string
}

func (t *CompileErr) Error() string {
	return "Compile error: " + t.errorMessage
}

func (t *CompileErr) Type() ScriptErrorType {
	return CompileError
}

func NewCompileErr(message string) ScriptError {
	return &CompileErr{errorMessage: message}
}
