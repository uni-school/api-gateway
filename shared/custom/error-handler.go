package custom

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/uni-school/api-gateaway/libs/response"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CustomHTTPErrorHandler(err error, ctx echo.Context) {
	logrus.Infof("error: %#v\n", err)
	he, ok := err.(*echo.HTTPError)
	if ok {
		ctx.JSON(he.Code, response.ResponseFail(fmt.Sprint(he.Message), he.Internal))
		return
	}
	ve, ok := err.(validator.ValidationErrors)
	if ok {
		ctx.JSON(http.StatusBadRequest, response.ResponseFail(ve.Error(), ve))
		return
	}
	me, ok := err.(*json.MarshalerError)
	if ok {
		ctx.JSON(http.StatusUnprocessableEntity, response.ResponseFail(me.Error(), me.Unwrap()))
		return
	}
	re, ok := err.(runtime.Error)
	if ok {
		ctx.JSON(http.StatusUnprocessableEntity, response.ResponseFail(re.Error(), re))
		return
	}
	se, ok := status.FromError(err)
	if ok {
		var httpStatusCode int = http.StatusInternalServerError
		switch se.Code() {
		case codes.AlreadyExists:
			httpStatusCode = http.StatusBadRequest
		}
		ctx.JSON(httpStatusCode, response.ResponseFail(se.Message(), se.Err()))
		return
	}
	ctx.JSON(http.StatusInternalServerError, response.ResponseFail(err.Error(), err))
}
