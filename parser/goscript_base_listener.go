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

// EnterFuncDef is called when production FuncDef is entered.
func (s *BasegoscriptListener) EnterFuncDef(ctx *FuncDefContext) {}

// ExitFuncDef is called when production FuncDef is exited.
func (s *BasegoscriptListener) ExitFuncDef(ctx *FuncDefContext) {}

// EnterTypeDef is called when production TypeDef is entered.
func (s *BasegoscriptListener) EnterTypeDef(ctx *TypeDefContext) {}

// ExitTypeDef is called when production TypeDef is exited.
func (s *BasegoscriptListener) ExitTypeDef(ctx *TypeDefContext) {}

// EnterGlobalDef is called when production GlobalDef is entered.
func (s *BasegoscriptListener) EnterGlobalDef(ctx *GlobalDefContext) {}

// ExitGlobalDef is called when production GlobalDef is exited.
func (s *BasegoscriptListener) ExitGlobalDef(ctx *GlobalDefContext) {}

// EnterEnumDef is called when production EnumDef is entered.
func (s *BasegoscriptListener) EnterEnumDef(ctx *EnumDefContext) {}

// ExitEnumDef is called when production EnumDef is exited.
func (s *BasegoscriptListener) ExitEnumDef(ctx *EnumDefContext) {}

// EnterRun is called when production Run is entered.
func (s *BasegoscriptListener) EnterRun(ctx *RunContext) {}

// ExitRun is called when production Run is exited.
func (s *BasegoscriptListener) ExitRun(ctx *RunContext) {}

// EnterFunctiondef is called when production functiondef is entered.
func (s *BasegoscriptListener) EnterFunctiondef(ctx *FunctiondefContext) {}

// ExitFunctiondef is called when production functiondef is exited.
func (s *BasegoscriptListener) ExitFunctiondef(ctx *FunctiondefContext) {}

// EnterParam is called when production param is entered.
func (s *BasegoscriptListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BasegoscriptListener) ExitParam(ctx *ParamContext) {}

// EnterTypeDefAlias is called when production TypeDefAlias is entered.
func (s *BasegoscriptListener) EnterTypeDefAlias(ctx *TypeDefAliasContext) {}

// ExitTypeDefAlias is called when production TypeDefAlias is exited.
func (s *BasegoscriptListener) ExitTypeDefAlias(ctx *TypeDefAliasContext) {}

// EnterTypeDefComplex is called when production TypeDefComplex is entered.
func (s *BasegoscriptListener) EnterTypeDefComplex(ctx *TypeDefComplexContext) {}

// ExitTypeDefComplex is called when production TypeDefComplex is exited.
func (s *BasegoscriptListener) ExitTypeDefComplex(ctx *TypeDefComplexContext) {}

// EnterSimpleTypeName is called when production SimpleTypeName is entered.
func (s *BasegoscriptListener) EnterSimpleTypeName(ctx *SimpleTypeNameContext) {}

// ExitSimpleTypeName is called when production SimpleTypeName is exited.
func (s *BasegoscriptListener) ExitSimpleTypeName(ctx *SimpleTypeNameContext) {}

// EnterGenericTypeName is called when production GenericTypeName is entered.
func (s *BasegoscriptListener) EnterGenericTypeName(ctx *GenericTypeNameContext) {}

// ExitGenericTypeName is called when production GenericTypeName is exited.
func (s *BasegoscriptListener) ExitGenericTypeName(ctx *GenericTypeNameContext) {}

// EnterEnumdef is called when production enumdef is entered.
func (s *BasegoscriptListener) EnterEnumdef(ctx *EnumdefContext) {}

// ExitEnumdef is called when production enumdef is exited.
func (s *BasegoscriptListener) ExitEnumdef(ctx *EnumdefContext) {}

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

// EnterExprEntry is called when production ExprEntry is entered.
func (s *BasegoscriptListener) EnterExprEntry(ctx *ExprEntryContext) {}

// ExitExprEntry is called when production ExprEntry is exited.
func (s *BasegoscriptListener) ExitExprEntry(ctx *ExprEntryContext) {}

// EnterFunctionAssign is called when production FunctionAssign is entered.
func (s *BasegoscriptListener) EnterFunctionAssign(ctx *FunctionAssignContext) {}

// ExitFunctionAssign is called when production FunctionAssign is exited.
func (s *BasegoscriptListener) ExitFunctionAssign(ctx *FunctionAssignContext) {}

// EnterMultiAssign is called when production MultiAssign is entered.
func (s *BasegoscriptListener) EnterMultiAssign(ctx *MultiAssignContext) {}

// ExitMultiAssign is called when production MultiAssign is exited.
func (s *BasegoscriptListener) ExitMultiAssign(ctx *MultiAssignContext) {}

