package entity


type User struct{
	Username string `json:"username"`
    Password string `json:"password"`
    Email string `json:"email"`
    Telphone string `json:"telphone"`
}
/* 函数名：JSONDecode
 * 参数： string json，表示收到的json字符串
 * 返回值： error
 * 根据所给的JSON修改执行的User的值
 */
 /*
func (user User) JSONDecode (jsonStr string) err{
	return json.Unmarshal([]byte(jsonStr),user)
}*/