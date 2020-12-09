grammar goscript;

// built-in
PUSHBACK: 'pushBack';
PUSHFRONT: 'pushFront';
DELETE: 'delete';
ENUMSTRING: 'enumString';
LEN: 'len';
TYPEOF: 'typeof';
UINT32: 'uint32';
UINT64: 'uint64';
INT32: 'int32';
INT64: 'int64';
FLOAT32: 'float32';
FLOAT64: 'float64';
STRING: 'string';
BYTES: 'bytes';
BOOL: 'bool';
UINT8: 'uint8';

NEW: 'new';

FOR: 'for';
BREAK: 'break';
CONTINUE: 'continue';

IF: 'if';
ELSE: 'else';

SWITCH: 'switch';
CASE: 'case';

RETURN: 'return';

VAR: 'var';
LOCAL: 'local';

FUNCTION: 'func';

TYPEDEF: 'type';

BOOLLITERAL
   : 'true'
   | 'false'
   ;

NULL
   : 'nil'
   ;

POW: '**';
MUL: '*';
DIV: '/';
MOD: '%';
ADD: '+';
SUB: '-';
UNARYADD: '++';
UNARYSUB: '--';

EQ: '==';
INEQ: '!=';
GT: '>';
GE: '>=';
LE: '<=';
LT: '<';
REGEX: '=~';

AND: '&&';
OR: '||';
NOT: '!';

ASSIGN: '=';
ADDEQUAL: '+=';
SUBEQUAL: '-=';
MULEQUAL: '*=';
DIVEQUAL: '/=';

INT
   : [0-9]+;

FLOAT
   : ([0-9]+)'.'[0-9]+;

STRINGLITERAL
   : '\'' ( ~('\''|'\\') | ('\\' .) )* '\'';

NAME: [a-zA-Z_]+[a-zA-Z0-9_]*;

DOT: '.';
WHITESPACE: [ \r\n\t]+ -> skip;
COMMENT :  '//' ~( '\r' | '\n' )* ( '\r' | '\n' ) -> skip;

// 1
program
    : statement+;
// 2
statement
    : functiondef # FuncDef
    | typedef ';' # TypeDef
    | enumdef # EnumDef
    | execution # Run
    ;

functiondef
    : FUNCTION NAME '(' inparam (',' inparam)*  ')' returntypename?  closure # FunctionDef
    | FUNCTION NAME '(' ')' returntypename?  block # FunctionDef
    | FUNCTION NAME '(' inparam (',' inparam)*  ')' '('returntypename (',' returntypename) *')' closure # FunctionDef
    | FUNCTION NAME '(' ')' '('returntypename (',' returntypename) *')'  closure # FunctionDef
    | FUNCTION NAME '(' inparam (',' inparam)*  ')' outparam?  closure # FunctionDef
    | FUNCTION NAME '(' ')' typename?  closure # FunctionDef
    | FUNCTION NAME '(' inparam (',' inparam)*  ')' '('outparam (',' outparam) *')' closure # FunctionDef
    | FUNCTION NAME '(' ')' '('outparam (',' outparam) *')'  closure # FunctionDef
    ;

lambda
    : FUNCTION  '(' inparam (',' inparam)*  ')' returntypename?  closure # LambdaDef
    | FUNCTION  '(' ')' returntypename?  closure # LambdaDef
    | FUNCTION  '(' inparam (',' inparam)*  ')' '('returntypename (',' returntypename) *')' closure # LambdaDef
    | FUNCTION  '(' ')' '('returntypename (',' returntypename) *')'  closure # LambdaDef
    | FUNCTION  '(' inparam (',' inparam)*  ')' outparam?  closure # LambdaDef
    | FUNCTION  '(' ')' typename?  closure # LambdaDef
    | FUNCTION  '(' inparam (',' inparam)*  ')' '('outparam (',' outparam) *')' closure # LambdaDef
    | FUNCTION  '(' ')' '('outparam (',' outparam) *')'  closure # LambdaDef
    ;

closure
    : block;

inparam
    : param;

outparam
    : param;

intypename
    : typename;

returntypename
    : typename;

param
    : NAME typename;

typedef
    : TYPEDEF NAME typename # TypeDefAlias
    | TYPEDEF NAME '{' (NAME typename ',')* '}' # TypeDefComplex
    ;

typename
    : NAME # SimpleTypeName
    | basicTypeName # SimpleTypeName
    | functionTypeName # FunctionType
    | 'map' '<' (NAME|basicTypeName) ',' typename '>' # MapTypeName
    | 'slice' '<' typename '>' # SliceTypeName
    ;

functionTypeName
    : FUNCTION  '(' intypename (',' intypename)*  ')' returntypename?
    | FUNCTION  '(' ')' returntypename?
    | FUNCTION  '(' intypename (',' intypename)*  ')' '('returntypename (',' returntypename) *')'
    | FUNCTION  '(' ')' '('returntypename (',' returntypename) *')'
    ;

