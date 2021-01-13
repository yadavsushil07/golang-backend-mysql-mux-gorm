package main

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "fmt"
    "log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
)

type Address struct{
    gorm.Model
    Area     string  `json:"area"`
    City     string  `json:"city"`
    State    string  `json:"state"`
    Pincode  string  `json:"pincode"`
}

type User struct {
    gorm.Model
    Id          string `json:"id"`
    Name        string `json:"name"`
    Mobile      string `json:"phone no"`
    Address     *Address

}
 
type Admin struct{
    gorm.Model
    Id          string `json:"id"`
    Username    string `json:"username`
    Password    string `json:"password"`
}

type Registration struct{
    Id        string `json:"id"`
    Name      string `json:"name"`
    Password  string `json:"pasword"`
    Email     string `json:"email"`
}


var User []User


func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    w.Header().Set("Content-Type","application/json")
    json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    vars := mux.Vars(r)
    key := vars["id"]

    // Loop over all of our Articles
    // if the article.Id equals the key we pass in
    // return the article encoded as JSON
    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article 
    json.Unmarshal(reqBody, &article)

    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    // once again, we will need to parse the path parameters
    vars := mux.Vars(r)
    // we will need to extract the `id` of the article we
    // wish to delete
    id := vars["id"]

    // we then need to loop through all our articles
    for index, article := range Articles {
        // if our id path parameter matches one of our
        // articles
        if article.Id == id {
            // updates our Articles array to remove the 
            // article
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
	}
}



func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
   
    myRouter := mux.NewRouter().StrictSlash(true) 
    myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
    log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
    Articles = []Article{
        Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}