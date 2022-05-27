package toolx

import "reflect"

func InArray(val interface{}, array interface{}) (exists bool) {
	exists = false
	kind := reflect.TypeOf(array).Kind()
	switch kind {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				exists = true
				return
			}
		}
	}
	return
}
