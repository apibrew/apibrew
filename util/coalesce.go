package util

func Coalesce[T interface{}](val *T, defaultVal T) T {
	if val == nil {
		return defaultVal
	}

	return *val
}

//func CoalesceP[T interface{}](val *T, defaultVal T) T {
//	if val == nil {
//		return defaultVal
//	}
//
//	return Coalesce(*val, defaultVal)
//}
//
//func CoalescePP[T interface{}](val **T, defaultVal T) T {
//	if val == nil {
//		return defaultVal
//	}
//
//	return CoalesceP(*val, defaultVal)
//}
