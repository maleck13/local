package test

import (
	"bytes"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/maleck13/local/app"
	"golang.org/x/net/context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// ListForCountyAndAreaCouncillorsOK runs the method ListForCountyAndArea of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListForCountyAndAreaCouncillorsOK(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CouncillorsController, area string, county string) (http.ResponseWriter, app.GoaLocalCouncillorCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{area}
		query["area"] = sliceVal
	}
	{
		sliceVal := []string{county}
		query["county"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/councillors"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{area}
		prms["area"] = sliceVal
	}
	{
		sliceVal := []string{county}
		prms["county"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CouncillorsTest"), rw, req, prms)
	listForCountyAndAreaCtx, err := app.NewListForCountyAndAreaCouncillorsContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.ListForCountyAndArea(listForCountyAndAreaCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.GoaLocalCouncillorCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(app.GoaLocalCouncillorCollection)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.GoaLocalCouncillorCollection", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// ListForCountyAndAreaCouncillorsUnauthorized runs the method ListForCountyAndArea of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListForCountyAndAreaCouncillorsUnauthorized(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CouncillorsController, area string, county string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{area}
		query["area"] = sliceVal
	}
	{
		sliceVal := []string{county}
		query["county"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/councillors"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{area}
		prms["area"] = sliceVal
	}
	{
		sliceVal := []string{county}
		prms["county"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CouncillorsTest"), rw, req, prms)
	listForCountyAndAreaCtx, err := app.NewListForCountyAndAreaCouncillorsContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.ListForCountyAndArea(listForCountyAndAreaCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 401 {
		t.Errorf("invalid response status code: got %+v, expected 401", rw.Code)
	}

	// Return results
	return rw
}

// ReadByIDCouncillorsOK runs the method ReadByID of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ReadByIDCouncillorsOK(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CouncillorsController, id string) (http.ResponseWriter, *app.GoaLocalCouncillor) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/councillors/%v", id),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CouncillorsTest"), rw, req, prms)
	readByIDCtx, err := app.NewReadByIDCouncillorsContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.ReadByID(readByIDCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.GoaLocalCouncillor
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.GoaLocalCouncillor)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.GoaLocalCouncillor", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// ReadByIDCouncillorsUnauthorized runs the method ReadByID of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ReadByIDCouncillorsUnauthorized(t *testing.T, ctx context.Context, service *goa.Service, ctrl app.CouncillorsController, id string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/councillors/%v", id),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["id"] = []string{fmt.Sprintf("%v", id)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CouncillorsTest"), rw, req, prms)
	readByIDCtx, err := app.NewReadByIDCouncillorsContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.ReadByID(readByIDCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 401 {
		t.Errorf("invalid response status code: got %+v, expected 401", rw.Code)
	}

	// Return results
	return rw
}
