package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nfnt/resize"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var TplCity = CityId
//var TplRegion = RegionId
var TplParam1945 = RoomsId
var TplParam1957 = TypeId
var TplParam2009 = HouseTypeId
var TplDistrict = DistrictId
var Lat string
var Lng string
var i = 0
var n = 300






func handler(city string) (err error) {

	dateup, _ := ioutil.ReadFile("date.txt")
	var dateto = string(dateup)

	fmt.Println(dateto, "-dateto 2")
	fmt.Println(city, "-city 2")
	fmt.Println(n, "-number")

	var mapSprStack = map[string]string{} //создаем пустую мапу и сразу ее инициализируем
	//fmt.Println("Создается мапка Стэк")

	numb := strconv.Itoa(n)
	re, err := http.Get("https://ads-api.ru/main/api?user=vega-dv@yandex.ru&token=570edf37a945b9434a919a51cb149648&category_id=2,3&date1=" + dateto + "&city=" + city + "&nedvigimost_type=1&source=1,4&withcoords=1&limit="+ numb +"&sort=asc")
	fmt.Println("запрос отправлен")
	if err != nil {
		fmt.Println(err, "ошибка запроса ADS-API")
		return err
	} else {
		defer re.Body.Close()
	}

	b, err := ioutil.ReadAll(re.Body)
	if err != nil {
		fmt.Println("ошибка чтения из запроса - ", err)
		return err
	}

	obj := AvitoJson{}
	err = json.Unmarshal(b, &obj)
	if err != nil {
		fmt.Println("ошибка преобразования AvitoJson - ", err)
		return err
	}

	fmt.Println(len(obj.Data), "массив с парсера")
	if len(obj.Data) == 0{
		return
	}


	//Создаем объект и наполняешь его данными
	var mapSend = map[string]string{}

	//lenData := len(obj.Data)
	//Запрашиваем справочники и проверяем на ошибки
	sprCity := ParsSpr(GetSpr(CityId))
	if len(sprCity) == 0{
		return  err
	}
	sprDistrict := ParsSpr(GetSpr(DistrictId))
	if len(sprDistrict) == 0{
		return  err
	}
	sprStreet := ParsSpr(GetSpr(StreetId))
	if len(sprStreet) == 0{
		return  err
	}
	sprRooms := ParsSpr(GetSpr(RoomsId))
	if len(sprRooms) == 0{
		return  err
	}
	sprType := ParsSpr(GetSpr(TypeId))
	if len(sprType)==0{
		return  err
	}
	sprHouseType := ParsSpr(GetSpr(HouseTypeId))
	if len(sprHouseType)==0{
		return  err
	}

	for _, val := range obj.Data {

		p1945 := fmt.Sprintf("%v", val.Param1945)
		p1945 = strings.TrimSpace(p1945)
		p2113 := fmt.Sprintf("%v", val.Param2113)
		p2213 := fmt.Sprintf("%v", val.Param2213)
		p12722 := fmt.Sprintf("%v", val.Param12722)
		p12721 := fmt.Sprintf("%v", val.Param12721)
		p2636 := fmt.Sprintf("%v", val.Param2636)
		p2736 := fmt.Sprintf("%v", val.Param2736)


		mapSend["url"] = val.URL
		mapSend["price"] = strconv.Itoa(val.Price) //преобразуем string в type
		//mapSend["time"] = obj.Data[0].Time.Format("02-Jan-2006")
		mapSend["time"] = val.Time
		dateto = strings.ReplaceAll(val.Time, " ", "+")
		mapSend["phone"] = val.Phone

		//Меняем содержимое строки в названии клиента
		j := strings.NewReplacer("Агенство недвижимости","АН","Бюро недвижимости","АН")
		mapSend["person"] = j.Replace(val.Person)

		mapSend["person_type"] = val.PersonType
		mapSend["coords"] = val.Coords.Lat + ":" + val.Coords.Lng

		//получаем адрес по координатам
		Lng = val.Coords.Lat
		Lat = val.Coords.Lng

		resultAddress := geoApi(city, Lat, Lng)

		mapSprStack["coords"] = resultAddress.Coords

		//вызываем справочник регион
		//sprRegion := ParsSpr(GetSpr(RegionId))
		//fmt.Println(sprRegion, "sprReg")

		////проверка на наличие в справчонике региона
		//if v, found := sprRegion[resultAddress.Region]; found {
		//	mapSend["region_pointvalue"] = resultAddress.Region
		//	mapSend["region_pointsrc"] = v
		////	fmt.Println(mapSend["region_pointvalue"], mapSend["region_pointsrc"], "mapsendRegion")
		//} else {
		//	resultKey, resultValue := createSprValue(resultAddress.Region, TplRegion, mapSprStack)
		//
		//	// добавляем соданный ид-региона в слайс справочника региона
		//
		//	sprRegion[resultKey] = resultValue
		//
		//	// добавляем поле региона в слайс объекта объявления для отправлки на создание в справочнике региона
		//	mapSend["region_pointvalue"] = resultValue
		//	mapSend["region_pointsrc"] = resultKey
		//}
		////добавляем в мапу пару строк с названиями и айди элементов связанных шаблонов
		//mapSprStack["region_pointvalue"] = mapSend["region_pointvalue"]
		//mapSprStack["region_pointsrc"] = mapSend["region_pointsrc"]

		//вызываем функцию запроса справочника города

		//fmt.Println(sprCity, "sprCity")

		//проверка на наличие в справочнике города
		if v, found := sprCity[resultAddress.City]; found {
			mapSend["city1_pointvalue"] = resultAddress.City
			mapSend["city1_pointsrc"] = v
			//fmt.Println(mapSend["city1_pointvalue"],"валуе город сити1 проверка", mapSend["city1_pointsrc"], "кей город сити1 проверка" )
			//fmt.Println(sprCity[v], "sprcity-v")
		} else {
			resultKey, resultValue := createSprValue(resultAddress.City, TplCity, mapSprStack)
			//fmt.Print(resultKey, resultValue, "ключ и значение города")

			// добавляем соданный ид-города в слайс справочника городов
			if _, found := sprCity[resultKey]; !found {
				sprCity[resultKey] = resultValue
				//fmt.Println(sprCity, "слайс справчоника городов")
			}

			// добавляем поле города в слайс объекта объявления для отправлки на создание в справочнике город
			mapSend["city1_pointvalue"] = resultValue
			mapSend["city1_pointsrc"] = resultKey
			//fmt.Println(mapSend["city1_pointvalue"], "валуе город сити1 создание", mapSend["city1_pointsrc"], "кей город сити1 создание /n" )
		}
		//добавляем в мапу пару строк с названиями и айди элементов связанных шаблонов.
		mapSprStack["city1_pointvalue"] = mapSend["city1_pointvalue"]
		mapSprStack["city1_pointsrc"] = mapSend["city1_pointsrc"]

		//mapSend["region_pointvalue"] = reg
		//mapSend["region_pointsrc"] = regSrc

		//вызываем функцию запроса справочника район города

		//fmt.Println(sprDistrict, "sprDis")

		if v, found := sprDistrict[resultAddress.District]; found {
			mapSend["district_pointvalue"] = resultAddress.District
			mapSend["district_pointsrc"] = v
			//fmt.Println(mapSend["district_pointvalue"],mapSend["district_pointsrc"], "Dist found" )
		} else {
			resultKey, resultValue := createSprValue(resultAddress.District, TplDistrict, mapSprStack)

			// добавляем соданный ид-региона в слайс справочника региона
			sprDistrict[resultKey] = resultValue
			//fmt.Print(sprDistrict, "добавляем ид региона в српавочник")

			// добавляем поле региона в слайс объекта объявления для отправлки на создание в справочнике региона
			mapSend["district_pointvalue"] = resultValue
			mapSend["district_pointsrc"] = resultKey
		}

		mapSprStack["district_pointvalue"] = mapSend["district_pointvalue"]
		mapSprStack["district_pointsrc"] = mapSend["district_pointsrc"]
		//fmt.Println(mapSprStack["district_pointvalue"],mapSprStack["district_pointsrc"], "Dist found 2" )

		//справочник улиц

		//fmt.Println(sprStreet, "sprStreet")

		if v, found := sprStreet[resultAddress.Street]; found {
			mapSend["street_pointvalue"] = resultAddress.Street
			mapSend["street_pointsrc"] = v
		} else {
			resultKey, resultValue := createSprValue(resultAddress.Street, TplStreet, mapSprStack)

			// добавляем соданный ид-региона в слайс справочника региона
			sprStreet[resultKey] = resultValue

			// добавляем поле региона в слайс объекта объявления для отправлки на создание в справочнике региона
			mapSend["street_pointvalue"] = resultValue
			mapSend["street_pointsrc"] = resultKey
		}
		mapSprStack["street_pointvalue"] = mapSend["street_pointvalue"]
		mapSprStack["street_pointsrc"] = mapSend["street_pointsrc"]

		//Проверяем приходят ли адрес сгеопарсера, если нет то переходим к след объявлению
		if resultAddress.House == "" {
			continue
		}

		if val.Param2836 == "0" || val.Param2313 == "0"{
			continue
		}


		//Меняем содержимое строки в адресе
		r:= strings.NewReplacer("улица", "ул", "проспект", "пр-кт", "переулок", "пер", "квартал", "кв-л")
		mapSend["house"] = r.Replace(resultAddress.House)
		//fmt.Println(mapSend["house"], "house" )

		mapSend["metro"] = val.Metro

		//fmt.Println(val.Address, "Address")
		if val.Address == "" {
			fmt.Println("нет адреса")
			continue
		}
		//Поле полный алрес
		//mapSend["address"] = val.Address
		mapSend["description"] = val.Description
		mapSend["nedvigimost_type"] = val.NedvigimostType
		mapSend["avitoid"] = fmt.Sprint(val.Avitoid)
		mapSend["source"] = val.Source
		mapSend["id"] = strconv.Itoa(val.ID)
		mapSend["cat1"] = val.Cat1
		mapSend["cat2"] = val.Cat2
		mapSend["param_1943"] = val.Param1943
		//mapSend["param_1945"] = strconv.Itoa(obj.Data[0].Param1945)

		//переводим int  в строку
		//fmt.Print(obj.Data[n].Param1945, "obj.Data[n].Param1945 - ")
		//comn := obj.Data[n].Param1945
		//fmt.Print(comn, "comn - ")

		//вызываем функцию запроса справочника кол-во комнат

		//fmt.Println(sprRooms, "sprRoom српавочник")

		//fmt.Println("проверка на комнаты")

		//проверка на категорию Комнаты

		if val.Cat2 == "Комнаты" {
			mapSend["param_1945_pointvalue"] = "Комната"
			mapSend["param_1945_pointsrc"] = "2020-02-24T04-07-58z03-00-0136a4"
			//проверка на наличие в справчонике кол-во комнат
		}else if p1945 =="" || p1945 =="<nil>"{
			fmt.Println(p1945, "Нет кол-ва комнат")
			continue
		}else {
			if v, found := sprRooms[p1945]; found {
				mapSend["param_1945_pointvalue"] = p1945
				mapSend["param_1945_pointsrc"] = v
			} else {
				resultKey, resultValue := createSprValue(p1945, TplParam1945, mapSprStack)

				// добавляем соданный ид-города в слайс справочника городов
				sprRooms[resultKey] = resultValue
				fmt.Print(sprRooms[resultKey], sprRooms[resultValue], "- sprRooms")

				// добавляем поле города в слайс объекта объявления для отправлки на создание в справочнике город
				mapSend["param_1945_pointvalue"] = resultValue
				mapSend["param_1945_pointsrc"] = resultKey
				//fmt.Println(mapSend["param_1945_pointvalue"], mapSend["param_1945_pointsrc"], "не проходит проверку Rooms")
			}
		}

		//вызываем функцию запроса справочника вид объекта

		//fmt.Println(sprType, "sprType")

		//проверка на наличие в справчонике вид объекта
		if val.Cat2 == "Комнаты" {
			mapSend["param_1957_pointvalue"] = "Вторичка"
			mapSend["param_1957_pointsrc"] = "2020-02-21T06-47-35z03-00-660c17"

		}else if val.Param1957 ==""{
			mapSend["param_1957_pointvalue"] = "Вторичка"
			mapSend["param_1957_pointsrc"] = "2020-02-21T06-47-35z03-00-660c17"
			fmt.Println(val.Param1957, "Нет вида квартиры")
		}else {
			if v, found := sprType[val.Param1957]; found {
				mapSend["param_1957_pointvalue"] = val.Param1957
				mapSend["param_1957_pointsrc"] = v
				//fmt.Println(mapSend["param_1957_pointvalue"], mapSend["param_1957_pointsrc"], "- вид объекта")

			} else {

				resultKey, resultValue := createSprValue(val.Param1957, TplParam1957, mapSprStack)

				// добавляем соданный ид-города в слайс справочника городов
				sprType[resultKey] = resultValue

				// добавляем поле города в слайс объекта объявления для отправлки на создание в справочнике
				mapSend["param_1957_pointvalue"] = resultValue
				mapSend["param_1957_pointsrc"] = resultKey
				fmt.Println(mapSend["param_1957_pointvalue"], mapSend["param_1957_pointsrc"], "не проходит проверку Type")
			}
		}

		//вызываем функцию запроса справочника тип дома

		//fmt.Println(sprHouseType, "sprHouseType")

		if val.Cat2 == "Квартиры"{
			//проверка на наличие в справчонике тип дома param_2009

			if v, found := sprHouseType[val.Param2009]; found {
				mapSend["param_2009_pointvalue"] = val.Param2009
				mapSend["param_2009_pointsrc"] = v
			}else if val.Param2009 ==""{
				mapSend["param_2009_pointvalue"] = "Не указано"
				mapSend["param_2009_pointsrc"] = "2020-03-03T11-41-19z03-00-1ca186"
				//fmt.Println(val.Param2009, "Нет типа стен квартиры")
			}else {
				resultKey, resultValue := createSprValue(val.Param2009, TplParam2009, mapSprStack)

				// добавляем соданный ид-города в слайс справочника городов
				sprHouseType[resultKey] = resultValue

				//// добавляем поле города в слайс объекта объявления для отправлки на создание в справочнике город
				mapSend["param_2009_pointvalue"] = resultValue
				mapSend["param_2009_pointsrc"] = resultKey
				//fmt.Println(mapSend["param_1957_pointvalue"], mapSend["param_1957_pointsrc"], "не проходит проверку Type")
			}
		}

		if val.Cat2 == "Комнаты"{
			//проверка на наличие в справчонике тип дома param_2009

			if v, found := sprHouseType[val.Param2567]; found {
				mapSend["param_2009_pointvalue"] = val.Param2567
				mapSend["param_2009_pointsrc"] = v
			}else if val.Param2567 ==""{
				mapSend["param_2009_pointvalue"] = "Не указано"
				mapSend["param_2009_pointsrc"] = "2020-03-03T11-41-19z03-00-1ca186"
				fmt.Println("Нет типа стен комнаты")
			}else {
				resultKey, resultValue := createSprValue(val.Param2567, TplParam2009, mapSprStack)

				// добавляем соданный ид-города в слайс справочника городов
				sprHouseType[resultKey] = resultValue

				//// добавляем поле города в слайс объекта объявления для отправлки на создание в справочнике город
				mapSend["param_2009_pointvalue"] = resultValue
				mapSend["param_2009_pointsrc"] = resultKey
				//fmt.Println(mapSend["param_1957_pointvalue"], mapSend["param_1957_pointsrc"], "не проходит проверку Type")
			}
		}

		var imgResult []string
		//Собираем урл картинок из массива в строку с разделителем ;
		for _, v := range val.Images {
			imgResult = append(imgResult, fmt.Sprint(v.Imgurl))
		}

		//Создаем картинку для превьюшки
 		if len(val.Images) != 0 && val.Images[0].Imgurl != ""  {
			smImg := imgresize(val.Images[0].Imgurl, val.ID)
			if smImg == "0" {
				mapSend["smimg"] = "/kvadrat/gui/upload/1kvadrat/noImg.jpg"
			} else {
				mapSend["smimg"] = smImg
			}
			//fmt.Println(mapSend["smimg"], "smImg")
		}else {
			mapSend["smimg"] = "/kvadrat/gui/upload/1kvadrat/noImg.jpg"
		}

		imgResultOut := strings.Join(imgResult[:], ";")

		mapSend["images"] = imgResultOut
		//mapSend["images"] = obj.Data[0].Images[0].Imgurl
		//fmt.Println(imgResultOut, "img")

		if val.Cat2 == "Комнаты" {
			mapSend["param_2313"] = val.Param2836
		} else {
			mapSend["param_2313"] = val.Param2313
		}

		if val.Cat2 == "Комнаты" {
			mapSend["param_2113"] = p2636
		} else {
			mapSend["param_2113"] = p2113
		}

		if val.Cat2 == "Комнаты" {
			mapSend["param_2213"] = p2736
		} else {
			mapSend["param_2213"] = p2213
		}

		mapSend["param_12721"] = p12721
		mapSend["param_12722"] = p12722

		if mapSend["param_2113"] == mapSend["param_2213"] || mapSend["param_2113"] == "1" {
			//mapSend["LastFloor"] = "checked"
			mapSend["LastFloor_pointvalue"] = "Крайний"
			mapSend["LastFloor_pointsrc"] = "2019-12-02T08-50-11Z-34692d"
		} else {
			mapSend["LastFloor_pointvalue"] = "Средний"
			mapSend["LastFloor_pointsrc"] = "2019-12-02T08-50-46Z-837703"
		}

		mapSend["count_ads_same_phone"] = fmt.Sprint(val.CountAdsSamePhone)
		mapSend["phone_protected"] = fmt.Sprint(val.PhoneProtected)
		//fmt.Println(mapSend["count_ads_same_phone"], mapSend["phone_protected"], "доп поля")

		if val.PersonType == "Агентство"{
			mapSend["owner"] = "0"
		}else{
			//&& val.PhoneProtected == 0
			if val.CountAdsSamePhone <= owner {
				mapSend["owner"] = "1"
			}else{
				mapSend["owner"] = "0"
			}
		}

		mapSend["title"] = val.Title
		mapSend["KmDoMetro"] = string(val.KmDoMetro)

		//Формируем краткий тайтл
		l := strings.NewReplacer("1","1-к квартира, ","2","2-к квартира, ","3","3-к квартира, ","4","4-к квартира, ","5","5-к квартира, ","6","6-к квартира, ","7","7-к квартира, ", "8","8-к квартира, ")
		mapSend["smtitle"] = l.Replace(mapSend["param_1945_pointvalue"])+mapSend["param_2313"]+"(м²), "+mapSend["param_2113"]+"/"+mapSend["param_2213"]+" эт"

		//Заполняем системные поля
		mapSend["data-source"] = "2019-11-07T04-48-08Z-10190e"
		mapSend["data-clientname"] = "Денис Оленич"
		mapSend["data-clientid"] = "f48d38aa-4160-e772-8082-2a76a573716f"
		mapSend["data-uid"] = fmt.Sprint(val.ID)

		fmt.Println(dateto, "dateto3")
		fmt.Println(val.ID, "ID 3")
		fmt.Println()
		time.Sleep(5000 * time.Millisecond)

		//Перевод объекта в json
		//fmt.Println(mapSend)
		jsonMap, err := json.Marshal(mapSend)
		if err != nil {
			fmt.Println("ошибка преобразования мапы", err)
			continue
		}

		//отправляем тело постом
		jsonR, err := http.Post("https://buildbox.app/kvadrat/gui/objs?format=json", "application/json", bytes.NewBuffer(jsonMap))
		if err != nil{
			fmt.Println("ошибка запроса jsonR", err)
			//fmt.Println(jsonR, "JsonR")
			continue
		}


		if jsonR == nil{
			fmt.Println("ошибка jsonR")
			continue
		}
		//if jsonR.Body == nil {
		//	fmt.Println("ошибка jsonR.Body")
		//	//return
		//}

		var result map[string]interface{}
		json.NewDecoder(jsonR.Body).Decode(&result)
		//fmt.Println(jsonR, "jsoR")
		//fmt.Println(jsonR.Body, "jsonRBody")
	}
		dateout := []byte(dateto)
		_ = ioutil.WriteFile("date.txt", dateout, 0644)
		//if err != nil{
		//	fmt.Println(err, "ошибка записи даты")
		//}

		i++
		fmt.Println(i, "дергаем хандлер")
		//t := time.Now()
		//fmt.Printf(t.Format(time.RFC3339))

		if len(obj.Data) == n{
			handler(city)
			return
		}else if len(obj.Data) == 0 {
			handler(city)
			return
		}else {
			fmt.Printf(t.Format(time.RFC3339))
			time.Sleep(30 * time.Minute)
			handler(city)
			return
		}



}

