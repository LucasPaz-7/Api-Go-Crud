package validation

import (
	"encoding/json"
	"errors"

	"github.com/LucasPaz-7/Api-Go-Crud/src/configuration/rest_err"
	
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate = validator.New()
	transl  ut.Translator
)

func init(){
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok{

		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(
	validation_err error,
) *rest_err.RestErr {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr){
		return rest_err.NewBadRequestError("invalid field type")
	} else if errors.As(validation_err, &jsonValidationError){
		errorCauses := []rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors){
			causes := rest_err.Causes{
				Message: e.Translate(transl),
				Field: e.Field(),
			}

			errorCauses = append(errorCauses, causes)
	}

	return rest_err.NewBadRequestValidationError("Some fields are invalid", errorCauses)

	} else{
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}

}
