// package main

// import "fmt"

// func main(){
// 	fmt.Println("hello worls")
// }

// package main

// import "fmt"

// func main(){
// 	 a,b := "10","15";
// 	fmt.Printf("%s %s",a,b);
// }

// package main

// import "fmt"

// func main(){
// 	 a := 11;
// 	 if (a % 2 == 0){
// 	fmt.Printf("%d even",a )
// 	 }	 else {
// 		fmt.Printf("%d odd",a )
// 	 }
// }

// package main
// import "fmt"

// func main(){
// 	var num int
// 	fmt.Printf("Enter the number :")
// 	fmt.Scanf("%d",&num)
// 	if num%2==0 {
// 		fmt.Printf("%d is even",num)
// 	}else {
// 		fmt.Printf("%d is odd",num)

// 	}
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main(){
// 	var num int
// 	fmt.Printf("Enter the number : ")
// 	fmt.Scanf("%d",&num)
// 	var data,data1 string =FizzBuzz(num);
// 	fmt.Printf("%s %s",data,data1)

// }

// func FizzBuzz(n int) (string,string) {
// 	if n%3==0 && n%5==0{
// 		return "FizzBuzz","FizzBuzz"
// 	}else if n%3==0{
// 		return "Fizz","FizzBuzz"
// 	}else if n%5==0{
// 		return "Buzz","FizzBuzz"
// 	}
// 	return strconv.Itoa(n),"FizzBuzz"
// }

// package main

// import "fmt"

// func main(){
// 	var c byte
// 	fmt.Printf(" do subscription ")
// 	fmt.Scanf("%c",&c)
// 	switch(c){
// 	case 'y':fmt.Println("Thanku")
// 	case 'n':fmt.Println("no problem")
// 	default:fmt.Println("invalid")
// 	}
// }
// package main

// import "fmt"

// func main() {
// 	v := "42" // change me!
// 	fmt.Printf("%T", v)
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func Sqrt(x float64) float64 {
// 	z := 1.0

// 	for i := 0; i < 10; i++ {
// 		z -= (z*z - x) / (2 * z)
// 		fmt.Printf("Iteration %d: z = %v\n", i+1, z)
// 	}

// 	return z
// }

// func main() {
// 	values := []float64{1, 2, 3, 4, 5}

// 	for _, x := range values {
// 		result := Sqrt(x)
// 		fmt.Printf("Sqrt(%v) = %v (math.Sqrt: %v)\n", x, result, math.Sqrt(x))
// 	}
// }

// package main

// import "fmt"
// func main() {
// 	var daysOfWeek[] string
// 	daysOfWeek = make([]string,3,4)
// 	daysOfWeek =append(daysOfWeek,"1", "2", "3", "4", "5", "6",);
// 	// weekdays_:= daysOfWeek[:]
// 	print(daysOfWeek)
// }
// func print(slice []string) {
// 	fmt.Printf("%v len:%d cap:%d\n", slice,len(slice),cap(slice))
// }

// package main

// import "fmt"

// func main() {
// 	var a int
// 	var b *int
// 	a=5
// 	b=&a
// 	*b=10
// 	fmt.Printf("%d %d",a,*b)
// }

// package main

// import "fmt"

// func main(){

// 	var input string

// 	wheelsMap  :=map[string]int{
// 		"car":4,
// 		"auto":3,
// 		"bicycle":2,
// 	}
// fmt.Scanf("%s",&input)

// wheel :=wheelsMap[input]
// fmt.Printf("%v",wheel)
// 	// fmt.Println(wheelsMap);
// 	// for vehicle, wheels := range wheelsMap {
// 	// 	fmt.Printf("%s : %d\n", vehicle, wheels)
// // }
// }

// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main(){
// readFile(os.Args[1])
// }

// func readFile(name string){
// 	var bytes [10]byte
// 	var err error
// 	f,err := os.Open(name)

// 	if err != nil{
// 		panic(err)
// 	}

// 	for {
// 		_, err = f.Read(bytes[:])
// 		if err == io.EOF {
// 			break
// 	}
// 	fmt.Printf("%s:",bytes)
// 	}

// }

// package main

// import "fmt"

// type Celsius float64

// type Farenheit float64

// func main() {
// var f Farenheit
// f=125
// fmt.Println(f)
// }

// package main

// import "fmt"

// type Celsius float32

// type Farenheit float32

// func (c Celsius) ToFarenheit() Farenheit {
// return Farenheit(9.0/5.0*c+32)
// }

// func main() {
// var c Celsius
// c=36.9
// f:=c.ToFarenheit()
// fmt.Println(f)
// }

// package main

// import (
// 	"fmt"
// 	"os"
// 	"time"
// )

// type Logger interface{
// 	Info(string)
// }

// const format ="%v: Info : %s"
// type ScreenLogger struct{

// }

// func (ScreenLogger) Info(message string) {
// fmt.Printf(format,time.Now(),message)
// }

// type FileLogger struct{
// 	File os.File
// }

// func main(){
// 	var log Logger
// 	log = &ScreenLogger{}
// }

// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )
// func main() {
// 	fmt.Println("Welcom")

// 	content :="this is data"

// 	file,err := os.Create("./test.txt")

// 	checkNilError(err)

// 	length,err := io.WriteString(file,content)

// 	checkNilError(err)

// fmt.Println("length is ",length)
//  defer file.Close()
//  readFile("./test.txt")
// }

// func readFile(filename string){
// 	databyte,err:=os.ReadFile(filename)
// 	checkNilError(err)
// 		fmt.Println("text data\n",string(databyte))
// }

// func checkNilError(err error){
// 	if err != nil {
// 		panic(err)
// 		}
// }

// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// const url="http://lco.dev"

// func main() {
// 	fmt.Println("LCO web request")

// 	responce,err := http.Get(url)

// 	if err != nil {
// 	panic(err)
// 	}
// fmt.Printf("responceis of type:%T\n",responce)
// defer responce.Body.Close()

// databytes,err:=io.ReadAll(responce.Body)

// if err != nil {
// panic(err)
// }

// content:=string(databytes)
// fmt.Println(content)

// }

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main(){
	fmt.Println("wellcome")
	PerformGetRequest()
}

func PerformGetRequest(){

	const myurl = "http://localhost:8000/get"

	response,err :=http.Get(myurl)

if err!=nil{

	panic(err)

}

defer response.Body.Close()

fmt.Println("staus code", response.Status)

content,_:=io.ReadAll(response.Body)

fmt.Println(string(content))

}

