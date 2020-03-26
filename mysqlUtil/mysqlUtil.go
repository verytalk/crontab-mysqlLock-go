package mysqlUtil


import (
	"crontask/configUtil"
	"crontask/fileUtil"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var dbConn *sql.DB;

func getConnection()(*sql.DB){
	mysqlUrl := ""+configUtil.Config.Mysql.User+":"+configUtil.Config.Mysql.Password+"@tcp("+configUtil.Config.Mysql.Host+":"+configUtil.Config.Mysql.Port+")"+"/INFORMATION_SCHEMA"
		//连接数据库查询
		db, err := sql.Open("mysql", mysqlUrl)
		dbConn = db
		if err != nil {
			log.Fatal(err.Error())
		}
	//最终关闭数据库
	//defer db.Close()
	return dbConn
}


func runSelect(mSql string){
	db := getConnection()
	defer db.Close()

	//mSql := "SELECT * FROM INFORMATION_SCHEMA.INNODB_LOCKS"
	rows, err := db.Query(mSql)
	if err != nil {
		log.Fatal(err.Error())
	}
	//获取列名
	columns, _ := rows.Columns()
	//定义一个切片,长度是字段的个数,切片里面的元素类型是sql.RawBytes
	values := make([]sql.RawBytes, len(columns))
	//定义一个切片,元素类型是interface{} 接口
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		//把sql.RawBytes类型的地址存进去了
		scanArgs[i] = &values[i]
	}

	//获取字段值
	var result []map[string]string
	for rows.Next() {
		res := make(map[string]string)
		rows.Scan(scanArgs...)
		for i, col := range values {
			res[columns[i]] = string(col)
		}
		result = append(result, res)
	}
	LoggingText := ""
	//遍历结果
	for _, r := range result {
		for k, v := range r {
			//log.Printf("%s = %s", k, v)
			LoggingText +=  k + " = " + v+"\n"
		}
	}
	if(LoggingText != ""){
		LoggingText = mSql+"\n"+LoggingText
		fileUtil.Tracefile(LoggingText);
		log.Printf(LoggingText)
	}
	rows.Close()
}


func Run(){
	runSelect("SELECT * FROM INFORMATION_SCHEMA.INNODB_LOCKS")
	runSelect("SELECT * FROM INFORMATION_SCHEMA.innodb_trx")
	runSelect("SELECT * FROM INFORMATION_SCHEMA.INNODB_LOCK_WAITS")
}
