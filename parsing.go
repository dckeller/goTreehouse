package main 

import ("fmt" 
				"net/http"
				"io/ioutil"
				"encoding/xml")

// Define the structure of the XML doc
type SitemapIndex struct {

	// Value  Slice/type  unmarshal tag (html tag)
	Locations []Location `xml:"sitemap"`
}

// 
type Location struct {
	// Value type   unmarshal tag (html tag)
	Loc string `xml:"loc"`
}

// Value receiver 
func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml") // Use an underscore if you define any variable you don't intend to use
	bytes, _ := ioutil.ReadAll(resp.Body)
	// string_body := string(bytes)
	// fmt.Println(string_body)
	resp.Body.Close()


	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations) 
}