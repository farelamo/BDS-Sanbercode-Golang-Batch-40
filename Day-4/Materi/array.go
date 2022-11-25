package main

import (
	"fmt"
)

func main(){
						/************* Array *************/
	// Init Array
	var names [4] string 
	names[0] = "Farlam"
	names[1] = "Fariz"
	names[2] = "Fahmi"
	names[3] = "Faturachmad"

	fmt.Println(names[0], names[1], names[2], names[3])

	// Init Array Langsung
	var fruits = [4] string{"Banana", "Semongko", "Pencit", "Extra Joss"}
	fmt.Println(fruits)

	// Init Array tanpa jumlah data / elemen
	var years = [...] int {2017, 2018, 2019, 2020, 2021}
	fmt.Println(years)

	// Array Multidimensi
	var numbers1 = [2][3]int{[3]int{1,2,3}, [3]int{4,5,6}}
	fmt.Println(numbers1)

	// u/ sub element array, bisa tdk dituliskan jumlah datanya
	var numbers2 = [2][3]int{{7,8,9},{10,11,12}}
	fmt.Println(numbers2)
	
							/************* Slice *************/
	/* Perbedaan slice dg array ()
		cuma beda di init jumlah datanya.
		Array harus init jumlah datanya, sedangkan slice tidak pakai
	*/

	var sports = [] string {"football", "basketball", "baseball"}
	fmt.Println(sports)

	/* Hubungan Array dg Slice & Operasi Slice
		Array dan Slice merupakan satu kesatuan,
		Array adalah nilai atau elemennya, sedangkan Slice sebagai refrence dari nilai / elemen tsb
	*/

	var people 		= [] string {"Arman", "Hakim", "Lestaluhu", "Desta"}
	var newPeople 	= people[0:2]
	fmt.Println(newPeople) // ["Arman", "Hakim"]
	
	// diatas adalah cara pengambilan nilai di dalam slice menggunakan teknik 2 Index
	// dimulai dari index ke 0 sampai sebelum index ke 2. 
	// Hasil yg didapat yaitu index ke 0 dan index ke 1

	/* Perbedaan pengambilan / get element dalam Array & Slice
		Array menggunakan people[0], people[1], dst..
		hal ini akan mengcopy dari data element aslinya

		Beda dengan Slice, yang menggunakan people[0:2] akan mengambil langsung dari
		data elementnya (tidak dicopy). 
		Nilai yang didapat adalah refrence dari data asli elementnya
	*/

	// List table get element slice
	// people[0:4] => ["Arman", "Hakim", "Lestaluhu", "Desta"]	get element start index 0 s/d sebelum index 4
	// people[0:0] => []										menghasilkan array kosong, karena tdk ada data start 0 dan sebelum 0
	// people[4:4] => []										menghasilkan array kosong, karena tdk ada data start 4 dan sebelum 4
	// people[4:0] => error										people[a:b], a harus lebih kecil dari b
	// people[:]   => ["Arman", "Hakim", "Lestaluhu", "Desta"]	[:] akan get semua data
	// people[2:]  => ["Lestaluhu", "Desta"]					get data mulai dari index ke 2
	// people[:2]  => ["Arman", "Hakim"]						get data sebelum index ke 2

	// Pembuktian Refrence pada Slice
	var newPeople2 = newPeople[0:1]
	newPeople[0]   = "AGUS" //data pertama newPeople2 diubah jadi AGUS
	
	fmt.Println(newPeople2) // ["AGUS"]
	fmt.Println(newPeople)  // ["AGUS", "Hakim"]
	fmt.Println(people) 	// ["AGUS", "Hakim", "Lestaluhu", "Desta"]

	// Data dari slice sebelumnya akan ikut berubah, hal ini menunjukkan bahwa data element
	// yang diambil dari slice akan merefrence dari data element aslinya (tidak dicopy)

	/* ======== Function pada Slice ======== */

	// 1. len()
	/* Untuk Menghitung jumlah data dalam slice */
	var lol = [] string {"haha", "hihi"}
	fmt.Println(len(lol)) // 2 

	// 2. cap()
	/* untuk meghitung lebar atau kapasitas maksimum dari suatu slice */

	var bms = []string{"Mas Pitra", "Mas Yogs", "Mas BingBong", "Mas Ndi",}
	fmt.Println(len(bms)) // 4
	fmt.Println(cap(bms)) // 4

	var aBms = bms[0:3]
	fmt.Println(len(aBms)) // 3
	fmt.Println(cap(aBms)) // 4

	var bBms = bms[1:4]
	fmt.Println(len(bBms)) // 3
	fmt.Println(cap(bBms)) // 3

	/* 
		aBms => index ke 0,1,2
		bBms => index ke 1,2,3

		FAQ 	: Kenapa aBms bisa 3 dan bBms lebarnya jadi 4 ??
		Answer 	: 

			Singkatnya, slicing pada aBms [0:3] akan mengambil data dari
			index ke 0 s/d sebelum index ke 3 yaitu index ke 2. 
			Sehingga, tersisa index ke 3 hal ini lah yang membuat lebar menjadi 4 
			karena tersisa 1 index 

				cap(aBms) => ["Mas Pitra", "Mas Yogs", "Mas BingBong", ....]

			Sedangkan, bBms dimulai dari index ke 1 sampai sebelum index ke 4.
			Hal ini akan mengambil data dari index ke 1 s/d index ke 3
			karna sudah tidak ada data lagi yang tersisa setelah index ke 3, 
			maka, lebar nya hanya 3 saja

				cap(bBms) => ...,["Mas Yogs", "Mas BingBong", "Mas Ndi"]
			
			intinya, hasil lebar / cap dari slice dimulai dari start index dan
			ditambah sesuai dengan sisa data yang ada pada slice aslinya

			for ex, 
				var bms = []string{"Mas Pitra", "Mas Yogs", "Mas BingBong", "Mas Ndi", "ole"}

				len(bms) => 5 
				cap(bms) => 5
				
				cBms = bms[0:3]
				cap(cBms) => 5 (
					karena start dimulai dari index ke 0 dan menyisakan 2 data.
					Yaitu "Mas Ndi", "ole"
				)

				dBms = [1:4]
				cap(dBms) => 4 (
					karena start dimulai dari index ke 1 dan menyisakan 1 data.
					Yaitu "ole"
				)
	*/

	// 3. append()
	//  u/ menambahkan element kedalam array slice setelah index paling akhir

	var words 		= []string{"anjay", "lol", "anjir", "kamu tanyea"}
	var newWords	= append(words, "slebew")
	fmt.Println(newWords)

	// 2 hal yang perlu diketahui, 
	//  - ketika len(words) === cap(words), akan mereturn refrence baru *contoh seperti newWords
	//  - ketika cap(words) > len(words) atau len(words) < cap(words)
	//  maka akan mengubah nilai dari array yang direfrence

	var sliceWords    = words[0:2] // words adalah array yang direfrence
	var newSliceWords = append(sliceWords, "wkwk")
	fmt.Println(newSliceWords) // [anjay, lol, wkwk]
	fmt.Println(words) // will be [anjay, lol, wkwk, kamu tanyea]
	

	// 4. copy()
	// u/ mengcopy element / data pada slice2 ke slice 1
	// copy(slice1, slice2)

	dst := []string{"potato", "potato", "potato"}
	src := []string{"watermelon", "pinnaple"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple potato
	fmt.Println(src) // watermelon pinnaple
	fmt.Println(n)   // 2

	/* ======== End Function pada Slice ======== */

	// Akses elemen / data pada slice dengan 3 index
	    // u/ akses element sekaligus menentukan kapasitasnya
		// fruits[0:1:1] dg aturan angka kapasitas(paling belakang),
		// tdk bole melebihi kapasitas slice nya (tengah)

		var holaFruits = []string{"apple", "grape", "banana"}
		var eFruits = holaFruits[0:2]
		var fFruits = holaFruits[0:2:2]

		fmt.Println(holaFruits)      // ["apple", "grape", "banana"]
		fmt.Println(len(holaFruits)) // len: 3
		fmt.Println(cap(holaFruits)) // cap: 3

		fmt.Println(eFruits)      // ["apple", "grape"]
		fmt.Println(len(eFruits)) // len: 2
		fmt.Println(cap(eFruits)) // cap: 3

		fmt.Println(fFruits)      // ["apple", "grape"]
		fmt.Println(len(fFruits)) // len: 2
		fmt.Println(cap(fFruits)) // cap: 2

	// bisa dibilang teknik 3 index handle dari teknik index 2 

	// Alokasi Element Slice dg keyword make
	var hehe = make([]int, 2) // (tipe data, jumlah data)
	hehe[0]  = 4
	hehe[1]  = 5
    // hehe[2]  = 6 bakal error karena cuma berkapasitas 2
	
							/************* Map *************/
	/*
		berbentuk key-value
		key bersifat unik, karena u. identifier
	*/

	var test = map[string]int{}
	test["halo"] = 2
	test["hehe"] = 3

	fmt.Println(test["halo"]) // 2
	fmt.Println(test["hehe"]) // 3

	
}