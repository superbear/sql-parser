#### ANTLR Go
https://github.com/antlr/antlr4/blob/master/doc/go-target.md

#### Spark ANTLR4 Grammar
https://github.com/apache/spark/blob/master/sql/api/src/main/antlr4/org/apache/spark/sql/catalyst/parser/SqlBaseParser.g4

https://github.com/apache/spark/blob/master/sql/api/src/main/antlr4/org/apache/spark/sql/catalyst/parser/SqlBaseLexer.g4


superior-spark-parser/src/main/antlr4/io/github/melin/superior/parser/spark/antlr4/SparkSqlParser.g4

#### Gen Spark SQL Parser
```sh
antlr4 -Dlanguage=Go -visitor -package spark *.g4
```
without visitor
```sh
antlr4 -Dlanguage=Go -no-visitor -package spark *.g4
```
