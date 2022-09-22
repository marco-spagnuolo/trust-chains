package main

import (
	"encoding/json"
	f "go2/functions"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Asset struct {
	Issuer         string `json:"issuer"`
	ID             string `json:"id"`
	CustomerID     string `json:"customerID"`
	Description    string `json:"description"`
	Key            string `json:"key"`
	ProductType    string `json:"productType"`
	SwltID         string `json:"swltID"`
	ActivationTime string `json:"activationTime"`
	ExpirationTime string `json:"expirationTime"`
	TotalCapacity  string `json:"totalCapacity"`
	UsedCapacity   string `json:"usedCapacity"`
	Owner          string `json:"owner"`
}
type UserPost struct {
	ID string `json:"id"`
}

var assets []Asset = []Asset{}
var userposts []UserPost

func main() {
	idUser := os.Args[1]

	log.Println("============ application-golang starts ============")

	log.Println("ID user:", idUser)

	router := mux.NewRouter()

	//licenses/cumulative

	router.HandleFunc("/appchain/licenses/cumulative/create", createItem).Methods("POST")
	router.HandleFunc("/appchain/licenses/cumulative/update/{id}", updateItem).Methods("POST")
	router.HandleFunc("/appchain/licenses/cumulative/transfer/{id}", transferItem).Methods("POST")
	router.HandleFunc("/appchain/licenses/cumulative/getall", getAllItem).Methods("GET")
	router.HandleFunc("/appchain/licenses/cumulative/get/{id}", getIDItem).Methods("GET")

	//idm

	router.HandleFunc("/appchain/idm/add", addUser).Methods("POST")
	router.HandleFunc("/appchain/idm/{id}/remove", removeUser).Methods("DELETE")
	router.HandleFunc("/appchain/idm/list", getAllWallet).Methods("GET")

	err := http.ListenAndServe(":8081", router)
	if err != nil {
		log.Println("ListenAndServe FAIL")
	}
	log.Println("============ application-golang ends ============")

	log.Println("============ application-golang ends ============")
}

//LICENSES
//POST LOGIC ok
func createItem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	log.Println("CREATE START")

	var newAsset Asset

	json.NewDecoder(r.Body).Decode(&newAsset)
	//prendi valori dal post e passali alla funzione
	assets = append(assets, newAsset)
	str, _ := json.MarshalIndent(newAsset, "", "")
	log.Println("ResultJSON: ", string(str))
	json.NewEncoder(w).Encode(&newAsset)

	_, err := f.Create(os.Args[1], newAsset.Issuer, newAsset.ID,
		newAsset.CustomerID, newAsset.Description,
		newAsset.Key, newAsset.ProductType,
		newAsset.SwltID, newAsset.ActivationTime,
		newAsset.ExpirationTime, newAsset.TotalCapacity,
		newAsset.UsedCapacity, newAsset.Owner)
	if err != nil {
		log.Println("CREATE FAIL")
	}
	log.Println("CREATE OK")

}

//POST LOGIC ok
func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("TRANSFER START")
	params := mux.Vars(r)
	for _, item := range assets {
		if item.ID == params["id"] {
			//assets = append(assets[:index], assets[index+1:]...)

			var newAsset Asset
			_ = json.NewDecoder(r.Body).Decode(&newAsset)
			newAsset.ID = params["id"]

			_, err := f.Update(os.Args[1], newAsset.Issuer, newAsset.ID,
				newAsset.CustomerID, newAsset.Description,
				newAsset.Key, newAsset.ProductType,
				newAsset.SwltID, newAsset.ActivationTime,
				newAsset.ExpirationTime, newAsset.TotalCapacity,
				newAsset.UsedCapacity, newAsset.Owner)

			if err != nil {
				log.Println("UPDATE FAIL")

			}
			str, _ := json.MarshalIndent(newAsset, "", "")
			log.Println("ResultJSON: ", string(str))

			assets = append(assets, newAsset)
			json.NewEncoder(w).Encode(&newAsset)
			log.Println("UPDATE OK")
			return
		}
	}
	json.NewEncoder(w).Encode(assets)

}

//POST LOGIC ok
func transferItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("TRANSFER START")
	params := mux.Vars(r)
	for index, item := range assets {
		if item.ID == params["id"] {
			assets = append(assets[:index], assets[index+1:]...)

			var newAsset Asset
			_ = json.NewDecoder(r.Body).Decode(&newAsset)
			newAsset.ID = params["id"]

			_, err := f.Transfer(os.Args[1], newAsset.ID, newAsset.Owner)
			if err != nil {
				log.Println("TRANSFER FAIL")

			}
			str, _ := json.MarshalIndent(newAsset, "", "")
			log.Println("ResultJSON: ", string(str))

			assets = append(assets, newAsset)
			json.NewEncoder(w).Encode(&newAsset)
			log.Println("TRANSFER OK")
			return
		}
	}
	json.NewEncoder(w).Encode(assets)

}

//GET LOGIC
func getIDItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("GET START")
	params := mux.Vars(r)
	for _, item := range assets {
		if item.ID == params["id"] {
			id := params["id"]

			f.GetItemFromId(os.Args[1], id)
			json.NewEncoder(w).Encode(item)
			str, _ := json.MarshalIndent(item, "", "")
			log.Println("GET ResultJSON: ", string(str))
			break
		}
		id := params["id"]
		log.Println("GET FAIL FOR ID :", id)
		return
	}
	json.NewEncoder(w).Encode(&Asset{})
}

//GET LOGIC
func getAllItem(w http.ResponseWriter, r *http.Request) {
	log.Println("GET ALL START ")

	w.Header().Set("Content-Type", "application/json")

	f.QueryAll(os.Args[1])

	json.NewEncoder(w).Encode(assets)

	//log.Println("RES", string(result))

	str, _ := json.MarshalIndent(assets, "", "")
	log.Println("ResultJSON: ", string(str))

}

//POST LOGIC
func addUser(w http.ResponseWriter, r *http.Request) {

	log.Println("ADD USER START")
	w.Header().Set("Content-Type", "application/json")

	var userpost UserPost

	_ = json.NewDecoder(r.Body).Decode(&userpost)
	f.AddToWallet(userpost.ID)
	userposts = append(userposts, userpost)
	json.NewEncoder(w).Encode(&userposts)

	log.Printf("USER ID: %v ADDED", userpost.ID)

}

func removeUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("REMOVE USER START")

	params := mux.Vars(r)
	for index, item := range userposts {
		if item.ID == params["id"] {
			id := params["id"]
			userposts = append(userposts[:index], userposts[index+1:]...)
			f.RemoveFromWallet(id)
			log.Printf("USER ID: %v REMOVED", id)
			break
		}
	}
	json.NewEncoder(w).Encode(userposts)
}

func getAllWallet(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	f.ListAll()
	json.NewEncoder(w).Encode(userposts)

	str, _ := json.MarshalIndent(userposts, "", "")
	log.Println("ResultJSON ID WALLET: ", string(str))
}
