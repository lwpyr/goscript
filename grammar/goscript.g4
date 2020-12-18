grammar goscript;

TYPEDEF: 'type';
MESSAGE: 'message';
ENUM: 'enum';
MAP: 'map';
ONEOF: 'oneof';
NEW: 'new';

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
ANY: 'any';

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
LOCALASSIGN: ':=';

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
    : (typedef|execution)+;

name:NAME|TYPEDEF|MAP|ONEOF|FOR|BREAK|CONTINUE|IF|ELSE|SWITCH|CASE|RETURN|VAR|LOCAL|CONST|MESSAGE|ENUM;
fieldname:NAME|TYPEDEF|MAP|ONEOF|UINT32|UINT64|INT32|INT64|FLOAT32|FLOAT64|STRING|BYTES|BOOL|UINT8|CHAN|ANY|FOR|BREAK|CONTINUE|IF|ELSE|SWITCH|CASE|RETURN|VAR|LOCAL|CONST|MESSAGE|ENUM;


functiondef
    : FUNCTION name '(' (inparam (',' inparam)* (TAILARRAY)?)?  ')' returntypename?  block # FunctionDef
    | FUNCTION name '(' (inparam (',' inparam)* (TAILARRAY)?)?  ')' '('returntypename (',' returntypename) *')' block # FunctionDef
    | FUNCTION name '(' (inparam (',' inparam)* (TAILARRAY)?)?  ')' outparam?  block # FunctionDef
    | FUNCTION name '(' (inparam (',' inparam)* (TAILARRAY)?)?  ')' '('outparam (',' outparam) *')' block # FunctionDef
    ;

lambda
    : FUNCTION  '(' (inparam (',' inparam)* (TAILARRAY)?)?  ')' returntypename?  block # LambdaDef
    | FUNCTION  '(' (inparam (',' inparam)* (TAILARRAY)?)?  ')' '('returntypename (',' returntypename) *')' block # LambdaDef
    | FUNCTION  '(' (inparam (',' inparam)* (TAILARRAY)?)?  ')' outparam?  block # LambdaDef
    | FUNCTION  '(' (inparam (',' inparam)* (TAILARRAY)?)?  ')' '('outparam (',' outparam) *')' block # LambdaDef
    ;

inparam
    : param;

outparam
    : param;

intypename
    : typename;

returntypename
    : typename;

param
    : name typename;

typename
    : (name|basicTypeName) # SimpleTypeName
    | functionTypeName # FunctionType
    | MAP '[' basicTypeName ']' typename # MapTypeName
    | '[]' typename # SliceTypeName
    | CHAN '[' typename ']' # ChanTypeName
    ;

functionTypeName
    : FUNCTION  '(' (intypename (',' intypename)* (TAILARRAY)?)?  ')' returntypename?
    | FUNCTION  '(' (intypename (',' intypename)* (TAILARRAY)?)?  ')' '('returntypename (',' returntypename) *')'
    ;

typedef
    : TYPEDEF name typenameindef # TypeDefAlias
    | TYPEDEF name MESSAGE '{' (messagefield (messagefield)*)? '}' # TypeDefMessage
    | TYPEDEF name ENUM '{' (name ':' INT)* '}' # TypeDefEnum
    ;

messagefield
    : fieldname typenameindef # FieldDef
    | ONEOF fieldname '{' oneoffield (oneoffield)* '}' # OneofDef
    ;

oneoffield
    : fieldname typenameindef;

typenameindef
    : simpleTypeNameindef
    | chanTypeNameindef
    | sliceTypeNameindef
    | mapTypeNameindef
    | functionTypeNameindef
    ;

simpleTypeNameindef
    :(name|basicTypeName);
chanTypeNameindef
    : CHAN '[' typenameindef ']';
sliceTypeNameindef
    : '[]' typenameindef;
mapTypeNameindef
    : MAP '[' basicTypeName ']' typenameindef;
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
    | functiondef # FunctionDefine
    ;

control
    : IF '(' expr ')' block (ELSE (block|control))? # If
    | SWITCH '(' expr ')' '{' (CASE constant ':' block)+ '}' # Switch
    | FOR '(' name 'in' expr ')' block # ForInSlice
    | FOR '(' name ',' name 'in' expr ')' block  # ForInMap
    | FOR '(' line ';' expr ';' restoreStack ')' block  # For
    | BREAK ';' # Break
    | CONTINUE ';' # Continue
    | RETURN ';'# ReturnVoid
    | RETURN expr (',' expr)*';' # ReturnVal
    ;

block
    : '{' (execution)* '}';

line
    : restoreStack # RestoreStackSp
    | symbolDef # SymbolDefine
    ;

restoreStack
    : keepStack;

keepStack
    : expr # ExprEntry
    | expr (',' expr)+ op=ASSIGN expr # FunctionAssign
    | expr op=(ASSIGN|ADDEQUAL|SUBEQUAL|MULEQUAL|DIVEQUAL) expr # Assign
    ;

symbol
    : name| '@';

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
    | symbol # Pass
    | lambda # Pass
    | constructor # Pass
    | (typename)? '{' (expr (',' expr)*)? '}' # InitSlice
    | (typename)? '{'  (expr ':' expr (',' expr ':' expr )*)? '}' # InitKV
    | expr DOT fieldname # Select
    | expr DOT '(' typename ')' # TypeAssert
    | expr '[?(' expr ')]' # SliceFilter
    | expr '[' expr ']' # Index
    | expr '[' indexs (',' indexs)* ']' # SliceMultiIndex
    | expr '[' '[' (expr (',' expr)*)? ']' ']' # MapMultiIndex
    | expr '(' (expr (',' expr)*)? ')' # Call
    | typename '(' expr ')' # TypeConvert
    | op=(UNARYADD|UNARYSUB|NOT|SUB) expr # LeftUnary
    | expr op=(UNARYADD|UNARYSUB) # RightUnary
    | <assoc=right> expr op=POW  expr # Binary
    | expr op=(MUL | DIV | MOD) expr # Binary
    | expr op=(ADD | SUB) expr # Binary
    | expr op=(EQ | INEQ | GT | GE | LT | LE | REGEX) expr # Binary
    | expr op=AND expr # Binary
    | expr op=OR expr # Binary
    | expr (CHANOP|CHANOPNONBLOCK) expr # Send
    | (CHANOP|CHANOPNONBLOCK) expr # Recv
    ;

basicTypeName
    : (UINT32|UINT64|INT32|INT64|FLOAT32|FLOAT64|STRING|BYTES|BOOL|UINT8|ANY);

constant
    : INT # ConstantInt
    | FLOAT # ConstantFloat
    | BOOLLITERAL # ConstantBool
    | NULL # ConstantNil
    | STRINGLITERAL # ConstantString
    ;

constructor
    : NEW typename '(' (expr (',' expr)*)? ')'
    ;

symbolDef
    : VAR name typename
    | VAR name typename '=' expr
    | VAR name '=' expr
    | LOCAL name typename
    | LOCAL name typename '=' expr
    | LOCAL name '=' expr
    | name LOCALASSIGN expr
    | CONST name typename '=' expr
    | CONST name '=' expr
    ;