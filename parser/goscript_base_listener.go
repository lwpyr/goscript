// Code generated from goscript.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // goscript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasegoscriptListener is a complete listener for a parse tree produced by goscriptParser.
type BasegoscriptListener struct{}

var _ goscriptListener = &BasegoscriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasegoscriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasegoscriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasegoscriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasegoscriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasegoscriptListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasegoscriptListener) ExitProgram(ctx *ProgramContext) {}

// EnterName is called when production name is entered.
func (s *BasegoscriptListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BasegoscriptListener) ExitName(ctx *NameContext) {}

// EnterFieldname is called when production fieldname is entered.
func (s *BasegoscriptListener) EnterFieldname(ctx *FieldnameContext) {}

// ExitFieldname is called when production fieldname is exited.
func (s *BasegoscriptListener) ExitFieldname(ctx *FieldnameContext) {}

// EnterFunctionDef is called when production FunctionDef is entered.
func (s *BasegoscriptListener) EnterFunctionDef(ctx *FunctionDefContext) {}

// ExitFunctionDef is called when production FunctionDef is exited.
func (s *BasegoscriptListener) ExitFunctionDef(ctx *FunctionDefContext) {}

// EnterLambdaDef is called when production LambdaDef is entered.
func (s *BasegoscriptListener) EnterLambdaDef(ctx *LambdaDefContext) {}

// ExitLambdaDef is called when production LambdaDef is exited.
func (s *BasegoscriptListener) ExitLambdaDef(ctx *LambdaDefContext) {}

// EnterClosure is called when production closure is entered.
func (s *BasegoscriptListener) EnterClosure(ctx *ClosureContext) {}

// ExitClosure is called when production closure is exited.
func (s *BasegoscriptListener) ExitClosure(ctx *ClosureContext) {}

// EnterInparam is called when production inparam is entered.
func (s *BasegoscriptListener) EnterInparam(ctx *InparamContext) {}

// ExitInparam is called when production inparam is exited.
func (s *BasegoscriptListener) ExitInparam(ctx *InparamContext) {}

// EnterOutparam is called when production outparam is entered.
func (s *BasegoscriptListener) EnterOutparam(ctx *OutparamContext) {}

// ExitOutparam is called when production outparam is exited.
func (s *BasegoscriptListener) ExitOutparam(ctx *OutparamContext) {}

// EnterIntypename is called when production intypename is entered.
func (s *BasegoscriptListener) EnterIntypename(ctx *IntypenameContext) {}

// ExitIntypename is called when production intypename is exited.
func (s *BasegoscriptListener) ExitIntypename(ctx *IntypenameContext) {}

// EnterReturntypename is called when production returntypename is entered.
func (s *BasegoscriptListener) EnterReturntypename(ctx *ReturntypenameContext) {}

// ExitReturntypename is called when production returntypename is exited.
func (s *BasegoscriptListener) ExitReturntypename(ctx *ReturntypenameContext) {}

// EnterParam is called when production param is entered.
func (s *BasegoscriptListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BasegoscriptListener) ExitParam(ctx *ParamContext) {}

// EnterSimpleTypeName is called when production SimpleTypeName is entered.
func (s *BasegoscriptListener) EnterSimpleTypeName(ctx *SimpleTypeNameContext) {}

// ExitSimpleTypeName is called when production SimpleTypeName is exited.
func (s *BasegoscriptListener) ExitSimpleTypeName(ctx *SimpleTypeNameContext) {}

// EnterFunctionType is called when production FunctionType is entered.
func (s *BasegoscriptListener) EnterFunctionType(ctx *FunctionTypeContext) {}

// ExitFunctionType is called when production FunctionType is exited.
func (s *BasegoscriptListener) ExitFunctionType(ctx *FunctionTypeContext) {}

// EnterMapTypeName is called when production MapTypeName is entered.
func (s *BasegoscriptListener) EnterMapTypeName(ctx *MapTypeNameContext) {}

