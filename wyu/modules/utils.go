package modules

import (
	cRand "crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/spf13/cast"
	"math/rand"
	"strings"
	"time"
)

func UtilsMergeToMap(data ...map[string]interface{}) (toMap map[string]interface{}) {
	if len(data) < 1 {
		return
	}

	toMap = map[string]interface{}{}
	for _, src := range data {
		for key, val := range src {
			toMap[key] = val
		}
	}

	return
}

func UtilsJsonToMap(data []byte) (toMap map[string]interface{}) {
	err := json.Unmarshal(data, &toMap)
	if err != nil {
		return
	}

	return
}

func UtilsStructToMap(data interface{}) (src map[string]interface{}, err error) {
	byteData, err := json.Marshal(data)
	if err != nil {
		return
	}

	err = json.Unmarshal(byteData, &src)
	return
}

func UtilsMapToStruct(src interface{}, data interface{}) (err error) {
	strJson, err := json.Marshal(src)
	if err != nil {
		return
	}

	return json.Unmarshal(strJson, &data)
}

func UtilsInterfaceToStringInMap(data map[interface{}]interface{}) (toMap map[string]interface{}) {
	if len(data) < 1 {
		return
	}

	toMap = make(map[string]interface{}, len(data))
	for key, val := range data {
		toMap[cast.ToString(key)] = val
	}

	return
}

func UtilsStrContains(str string, src ...interface{}) (ok bool, err error) {
	if len(src) < 1 {
		err = errors.New("source is nil")
		return
	}

	for _, val := range src {
		if strings.Contains(str, cast.ToString(val)) {
			ok = true
			return
		}
	}

	return
}

func UtilsRandUUID(nums int) (strUUID string, err error) {
	if nums < 1 {
		err = errors.New("nums error")
		return
	}

	uuid := make([]byte, nums)
	uuidNums, err := cRand.Read(uuid)
	if uuidNums != len(uuid) || err != nil {
		return
	}

	strUUID = hex.EncodeToString(uuid)
	return
}

func UtilsRandInt(min, max int) (intRand int) {
	rand.Seed(time.Now().Unix())
	intRand = rand.Intn(max-min) + min
	return
}