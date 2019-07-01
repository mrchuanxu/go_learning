package go_mysql

// import module
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	// "reflect"
	"testing"
)
// 配置打开文件
type SqlOperating struct{
	db *sql.DB
}

func OpenSql(sql_name string,user_settings string)(*sql.DB,error){
	db,err := sql.Open(sql_name,user_settings)
	return db,err
}
// 增删改
func (sqlDb *SqlOperating)Inse_Dele_UpData(str string){
	_,err := sqlDb.db.Query(str)
	CheckErr(err)
}

// 查
func (sqlDb *SqlOperating)SelectData(str string)(*sql.Rows,error){
	result,err := sqlDb.db.Query(str)
	return result,err
}

// check error
func CheckErr(err error){
	if err != nil{
		fmt.Println(err)
		panic("connect wrong!")
	}
}

func printResult(query *sql.Rows){
	column,_ := query.Columns()
	values := make([][]byte,len(column))
	scans := make([]interface{},len(column))

	for i := range values{
		scans[i] = &values[i]
	}

	results := make(map[int]map[string]string)

	i:=0
	for query.Next(){
		if err := query.Scan(scans...);err != nil{
			fmt.Println(err)
			return
		}
		row := make(map[string]string)

		for k,v:=range values{
			key := column[k]
			row[key] = string(v)
		}
		results[i] = row
		i++
	}
	for k,v := range results{
		fmt.Println(k,v)
	}
}

func TestMysql(t *testing.T){
	sqlDbase := new(SqlOperating)
	var err error
	sqlDbase.db,err = OpenSql("mysql","root:rootroot@tcp(127.0.0.1:3306)/?charset=utf8")
	CheckErr(err)
	// sqlDbase.Inse_Dele_UpData("create table tmpdb.transdb(id int, name varchar(20), pass varchar(20))")
	sqlDbase.Inse_Dele_UpData("insert into tmpdb.transdb values(101, '姓名1', 'address1'), (102, '姓名2', 'address2'), (103, '姓名3', 'address3'), (104, '姓名4', 'address4')")
	query,qerr := sqlDbase.SelectData("select * from tmpdb.transdb")
	CheckErr(qerr)
	// v := reflect.ValueOf(query)
	// t.Log(v)
	printResult(query)
	defer sqlDbase.db.Close() // 配对以上的mysql
}

/***
* 更多关于数据库操作，参照以下链接
* https://golang.org/pkg/database/sql/
***/