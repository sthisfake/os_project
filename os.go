package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type process struct {
	id int
	firstArriveTime int
	arriveTime int
	arrivedState bool
	cpuTime1 int
	cpuTime1State bool
	ioTime int
	ioTimeState bool
	cpuTime2 int
	cpuTime2State bool
	isDone bool
	responseTime int
	turnaroundTime int
	waitingTime int
	lastUsed bool
}



func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file " + filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for " + filePath, err)
	}

	return records
}

func FCFS(processes [5] process) {

	var chart []int
	currentTime:= 0
	flag := 0
	sign := 0

	for flag < 5{



		i:=0
		for i < 5{
			if  processes[i].arriveTime <= currentTime  && processes[i].isDone == false {
				if processes[i].cpuTime1State == false {
					for j := 0; j <processes[i].cpuTime1 ; j++ {

						chart = append(chart , processes[i].id)
						currentTime++

					}
					processes[i].cpuTime1State = true
					processes[i].arriveTime = currentTime + processes[i].ioTime
					sign = 1
					break
				} else {
					for j := 0; j < processes[i].cpuTime2 ; j++ {
						chart = append(chart , processes[i].id)
						currentTime++

					}
					processes[i].cpuTime2State = true
					processes[i].isDone = true
					flag++
					sign = 1
					break
				}
			}
			if i == 4 {
				break

			} else {
				i++

			}



		}
		if sign == 0 {
			chart = append(chart , 0)
			currentTime++
		}

		sign = 0

	}

	printtt(processes , chart , "FCFS")

}

func RR(processes [5] process){
	var chart2 []int
	currentTime := 0
	flag := 0
	sign := 0
	timeQ := 5

	for i :=0 ; i<5 ; i++ {
		processes[i].arriveTime = processes[i].firstArriveTime
	}

	i:=-1

	for flag < 5{
		i=-1
		for i < 5{
			i++
			if  processes[i].arriveTime <= currentTime  && processes[i].isDone == false {
				if processes[i].cpuTime1State == false {
					if timeQ >= processes[i].cpuTime1 {
						for j := 0; j <processes[i].cpuTime1 ; j++ {

							chart2 = append(chart2 , processes[i].id)
							currentTime++

						}
						processes[i].cpuTime1State = true
						processes[i].arriveTime = currentTime + processes[i].ioTime
						sign = 1
						break
					}else{
						for j := 0; j <timeQ ; j++ {

							chart2 = append(chart2 , processes[i].id)
							currentTime++

						}
						processes[i].cpuTime1 = processes[i].cpuTime1 - timeQ
						sign = 1
						break

					}

				} else {
					if timeQ >= processes[i].cpuTime2 {
						for j := 0; j <processes[i].cpuTime2 ; j++ {

							chart2 = append(chart2 , processes[i].id)
							currentTime++

						}
						processes[i].cpuTime2State = true
						processes[i].isDone = true
						flag++
						sign = 1
						break
					}else{
						for j := 0; j <timeQ ; j++ {

							chart2 = append(chart2 , processes[i].id)
							currentTime++

						}
						processes[i].cpuTime2 = processes[i].cpuTime2 - timeQ
						sign = 1
						break

					}
				}
			}
			if i == 4 {
				i=0
				break

			}

		}
		if sign == 0 {
			chart2 = append(chart2 , 0)
			currentTime++
		}

		sign = 0

	}

	printtt(processes , chart2 , "RR")



}

func SJF(processes [5] process){

	var chart []int
	var ready []int
	currentTime:= 0
	flag := 0

	for flag < 5 {

		for i:= 0 ; i<5 ; i++{
			if currentTime>=processes[i].arriveTime && processes[i].isDone == false {
				ready = append(ready , i)
			}
		}

		//if zero in ready
		if len(ready)== 0 {
			chart = append(chart , 0)
			currentTime++
		}
		//if one in ready
		if len(ready) == 1 {
			if processes[ready[0]].cpuTime1State == false {
				for j := 0; j < processes[ready[0]].cpuTime1; j++ {
					chart = append(chart, processes[ready[0]].id)
					currentTime++
				}
				processes[ready[0]].cpuTime1State = true
				processes[ready[0]].arriveTime = currentTime + processes[ready[0]].ioTime
				ready = nil
			} else {
				for j := 0; j < processes[ready[0]].cpuTime2; j++ {
					chart = append(chart, processes[ready[0]].id)
					currentTime++
				}
				processes[ready[0]].cpuTime2State = true
				processes[ready[0]].isDone = true
				flag++
				ready = nil
			}

		}

		if len(ready) > 1 {
			temp:=100
			index:=0
			numb :=0
			for j:=0 ; j<len(ready) ; j++{
				if processes[ready[j]].cpuTime1State == false {
					if temp>processes[ready[j]].cpuTime1 {
						temp = processes[ready[j]].cpuTime1
						index = ready[j]
						numb = 1
					}

				}else {
					if temp>processes[ready[j]].cpuTime2 {
						temp = processes[ready[j]].cpuTime2
						index = ready[j]
						numb = 2
					}

				}
			}
			if numb == 1{
				for j := 0; j < processes[index].cpuTime1; j++ {
					chart = append(chart, processes[index].id)
					currentTime++
				}
				processes[index].cpuTime1State = true
				processes[index].arriveTime = currentTime + processes[index].ioTime
				ready = nil

			} else{
				for j := 0; j < processes[index].cpuTime2; j++ {
					chart = append(chart, processes[index].id)
					currentTime++
				}
				processes[index].cpuTime2State = true
				processes[index].isDone = true
				flag++
				ready = nil

			}



		}

	}

	printtt(processes , chart , "SJF")

}

