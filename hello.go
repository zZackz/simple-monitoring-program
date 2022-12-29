package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	showIntroduction()

	for {
		showMenu()
		command := setCommand()
		switch command {
		case 1:
			initMonitoring()
		case 2:
			showLog()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid Command")
		}
	}

}

func showIntroduction() {
	name := "Zack"
	version := 1.1
	fmt.Println("Hell,", name)
	fmt.Println("Version program: ", version)
}

func showMenu() {
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit")
}

func setCommand() int {
	var command int

	fmt.Scan(&command)

	return command
}

func initMonitoring() {
	fmt.Println("Monitoring...")
	// websites := []string{"https://www.alura.com.br", "https://www.google.com", "https://www.facebook.com"}

	websites := readFileWebsite()

	for {

		for _, website := range websites {
			resp, err := http.Get(website)

			if err != nil {
				fmt.Println(err)
			}

			if resp.StatusCode == 200 {
				fmt.Println("Website: ", website, "successfully loaded")
				registerLog(website, true)
			} else {
				fmt.Println("Website: ", website, "loading failure")
				registerLog(website, false)
			}
		}

		time.Sleep(5 * time.Second)
	}

}

func readFileWebsite() []string {

	var websites []string

	file, err := os.Open("sites.txt")
	// file, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		websites = append(websites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return websites

}

func registerLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05 ") + site + "- online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func showLog() {

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))

}
