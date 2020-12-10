grammar goscript;

// built-in
PUSHBACK: 'pushBack';
PUSHFRONT: 'pushFront';
DELETE: 'delete';
ENUMSTRING: 'enumString';
LEN: 'len';
TYPEOF: 'typeof';

// kind
MAP: 'map';
SLICE: 'slice';
MESSAGE: 'message';
ENUM: 'enum';

// basic type name
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
CHAN: 'chan';
OBJECT: 'object';

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
CONST: 'const';

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

CHANOP: '<-';
CHANOPNONBLOCK: '<<-';

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
TAILARRAY : '...';

WHITESPACE: [ \r\n\t]+ -> skip;
COMMENT :  '//' ~( '\r' | '\n' )* ( '\r' | '\n' ) -> skip;

// 1
program
    : (functiondef|typedef|execution)+;

functiondef
    : FUNCTION NAME '(' (inparam (',' inparam)* (TAILARRAY)?)?  ')' returntypename?  closure # FunctionDef
    | FUNCTION NAME '(' (inparam (',' inparam)* (TAILARRAY)?)?  ')' '('returntypename (',' returntypename) *')' closure # FunctionDef
    | FUNCTION NAME '(' (inparam (',' inparam)* (TAILARRAY)?)?  ')' outparam?  closure # FunctionDef
    | FUNCTION NAME '(' (inparam (',' inparam)* (TAILARRAY)?)?  ')' '('outparam (',' outparam) *')' closure # FunctionDef
    ;

lambda
    : FUNCTION  '(' (inparam (',' inparam)* (TAILARRAY)?)?  ')' returntypename?  closure # LambdaDef
    | FUNCTION  '(' (inparam (',' inparam)* (TAILARRAY)?)?  ')' '('returntypename (',' returntypename) *')' closure # LambdaDef
    | FUNCTION  '(' (inparam (',' inparam)* (TAILARRAY)?)?  ')' outparam?  closure # LambdaDef
    | FUNCTION  '(' (inparam (',' inparam)* (TAILARRAY)?)?  ')' '('outparam (',' outparam) *')' closure # LambdaDef
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

typename
    : (NAME|basicTypeName) # SimpleTypeName
    | functionTypeName # FunctionType
    | MAP '<' basicTypeName ',' typename '>' # MapTypeName
    | SLICE '<' typename '>' # SliceTypeName
    | CHAN '<' typename '>' # ChanTypeName
    ;

functionTypeName
    : FUNCTION  '(' (intypename (',' intypename)* (TAILARRAY)?)?  ')' returntypename?
    | FUNCTION  '(' (intypename (',' intypename)* (TAILARRAY)?)?  ')' '('returntypename (',' returntypename) *')'
    ;

typedef
    : TYPEDEF NAME MAP '<'  basicTypeName ',' typenameindef '>'# TypeDefMap
    | TYPEDEF NAME SLICE '<' typenameindef '>'# TypeDefSlice
    | TYPEDEF NAME MESSAGE '{' (messagefield (messagefield)*)? '}' # TypeDefMessage
    | TYPEDEF NAME ENUM '{' (NAME ':' INT)* '}' # TypeDefEnum
    | TYPEDEF NAME functionTypeNameindef # TypeDefFunction
    ;

messagefield
    : NAME typenameindef;

typenameindef
    : (NAME|basicTypeName) # SimpleTypeNameInDef
    | functionTypeNameindef # FunctionTypeInDef
    | MAP '<' basicTypeName ',' typenameindef '>' # MapTypeNameInDef
    | SLICE '<' typenameindef '>' # SliceTypeNameInDef
    | CHAN '<' typenameindef '>' # ChanTypeNameInDef
    ;

functionTypeNameindef
    : FUNCTION  '(' (intypenameindef (',' intypenameindef)* (TAILARRAY)?)?  ')' returntypenameindef?
    | FUNCTION  '(' (intypenameindef (',' intypenameindef)* (TAILARRAY)?)?  ')' '('returntypenameindef (',' returntypenameindef) *')'
    ;

intypenameindef
    : typenameindef;

returntypenameindef
    : typenameindef;

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
    | vardef # VarDef
    | constdef # ConstDef
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
    | variable '(' (expr (',' expr)*)? ')' # DirectCall
    | variable DOT '(' asserted ')' # TypeAssert
    | NAME # VariableName
    | '@' # VariableName
    ;

asserted: typename;

filter: expr;

indexs
    : expr ':' expr ':' expr # IndexType1
    | expr ':' expr # IndexType2
    | expr ':' # IndexType3
    | expr # IndexType4
    | ':' expr # IndexType5
    ;

expr
    : '(' expr ')' # Pass
    | constant # Pass
    | variable # Pass
    | lambda # Pass
    | builtin # Pass
    | op=(UNARYADD|UNARYSUB) variable # LeftUnary
    | op=(NOT|SUB) expr # LeftUnary
    | variable op=(UNARYADD|UNARYSUB) # RightUnary
    | <assoc=right> expr op=POW  expr # Binary
    | expr op=(MUL | DIV | MOD) expr # Binary
    | expr op=(ADD | SUB) expr # Binary
    | expr op=(EQ | INEQ | GT | GE | LT | LE | REGEX) expr # Binary
    | expr op=AND expr # Binary
    | expr op=OR expr # Binary
    | variable (CHANOP|CHANOPNONBLOCK) variable # Send
    | (CHANOP|CHANOPNONBLOCK) variable # Recv
    | <assoc=right> lhs op=(ASSIGN|ADDEQUAL|SUBEQUAL|MULEQUAL|DIVEQUAL) expr # Binary
    | <assoc=right> lhs op=ASSIGN initializationListBegin # AssignInitializationlist
    | constructor # Construct
    ;

basicTypeName
    : (UINT32|UINT64|INT32|INT64|FLOAT32|FLOAT64|STRING|BYTES|BOOL|UINT8|OBJECT);

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
    : '[' (initializationList (',' initializationList)*)? ']' # InitSlice
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
    | VAR NAME typename '=' initializationListBegin
    | LOCAL NAME typename
    | LOCAL NAME typename '=' expr
    | LOCAL NAME typename '=' initializationListBegin
    ;

constdef
    : CONST NAME basicTypeName '=' constant
    ;