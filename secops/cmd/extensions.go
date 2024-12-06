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

var (
	extensionId string
)

// extensionsCmd represents the extensions command
var extensionsCmd = &cobra.Command{
	Use: "extensions",
	Aliases: []string{
		"extension",
		"ext",
	},
	Short: "A brief description of your command",
	Long:  ``,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		extension, err := cmd.Flags().GetString("extension")
		if err != nil {
			return err
		}
		request.URL = request.URL.JoinPath("parserExtensions", extension)
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var listExtensionsCmd = newListCmd(
	"List parser extensions",
	"",
)

var getExtensionCmd = newGetCmd(
	"Get a parser extension",
	"",
)

func init() {
	logtypesCmd.AddCommand(extensionsCmd)
	extensionsCmd.PersistentFlags().StringVarP(&extensionId, "extension", "x", "", "Parser extension ID")
	extensionsCmd.MarkPersistentFlagRequired("extension")

	// get
	extensionsCmd.AddCommand(getExtensionCmd)

	// list
	extensionsCmd.AddCommand(listExtensionsCmd)
	listExtensionsCmd.Flags().StringP("logtype", "l", "-", "Parser log type label")
	listExtensionsCmd.PreRun = func(cmd *cobra.Command, _ []string) {
		flagMarkUnrequired(cmd, "extension")
	}
}
