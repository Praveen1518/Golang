
package main
import "fmt"
func main(){
	var revenue float64
	var expenses float64
	var taxRate float64
	fmt.Print("revenue ")
	fmt.Scan(&revenue)
	fmt.Print("expenses ")
	fmt.Scan(&expenses)
	fmt.Print("taxRate ")
	fmt.Scan(&taxRate)
	ebt:= revenue-expenses
	profit:= ebt *(1-taxRate/100)
	ratio:=ebt/profit
	// fmt.Println("ebt",ebt)
	fmt.Printf("ebt %v \n",ebt)
	fmt.Println("profit",profit)
	fmt.Println("ratio",ratio)
}