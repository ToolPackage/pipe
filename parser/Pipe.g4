grammar Pipe;

// script entry

script: funcDef* EOF;

funcDef: 'def' funcName funcParamsDef funcBody;

funcName: Identifier ('.' Identifier)*;

funcParamsDef: '(' (funcParamDef (',' funcParamDef)*)? ')';

funcParamDef: funcParamName optionalParamFlag? ':' funcParamType;

funcParamName: Identifier;

optionalParamFlag: '?';

funcParamType: ValueType;

funcBody: '{' (multiPipe PipeSeparator)? returnStatement '}';

returnStatement: 'return' pipe?;

// comand line entry

cmd: multiPipe EOF;

multiPipe: pipe (PipeSeparator pipe)*;

pipe: pipeNode (Connect pipeNode)*;

pipeNode: variableNode | streamNode | multiFunctionNode;

variableNode: '$' Identifier;

streamNode: '/' streamSplitter ',' pipe (',' streamCollector)? '/';

streamSplitter: functionNode;

streamCollector: functionNode;

multiFunctionNode: functionNode (',' functionNode)*;

functionNode: functionName functionParameters?;

functionName: Identifier ('.' Identifier)*;

functionParameters: '(' (functionParameter (',' functionParameter)* ','?)? ')';

functionParameter: functionParameterLabel? functionParameterValue;

functionParameterLabel: Identifier ':';

functionParameterValue: variableValue | dictValue | numberValue | stringValue | booleanValue;

variableValue: '$' Identifier;

Connect: [=];

PipeSeparator: [;&];

ValueType: 'string' | 'integer' | 'float' | 'bool' | 'dict';

// dict

dictValue: '[' dictEntries? ']';

dictEntries: dictEntry (',' dictEntry)*;

dictEntry: (dictEntryLabel ':')? dictEntryValue;

dictEntryLabel: Identifier;

dictEntryValue: numberValue | decimalValue | stringValue | booleanValue | variableValue;

// number literal

numberValue: integerValue | decimalValue;

integerValue: IntegerLiteral;

decimalValue: DecimalLiteral;

IntegerLiteral: '0' | '-'? NonZeroDigit Digits?;

DecimalLiteral: '-'? Digits '.' Digits? | '.' Digits;

fragment Digits: Digit+;

fragment Digit: '0' | NonZeroDigit;

fragment NonZeroDigit: [1-9];

// string literal

stringValue: StringLiteral;

StringLiteral: '\'' StringCharacters? '\'';

fragment StringCharacters: StringCharacter+;

fragment StringCharacter: ~['];

// boolean literal

booleanValue: BooleanLiteral;

BooleanLiteral: True | False;

True: 'true';

False: 'false';

// identifier

Identifier: Letter LetterOrDigit*;

fragment Letter: [a-zA-Z];

fragment LetterOrDigit: [a-zA-Z0-9];

// whitespace

WS: [ \t\r\n\u000C]+ -> skip;