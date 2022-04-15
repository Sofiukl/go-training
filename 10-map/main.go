package main

import "fmt"

// similar like java map, python dictionary
// key value pair
type department map[string]string

func main() {
	departmentMap := make(department)
	logMap(departmentMap)
	departmentMap.put("dep1", "backend")
	departmentMap.put("dep2", "frontend")
	logMap(departmentMap)
	delete(departmentMap, "dep1")
	logMap(departmentMap)
}

// method on type department, hold the reference
func (d department) put(key string, val string) {
	d[key] = val
}

// func function_name( [parameter list] ) [return_types]
// {
//    body of the function
// }
func logMap(m map[string]string) {
	// range loop
	for key, val := range m {
		fmt.Println(key, ":", val)
	}
}
