package main

import (
	"fmt"
	"encoding/json"
)

func main(){
	// Decode JSON
	/* Decode JSON bisa dikonversi ke bentuk object struct atau map[string]interface{} */

	/********* DECODE OBJECT JSON ke OBJECT STRUCT *********/
	// Struct penampung data decode harus dibuat publik
	type User struct { 
		Fullname string `json:"Name"` // json: "Name" u/ tag pada key di json
		Age int // tdk perlu tag, karna udah sama dengan key di decode
	}

	var jsonString = `{"Name": "John Doe", "Age": 27}` //Key Name menjadi tag di Struct
	var jsonData   = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data) 
	// fungsi Unmarshal hanya menerima data bentuk []byte
	// maka dari itu jsonString diatas, dicasting dahulu ke []byte
	// untuk parameter ke 2 Unmarshal harus diisi dengan pointer untuk menampung hasil decode
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("========= OBEJCT STRUCT =========")
	fmt.Println("user :", data.Fullname)
	fmt.Println("age  :", data.Age)

	/********* DECODE JSON ke MAP[STRING] INTERFACE{} & INTERFACE{} *********/
	// Dengan map[string]interface{}

	var data1 map[string]interface{}

	fmt.Println("========= MAP 	INTERFACE KOSONG =========")
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data1["Name"])
	fmt.Println("age  :", data1["Age"])

	// Dengan interface{}
	// Syaratnya harus dicasting ke map[string]interface saat pengaksesannya

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodedData = data2.(map[string]interface{})

	fmt.Println("========= INTERFACE KOSONG =========")
	fmt.Println("user :", decodedData["Name"])
	fmt.Println("age  :", decodedData["Age"])

	/********* DECODE ARRAY JSON ke ARRAY OBJECT *********/
	// decode tipe array bisa ke slice/map ya ges proses sama kek yang object og

	var jsonArray = `[
		{"Name": "Zaskia", "Age": 20},
		{"Name": "Azka", "Age": 18}
	]`

	var dataDecode []User // init slice dengan tipe data user

	var error = json.Unmarshal([]byte(jsonArray), &dataDecode)
	if error != nil {
		fmt.Println(error.Error())
		return
	}

	fmt.Println("========= DECODE ARRAY =========")
	fmt.Println("User 1 : ", dataDecode[0].Fullname)
	fmt.Println("User 2 : ", dataDecode[1].Fullname)

	/********* ENCODE OBJECT ke JSON STRING *********/
	var object = []User{{"jhon doe", 20}, {"karima", 28}}

	var convertJson, errorConvert = json.Marshal(object)
	if errorConvert != nil {
		fmt.Println(errorConvert.Error())
		return
	}

	fmt.Println("========= ENCODE JSON =========")
	fmt.Println(string(convertJson))
}

