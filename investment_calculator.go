package main
import "math"
import "fmt"
func main() {
    // investmentAmount,year :=2000,10
    var investmentAmount float64
    var year float64
    expetcedReturnRate:=5.5
    const inflationRate =2.5
    
    fmt.Print("Investment amount: ")
    fmt.Scan(&investmentAmount)
    fmt.Print("Expected return rate ")
    fmt.Scan(&expetcedReturnRate)
    fmt.Print("Years: ")
    fmt.Scan(&year)

    furtureValue:=float64(investmentAmount) *math.Pow(1+expetcedReturnRate /100,float64(year))
    futureRealValue := furtureValue/math.Pow(1+inflationRate/100,year)
    fmt.Print("furtureValue")
    fmt.Println(furtureValue)
    fmt.Print("futureRealValue")
    fmt.Println(futureRealValue)
}