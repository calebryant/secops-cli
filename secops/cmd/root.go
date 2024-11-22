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
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/spf13/cobra"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/googleapi"
)

const (
	dataplaneUrl   string = "chronicle.googleapis.com"
	dataplaneScope string = "https://www.googleapis.com/auth/cloud-platform"
)

var (
	request *http.Request = &http.Request{}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "secops",
	Short: "CLI tool to interact with SecOps APIs",
	Long:  ``,
	// all child commands run this function before Run unless it is overwritten
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var version, project, location, instance string
		var err error
		if version, err = cmd.Flags().GetString("version"); err != nil {
			return err
		}
		if location, err = cmd.Flags().GetString("location"); err != nil {
			return err
		}
		if project, err = cmd.Flags().GetString("project"); err != nil {
			return err
		}
		if instance, err = cmd.Flags().GetString("instance"); err != nil {
			return err
		}
		// initialize the request object
		u := &url.URL{
			Scheme: "https",
			Host:   location + "-" + dataplaneUrl,
			Path:   path.Join(version, "projects", project, "locations", location, "instances", instance),
		}
		request, err = http.NewRequest("", u.String(), nil)
		if err != nil {
			return err
		}
		return nil
	},
	// all child commands run this function after Run unless it is overwritten
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		// create the auth client
		client, err := google.DefaultClient(context.Background(), dataplaneScope)
		if err != nil {
			return err
		}
		// make the request
		response, err := client.Do(request)
		// if err != nil && strings.Contains(err.Error(), "reauth") {
		// 	reauthCmd := exec.Command("gcloud", "auth", "login", "--update-adc")
		// 	reauthCmd.Run()
		// 	response, err = client.Do(request)
		// }
		if err != nil {
			return err
		}
		// print the response to terminal
		if response == nil {
			return nil
		}
		defer log.Println(response.Status, request.Method, request.URL.String())
		err = googleapi.CheckMediaResponse(response)
		if err != nil {
			return err
		}
		if response.Body != nil {
			defer response.Body.Close()
		}
		data, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(data))
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.ExecuteContext(context.Background())
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("version", "v", "v1alpha", "Chronicle API version (v1, v1alpha, v1beta)")

	rootCmd.PersistentFlags().StringP("location", "R", "us", "SecOps availability region")

	rootCmd.PersistentFlags().StringP("project", "P", "", "GCP project ID linked to the SecOps tenant")
	rootCmd.MarkPersistentFlagRequired("project")

	rootCmd.PersistentFlags().StringP("instance", "I", "", "SecOps customer ID")
	rootCmd.MarkPersistentFlagRequired("instance")

	// child commands
	rootCmd.AddCommand(logtypesCmd)
}
