### lulang

Just for fun!

#### TODO List

* [x] lexer
* [x] parser
* [x] evaluator
* [x] builtins
* [ ] bytecode
* [ ] compilation

#### Install

```shell
./install.sh
```

#### Run Interpreter

```shell
lulang interpreter
```

#### Run Compiler

```shell
lulang compile -s <source_file>
```

#### The heart of the parser

```go
func (p *Parser) parseExpression(precedence int) ast.Expression {
    prefix := p.prefixParseFns[p.curToken.Type]
    if prefix == nil {
        p.noPrefixParseFnError(p.curToken.Type)
        return nil
    }

    leftExp := prefix()
    // the heart of the parser
    for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
        infix := p.infixParseFns[p.peekToken.Type]
        if infix == nil {
            return leftExp
        }

        p.nextToken()
        leftExp = infix(leftExp)
    }

    return leftExp
}
```

#### Major Parts

1. string -> tokens -> ast -> objects (lexer -> parser -> evaluator)
2. string -> tokens -> ast -> bytecode -> objects (lexer -> parser -> compiler -> vm)

#### OpCode

1. OpConstant
2. OpAdd

#### Referring

[Top Down Operator Precedence](https://tdop.github.io/)
[Von Neumann architecture](https://en.wikipedia.org/wiki/Von_Neumann_architecture)

