package BaseT

import "strings"

func Contains(arr *[]string,target string, isSensitive bool) bool  {
	if arr == nil{
		return false
	}
	if !isSensitive{
		target = strings.ToUpper(target)
		for _,item := range *arr{
			if strings.ToUpper(item) == target{
				return  true
			}else {
				continue
			}
		}
	}else{
		for _,item := range *arr{
			if item == target{
				return  true
			}else {
				continue
			}
		}
	}
	return  false

}
