package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	returnSitesFromTxt()
	for {
		showMenu()
		selectItem()
	}
}

func showMenu() {
	fmt.Println("\n-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
	fmt.Println("1- Monitore a site")
	fmt.Println("2- Monitore all sites from targets.txt")
	fmt.Println("3- Show Logs")
	fmt.Println("0- Exit")
	fmt.Print("Select an option: ")
}

func selectItem() {
	var command int
	fmt.Scan(&command)

	switch command {
	case 0:
		fmt.Println("Bye Bye!")
		os.Exit(0)
	case 1:
		checkSite()
	case 2:
		checkSiteAtFile()
	case 3:
		showLogs()
	default:
		fmt.Println("Unknown command.")
		os.Exit(-1)
	}
}

func checkSite() {
	var url string
	fmt.Print("Paste the all url: ")
	fmt.Scan(&url)

	resp, err := http.Get(url)
	registerLog(url, resp.StatusCode)

	if err != nil || resp.StatusCode != 200 {
		fmt.Println(url, "is uncessable.")
		return
	}
	fmt.Println(url, "is ok.")
}

func checkSiteAtFile() {
	sites := returnSitesFromTxt()

	for _, url := range sites {
		resp, err := http.Get(url)

		if err != nil || resp.StatusCode != 200 {
			fmt.Println("[ ??? ]", url, "is uncessable.")
			registerLog(url, 400)
		} else {
			fmt.Println("[", resp.StatusCode, "]", url, "is ok.")
			registerLog(url, resp.StatusCode)
		}
	}
}

func returnSitesFromTxt() (sites []string) {
	file, err := os.Open(("pkg/targets.txt"))
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println("File 'pkg/targets.txt' was not found.")
		os.Exit(-1)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		sites = append(sites, strings.TrimSpace(line))
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println("File are not valid.")
			os.Exit(-1)
		}
	}

	file.Close()
	return
}

func registerLog(site string, status any) {
	file, err := os.OpenFile("pkg/logs.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println("File 'pkg/logs.txt' was not found.")
		os.Exit(-1)
	}

	if status == 200 {
		file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " -> [ 200 ] " + site + " is ok!\n")
	} else {
		file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " -> [ ??? ] " + site + " IS DOWN\n")
	}

	file.Close()
}

func showLogs() {
	file, err := os.Open("pkg/logs.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println("File 'pkg/logs.txt' was not found.")
		os.Exit(-1)
	}

	fmt.Println("\n-=-=-=-=-=-=-=-=-=Logs-=-=-=-=-=-=-=-=-=")
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		fmt.Println(strings.TrimSpace(line))
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println("File are not valid.")
			os.Exit(-1)
		}
	}
	fmt.Println("\n-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	file.Close()
}
