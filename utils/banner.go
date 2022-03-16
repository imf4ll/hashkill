package utils

import "fmt"

func Banner(mode, algo, filename, wordlist, file, hash string) {
    fmt.Printf(`=====================================================
[>] Mode -> %v
[>] Algorithm -> %v
[>] Filename -> %v
[>] Wordlist -> %v
[>] File -> %v
[>] Hash -> %v
[>] Version -> 0.1.2
[>] Repo -> https://github.com/z3oxs/hashkill
=====================================================

`, mode, algo, filename, wordlist, file, hash)
}
