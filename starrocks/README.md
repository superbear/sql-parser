#### ANTLR Go
https://github.com/antlr/antlr4/blob/master/doc/go-target.md

#### StarRocks ANTLR4 Grammar
https://github.com/StarRocks/starrocks/blob/main/fe/fe-grammar/src/main/antlr/com/starrocks/grammar/StarRocks.g4

https://github.com/StarRocks/starrocks/blob/main/fe/fe-grammar/src/main/antlr/com/starrocks/grammar/StarRocksLex.g4

#### Gen StarRocks SQL Parser
```sh
antlr4 -Dlanguage=Go -visitor -package starrocks *.g4
```
without visitor
```sh
antlr4 -Dlanguage=Go -no-visitor -package starrocks *.g4
```
