package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/puerkitobio/goquery"
)

const (
	POSTDATA    = "&ctl00%24cphMain%24TabContainer1%24Marks%24ddlYear=2014&ctl00%24cphMain%24TabContainer1%24Marks%24btnFind=Find"
	KNEC_URL    = "http://www.knec-portal.ac.ke/RESULTS/ResultKCPE.aspx"
	MAXROUTINES = 1000
)

/*given an index number return a candidates results in a html page */
type PageResult struct {
	Page  string
	Index string
}

func getCandidateResults(index string, client *http.Client) (htmlPage string, err error) {

	data := getPreData() + index + POSTDATA
	bb := bytes.NewBuffer([]byte(data))
	req, err := http.NewRequest("POST", KNEC_URL, bb)
	if err != nil {
		return "", errors.New("Couldn't make post request")
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	if err != nil {
		return "", errors.New("Making request error")
	}

	restr, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", errors.New("Body parse error")
	}
	return string(restr), nil
}

func getIndexNumbers() chan string {

	ch := make(chan string) // for generator

	COUNTY_NUMBERS := []string{
		"01101", "01113", "01114", "01115", "02105", "02109", "02110", "03106", "03108",
		"03120", "03121", "04102", "04107", "04111", "04116", "04119", "04122", "05103",
		"05112", "05117", "06104", "06118", "07201", "07209", "07213", "07214", "07215",
		"07216", "07225", "08202", "08210", "08217", "08218", "08219", "08220", "08221",
		"08237", "09203", "09222", "09223", "09224", "09239", "10204", "10208", "10225",
		"10227", "10228", "10229", "10234", "10238", "11205", "11207", "11211", "11212",
		"11230", "11231", "11232", "11233", "11235", "11235", "12301", "12314", "12315",
		"12315", "12329", "12330", "12343", "12345", "13302", "13310", "13317", "13328",
		"13331", "13332", "13338", "13339", "13344", "13350", "13351", "13352", "13353",
		"13354", "13357", "13350", "14303", "14312", "14333", "14341", "14355", "15304",
		"15309", "15318", "15319", "15327", "15334", "15337", "15349", "15305", "15311",
		"15320", "15321", "15340", "16358", "15359", "17305", "17322", "17355", "18307",
		"18323", "18324", "18325", "18335", "18336", "18346", "18347", "18348", "19308",
		"19313", "19325", "19342", "20401", "20402", "20403", "20404", "20405", "20405",
		"20407", "20408", "20409", "21501", "21524", "21525", "21548", "21549", "21550",
		"22502", "22525", "22527", "23503", "23528", "23529", "24505", "24530", "24531",
		"24568", "25508", "25533", "25551", "25553", "25509", "25534", "25535", "27511",
		"27535", "27537", "27538", "27552", "27554", "27555", "27559", "27570", "28512",
		"28522", "28539", "28553", "28571", "29513", "29523", "29540", "29541", "29542",
		"30514", "30543", "30544", "30555", "30555", "31515", "31545", "31557", "31566",
		"31567", "32516", "32519", "32546", "32560", "33517", "33521", "33532", "33547",
		"33558", "33562", "34518", "34520", "34559", "34561", "35601", "35606", "35609",
		"35610", "35620", "35623", "35629", "36602", "36605", "36611", "36612", "36613",
		"36621", "36626", "36628", "36630", "37603", "37607", "37608", "37614", "37615",
		"37616", "37617", "37624", "37625", "37627", "37631", "37532", "38504", "38618",
		"38519", "38622", "39701", "39713", "39714", "39733", "39734", "39737", "40703",
		"40711", "40715", "40719", "40723", "40727", "40732", "40735", "40740", "41704",
		"41709", "41710", "41724", "41730", "41731", "42705", "42712", "42721", "42725",
		"42725", "42738", "43705", "43715", "43720", "43722", "43728", "44707", "44708",
		"44717", "44718", "44729", "44736", "44739", "45301", "45304", "45305", "45806",
		"45815", "45816", "45821", "46802", "46807", "46308", "46809", "46813", "46818",
		"46819", "46320", "47803", "47810", "47311", "47812", "47814", "47817"}

	go func() {
		for _, countyIndex := range COUNTY_NUMBERS {
			// loop through all possible schools
			for i := 1; i < 1000; i++ {
				//format this in a way where we can fast forward to a new school
				var schoolIndex string = "000" + strconv.Itoa(i)
				schoolIndex = schoolIndex[(len(schoolIndex) - 3):]
				// for each school, loop through every possible student
				for s := 1; s < 1000; s++ {
					var candidateIndex string = "000" + strconv.Itoa(s)
					candidateIndex = candidateIndex[(len(candidateIndex) - 3):]
					ch <- (countyIndex + schoolIndex + candidateIndex)
				}
			}
		}
	}()

	return ch
}

