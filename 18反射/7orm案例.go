package main
import(
	"errors"
	"fmt"
	"reflect"
	"strings"
)
type Student struct{
	Name string `feng-orm:"name"`
	Age int 	`feng-orm:"age"`
}
type UserInfo struct{
	Id int `feng-orm:"id"`
	Name string `feng-orm:"name"`
	Age int `feng-orm:"age"`
}
// Find(Student, "name = ?", "fengfeng")
//希望生成 select name, age from Students where name = "fengfeng"
func Find(obj any, query ...any)(sql string, err error){
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct{
		err = errors.New("非结构体")
		return
	}
	// 对第二个参数query[]进行分析
	var where string
	if (len(query) > 0){	// 有where筛选条件
		q := query[0]
		if strings.Count(q.(string), "?")+1 != len(query){
			err = errors.New("参数个数不对")
			return
		}
		// 拼接where语句
		for _, a:= range query[1:]{
			at := reflect.TypeOf(a)
			switch at.Kind(){
			case reflect.Int:
				q = strings.Replace(q.(string), "?", fmt.Sprintf("%d", a.(int)), 1)
			case reflect.String:
				q = strings.Replace(q.(string), "?", fmt.Sprintf("'%s'", a.(string)), 1)
			}
		}
		where += "where" + q.(string)
	}

	var columns []string
	for i := 0; i < t.NumField(); i++{
		field := t.Field(i)
		f := field.Tag.Get("feng-orm")
		columns = append(columns, f)
	}
	name := strings.ToLower(t.Name()) + "s"
	sql = fmt.Sprintf("select %s from %s %s", strings.Join(columns, ","), name, where)
	return
}
func main(){
	sql, err := Find(Student{}, "name = ? and age = ?", "枫枫", 23)
	fmt.Println(sql, err)
	sql, err = Find(UserInfo{}, "id = ?", 1)
	fmt.Println(sql, err)
}