//func static(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	w.Write([]byte(json1))
//}


// возвращаем необходимый значение атрибута для объекта если он есть, инае пусто
// а также из заголовка объекта
func (p *Data) Attr(name, element string) (result string, found bool) {

	if _, found := p.Attributes[name]; found {

		// фикс для тех объектов, на которых добавлено скрытое поле Uid
		if name == "uid" {
			return p.Uid, true
		}

		switch element {
		case "src":
			return p.Attributes[name].Src, true
		case "value":
			return p.Attributes[name].Value, true
		case "tpls":
			return p.Attributes[name].Tpls, true
		case "rev":
			return p.Attributes[name].Rev, true
		case "status":
			return p.Attributes[name].Status, true
		case "uid":
			return p.Uid, true
		case "source":
			return p.Source, true
		case "id":
			return p.Id, true
		case "title":
			return p.Title, true
		case "type":
			return p.Type, true
		}
	} else {
		switch name {
		case "uid":
			return p.Uid, true
		case "source":
			return p.Source, true
		case "id":
			return p.Id, true
		case "title":
			return p.Title, true
		case "type":
			return p.Type, true
		}
	}
	return "", false
}


func createSprValue(value, tpl string, stack map[string]string) (resKey, resValue string){
		var mapSendCity = map[string]string{}
		mapSendCity["title"] = value
		mapSendCity["data-source"] = tpl
		mapSendCity["data-clientname"] = "Денис Оленич"
		mapSendCity["data-clientid"] = "f48d38aa-4160-e772-8082-2a76a573716f"

		//кладем в мапку значение связаного справочника (типа город-региона)
		for k, v := range stack {
			mapSendCity[k] = v
		}

		mapSendCityOut, err := json.Marshal(mapSendCity)
		if err != nil{
			fmt.Println("ошибка преобразования города" , err)
			return
		}

	//запрос на создание объекта город в справочнике
	jsonCity, err := http.Post("https://buildbox.app/kvadrat/gui/objs?format=json", "application/json", bytes.NewBuffer(mapSendCityOut))
	if err != nil{
		return
	}

	var resultCity ResponseData
	json.NewDecoder(jsonCity.Body).Decode(&resultCity)


	resKey = resultCity.Data[0].Uid
	resValue = resultCity.Data[0].Attributes["title"].Value
	//fmt.Println(resKey, resValue, "- отправляем данные в справочник города")

	return
}

func imgresize(Imgurl string, ID int)(smImg string){
	f, err := http.Get(Imgurl)
	if err != nil {
		fmt.Println(err, "error1")
		smImg = "0"
		return smImg
	}

	defer f.Body.Close()

	m, err := jpeg.Decode(f.Body)
	if err != nil {
		fmt.Println(err, "error2")
		smImg = "0"
		return smImg
	}

	resImg := resize.Resize(500, 375, m, resize.NearestNeighbor)

	strID := strconv.Itoa(ID)

	img, err := os.Create("img/"+strID+".jpg")
	//fmt.Println(strID, "strID")
	if err != nil {
		fmt.Println(err, "error3")
		smImg = "0"
		return smImg
	}
	defer img.Close()

	jpeg.Encode(img, resImg, nil )
	if err != nil {
		fmt.Println(err, "error4")
		smImg = "0"
		return smImg
	}

	//Проверяем размер полученного изображения.
	ImgSize, _ := os.Stat("img/"+strID+".jpg")
	if ImgSize.Size() == 0{
		smImg = "0"
		return smImg
	}

	smImg = "http://cdn.domotchet.ru/"+strID+".jpg"

	return smImg
}