// Code generated from goscript.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // goscript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// goscriptListener is a complete listener for a parse tree produced by goscriptParser.
type goscriptListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterFieldname is called when entering the fieldname production.
	EnterFieldname(c *FieldnameContext)

	// EnterFunctionDef is called when entering the FunctionDef production.
	EnterFunctionDef(c *FunctionDefContext)

	// EnterLambdaDef is called when entering the LambdaDef production.
	EnterLambdaDef(c *LambdaDefContext)

	// EnterClosure is called when entering the closure production.
	EnterClosure(c *ClosureContext)

	// EnterInparam is called when entering the inparam production.
	EnterInparam(c *InparamContext)

	// EnterOutparam is called when entering the outparam production.
	EnterOutparam(c *OutparamContext)

	// EnterIntypename is called when entering the intypename production.
	EnterIntypename(c *IntypenameContext)

	// EnterReturntypename is called when entering the returntypename production.
	EnterReturntypename(c *ReturntypenameContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterSimpleTypeName is called when entering the SimpleTypeName production.
	EnterSimpleTypeName(c *SimpleTypeNameContext)

	// EnterFunctionType is called when entering the FunctionType production.
	EnterFunctionType(c *FunctionTypeContext)

	// EnterMapTypeName is called when entering the MapTypeName production.
	EnterMapTypeName(c *MapTypeNameContext)

	// EnterSliceTypeName is called when entering the SliceTypeName production.
	EnterSliceTypeName(c *SliceTypeNameContext)

	// EnterChanTypeName is called when entering the ChanTypeName production.
	EnterChanTypeName(c *ChanTypeNameContext)

	// EnterFunctionTypeName is called when entering the functionTypeName production.
	EnterFunctionTypeName(c *FunctionTypeNameContext)

	// EnterTypeDefAlias is called when entering the TypeDefAlias production.
	EnterTypeDefAlias(c *TypeDefAliasContext)

	// EnterTypeDefMessage is called when entering the TypeDefMessage production.
	EnterTypeDefMessage(c *TypeDefMessageContext)

	// EnterTypeDefEnum is called when entering the TypeDefEnum production.
	EnterTypeDefEnum(c *TypeDefEnumContext)

	// EnterFieldDef is called when entering the FieldDef production.
	EnterFieldDef(c *FieldDefContext)

	// EnterOneofDef is called when entering the OneofDef production.
	EnterOneofDef(c *OneofDefContext)

	// EnterOneoffield is called when entering the oneoffield production.
	EnterOneoffield(c *OneoffieldContext)

	// EnterTypenameindef is called when entering the typenameindef production.
	EnterTypenameindef(c *TypenameindefContext)

	// EnterSimpleTypeNameindef is called when entering the simpleTypeNameindef production.
	EnterSimpleTypeNameindef(c *SimpleTypeNameindefContext)

	// EnterChanTypeNameindef is called when entering the chanTypeNameindef production.
	EnterChanTypeNameindef(c *ChanTypeNameindefContext)

	// EnterSliceTypeNameindef is called when entering the sliceTypeNameindef production.
	EnterSliceTypeNameindef(c *SliceTypeNameindefContext)

	// EnterMapTypeNameindef is called when entering the mapTypeNameindef production.
	EnterMapTypeNameindef(c *MapTypeNameindefContext)

	// EnterFunctionTypeNameindef is called when entering the functionTypeNameindef production.
	EnterFunctionTypeNameindef(c *FunctionTypeNameindefContext)

	// EnterIntypenameindef is called when entering the intypenameindef production.
	EnterIntypenameindef(c *IntypenameindefContext)

	// EnterReturntypenameindef is called when entering the returntypenameindef production.
	EnterReturntypenameindef(c *ReturntypenameindefContext)

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

	// EnterRestoreStackSp is called when entering the RestoreStackSp production.
	EnterRestoreStackSp(c *RestoreStackSpContext)

	// EnterVarDef is called when entering the VarDef production.
	EnterVarDef(c *VarDefContext)

	// EnterConstDef is called when entering the ConstDef production.
	EnterConstDef(c *ConstDefContext)

	// EnterRestoreStack is called when entering the restoreStack production.
	EnterRestoreStack(c *RestoreStackContext)

	// EnterExprEntry is called when entering the ExprEntry production.
	EnterExprEntry(c *ExprEntryContext)

	// EnterFunctionAssign is called when entering the FunctionAssign production.
	EnterFunctionAssign(c *FunctionAssignContext)

	// EnterVariableName is called when entering the VariableName production.
	EnterVariableName(c *VariableNameContext)

	// EnterAsserted is called when entering the asserted production.
	EnterAsserted(c *AssertedContext)

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

	// EnterRecv is called when entering the Recv production.
	EnterRecv(c *RecvContext)

	// EnterSliceFilter is called when entering the SliceFilter production.
	EnterSliceFilter(c *SliceFilterContext)

	// EnterDirectCall is called when entering the DirectCall production.
	EnterDirectCall(c *DirectCallContext)

	// EnterIndex is called when entering the Index production.
	EnterIndex(c *IndexContext)

	// EnterMapMultiIndex is called when entering the MapMultiIndex production.
	EnterMapMultiIndex(c *MapMultiIndexContext)

	// EnterSend is called when entering the Send production.
	EnterSend(c *SendContext)

	// EnterPass is called when entering the Pass production.
	EnterPass(c *PassContext)

	// EnterSelect is called when entering the Select production.
	EnterSelect(c *SelectContext)

	// EnterConstruct is called when entering the Construct production.
	EnterConstruct(c *ConstructContext)

	// EnterBinary is called when entering the Binary production.
	EnterBinary(c *BinaryContext)

	// EnterLeftUnary is called when entering the LeftUnary production.
	EnterLeftUnary(c *LeftUnaryContext)

	// EnterTypeAssert is called when entering the TypeAssert production.
	EnterTypeAssert(c *TypeAssertContext)

	// EnterSliceMultiIndex is called when entering the SliceMultiIndex production.
	EnterSliceMultiIndex(c *SliceMultiIndexContext)

	// EnterRightUnary is called when entering the RightUnary production.
	EnterRightUnary(c *RightUnaryContext)

	// EnterAssignInitializationlist is called when entering the AssignInitializationlist production.
	EnterAssignInitializationlist(c *AssignInitializationlistContext)

	// EnterBasicTypeName is called when entering the basicTypeName production.
	EnterBasicTypeName(c *BasicTypeNameContext)

	// EnterBuiltin is called when entering the builtin production.
	EnterBuiltin(c *BuiltinContext)

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

	// EnterConstructor is called when entering the constructor production.
	EnterConstructor(c *ConstructorContext)

	// EnterVardef is called when entering the vardef production.
	EnterVardef(c *VardefContext)

	// EnterConstdef is called when entering the constdef production.
	EnterConstdef(c *ConstdefContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitFieldname is called when exiting the fieldname production.
	ExitFieldname(c *FieldnameContext)

	// ExitFunctionDef is called when exiting the FunctionDef production.
	ExitFunctionDef(c *FunctionDefContext)

	// ExitLambdaDef is called when exiting the LambdaDef production.
	ExitLambdaDef(c *LambdaDefContext)

	// ExitClosure is called when exiting the closure production.
	ExitClosure(c *ClosureContext)

	// ExitInparam is called when exiting the inparam production.
	ExitInparam(c *InparamContext)

	// ExitOutparam is called when exiting the outparam production.
	ExitOutparam(c *OutparamContext)

	// ExitIntypename is called when exiting the intypename production.
	ExitIntypename(c *IntypenameContext)

	// ExitReturntypename is called when exiting the returntypename production.
	ExitReturntypename(c *ReturntypenameContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitSimpleTypeName is called when exiting the SimpleTypeName production.
	ExitSimpleTypeName(c *SimpleTypeNameContext)

	// ExitFunctionType is called when exiting the FunctionType production.
	ExitFunctionType(c *FunctionTypeContext)

	// ExitMapTypeName is called when exiting the MapTypeName production.
	ExitMapTypeName(c *MapTypeNameContext)

	// ExitSliceTypeName is called when exiting the SliceTypeName production.
	ExitSliceTypeName(c *SliceTypeNameContext)

	// ExitChanTypeName is called when exiting the ChanTypeName production.
	ExitChanTypeName(c *ChanTypeNameContext)

	// ExitFunctionTypeName is called when exiting the functionTypeName production.
	ExitFunctionTypeName(c *FunctionTypeNameContext)

	// ExitTypeDefAlias is called when exiting the TypeDefAlias production.
	ExitTypeDefAlias(c *TypeDefAliasContext)

	// ExitTypeDefMessage is called when exiting the TypeDefMessage production.
	ExitTypeDefMessage(c *TypeDefMessageContext)

	// ExitTypeDefEnum is called when exiting the TypeDefEnum production.
	ExitTypeDefEnum(c *TypeDefEnumContext)

	// ExitFieldDef is called when exiting the FieldDef production.
	ExitFieldDef(c *FieldDefContext)

	// ExitOneofDef is called when exiting the OneofDef production.
	ExitOneofDef(c *OneofDefContext)

	// ExitOneoffield is called when exiting the oneoffield production.
	ExitOneoffield(c *OneoffieldContext)

	// ExitTypenameindef is called when exiting the typenameindef production.
	ExitTypenameindef(c *TypenameindefContext)

	// ExitSimpleTypeNameindef is called when exiting the simpleTypeNameindef production.
	ExitSimpleTypeNameindef(c *SimpleTypeNameindefContext)

	// ExitChanTypeNameindef is called when exiting the chanTypeNameindef production.
	ExitChanTypeNameindef(c *ChanTypeNameindefContext)

	// ExitSliceTypeNameindef is called when exiting the sliceTypeNameindef production.
	ExitSliceTypeNameindef(c *SliceTypeNameindefContext)

	// ExitMapTypeNameindef is called when exiting the mapTypeNameindef production.
	ExitMapTypeNameindef(c *MapTypeNameindefContext)

	// ExitFunctionTypeNameindef is called when exiting the functionTypeNameindef production.
	ExitFunctionTypeNameindef(c *FunctionTypeNameindefContext)

	// ExitIntypenameindef is called when exiting the intypenameindef production.
	ExitIntypenameindef(c *IntypenameindefContext)

	// ExitReturntypenameindef is called when exiting the returntypenameindef production.
	ExitReturntypenameindef(c *ReturntypenameindefContext)

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

	// ExitRestoreStackSp is called when exiting the RestoreStackSp production.
	ExitRestoreStackSp(c *RestoreStackSpContext)

	// ExitVarDef is called when exiting the VarDef production.
	ExitVarDef(c *VarDefContext)

	// ExitConstDef is called when exiting the ConstDef production.
	ExitConstDef(c *ConstDefContext)

	// ExitRestoreStack is called when exiting the restoreStack production.
	ExitRestoreStack(c *RestoreStackContext)

	// ExitExprEntry is called when exiting the ExprEntry production.
	ExitExprEntry(c *ExprEntryContext)

	// ExitFunctionAssign is called when exiting the FunctionAssign production.
	ExitFunctionAssign(c *FunctionAssignContext)

	// ExitVariableName is called when exiting the VariableName production.
	ExitVariableName(c *VariableNameContext)

	// ExitAsserted is called when exiting the asserted production.
	ExitAsserted(c *AssertedContext)

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

	// ExitRecv is called when exiting the Recv production.
	ExitRecv(c *RecvContext)

	// ExitSliceFilter is called when exiting the SliceFilter production.
	ExitSliceFilter(c *SliceFilterContext)

	// ExitDirectCall is called when exiting the DirectCall production.
	ExitDirectCall(c *DirectCallContext)

	// ExitIndex is called when exiting the Index production.
	ExitIndex(c *IndexContext)

	// ExitMapMultiIndex is called when exiting the MapMultiIndex production.
	ExitMapMultiIndex(c *MapMultiIndexContext)

	// ExitSend is called when exiting the Send production.
	ExitSend(c *SendContext)

	// ExitPass is called when exiting the Pass production.
	ExitPass(c *PassContext)

	// ExitSelect is called when exiting the Select production.
	ExitSelect(c *SelectContext)

	// ExitConstruct is called when exiting the Construct production.
	ExitConstruct(c *ConstructContext)

	// ExitBinary is called when exiting the Binary production.
	ExitBinary(c *BinaryContext)

	// ExitLeftUnary is called when exiting the LeftUnary production.
	ExitLeftUnary(c *LeftUnaryContext)

	// ExitTypeAssert is called when exiting the TypeAssert production.
	ExitTypeAssert(c *TypeAssertContext)

	// ExitSliceMultiIndex is called when exiting the SliceMultiIndex production.
	ExitSliceMultiIndex(c *SliceMultiIndexContext)

	// ExitRightUnary is called when exiting the RightUnary production.
	ExitRightUnary(c *RightUnaryContext)

	// ExitAssignInitializationlist is called when exiting the AssignInitializationlist production.
	ExitAssignInitializationlist(c *AssignInitializationlistContext)

	// ExitBasicTypeName is called when exiting the basicTypeName production.
	ExitBasicTypeName(c *BasicTypeNameContext)

	// ExitBuiltin is called when exiting the builtin production.
	ExitBuiltin(c *BuiltinContext)

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

	// ExitConstructor is called when exiting the constructor production.
	ExitConstructor(c *ConstructorContext)

	// ExitVardef is called when exiting the vardef production.
	ExitVardef(c *VardefContext)

	// ExitConstdef is called when exiting the constdef production.
	ExitConstdef(c *ConstdefContext)
}
