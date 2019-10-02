package modules

import (
	cryptorand "crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

func UtilsMergeToMap(data ...map[string]interface{}) map[string]interface{} {
	if len(data) < 1 {
		return nil
	}

	var toMap map[string]interface{} = map[string]interface{}{}
	for _, src := range data {
		for key, val := range src {
			toMap[key] = val
		}
	}

	return toMap
}

func UtilsJsonToMap(data []byte) map[string]interface{} {
	var toMap map[string]interface{}
	json.Unmarshal(data, &toMap)

	return toMap
}

func UtilsStructToMap(data interface{}) (map[string]interface{}, error) {
	byteData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var src map[string]interface{}
	json.Unmarshal(byteData, &src)

	return src, nil
}

func UtilsMapToStruct(src interface{}, data interface{}) error {
	strJson, err := json.Marshal(src)
	if err != nil {
		return err
	}

	return json.Unmarshal(strJson, &data)
}

func UtilsInterfaceToStringInMap(data map[interface{}]interface{}) map[string]interface{} {
	if len(data) < 1 {
		return nil
	}

	var toMap map[string]interface{} = make(map[string]interface{}, len(data))

	for key, val := range data {
		toMap[key.(string)] = val
	}
	return toMap
}

func UtilsStrContains(str string, src ...interface{}) (bool, error) {
	if len(src) < 1 {
		return false, errors.New("source is nil")
	}

	for _, val := range src {
		if strings.Contains(str, val.(string)) {
			return true, nil
		}
	}

	return false, nil
}

func UtilsRandUUID(nums int) (string, error) {
	if nums < 1 {
		return "", errors.New("nums error")
	}

	uuid := make([]byte, nums)
	uuidNums, err := cryptorand.Read(uuid)
	if uuidNums != len(uuid) || err != nil {
		return "", err
	}

	return hex.EncodeToString(uuid), nil
}

func UtilsRandInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func UtilsIsset(key interface{}, arr interface{}, params ...interface{}) (bool, interface{}) {
	switch reflect.TypeOf(key).Kind() {

	case reflect.Int:

		switch reflect.TypeOf(arr).Kind() {

		case reflect.Map:
			if reflect.TypeOf(arr).String() == "map[int]int" {
				data, ok := arr.(map[int]int)[key.(int)]
				return ok, data
			} else {
				return false, nil
			}

		default:
			return false, nil
		}

	case reflect.String:

		switch reflect.TypeOf(arr).Kind() {

		case reflect.Map:
			if reflect.TypeOf(arr).String() == "gin.H" {
				data, ok := arr.(gin.H)[key.(string)]
				if ok == false {
					return ok, nil
				}

				if data == key.(string) {
					return true, data
				} else {
					return false, nil
				}
			} else {
				return false, nil
			}

		default:
			return false, nil
		}

	default:
		return false, nil

	}
}