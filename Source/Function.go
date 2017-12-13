package main 

import (
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

var listmakanan []Food

var listminuman []Drink

var listcanteen []Canteen

// Method GET SPECIFIC

func GetMakanan(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range listmakanan {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Food{})
}

func GetMinuman(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range listminuman {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Food{})
}

func GetCanteen(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range listcanteen {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Food{})
}

//METHOD GET ALL
func GetAllMakanan(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(listmakanan)
}


func GetAllMinuman(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(listminuman)
}

func Getlistcanteen(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(listcanteen)
}


//METHOD POST
func AddMakanan(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
    var makanan Food
    _ = json.NewDecoder(req.Body).Decode(&makanan)
    makanan.ID = params["id"]
    makanan.Nama_makanan = params["Nama_makanan"]
    makanan.Harga = params["Harga"]
    listmakanan = append(listmakanan, makanan)
    json.NewEncoder(w).Encode(listmakanan)

}

func AddMinuman(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
    var minuman Drink
    _ = json.NewDecoder(req.Body).Decode(&minuman)
    minuman.ID = params["id"]
    minuman.Nama_minuman = params["Nama_minuman"]
    minuman.Harga = params["Harga"]
    listminuman = append(listminuman, minuman)
    json.NewEncoder(w).Encode(listminuman)
}

func AddCanteen(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
    var canteen Canteen
    _ = json.NewDecoder(req.Body).Decode(&canteen)
    canteen.ID = params["id"]
    canteen.Nama_canteen = params["Nama_canteen"]
    canteen.Location = params["Location"]
    listcanteen = append(listcanteen, canteen)
    json.NewEncoder(w).Encode(listcanteen)
}

//METHOD DELETE
func DeleteFood(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range listmakanan {
        if item.ID == params["id"] {
            listmakanan = append(listmakanan[:index], listmakanan[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(listmakanan)
}

func DeleteDrink(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range listminuman {
        if item.ID == params["id"] {
            listminuman = append(listminuman[:index], listminuman[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(listminuman)
}

func DeleteCanteen(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range listcanteen {
        if item.ID == params["id"] {
            listcanteen = append(listcanteen[:index], listcanteen[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(listcanteen)
}