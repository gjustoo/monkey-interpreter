# TODO


- Clean up the code.
- Better folder structure.


## Known issues

- The lexer does not support UTF-8/unicode.
- The repl freezes when the user enters an expression without a semicolon at the end. --> Important!

## Lexer

- readIdent -> Implement UTF-8/unicode support.
- readDigit -> Implement floats, hex, octal support.
- nextToken -> Extract method for the twoCharToken case.
- Add support for comments.
- Add line and column information to the token for better error messages.

## AST

## Parser


## Evaluator

- Better error handling