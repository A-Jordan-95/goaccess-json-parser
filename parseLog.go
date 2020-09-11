package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Analytics struct {
	Regions        [6]Region
	OpSystems      [9]OpSystem
	UniqueVisitors float64
}

type Region struct {
	Name       string
	VisCount   float64
	VisPercent float64
}

type OpSystem struct {
	Name       string
	VisCount   float64
	VisPercent float64
}

func getVisitorsAndData(theMap map[string]interface{}) (string, map[string]interface{}) {
	var data string
	var visitors map[string]interface{}
	for k, v := range theMap {
		switch k {
		case "visitors":
			visitors = v.(map[string]interface{})
		case "data":
			data = v.(string)
		}
	}
	return data, visitors
}

func getVisCountAndPercent(theMap map[string]interface{}) (float64, float64) {
	var count float64
	var percent float64
	for k, v := range theMap {
		switch k {
		case "count":
			count = v.(float64)
		case "percent":
			percent = v.(float64)
		}
	}
	return count, percent
}

func main() {
	//open json file:
	jsonFile, err := os.Open("report.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened report.json")
	//defer closing of json file:
	defer jsonFile.Close()

	//package json blob as bytearray:
	byteValue, _ := ioutil.ReadAll(jsonFile)

	//parse json blob into an interface:
	var result interface{}
	json.Unmarshal(byteValue, &result)

	//map the interfaced json blob:
	m := result.(map[string]interface{})

	//extract unique visitors information from mapped json blob:
	uniqueVisitors := m["general"].(map[string]interface{})["unique_visitors"].(float64)

	//map geolocation data from mapped json blob:
	geolocationData := m["geolocation"].(map[string]interface{})["data"]

	//map each region of geolocation data:
	regionMap0 := geolocationData.([]interface{})[0].(map[string]interface{})
	regionMap1 := geolocationData.([]interface{})[1].(map[string]interface{})
	regionMap2 := geolocationData.([]interface{})[2].(map[string]interface{})
	regionMap3 := geolocationData.([]interface{})[3].(map[string]interface{})
	regionMap4 := geolocationData.([]interface{})[4].(map[string]interface{})
	regionMap5 := geolocationData.([]interface{})[5].(map[string]interface{})

	//extract relevant analytical data from each region into a Region struct:
	region0Data, region0Visitors := getVisitorsAndData(regionMap0)
	region0VisCount, region0VisPercent := getVisCountAndPercent(region0Visitors)
	r0 := Region{region0Data, region0VisCount, region0VisPercent}

	region1Data, region1Visitors := getVisitorsAndData(regionMap1)
	region1VisCount, region1VisPercent := getVisCountAndPercent(region1Visitors)
	r1 := Region{region1Data, region1VisCount, region1VisPercent}

	region2Data, region2Visitors := getVisitorsAndData(regionMap2)
	region2VisCount, region2VisPercent := getVisCountAndPercent(region2Visitors)
	r2 := Region{region2Data, region2VisCount, region2VisPercent}

	region3Data, region3Visitors := getVisitorsAndData(regionMap3)
	region3VisCount, region3VisPercent := getVisCountAndPercent(region3Visitors)
	r3 := Region{region3Data, region3VisCount, region3VisPercent}

	region4Data, region4Visitors := getVisitorsAndData(regionMap4)
	region4VisCount, region4VisPercent := getVisCountAndPercent(region4Visitors)
	r4 := Region{region4Data, region4VisCount, region4VisPercent}

	region5Data, region5Visitors := getVisitorsAndData(regionMap5)
	region5VisCount, region5VisPercent := getVisCountAndPercent(region5Visitors)
	r5 := Region{region5Data, region5VisCount, region5VisPercent}

	//add each Region struct instantiation into an array of Regions
	regions := [6]Region{r0, r1, r2, r3, r4, r5}

	//map OpSys data from mapped json blob:
	osData := m["os"].(map[string]interface{})["data"]

	//map each OpSys of operating system data:
	osMap0 := osData.([]interface{})[0].(map[string]interface{})
	osMap1 := osData.([]interface{})[1].(map[string]interface{})
	osMap2 := osData.([]interface{})[2].(map[string]interface{})
	osMap3 := osData.([]interface{})[3].(map[string]interface{})
	osMap4 := osData.([]interface{})[4].(map[string]interface{})
	osMap5 := osData.([]interface{})[5].(map[string]interface{})
	osMap6 := osData.([]interface{})[6].(map[string]interface{})
	osMap7 := osData.([]interface{})[7].(map[string]interface{})
	osMap8 := osData.([]interface{})[8].(map[string]interface{})

	//extract relevant analytical data from each operating system into
	//an OpSystem struct:
	os0Data, os0Visitors := getVisitorsAndData(osMap0)
	os0VisCount, os0VisPercent := getVisCountAndPercent(os0Visitors)
	os0 := OpSystem{os0Data, os0VisCount, os0VisPercent}

	os1Data, os1Visitors := getVisitorsAndData(osMap1)
	os1VisCount, os1VisPercent := getVisCountAndPercent(os1Visitors)
	os1 := OpSystem{os1Data, os1VisCount, os1VisPercent}

	os2Data, os2Visitors := getVisitorsAndData(osMap2)
	os2VisCount, os2VisPercent := getVisCountAndPercent(os2Visitors)
	os2 := OpSystem{os2Data, os2VisCount, os2VisPercent}

	os3Data, os3Visitors := getVisitorsAndData(osMap3)
	os3VisCount, os3VisPercent := getVisCountAndPercent(os3Visitors)
	os3 := OpSystem{os3Data, os3VisCount, os3VisPercent}

	os4Data, os4Visitors := getVisitorsAndData(osMap4)
	os4VisCount, os4VisPercent := getVisCountAndPercent(os4Visitors)
	os4 := OpSystem{os4Data, os4VisCount, os4VisPercent}

	os5Data, os5Visitors := getVisitorsAndData(osMap5)
	os5VisCount, os5VisPercent := getVisCountAndPercent(os5Visitors)
	os5 := OpSystem{os5Data, os5VisCount, os5VisPercent}

	os6Data, os6Visitors := getVisitorsAndData(osMap6)
	os6VisCount, os6VisPercent := getVisCountAndPercent(os6Visitors)
	os6 := OpSystem{os6Data, os6VisCount, os6VisPercent}

	os7Data, os7Visitors := getVisitorsAndData(osMap7)
	os7VisCount, os7VisPercent := getVisCountAndPercent(os7Visitors)
	os7 := OpSystem{os7Data, os7VisCount, os7VisPercent}

	os8Data, os8Visitors := getVisitorsAndData(osMap8)
	os8VisCount, os8VisPercent := getVisCountAndPercent(os8Visitors)
	os8 := OpSystem{os8Data, os8VisCount, os8VisPercent}

	//add each OpSystem struct into an array of OpSystems:
	opSystems := [9]OpSystem{os0, os1, os2, os3, os4, os5, os6, os7, os8}

	//package unique visitors, regions, and operating systems into
	//a single 'Analytics' struct:
	analytics := Analytics{regions, opSystems, uniqueVisitors}

	//marshall 'Analytics' struct into much smaller more useful json blob
	//than we originally started with:
	file, _ := json.MarshalIndent(analytics, "", " ")
	//write the json blob to a file which we can add to database for use later:
	_ = ioutil.WriteFile("analytics.json", file, 0644)
}