// ExitMapTypeName is called when production MapTypeName is exited.
func (s *BasegoscriptListener) ExitMapTypeName(ctx *MapTypeNameContext) {}

// EnterSliceTypeName is called when production SliceTypeName is entered.
func (s *BasegoscriptListener) EnterSliceTypeName(ctx *SliceTypeNameContext) {}

// ExitSliceTypeName is called when production SliceTypeName is exited.
func (s *BasegoscriptListener) ExitSliceTypeName(ctx *SliceTypeNameContext) {}

// EnterChanTypeName is called when production ChanTypeName is entered.
func (s *BasegoscriptListener) EnterChanTypeName(ctx *ChanTypeNameContext) {}

// ExitChanTypeName is called when production ChanTypeName is exited.
func (s *BasegoscriptListener) ExitChanTypeName(ctx *ChanTypeNameContext) {}

// EnterFunctionTypeName is called when production functionTypeName is entered.
func (s *BasegoscriptListener) EnterFunctionTypeName(ctx *FunctionTypeNameContext) {}

// ExitFunctionTypeName is called when production functionTypeName is exited.
func (s *BasegoscriptListener) ExitFunctionTypeName(ctx *FunctionTypeNameContext) {}

// EnterTypeDefAlias is called when production TypeDefAlias is entered.
func (s *BasegoscriptListener) EnterTypeDefAlias(ctx *TypeDefAliasContext) {}

// ExitTypeDefAlias is called when production TypeDefAlias is exited.
func (s *BasegoscriptListener) ExitTypeDefAlias(ctx *TypeDefAliasContext) {}

// EnterTypeDefMessage is called when production TypeDefMessage is entered.
func (s *BasegoscriptListener) EnterTypeDefMessage(ctx *TypeDefMessageContext) {}

// ExitTypeDefMessage is called when production TypeDefMessage is exited.
func (s *BasegoscriptListener) ExitTypeDefMessage(ctx *TypeDefMessageContext) {}

// EnterTypeDefEnum is called when production TypeDefEnum is entered.
func (s *BasegoscriptListener) EnterTypeDefEnum(ctx *TypeDefEnumContext) {}

// ExitTypeDefEnum is called when production TypeDefEnum is exited.
func (s *BasegoscriptListener) ExitTypeDefEnum(ctx *TypeDefEnumContext) {}

// EnterFieldDef is called when production FieldDef is entered.
func (s *BasegoscriptListener) EnterFieldDef(ctx *FieldDefContext) {}

// ExitFieldDef is called when production FieldDef is exited.
func (s *BasegoscriptListener) ExitFieldDef(ctx *FieldDefContext) {}

// EnterOneofDef is called when production OneofDef is entered.
func (s *BasegoscriptListener) EnterOneofDef(ctx *OneofDefContext) {}

// ExitOneofDef is called when production OneofDef is exited.
func (s *BasegoscriptListener) ExitOneofDef(ctx *OneofDefContext) {}

// EnterOneoffield is called when production oneoffield is entered.
func (s *BasegoscriptListener) EnterOneoffield(ctx *OneoffieldContext) {}

// ExitOneoffield is called when production oneoffield is exited.
func (s *BasegoscriptListener) ExitOneoffield(ctx *OneoffieldContext) {}

// EnterTypenameindef is called when production typenameindef is entered.
func (s *BasegoscriptListener) EnterTypenameindef(ctx *TypenameindefContext) {}

// ExitTypenameindef is called when production typenameindef is exited.
func (s *BasegoscriptListener) ExitTypenameindef(ctx *TypenameindefContext) {}

// EnterSimpleTypeNameindef is called when production simpleTypeNameindef is entered.
func (s *BasegoscriptListener) EnterSimpleTypeNameindef(ctx *SimpleTypeNameindefContext) {}

