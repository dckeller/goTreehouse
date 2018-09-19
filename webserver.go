package main 

import ("fmt" 
				"net/http")

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, go is neat!</h1>") // format based on what you specify
	fmt.Fprintf(w, "<p>Below is a pic of a panda:</p>")
	fmt.Fprintf(w, "<img width='300' src='https://www.thehindu.com/migration_catalog/article10107232.ece/alternates/FREE_960/pandas'>")
	fmt.Fprintf(w, "<p>Isn't he %s cute?<br/>", "<strong>SUPER</strong>")
	fmt.Fprintf(w, "<a href='http://localhost:8000/about/'>About</a>")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the About Page") // format based on what you specify
}

func main() {
	http.HandleFunc("/", index_handler) // root path to page
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil) // 2nd arg have to pass something
}



