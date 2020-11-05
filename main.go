package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	data, err := ioutil.ReadFile("emails.txt")
	if err != nil {
		fmt.Println(`File "emails.txt" not found`, err)
		time.Sleep(5 * time.Second)
		return
	}
	re := regexp.MustCompile("[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,4}")

	count := len(re.FindAllIndex(data, -1))
	if count < 1 {
		fmt.Println("Not found email.")
		time.Sleep(5 * time.Second)
		panic("end")
	}
	emails := re.FindAll(data, -1)

	domainRegex := regexp.MustCompile("@[A-Za-z0-9.-]+\\.[A-Za-z]{2,4}")
	domain := domainRegex.Find([]byte(data))

	f, err := os.Create("./" + string(domain) + ".txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := "Count: " + strconv.Itoa(count)
	for _, v := range emails {
		w += "\n" + string(v)
		fmt.Println(string(v))
	}

	b := w
	f.Write([]byte(b))
	time.Sleep(20 * time.Second)
	return
}