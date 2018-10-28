package entity

import(
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
)

type User struct{
	Username string `json:"username"`
    Password string `json:"password"`
    Email string `json:"email"`
    Telphone string `json:"telphone"`
}
/* 函数名：
 * 参数： 
 * 返回值： users数组，error
 * 
 */
 
func ReadUserFromFile (filePath string) ([]User,error) {
	var users []User
	fmt.Println("In reading user file")
	str,err := ioutil.ReadFile(filePath)
	if err!=nil {
		return users,err
	}
	jsonStr := string(str)
	fmt.Printf("%s\n",jsonStr)
	
	json.Unmarshal([]byte(jsonStr),&users)
	return users,nil
}

func WriteUserToFile (filePath string, users []User) error{
	if data,err:=json.Marshal(users);err==nil{
		return ioutil.WriteFile(filePath,[]byte(data),os.ModeAppend)
	} else{
		return err
	}
}