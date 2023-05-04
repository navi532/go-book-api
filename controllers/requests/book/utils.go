package requests

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Parse(ctx *gin.Context, request interface{}) error {

	if ctx.ShouldBindJSON(request) != nil {
		return errors.New("unable to parse create request")
	}

	switch v := request.(type) {
	case *[]CreateRequest:
		for _, eachRequest := range *v {
			if err := validate.Struct(eachRequest); err != nil {
				return err
			}
		}

	default:
		if err := validate.Struct(request); err != nil {
			return err
		}
	}

	return nil
}
