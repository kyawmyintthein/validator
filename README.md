# Validator
Golang validator with customizable error message template.

## Usage
``` data := map[string]interface{}{"attr": "id", "value": id}
    tmpl := goerror.NewTemplate("Required", "{{attr}} should no be blank. {{attr}} is {{value}}. Please check your data input.")
    err = validator.RequiredWithTemplate(id, data, tmpl)
    fmt.Println(err.Error())
```