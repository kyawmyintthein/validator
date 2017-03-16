package validator

import(
    "github.com/kyawmyintthein/goerror"
)

type ValidatorFunc func(v interface{}) bool
type Validator struct {
    Name            string
    MessageTemplate string
    Data            map[string]interface{}
    Fn              ValidatorFunc
}

func NewValidator(name string, messageTmpl string, fn ValidatorFunc) Validator {
    return Validator{
        Name:            name,
        MessageTemplate: messageTmpl,
        Fn:              fn,
    }
}

func (v *Validator) Validate(i interface{}, data map[string]interface{}) goerror.Error {
    var err goerror.Error
    if v.Fn(i) {
        tmpl := goerror.NewTemplate(v.Name, v.MessageTemplate)
        return tmpl.Error(data)
    }
    return err
}

func (v *Validator) ValidateDefaultError(i interface{}, data map[string]interface{}) error {
    var err error
    if v.Fn(i) {
        tmpl := goerror.NewTemplate(v.Name, v.MessageTemplate)
        return tmpl.DefaultError(data)
    }
    return err
}