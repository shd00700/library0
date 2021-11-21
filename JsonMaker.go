package Library

import (
	"encoding/json"
	"github.com/enitt-dev/go-utils/convert"
)

type ReadCoilStruct struct {
	Code_name     string
	Start_Address uint16
	Last_Address  uint16
}

type WriteCoilStruct struct {
	Code_name     string
	Start_Address uint16
	Last_Address  uint16
}

type ReadRegStruct struct {
	Code_name     string
	Start_Address uint16
	Last_Address  uint16
}

type WriteRegStruct struct {
	Code_name     string
	Start_Address uint16
	Last_Address  uint16
}

type ErrJson struct {
	Errormessage string
}

type ReadRegInStruct struct {
	Alias   string
	Address uint16
	Value   uint16
}

var ReadCoil_alias = []string{"Un_name0", "Un_name1", "Un_name2", "Un_name3", "Un_name4", "Un_name5", "Un_name6",
	"Un_name7", "Un_name8", "Un_name9", "Un_name10", "Un_name11", "Un_name12", "Un_name13", "Un_name14",
	"Un_name15", "Un_name16", "Un_name17", "Un_name18", "Un_name19", "Un_name20"}

var ReadCoilIn_alias = []string{"Un_name0", "Un_name1", "Un_name2", "Un_name3", "Un_name4", "Un_name5", "Un_name6",
	"Un_name7", "Un_name8", "Un_name9", "Un_name10", "Un_name11", "Un_name12", "Un_name13", "Un_name14",
	"Un_name15", "Un_name16", "Un_name17", "Un_name18", "Un_name19", "Un_name20"}

var ReadRegIn_alias = []string{"Un_name0", "Un_name1", "Un_name2", "Un_name3", "Un_name4", "Un_name5", "Un_name6",
	"Un_name7", "Un_name8", "Un_name9", "Un_name10", "Un_name11", "Un_name12", "Un_name13", "Un_name14",
	"Un_name15", "Un_name16", "Un_name17", "Un_name18", "Un_name19", "Un_name20"}

var ReadHoldReg_alias = []string{"Un_name0", "Un_name1", "Un_name2", "Un_name3", "Un_name4", "Un_name5", "Un_name6",
	"Un_name7", "Un_name8", "Un_name9", "Un_name10", "Un_name11", "Cn_name12", "Un_name13", "Un_name14",
	"Un_name15", "Un_name16", "Un_name17", "Un_name18", "Un_name19", "Un_name20"}

var SensorData = []string{"CO2", "Temperature", "Temperature_object", "Humidity"}

func ReadCoilJsonMaker(a uint16, b []int, leng uint16) interface{} {

	alias_arr := map[string]int{}
	println("aa", a, b, leng)

	for i := 0; i < int(leng); i++ {
		alias_arr[ReadCoil_alias[int(a)+i]] = b[i]
	}
	jsonmaker, err := json.Marshal(alias_arr)
	if err != nil {
		panic(err)
	}
	return jsonmaker
}

func ReadCoilInJsonMaker(a uint16, b []int, leng uint16) []byte {

	alias_arr := map[string]interface{}{}
	println("aa", a, b, leng)
	println("")

	for i := 0; i < int(leng); i++ {
		alias_arr[ReadCoil_alias[int(a)+i]] = b[i]
	}
	jsonmaker, err := json.Marshal(alias_arr)
	if err != nil {
		panic(err)
	}

	println("Jsonmaker Type:", jsonmaker)
	return jsonmaker
}

func ReadHoldRegJsonMaker(a uint16, b []uint16, leng uint16) interface{} {

	alias_arr := map[string]uint16{}
	println("aa", a, b, leng)

	for i := 0; i < int(leng); i++ {
		alias_arr[ReadHoldReg_alias[int(a)+i]] = b[i]
	}
	jsonmaker, err := json.Marshal(alias_arr)
	if err != nil {
		panic(err)
	}
	return jsonmaker
}

func ReadRegInJsonMaker(a uint16, b []uint16, leng uint16) interface{} {
	alias_arr := map[string]uint16{}

	println("aa", a, b, leng)

	for i := 0; i < int(leng); i++ {
		alias_arr[ReadRegIn_alias[int(a)+i]] = b[i]
	}

	jsonmaker, err := json.Marshal(alias_arr)
	if err != nil {
		panic(err)
	}
	return jsonmaker
}

func JsonMaker(bytes []byte) ([]byte, []byte) {
	sensordater := map[string]interface{}{}
	result := []map[string]interface{}{}

	for {
		if bytes[4] == 0x30 {
			sensordater[SensorData[int(0)+1]] = convert.BytesToFloat32(bytes[6:])
		}
		if bytes[4] == 0x10 {
			sensordater[SensorData[int(0)+0]] = convert.BytesToFloat32(bytes[6:])
		}
		if bytes[4] == 0x20 {
			sensordater[SensorData[int(0)+3]] = convert.BytesToFloat32(bytes[6:])
		}

	}

	result = append(result, sensordater)

	println("result:", result)
	jsonmaker, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	return jsonmaker
}
