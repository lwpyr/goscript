// Code generated from goscript.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // goscript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// goscriptListener is a complete listener for a parse tree produced by goscriptParser.
type goscriptListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterFuncDef is called when entering the FuncDef production.
	EnterFuncDef(c *FuncDefContext)

	// EnterTypeDef is called when entering the TypeDef production.
	EnterTypeDef(c *TypeDefContext)

	// EnterGlobalDef is called when entering the GlobalDef production.
	EnterGlobalDef(c *GlobalDefContext)

	// EnterEnumDef is called when entering the EnumDef production.
	EnterEnumDef(c *EnumDefContext)

	// EnterRun is called when entering the Run production.
	EnterRun(c *RunContext)

	// EnterFunctionDef is called when entering the FunctionDef production.
	EnterFunctionDef(c *FunctionDefContext)

	// EnterFunctionDefWithReturnName is called when entering the FunctionDefWithReturnName production.
	EnterFunctionDefWithReturnName(c *FunctionDefWithReturnNameContext)

	// EnterInparam is called when entering the inparam production.
	EnterInparam(c *InparamContext)

	// EnterOutparam is called when entering the outparam production.
	EnterOutparam(c *OutparamContext)

	// EnterReturntypename is called when entering the returntypename production.
	EnterReturntypename(c *ReturntypenameContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterTypeDefAlias is called when entering the TypeDefAlias production.
	EnterTypeDefAlias(c *TypeDefAliasContext)

	// EnterTypeDefComplex is called when entering the TypeDefComplex production.
	EnterTypeDefComplex(c *TypeDefComplexContext)

	// EnterSimpleTypeName is called when entering the SimpleTypeName production.
	EnterSimpleTypeName(c *SimpleTypeNameContext)

	// EnterMapTypeName is called when entering the MapTypeName production.
	EnterMapTypeName(c *MapTypeNameContext)

	// EnterSliceTypeName is called when entering the SliceTypeName production.
	EnterSliceTypeName(c *SliceTypeNameContext)

	// EnterEnumdef is called when entering the enumdef production.
	EnterEnumdef(c *EnumdefContext)

	// EnterCtrl is called when entering the Ctrl production.
	EnterCtrl(c *CtrlContext)

	// EnterLineProgram is called when entering the LineProgram production.
	EnterLineProgram(c *LineProgramContext)

	// EnterIf is called when entering the If production.
	EnterIf(c *IfContext)

	// EnterSwitch is called when entering the Switch production.
	EnterSwitch(c *SwitchContext)

	// EnterForInSlice is called when entering the ForInSlice production.
	EnterForInSlice(c *ForInSliceContext)

	// EnterForInMap is called when entering the ForInMap production.
	EnterForInMap(c *ForInMapContext)

	// EnterFor is called when entering the For production.
	EnterFor(c *ForContext)

	// EnterBreak is called when entering the Break production.
	EnterBreak(c *BreakContext)

	// EnterContinue is called when entering the Continue production.
	EnterContinue(c *ContinueContext)

	// EnterReturnVoid is called when entering the ReturnVoid production.
	EnterReturnVoid(c *ReturnVoidContext)

	// EnterReturnVal is called when entering the ReturnVal production.
	EnterReturnVal(c *ReturnValContext)

	// EnterCollection is called when entering the collection production.
	EnterCollection(c *CollectionContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterExprEntry is called when entering the ExprEntry production.
	EnterExprEntry(c *ExprEntryContext)

	// EnterFunctionAssign is called when entering the FunctionAssign production.
	EnterFunctionAssign(c *FunctionAssignContext)

	// EnterMultiAssign is called when entering the MultiAssign production.
	EnterMultiAssign(c *MultiAssignContext)

	// EnterLhs is called when entering the lhs production.
	EnterLhs(c *LhsContext)

	// EnterSliceFilter is called when entering the SliceFilter production.
	EnterSliceFilter(c *SliceFilterContext)

	// EnterVariableName is called when entering the VariableName production.
	EnterVariableName(c *VariableNameContext)

	// EnterSelect is called when entering the Select production.
	EnterSelect(c *SelectContext)

	// EnterIndex is called when entering the Index production.
	EnterIndex(c *IndexContext)

	// EnterMapMultiIndex is called when entering the MapMultiIndex production.
	EnterMapMultiIndex(c *MapMultiIndexContext)

	// EnterSliceMultiIndex is called when entering the SliceMultiIndex production.
	EnterSliceMultiIndex(c *SliceMultiIndexContext)

	// EnterFilter is called when entering the filter production.
	EnterFilter(c *FilterContext)

	// EnterIndexType1 is called when entering the IndexType1 production.
	EnterIndexType1(c *IndexType1Context)

	// EnterIndexType2 is called when entering the IndexType2 production.
	EnterIndexType2(c *IndexType2Context)

	// EnterIndexType3 is called when entering the IndexType3 production.
	EnterIndexType3(c *IndexType3Context)

	// EnterIndexType4 is called when entering the IndexType4 production.
	EnterIndexType4(c *IndexType4Context)

	// EnterIndexType5 is called when entering the IndexType5 production.
	EnterIndexType5(c *IndexType5Context)

	// EnterPass is called when entering the Pass production.
	EnterPass(c *PassContext)

	// EnterConstruct is called when entering the Construct production.
	EnterConstruct(c *ConstructContext)

	// EnterBinary is called when entering the Binary production.
	EnterBinary(c *BinaryContext)

	// EnterLeftUnary is called when entering the LeftUnary production.
	EnterLeftUnary(c *LeftUnaryContext)

	// EnterRightUnary is called when entering the RightUnary production.
	EnterRightUnary(c *RightUnaryContext)

	// EnterAssignInitializationlist is called when entering the AssignInitializationlist production.
	EnterAssignInitializationlist(c *AssignInitializationlistContext)

	// EnterLocalDef is called when entering the LocalDef production.
	EnterLocalDef(c *LocalDefContext)

	// EnterInitializationListBegin is called when entering the initializationListBegin production.
	EnterInitializationListBegin(c *InitializationListBeginContext)

	// EnterInitSlice is called when entering the InitSlice production.
	EnterInitSlice(c *InitSliceContext)

	// EnterInitMessage is called when entering the InitMessage production.
	EnterInitMessage(c *InitMessageContext)

	// EnterInitMap is called when entering the InitMap production.
	EnterInitMap(c *InitMapContext)

	// EnterInitConstant is called when entering the InitConstant production.
	EnterInitConstant(c *InitConstantContext)

	// EnterConstantInt is called when entering the ConstantInt production.
	EnterConstantInt(c *ConstantIntContext)

	// EnterConstantFloat is called when entering the ConstantFloat production.
	EnterConstantFloat(c *ConstantFloatContext)

	// EnterConstantBool is called when entering the ConstantBool production.
	EnterConstantBool(c *ConstantBoolContext)

	// EnterConstantNil is called when entering the ConstantNil production.
	EnterConstantNil(c *ConstantNilContext)

	// EnterConstantString is called when entering the ConstantString production.
	EnterConstantString(c *ConstantStringContext)

	// EnterFunctions is called when entering the functions production.
	EnterFunctions(c *FunctionsContext)

	// EnterConstructor is called when entering the constructor production.
	EnterConstructor(c *ConstructorContext)

	// EnterVardef is called when entering the vardef production.
	EnterVardef(c *VardefContext)

	// EnterLocal is called when entering the Local production.
	EnterLocal(c *LocalContext)

	// EnterLocalAssign is called when entering the LocalAssign production.
	EnterLocalAssign(c *LocalAssignContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitFuncDef is called when exiting the FuncDef production.
	ExitFuncDef(c *FuncDefContext)

	// ExitTypeDef is called when exiting the TypeDef production.
	ExitTypeDef(c *TypeDefContext)

	// ExitGlobalDef is called when exiting the GlobalDef production.
	ExitGlobalDef(c *GlobalDefContext)

	// ExitEnumDef is called when exiting the EnumDef production.
	ExitEnumDef(c *EnumDefContext)

	// ExitRun is called when exiting the Run production.
	ExitRun(c *RunContext)

	// ExitFunctionDef is called when exiting the FunctionDef production.
	ExitFunctionDef(c *FunctionDefContext)

	// ExitFunctionDefWithReturnName is called when exiting the FunctionDefWithReturnName production.
	ExitFunctionDefWithReturnName(c *FunctionDefWithReturnNameContext)

	// ExitInparam is called when exiting the inparam production.
	ExitInparam(c *InparamContext)

	// ExitOutparam is called when exiting the outparam production.
	ExitOutparam(c *OutparamContext)

	// ExitReturntypename is called when exiting the returntypename production.
	ExitReturntypename(c *ReturntypenameContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitTypeDefAlias is called when exiting the TypeDefAlias production.
	ExitTypeDefAlias(c *TypeDefAliasContext)

	// ExitTypeDefComplex is called when exiting the TypeDefComplex production.
	ExitTypeDefComplex(c *TypeDefComplexContext)

	// ExitSimpleTypeName is called when exiting the SimpleTypeName production.
	ExitSimpleTypeName(c *SimpleTypeNameContext)

	// ExitMapTypeName is called when exiting the MapTypeName production.
	ExitMapTypeName(c *MapTypeNameContext)

	// ExitSliceTypeName is called when exiting the SliceTypeName production.
	ExitSliceTypeName(c *SliceTypeNameContext)

	// ExitEnumdef is called when exiting the enumdef production.
	ExitEnumdef(c *EnumdefContext)

	// ExitCtrl is called when exiting the Ctrl production.
	ExitCtrl(c *CtrlContext)

	// ExitLineProgram is called when exiting the LineProgram production.
	ExitLineProgram(c *LineProgramContext)

	// ExitIf is called when exiting the If production.
	ExitIf(c *IfContext)

	// ExitSwitch is called when exiting the Switch production.
	ExitSwitch(c *SwitchContext)

	// ExitForInSlice is called when exiting the ForInSlice production.
	ExitForInSlice(c *ForInSliceContext)

	// ExitForInMap is called when exiting the ForInMap production.
	ExitForInMap(c *ForInMapContext)

	// ExitFor is called when exiting the For production.
	ExitFor(c *ForContext)

	// ExitBreak is called when exiting the Break production.
	ExitBreak(c *BreakContext)

	// ExitContinue is called when exiting the Continue production.
	ExitContinue(c *ContinueContext)

	// ExitReturnVoid is called when exiting the ReturnVoid production.
	ExitReturnVoid(c *ReturnVoidContext)

	// ExitReturnVal is called when exiting the ReturnVal production.
	ExitReturnVal(c *ReturnValContext)

	// ExitCollection is called when exiting the collection production.
	ExitCollection(c *CollectionContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitExprEntry is called when exiting the ExprEntry production.
	ExitExprEntry(c *ExprEntryContext)

	// ExitFunctionAssign is called when exiting the FunctionAssign production.
	ExitFunctionAssign(c *FunctionAssignContext)

	// ExitMultiAssign is called when exiting the MultiAssign production.
	ExitMultiAssign(c *MultiAssignContext)

	// ExitLhs is called when exiting the lhs production.
	ExitLhs(c *LhsContext)

	// ExitSliceFilter is called when exiting the SliceFilter production.
	ExitSliceFilter(c *SliceFilterContext)

	// ExitVariableName is called when exiting the VariableName production.
	ExitVariableName(c *VariableNameContext)

	// ExitSelect is called when exiting the Select production.
	ExitSelect(c *SelectContext)

	// ExitIndex is called when exiting the Index production.
	ExitIndex(c *IndexContext)

	// ExitMapMultiIndex is called when exiting the MapMultiIndex production.
	ExitMapMultiIndex(c *MapMultiIndexContext)

	// ExitSliceMultiIndex is called when exiting the SliceMultiIndex production.
	ExitSliceMultiIndex(c *SliceMultiIndexContext)

	// ExitFilter is called when exiting the filter production.
	ExitFilter(c *FilterContext)

	// ExitIndexType1 is called when exiting the IndexType1 production.
	ExitIndexType1(c *IndexType1Context)

	// ExitIndexType2 is called when exiting the IndexType2 production.
	ExitIndexType2(c *IndexType2Context)

	// ExitIndexType3 is called when exiting the IndexType3 production.
	ExitIndexType3(c *IndexType3Context)

	// ExitIndexType4 is called when exiting the IndexType4 production.
	ExitIndexType4(c *IndexType4Context)

	// ExitIndexType5 is called when exiting the IndexType5 production.
	ExitIndexType5(c *IndexType5Context)

	// ExitPass is called when exiting the Pass production.
	ExitPass(c *PassContext)

	// ExitConstruct is called when exiting the Construct production.
	ExitConstruct(c *ConstructContext)

	// ExitBinary is called when exiting the Binary production.
	ExitBinary(c *BinaryContext)

	// ExitLeftUnary is called when exiting the LeftUnary production.
	ExitLeftUnary(c *LeftUnaryContext)

	// ExitRightUnary is called when exiting the RightUnary production.
	ExitRightUnary(c *RightUnaryContext)

	// ExitAssignInitializationlist is called when exiting the AssignInitializationlist production.
	ExitAssignInitializationlist(c *AssignInitializationlistContext)

	// ExitLocalDef is called when exiting the LocalDef production.
	ExitLocalDef(c *LocalDefContext)

	// ExitInitializationListBegin is called when exiting the initializationListBegin production.
	ExitInitializationListBegin(c *InitializationListBeginContext)

	// ExitInitSlice is called when exiting the InitSlice production.
	ExitInitSlice(c *InitSliceContext)

	// ExitInitMessage is called when exiting the InitMessage production.
	ExitInitMessage(c *InitMessageContext)

	// ExitInitMap is called when exiting the InitMap production.
	ExitInitMap(c *InitMapContext)

	// ExitInitConstant is called when exiting the InitConstant production.
	ExitInitConstant(c *InitConstantContext)

	// ExitConstantInt is called when exiting the ConstantInt production.
	ExitConstantInt(c *ConstantIntContext)

	// ExitConstantFloat is called when exiting the ConstantFloat production.
	ExitConstantFloat(c *ConstantFloatContext)

	// ExitConstantBool is called when exiting the ConstantBool production.
	ExitConstantBool(c *ConstantBoolContext)

	// ExitConstantNil is called when exiting the ConstantNil production.
	ExitConstantNil(c *ConstantNilContext)

	// ExitConstantString is called when exiting the ConstantString production.
	ExitConstantString(c *ConstantStringContext)

	// ExitFunctions is called when exiting the functions production.
	ExitFunctions(c *FunctionsContext)

	// ExitConstructor is called when exiting the constructor production.
	ExitConstructor(c *ConstructorContext)

	// ExitVardef is called when exiting the vardef production.
	ExitVardef(c *VardefContext)

	// ExitLocal is called when exiting the Local production.
	ExitLocal(c *LocalContext)

	// ExitLocalAssign is called when exiting the LocalAssign production.
	ExitLocalAssign(c *LocalAssignContext)
}
