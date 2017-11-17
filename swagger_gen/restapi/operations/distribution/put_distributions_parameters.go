// Code generated by go-swagger; DO NOT EDIT.

package distribution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/checkr/flagr/swagger_gen/models"
)

// NewPutDistributionsParams creates a new PutDistributionsParams object
// with the default values initialized.
func NewPutDistributionsParams() PutDistributionsParams {
	var ()
	return PutDistributionsParams{}
}

// PutDistributionsParams contains all the bound params for the put distributions operation
// typically these are obtained from a http.Request
//
// swagger:parameters putDistributions
type PutDistributionsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*array of distributions
	  Required: true
	  In: body
	*/
	Body *models.PutDistributionsRequest
	/*numeric ID of the flag
	  Required: true
	  Minimum: 1
	  In: path
	*/
	FlagID int64
	/*numeric ID of the segment
	  Required: true
	  Minimum: 1
	  In: path
	*/
	SegmentID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PutDistributionsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PutDistributionsRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}

	} else {
		res = append(res, errors.Required("body", "body"))
	}

	rFlagID, rhkFlagID, _ := route.Params.GetOK("flagID")
	if err := o.bindFlagID(rFlagID, rhkFlagID, route.Formats); err != nil {
		res = append(res, err)
	}

	rSegmentID, rhkSegmentID, _ := route.Params.GetOK("segmentID")
	if err := o.bindSegmentID(rSegmentID, rhkSegmentID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutDistributionsParams) bindFlagID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("flagID", "path", "int64", raw)
	}
	o.FlagID = value

	if err := o.validateFlagID(formats); err != nil {
		return err
	}

	return nil
}

func (o *PutDistributionsParams) validateFlagID(formats strfmt.Registry) error {

	if err := validate.MinimumInt("flagID", "path", int64(o.FlagID), 1, false); err != nil {
		return err
	}

	return nil
}

func (o *PutDistributionsParams) bindSegmentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("segmentID", "path", "int64", raw)
	}
	o.SegmentID = value

	if err := o.validateSegmentID(formats); err != nil {
		return err
	}

	return nil
}

func (o *PutDistributionsParams) validateSegmentID(formats strfmt.Registry) error {

	if err := validate.MinimumInt("segmentID", "path", int64(o.SegmentID), 1, false); err != nil {
		return err
	}

	return nil
}
