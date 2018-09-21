package main 

import ( 
				"net/http"
				"io/ioutil"
				"encoding/xml")

// Define the structure of the XML doc
type SitemapIndex struct {

	// Value  Slice/type  unmarshal tag (html tag)
	Locations []string `xml:"sitemap>loc"`
}

// Since location is just a string, this can be removed and placed in the SiteMapIndex struct
//type Location struct {
	// Value type   unmarshal tag (html tag)
// 	Loc string `xml:"loc"`
// }

// This can be removed as well after adding in loc from the Location struct
// Value receiver 
// func (l Location) String() string {
// 	return fmt.Sprintf(l.Loc)
// }

type News struct {
	Title []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func main() {
	s := SitemapIndex{}
	n := News{}

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml") // Use an underscore if you define any variable you don't intend to use
	bytes, _ := ioutil.ReadAll(resp.Body)
	// string_body := string(bytes)
	// fmt.Println(string_body)
	resp.Body.Close()


	xml.Unmarshal(bytes, &s)

	//fmt.Println(s.Locations) 

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		xml.Unmarshal(bytes, &n)
	}
}