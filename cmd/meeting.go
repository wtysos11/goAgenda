// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	//"github.com/wtysos11/goAgenda/entity"
	"github.com/spf13/cobra"
)

// meetingCmd represents the meeting command
var meetingCmd = &cobra.Command{
	Use:   "meeting",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("meeting called")
		if len(args)>0{
			switch (args[0]){
				case "create":{
					fmt.Println("create")
				}
				case "addUser":{
					fmt.Println("add user")
				}
				case "deleteUser":{
					fmt.Println("delete user")
				}
				case "lookup":{
					fmt.Println("meeting lookup")
				}
				case "cancel":{
					fmt.Println("meeting cancel")
				}
				case "exit":{
				
					fmt.Println("meeting exit")
				}
				case "clear":{
					fmt.Println("meeting clear")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(meetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// meetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// meetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
