// cmd/cryptosecuritytoolkittechpro/main.go
package main

import (
"flag"
"log"
"os"

"cryptosecuritytoolkittechpro/internal/cryptosecuritytoolkittechpro"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := cryptosecuritytoolkittechpro.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
