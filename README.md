package main  
import (  
   "net/http"  
   "github.com/go-chi/chi"  
   "github.com/go-chi/cors"  
   "github.com/go-chi/render" 
)  
type Product struct {  
   ID        string   'json:"id,omitempty"'  
   Productname string   'json:"Productname,omitempty"'  
   Producttype  string   'json:"Producttype,omitempty"'  
   Price   *Price 'json:"Price,omitempty"'  
}  
type Price  struct {  
   Price string 'json:"Price,omitempty"'  
   Currency string 'json:"Currency,omitempty"'

}  
var Prod []Product  
func GetProdIdEndpoint(w http.ResponseWriter, req *http.Request) {  
   params := chi.Vars(req)  
   for _, item := range Prod {  
      if item.ID == params["id"] {  
         json.NewEncoder(w).Encode(item)  
         return  
      }  
   }  
   json.NewEncoder(w).Encode(&Product{})  
}  
func GetProductEndpoint(w http.ResponseWriter, req *http.Request) {  
   json.NewEncoder(w).Encode(Prod)  
}  
func CreateProductEndpoint(w http.ResponseWriter, req *http.Request) {  
   params := chi.Vars(req)  
   var Devices Product  
   _ = json.NewDecoder(req.Body).Decode(&Devices)  
   Devices.ID = params["id"]  
   Prod = append(Prod, Devices)  
   json.NewEncoder(w).Encode(Prod)  
}  
func DeleteProductEndpoint(w http.ResponseWriter, req *http.Request) {  
   params := chi.Vars(req)  
   for index, item := range Prod {  
      if item.ID == params["id"] {  
         Prod = append Prod[:index], Prod[index+1:]...)  
         break  
      }  
   }  
   json.NewEncoder(w).Encode(Prod)  
}  
func main() {  
   router := chi.NewRouter()  
   Prod = append(Prod, Product{ID: "1", Productname: "hpLaptop",Producttype : "Laptop",   
   Price: Price{Price: "$600", Currency: "USD"}})  
   Prod = append(Prod, Product{ID: "2", Productname: "Carpet", Producttype: "Household"})  
   router.HandleFunc("/Product", GetProductEndpoint).Methods("GET")  
   router.HandleFunc("/Product/{id}", GetProdIdEndpoint).Methods("GET")  
   router.HandleFunc("/Product/{id}", CreateProductEndpoint).Methods("POST")  
   router.HandleFunc("/Product/{id}", DeleteProductEndpoint).Methods("DELETE")  
   log.Fatal(http.ListenAndServe(":12345", router))  
}  


	


