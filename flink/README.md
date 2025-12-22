#### ANTLR Go
https://github.com/antlr/antlr4/blob/master/doc/go-target.md

#### Flink ANTLR4 Grammar
https://github.com/DTStack/dt-sql-parser/tree/main/src/grammar/flink

#### Gen StarRocks SQL Parser
```sh
antlr4 -Dlanguage=Go -visitor -package flink *.g4
```
without visitor
```sh
antlr4 -Dlanguage=Go -no-visitor -package flink *.g4
```
