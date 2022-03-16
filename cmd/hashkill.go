package cmd

import (
    "sync"
    "fmt"

    "github.com/z3oxs/hashkill/utils"
    "github.com/spf13/cobra"
)

var wg sync.WaitGroup

func Hashkill(c *cobra.Command, args []string) {
    algo, _ := c.Flags().GetString("algo")
    hashtype, _ := c.Flags().GetBool("hashtype")
    wordlist, _ := c.Flags().GetString("wordlist")
    version, _ := c.Flags().GetBool("version")
    hashedwl, _ := c.Flags().GetString("name")
    file, _ := c.Flags().GetString("file")
    hash, _ := c.Flags().GetString("hash")
    createM, _ := c.Flags().GetBool("create")
    crackM, _ := c.Flags().GetBool("crack")

    if hashtype {
        utils.Hashtype()

        return
    }

    if version {
        fmt.Println("v0.1.2 (03/16/2022)")

        return
    }

    if createM {
        utils.Banner("Create", algo, hashedwl, wordlist, file, hash)
        create(algo, wordlist, hashedwl)

    } else if crackM {
        utils.Banner("Crack", algo, "", wordlist, file, hash)
        
        if file != "" {
            crack(algo, wordlist, file, "file")

        } else {
            crack(algo, wordlist, hash, "hash")

        }
    }
}
