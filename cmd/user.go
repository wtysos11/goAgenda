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
	"github.com/wtysos11/goAgenda/entity"
	"github.com/spf13/cobra"
	"encoding/json"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("agenda user")
		username, _ := cmd.Flags().GetString("user")
		password, _:= cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")

		fmt.Println("register called by " + username)
		fmt.Println("register's password is "+password)
		fmt.Println("register's email is "+email)
		if(len(args)>0){
			switch(args[0]){
				case "register":{
					fmt.Println("register")
				}
				case "login":{
					fmt.Println("user login")
				}
				case "logout":{
					fmt.Println("user logout")
				}
				case "lookup":{
					fmt.Println("user lookup")
				}
				case "delete":{
					fmt.Println("user delete")
				}
			}
		}

		var myuser []entity.User
		fmt.Println("JSON decode to structure test1")
		jsonStr := `[{"username":"wtysos11","password":"123456","email":"wtysos11"},{"username":"wtysos11","password":"123456","email":"wtysos11"}]`
		json.Unmarshal([]byte(jsonStr),&myuser)
		fmt.Println(myuser)
		
		if data,err:=json.Marshal(myuser);err==nil{
			fmt.Printf("%s\n",data)
		}

	},
}


func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.Flags().StringP("user","u","Anonymous","Help message for username")
	userCmd.Flags().StringP("password","p","123456","Help message for password")
	userCmd.Flags().StringP("email","e","your_email@no.work","Help message for email")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
