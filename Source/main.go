package main

import (
    "fmt"
    "log"
    "github.com/gorilla/mux"
    "net/http"
)   

func controller(){
        //Handler untuk FoodFinder
        //Food
    router := mux.NewRouter()
    router.HandleFunc("/food", GetAllMakanan).Methods("GET")
    router.HandleFunc("/food/{id}", GetMakanan).Methods("GET")
    router.HandleFunc("/food/{id}/{Nama_makanan}:{Harga}/{Nama_canteen}", AddMakanan).Methods("POST")
    router.HandleFunc("/food/{id}", DeleteFood).Methods("DELETE")
        //Drink
    router.HandleFunc("/drink", GetAllMinuman).Methods("GET")
    router.HandleFunc("/drink/{id}", GetMinuman).Methods("GET")
    router.HandleFunc("/drink/{id}/{Nama_minuman}:{Harga}/{Nama_canteen}", AddMinuman).Methods("POST")
    router.HandleFunc("/drink/{id}", DeleteDrink).Methods("DELETE")
        //Canteen
    router.HandleFunc("/canteen", Getlistcanteen).Methods("GET")
    router.HandleFunc("/canteen/{id}", GetCanteen).Methods("GET")
    router.HandleFunc("/canteen/{id}/{Nama_canteen}:{Location}", AddCanteen).Methods("POST")
    router.HandleFunc("/canteen/{id}", DeleteCanteen).Methods("DELETE")

        //Memberikan Suggestion
    router.HandleFunc("/SuggestFood/{Harga}", SuggestFood).Methods("GET")
    router.HandleFunc("/SuggestDrink/{Harga}", SuggestDrink).Methods("GET")
    log.Fatal(http.ListenAndServe(":12345", router))
    return
}

func main() {
	fmt.Printf("server is up and running on port 12345")
    
    controller()

    
}

