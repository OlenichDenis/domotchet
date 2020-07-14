package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const CityId string = "2019-11-12T06-53-35Z-1ba40c"
//const RegionId string = "2019-11-07T04-35-26Z-b54525"
const RoomsId string = "2019-11-07T04-45-07Z-ac97fe"
const TypeId string  = "2019-11-07T04-43-52Z-083e41"
const StreetId string  = "2019-11-14T06-14-09Z-8452ed"
const  DistrictId string = "2019-11-16T23-20-18Z-4888f2"
const  HouseTypeId string = "2019-11-07T04-52-08Z-de9759"


func GetSpr(tplID string) (result ResponseData){
	sp, err := http.Post("http://domotchet.ru/kvadrat/gui/query/query_source_link?source="+tplID, "application/json", nil)
	if err != nil {
		fmt.Println("ошибка запроса справочника", err)
		time.Sleep(10 * time.Minute)
		return
	}
	json.NewDecoder(sp.Body).Decode(&result)

	//fmt.Println(result, "sp")
	return result

	}

func ParsSpr(inpData ResponseData) map[string]string{
	var outMap = make(map[string]string)
	//fmt.Println(inpData, "inpData")

	if len(inpData.Data) == 0{
		//fmt.Println(inpData, "Пустое значение intDate")
		//return nil
	}
	//fgf := inpData.Data[0].Attributes["title"].Value
	//dfd := inpData.Data[0].Uid

	//fmt.Print(fgf,dfd, "значения inpData")

	for k,_ := range inpData.Data {
		outMap[inpData.Data[k].Attributes["title"].Value] = inpData.Data[k].Uid
		//fmt.Print(inpData.Data[k].Attributes["title"].Value, "-кей inpData", inpData.Data[k].Uid, "-value inpData")
	}
	return outMap
}
