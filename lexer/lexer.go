package lexer

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading  position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// returns next character and advance pos in input string
func (l *Lexer) readChar() {
	// check if end of input reached
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL", haven't read anything or EOF
	} else { // haven't reached end, so l.ch to next character
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition // points to pos we last read
	l.readPosition += 1         // always points to position which will be read next
}
