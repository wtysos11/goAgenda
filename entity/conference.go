package entity

import (
    "io/ioutil"
    "fmt"
    "encoding/json"
    "os"
)

type Meeting struct{
    Creator string
    StartTime string
    EndTime string
    Title string
    UserList []string
}

type time struct{
    year int
    month int
    day int
    hour int
    minute int
    second int
}

func ReadMeetingFromFile (filePath string) ([]Meeting,error) {
	var conference []Meeting
	fmt.Println("In reading Meeting file")
	str,err := ioutil.ReadFile(filePath)
	if err!=nil {
		return conference,err
	}
	jsonStr := string(str)
	fmt.Printf("%s\n",jsonStr)
	
	json.Unmarshal([]byte(jsonStr),&conference)
	return conference,nil
}

func WriteMeetingToFile (filePath string, meeting []Meeting) error{
	if data,err:=json.Marshal(meeting);err==nil{
		return ioutil.WriteFile(filePath,[]byte(data),os.ModeAppend)
	} else{
		return err
	}
}