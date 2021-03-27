package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	pw "github.com/jerwheaton/SimplePasswordUtil/pkg/password"
)

const (
	bl            = "b"
	blLong        = "bloom"
	blDescription = "Use bloom filter"
)

var (
	useBloom    bool
	scalesInput string
	thumbsInput string

	rootCmd = &cobra.Command{
		Use:   "password",
		Short: "A utility that helps to create secure passwords",
		Long:  `PasswordService provides tools that can be used to create secure passwords that are safe to use on IoT services.`,
	}

	checkCmd = &cobra.Command{
		Use:   "check [list path] [password]",
		Short: "",
		Long:  ``,
		Args:  cobra.ExactArgs(2),
		Run:   RunCheck,
	}

	rateCmd = &cobra.Command{
		Use:   "rate [password]",
		Short: "",
		Long:  ``,
		Args:  cobra.ExactArgs(1),
		Run:   RunRate,
	}
)

func RunCheck(cmd *cobra.Command, args []string) {
	match, err := pw.Check(args[0], args[1], useBloom)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Matched: ", match)
}

func RunRate(cmd *cobra.Command, args []string) {
	r := pw.Rate(args[0])
	fmt.Println("Rated: ", r)
}

func NewCommand() *cobra.Command {
	checkCmd.Flags().BoolVarP(&useBloom, blLong, bl, false, blDescription)
	rootCmd.AddCommand(rateCmd)
	rootCmd.AddCommand(checkCmd)

	return rootCmd
}
