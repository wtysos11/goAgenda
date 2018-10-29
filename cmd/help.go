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

	"github.com/spf13/cobra"
)


var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "helo user to do something",
	Long: `

	you can use this app to create or remove meetings.Also you must register a user to have the rights to use the functions.

	Usage:
		agenda [command]

	Available Commands:
		add         To add Participator of the meeting
		clear       clear all the meeting created by the current user
		create      To create a new meeting
		delete      To delete your account in Agenda
		deleteM     delete meeting with the title [title]
		help        Help about any command
		login       Using UserName with PassWord to login Agenda.
		logout      To logout Agenda
		queryM      To query all the meeting have attended during [StartTime] and [EndTime]
		queryU      To query all the users' names
		quit        quit the meeting with the title [title]
		register    Register a new user
		remove      To remove Participator from the meeting

	Flags:
		--config string   config file (default is $HOME/.agenda.yaml)
		-h, --help            help for agenda
		-t, --toggle          Help message for toggle

	Use "agenda [command] --help" for more information about a command.

	`,
	Run: func(cmd *cobra.Command, args []string) {
		
		fmt.Println("help called")
		fmt.Println(args)
		if(len(args) == 2 && args[0]=="user" && args[1] == "register"){
			fmt.Println("register here")
		}else{
			fmt.Println("Otherwise, help here")
		}
		
	},
}


func init() {
	rootCmd.AddCommand(helpCmd)
	//rootCmd.SetHelpCommand(helpCmd)


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
