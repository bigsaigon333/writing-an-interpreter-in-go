package lexer

type Lexer struct {
	input        string // 입력 문자열
	position     int    // 입력 문자열에서 현재 문자의 위치
	readPosition int    // 입력 문자열에서 다음 문자의 위치
	ch           byte   // 현재 문자
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
