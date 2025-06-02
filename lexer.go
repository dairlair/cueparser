// // The `Lexer` provides tools for parsing CUE sheet lines.
// //
// // A CUE sheet is a plain-text metadata format used to describe the layout of tracks
// // in a CD or disk image. Common directives include FILE, TRACK, TITLE, INDEX, etc.
// // Each directive is expected to appear on a separate line in the file.
// //
// // The lexer implemented in this file (lexer.go), specifically the `Lexer` struct
// // returned by the `NewLexer` function, processes one line at a time and tokenizes it.
// // The input line must not contain newline characters (\n or \r\n).
// //
// // According to the CUE sheet format specification:
// // https://github.com/libyal/libodraw/blob/main/documentation/CUE%20sheet%20format.asciidoc
// //
// // > Each line in the cue sheet file defines a command, such as FILE, TRACK, INDEX, TITLE, PERFORMER, etc.
// //
// // While the specification does not explicitly prohibit multiline commands or line continuations,
// // all examples and structural explanations treat each line as an atomic command. Therefore, this lexer
// // assumes a one-command-per-line structure.
// //
// // The lexer supports:
// //
// //   - Uppercase keywords (e.g., FILE, TRACK, TITLE)
// //   - Quoted strings with support for spaces (e.g., "My Song.flac")
// //   - Simple argument tokenization for numeric or unquoted parameters
// //   - Comments using the following formats:
// //       - `;`, e.g., `; Just commenting`
// //       - `//`, e.g., `// Just commenting - such style is used by `cdrdao`
// //
// //     Comments do not have to be on a line by themselves — inline comments following
// //     valid tokens are also recognized and handled appropriately.
// //
// // This Lexer is intended to be used as a low-level component for building
// // full parsers that operate on entire .cue files.
// //
// // Lex takes a single line from a CUE sheet and returns a slice of tokens.
// // The input string must not contain a trailing newline.
package cueparser

//
//import (
//	"bufio"
//	"io"
//)
//
//// Command represents a single parsed command line from a CUE sheet.
//type Command struct {
//	Keyword string   // The command keyword (e.g., FILE, TRACK, TITLE)
//	Args    []string // Arguments for the keyword, in order
//	Comment string   // Optional comment (from ; or //), without the comment prefix
//	Raw     string   // The original raw line, if needed for reference/debugging
//}
//
//type Lexer struct {
//	input bufio.Reader
//}
//
//func NewLexer(r io.Reader) *Lexer {
//	input := bufio.NewReader(r)
//	return &Lexer{
//		input: *input,
//	}
//}
//
//func (l *Lexer) Lex() Command {
//	return Command{}
//}
