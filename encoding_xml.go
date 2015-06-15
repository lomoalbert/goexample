package main
import (
    "encoding/xml"
    "log"
)
type Result struct {
    XMLName xml.Name `xml:"persons"`
    Persons []Person `xml:"person"`
}
type Person struct {
    Name string `xml:"name,attr"`
    Age int `xml:"age,attr"`
    Career string `xml:"career"`
    Interests []string `xml:"interests>interest"`
}
func main() {
    content := []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<persons>
<person name="polaris" age="28">
<career>无业游民</career>
<interests>
<interest>编程</interest>
<interest>下棋</interest>
</interests>
</person>
<person name="studygolang" age="27">
<career>码农</career>
<interests>
<interest>编程</interest>
<interest>下棋</interest>
</interests>
</person>
</persons>
    `)
    var result Result
    err := xml.Unmarshal(content, &result)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(result)
    log.Println(result.Persons[0].Name)
}