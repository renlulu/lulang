### lulang

Just for fun!

#### TODO List

* [x] lexer
* [x] parser
* [x] evaluator
* [x] builtins
* [ ] bytecode
* [ ] compilation

#### The heart of the parser

```go
for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
    infix := p.infixParseFns[p.peekToken.Type]
    if infix == nil {
        return leftExp
    }

    p.nextToken()
    leftExp = infix(leftExp)
}
```

#### Major Parts

1. string -> tokens -> ast -> objects (lexer -> parser -> evaluator)
2. string -> tokens -> ast -> bytecode -> objects (lexer -> parser -> compiler -> vm)

#### Referring

[Top Down Operator Precedence](https://tdop.github.io/)
[Von Neumann architecture](https://en.wikipedia.org/wiki/Von_Neumann_architecture)

