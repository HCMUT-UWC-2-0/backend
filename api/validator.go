package api

import (
	"github.com/HCMUT-UWC-2-0/backend/util"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var validDate validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if date, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsDateString(date)
	}
	return false
}

func (server *Server) setupGinDateValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("date", validDate)
	}
}


// TODO: Add validator for gender, major, faculty, and year