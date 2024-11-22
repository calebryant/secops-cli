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
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

func newActivateCmd(short, long string) *cobra.Command {
	c := &cobra.Command{
		Use:   "activate",
		Short: short,
		Long:  long,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// initialize the request object
			request.Method = http.MethodPost
			request.Body = nil
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			request.URL.Path += ":activate"
			return nil
		},
	}
	return c
}

func newDeactivateCmd(short, long string) *cobra.Command {
	c := &cobra.Command{
		Use:   "deactivate",
		Short: short,
		Long:  long,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// initialize the request object
			request.Method = http.MethodPost
			request.Body = nil
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			request.URL.Path += ":deactivate"
			return nil
		},
	}
	return c
}

func newGetCmd(short, long string) *cobra.Command {
	c := &cobra.Command{
		Use:   "get",
		Short: short,
		Long:  long,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// initialize the request object
			request.Method = http.MethodGet
			request.Body = nil
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return c
}

func newCreateCmd(short, long string) *cobra.Command {
	c := &cobra.Command{
		Use: "create",
		Aliases: []string{
			"new",
			"submit",
		},
		Short: short,
		Long:  long,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// initialize the request object
			request.Method = http.MethodPost
			request.Body = nil
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return c
}

func newDeleteCmd(short, long string) *cobra.Command {
	c := &cobra.Command{
		Use: "delete",
		Aliases: []string{
			"del",
		},
		Short: short,
		Long:  long,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// initialize the request object
			request.Method = http.MethodDelete
			request.Body = nil
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return c
}

func newListCmd(short, long string) *cobra.Command {
	c := &cobra.Command{
		Use:   "list",
		Short: short,
		Long:  long,
		Aliases: []string{
			"ls",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			var pageSize int
			var pageToken string
			// initialize the request object
			request.Method = http.MethodGet
			request.Body = nil
			// apply query params
			if pageSize, err = cmd.Flags().GetInt("size"); err != nil {
				panic(err)
			}
			if pageToken, err = cmd.Flags().GetString("next"); err != nil {
				panic(err)
			}
			q := request.URL.Query()
			if pageSize > 0 {
				q.Set("pageSize", strconv.Itoa(pageSize))
			}
			if pageToken != "" {
				q.Set("pageToken", pageToken)
			}
			request.URL.RawQuery = q.Encode()
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	c.Flags().Int("size", 0, "Number of results to return")
	c.Flags().String("next", "", "Next page token value for paginated results")
	return c
}
