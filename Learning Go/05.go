package main

func main() {
	var mapMaybe map[int]string
	mapMaybe = make(map[int]string)
	mapMaybe[1] = "one"
	mapMaybe[2] = "two"

	for k, v := range mapMaybe {
		println(k, v)
	}

	var crazyMap map[string]map[string]int
	crazyMap = make(map[string]map[string]int)

	sans := make(map[string]int)
	sans["hp"] = 1
	sans["atk"] = 50
	sans["def"] = 9999

	asgore := map[string]int{
		"hp":  9999,
		"atk": 80,
		"def": 80,
	}

	crazyMap["sans"] = sans
	crazyMap["asgore"] = asgore

	for k, v := range crazyMap {
		println(k)
		for k2, v2 := range v {
			println(k2, v2)
		}
	}
}