// ExitSimpleTypeNameindef is called when production simpleTypeNameindef is exited.
func (s *BasegoscriptListener) ExitSimpleTypeNameindef(ctx *SimpleTypeNameindefContext) {}

// EnterChanTypeNameindef is called when production chanTypeNameindef is entered.
func (s *BasegoscriptListener) EnterChanTypeNameindef(ctx *ChanTypeNameindefContext) {}

// ExitChanTypeNameindef is called when production chanTypeNameindef is exited.
func (s *BasegoscriptListener) ExitChanTypeNameindef(ctx *ChanTypeNameindefContext) {}

// EnterSliceTypeNameindef is called when production sliceTypeNameindef is entered.
func (s *BasegoscriptListener) EnterSliceTypeNameindef(ctx *SliceTypeNameindefContext) {}

// ExitSliceTypeNameindef is called when production sliceTypeNameindef is exited.
func (s *BasegoscriptListener) ExitSliceTypeNameindef(ctx *SliceTypeNameindefContext) {}

// EnterMapTypeNameindef is called when production mapTypeNameindef is entered.
func (s *BasegoscriptListener) EnterMapTypeNameindef(ctx *MapTypeNameindefContext) {}

// ExitMapTypeNameindef is called when production mapTypeNameindef is exited.
func (s *BasegoscriptListener) ExitMapTypeNameindef(ctx *MapTypeNameindefContext) {}

// EnterFunctionTypeNameindef is called when production functionTypeNameindef is entered.
func (s *BasegoscriptListener) EnterFunctionTypeNameindef(ctx *FunctionTypeNameindefContext) {}

// ExitFunctionTypeNameindef is called when production functionTypeNameindef is exited.
func (s *BasegoscriptListener) ExitFunctionTypeNameindef(ctx *FunctionTypeNameindefContext) {}

// EnterIntypenameindef is called when production intypenameindef is entered.
func (s *BasegoscriptListener) EnterIntypenameindef(ctx *IntypenameindefContext) {}

// ExitIntypenameindef is called when production intypenameindef is exited.
func (s *BasegoscriptListener) ExitIntypenameindef(ctx *IntypenameindefContext) {}

// EnterReturntypenameindef is called when production returntypenameindef is entered.
func (s *BasegoscriptListener) EnterReturntypenameindef(ctx *ReturntypenameindefContext) {}

// ExitReturntypenameindef is called when production returntypenameindef is exited.
func (s *BasegoscriptListener) ExitReturntypenameindef(ctx *ReturntypenameindefContext) {}

// EnterCtrl is called when production Ctrl is entered.
func (s *BasegoscriptListener) EnterCtrl(ctx *CtrlContext) {}

// ExitCtrl is called when production Ctrl is exited.
func (s *BasegoscriptListener) ExitCtrl(ctx *CtrlContext) {}

// EnterLineProgram is called when production LineProgram is entered.
func (s *BasegoscriptListener) EnterLineProgram(ctx *LineProgramContext) {}

// ExitLineProgram is called when production LineProgram is exited.
func (s *BasegoscriptListener) ExitLineProgram(ctx *LineProgramContext) {}

// EnterIf is called when production If is entered.
func (s *BasegoscriptListener) EnterIf(ctx *IfContext) {}

// ExitIf is called when production If is exited.
func (s *BasegoscriptListener) ExitIf(ctx *IfContext) {}

// EnterSwitch is called when production Switch is entered.
func (s *BasegoscriptListener) EnterSwitch(ctx *SwitchContext) {}

// ExitSwitch is called when production Switch is exited.
func (s *BasegoscriptListener) ExitSwitch(ctx *SwitchContext) {}

// EnterForInSlice is called when production ForInSlice is entered.
func (s *BasegoscriptListener) EnterForInSlice(ctx *ForInSliceContext) {}

// ExitForInSlice is called when production ForInSlice is exited.
func (s *BasegoscriptListener) ExitForInSlice(ctx *ForInSliceContext) {}