// EnterLhs is called when production lhs is entered.
func (s *BasegoscriptListener) EnterLhs(ctx *LhsContext) {}

// ExitLhs is called when production lhs is exited.
func (s *BasegoscriptListener) ExitLhs(ctx *LhsContext) {}

// EnterSliceFilter is called when production SliceFilter is entered.
func (s *BasegoscriptListener) EnterSliceFilter(ctx *SliceFilterContext) {}

// ExitSliceFilter is called when production SliceFilter is exited.
func (s *BasegoscriptListener) ExitSliceFilter(ctx *SliceFilterContext) {}

// EnterVariableName is called when production VariableName is entered.
func (s *BasegoscriptListener) EnterVariableName(ctx *VariableNameContext) {}

// ExitVariableName is called when production VariableName is exited.
func (s *BasegoscriptListener) ExitVariableName(ctx *VariableNameContext) {}

// EnterSelect is called when production Select is entered.
func (s *BasegoscriptListener) EnterSelect(ctx *SelectContext) {}

// ExitSelect is called when production Select is exited.
func (s *BasegoscriptListener) ExitSelect(ctx *SelectContext) {}

// EnterIndex is called when production Index is entered.
func (s *BasegoscriptListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production Index is exited.
func (s *BasegoscriptListener) ExitIndex(ctx *IndexContext) {}

// EnterMapMultiIndex is called when production MapMultiIndex is entered.
func (s *BasegoscriptListener) EnterMapMultiIndex(ctx *MapMultiIndexContext) {}

// ExitMapMultiIndex is called when production MapMultiIndex is exited.
func (s *BasegoscriptListener) ExitMapMultiIndex(ctx *MapMultiIndexContext) {}

// EnterSliceMultiIndex is called when production SliceMultiIndex is entered.
func (s *BasegoscriptListener) EnterSliceMultiIndex(ctx *SliceMultiIndexContext) {}

// ExitSliceMultiIndex is called when production SliceMultiIndex is exited.
func (s *BasegoscriptListener) ExitSliceMultiIndex(ctx *SliceMultiIndexContext) {}

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

// EnterPass is called when production Pass is entered.
func (s *BasegoscriptListener) EnterPass(ctx *PassContext) {}

// ExitPass is called when production Pass is exited.
func (s *BasegoscriptListener) ExitPass(ctx *PassContext) {}

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

// EnterRightUnary is called when production RightUnary is entered.
func (s *BasegoscriptListener) EnterRightUnary(ctx *RightUnaryContext) {}

// ExitRightUnary is called when production RightUnary is exited.
func (s *BasegoscriptListener) ExitRightUnary(ctx *RightUnaryContext) {}

// EnterAssignInitializationlist is called when production AssignInitializationlist is entered.
func (s *BasegoscriptListener) EnterAssignInitializationlist(ctx *AssignInitializationlistContext) {}

// ExitAssignInitializationlist is called when production AssignInitializationlist is exited.
func (s *BasegoscriptListener) ExitAssignInitializationlist(ctx *AssignInitializationlistContext) {}

// EnterLocalDef is called when production LocalDef is entered.
func (s *BasegoscriptListener) EnterLocalDef(ctx *LocalDefContext) {}

// ExitLocalDef is called when production LocalDef is exited.
func (s *BasegoscriptListener) ExitLocalDef(ctx *LocalDefContext) {}

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

// EnterFunctions is called when production functions is entered.
func (s *BasegoscriptListener) EnterFunctions(ctx *FunctionsContext) {}

// ExitFunctions is called when production functions is exited.
func (s *BasegoscriptListener) ExitFunctions(ctx *FunctionsContext) {}

// EnterConstructor is called when production constructor is entered.
func (s *BasegoscriptListener) EnterConstructor(ctx *ConstructorContext) {}

// ExitConstructor is called when production constructor is exited.
func (s *BasegoscriptListener) ExitConstructor(ctx *ConstructorContext) {}

// EnterVardef is called when production vardef is entered.
func (s *BasegoscriptListener) EnterVardef(ctx *VardefContext) {}

// ExitVardef is called when production vardef is exited.
func (s *BasegoscriptListener) ExitVardef(ctx *VardefContext) {}

// EnterLocal is called when production Local is entered.
func (s *BasegoscriptListener) EnterLocal(ctx *LocalContext) {}

// ExitLocal is called when production Local is exited.
func (s *BasegoscriptListener) ExitLocal(ctx *LocalContext) {}

// EnterLocalAssign is called when production LocalAssign is entered.
func (s *BasegoscriptListener) EnterLocalAssign(ctx *LocalAssignContext) {}

// ExitLocalAssign is called when production LocalAssign is exited.
func (s *BasegoscriptListener) ExitLocalAssign(ctx *LocalAssignContext) {}
