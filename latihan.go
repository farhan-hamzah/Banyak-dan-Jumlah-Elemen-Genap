package main
import "fmt"

func main(){
	const NMAX int = 10
	var A[NMAX]int
	var i, n, jumBilGenap, banyakElemen int
	fmt.Scan(&n)

	for i= 0; i < n; i++{
		fmt.Scan(&A[i])
	}
	
	for i = 0; i < n; i++{
		if A[i]%2 == 0{
			jumBilGenap+=1
			banyakElemen +=A[i]
		}
	}
	fmt.Print(jumBilGenap, banyakElemen)
}