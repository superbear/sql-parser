#### ANTLR Go
https://github.com/antlr/antlr4/blob/master/doc/go-target.md

#### StarRocks ANTLR4 Grammar
https://github.com/StarRocks/starrocks/blob/main/fe/fe-core/src/main/java/com/starrocks/sql/parser/StarRocks.g4

https://github.com/StarRocks/starrocks/blob/main/fe/fe-core/src/main/java/com/starrocks/sql/parser/StarRocksLex.g4

#### Gen StarRocks SQL Parser
```sh
antlr4 -Dlanguage=Go -visitor -package starrocks *.g4
```
without visitor
```sh
antlr4 -Dlanguage=Go -no-visitor -package starrocks *.g4
```
