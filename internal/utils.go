package internal

import (
	"bufio"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"sort"
	"time"
)

func SortByResponseTime(validProxy []ValidProxy) []ValidProxy {

	sort.Slice(validProxy, func(i, j int) bool {
		return validProxy[i].ResponseTime < validProxy[j].ResponseTime
	})

	return validProxy

}

// AppendTo append string to a file
func AppendTo(filename string, data string) (string, error) {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}
	if _, err := f.Write([]byte(data + "\n")); err != nil {
		return "", err
	}
	if err := f.Close(); err != nil {
		return "", err
	}
	return filename, nil
}

func ReadFile(path string) ([]string, error) {
	inFile, err := os.Open(path)
	var result []string
	if err != nil {
		fmt.Println(err.Error() + `: ` + path)
		return nil, err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		//fmt.Println(scanner.Text()) // the line
		result = append(result, scanner.Text())
	}

	return result, nil
}

func WriteProxiesToFile(proxies []ValidProxy, output string) {

	file, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	if err != nil {
		log.Fatal().Err(err).Msg("failed creating file")
	}

	lineWriter := bufio.NewWriter(file)

	for _, proxy := range proxies {
		_, _ = lineWriter.WriteString(proxy.Address + "\n")
	}

	lineWriter.Flush()

}

func Unique(input []string) []string {
	unique := make(map[string]bool, len(input))
	list := make([]string, len(unique))
	for _, el := range input {
		if len(el) != 0 {
			if !unique[el] {
				list = append(list, el)
				unique[el] = true
			}
		}
	}
	return list
}

func GenerateOutputName(output string) string {

	t := time.Now()
	result := fmt.Sprintf("%d-%02d-%02dT%02d-%02d-%02d-%s",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(), output)

	return result
}

//
//func GetCurrentIP()  string {
//
//	resp, err :=  http.Get("https://api.ipify.org/?format=text")
//
//	if err!=nil{
//
//		log.Fatal().Err(err).Msg("Can't get current ip")
//	}
//	defer resp.Body.Close()
//
//	body , _ := ioutil.ReadAll(resp.Body)
//
//	return string(body)
//
//}
