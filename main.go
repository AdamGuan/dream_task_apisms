package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() { //main函数
	mode := "onlineProd"
	//default
	dbUsername := "root"
	dbPwd := "root"
	dbName := "dream_api_sms_v2"
	dbPort := "3306"
	switch mode {
	case "dev":
		dbUsername = "root"
		dbPwd = "root"
		dbName = "dream_api_sms_v2"
	case "155Dev":
		dbUsername = "root"
		dbPwd = "root"
		dbName = "dream_api_sms_v2"
	case "onlineTest":
		dbUsername = "v2_test"
		dbPwd = "root"
		dbName = "dream_api_sms_v2_test"
	case "onlineProd":
		dbUsername = "center_user_v2"
		dbPwd = "root"
		dbName = "dream_api_sms_v2"
	}

	db, err := sql.Open("mysql", dbUsername+":"+dbPwd+"@tcp(localhost:"+dbPort+")/"+dbName+"?charset=utf8&loc=Asia%2FShanghai")
	if err != nil {
		panic(err.Error())       //抛出异常
		fmt.Println(err.Error()) //仅仅是显示异常
	}
	defer db.Close() //只有在前面用了 panic 这时defer才能起作用，如果链接数据的时候出问题，他会往err写数据
	delete_sql := "DELETE t_class.* FROM t_class,(SELECT t_class.F_class_id FROM t_class LEFT JOIN t_user ON t_class.F_class_id = t_user.F_class_id GROUP BY t_class.F_class_id HAVING count(t_user.F_user_name) <= 0) as t WHERE t_class.F_class_id = t.F_class_id"
	_, err = db.Exec(delete_sql)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("success1!")
	}

	delete_sql = "DELETE FROM t_sms_action_valid WHERE F_last_timestamp <= '"+GetDateTimeBeforeMinute(2)+"'"
	_, err = db.Exec(delete_sql)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("success2!")
	}

	delete_sql = "DELETE FROM t_email_action_valid WHERE F_last_timestamp <= '"+GetDateTimeBeforeMinute(2)+"'"
	_, err = db.Exec(delete_sql)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("success3!")
	}

	db.Close() //关闭数据库
}

func GetDateTimeBeforeMinute(num int)string{
	return time.Now().Add(-time.Minute * time.Duration(num)).Format("2006-01-02 15:04:05")
}
