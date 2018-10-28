package entity


type Meeting struct{
    StartTime string
    EndTime string
    Title string
    UserList []User
}

type time struct{
    year int
    month int
    day int
    hour int
    minute int
    second int
}