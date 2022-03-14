package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command {
    Short: "Simple CLI to create hashlists and fast crack any hash with their.",
    Run: Hashkill,
}

func init() {
    rootCmd.PersistentFlags().StringP("wordlist", "w", "", "Wordlist")
    rootCmd.PersistentFlags().BoolP("hashtype", "t", false, "Print available, hashtypes")
    rootCmd.PersistentFlags().StringP("file", "f", "", "Hashes file")
    rootCmd.PersistentFlags().StringP("hash", "H", "", "Hash string")
    rootCmd.PersistentFlags().StringP("name", "n", "hashlist.txt", "Hash wordlist filename")
    rootCmd.PersistentFlags().BoolP("create", "c", false, "Create hashed wordlist")
    rootCmd.PersistentFlags().BoolP("crack", "C", false, "Break hashs mode")
    rootCmd.PersistentFlags().StringP("algo", "a", "md5", "Hash algorithm")
}

func Execute() {
    rootCmd.Execute()
}
