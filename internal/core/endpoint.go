package core



func GetAllByKey(parsedjson map[string]any, key string) (any){
	for i, e := range parsedjson{
		if key == i {
			return  e
		}
	}
		
	return nil
}

func GetElementById(parsedjson map[string]any, key string, id int) (map[string]any){
	all:= GetAllByKey(parsedjson,key)
	
	for _,v := range all.([]any) {
		if j,ok := v.(map[string]any); ok {
			if float64(id) == j["id"] {
				return j
			}
		}
	}

	return nil
}

