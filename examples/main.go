package main

import "fmt"
import "validator"

type Animal struct {
    Name string
    Age  int
}

func main() {
    id := 0
    err := errtmpl.Required(id, "user ID")
    fmt.Println(reqErr.Error())

    data := map[string]interface{}{"attr": "id", "value": id}
    tmpl := validator.NewTemplate("ID required", "{{attr}} should no be blank. {{attr}} is {{value}}. Please check your data input.")
    err = errtmpl.RequiredWithTemplate(id, cusdata, tmpl)
    fmt.Println(err.Error())
}