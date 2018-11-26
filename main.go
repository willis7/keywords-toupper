package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var sqlkeywords = []string{
"all", "alter", "and", "any", "array", "arrow", "as", "asc", "at",
"begin", "between", "by",
"case", "check", "clusters", "cluster", "colauth", "columns", "compress", "connect", "crash", "create", "current",
"decimal", "declare", "default", "delete", "desc", "distinct", "drop",
"else", "end", "exception", "exclusive", "exists",
"fetch", "form", "for", "from",
"goto", "grant", "group",
"having",
"identified", "if", "in", "indexes", "index", "insert", "intersect", "into", "is",
"like", "lock",
"minus", "mode",
"nocompress", "not", "nowait", "null",
"of", "on", "option", "or", "order", "overlaps",
"prior", "procedure", "public",
"range", "record", "resource", "revoke",
"select", "share", "size", "sql", "start", "subtype",
"tabauth", "table", "then", "to", "type",
"union", "unique", "update", "use",
"values", "view", "views",
"when", "where", "with",
}

var re = regexp.MustCompile(`\b(` + strings.Join(sqlkeywords, "|") + `)\b`)

func KeywordsToUpper(src string) string {
	return re.ReplaceAllStringFunc(src, func(w string) string {
		return strings.ToUpper(w)
	})
}

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	for err == nil {
		fmt.Fprint(os.Stdout, KeywordsToUpper(line))
		line, err = r.ReadString('\n')
	}
	if err != io.EOF {
		fmt.Println(err)
		return
	}
}

