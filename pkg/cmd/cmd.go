package cmd

import (
	"fmt"
	"log"

	pw "github.com/jerwheaton/SimplePasswordUtil/pkg/password"
	"github.com/spf13/cobra"
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
		Use:   "check [password] [optional: list path]",
		Short: "",
		Long:  ``,
		Args:  cobra.RangeArgs(1, 2),
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
	path := ""
	if len(args) == 2 {
		path = args[1]
	}
	match, err := pw.Check(path, args[0], useBloom)
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
