package main

import "fmt"
import "validator"
import "github.com/kyawmyintthein/goerror"

type Animal struct {
    Name string
    Age  int
}

func main() {
    id := 0
    err := validator.Required(id, "user ID")
    fmt.Println(err.Error())

    data := map[string]interface{}{"attr": "id", "value": id}
    tmpl := goerror.NewTemplate("Required", "{{attr}} should no be blank. {{attr}} is {{value}}. Please check your data input.")
    err = validator.RequiredWithTemplate(id, data, tmpl)
    fmt.Println(err.Error())
}