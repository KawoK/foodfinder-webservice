// Nama 		: Ikhsan Widi Adyatma
// NIM 			: 18215047


package main

// Model data untuk webservice foodfinder 
type Food struct {
    ID        		string   	`json:"id,omitempty"`
    Nama_makanan 	string   	`json:"Nama_makanan,omitempty"`
    Harga  			string   	`json:"Harga,omitempty"`
    
}
type Drink struct {
    ID        		string   `json:"id,omitempty"`
    Nama_minuman 	string   `json:"Nama_minuman,omitempty"`
    Harga  			string   `json:"Harga,omitempty"`
    
}
type Canteen struct {
    ID       		string   `json:"id,omitempty"`
    Nama_canteen 	string   `json:"Nama_canteen,omitempty"`
    Location 		string   `json:"Location,omitempty"`
}