func printtt(processes [5] process , chart []int , name string){

	averageResponseTime , processes := responseTime(processes , chart)
	averageTurnAroundTime , processes := turnAroundTime(processes, chart)
	averageWaitingTime , processes := waitingTime(processes )

	totalTime := len(chart)
	var idleTime =0
	for i:= 0 ; i<len(chart) ; i++{
		if chart[i] == 0{
			idleTime++
		}
	}


	burstTime := totalTime - idleTime
	cpuRandeman := float64(burstTime) / float64(totalTime)
	throuput := float64(burstTime) / float64(len(processes))

	fmt.Println("\n////////////////////////////////////////////////////////")
	fmt.Println("\n                   		",name ,"                \n")
	fmt.Println("	     response	   turnaround	      waiting")
	for i := 0 ; i<5 ; i++ {
		fmt.Println("p",processes[i].id , "		" , processes[i].responseTime , "		" , processes[i].turnaroundTime,
			"		" , processes[i].waitingTime)
	}
	fmt.Println("----------------------------------------------------------")
	fmt.Println("avg            " , averageResponseTime , "            " , averageTurnAroundTime , "          " , averageWaitingTime, "\n")
	fmt.Println("Total Time : " , totalTime)
	fmt.Println("idle  Time : " , idleTime)
	fmt.Println("Burst Time : " , burstTime)
	fmt.Printf("CPU Utilization : %.2f\n" , cpuRandeman)
	fmt.Println("Throughput : " , throuput)

	fmt.Print("\n\n|")
	for i:=0; i< len(chart) ; i++ {

		temp:= chart[i]
		if temp == 0{
			print("''''''")
		}
		if temp == 1{
			print("  P1  ")
		}
		if temp == 2{
			print("  P2  ")
		}
		if temp == 3{
			print("  P3  ")
		}
		if temp == 4{
			print("  P4  ")
		}
		if temp == 5{
			print("  P5  ")
		}

		fmt.Print("|")


	}
	fmt.Println()


}

func responseTime(processes [5] process , chart []int ) (float64, [5]process) {


	sum:=0
	for i:=0 ; i<5 ; i++ {
		for j:=0; j< len(chart) ; j++{
			if chart[j] == processes[i].id{
				response := j - processes[i].firstArriveTime
				processes[i].responseTime = response
				sum = processes[i].responseTime + sum
				break
			}
		}

	}
	 averageResponseTime := float64(sum) / 5.0
	 return  averageResponseTime , processes


}

func turnAroundTime(processes [5] process , chart []int ) (float64, [5]process)  {


	sum := 0
	for i:=0 ; i<5 ; i++ {

		for j:= len(chart)- 1 ; j>-1 ; j--{
			if chart[j] == processes[i].id{
				turn := j - processes[i].firstArriveTime
				processes[i].turnaroundTime = turn + 1
				sum = processes[i].turnaroundTime + sum
				break
			}
		}
	}
	averageTurnAroundTime := float64(sum) / 5.0
	return averageTurnAroundTime , processes


}

func waitingTime(processes [5] process  ) (float64, [5]process)  {

	// calculating waiting time
	sum := 0
	for i:=0 ; i<5 ; i++ {
		sth := processes[i].cpuTime1 + processes[i].cpuTime2
		mines := processes[i].turnaroundTime - sth
		processes[i].waitingTime = mines
		sum = processes[i].waitingTime + sum
	}
	averageWaitingTime := float64(sum) / 5.0
	return averageWaitingTime , processes



}

func main() {

	records := readCsvFile("proces_inputs.csv")
	var sth [5][5] int
	var processes [5] process

	for i:=0; i<5; i++{
		for j:=0 ; j<5 ; j++{
			sth[i][j], _ = strconv.Atoi(records[i+1][j])
		}
		processes[i] = process{
			id :  sth[i][0],
			arriveTime : sth[i][1],
			firstArriveTime: sth[i][1] ,
			arrivedState: false,
			cpuTime1 : sth[i][2],

			cpuTime1State: false,
			ioTime : sth[i][3],
			ioTimeState: false,
			cpuTime2 : sth[i][4],
			cpuTime2State: false,
			isDone: false,
			lastUsed: false,
		}

	}

	FCFS(processes)
	RR(processes)
	SJF(processes)




}
