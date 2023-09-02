package utils

import "regexp"

const RegHan = `[\p{Han}]+`

func FilterHan(str string) string {

	reg := regexp.MustCompile(RegHan)
	return reg.FindString(str)

}
