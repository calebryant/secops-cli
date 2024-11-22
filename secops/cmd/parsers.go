/*
Copyright © 2024 calebryant

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// parsersCmd represents the parsers command
var parsersCmd = &cobra.Command{
	Use: "parsers",
	Aliases: []string{
		"parser",
		"cbn",
	},
	Short: "A brief description of your command",
	Long:  ``,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		parser, err := cmd.Flags().GetString("parser")
		if err != nil {
			return err
		}
		// mark the 'parser' flag as optional for list and create commands
		switch cmd.CommandPath() {
		case "secops logtypes parsers list", "secops logtypes parsers create":
			if err := cmd.Flags().SetAnnotation("parser", cobra.BashCompOneRequiredFlag, []string{"false"}); err != nil {
				return err
			}
		}
		request.URL = request.URL.JoinPath("parsers", parser)
		return nil
	},
}

func init() {
	logtypesCmd.AddCommand(parsersCmd)

	parsersCmd.PersistentFlags().StringP("parser", "p", "", "Parser ID")
	parsersCmd.MarkPersistentFlagRequired("parser")

	// listParsersCmd.Flags().StringP("logtype", "l", "-", "SecOps log type label") // override the logtype flag so it can have a "-" default value
	// listParsersCmd.Flags().StringP("filter", "f", "", "A filter which should follow the guidelines of AIP-160 (https://google.aip.dev/160)")
	// parsersCmd.AddCommand(&listParsersCmd)

	// var createParserCmd = *postCmd
	// parsersCmd.AddCommand(&createParserCmd)

	// var getParsersCmd = *getCmd
	// parsersCmd.AddCommand(&getParsersCmd)

	// var activateParserCmd = activateCmd
	// parsersCmd.AddCommand(&activateParserCmd)

	// var deactivateParserCmd = deactivateCmd
	// parsersCmd.AddCommand(&deactivateParserCmd)

	// var deleteParserCmd = *deleteCmd
	// deleteParserCmd.Flags().BoolP("force", "f", false, "Force delete an active parser")
	// parsersCmd.AddCommand(&deleteParserCmd)
}
