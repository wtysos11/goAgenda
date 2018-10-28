package entity

import (
    "io/ioutil"
    "fmt"
    "encoding/json"
    "os"
    "errors"
    "strings"
    "strconv"
)

type Meeting struct{
    Creator string
    StartTime string
    EndTime string
    Title string
    UserList []string
}

type AgendaTime struct{
    Year int
    Month int
    Day int
    Hour int
    Minute int
    Second int
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

func String2Time (time string) (AgendaTime,error){
    mainSpe := strings.Index(time,"/")
    if mainSpe == -1 {
        return AgendaTime{},errors.New("Time format incorrect, don't have /")
    }

    front := strings.Split(time[:mainSpe],"-")
    back := strings.Split(time[mainSpe+1:],":")

    year,yerr := strconv.Atoi(front[0])
    month,merr := strconv.Atoi(front[1])
    day,derr := strconv.Atoi(front[2])
    hour,herr := strconv.Atoi(back[0])
    minute, mierr := strconv.Atoi(back[1])
    second, serr := strconv.Atoi(back[2])

    if yerr!=nil || merr!=nil || derr!=nil || herr!=nil || mierr!=nil || serr!=nil {
        return AgendaTime{},errors.New("Having trouble converting time string to int")
    }

    return AgendaTime{year,month,day,hour,minute,second},nil
}

//time1<time2 -1 time1==time2 0 time1>time2 1
func CompareTime (time1 AgendaTime, time2 AgendaTime) int{
    if time1.Year<time2.Year{
        return -1
    } else if time1.Year>time2.Year{
        return 1
    } else{
        if time1.Month<time2.Month{
            return -1
        } else if time1.Month>time2.Month{
            return 1
        } else{
            if time1.Day < time2.Day{
                return -1
            } else if time1.Day>time2.Day{
                return 1
            } else{
                if time1.Hour<time2.Hour{
                    return -1
                } else if time1.Hour>time2.Hour{
                    return 1
                } else {
                    if time1.Minute < time2.Minute{
                        return -1
                    } else if time1.Minute > time2.Minute{
                        return 1
                    } else{
                        if time1.Second < time2.Second{
                            return -1
                        } else if time1.Second > time2.Second {
                            return 1
                        } else {
                            return 0
                        }
                    }
                }
            }
        }
    }
}