grammar goscript;

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

BOOL
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

STRING
   : '\'' ( ~('\''|'\\') | ('\\' .) )* '\'';

NAME: [a-zA-Z_]+[a-zA-Z0-9_]*;

DOT: '.';
WHITESPACE: [ \r\n\t]+ -> skip;

program
    : statement+;

statement
    : functiondef # FuncDef
    | typedef ';' # TypeDef
    | vardef # GlobalDef
    | enumdef # EnumDef
    | execution # Run
    ;

functiondef
    : FUNCTION NAME '(' param (',' param)*  ')' NAME?  block
    | FUNCTION NAME '(' ')' NAME?  block;

param
    : NAME NAME;

typedef
    : TYPEDEF NAME typename # TypeDefAlias
    | TYPEDEF NAME '{' (NAME typename ',')* '}' # TypeDefComplex
    ;

typename
    : NAME # SimpleTypeName
    | NAME '<'typename (',' typename)*'>' # GenericTypeName;

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
    | FOR '(' expr ';' expr ';' expr ')' block  # For
    | BREAK ';' # Break
    | CONTINUE ';' # Continue
    | RETURN ';'# ReturnVoid
    | RETURN expr ';' # ReturnVal
    ;

collection
    : expr;

block
    : '{' (execution|vardef)* '}';

line
    : expr # ExprEntry
    | lhs (',' lhs)* op=ASSIGN functions # FunctionAssign
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
    | functions # Pass
    | constructor # Construct
    | localdef # LocalDef
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
    | BOOL # ConstantBool
    | NULL # ConstantNil
    | STRING # ConstantString
    ;

functions
    : NAME '(' expr (',' expr)* ')'
    | variable DOT NAME '(' ')'
    | variable DOT NAME '(' expr (',' expr)* ')'
    | NAME '(' ')'
    ;

constructor
    : NEW NAME '(' ')'
    | NEW NAME '('expr (',' expr)*')';

vardef
    : VAR NAME NAME ';';

localdef
    : LOCAL NAME NAME # Local
    | LOCAL NAME NAME '=' expr # LocalAssign
    ;