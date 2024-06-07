package utils

import "fmt"

func Banner(mode, algo, filename, wordlist, file, hash string) {
    banner := "=====================================================\n"
    if mode != "" { banner += fmt.Sprintf("➤ Mode: %v\n", mode) }
    if algo != "" { banner += fmt.Sprintf("➤ Algorithm: %v\n", algo) }
    if filename != "" { banner += fmt.Sprintf("➤ Filename: %v\n", filename) }
    if wordlist != "" { banner += fmt.Sprintf("➤ Wordlist: %v\n", wordlist) }
    if file != "" { banner += fmt.Sprintf("➤ File: %v\n", file) }
    if hash != "" { banner += fmt.Sprintf("➤ Hash: %v\n", hash) }
    banner += "=====================================================\n"

    fmt.Println(banner)
}
