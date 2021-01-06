// Code generated from ./qasm3.g4 by ANTLR 4.9. DO NOT EDIT.

package qasm3 // qasm3
import "github.com/antlr/antlr4/runtime/Go/antlr"

// qasm3Listener is a complete listener for a parse tree produced by qasm3Parser.
type qasm3Listener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterHeader is called when entering the header production.
	EnterHeader(c *HeaderContext)

	// EnterVersion is called when entering the version production.
	EnterVersion(c *VersionContext)

	// EnterInclude is called when entering the include production.
	EnterInclude(c *IncludeContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterGlobalStatement is called when entering the globalStatement production.
	EnterGlobalStatement(c *GlobalStatementContext)

	// EnterDeclarationStatement is called when entering the declarationStatement production.
	EnterDeclarationStatement(c *DeclarationStatementContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterReturnSignature is called when entering the returnSignature production.
	EnterReturnSignature(c *ReturnSignatureContext)

	// EnterProgramBlock is called when entering the programBlock production.
	EnterProgramBlock(c *ProgramBlockContext)

	// EnterDesignator is called when entering the designator production.
	EnterDesignator(c *DesignatorContext)

	// EnterDoubleDesignator is called when entering the doubleDesignator production.
	EnterDoubleDesignator(c *DoubleDesignatorContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterIndexIdentifier is called when entering the indexIdentifier production.
	EnterIndexIdentifier(c *IndexIdentifierContext)

	// EnterIndexIdentifierList is called when entering the indexIdentifierList production.
	EnterIndexIdentifierList(c *IndexIdentifierListContext)

	// EnterAssociation is called when entering the association production.
	EnterAssociation(c *AssociationContext)

	// EnterQuantumType is called when entering the quantumType production.
	EnterQuantumType(c *QuantumTypeContext)

	// EnterQuantumDeclaration is called when entering the quantumDeclaration production.
	EnterQuantumDeclaration(c *QuantumDeclarationContext)

	// EnterQuantumArgument is called when entering the quantumArgument production.
	EnterQuantumArgument(c *QuantumArgumentContext)

	// EnterQuantumArgumentList is called when entering the quantumArgumentList production.
	EnterQuantumArgumentList(c *QuantumArgumentListContext)

	// EnterBitType is called when entering the bitType production.
	EnterBitType(c *BitTypeContext)

	// EnterSingleDesignatorType is called when entering the singleDesignatorType production.
	EnterSingleDesignatorType(c *SingleDesignatorTypeContext)

	// EnterDoubleDesignatorType is called when entering the doubleDesignatorType production.
	EnterDoubleDesignatorType(c *DoubleDesignatorTypeContext)

	// EnterNoDesignatorType is called when entering the noDesignatorType production.
	EnterNoDesignatorType(c *NoDesignatorTypeContext)

	// EnterClassicalType is called when entering the classicalType production.
	EnterClassicalType(c *ClassicalTypeContext)

	// EnterConstantDeclaration is called when entering the constantDeclaration production.
	EnterConstantDeclaration(c *ConstantDeclarationContext)

	// EnterSingleDesignatorDeclaration is called when entering the singleDesignatorDeclaration production.
	EnterSingleDesignatorDeclaration(c *SingleDesignatorDeclarationContext)

	// EnterDoubleDesignatorDeclaration is called when entering the doubleDesignatorDeclaration production.
	EnterDoubleDesignatorDeclaration(c *DoubleDesignatorDeclarationContext)

	// EnterNoDesignatorDeclaration is called when entering the noDesignatorDeclaration production.
	EnterNoDesignatorDeclaration(c *NoDesignatorDeclarationContext)

	// EnterBitDeclaration is called when entering the bitDeclaration production.
	EnterBitDeclaration(c *BitDeclarationContext)

	// EnterClassicalVariableDeclaration is called when entering the classicalVariableDeclaration production.
	EnterClassicalVariableDeclaration(c *ClassicalVariableDeclarationContext)

	// EnterClassicalDeclaration is called when entering the classicalDeclaration production.
	EnterClassicalDeclaration(c *ClassicalDeclarationContext)

	// EnterClassicalTypeList is called when entering the classicalTypeList production.
	EnterClassicalTypeList(c *ClassicalTypeListContext)

	// EnterClassicalArgument is called when entering the classicalArgument production.
	EnterClassicalArgument(c *ClassicalArgumentContext)

	// EnterClassicalArgumentList is called when entering the classicalArgumentList production.
	EnterClassicalArgumentList(c *ClassicalArgumentListContext)

	// EnterAliasStatement is called when entering the aliasStatement production.
	EnterAliasStatement(c *AliasStatementContext)

	// EnterConcatenateExpression is called when entering the concatenateExpression production.
	EnterConcatenateExpression(c *ConcatenateExpressionContext)

	// EnterRangeDefinition is called when entering the rangeDefinition production.
	EnterRangeDefinition(c *RangeDefinitionContext)

	// EnterQuantumGateDefinition is called when entering the quantumGateDefinition production.
	EnterQuantumGateDefinition(c *QuantumGateDefinitionContext)

	// EnterQuantumGateSignature is called when entering the quantumGateSignature production.
	EnterQuantumGateSignature(c *QuantumGateSignatureContext)

	// EnterQuantumBlock is called when entering the quantumBlock production.
	EnterQuantumBlock(c *QuantumBlockContext)

	// EnterQuantumStatement is called when entering the quantumStatement production.
	EnterQuantumStatement(c *QuantumStatementContext)

	// EnterQuantumInstruction is called when entering the quantumInstruction production.
	EnterQuantumInstruction(c *QuantumInstructionContext)

	// EnterQuantumMeasurement is called when entering the quantumMeasurement production.
	EnterQuantumMeasurement(c *QuantumMeasurementContext)

	// EnterQuantumMeasurementDeclaration is called when entering the quantumMeasurementDeclaration production.
	EnterQuantumMeasurementDeclaration(c *QuantumMeasurementDeclarationContext)

	// EnterQuantumBarrier is called when entering the quantumBarrier production.
	EnterQuantumBarrier(c *QuantumBarrierContext)

	// EnterQuantumGateModifier is called when entering the quantumGateModifier production.
	EnterQuantumGateModifier(c *QuantumGateModifierContext)

	// EnterQuantumGateCall is called when entering the quantumGateCall production.
	EnterQuantumGateCall(c *QuantumGateCallContext)

	// EnterQuantumGateName is called when entering the quantumGateName production.
	EnterQuantumGateName(c *QuantumGateNameContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// EnterBinaryOperator is called when entering the binaryOperator production.
	EnterBinaryOperator(c *BinaryOperatorContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpressionTerminator is called when entering the expressionTerminator production.
	EnterExpressionTerminator(c *ExpressionTerminatorContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterBuiltInMath is called when entering the builtInMath production.
	EnterBuiltInMath(c *BuiltInMathContext)

	// EnterCastOperator is called when entering the castOperator production.
	EnterCastOperator(c *CastOperatorContext)

	// EnterIncrementor is called when entering the incrementor production.
	EnterIncrementor(c *IncrementorContext)

	// EnterAssignmentExpression is called when entering the assignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterMembershipTest is called when entering the membershipTest production.
	EnterMembershipTest(c *MembershipTestContext)

	// EnterSetDeclaration is called when entering the setDeclaration production.
	EnterSetDeclaration(c *SetDeclarationContext)

	// EnterLoopBranchBlock is called when entering the loopBranchBlock production.
	EnterLoopBranchBlock(c *LoopBranchBlockContext)

	// EnterBranchingStatement is called when entering the branchingStatement production.
	EnterBranchingStatement(c *BranchingStatementContext)

	// EnterLoopStatement is called when entering the loopStatement production.
	EnterLoopStatement(c *LoopStatementContext)

	// EnterControlDirectiveStatement is called when entering the controlDirectiveStatement production.
	EnterControlDirectiveStatement(c *ControlDirectiveStatementContext)

	// EnterControlDirective is called when entering the controlDirective production.
	EnterControlDirective(c *ControlDirectiveContext)

	// EnterKernelDeclaration is called when entering the kernelDeclaration production.
	EnterKernelDeclaration(c *KernelDeclarationContext)

	// EnterSubroutineDefinition is called when entering the subroutineDefinition production.
	EnterSubroutineDefinition(c *SubroutineDefinitionContext)

	// EnterSubroutineArgumentList is called when entering the subroutineArgumentList production.
	EnterSubroutineArgumentList(c *SubroutineArgumentListContext)

	// EnterPragma is called when entering the pragma production.
	EnterPragma(c *PragmaContext)

	// EnterTimeUnit is called when entering the timeUnit production.
	EnterTimeUnit(c *TimeUnitContext)

	// EnterTimingType is called when entering the timingType production.
	EnterTimingType(c *TimingTypeContext)

	// EnterTimingBox is called when entering the timingBox production.
	EnterTimingBox(c *TimingBoxContext)

	// EnterTimeTerminator is called when entering the timeTerminator production.
	EnterTimeTerminator(c *TimeTerminatorContext)

	// EnterTimeIdentifier is called when entering the timeIdentifier production.
	EnterTimeIdentifier(c *TimeIdentifierContext)

	// EnterTimeInstructionName is called when entering the timeInstructionName production.
	EnterTimeInstructionName(c *TimeInstructionNameContext)

	// EnterTimeInstruction is called when entering the timeInstruction production.
	EnterTimeInstruction(c *TimeInstructionContext)

	// EnterTimeStatement is called when entering the timeStatement production.
	EnterTimeStatement(c *TimeStatementContext)

	// EnterCalibration is called when entering the calibration production.
	EnterCalibration(c *CalibrationContext)

	// EnterCalibrationGrammarDeclaration is called when entering the calibrationGrammarDeclaration production.
	EnterCalibrationGrammarDeclaration(c *CalibrationGrammarDeclarationContext)

	// EnterCalibrationDefinition is called when entering the calibrationDefinition production.
	EnterCalibrationDefinition(c *CalibrationDefinitionContext)

	// EnterCalibrationGrammar is called when entering the calibrationGrammar production.
	EnterCalibrationGrammar(c *CalibrationGrammarContext)

	// EnterCalibrationArgumentList is called when entering the calibrationArgumentList production.
	EnterCalibrationArgumentList(c *CalibrationArgumentListContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitHeader is called when exiting the header production.
	ExitHeader(c *HeaderContext)

	// ExitVersion is called when exiting the version production.
	ExitVersion(c *VersionContext)

	// ExitInclude is called when exiting the include production.
	ExitInclude(c *IncludeContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitGlobalStatement is called when exiting the globalStatement production.
	ExitGlobalStatement(c *GlobalStatementContext)

	// ExitDeclarationStatement is called when exiting the declarationStatement production.
	ExitDeclarationStatement(c *DeclarationStatementContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitReturnSignature is called when exiting the returnSignature production.
	ExitReturnSignature(c *ReturnSignatureContext)

	// ExitProgramBlock is called when exiting the programBlock production.
	ExitProgramBlock(c *ProgramBlockContext)

	// ExitDesignator is called when exiting the designator production.
	ExitDesignator(c *DesignatorContext)

	// ExitDoubleDesignator is called when exiting the doubleDesignator production.
	ExitDoubleDesignator(c *DoubleDesignatorContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitIndexIdentifier is called when exiting the indexIdentifier production.
	ExitIndexIdentifier(c *IndexIdentifierContext)

	// ExitIndexIdentifierList is called when exiting the indexIdentifierList production.
	ExitIndexIdentifierList(c *IndexIdentifierListContext)

	// ExitAssociation is called when exiting the association production.
	ExitAssociation(c *AssociationContext)

	// ExitQuantumType is called when exiting the quantumType production.
	ExitQuantumType(c *QuantumTypeContext)

	// ExitQuantumDeclaration is called when exiting the quantumDeclaration production.
	ExitQuantumDeclaration(c *QuantumDeclarationContext)

	// ExitQuantumArgument is called when exiting the quantumArgument production.
	ExitQuantumArgument(c *QuantumArgumentContext)

	// ExitQuantumArgumentList is called when exiting the quantumArgumentList production.
	ExitQuantumArgumentList(c *QuantumArgumentListContext)

	// ExitBitType is called when exiting the bitType production.
	ExitBitType(c *BitTypeContext)

	// ExitSingleDesignatorType is called when exiting the singleDesignatorType production.
	ExitSingleDesignatorType(c *SingleDesignatorTypeContext)

	// ExitDoubleDesignatorType is called when exiting the doubleDesignatorType production.
	ExitDoubleDesignatorType(c *DoubleDesignatorTypeContext)

	// ExitNoDesignatorType is called when exiting the noDesignatorType production.
	ExitNoDesignatorType(c *NoDesignatorTypeContext)

	// ExitClassicalType is called when exiting the classicalType production.
	ExitClassicalType(c *ClassicalTypeContext)

	// ExitConstantDeclaration is called when exiting the constantDeclaration production.
	ExitConstantDeclaration(c *ConstantDeclarationContext)

	// ExitSingleDesignatorDeclaration is called when exiting the singleDesignatorDeclaration production.
	ExitSingleDesignatorDeclaration(c *SingleDesignatorDeclarationContext)

	// ExitDoubleDesignatorDeclaration is called when exiting the doubleDesignatorDeclaration production.
	ExitDoubleDesignatorDeclaration(c *DoubleDesignatorDeclarationContext)

	// ExitNoDesignatorDeclaration is called when exiting the noDesignatorDeclaration production.
	ExitNoDesignatorDeclaration(c *NoDesignatorDeclarationContext)

	// ExitBitDeclaration is called when exiting the bitDeclaration production.
	ExitBitDeclaration(c *BitDeclarationContext)

	// ExitClassicalVariableDeclaration is called when exiting the classicalVariableDeclaration production.
	ExitClassicalVariableDeclaration(c *ClassicalVariableDeclarationContext)

	// ExitClassicalDeclaration is called when exiting the classicalDeclaration production.
	ExitClassicalDeclaration(c *ClassicalDeclarationContext)

	// ExitClassicalTypeList is called when exiting the classicalTypeList production.
	ExitClassicalTypeList(c *ClassicalTypeListContext)

	// ExitClassicalArgument is called when exiting the classicalArgument production.
	ExitClassicalArgument(c *ClassicalArgumentContext)

	// ExitClassicalArgumentList is called when exiting the classicalArgumentList production.
	ExitClassicalArgumentList(c *ClassicalArgumentListContext)

	// ExitAliasStatement is called when exiting the aliasStatement production.
	ExitAliasStatement(c *AliasStatementContext)

	// ExitConcatenateExpression is called when exiting the concatenateExpression production.
	ExitConcatenateExpression(c *ConcatenateExpressionContext)

	// ExitRangeDefinition is called when exiting the rangeDefinition production.
	ExitRangeDefinition(c *RangeDefinitionContext)

	// ExitQuantumGateDefinition is called when exiting the quantumGateDefinition production.
	ExitQuantumGateDefinition(c *QuantumGateDefinitionContext)

	// ExitQuantumGateSignature is called when exiting the quantumGateSignature production.
	ExitQuantumGateSignature(c *QuantumGateSignatureContext)

	// ExitQuantumBlock is called when exiting the quantumBlock production.
	ExitQuantumBlock(c *QuantumBlockContext)

	// ExitQuantumStatement is called when exiting the quantumStatement production.
	ExitQuantumStatement(c *QuantumStatementContext)

	// ExitQuantumInstruction is called when exiting the quantumInstruction production.
	ExitQuantumInstruction(c *QuantumInstructionContext)

	// ExitQuantumMeasurement is called when exiting the quantumMeasurement production.
	ExitQuantumMeasurement(c *QuantumMeasurementContext)

	// ExitQuantumMeasurementDeclaration is called when exiting the quantumMeasurementDeclaration production.
	ExitQuantumMeasurementDeclaration(c *QuantumMeasurementDeclarationContext)

	// ExitQuantumBarrier is called when exiting the quantumBarrier production.
	ExitQuantumBarrier(c *QuantumBarrierContext)

	// ExitQuantumGateModifier is called when exiting the quantumGateModifier production.
	ExitQuantumGateModifier(c *QuantumGateModifierContext)

	// ExitQuantumGateCall is called when exiting the quantumGateCall production.
	ExitQuantumGateCall(c *QuantumGateCallContext)

	// ExitQuantumGateName is called when exiting the quantumGateName production.
	ExitQuantumGateName(c *QuantumGateNameContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitBinaryOperator is called when exiting the binaryOperator production.
	ExitBinaryOperator(c *BinaryOperatorContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpressionTerminator is called when exiting the expressionTerminator production.
	ExitExpressionTerminator(c *ExpressionTerminatorContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitBuiltInMath is called when exiting the builtInMath production.
	ExitBuiltInMath(c *BuiltInMathContext)

	// ExitCastOperator is called when exiting the castOperator production.
	ExitCastOperator(c *CastOperatorContext)

	// ExitIncrementor is called when exiting the incrementor production.
	ExitIncrementor(c *IncrementorContext)

	// ExitAssignmentExpression is called when exiting the assignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitMembershipTest is called when exiting the membershipTest production.
	ExitMembershipTest(c *MembershipTestContext)

	// ExitSetDeclaration is called when exiting the setDeclaration production.
	ExitSetDeclaration(c *SetDeclarationContext)

	// ExitLoopBranchBlock is called when exiting the loopBranchBlock production.
	ExitLoopBranchBlock(c *LoopBranchBlockContext)

	// ExitBranchingStatement is called when exiting the branchingStatement production.
	ExitBranchingStatement(c *BranchingStatementContext)

	// ExitLoopStatement is called when exiting the loopStatement production.
	ExitLoopStatement(c *LoopStatementContext)

	// ExitControlDirectiveStatement is called when exiting the controlDirectiveStatement production.
	ExitControlDirectiveStatement(c *ControlDirectiveStatementContext)

	// ExitControlDirective is called when exiting the controlDirective production.
	ExitControlDirective(c *ControlDirectiveContext)

	// ExitKernelDeclaration is called when exiting the kernelDeclaration production.
	ExitKernelDeclaration(c *KernelDeclarationContext)

	// ExitSubroutineDefinition is called when exiting the subroutineDefinition production.
	ExitSubroutineDefinition(c *SubroutineDefinitionContext)

	// ExitSubroutineArgumentList is called when exiting the subroutineArgumentList production.
	ExitSubroutineArgumentList(c *SubroutineArgumentListContext)

	// ExitPragma is called when exiting the pragma production.
	ExitPragma(c *PragmaContext)

	// ExitTimeUnit is called when exiting the timeUnit production.
	ExitTimeUnit(c *TimeUnitContext)

	// ExitTimingType is called when exiting the timingType production.
	ExitTimingType(c *TimingTypeContext)

	// ExitTimingBox is called when exiting the timingBox production.
	ExitTimingBox(c *TimingBoxContext)

	// ExitTimeTerminator is called when exiting the timeTerminator production.
	ExitTimeTerminator(c *TimeTerminatorContext)

	// ExitTimeIdentifier is called when exiting the timeIdentifier production.
	ExitTimeIdentifier(c *TimeIdentifierContext)

	// ExitTimeInstructionName is called when exiting the timeInstructionName production.
	ExitTimeInstructionName(c *TimeInstructionNameContext)

	// ExitTimeInstruction is called when exiting the timeInstruction production.
	ExitTimeInstruction(c *TimeInstructionContext)

	// ExitTimeStatement is called when exiting the timeStatement production.
	ExitTimeStatement(c *TimeStatementContext)

	// ExitCalibration is called when exiting the calibration production.
	ExitCalibration(c *CalibrationContext)

	// ExitCalibrationGrammarDeclaration is called when exiting the calibrationGrammarDeclaration production.
	ExitCalibrationGrammarDeclaration(c *CalibrationGrammarDeclarationContext)

	// ExitCalibrationDefinition is called when exiting the calibrationDefinition production.
	ExitCalibrationDefinition(c *CalibrationDefinitionContext)

	// ExitCalibrationGrammar is called when exiting the calibrationGrammar production.
	ExitCalibrationGrammar(c *CalibrationGrammarContext)

	// ExitCalibrationArgumentList is called when exiting the calibrationArgumentList production.
	ExitCalibrationArgumentList(c *CalibrationArgumentListContext)
}
