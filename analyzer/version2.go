package analyzer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func V2(){
	file, err := os.Open("data/logs.log")
	
	if err != nil {
		fmt.Println("ERROR: ",err)
		return
	}

	scanner := bufio.NewScanner(file)
	
	var data_count map[string]int = make(map[string]int)
	var error_count map[string]int = make(map[string]int)
	for scanner.Scan(){
		if strings.Contains(scanner.Text(), "INFO"){
			data_count["INFO"]++
		}else if strings.Contains(scanner.Text(), "WARN"){
			data_count["WARN"]++
		}else if strings.Contains(scanner.Text(), "ERROR"){
			data_count["ERROR"]++
			parts := strings.Split(scanner.Text()," ERROR ")
			if len(parts) < 2{
				error_count["ERROR"]++
			}else{
				error_count[parts[1]]++
			}
		}

	}
	var total int64 = int64(data_count["INFO"])+int64(data_count["WARN"])+int64(data_count["ERROR"])
	if total == 0{
		fmt.Println("No Data Found!")
		return
	}
	fmt.Println("INFO :",data_count["INFO"])
	fmt.Println("WARNING :",data_count["WARN"])
	fmt.Println("ERROR :",data_count["ERROR"])

	var max int
	var most string
	for key, value := range error_count{
		if value > max{
			most = key
			max = value
		}
	}
	var errorPercent float64 = (float64(error_count[most])/float64(data_count["ERROR"]))*100

	fmt.Printf("Most Common Error: %s\nOccurrance: %.2f%%\n",most,errorPercent)
}