package main

import (
	"fmt"
	"time"
	"os"
	"bufio"
	"strings"
)
func main() {
	fmt.Printf("Give an adress for the server:\n")
	reader := bufio.NewReader(os.Stdin)
    r, _ := reader.ReadString('\n')
	r = strings.Trim(r, "\n")
	dir, size , err := dfs(r,r)
	if err != nil {
		fmt.Println("Failed to access server is down.") //since is never specified what happens when you can't even get into the server to begin with i'll just end the program there
		return
	}
	ss, err := os.Create("successful.txt")
	if err != nil {
		panic(err)
	}
	ff, err := os.Create("not-successful.txt")
	if err != nil {
		panic(err)
	}
	//time.Sleep(10*time.Second) //toggle this to test the "disconnection" system :D
 	for i := 0; i<size; i++ {
		err := dw(dir[i])
		start := time.Now()
		att := 1
		for ; time.Since(start) < 10*time.Minute; { //fix to 10 mins after done testing
			if err != nil {
				if att%12 == 0 {
					time.Sleep(time.Minute)
				} else {
					time.Sleep(5*time.Second)
				}
				att++
				err = dw(dir[i])
			} else {
				break
			}
		}
		if err != nil {
			_, err = ff.WriteString(dir[i] + "\n")
			if err != nil {
				panic(err)
			}
		} else {
			_, err = ss.WriteString(dir[i] + "\n")
			if err != nil {
				panic(err)
			}
		}
	}
	defer ss.Close()
	defer ff.Close()
}