// EnterForInMap is called when production ForInMap is entered.
func (s *BasegoscriptListener) EnterForInMap(ctx *ForInMapContext) {}

// ExitForInMap is called when production ForInMap is exited.
func (s *BasegoscriptListener) ExitForInMap(ctx *ForInMapContext) {}

// EnterFor is called when production For is entered.
func (s *BasegoscriptListener) EnterFor(ctx *ForContext) {}

// ExitFor is called when production For is exited.
func (s *BasegoscriptListener) ExitFor(ctx *ForContext) {}

// EnterBreak is called when production Break is entered.
func (s *BasegoscriptListener) EnterBreak(ctx *BreakContext) {}

// ExitBreak is called when production Break is exited.
func (s *BasegoscriptListener) ExitBreak(ctx *BreakContext) {}

// EnterContinue is called when production Continue is entered.
func (s *BasegoscriptListener) EnterContinue(ctx *ContinueContext) {}

// ExitContinue is called when production Continue is exited.
func (s *BasegoscriptListener) ExitContinue(ctx *ContinueContext) {}

// EnterReturnVoid is called when production ReturnVoid is entered.
func (s *BasegoscriptListener) EnterReturnVoid(ctx *ReturnVoidContext) {}

// ExitReturnVoid is called when production ReturnVoid is exited.
func (s *BasegoscriptListener) ExitReturnVoid(ctx *ReturnVoidContext) {}

// EnterReturnVal is called when production ReturnVal is entered.
func (s *BasegoscriptListener) EnterReturnVal(ctx *ReturnValContext) {}

// ExitReturnVal is called when production ReturnVal is exited.
func (s *BasegoscriptListener) ExitReturnVal(ctx *ReturnValContext) {}

// EnterCollection is called when production collection is entered.
func (s *BasegoscriptListener) EnterCollection(ctx *CollectionContext) {}

// ExitCollection is called when production collection is exited.
func (s *BasegoscriptListener) ExitCollection(ctx *CollectionContext) {}

// EnterBlock is called when production block is entered.
func (s *BasegoscriptListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasegoscriptListener) ExitBlock(ctx *BlockContext) {}

// EnterRestoreStackSp is called when production RestoreStackSp is entered.
func (s *BasegoscriptListener) EnterRestoreStackSp(ctx *RestoreStackSpContext) {}

// ExitRestoreStackSp is called when production RestoreStackSp is exited.
func (s *BasegoscriptListener) ExitRestoreStackSp(ctx *RestoreStackSpContext) {}

// EnterVarDef is called when production VarDef is entered.
func (s *BasegoscriptListener) EnterVarDef(ctx *VarDefContext) {}

// ExitVarDef is called when production VarDef is exited.
func (s *BasegoscriptListener) ExitVarDef(ctx *VarDefContext) {}

// EnterConstDef is called when production ConstDef is entered.
func (s *BasegoscriptListener) EnterConstDef(ctx *ConstDefContext) {}

// ExitConstDef is called when production ConstDef is exited.
func (s *BasegoscriptListener) ExitConstDef(ctx *ConstDefContext) {}

// EnterRestoreStack is called when production restoreStack is entered.
func (s *BasegoscriptListener) EnterRestoreStack(ctx *RestoreStackContext) {}

// ExitRestoreStack is called when production restoreStack is exited.
func (s *BasegoscriptListener) ExitRestoreStack(ctx *RestoreStackContext) {}

// EnterExprEntry is called when production ExprEntry is entered.
func (s *BasegoscriptListener) EnterExprEntry(ctx *ExprEntryContext) {}

// ExitExprEntry is called when production ExprEntry is exited.
func (s *BasegoscriptListener) ExitExprEntry(ctx *ExprEntryContext) {}

// EnterFunctionAssign is called when production FunctionAssign is entered.
func (s *BasegoscriptListener) EnterFunctionAssign(ctx *FunctionAssignContext) {}

