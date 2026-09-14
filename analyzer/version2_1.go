package analyzer

import (
		"strings"
		"os"
		"bufio")

type LogDetails struct{
	InfoCount int64
	WarningCount int64
	ErrorCount int64
	MostFreqError string
	ErrorPercent float64

}

func V2_1()(LogDetails){
	file, err := os.Open("data/logs.log")
	
	if err != nil {
		panic(err)
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
		panic("No Data Found!")
	}
	
	var max int
	var most string
	for key, value := range error_count{
		if value > max{
			most = key
			max = value
		}
	}
	
	var errorPercent float64 = (float64(error_count[most])/float64(data_count["ERROR"]))*100

	return LogDetails{
		InfoCount: int64(data_count["INFO"]),
		WarningCount: int64(data_count["WARN"]),
		ErrorCount: int64(data_count["ERROR"]),
		MostFreqError: most,
		ErrorPercent: errorPercent,
	}
}