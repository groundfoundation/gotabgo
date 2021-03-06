/*
Copyright © 2021 The Authors of gotabgo

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// serverinfoCmd represents the serverinfo command
var serverinfoCmd = &cobra.Command{
	Use:   "serverinfo",
	Short: "returns information about your Tableau server",
	Long: `Makes a call to your Tableau server and requests information
	about the server. Response will include server version, build and API
	version.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("serverinfo called")
		fmt.Println(options.server)
		fmt.Println(tabApi)
		si, e := tabApi.ServerInfo()
		if e != nil {
			return e
		}
		log.Debugf("serverinformatio: %v", si)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverinfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverinfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverinfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
