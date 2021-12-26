package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var yazi string
var kod int

func giris() {
	fmt.Println("ad giriniz:")
	fmt.Scan(&yazi)
}

func ekle() {
	vt, _ := sql.Open("sqlite3", "veritabani.db")
	işlem, _ := vt.Prepare("INSERT INTO kisiler(ad) values(?)")
	veri, _ := işlem.Exec(yazi)
	id, _ := veri.LastInsertId() //Son girişin id numarısını aldık
	fmt.Println("Son kişinin id'si", id)
	vt.Close() //İşimiz bittikten sonra veri tabanımızı kapatıyoruz
}

func tablo() {
	vt, _ := sql.Open("sqlite3", "veritabani.db")
	islem, _ := vt.Query("SELECT * FROM kisiler")

	var id int
	var name string

	for islem.Next() {
		aktarma := islem.Scan(&id, &name)
		if aktarma == nil {
			fmt.Println(id, name)
		}

	}

	islem.Close()
	vt.Close()
}

func guncelleme() {
	var id int
	var tekrar string
	vt, _ := sql.Open("sqlite3", "veritabani.db")
	islem, _ := vt.Prepare("UPDATE kisiler set ad=? where id=?")
	fmt.Println("id: ")
	fmt.Scan(&id)
	fmt.Println("Degisecek ad: ")
	fmt.Scan(&tekrar)
	veri, _ := islem.Exec(tekrar, id)
	if veri != nil {
		fmt.Println(id, tekrar)
	}

	vt.Close()
}

func silme() {
	var id int
	vt, _ := sql.Open("sqlite3", "veritabani.db")
	islem, _ := vt.Prepare("DELETE FROM kisiler where id=?")
	fmt.Println("silinece id: ")
	fmt.Scan(&id)
	veri, _ := islem.Exec(id)
	sil, _ := veri.RowsAffected()
	fmt.Println("Silinen id: ", sil)
	vt.Close()
}

func main() {
	for true {
		fmt.Println("Işlem Numarası Seçiniz.\n 1-Veri Ekleme \n 2-Güncel Tablo \n 3-Veri Güncelleme \n 4-Veri Silme \n  ")
		fmt.Scan(&kod)
		fmt.Println("***********************************************")
		if kod == 1 {
			giris()
			ekle()

			fmt.Println("***********************************************")
			tablo()
			fmt.Println("***********************************************")
		} else if kod == 2 {
			tablo()

		} else if kod == 3 {
			guncelleme()
			fmt.Println("***********************************************")
			tablo()
			fmt.Println("***********************************************")
		} else if kod == 4 {
			silme()
			fmt.Println("***********************************************")
			tablo()
			fmt.Println("***********************************************")

		}

		//Hangi bölüme eklenecekse yukarıda orayı belirtiyoruz
		//Eklenecek değer
	}
}
