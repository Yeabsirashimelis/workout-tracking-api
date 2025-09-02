package migrations

import "embed"

/*
This imports Go’s built-in embed package.
embed lets you embed files into your compiled binary, so you don’t need to read them from disk at runtime.

// go:embed *.sql
   This is a special comment that Go understands at compile time.
   *.sql means include all .sql files in the same folder.
   These files will be embedded into the binary and accessible through the variable declared below it.

*/

//go:embed *.sql
var FS embed.FS