// Code generated from ./qasm3.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // qasm3

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Baseqasm3Listener is a complete listener for a parse tree produced by qasm3Parser.
type Baseqasm3Listener struct{}

var _ qasm3Listener = &Baseqasm3Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Baseqasm3Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Baseqasm3Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Baseqasm3Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Baseqasm3Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *Baseqasm3Listener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *Baseqasm3Listener) ExitProgram(ctx *ProgramContext) {}

// EnterHeader is called when production header is entered.
func (s *Baseqasm3Listener) EnterHeader(ctx *HeaderContext) {}

// ExitHeader is called when production header is exited.
func (s *Baseqasm3Listener) ExitHeader(ctx *HeaderContext) {}

// EnterVersion is called when production version is entered.
func (s *Baseqasm3Listener) EnterVersion(ctx *VersionContext) {}

// ExitVersion is called when production version is exited.
func (s *Baseqasm3Listener) ExitVersion(ctx *VersionContext) {}

// EnterInclude is called when production include is entered.
func (s *Baseqasm3Listener) EnterInclude(ctx *IncludeContext) {}

// ExitInclude is called when production include is exited.
func (s *Baseqasm3Listener) ExitInclude(ctx *IncludeContext) {}

// EnterStatement is called when production statement is entered.
func (s *Baseqasm3Listener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *Baseqasm3Listener) ExitStatement(ctx *StatementContext) {}

// EnterGlobalStatement is called when production globalStatement is entered.
func (s *Baseqasm3Listener) EnterGlobalStatement(ctx *GlobalStatementContext) {}

// ExitGlobalStatement is called when production globalStatement is exited.
func (s *Baseqasm3Listener) ExitGlobalStatement(ctx *GlobalStatementContext) {}

// EnterDeclarationStatement is called when production declarationStatement is entered.
func (s *Baseqasm3Listener) EnterDeclarationStatement(ctx *DeclarationStatementContext) {}

// ExitDeclarationStatement is called when production declarationStatement is exited.
func (s *Baseqasm3Listener) ExitDeclarationStatement(ctx *DeclarationStatementContext) {}

// EnterComment is called when production comment is entered.
func (s *Baseqasm3Listener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *Baseqasm3Listener) ExitComment(ctx *CommentContext) {}

// EnterReturnSignature is called when production returnSignature is entered.
func (s *Baseqasm3Listener) EnterReturnSignature(ctx *ReturnSignatureContext) {}

// ExitReturnSignature is called when production returnSignature is exited.
func (s *Baseqasm3Listener) ExitReturnSignature(ctx *ReturnSignatureContext) {}

// EnterProgramBlock is called when production programBlock is entered.
func (s *Baseqasm3Listener) EnterProgramBlock(ctx *ProgramBlockContext) {}

// ExitProgramBlock is called when production programBlock is exited.
func (s *Baseqasm3Listener) ExitProgramBlock(ctx *ProgramBlockContext) {}

// EnterDesignator is called when production designator is entered.
func (s *Baseqasm3Listener) EnterDesignator(ctx *DesignatorContext) {}

// ExitDesignator is called when production designator is exited.
func (s *Baseqasm3Listener) ExitDesignator(ctx *DesignatorContext) {}

// EnterDoubleDesignator is called when production doubleDesignator is entered.
func (s *Baseqasm3Listener) EnterDoubleDesignator(ctx *DoubleDesignatorContext) {}

