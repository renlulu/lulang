### lulang

Just for fun!

#### TODO List

* [x] lexer
* [ ] parser
* [ ] evaluation
* [ ] compilation
* [ ] use `chan` to improve performance

#### Project layout

##### ast

Definition of Abstract Syntax Tree.

##### lexer

Lexer that scan input string to produce `Token`.

##### Parser

Parser consumes `Token` and produces AST.

##### repl

A simple interactive environment of `lulang`.

##### token

Basic unit of `lulang`.

#### Some notes (terminology) about the parser

*prefix operator* is an operator in the front of its operand:

--5

*postfix operator* is an operator after its operand:

foobar++

*infix operators* appear in binary expressions:

5 * 8

*order of operations* means which priority do different operators have:

5 + 5 * 10

