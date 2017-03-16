package validator

import(
    "github.com/kyawmyintthein/goerror"
)

const (
    DEFAULT_REQUIRED_ERROR_MESSAGE_TEMPLATE = "{{attr}} is required."
)

/*
 * name   : Required
 * desc   : Validate empty or null value in golang type such as int, string, float, struct
 * return : Return custom error type TError
 */
func Required(i interface{}, attr string) goerror.Error {
    requiredValidator := NewValidator("Required", DEFAULT_REQUIRED_ERROR_MESSAGE_TEMPLATE, func(v interface{}) bool {
        return isEmpty(v)
    })
    data := map[string]interface{}{"attr": attr}
    return requiredValidator.Validate(i, data)
}

/*
 * name   : RequiredDefaultError
 * desc   : Validate empty or null value in golang type such as int, string, float, struct
 * return : Return custom error type TError
 */
func RequiredDefaultError(i interface{}, attr string) error {
    requiredValidator := NewValidator("Required", DEFAULT_REQUIRED_ERROR_MESSAGE_TEMPLATE, func(v interface{}) bool {
        return isEmpty(v)
    })

    data := map[string]interface{}{"attr": attr}
    return requiredValidator.ValidateDefaultError(i, data)
}

/*
 * name   : Required
 * desc   : Validate empty or null value in golang type such as int, string, float, struct
 * return : Return custom error type TError
 */
func RequiredDefaultErrorWithTemplate(i interface{}, data map[string]interface{}, tmpl goerror.ErrorTemplate) error {
    requiredValidator := NewValidator(tmpl.Name, tmpl.Layout, func(v interface{}) bool { return isEmpty(v) })
    return requiredValidator.ValidateDefaultError(i, data)
}

/*
 * name   : RequiredWithCustomTemplate
 * desc   : Validate empty or null value in golang type such as int, string, float, struct with custom *          message template
 * return : Return custom error type TError
 */
func RequiredWithTemplate(i interface{}, data map[string]interface{}, tmpl goerror.ErrorTemplate) goerror.Error {
    requiredValidator := NewValidator(tmpl.Name, tmpl.Layout, func(v interface{}) bool {
        return isEmpty(v)
    })
    return requiredValidator.Validate(i, data)
}