package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const numberResult  = 7


func geoApi(city, Lat, Lng string) (resultGeo FinalGeo){
	var apigeo = "07ca80d3-7d62-419c-b24b-7453b7e06f0c"

	geo, err := http.Get("https://geocode-maps.yandex.ru/1.x/?apikey="+apigeo+"&geocode="+Lat+","+Lng+"&format=json"+"&results="+strconv.Itoa(numberResult))
	if err != nil {
		fmt.Println("ошибка запроса адреса по гео")
		return
	} else {
		defer geo.Body.Close()
	}

	geoAll, err := ioutil.ReadAll(geo.Body)
	if err != nil {
		fmt.Println("ошибка чтения из запроса geo")
		return
	}

	objgeo := YandexGeo{}
	err = json.Unmarshal(geoAll, &objgeo)

	if err != nil {
		fmt.Println("ошибка преобразования YandexGeo", err)
		return
	}
	//Составной райно города
	//sliceDistrict := []string{}
	f:=len(objgeo.Response.GeoObjectCollection.FeatureMember)

	for i := 0; i<f; i++{

		rrr := objgeo.Response.GeoObjectCollection.FeatureMember[i].GeoObject.MetaDataProperty.GeocoderMetaData.Kind
		qqq := objgeo.Response.GeoObjectCollection.FeatureMember[i].GeoObject.Name
		//kkk := objgeo.Response.GeoObjectCollection.FeatureMember[0].GeoObject.MetaDataProperty.GeocoderMetaData.Address.Components


	if rrr == "house"{
		resultGeo.House = qqq
		//fmt.Print(resultGeo.House)
		continue
	}else if rrr == "street"{
		resultGeo.Street = qqq
		continue
	}else if rrr == "district" {
		//добавляет в слайс значения дикстрикта
		resultGeo.District = qqq
		//sliceDistrict = append(sliceDistrict, qqq)
		//fmt.Println(resultGeo.District, "resultGeo.District1")
		continue
	}else if rrr == "locality" && qqq != city{
		resultGeo.District = qqq
		//fmt.Println(resultGeo.District, "resultGeo.District2")
		resultGeo.City = city
		//fmt.Println(resultGeo.City, "resultGeo.City1")
		continue
	}else if rrr ==  "locality" && qqq == city{
		resultGeo.City = qqq
		//fmt.Println(resultGeo.City, "resultGeo.City2")
		continue
	}
	//else if rrr == "province"{
	//	resultGeo.Region = qqq
	//	continue
	//	//fmt.Println(resultGeo.Region, "resultGeo.Region")
	//}





	}
	//rrr := objgeo.Response.GeoObjectCollection.FeatureMember[0].GeoObject.MetaDataProperty.GeocoderMetaData.Kind
	//qqq := objgeo.Response.GeoObjectCollection.FeatureMember[0].GeoObject.Name
	////lll := objgeo.Response.GeoObjectCollection.FeatureMember
	//
	//
	//fmt.Println(rrr, "rrr")
	//sliceDistrict := []string{}
	//for _, v:= range rrr{
	//	kind := v.Kind
	//	//coord := v.GeoObject.Point.Pos
	//	name := v.Name
	//
	//	//resultGeo.Coords = coord
	//	if kind == "province"{
	//		resultGeo.Region = name
	//		//fmt.Println(resultGeo.Region, "resultGeo.Region")
	//	}else if kind == "locality"{
	//		resultGeo.City = name
	//		//fmt.Println(resultGeo.City, "resultGeo.City")
	//	}else if kind == "district"{
	//		//добавляет в слайс значения дикстрикта
	//		sliceDistrict = append(sliceDistrict, name)
	//		//fmt.Println(resultGeo.District, "resultGeo.District")
	//	}else if kind == "street"{
	//		resultGeo.Street = name
	//		fmt.Println(resultGeo.Street, "resultGeo.Street")
	//	}else if kind == "house"{
	//		resultGeo.House = name
	//		//fmt.Println(resultGeo.House, "resultGeo.House")
	//	}else {
	//		fmt.Println("нет данных с ЯндексГеокодера")
	//	}
	//}

	//переводим слайс в комбинированную строку с разделителем , (составной район города)
	//resultGeo.District = strings.Join(sliceDistrict, ", ")

	return

}
