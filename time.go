package main

import (
	"bufio"
	"fmt"

	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	//"time"
)

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
}

func PlayList(device Player, song []string) {
	for _,song := range song{
		device.Play(song)
	}
	device.Stop()

}

func OpenFile(fileName string) (*os.File, error) {
	fmt.Println("Opening", fileName)
	return os.Open(fileName)
}

func CloseFile(file *os.File) {
	fmt.Println("Closing File")
}

func abc(channel chan string) { 
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func GetFloats(fileName string) ([]float64,error) {
	var numbers []float64
	file, err := OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	defer CloseFile(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	
	if scanner.Err() != nil{
		return nil, scanner.Err()
	}
	return numbers, nil
}

type Page struct {
	URL string
	Size int
}


func main() {
	// numbers,err := GetFloats(os.Args[1])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var sum float64 = 0
	// for _, number := range numbers{
	// 	sum += number
	// }
	// fmt.Printf("Sum: %0.2f\n", sum)
	pages := make(chan Page)
	urls := []string{"https://example.com/","https://golang.org/","https://golang.org/doc"}
	for _, url := range urls{
		go responseSize(url, pages)
	}
	for i :=0; i < len(urls); i++{
		fmt.Println(<- pages)
	}
	
	
}

func responseSize(url string, channel chan Page) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)}
	// myChanel := make(chan string)
	// go send(myChanel)
	// reportNap("Receiving gorutine", 5)
	// fmt.Println(<-myChanel)
	// fmt.Println(<-myChanel)
	
// }

// func reportNap(name string, delay int) {
// 	for i := 0; i < delay; i++ {
// 		fmt.Println(name, "sleeping")
// 		time.Sleep(1 * time.Second)
// 	}
// 	fmt.Println(name,"Wakes up")
// }

// func send(myChanel chan string) {
// 	reportNap("sending gorutine", 2)
// 	fmt.Println("***Sending value***")
// 	myChanel <- "a"
// 	fmt.Println("***Sending value***")
// 	myChanel <- "b"
}
