package csObject

import (
	"encoding/json"
	"errors"
	"fmt"

	A "github.com/CritsendGo/ApiClient"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var ApiToken string
var apiClient = A.NewClient("")

type objectStatus struct {
	Name   string
	IsSave bool
}

var (
	oStatusNew     = objectStatus{"New", false}
	oStatusSaved   = objectStatus{"Saved", true}
	oStatusUpdated = objectStatus{"Updated", false}
)

var (
	ErrorEmptyToken   = errors.New("No token set on ApiToken")
	ErrorJsonUnFormat = errors.New("Unable to set json from the object")
)

func FromJson(f json.RawMessage) {
	// Retrieve tag format to set the good format into the struct
}
func checkOrSetToken() {
	if ApiToken == "" {
		fmt.Println("No token")
	}
	if apiClient.Token == "" {
		apiClient.Token = ApiToken
	}
}

func Save(o interface{}) error {
	SaveObject(o)
	return nil
}
func Get(o interface{}) error {
	checkOrSetToken()
	status := CheckObject(o)
	if status == "OPTIONAL" {
		return errors.New("all non optional field is not set")
	}
	if status == "UNKNOWN" {
		// Primary is not set get by its non optional
		return GetByNonOptional(o)
	}
	// Primary is set get by its id
	return GetById(o)

	//result,e:=apiClient.Get(api+"/?"+api+"_name="+name)
	//data:=MapToObject(result,o)

}
func CheckObject(o interface{}) string {
	// Used to check that all fields with optional = false is not empty
	// Return :
	//	- Error : If all field is not set
	//  - Unknown : If all field is ok but primary empty
	//  - Ok : If primary and optional is ok
	oValue := reflect.ValueOf(o)
	oModel := reflect.TypeOf(o)
	oOrigin := oValue
	if oValue.Kind() == reflect.Ptr {
		oValue = oValue.Elem()
	}
	var isOkOptional = false
	var isOkPrimary = false
	for i := 0; i < oModel.Elem().NumField(); i++ {
		isPrimary, _ := strconv.ParseBool(oModel.Elem().Field(i).Tag.Get("oPrimary"))
		isOptional, _ := strconv.ParseBool(oModel.Elem().Field(i).Tag.Get("oOptional"))
		fieldValue := strings.TrimSpace(fmt.Sprint(oOrigin.Elem().Field(i)))
		if isPrimary == true && fieldValue != "0" {
			isOkPrimary = true
		}
		if isOptional == false && len(fieldValue) > 0 {
			isOkOptional = true
		}
	}
	if isOkOptional == false {
		return "OPTIONAL"
	}
	if isOkPrimary == false {
		return "UNKNOWN"
	}
	return "OK"
}

func GetById(o interface{}) error {

	return nil
}
func GetByNonOptional(o interface{}) error {
	oValue := reflect.ValueOf(o)
	oModel := reflect.TypeOf(o)
	apiBrut := fmt.Sprint(reflect.TypeOf(o))
	apiName := strings.ToLower(strings.Replace(strings.Replace(apiBrut, "csObject.", "", 1), "*", "", 1))

	oOrigin := oValue
	if oValue.Kind() == reflect.Ptr {
		oValue = oValue.Elem()
	}
	search := make(map[string]string)

	for i := 0; i < oModel.Elem().NumField(); i++ {
		isOptional, _ := strconv.ParseBool(oModel.Elem().Field(i).Tag.Get("oOptional"))
		tagType := oModel.Elem().Field(i).Tag.Get("oType")
		tagName := strings.Split(oModel.Elem().Field(i).Tag.Get("json"), ",omit")[0]
		fieldValue := strings.TrimSpace(fmt.Sprint(oOrigin.Elem().Field(i)))
		if tagType == "struct" {
			if fieldValue != "<nil>" {
				fieldValue = strings.TrimSpace(fmt.Sprint(oOrigin.Elem().Field(i).Elem().Field(0)))
			} else {
				fieldValue = ""
			}
		}

		if isOptional == false {
			search[tagName] = fieldValue
		}
	}
	var mapField []string
	for name, value := range search {
		mapField = append(mapField, name+"="+value)
	}
	path := apiName + "/?" + strings.Join(mapField, "&")
	//fmt.Println("GET ",apiName,path)
	result, e := apiClient.Get(path)

	if e != nil {
		// Not found need to be save
		fmt.Println("Error in apiclient get ", e)
		return e
	}
	if len(result) != 1 {
		return errors.New("too many result from the non optional value")
	}
	data := result[0]
	//fmt.Println(data)
	err := MapToObject(data, o)
	//mt.Print(o)
	return err

}

func GetAll(o interface{}) (i []interface{}) {
	checkOrSetToken()
	apiBrut := fmt.Sprint(reflect.TypeOf(o))
	apiName := strings.ToLower(strings.Replace(strings.Replace(apiBrut, "csObject.", "", 1), "*", "", 1))
	path := apiName + "/"
	//fmt.Println("GET ",apiName,path)
	fmt.Println(path)
	result, e := apiClient.Get(path)
	if e != nil {
		// Not found need to be save
		fmt.Println("Error in apiclient get ", e)
		return i
	}
	if len(result) < 1 {
		fmt.Println("No result ", e)
		return
	}
	for _, res := range result {
		fmt.Println(res)
		newO := reflect.New(reflect.TypeOf(o))
		err := MapToObject(res, newO)
		if err != nil {
			i = append(i, newO)
		} else {
			fmt.Println(err)
		}
	}
	return i
}

func MapToObject(a map[string]string, o interface{}) error {
	model := reflect.TypeOf(o)
	reflect.New(model)
	oValue := reflect.ValueOf(o)
	if oValue.Kind() == reflect.Ptr {
		oValue = oValue.Elem()
	}
	for fieldName, fieldValue := range a {
		for i := 0; i < model.Elem().NumField(); i++ {
			tagName := strings.Split(model.Elem().Field(i).Tag.Get("json"), ",omit")[0]
			tagType := model.Elem().Field(i).Tag.Get("oType")
			if tagName == fieldName {
				if tagType == "string" {
					oValue.Field(i).SetString(fieldValue)
				}
				if tagType == "int" {
					b, _ := strconv.Atoi(fieldValue)

					oValue.Field(i).SetInt(int64(b))
				}
				if tagType == "time" {
					f, _ := time.Parse("2006-01-02 15:04:05", fieldValue)
					oValue.Field(i).Set(reflect.ValueOf(f))
				}
			}
		}
	}
	return nil
}

func SaveObject(f interface{}) (e error) {
	apiBrut := fmt.Sprint(reflect.TypeOf(f))
	apiName := strings.ToLower(strings.Replace(strings.Replace(apiBrut, "csObject.", "", 1), "*", "", 1))
	//fmt.Println(apiName,"POST")
	if ApiToken == "" {
		fmt.Println("No token")
		return ErrorEmptyToken
	}
	if apiClient.Token == "" {
		apiClient.Token = ApiToken
	}
	jsonString, err := ToJson(f)
	if err != nil {
		fmt.Println("Unable to set JSon")
		return ErrorJsonUnFormat
	}
	body := "[" + jsonString + "]"
	//fmt.Println(apiName,"POST",body)
	mapping, err := apiClient.Insert(apiName, body)
	/*id,_:=strconv.Atoi(mapping[0][apiName+"_id"])
	v:=int64(id)
	g:=reflect.ValueOf(f)
	if g.Kind() == reflect.Ptr {
		g = g.Elem()
	}
	for i := 0; i < reflect.TypeOf(g).NumField(); i++ {
		if reflect.TypeOf(g).Field(i).Name=="Id"{
			g.Field(i).SetInt(v)
		}
	}
	return
	*/
	return MapToObject(mapping[0], f)
}

func ToJson(f interface{}) (string, error) {
	g := reflect.ValueOf(f)
	oModel := reflect.TypeOf(f)
	if g.Kind() == reflect.Ptr {
		g = g.Elem()
	}
	var jsonMap = make(map[string]string)
	for i := 0; i < oModel.Elem().NumField(); i++ {
		tagType := oModel.Elem().Field(i).Tag.Get("oType")
		tagName := strings.Split(oModel.Elem().Field(i).Tag.Get("json"), ",omit")[0]
		fieldValue := ""
		if tagType == "struct" {
			if g.Field(i).Elem().Kind() != reflect.Invalid {
				fieldValue = fmt.Sprint(g.Field(i).Elem().Field(0))
				jsonMap[tagName] = fieldValue
			}

		} else if tagType == "time" {
			timeVal := strings.Split(fmt.Sprint(g.Field(i)), " m=")[0]
			timeFormat := "2006-01-02 15:04:05.999999999 -0700 MST"
			r, _ := time.Parse(timeFormat, timeVal)
			timeString := r.Format("2006-01-02 15:04:05")
			if timeString == "0001-01-01 00:00:00" {
				timeString = time.Now().Format("2006-01-02 15:04:05")
			}
			fieldValue = timeString
			jsonMap[tagName] = fieldValue

		} else if tagType == "string" {
			fieldValue = g.Field(i).String()
			if fieldValue != "" {
				jsonMap[tagName] = fieldValue
			}
		} else if tagType == "int" {
			fieldValue := fmt.Sprint(g.Field(i))
			if fieldValue != "0" {
				jsonMap[tagName] = fieldValue
			}

		}

	}

	jsonString, _ := json.Marshal(jsonMap)
	return string(jsonString), nil
}
func IsSubStruct(g interface{}) bool {
	numFields := reflect.TypeOf(g).NumField()
	for i := 0; i < numFields; i++ {
		if reflect.ValueOf(g).Field(i).Kind() == reflect.Struct {
			if reflect.TypeOf(g).Field(i).Name != "status" {
				return true
			}
		}
	}
	return false
}
