package entryRequest

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris"
)

type ErrParams struct {
	Code int
	Msg  string
}

var NorMal = ErrParams{0, ""}

func RequestParams(ctx iris.Context, validate *validator.Validate, params interface{}) (interface{}, *ErrParams) {
	if err := ctx.ReadJSON(params); err != nil {
		return nil, &ErrParams{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	err := validate.Struct(params)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return nil, &ErrParams{
				Code: -1,
				Msg:  err.Error(),
			}
		}
		for _, err := range errs {
			return nil, &ErrParams{
				Code: -1,
				Msg:  err.ActualTag(),
			}
		}
	}
	return params, &NorMal
}
