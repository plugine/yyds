grammar Music;

// Parser
music: content EOF;
content: expression | content expression;
expression: controlAssign | pattern;
controlAssign: SET key = TOKEN '=' value = TOKEN;
pattern:
	PATTERN L_CURLY list = noteList R_CURLY
	| PATTERN LEFT_RBRACKET argsContext = args RIGHT_RBRACKET L_CURLY list = noteList R_CURLY;
args:
	assignContext = assign
	| argsContext = args COMMA assignContext = assign;
assign: key = TOKEN '=' value = TOKEN;
noteList: noteLine+;
noteLine: signContext = TOKEN COLON notesContext = notes;

notes: noteWithSign+;

noteWithSign:
	tokenContext = TOKEN
	| tokenContext = TOKEN MULTIPLE multipleContext = TOKEN
	| tokenContext = TOKEN L_CURLY noteSignContext = TOKEN R_CURLY
	| tokenContext = TOKEN L_CURLY noteSignContext = TOKEN R_CURLY MULTIPLE multipleContext =
		TOKEN;

// Lexer

LEFT_RBRACKET: '(';
RIGHT_RBRACKET: ')';
MULTIPLE: '*';
PATTERN: 'pattern';
SET: 'set';
TOKEN: [a-zA-Z0-9_+]+;
L_CURLY: '{';
R_CURLY: '}';
COMMA: ',';
COLON: ':';
WHITESPACE: [ \r\n\t]+ -> skip;
LINE_COMMENT: '//' ~[\r\n]* -> channel(HIDDEN);