package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
)

func main() {
	fmt.Print("Enter DOMAIN: ")
	var url string
	fmt.Scanln(&url)
	client := &http.Client{}
	lines, _ := readLines("wordlist.txt")
	fmt.Println(len(lines))
	fmt.Println("Starting...")
	var wg sync.WaitGroup
	for _, line := range lines {

		wg.Add(1)
		go requestq(line, url, client, &wg)

	}
	wg.Wait()
	fmt.Println("Done!")
}

func requestq(line string, url string, client *http.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	req, _ := http.NewRequest("GET", "http://"+line+"."+url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36")
	resp, err := client.Do(req)
	if err == nil {
		if resp.StatusCode < 400 {
			fmt.Println("Found: " + "http://" + line + "." + url + " Status: " + resp.Status)
		}
	} else {
	}
}

// I take the function from => https://stackoverflow.com/a/18479916/14779443
// I take the wordlist from => https://raw.githubusercontent.com/8sp/Subdomain-Finder/main/wordlist
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
