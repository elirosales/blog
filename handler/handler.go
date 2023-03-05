package handler

import (
	"errors"
	"fmt"

	"github.com/elizabethrosales/blog/service"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	log *logrus.Logger
	svc *service.Service
}

func New(svc *service.Service, log *logrus.Logger) *Handler {
	return &Handler{
		log: log,
		svc: svc,
	}
}

func (h *Handler) translateBindingErr(err error) string {
	if err.Error() == "EOF" {
		return "Missing request body"
	}

	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}

	var errMsg string
	if errors.As(err, &errs) {
		for _, fe := range errs {
			if errMsg != "" {
				errMsg += ", "
			}

			errMsg += getErrorMsg(fe)
		}
	}

	return errMsg
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("`%s` is a required field", fe.Field())
	}

	return fe.Tag()
}