// ExitFunctionAssign is called when production FunctionAssign is exited.
func (s *BasegoscriptListener) ExitFunctionAssign(ctx *FunctionAssignContext) {}

// EnterVariableName is called when production VariableName is entered.
func (s *BasegoscriptListener) EnterVariableName(ctx *VariableNameContext) {}

// ExitVariableName is called when production VariableName is exited.
func (s *BasegoscriptListener) ExitVariableName(ctx *VariableNameContext) {}

// EnterAsserted is called when production asserted is entered.
func (s *BasegoscriptListener) EnterAsserted(ctx *AssertedContext) {}

// ExitAsserted is called when production asserted is exited.
func (s *BasegoscriptListener) ExitAsserted(ctx *AssertedContext) {}

// EnterFilter is called when production filter is entered.
func (s *BasegoscriptListener) EnterFilter(ctx *FilterContext) {}

// ExitFilter is called when production filter is exited.
func (s *BasegoscriptListener) ExitFilter(ctx *FilterContext) {}

// EnterIndexType1 is called when production IndexType1 is entered.
func (s *BasegoscriptListener) EnterIndexType1(ctx *IndexType1Context) {}

// ExitIndexType1 is called when production IndexType1 is exited.
func (s *BasegoscriptListener) ExitIndexType1(ctx *IndexType1Context) {}

// EnterIndexType2 is called when production IndexType2 is entered.
func (s *BasegoscriptListener) EnterIndexType2(ctx *IndexType2Context) {}

// ExitIndexType2 is called when production IndexType2 is exited.
func (s *BasegoscriptListener) ExitIndexType2(ctx *IndexType2Context) {}

// EnterIndexType3 is called when production IndexType3 is entered.
func (s *BasegoscriptListener) EnterIndexType3(ctx *IndexType3Context) {}

// ExitIndexType3 is called when production IndexType3 is exited.
func (s *BasegoscriptListener) ExitIndexType3(ctx *IndexType3Context) {}

// EnterIndexType4 is called when production IndexType4 is entered.
func (s *BasegoscriptListener) EnterIndexType4(ctx *IndexType4Context) {}

// ExitIndexType4 is called when production IndexType4 is exited.
func (s *BasegoscriptListener) ExitIndexType4(ctx *IndexType4Context) {}

// EnterIndexType5 is called when production IndexType5 is entered.
func (s *BasegoscriptListener) EnterIndexType5(ctx *IndexType5Context) {}

// ExitIndexType5 is called when production IndexType5 is exited.
func (s *BasegoscriptListener) ExitIndexType5(ctx *IndexType5Context) {}

// EnterRecv is called when production Recv is entered.
func (s *BasegoscriptListener) EnterRecv(ctx *RecvContext) {}

// ExitRecv is called when production Recv is exited.
func (s *BasegoscriptListener) ExitRecv(ctx *RecvContext) {}

// EnterSliceFilter is called when production SliceFilter is entered.
func (s *BasegoscriptListener) EnterSliceFilter(ctx *SliceFilterContext) {}

// ExitSliceFilter is called when production SliceFilter is exited.
func (s *BasegoscriptListener) ExitSliceFilter(ctx *SliceFilterContext) {}

// EnterDirectCall is called when production DirectCall is entered.
func (s *BasegoscriptListener) EnterDirectCall(ctx *DirectCallContext) {}

// ExitDirectCall is called when production DirectCall is exited.
func (s *BasegoscriptListener) ExitDirectCall(ctx *DirectCallContext) {}

