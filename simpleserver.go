package main

import (
    "fmt"
    "html"
    "io"
    "log"
    "net/http"
)


//function that searches the database for malware URLs
//should search the database using something like this: GET /urlinfo/1/{hostname_and_port}/{original_path_and_query_string}
func DataBaseValidation(url string) bool {
    return true //just getting one response for testing if code is working
}
// add a function to create a 2 flows (1 for malware URLs and one for ok URLs)
func urlinfo(w io.Writer,url string) {
    validUrl := DataBaseValidation(url)
    if(validUrl == true) {
         fmt.Fprintf(w,"URL IS GOOD")
    } else {
        fmt.Fprintf(w,"URL IS BAD")
    }
}
func main() {

//this is handling the urlinfo function to handle url requests
   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        urlinfo(w, html.EscapeString(r.URL.Path))
    })         

    log.Fatal(http.ListenAndServe(":8081", nil))

}