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

	1. To add a register: 
	
	instruction:	register -u [UserName] -p [Pass] -e [Email] -t [Phone]
		
		[Username] register's name
		[Pass] register's Password
		[Email] register's email
		[Phone] register's phone-number
	
	2. To remove a register

	instruction:	delete

	(attention: you can delete your account in the database of Agenda)

	3. To query user

	instruction:	queryU

	(attention: you can query all user's name who has registed)

	4. To login 

	instruction:	login -u [UserName] -p [PassWord]

		[Username] register's name
		[Pass] register's Password

	5. To logout

	instruction:	logout

	`,
	Run: func(cmd *cobra.Command, args []string) {
		
		fmt.Println("help called")
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