// EnterIndex is called when production Index is entered.
func (s *BasegoscriptListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production Index is exited.
func (s *BasegoscriptListener) ExitIndex(ctx *IndexContext) {}

// EnterMapMultiIndex is called when production MapMultiIndex is entered.
func (s *BasegoscriptListener) EnterMapMultiIndex(ctx *MapMultiIndexContext) {}

// ExitMapMultiIndex is called when production MapMultiIndex is exited.
func (s *BasegoscriptListener) ExitMapMultiIndex(ctx *MapMultiIndexContext) {}

// EnterSend is called when production Send is entered.
func (s *BasegoscriptListener) EnterSend(ctx *SendContext) {}

// ExitSend is called when production Send is exited.
func (s *BasegoscriptListener) ExitSend(ctx *SendContext) {}

// EnterPass is called when production Pass is entered.
func (s *BasegoscriptListener) EnterPass(ctx *PassContext) {}

// ExitPass is called when production Pass is exited.
func (s *BasegoscriptListener) ExitPass(ctx *PassContext) {}

// EnterSelect is called when production Select is entered.
func (s *BasegoscriptListener) EnterSelect(ctx *SelectContext) {}

// ExitSelect is called when production Select is exited.
func (s *BasegoscriptListener) ExitSelect(ctx *SelectContext) {}

// EnterConstruct is called when production Construct is entered.
func (s *BasegoscriptListener) EnterConstruct(ctx *ConstructContext) {}

// ExitConstruct is called when production Construct is exited.
func (s *BasegoscriptListener) ExitConstruct(ctx *ConstructContext) {}

// EnterBinary is called when production Binary is entered.
func (s *BasegoscriptListener) EnterBinary(ctx *BinaryContext) {}

// ExitBinary is called when production Binary is exited.
func (s *BasegoscriptListener) ExitBinary(ctx *BinaryContext) {}

// EnterLeftUnary is called when production LeftUnary is entered.
func (s *BasegoscriptListener) EnterLeftUnary(ctx *LeftUnaryContext) {}

// ExitLeftUnary is called when production LeftUnary is exited.
func (s *BasegoscriptListener) ExitLeftUnary(ctx *LeftUnaryContext) {}

// EnterTypeAssert is called when production TypeAssert is entered.
func (s *BasegoscriptListener) EnterTypeAssert(ctx *TypeAssertContext) {}

// ExitTypeAssert is called when production TypeAssert is exited.
func (s *BasegoscriptListener) ExitTypeAssert(ctx *TypeAssertContext) {}

// EnterSliceMultiIndex is called when production SliceMultiIndex is entered.
func (s *BasegoscriptListener) EnterSliceMultiIndex(ctx *SliceMultiIndexContext) {}

// ExitSliceMultiIndex is called when production SliceMultiIndex is exited.
func (s *BasegoscriptListener) ExitSliceMultiIndex(ctx *SliceMultiIndexContext) {}

// EnterRightUnary is called when production RightUnary is entered.
func (s *BasegoscriptListener) EnterRightUnary(ctx *RightUnaryContext) {}

// ExitRightUnary is called when production RightUnary is exited.
func (s *BasegoscriptListener) ExitRightUnary(ctx *RightUnaryContext) {}

// EnterAssignInitializationlist is called when production AssignInitializationlist is entered.
func (s *BasegoscriptListener) EnterAssignInitializationlist(ctx *AssignInitializationlistContext) {}

// ExitAssignInitializationlist is called when production AssignInitializationlist is exited.
func (s *BasegoscriptListener) ExitAssignInitializationlist(ctx *AssignInitializationlistContext) {}

// EnterBasicTypeName is called when production basicTypeName is entered.
func (s *BasegoscriptListener) EnterBasicTypeName(ctx *BasicTypeNameContext) {}

// ExitBasicTypeName is called when production basicTypeName is exited.
func (s *BasegoscriptListener) ExitBasicTypeName(ctx *BasicTypeNameContext) {}

// EnterBuiltin is called when production builtin is entered.
func (s *BasegoscriptListener) EnterBuiltin(ctx *BuiltinContext) {}

// ExitBuiltin is called when production builtin is exited.
func (s *BasegoscriptListener) ExitBuiltin(ctx *BuiltinContext) {}

// EnterInitializationListBegin is called when production initializationListBegin is entered.
func (s *BasegoscriptListener) EnterInitializationListBegin(ctx *InitializationListBeginContext) {}

// ExitInitializationListBegin is called when production initializationListBegin is exited.
func (s *BasegoscriptListener) ExitInitializationListBegin(ctx *InitializationListBeginContext) {}

// EnterInitSlice is called when production InitSlice is entered.
func (s *BasegoscriptListener) EnterInitSlice(ctx *InitSliceContext) {}

// ExitInitSlice is called when production InitSlice is exited.
func (s *BasegoscriptListener) ExitInitSlice(ctx *InitSliceContext) {}

// EnterInitMessage is called when production InitMessage is entered.
func (s *BasegoscriptListener) EnterInitMessage(ctx *InitMessageContext) {}

// ExitInitMessage is called when production InitMessage is exited.
func (s *BasegoscriptListener) ExitInitMessage(ctx *InitMessageContext) {}

// EnterInitMap is called when production InitMap is entered.
func (s *BasegoscriptListener) EnterInitMap(ctx *InitMapContext) {}

// ExitInitMap is called when production InitMap is exited.
func (s *BasegoscriptListener) ExitInitMap(ctx *InitMapContext) {}

// EnterInitConstant is called when production InitConstant is entered.
func (s *BasegoscriptListener) EnterInitConstant(ctx *InitConstantContext) {}

// ExitInitConstant is called when production InitConstant is exited.
func (s *BasegoscriptListener) ExitInitConstant(ctx *InitConstantContext) {}

// EnterConstantInt is called when production ConstantInt is entered.
func (s *BasegoscriptListener) EnterConstantInt(ctx *ConstantIntContext) {}

// ExitConstantInt is called when production ConstantInt is exited.
func (s *BasegoscriptListener) ExitConstantInt(ctx *ConstantIntContext) {}

// EnterConstantFloat is called when production ConstantFloat is entered.
func (s *BasegoscriptListener) EnterConstantFloat(ctx *ConstantFloatContext) {}

// ExitConstantFloat is called when production ConstantFloat is exited.
func (s *BasegoscriptListener) ExitConstantFloat(ctx *ConstantFloatContext) {}

// EnterConstantBool is called when production ConstantBool is entered.
func (s *BasegoscriptListener) EnterConstantBool(ctx *ConstantBoolContext) {}

// ExitConstantBool is called when production ConstantBool is exited.
func (s *BasegoscriptListener) ExitConstantBool(ctx *ConstantBoolContext) {}

// EnterConstantNil is called when production ConstantNil is entered.
func (s *BasegoscriptListener) EnterConstantNil(ctx *ConstantNilContext) {}

// ExitConstantNil is called when production ConstantNil is exited.
func (s *BasegoscriptListener) ExitConstantNil(ctx *ConstantNilContext) {}

// EnterConstantString is called when production ConstantString is entered.
func (s *BasegoscriptListener) EnterConstantString(ctx *ConstantStringContext) {}

// ExitConstantString is called when production ConstantString is exited.
func (s *BasegoscriptListener) ExitConstantString(ctx *ConstantStringContext) {}

// EnterConstructor is called when production constructor is entered.
func (s *BasegoscriptListener) EnterConstructor(ctx *ConstructorContext) {}

// ExitConstructor is called when production constructor is exited.
func (s *BasegoscriptListener) ExitConstructor(ctx *ConstructorContext) {}

// EnterVardef is called when production vardef is entered.
func (s *BasegoscriptListener) EnterVardef(ctx *VardefContext) {}

// ExitVardef is called when production vardef is exited.
func (s *BasegoscriptListener) ExitVardef(ctx *VardefContext) {}

// EnterConstdef is called when production constdef is entered.
func (s *BasegoscriptListener) EnterConstdef(ctx *ConstdefContext) {}

// ExitConstdef is called when production constdef is exited.
func (s *BasegoscriptListener) ExitConstdef(ctx *ConstdefContext) {}
