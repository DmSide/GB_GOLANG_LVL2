package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"./easyJSONPackage"
)

func SQLConverter(sql string, args ...interface{}) (string, []interface{}) {
	var (
		resultArgs []interface{}
		resultSQL  string
	)
	chunks := strings.Split(sql, "?")
	if len(chunks) != len(args)+1 {
		fmt.Println(errors.New("length of sql string and args should be the same"))
	}
	for i, chunk := range chunks {
		resultSQL = resultSQL + chunk
		if i != len(chunks)-1 {
			resultSQL = resultSQL + "?"
		}

		if i >= len(args)-1 {
			continue
		}

		if arr, ok := args[i].([]int); ok {
			for j, val := range arr {
				if j != 0 {
					resultSQL = resultSQL + ",?"
				}

				resultArgs = append(resultArgs, val)
			}
		} else if arr, ok := args[i].([]string); ok {
			for j, val := range arr {
				if j != 0 {
					resultSQL = resultSQL + ",?"
				}

				resultArgs = append(resultArgs, val)
			}
		} else {
			resultArgs = append(resultArgs, args[i])
		}
	}
	return resultSQL, resultArgs
}

func main() {
	//Напишите функцию, которая на вход получает запрос SQL и произвольные параметры, среди которых могут быть как обычные значения (строка, число) так и слайсы таких значений.
	//Позиция каждого переданного параметра в запросе SQL обозначается знаком "?".
	//Функция должна вернуть запрос SQL, в котором для каждого параметра-слайса количество знаков "?" будет через запятую размножено до количества элементов слайса, а вторым ответом вернуть слайс из параметров, которые соответствуют новым позициям знаков "?".
	//Пример:
	//Вызов: func ( "SELECT * FROM table WHERE deleted = ? AND id IN(?) AND count < ?", false, []int{1, 6, 234}, 555 )
	//Ответ: "SELECT * FROM table WHERE deleted = ? AND id IN(?,?,?) AND count < ?", []interface{}{ false, 1, 6, 234, 555 }
	a, b := SQLConverter("SELECT * FROM table WHERE deleted = ? AND id IN(?) AND count < ?", false, []int{1, 6, 234}, 555)
	fmt.Println(a)
	fmt.Println(b)

	//Сделайте кодогенерацию с помощью easyjson для любой Вашей структуры.
	d := &easyJSONPackage.JSONData{}
	//_ = d.UnmarshalJSON([]byte(`{"Data" : ["One", "Two", "Three"]} `))
	// Or you could also use
	data := []byte(`{"Data" : ["One", "Two", "Three"]} `)
	_ = json.Unmarshal(data, d) // this will also call this d.UnmarshalJSON
	fmt.Println(d)
}