/* Parse html page and give  results in a CSV format */

func getPreData() (predat string) {

	predata := "ctl00_cphMain_ToolkitScriptManager1_HiddenField=%3B%3BAjaxControlToolkit"
	predata += "%2C+Version%3D3.5.40412.0%2C+Culture%3Dneutral%2C+PublicKeyToken%3D28f01b0e84b6d53e"
	predata += "%3Aen-GB%3A1547e793-5b7e-48fe-8490-03a375b13a"
	predata += "33%3A475a4ef5%3Aeffe2a26%3A8e94f951%3A1d3ed089&ctl00_cphMain_TabContainer1_ClientState=%"
	predata += "7B%22ActiveTabIndex%22%3A0%2C%22TabState%22%3A%5Btrue%5D%7D&__EVENTTARGET=&__EVENTARGUMENT=&__VIEWSTATE"
	predata += "=%2FwEPDwUJNzIwODUwMzE2D2QWAmYPZBYCAgQPZBYEAgEPDxYCHgRUZXh0BQ8yMSBKYW51YXJ5IDIwMTVkZAILD2QWAgIBD2QWAmYPZBY"
	predata += "CAgEPZBYCAhsPPCsADQBkGAMFHl9fQ29udHJvbHNSZXF1aXJlUG9zdEJhY2tLZXlfXxYBBRtjdGwwMCRjcGhNYWluJFRhYkNvbnRhaW5lcjEFG2N0bD"
	predata += "AwJGNwaE1haW4kVGFiQ29udGFpbmVyMQ8PZGZkBStjdGwwMCRjcGhNYWluJFRhYkNvbnRhaW5lcjEkTWFya3MkR3JpZHZpZXcxD2dkCcktLvt%2FXIMC"
	predata += "IziOu3%2Bqi8MlsNs%3D&__VIEWSTATEGENERATOR=8A3A71A8&__EVENTVALIDATION=%2FwEWBAKgps3rDQLS1aG9AgKX%2F9f8CwK2lIrpBn45EiR"
	predata += "ZIqaQ%2FpmQ8A5Ae9qmxnMY&ctl00%24cphMain%24TabContainer1%24Marks%24txtIndex="

	return predata

}

func parsePage(pageRes *PageResult) (stud map[string]string, err error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBufferString(pageRes.Page))
	student := make(map[string]string)
	if err != nil {
		return student, err
	}

	fields := map[string]string{"total": "#ctl00_cphMain_TabContainer1_Marks_txtTotal",
		"name":       "#ctl00_cphMain_TabContainer1_Marks_txtName",
		"eng":        "#ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl02_MKS",
		"kis":        "#ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl03_MKS",
		"mat":        "#ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl04_MKS",
		"sci":        "#ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl05_MKS",
		"ssr":        "#ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl06_MKS",
		"schoolName": "#ctl00_cphMain_TabContainer1_Marks_txtSchool",
		"gender":     "#ctl00_cphMain_TabContainer1_Marks_txtGender",
		"engGrade":   "#ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl02_GRADE",
		"kisGrade":   "#ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl03_GRADE",
		"matGrade":   "#ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl04_GRADE",
		"sciGrade":   "#ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl05_GRADE",
		"ssrGrade":   "#ctl00_cphMain_TabContainer1_Marks_Gridview1_ctl06_GRADE",
	}
	for subj, param := range fields {

		f, ok := doc.Find(param).Attr("value")
		if !ok {
			return student, errors.New("Bad page")
		}
		student[subj] = f
	}
	student["index"] = pageRes.Index
	return student, nil
}

func main() {

	tasks := getIndexNumbers()
	results := make(chan string)

	var wg sync.WaitGroup
	for i := 0; i < MAXROUTINES; i++ {
		wg.Add(1)
		go func() {
			for index := range tasks {
				res, err := getCandidateResults(index, client)
				if err != nil {

					fmt.Println("Error with the connection")

				}
				p := &PageResult{Page: res, Index: index}
				results <- p

			}
			wg.Done()
		}()

	}

	wg.Wait()

	//get html from results channel and parse/ put them in CSV

	/*/FOR TESTING
	client := &http.Client{}
	index := "01113118001"
	res, _ := getCandidateResults(index, client)
	p := &PageResult{Page: res, Index: index}
	fmt.Println(parsePage(p))*/

}
