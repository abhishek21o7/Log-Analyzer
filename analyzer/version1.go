package analyzer
import ("fmt"
		"strings")

func V1(logs []string){
	var count map[string]int = map[string]int{"INFO":0,"WARNING":0,"ERROR":0}
	
	for _, log := range logs{
		if strings.Contains(log,"INFO"){
			count["INFO"]++
		}else if strings.Contains(log,"WARNING"){
			count["WARNING"]++
		}else{
			count["ERROR"]++
		}
	}
	var max int
	var most string
	for key, value := range count{
		if value > max{
			most = key
			max = value
		}
		fmt.Printf("%s : %d\n",key, value)
	}
	fmt.Printf("Most common error:\n%s\n",most)
}