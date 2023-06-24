package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	var UID string = ""
	var Input string = ""
	fmt.Println("Enter App UID or full app store URL")
	fmt.Scanln(&Input)
	

	isURL := strings.HasPrefix(Input, "https://apps.apple.com/")
	if isURL == true {
		UID, cntry := extractAppURL(Input)
		output := extractKey(getBundleID(UID,cntry), "bundleId")
		fmt.Println(output)
	} else {
		UID = Input
		output := extractKey(getBundleID(UID,"US"), "bundleId")
		fmt.Println(output)
	}

}

func extractAppURL(URL string) (UID, cntry string){

	u := regexp.MustCompile(".*id")
	UID = u.ReplaceAllString(URL, "")
	cntry = strings.TrimPrefix(URL,"https://apps.apple.com/")
	c := regexp.MustCompile("/.*")
	cntry = c.ReplaceAllString(cntry, "")

	return

}
func getBundleID(UID string, cntry string) string{
	url := "https://itunes.apple.com/lookup?id="+UID+"&country="+cntry
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	output := (string(body))
	return output
}

func extractKey(body string, key string) string {
	keystr := "\"" + key + "\":[^,;\\]}]*"
	r, _ := regexp.Compile(keystr)
	match := r.FindString(body)
	keyValMatch := strings.Split(match, ":")
	return strings.ReplaceAll(keyValMatch[1], "\"", "")
}