// ExitDoubleDesignator is called when production doubleDesignator is exited.
func (s *Baseqasm3Listener) ExitDoubleDesignator(ctx *DoubleDesignatorContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *Baseqasm3Listener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *Baseqasm3Listener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterIndexIdentifier is called when production indexIdentifier is entered.
func (s *Baseqasm3Listener) EnterIndexIdentifier(ctx *IndexIdentifierContext) {}

// ExitIndexIdentifier is called when production indexIdentifier is exited.
func (s *Baseqasm3Listener) ExitIndexIdentifier(ctx *IndexIdentifierContext) {}

// EnterIndexIdentifierList is called when production indexIdentifierList is entered.
func (s *Baseqasm3Listener) EnterIndexIdentifierList(ctx *IndexIdentifierListContext) {}

// ExitIndexIdentifierList is called when production indexIdentifierList is exited.
func (s *Baseqasm3Listener) ExitIndexIdentifierList(ctx *IndexIdentifierListContext) {}

// EnterAssociation is called when production association is entered.
func (s *Baseqasm3Listener) EnterAssociation(ctx *AssociationContext) {}

// ExitAssociation is called when production association is exited.
func (s *Baseqasm3Listener) ExitAssociation(ctx *AssociationContext) {}

// EnterQuantumType is called when production quantumType is entered.
func (s *Baseqasm3Listener) EnterQuantumType(ctx *QuantumTypeContext) {}

// ExitQuantumType is called when production quantumType is exited.
func (s *Baseqasm3Listener) ExitQuantumType(ctx *QuantumTypeContext) {}

// EnterQuantumDeclaration is called when production quantumDeclaration is entered.
func (s *Baseqasm3Listener) EnterQuantumDeclaration(ctx *QuantumDeclarationContext) {}

// ExitQuantumDeclaration is called when production quantumDeclaration is exited.
func (s *Baseqasm3Listener) ExitQuantumDeclaration(ctx *QuantumDeclarationContext) {}

// EnterQuantumArgument is called when production quantumArgument is entered.
func (s *Baseqasm3Listener) EnterQuantumArgument(ctx *QuantumArgumentContext) {}

// ExitQuantumArgument is called when production quantumArgument is exited.
func (s *Baseqasm3Listener) ExitQuantumArgument(ctx *QuantumArgumentContext) {}

// EnterQuantumArgumentList is called when production quantumArgumentList is entered.
func (s *Baseqasm3Listener) EnterQuantumArgumentList(ctx *QuantumArgumentListContext) {}

// ExitQuantumArgumentList is called when production quantumArgumentList is exited.
func (s *Baseqasm3Listener) ExitQuantumArgumentList(ctx *QuantumArgumentListContext) {}

// EnterBitType is called when production bitType is entered.
func (s *Baseqasm3Listener) EnterBitType(ctx *BitTypeContext) {}

// ExitBitType is called when production bitType is exited.
func (s *Baseqasm3Listener) ExitBitType(ctx *BitTypeContext) {}

// EnterSingleDesignatorType is called when production singleDesignatorType is entered.
func (s *Baseqasm3Listener) EnterSingleDesignatorType(ctx *SingleDesignatorTypeContext) {}

// ExitSingleDesignatorType is called when production singleDesignatorType is exited.
func (s *Baseqasm3Listener) ExitSingleDesignatorType(ctx *SingleDesignatorTypeContext) {}

// EnterDoubleDesignatorType is called when production doubleDesignatorType is entered.
func (s *Baseqasm3Listener) EnterDoubleDesignatorType(ctx *DoubleDesignatorTypeContext) {}

// ExitDoubleDesignatorType is called when production doubleDesignatorType is exited.
func (s *Baseqasm3Listener) ExitDoubleDesignatorType(ctx *DoubleDesignatorTypeContext) {}

// EnterNoDesignatorType is called when production noDesignatorType is entered.
func (s *Baseqasm3Listener) EnterNoDesignatorType(ctx *NoDesignatorTypeContext) {}

// ExitNoDesignatorType is called when production noDesignatorType is exited.
func (s *Baseqasm3Listener) ExitNoDesignatorType(ctx *NoDesignatorTypeContext) {}

// EnterClassicalType is called when production classicalType is entered.
func (s *Baseqasm3Listener) EnterClassicalType(ctx *ClassicalTypeContext) {}

// ExitClassicalType is called when production classicalType is exited.
func (s *Baseqasm3Listener) ExitClassicalType(ctx *ClassicalTypeContext) {}

// EnterConstantDeclaration is called when production constantDeclaration is entered.
func (s *Baseqasm3Listener) EnterConstantDeclaration(ctx *ConstantDeclarationContext) {}

// ExitConstantDeclaration is called when production constantDeclaration is exited.
func (s *Baseqasm3Listener) ExitConstantDeclaration(ctx *ConstantDeclarationContext) {}

// EnterSingleDesignatorDeclaration is called when production singleDesignatorDeclaration is entered.
func (s *Baseqasm3Listener) EnterSingleDesignatorDeclaration(ctx *SingleDesignatorDeclarationContext) {
}

// ExitSingleDesignatorDeclaration is called when production singleDesignatorDeclaration is exited.
func (s *Baseqasm3Listener) ExitSingleDesignatorDeclaration(ctx *SingleDesignatorDeclarationContext) {
}

// EnterDoubleDesignatorDeclaration is called when production doubleDesignatorDeclaration is entered.
func (s *Baseqasm3Listener) EnterDoubleDesignatorDeclaration(ctx *DoubleDesignatorDeclarationContext) {
}

// ExitDoubleDesignatorDeclaration is called when production doubleDesignatorDeclaration is exited.
func (s *Baseqasm3Listener) ExitDoubleDesignatorDeclaration(ctx *DoubleDesignatorDeclarationContext) {
}

// EnterNoDesignatorDeclaration is called when production noDesignatorDeclaration is entered.
func (s *Baseqasm3Listener) EnterNoDesignatorDeclaration(ctx *NoDesignatorDeclarationContext) {}

// ExitNoDesignatorDeclaration is called when production noDesignatorDeclaration is exited.
func (s *Baseqasm3Listener) ExitNoDesignatorDeclaration(ctx *NoDesignatorDeclarationContext) {}

// EnterBitDeclaration is called when production bitDeclaration is entered.
func (s *Baseqasm3Listener) EnterBitDeclaration(ctx *BitDeclarationContext) {}

// ExitBitDeclaration is called when production bitDeclaration is exited.
func (s *Baseqasm3Listener) ExitBitDeclaration(ctx *BitDeclarationContext) {}

// EnterClassicalVariableDeclaration is called when production classicalVariableDeclaration is entered.
func (s *Baseqasm3Listener) EnterClassicalVariableDeclaration(ctx *ClassicalVariableDeclarationContext) {
}

// ExitClassicalVariableDeclaration is called when production classicalVariableDeclaration is exited.
func (s *Baseqasm3Listener) ExitClassicalVariableDeclaration(ctx *ClassicalVariableDeclarationContext) {
}

// EnterClassicalDeclaration is called when production classicalDeclaration is entered.
func (s *Baseqasm3Listener) EnterClassicalDeclaration(ctx *ClassicalDeclarationContext) {}

// ExitClassicalDeclaration is called when production classicalDeclaration is exited.
func (s *Baseqasm3Listener) ExitClassicalDeclaration(ctx *ClassicalDeclarationContext) {}

// EnterClassicalTypeList is called when production classicalTypeList is entered.
func (s *Baseqasm3Listener) EnterClassicalTypeList(ctx *ClassicalTypeListContext) {}

// ExitClassicalTypeList is called when production classicalTypeList is exited.
func (s *Baseqasm3Listener) ExitClassicalTypeList(ctx *ClassicalTypeListContext) {}

// EnterClassicalArgument is called when production classicalArgument is entered.
func (s *Baseqasm3Listener) EnterClassicalArgument(ctx *ClassicalArgumentContext) {}

// ExitClassicalArgument is called when production classicalArgument is exited.
func (s *Baseqasm3Listener) ExitClassicalArgument(ctx *ClassicalArgumentContext) {}

// EnterClassicalArgumentList is called when production classicalArgumentList is entered.
func (s *Baseqasm3Listener) EnterClassicalArgumentList(ctx *ClassicalArgumentListContext) {}

// ExitClassicalArgumentList is called when production classicalArgumentList is exited.
func (s *Baseqasm3Listener) ExitClassicalArgumentList(ctx *ClassicalArgumentListContext) {}

// EnterAliasStatement is called when production aliasStatement is entered.
func (s *Baseqasm3Listener) EnterAliasStatement(ctx *AliasStatementContext) {}

// ExitAliasStatement is called when production aliasStatement is exited.
func (s *Baseqasm3Listener) ExitAliasStatement(ctx *AliasStatementContext) {}

// EnterConcatenateExpression is called when production concatenateExpression is entered.
func (s *Baseqasm3Listener) EnterConcatenateExpression(ctx *ConcatenateExpressionContext) {}

// ExitConcatenateExpression is called when production concatenateExpression is exited.
func (s *Baseqasm3Listener) ExitConcatenateExpression(ctx *ConcatenateExpressionContext) {}

// EnterRangeDefinition is called when production rangeDefinition is entered.
func (s *Baseqasm3Listener) EnterRangeDefinition(ctx *RangeDefinitionContext) {}

// ExitRangeDefinition is called when production rangeDefinition is exited.
func (s *Baseqasm3Listener) ExitRangeDefinition(ctx *RangeDefinitionContext) {}

// EnterQuantumGateDefinition is called when production quantumGateDefinition is entered.
func (s *Baseqasm3Listener) EnterQuantumGateDefinition(ctx *QuantumGateDefinitionContext) {}

// ExitQuantumGateDefinition is called when production quantumGateDefinition is exited.
func (s *Baseqasm3Listener) ExitQuantumGateDefinition(ctx *QuantumGateDefinitionContext) {}

// EnterQuantumGateSignature is called when production quantumGateSignature is entered.
func (s *Baseqasm3Listener) EnterQuantumGateSignature(ctx *QuantumGateSignatureContext) {}

// ExitQuantumGateSignature is called when production quantumGateSignature is exited.
func (s *Baseqasm3Listener) ExitQuantumGateSignature(ctx *QuantumGateSignatureContext) {}

// EnterQuantumBlock is called when production quantumBlock is entered.
func (s *Baseqasm3Listener) EnterQuantumBlock(ctx *QuantumBlockContext) {}

// ExitQuantumBlock is called when production quantumBlock is exited.
func (s *Baseqasm3Listener) ExitQuantumBlock(ctx *QuantumBlockContext) {}

// EnterQuantumStatement is called when production quantumStatement is entered.
func (s *Baseqasm3Listener) EnterQuantumStatement(ctx *QuantumStatementContext) {}

// ExitQuantumStatement is called when production quantumStatement is exited.
func (s *Baseqasm3Listener) ExitQuantumStatement(ctx *QuantumStatementContext) {}

// EnterQuantumInstruction is called when production quantumInstruction is entered.
func (s *Baseqasm3Listener) EnterQuantumInstruction(ctx *QuantumInstructionContext) {}

// ExitQuantumInstruction is called when production quantumInstruction is exited.
func (s *Baseqasm3Listener) ExitQuantumInstruction(ctx *QuantumInstructionContext) {}

// EnterQuantumMeasurement is called when production quantumMeasurement is entered.
func (s *Baseqasm3Listener) EnterQuantumMeasurement(ctx *QuantumMeasurementContext) {}

// ExitQuantumMeasurement is called when production quantumMeasurement is exited.
func (s *Baseqasm3Listener) ExitQuantumMeasurement(ctx *QuantumMeasurementContext) {}

// EnterQuantumMeasurementDeclaration is called when production quantumMeasurementDeclaration is entered.
func (s *Baseqasm3Listener) EnterQuantumMeasurementDeclaration(ctx *QuantumMeasurementDeclarationContext) {
}

// ExitQuantumMeasurementDeclaration is called when production quantumMeasurementDeclaration is exited.
func (s *Baseqasm3Listener) ExitQuantumMeasurementDeclaration(ctx *QuantumMeasurementDeclarationContext) {
}

// EnterQuantumBarrier is called when production quantumBarrier is entered.
func (s *Baseqasm3Listener) EnterQuantumBarrier(ctx *QuantumBarrierContext) {}

// ExitQuantumBarrier is called when production quantumBarrier is exited.
func (s *Baseqasm3Listener) ExitQuantumBarrier(ctx *QuantumBarrierContext) {}

// EnterQuantumGateModifier is called when production quantumGateModifier is entered.
func (s *Baseqasm3Listener) EnterQuantumGateModifier(ctx *QuantumGateModifierContext) {}

// ExitQuantumGateModifier is called when production quantumGateModifier is exited.
func (s *Baseqasm3Listener) ExitQuantumGateModifier(ctx *QuantumGateModifierContext) {}

// EnterQuantumGateCall is called when production quantumGateCall is entered.
func (s *Baseqasm3Listener) EnterQuantumGateCall(ctx *QuantumGateCallContext) {}

// ExitQuantumGateCall is called when production quantumGateCall is exited.
func (s *Baseqasm3Listener) ExitQuantumGateCall(ctx *QuantumGateCallContext) {}

// EnterQuantumGateName is called when production quantumGateName is entered.
func (s *Baseqasm3Listener) EnterQuantumGateName(ctx *QuantumGateNameContext) {}

// ExitQuantumGateName is called when production quantumGateName is exited.
func (s *Baseqasm3Listener) ExitQuantumGateName(ctx *QuantumGateNameContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *Baseqasm3Listener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *Baseqasm3Listener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterBinaryOperator is called when production binaryOperator is entered.
func (s *Baseqasm3Listener) EnterBinaryOperator(ctx *BinaryOperatorContext) {}

// ExitBinaryOperator is called when production binaryOperator is exited.
func (s *Baseqasm3Listener) ExitBinaryOperator(ctx *BinaryOperatorContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *Baseqasm3Listener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *Baseqasm3Listener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *Baseqasm3Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *Baseqasm3Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpressionTerminator is called when production expressionTerminator is entered.
func (s *Baseqasm3Listener) EnterExpressionTerminator(ctx *ExpressionTerminatorContext) {}

// ExitExpressionTerminator is called when production expressionTerminator is exited.
func (s *Baseqasm3Listener) ExitExpressionTerminator(ctx *ExpressionTerminatorContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *Baseqasm3Listener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *Baseqasm3Listener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterCall is called when production call is entered.
func (s *Baseqasm3Listener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *Baseqasm3Listener) ExitCall(ctx *CallContext) {}

// EnterBuiltInMath is called when production builtInMath is entered.
func (s *Baseqasm3Listener) EnterBuiltInMath(ctx *BuiltInMathContext) {}

// ExitBuiltInMath is called when production builtInMath is exited.
func (s *Baseqasm3Listener) ExitBuiltInMath(ctx *BuiltInMathContext) {}

// EnterCastOperator is called when production castOperator is entered.
func (s *Baseqasm3Listener) EnterCastOperator(ctx *CastOperatorContext) {}

// ExitCastOperator is called when production castOperator is exited.
func (s *Baseqasm3Listener) ExitCastOperator(ctx *CastOperatorContext) {}

// EnterIncrementor is called when production incrementor is entered.
func (s *Baseqasm3Listener) EnterIncrementor(ctx *IncrementorContext) {}

// ExitIncrementor is called when production incrementor is exited.
func (s *Baseqasm3Listener) ExitIncrementor(ctx *IncrementorContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *Baseqasm3Listener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *Baseqasm3Listener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *Baseqasm3Listener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *Baseqasm3Listener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterMembershipTest is called when production membershipTest is entered.
func (s *Baseqasm3Listener) EnterMembershipTest(ctx *MembershipTestContext) {}

// ExitMembershipTest is called when production membershipTest is exited.
func (s *Baseqasm3Listener) ExitMembershipTest(ctx *MembershipTestContext) {}

// EnterSetDeclaration is called when production setDeclaration is entered.
func (s *Baseqasm3Listener) EnterSetDeclaration(ctx *SetDeclarationContext) {}

// ExitSetDeclaration is called when production setDeclaration is exited.
func (s *Baseqasm3Listener) ExitSetDeclaration(ctx *SetDeclarationContext) {}

// EnterLoopBranchBlock is called when production loopBranchBlock is entered.
func (s *Baseqasm3Listener) EnterLoopBranchBlock(ctx *LoopBranchBlockContext) {}

// ExitLoopBranchBlock is called when production loopBranchBlock is exited.
func (s *Baseqasm3Listener) ExitLoopBranchBlock(ctx *LoopBranchBlockContext) {}

// EnterBranchingStatement is called when production branchingStatement is entered.
func (s *Baseqasm3Listener) EnterBranchingStatement(ctx *BranchingStatementContext) {}

// ExitBranchingStatement is called when production branchingStatement is exited.
func (s *Baseqasm3Listener) ExitBranchingStatement(ctx *BranchingStatementContext) {}

// EnterLoopStatement is called when production loopStatement is entered.
func (s *Baseqasm3Listener) EnterLoopStatement(ctx *LoopStatementContext) {}

// ExitLoopStatement is called when production loopStatement is exited.
func (s *Baseqasm3Listener) ExitLoopStatement(ctx *LoopStatementContext) {}

// EnterControlDirectiveStatement is called when production controlDirectiveStatement is entered.
func (s *Baseqasm3Listener) EnterControlDirectiveStatement(ctx *ControlDirectiveStatementContext) {}

// ExitControlDirectiveStatement is called when production controlDirectiveStatement is exited.
func (s *Baseqasm3Listener) ExitControlDirectiveStatement(ctx *ControlDirectiveStatementContext) {}

// EnterControlDirective is called when production controlDirective is entered.
func (s *Baseqasm3Listener) EnterControlDirective(ctx *ControlDirectiveContext) {}

// ExitControlDirective is called when production controlDirective is exited.
func (s *Baseqasm3Listener) ExitControlDirective(ctx *ControlDirectiveContext) {}

// EnterKernelDeclaration is called when production kernelDeclaration is entered.
func (s *Baseqasm3Listener) EnterKernelDeclaration(ctx *KernelDeclarationContext) {}

// ExitKernelDeclaration is called when production kernelDeclaration is exited.
func (s *Baseqasm3Listener) ExitKernelDeclaration(ctx *KernelDeclarationContext) {}

// EnterSubroutineDefinition is called when production subroutineDefinition is entered.
func (s *Baseqasm3Listener) EnterSubroutineDefinition(ctx *SubroutineDefinitionContext) {}

// ExitSubroutineDefinition is called when production subroutineDefinition is exited.
func (s *Baseqasm3Listener) ExitSubroutineDefinition(ctx *SubroutineDefinitionContext) {}

// EnterSubroutineArgumentList is called when production subroutineArgumentList is entered.
func (s *Baseqasm3Listener) EnterSubroutineArgumentList(ctx *SubroutineArgumentListContext) {}

// ExitSubroutineArgumentList is called when production subroutineArgumentList is exited.
func (s *Baseqasm3Listener) ExitSubroutineArgumentList(ctx *SubroutineArgumentListContext) {}

// EnterPragma is called when production pragma is entered.
func (s *Baseqasm3Listener) EnterPragma(ctx *PragmaContext) {}

// ExitPragma is called when production pragma is exited.
func (s *Baseqasm3Listener) ExitPragma(ctx *PragmaContext) {}

// EnterTimeUnit is called when production timeUnit is entered.
func (s *Baseqasm3Listener) EnterTimeUnit(ctx *TimeUnitContext) {}

// ExitTimeUnit is called when production timeUnit is exited.
func (s *Baseqasm3Listener) ExitTimeUnit(ctx *TimeUnitContext) {}

// EnterTimingType is called when production timingType is entered.
func (s *Baseqasm3Listener) EnterTimingType(ctx *TimingTypeContext) {}

// ExitTimingType is called when production timingType is exited.
func (s *Baseqasm3Listener) ExitTimingType(ctx *TimingTypeContext) {}

// EnterTimingBox is called when production timingBox is entered.
func (s *Baseqasm3Listener) EnterTimingBox(ctx *TimingBoxContext) {}

// ExitTimingBox is called when production timingBox is exited.
func (s *Baseqasm3Listener) ExitTimingBox(ctx *TimingBoxContext) {}

// EnterTimeTerminator is called when production timeTerminator is entered.
func (s *Baseqasm3Listener) EnterTimeTerminator(ctx *TimeTerminatorContext) {}

// ExitTimeTerminator is called when production timeTerminator is exited.
func (s *Baseqasm3Listener) ExitTimeTerminator(ctx *TimeTerminatorContext) {}

// EnterTimeIdentifier is called when production timeIdentifier is entered.
func (s *Baseqasm3Listener) EnterTimeIdentifier(ctx *TimeIdentifierContext) {}

// ExitTimeIdentifier is called when production timeIdentifier is exited.
func (s *Baseqasm3Listener) ExitTimeIdentifier(ctx *TimeIdentifierContext) {}

// EnterTimeInstructionName is called when production timeInstructionName is entered.
func (s *Baseqasm3Listener) EnterTimeInstructionName(ctx *TimeInstructionNameContext) {}

// ExitTimeInstructionName is called when production timeInstructionName is exited.
func (s *Baseqasm3Listener) ExitTimeInstructionName(ctx *TimeInstructionNameContext) {}

// EnterTimeInstruction is called when production timeInstruction is entered.
func (s *Baseqasm3Listener) EnterTimeInstruction(ctx *TimeInstructionContext) {}

// ExitTimeInstruction is called when production timeInstruction is exited.
func (s *Baseqasm3Listener) ExitTimeInstruction(ctx *TimeInstructionContext) {}

// EnterTimeStatement is called when production timeStatement is entered.
func (s *Baseqasm3Listener) EnterTimeStatement(ctx *TimeStatementContext) {}

// ExitTimeStatement is called when production timeStatement is exited.
func (s *Baseqasm3Listener) ExitTimeStatement(ctx *TimeStatementContext) {}

// EnterCalibration is called when production calibration is entered.
func (s *Baseqasm3Listener) EnterCalibration(ctx *CalibrationContext) {}

// ExitCalibration is called when production calibration is exited.
func (s *Baseqasm3Listener) ExitCalibration(ctx *CalibrationContext) {}

// EnterCalibrationGrammarDeclaration is called when production calibrationGrammarDeclaration is entered.
func (s *Baseqasm3Listener) EnterCalibrationGrammarDeclaration(ctx *CalibrationGrammarDeclarationContext) {
}

// ExitCalibrationGrammarDeclaration is called when production calibrationGrammarDeclaration is exited.
func (s *Baseqasm3Listener) ExitCalibrationGrammarDeclaration(ctx *CalibrationGrammarDeclarationContext) {
}

// EnterCalibrationDefinition is called when production calibrationDefinition is entered.
func (s *Baseqasm3Listener) EnterCalibrationDefinition(ctx *CalibrationDefinitionContext) {}

// ExitCalibrationDefinition is called when production calibrationDefinition is exited.
func (s *Baseqasm3Listener) ExitCalibrationDefinition(ctx *CalibrationDefinitionContext) {}

// EnterCalibrationGrammar is called when production calibrationGrammar is entered.
func (s *Baseqasm3Listener) EnterCalibrationGrammar(ctx *CalibrationGrammarContext) {}

// ExitCalibrationGrammar is called when production calibrationGrammar is exited.
func (s *Baseqasm3Listener) ExitCalibrationGrammar(ctx *CalibrationGrammarContext) {}

// EnterCalibrationArgumentList is called when production calibrationArgumentList is entered.
func (s *Baseqasm3Listener) EnterCalibrationArgumentList(ctx *CalibrationArgumentListContext) {}

// ExitCalibrationArgumentList is called when production calibrationArgumentList is exited.
func (s *Baseqasm3Listener) ExitCalibrationArgumentList(ctx *CalibrationArgumentListContext) {}
