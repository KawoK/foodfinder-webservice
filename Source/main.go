package main

import (
    
    "log"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "database/sql"

)

func main() {
	fmt.Printf("server is up and running on port 12345")
    db, err := sql.Open("mysql ",
        "root:root@tcp(127.0.0.1:3306)/FoodFinder")
    if err != nil {
        log.Fatal(err)
    }
    router := mux.NewRouter()
    listmakanan = append(listmakanan, Food{ID: "1", Nama_makanan: "Cireng", Harga :"3000"})
    listmakanan = append(listmakanan, Food{ID: "2", Nama_makanan: "Bakso", Harga :"4000"})
    listmakanan = append(listmakanan, Food{ID: "3", Nama_makanan: "Tahu", Harga :"5000"})
    
    listminuman = append(listminuman, Drink{ID: "1", Nama_minuman: "Thai-Tea", Harga :"3000"})
    listminuman = append(listminuman, Drink{ID: "2", Nama_minuman: "Aqua", Harga :"4000"})
    listminuman = append(listminuman, Drink{ID: "3", Nama_minuman: "Teh Botol", Harga :"5000"})

    listcanteen = append(listcanteen, Canteen{ID: "1", Nama_canteen: "Kantin Saraga", Location :"Saraga ITB"})
    listcanteen = append(listcanteen, Canteen{ID: "2", Nama_canteen: "Kantin GKU Barat", Location :"GKU Barat"})
    listcanteen = append(listcanteen, Canteen{ID: "3", Nama_canteen: "Etitu", Location :"Gedung CC Timur"})

    //Handler untuk FoodFinder
    	//Food
    router.HandleFunc("/food", GetAllMakanan).Methods("GET")
    router.HandleFunc("/food/{id}", GetMakanan).Methods("GET")
    router.HandleFunc("/food/{id}/{Nama_makanan}:{Harga}", AddMakanan).Methods("POST")
    router.HandleFunc("/food/{id}", DeleteFood).Methods("DELETE")
    	//Drink
    router.HandleFunc("/drink", GetAllMinuman).Methods("GET")
    router.HandleFunc("/drink/{id}", GetMinuman).Methods("GET")
    router.HandleFunc("/drink/{id}/{Nama_minuman}:{Harga}", AddMinuman).Methods("POST")
    router.HandleFunc("/drink/{id}", DeleteDrink).Methods("DELETE")
    	//Canteen
    router.HandleFunc("/canteen", Getlistcanteen).Methods("GET")
    router.HandleFunc("/canteen/{id}", GetCanteen).Methods("GET")
    router.HandleFunc("/canteen/{id}/{Nama_canteen}:{Location}", AddCanteen).Methods("POST")
    router.HandleFunc("/canteen/{id}", DeleteCanteen).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":12345", router))
    defer db.Close()
}