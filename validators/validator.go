package validators

import (
	"github.com/go-playground/validator"
)


var validate *validator.Validate


func init(){
	validate = validator.New()
}

func ModelValidator(model any) error {

	err := validate.Struct(model)

	if err != nil {
		
		return err
	}

	return nil 
	
}