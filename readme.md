### lulang

Just for fun!

#### TODO List

* [x] lexer
* [x] parser
* [x] evaluator
* [ ] builtins
* [ ] compilation
* [ ] use `chan` to improve performance

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

#### Referring

[Top Down Operator Precedence](https://tdop.github.io/)

