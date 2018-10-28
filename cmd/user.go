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
	"io/ioutil"
	"os"
	. "github.com/wtysos11/goAgenda/entity"
	"github.com/spf13/cobra"
)

const userPlace = "user.txt"
const cachePlace = "cache.txt"

//legal check, don't implement yet
func userLegalCheck(userinfo []User,username string, password string,email string ,telphone string) (bool,error){
	return true,nil
}

func checklogin() (bool,error){
	b,err := ioutil.ReadFile(cachePlace)
	if err!=nil {
		return false,err
	}
	str := string(b)

	if str == "logout"{
		return false,nil
	} else{
		return true,nil
	}
}

func getLoginUsername() (string,error){
	b,err := ioutil.ReadFile(cachePlace)
	if err!=nil {
		return "",err
	}
	return string(b),nil
}

func userLogin(username string) error{
	return ioutil.WriteFile(cachePlace,[]byte(username),os.ModeAppend)
}

func userLogout() error{
	return ioutil.WriteFile(cachePlace,[]byte("logout"),os.ModeAppend)
}


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
		//reading file from store place
		userinfo,userReadingerr := ReadUserFromFile(userPlace)
		if userReadingerr!=nil {
			fmt.Println(userReadingerr)
		}

		//get flags
		username, _ := cmd.Flags().GetString("user")
		password, _:= cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		telphone,_ := cmd.Flags().GetString("telphone")

		/*
		fmt.Println("In userCmd begin Flag test.")
		fmt.Println("register called by " + username)
		fmt.Println("register's password is "+password)
		fmt.Println("register's email is "+email)
		fmt.Println("register's telphone is "+telphone)
		fmt.Println("Flag test over\n")
*/

		//
		if len(args)>0 {
			switch (args[0]){
				case "register":{
					fmt.Println("register")
					//legal check for username(unique),password,email,telphone
					if pass,error := userLegalCheck(userinfo,username,password,email,telphone); pass == false {
						fmt.Println("Register Failed")
						fmt.Println(error)
						return
					}

					//if pass legal check, add it to userFile
					userinfo = append(userinfo,User{username,password,email,telphone})
					//store the user file into userPlace
					WriteUserToFile(userPlace,userinfo)
					fmt.Println("Register success")
				}
				case "login":{
					fmt.Println("user login")
					//check from cache whether the status is login.
					if check,error := checklogin(); error!=nil{
						fmt.Println(error)
						return
					} else if check {
						fmt.Println("Already Login")
						return
					}
					//validate username and password
					pass := false
					for _,user := range userinfo{
						if user.Username == username && user.Password == password{
							userLogin(user.Username)
							pass = true
							break
						}
					}
					//if no pass, report
					if !pass {
						fmt.Println("login failed")
						return
					}
					
					fmt.Println("Login success. Welcome!")
				}
				case "logout":{
					fmt.Println("user logout")
					//if status is login, make the status logout

					pass,err := checklogin()
					if err!=nil{
						fmt.Println(err)
						return
					} else if !pass {
						fmt.Println("Please login first.")
						return
					}

					userLogout()
					fmt.Println("Logout success")
					
				}
				case "lookup":{
					fmt.Println("user lookup")
					//check the status (login)
					pass,err := checklogin()
					if err!=nil{
						fmt.Println(err)
						return
					} else if !pass {
						fmt.Println("Please login first.")
						return
					}
					//if pass validation, give all info from all users
					fmt.Println("Here is users' info:")
					for _,user := range userinfo{
						fmt.Println(user.Username,user.Email,user.Telphone)
					}
				}
				case "delete":{
					fmt.Println("user delete")
					//check status login
					pass,err := checklogin()
					if err!=nil{
						fmt.Println(err)
						return
					} else if !pass {
						fmt.Println("Please login first.")
						return
					}
					loginUsername,loginErr := getLoginUsername()
					if loginErr!=nil{
						fmt.Println(loginErr)
						return
					}
					//if pass, delete this user and logout
					for i,user := range userinfo{
						if loginUsername == user.Username{
							userinfo = append(userinfo[:i],userinfo[i+1:]...)
							break
						}
					}
					//update the userPlace
					WriteUserToFile(userPlace,userinfo)
				}
				default:{
					fmt.Println("Unknown command")
				}
			}
		}
		
		/*
		fmt.Println("\nio reading and writing test.")
		myuser,err := entity.ReadUserFromFile("user.txt");
		if  err==nil{
			fmt.Println(myuser)
		} else{
			fmt.Println(err)
		}
		entity.WriteUserToFile("test.txt",myuser)
*/

	},
}


func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.Flags().StringP("user","u","Anonymous","Help message for username")
	userCmd.Flags().StringP("password","p","123456","Help message for password")
	userCmd.Flags().StringP("email","e","your_email@no.work","Help message for email")
	userCmd.Flags().StringP("telphone","t","11022223333","Help message for telphone")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
