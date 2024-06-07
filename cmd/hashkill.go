package cmd

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/spf13/cobra"
	"github.com/z3oxs/hashkill/utils"
)

var wg sync.WaitGroup
var l = log.New(os.Stdout, "", 0)

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
        fmt.Println("\033[1;32m[🗸]\033[m v0.1.4 (03/31/2022)")

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
