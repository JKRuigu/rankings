package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/puerkitobio/goquery"
)

const (
	POSTDATA     = "&ctl00%24cphMain%24TabContainer1%24Marks%24ddlYear=2014&ctl00%24cphMain%24TabContainer1%24Marks%24btnFind=Find"
	KNEC_URL     = "http://www.knec-portal.ac.ke/RESULTS/ResultKCPE.aspx"
	STUDENTERROR = 5
	DEBUG        = true
)

/*given an index number return a candidates results in a html page */
type PageResult struct {
	Page  string
	Index string
}

func debug(stuff string) {
	if DEBUG {
		fmt.Println(string)
	}
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

func getCountyNumbers() chan string {

	ch := make(chan string) // for generator

	COUNTY_NUMBERS := []string{"01113"}
	/*COUNTY_NUMBERS := []string{"01101"} , "01113", "01114", "01115", "02105", "02109", "02110", "03106", "03108",
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
	"46819", "46320", "47803", "47810", "47311", "47812", "47814", "47817"}*/

	go func() {
		for _, v := range COUNTY_NUMBERS {
			ch <- v
		}
		close(ch)
	}()

	return ch

}

func getStudentDetails(countyIndex string, client *http.Client) (students chan map[string]string) {

	ch := make(chan map[string]string)

	go func() {

		// loop through all possible schools
		var wg sync.WaitGroup
		for i := 117; i < 130; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				var errCount uint64 = 0
				fmt.Println("New School")
				var schoolIndex string = "000" + strconv.Itoa(i)
				schoolIndex = schoolIndex[(len(schoolIndex) - 3):]
				candidates := genCandidateIndex()

				for j := 0; j < len(candidates); j++ {

					for n := 0; n < len(candidates[j]); n++ {

						candidateIndex := "000" + strconv.Itoa(candidates[j][n])
						candidateIndex = candidateIndex[(len(candidateIndex) - 3):]
						fmt.Println("waiting")
						stud := countyIndex + schoolIndex + string(candidateIndex)
						fmt.Println(stud)
						res, err := getCandidateResults(stud, client)
						if err != nil {
							log.Fatal("Error with the connection")

						}
						//fmt.Println("here 1")
						p := &PageResult{Page: res, Index: stud}

						student, err := parsePage(p)
						//fmt.Println("here 2")
						fmt.Println(parsePage(p))

						if err != nil {
							fmt.Println("error")
							errCount += 1
							if errCount%STUDENTERROR == 0 {
								break
								fmt.Println("Ma bad")

							}

						} else {
							ch <- student

						}

					}
				}
			}(i)
		}
		wg.Wait()
		close(ch)

	}()

	return ch
}

func genCandidateIndex() map[int][]int {

	//res := make(chan string)
	candidates := make(map[int][]int)

	candidates[0] = append(candidates[0], 1)
	candidates[0] = append(candidates[0], 2)
	candidates[0] = append(candidates[0], 3)
	candidates[1] = append(candidates[1], 11)
	candidates[1] = append(candidates[1], 12)
	candidates[1] = append(candidates[1], 13)
	candidates[2] = append(candidates[2], 100)
	candidates[2] = append(candidates[2], 101)
	candidates[2] = append(candidates[2], 102)
	/*for i := 0; i < 300; i++ {
		candidates[0] = append(candidates[0], i+1)
	}
	for i := 300; i < 600; i++ {
		candidates[1] = append(candidates[1], i+1)
	}
	for i := 700; i < 800; i++ {
		candidates[2] = append(candidates[1], i+1)
	}
	for i := 800; i < 900; i++ {
		candidates[3] = append(candidates[1], i+1)
	}*/
	/*go func() {
	Cands:
		for _, val := range candidates {

			for _, index := range val {
				fmt.Println("moving waiting")
				yes, ok := <-move
				if !ok {
					break Cands
				}
				fmt.Println("moving", yes)
				if yes {
					continue Cands
				}
				candidateIndex := "000" + strconv.Itoa(index)
				candidateIndex = candidateIndex[(len(candidateIndex) - 3):]
				res <- candidateIndex
			}

		}
		/*
			res <- "001"
			res <- "002"
		close(res)
	}()*/

	return candidates

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

	fields := map[string]string{
		"total":      "#ctl00_cphMain_TabContainer1_Marks_txtTotal",
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
		student[subj] = strings.TrimSpace(f)
	}
	student["index"] = pageRes.Index
	return student, nil
}

func main() {

	client := &http.Client{}
	counties := getCountyNumbers()
	fields := []string{"total", "name", "eng",
		"kis", "mat", "sci", "ssr",
		"schoolName", "gender", "engGrade", "kisGrade",
		"matGrade", "sciGrade", "ssrGrade"}

	fmt.Println("total,name,eng,kis,mat,sci,ssr,schoolName,gender,engGrade,kisGrade,matGrade,sciGrade,ssrGrade")
	var wg sync.WaitGroup
	for countyIndex := range counties {
		wg.Add(1)
		debug("Country index(main): " + countyIndex)
		//start a routine per county
		go func(countyIndex string) {
			defer wg.Done()

			//indexes for all candidates in the county
			students := getStudentDetails(countyIndex, client)
			//print out students
			for student := range students {
				out := ""
				for i, details := range fields {

					if i < len(fields)-1 {
						out += student[details] + ","
					} else {

						out += student[details]
					}
				}

				fmt.Println(out)
			}

		}(countyIndex)

	}

	wg.Wait()

}
