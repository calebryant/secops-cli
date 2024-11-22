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
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// logtypesCmd represents the logtypes command
var logtypesCmd = &cobra.Command{
	Use: "logtypes",
	Aliases: []string{
		"lt",
	},
	Short: "Manage SecOps log types and child resources of log types",
	Long:  ``,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		logtype, err := cmd.Flags().GetString("logtype")
		if err != nil {
			return err
		}
		request.URL = request.URL.JoinPath("logTypes", logtype)
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(request.URL.String())
		os.Exit(0)
		return nil
	},
}

var listLogTypesCmd = newListCmd(
	"List log types",
	"",
)

func init() {
	logtypesCmd.PersistentFlags().StringP("logtype", "l", "", "SecOps log type label")
	logtypesCmd.MarkPersistentFlagRequired("logtype")

	// method commands
	logtypesCmd.AddCommand(listLogTypesCmd)
	listLogTypesCmd.PreRun = func(cmd *cobra.Command, _ []string) {
		flagMarkUnrequired(cmd, "logtype")
	}
}

func flagMarkUnrequired(cmd *cobra.Command, name string) {
	cmd.Flags().SetAnnotation(name, cobra.BashCompOneRequiredFlag, []string{"false"})
}
