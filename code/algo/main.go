package src

import (
	"strings"
	"fmt"
)


//func main() {
//	input := []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}
//	result := numUniqueEmails(input)
//	fmt.Println(result)
//}


func numUniqueEmails(emails []string) int {
	length := len(emails)
	result := []string{}
	for  i:=0;i<length;i++{
		part := strings.Split(emails[i],"@")
		local :=part[0]
		tmps :=strings.Split(local,"+")
		tmp := tmps[0]
		addr := strings.Replace(tmp, ".","",-1)
		addr = addr+part[1]
		result = append(result,addr)
	}
	uniqueResult := unique(result)
	return len(uniqueResult)
}

func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
			fmt.Println(entry)

		}
	}
	return list
}