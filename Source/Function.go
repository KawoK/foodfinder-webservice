package main 

import (
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

var listmakanan []Food

var listminuman []Drink

var listcanteen []Canteen

// Method GET SPECIFIC

func GetMakanan(w http.ResponseWriter, req *http.Request) {
    //Membuka koneksi ke database "Foodfinder"
    params := mux.Vars(req)
    dbms, err := sql.Open("mysql",
              "root:root@tcp(127.0.0.1:3306)/Foodfinder")
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer dbms.Close()
    //Inisialisasi data makanan
    makanan := Food{}
    //Query untuk Mendapatkan seluruh data makanan
    baris, err := dbms.Query("select * from Food where id = ?",params["id"])
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer baris.Close()
    //Scanning Rows
    for baris.Next(){
        err:= baris.Scan(&makanan.ID, &makanan.Nama_makanan, &makanan.Harga, &makanan.Nama_canteen)
        //Error Handling
        if err != nil {
            log.Fatal(err)
        }
        json.NewEncoder(w).Encode(&makanan)
    }
    err=baris.Err()
    if err != nil {
        log.Fatal(err)
    }
}

func GetMinuman(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
     //Membuka koneksi ke database "Foodfinder"
    dbms, err := sql.Open("mysql",
              "root:root@tcp(127.0.0.1:3306)/Foodfinder")
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer dbms.Close()
    //Inisialisasi data makanan
    minuman := Drink{}
    //Query untuk Mendapatkan seluruh data makanan
    baris, err := dbms.Query("select * from Drink where id = ?",params["id"])
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer baris.Close()
    //Scanning Rows
    for baris.Next(){
        err:= baris.Scan(&minuman.ID, &minuman.Nama_minuman, &minuman.Harga, &minuman.Nama_canteen)
        //Error Handling
        if err != nil {
            log.Fatal(err)
        }
        json.NewEncoder(w).Encode(&minuman)
    }
    err=baris.Err()
    if err != nil {
        log.Fatal(err)
    }
}

func GetCanteen(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
   //Membuka koneksi ke database "Foodfinder"
    dbms, err := sql.Open("mysql",
              "root:root@tcp(127.0.0.1:3306)/Foodfinder")
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer dbms.Close()
    //Inisialisasi data makanan
    kantin := Canteen{}
    //Query untuk Mendapatkan seluruh data makanan
    baris, err := dbms.Query("select * from Canteen where id= ?",params["id"])
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer baris.Close()
    //Scanning Rows
    for baris.Next(){
        err:= baris.Scan(&kantin.ID, &kantin.Nama_canteen, &kantin.Location)
        //Error Handling
        if err != nil {
            log.Fatal(err)
        }
        json.NewEncoder(w).Encode(&kantin)
    }
    err=baris.Err()
    if err != nil {
        log.Fatal(err)
    }
}

//METHOD GET ALL
func GetAllMakanan(w http.ResponseWriter, req *http.Request) {
    //Membuka koneksi ke database "Foodfinder"
    dbms, err := sql.Open("mysql",
              "root:root@tcp(127.0.0.1:3306)/Foodfinder")
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer dbms.Close()
    //Inisialisasi data makanan
    makanan := Food{}
    //Query untuk Mendapatkan seluruh data makanan
    baris, err := dbms.Query("select * from Food")
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer baris.Close()
    //Scanning Rows
    for baris.Next(){
        err:= baris.Scan(&makanan.ID, &makanan.Nama_makanan, &makanan.Harga, &makanan.Nama_canteen)
        //Error Handling
        if err != nil {
            log.Fatal(err)
        }
        json.NewEncoder(w).Encode(&makanan)
    }
    err=baris.Err()
    if err != nil {
        log.Fatal(err)
    }
}


func GetAllMinuman(w http.ResponseWriter, req *http.Request) {
    //Membuka koneksi ke database "Foodfinder"
    dbms, err := sql.Open("mysql",
              "root:root@tcp(127.0.0.1:3306)/Foodfinder")
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer dbms.Close()
    //Inisialisasi data makanan
    minuman := Drink{}
    //Query untuk Mendapatkan seluruh data makanan
    baris, err := dbms.Query("select * from Drink")
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer baris.Close()
    //Scanning Rows
    for baris.Next(){
        err:= baris.Scan(&minuman.ID, &minuman.Nama_minuman, &minuman.Harga, &minuman.Nama_canteen)
        //Error Handling
        if err != nil {
            log.Fatal(err)
        }
        json.NewEncoder(w).Encode(&minuman)
    }
    err=baris.Err()
    if err != nil {
        log.Fatal(err)
    }
}

func Getlistcanteen(w http.ResponseWriter, req *http.Request) {
    //Membuka koneksi ke database "Foodfinder"
    dbms, err := sql.Open("mysql",
              "root:root@tcp(127.0.0.1:3306)/Foodfinder")
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer dbms.Close()
    //Inisialisasi data makanan
    kantin := Canteen{}
    //Query untuk Mendapatkan seluruh data makanan
    baris, err := dbms.Query("select * from Canteen")
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer baris.Close()
    //Scanning Rows
    for baris.Next(){
        err:= baris.Scan(&kantin.ID, &kantin.Nama_canteen, &kantin.Location)
        //Error Handling
        if err != nil {
            log.Fatal(err)
        }
        json.NewEncoder(w).Encode(&kantin)
    }
    err=baris.Err()
    if err != nil {
        log.Fatal(err)
    }
}


//METHOD POST
func AddMakanan(w http.ResponseWriter, req *http.Request){
    params := mux.Vars(req)
    var makananbaru Food
    _ = json.NewDecoder(req.Body).Decode(&makananbaru)
    makananbaru.ID = params["id"]
    makananbaru.Nama_makanan = params["Nama_makanan"]
    makananbaru.Harga = params["Harga"]
    makananbaru.Nama_canteen = params["Nama_canteen"]
    dbms, err := sql.Open("mysql",
              "root:root@tcp(127.0.0.1:3306)/Foodfinder")
    //Error Handling
    if err != nil{
        log.Fatal(err)
    }
    //Query untuk melakukan insert into tabel food 
    baris,err := dbms.Prepare("insert into Food (ID,Nama_makanan, Harga,Nama_canteen) VALUES (?,?,?,?)")
    if err != nil {
        log.Fatal(err)
    }
    _,err = baris.Exec(makananbaru.ID, makananbaru.Nama_makanan, makananbaru.Harga, makananbaru.Nama_canteen)
}

func AddMinuman(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var minumanbaru Drink
    _ = json.NewDecoder(req.Body).Decode(&minumanbaru)
    minumanbaru.ID = params["id"]
    minumanbaru.Nama_minuman = params["Nama_minuman"]
    minumanbaru.Harga = params["Harga"]
    minumanbaru.Nama_canteen = params["Nama_canteen"]
    dbms, err := sql.Open("mysql",
              "root:root@tcp(127.0.0.1:3306)/Foodfinder")
    //Error Handling
    if err != nil{
        log.Fatal(err)
    }
    //Query untuk melakukan insert into tabel drink 
    baris,err := dbms.Prepare("insert into Drink (ID,Nama_minuman, Harga,Nama_canteen) VALUES (?,?,?,?)")
    if err != nil {
        log.Fatal(err)
    }
     _,err = baris.Exec(minumanbaru.ID, minumanbaru.Nama_minuman, minumanbaru.Harga, minumanbaru.Nama_canteen)
}

func AddCanteen(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
    var kantinbaru Canteen
    _ = json.NewDecoder(req.Body).Decode(&kantinbaru)
    kantinbaru.ID = params["id"]
    kantinbaru.Nama_canteen = params["Nama_kantin"]
    kantinbaru.Location = params["Location"]
    dbms, err := sql.Open("mysql",
              "root:root@tcp(127.0.0.1:3306)/Foodfinder")
    //Error Handling
    if err != nil{
        log.Fatal(err)
    }
    //Query untuk melakukan insert into tabel canteen
    baris,err := dbms.Prepare("insert into Canteen (ID,Nama_canteen, Location) VALUES (?,?,?)")
    if err != nil {
        log.Fatal(err)
    }
     _,err = baris.Exec(kantinbaru.ID, kantinbaru.Nama_canteen, kantinbaru.Location)
}

//METHOD DELETE
func DeleteFood(w http.ResponseWriter, req *http.Request) {
    //Membuka koneksi ke database "Foodfinder"
    params := mux.Vars(req)
    dbms, err := sql.Open("mysql",
              "root:root@tcp(127.0.0.1:3306)/Foodfinder")
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer dbms.Close()
    //Inisialisasi data makanan
    makanan := Food{}
    //Query untuk Mendapatkan seluruh data makanan
    baris, err := dbms.Query("DELETE from Food where id = ?",params["id"])
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer baris.Close()
    //Scanning Rows
    for baris.Next(){
        err:= baris.Scan(&makanan.ID, &makanan.Nama_makanan, &makanan.Harga, &makanan.Nama_canteen)
        //Error Handling
        if err != nil {
            log.Fatal(err)
        }
        json.NewEncoder(w).Encode(&makanan)
    }
    err=baris.Err()
    if err != nil {
        log.Fatal(err)
    }
}

func DeleteDrink(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
     //Membuka koneksi ke database "Foodfinder"
    dbms, err := sql.Open("mysql",
              "root:root@tcp(127.0.0.1:3306)/Foodfinder")
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer dbms.Close()
    //Inisialisasi data makanan
    minuman := Drink{}
    //Query untuk Mendapatkan seluruh data makanan
    baris, err := dbms.Query("DELETE from Drink where id = ?",params["id"])
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer baris.Close()
    //Scanning Rows
    for baris.Next(){
        err:= baris.Scan(&minuman.ID, &minuman.Nama_minuman, &minuman.Harga, &minuman.Nama_canteen)
        //Error Handling
        if err != nil {
            log.Fatal(err)
        }
        json.NewEncoder(w).Encode(&minuman)
    }
    err=baris.Err()
    if err != nil {
        log.Fatal(err)
    }
}

func DeleteCanteen(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
   //Membuka koneksi ke database "Foodfinder"
    dbms, err := sql.Open("mysql",
              "root:root@tcp(127.0.0.1:3306)/Foodfinder")
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer dbms.Close()
    //Inisialisasi data makanan
    kantin := Canteen{}
    //Query untuk Mendapatkan seluruh data makanan
    baris, err := dbms.Query("DELETE from Canteen where id= ?",params["id"])
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer baris.Close()
    //Scanning Rows
    for baris.Next(){
        err:= baris.Scan(&kantin.ID, &kantin.Nama_canteen, &kantin.Location)
        //Error Handling
        if err != nil {
            log.Fatal(err)
        }
        json.NewEncoder(w).Encode(&kantin)
    }
    err=baris.Err()
    if err != nil {
        log.Fatal(err)
    }
}

//Suggestion Method
func SuggestFood(w http.ResponseWriter, req *http.Request) {
    //Membuka koneksi ke database "Foodfinder"
    params := mux.Vars(req)
    dbms, err := sql.Open("mysql",
              "root:root@tcp(127.0.0.1:3306)/Foodfinder")
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer dbms.Close()
    //Inisialisasi data makanan
    makanan := Food{}
    //Query untuk Mendapatkan seluruh data makanan
    baris, err := dbms.Query("select * from Food where Harga <= ?",params["Harga"])
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer baris.Close()
    //Scanning Rows
    for baris.Next(){
        err:= baris.Scan(&makanan.ID, &makanan.Nama_makanan, &makanan.Harga, &makanan.Nama_canteen)
        //Error Handling
        if err != nil {
            log.Fatal(err)
        }
        json.NewEncoder(w).Encode(&makanan)
    }
    err=baris.Err()
    if err != nil {
        log.Fatal(err)
    }
}

func SuggestDrink(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
     //Membuka koneksi ke database "Foodfinder"
    dbms, err := sql.Open("mysql",
              "root:root@tcp(127.0.0.1:3306)/Foodfinder")
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer dbms.Close()
    //Inisialisasi data makanan
    minuman := Drink{}
    //Query untuk Mendapatkan seluruh data makanan
    baris, err := dbms.Query("select * from Drink where Harga <= ?",params["Harga"])
    //Error Handling
    if err != nil {
        log.Fatal(err)
    }
    defer baris.Close()
    //Scanning Rows
    for baris.Next(){
        err:= baris.Scan(&minuman.ID, &minuman.Nama_minuman, &minuman.Harga, &minuman.Nama_canteen)
        //Error Handling
        if err != nil {
            log.Fatal(err)
        }
        json.NewEncoder(w).Encode(&minuman)
    }
    err=baris.Err()
    if err != nil {
        log.Fatal(err)
    }
}