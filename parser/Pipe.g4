grammar Pipe;

script: function (Connect function)*;

function: functionName functionParameters?;

functionName: Identifier ('.' Identifier)*;

functionParameters: '(' (functionParameter (',' functionParameter)* ','?)? ')';

functionParameter: functionParameterLabel? (listValue | mapValue | numberValue | stringValue | booleanValue);

functionParameterLabel: Identifier ':';

Connect: [=];

// list

listValue: '[' listElements? ']';

listElements: (numberValue | stringValue | booleanValue) (',' (numberValue | stringValue | booleanValue));

// map

mapValue: '[' mapEntries? ']';

mapEntries: mapEntry (',' mapEntry)*;

mapEntry: Identifier ':' (numberValue | stringValue | booleanValue);

// number literal

numberValue: NumberLiteral;

NumberLiteral: IntegerLiteral | FloatingPointLiteral;

fragment IntegerLiteral: '0' | '-'? NonZeroDigit Digits?;

fragment FloatingPointLiteral: '-'? Digits '.' Digits? | '.' Digits;

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