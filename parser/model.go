package main

type YandexGeo struct {
	Response struct {
		GeoObjectCollection struct {
			MetaDataProperty struct {
				GeocoderResponseMetaData struct {
					Point struct {
						Pos string `json:"pos"`
					} `json:"Point"`
					Request string `json:"request"`
					Results string `json:"results"`
					Found   string `json:"found"`
				} `json:"GeocoderResponseMetaData"`
			} `json:"metaDataProperty"`
			FeatureMember []struct {
				GeoObject struct {
					MetaDataProperty struct {
						GeocoderMetaData struct {
							Precision string `json:"precision"`
							Text      string `json:"text"`
							Kind      string `json:"kind"`
							Address   struct {
								CountryCode string `json:"country_code"`
								Formatted   string `json:"formatted"`
								PostalCode  string `json:"postal_code"`
								Components  []struct {
									Kind string `json:"kind"`
									Name string `json:"name"`
								} `json:"Components"`
							} `json:"Address"`
							AddressDetails interface{} `json:"AddressDetails"`
						} `json:"GeocoderMetaData"`
					} `json:"metaDataProperty"`
					Name        string `json:"name"`
					Description string `json:"description"`
					BoundedBy   struct {
						Envelope struct {
							LowerCorner string `json:"lowerCorner"`
							UpperCorner string `json:"upperCorner"`
						} `json:"Envelope"`
					} `json:"boundedBy"`
					Point struct {
						Pos string `json:"pos"`
					} `json:"Point"`
				} `json:"GeoObject"`
			} `json:"featureMember"`
		} `json:"GeoObjectCollection"`
	} `json:"response"`
}

type AvitoJson struct {
	Code int `json:"code"`
	Data []struct {
	URL             string `json:"url"`
	Title           string `json:"title"`
	Price           int    `json:"price"`
	Time            string `json:"time"`
	Phone           string `json:"phone"`
	Person          string `json:"person"`
	PersonType      string `json:"person_type"`
	City            string `json:"city"`
	Metro           string `json:"metro"`
	Address         string `json:"address"`
	House         string `json:"house"`
	Description     string `json:"description"`
	NedvigimostType string `json:"nedvigimost_type"`
	Avitoid         interface{} `json:"avitoid"`
	Cat1ID          int    `json:"cat1_id"`
	Cat2ID          int    `json:"cat2_id"`
	Source          string `json:"source"`
	PhoneProtected  int    `json:"phone_protected"`
	ID              int    `json:"id"`
	Cat1            string `json:"cat1"`
	Cat2            string `json:"cat2"`
	Images          []struct {
	Imgurl string `json:"imgurl"`
} `json:"images"`
	Param1943  string `json:"param_1943"`
	Param1945  interface{}    `json:"param_1945"`
	Param1957  string `json:"param_1957"`
	Param2009  string `json:"param_2009"`
	Param2313  string    `json:"param_2313"`
	Param2113  interface{}     `json:"param_2113"`
	Param2213  interface{}    `json:"param_2213"`
	Param12721 interface{}    `json:"param_12721"`
	Param12722 interface{}    `json:"param_12722"`
	Param2545 interface{} `json:"param_2545"`
	Param2567 string `json:"param_2567"`
	Param2636 interface{} `json:"param_2636"`
	Param2736 interface{} `json:"param_2736"`
	Param2836 string `json:"param_2836"`
	Param2837 string `json:"param_2837"`
	Coords     struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
} `json:"coords"`
	PersonTypeID      int    `json:"person_type_id"`
	KmDoMetro      int    `json:"km_do_metro"`
	NedvigimostTypeID interface{}    `json:"nedvigimost_type_id"`
	SourceID          int    `json:"source_id"`
	Region            string `json:"region"`
	City1             string `json:"city1"`
	PhoneOperator     string `json:"phone_operator"`
	Params            interface{} `json:"params"`
	CountAdsSamePhone int `json:"count_ads_same_phone"`
} `json:"data"`
	Status string `json:"status"`
}

type FinalGeo struct {
	Region string `json:"region"`
	City string `json:"city"`
	Street string `json:"street"`
	House string `json:"house"`
	District string `json:"district"`
	Coords string `json:"coords"`
}


type ResponseData struct {
	Data      []Data        `json:"data"`
	Res      interface{}  `json:"res"`
	Status    RestStatus    `json:"status"`
	Metrics   Metrics   `json:"metrics"`
}

type Metrics struct {
	ResultSize      int `json:"result_size"`
	ResultCount     int `json:"result_count"`
	ResultOffset    int `json:"result_offset"`
	ResultLimit     int `json:"result_limit"`
	ResultPage   int `json:"result_page"`
	TimeExecution   string `json:"time_execution"`
	TimeQuery    string `json:"time_query"`
}

type Attribute struct {
	Value  string `json:"value"`
	Src    string `json:"src"`
	Tpls   string `json:"tpls"`
	Status string `json:"status"`
	Rev    string `json:"rev"`
	Uuid   string `json:"uuid"`
}

type Data struct {
	Uid        string               `json:"uid"`
	Id         string               `json:"id"`
	Source     string               `json:"source"`
	Parent     string               `json:"parent"`
	Type       string               `json:"type"`
	Title      string               `json:"title"`
	Rev        string               `json:"rev"`
	Attributes map[string]Attribute `json:"attributes"`
	Linkinid       string `json:"linkinid"`
	Linkinobj       []Data `json:"linkinobj"`
}


type RestStatus struct {
	Description string `json:"description"`
	Status      int    `json:"status"`
	Code        string `json:"code"`
	Error       string  `json:"error"`
}