enumdef
    : TYPEDEF 'enum' NAME '{' (NAME INT ',')* '}';

execution
    : control # Ctrl
    | line ';' # LineProgram
    ;

control
    : IF '(' expr ')' block (ELSE (block|control))? # If
    | SWITCH '(' expr ')' '{' (CASE constant ':' block)+ '}' # Switch
    | FOR '(' NAME 'in' collection ')' block # ForInSlice
    | FOR '(' NAME ',' NAME 'in' collection ')' block  # ForInMap
    | FOR '(' line ';' expr ';' restoreStack ')' block  # For
    | BREAK ';' # Break
    | CONTINUE ';' # Continue
    | RETURN ';'# ReturnVoid
    | RETURN expr (',' expr)*';' # ReturnVal
    ;

collection
    : expr;

block
    : '{' (execution)* '}';

line
    : restoreStack # RestoreStackSp
    | vardef # GlobalVarDef
    | localdef # LocalVarDef
    ;

restoreStack
    : keepStack;

keepStack
    : expr # ExprEntry
    | lhs (',' lhs)* op=ASSIGN variable # FunctionAssign
    | lhs (',' lhs)* op=ASSIGN expr (',' expr)* # MultiAssign
    ;

lhs
    : variable;

variable
    : variable DOT NAME # Select
    | variable '[?(' filter ')]' # SliceFilter
    | variable '[' expr ']' # Index
    | variable '[' indexs (',' indexs)* ']' # SliceMultiIndex
    | variable '[' '[' expr (',' expr)* ']' ']' # MapMultiIndex
    | variable '(' expr (',' expr)* ')' # DirectCall
    | variable '(' ')' # DirectCall
    | NAME # VariableName
    | '@' # VariableName
    ;

filter
    : expr;

indexs
    : expr ':' expr ':' expr # IndexType1
    | expr ':' expr # IndexType2
    | expr ':' # IndexType3
    | expr # IndexType4
    | ':' expr # IndexType5
    ;

expr
    : '(' expr ')' # Pass
    | op=(UNARYADD|UNARYSUB) variable # LeftUnary
    | op=(NOT|SUB) expr # LeftUnary
    | variable op=(UNARYADD|UNARYSUB) # RightUnary
    | <assoc=right> expr op=POW  expr # Binary
    | expr op=(MUL | DIV | MOD) expr # Binary
    | expr op=(ADD | SUB) expr # Binary
    | expr op=(EQ | INEQ | GT | GE | LT | LE | REGEX) expr # Binary
    | expr op=AND expr # Binary
    | expr op=OR expr # Binary
    | <assoc=right> lhs op=(ASSIGN|ADDEQUAL|SUBEQUAL|MULEQUAL|DIVEQUAL) expr # Binary
    | <assoc=right> lhs op=ASSIGN initializationListBegin # AssignInitializationlist
    | constant # Pass
    | variable # Pass
    | lambda # Pass
    | builtin # Pass
    | constructor # Construct
    ;

basicTypeName
    : UINT32
    | UINT64
    | INT32
    | INT64
    | FLOAT32
    | FLOAT64
    | STRING
    | BYTES
    | BOOL
    | UINT8
    ;

builtin
    : PUSHBACK '(' variable',' expr ')'
    | PUSHFRONT '(' variable',' expr ')'
    | DELETE '(' variable',' expr ')'
    | ENUMSTRING '(' variable ')'
    | LEN '(' variable ')'
    | TYPEOF '(' variable ')'
    | UINT32 '(' expr ')'
    | UINT64 '(' expr ')'
    | INT32 '(' expr ')'
    | INT64 '(' expr ')'
    | FLOAT32 '(' expr ')'
    | FLOAT64 '(' expr ')'
    | STRING '(' expr ')'
    | BYTES '(' expr ')'
    | BOOL '(' expr ')'
    | UINT8 '(' expr ')'
    ;

initializationListBegin
    : initializationList;

initializationList
    : '[' initializationList (',' initializationList)* ']' # InitSlice
    | '{'  NAME '('initializationList')' (',' NAME '(' initializationList ')' )* '}' # InitMessage
    | '{'  initializationList ':' initializationList (',' initializationList ':' initializationList )* '}' # InitMap
    | expr # InitConstant
    ;

constant
    : INT # ConstantInt
    | FLOAT # ConstantFloat
    | BOOLLITERAL # ConstantBool
    | NULL # ConstantNil
    | STRINGLITERAL # ConstantString
    ;

constructor
    : NEW typename '(' ')'
    | NEW typename '('expr (',' expr)*')'
    ;

vardef
    : VAR NAME typename
    | VAR NAME typename '=' expr
    ;

localdef
    : LOCAL NAME typename
    | LOCAL NAME typename '=' expr
    ;