package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// CREATED AS AN APP METHOD, TAKES THE REQUEST CONTEXT, RETURNS INT64 AND ERROR
func (app *application) readIDParam(r *http.Request) (int64, error) {
	
	// HTTPROUTER USES PARAMSFROMCONTEXT TO RETRIEVE PARAM NAMES AND VALUES FROM REQUEST CONTEXT
	params := httprouter.ParamsFromContext(r.Context())

	// BYNAME() GETS THE ID VALUE FROM THE SLICE BUT RETURNS A STRING
	// STRCONV CONVERTS RETURNED VALUE FROM STRING TO INT - BASE 10, 64 BIT
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid ID parameter")
	}

	return id, nil
}