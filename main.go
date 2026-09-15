package main

import (
    "fmt"
    "mini_log_analyzer/analyzer"
)
func main(){
    //For version 1 and version1_1
    /*logs := []string{
        "INFO User login",
        "ERROR Database failed",
        "INFO User logout",
        "WARNING High memory",
        "ERROR Database failed",
    }
    analyzer.V1()
    analyzer.V1_1()*/
	//analyzer.V2()

    logDetails, err := analyzer.V2_1()

    if err != nil{
        fmt.Println("Error : ",err)
        return
    }

    fmt.Printf("INFO : %d\nWARNING : %d\nERROR : %d\n",
                logDetails.InfoCount,
                logDetails.WarningCount,
                logDetails.ErrorCount)
    fmt.Printf("Most Common Error : %s\nOccurrence : %.2f%%\n",
                logDetails.MostFreqError,
                logDetails.ErrorPercent)

}
