package main

import ("fmt")

func main() {
	Pemilik := "Gus Samsudin" 
	Kereta  := "Mataram" 
	Perjalanan := "Jakarta-Surabaya"

	tiket := setKodeTiket("H1D1C9100", map[string]string{
		"pemilik" 	 : Pemilik,
		"kereta"  	 : Kereta,
		"perjalanan" : Perjalanan,
	})
	fmt.Println(tiket)
	
	ubahNama := ubahNamaPemilik("Doni Riyadi", tiket)
	fmt.Println(ubahNama)
}

func setKodeTiket(kodeTiket string, tiket map[string]string) map[string]string {
	tiket["kode_tiket"] = kodeTiket
	return cetakTiket(tiket)
}

func cetakTiket(tiket map[string]string) map[string]string {
	return tiket
}

func ubahNamaPemilik(nama string, tiket map[string]string) (tiketBaru map[string]string) {
	tiketBaru = tiket
	tiketBaru["pemilik"] = nama
	tiketBaru = cetakTiket(tiketBaru)
	return
}