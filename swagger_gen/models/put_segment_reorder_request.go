// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PutSegmentReorderRequest put segment reorder request
// swagger:model putSegmentReorderRequest

type PutSegmentReorderRequest struct {

	// segment ids
	// Required: true
	// Min Items: 1
	SegmentIds []int64 `json:"segmentIDs"`
}

/* polymorph putSegmentReorderRequest segmentIDs false */

// Validate validates this put segment reorder request
func (m *PutSegmentReorderRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSegmentIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutSegmentReorderRequest) validateSegmentIds(formats strfmt.Registry) error {

	if err := validate.Required("segmentIDs", "body", m.SegmentIds); err != nil {
		return err
	}

	iSegmentIdsSize := int64(len(m.SegmentIds))

	if err := validate.MinItems("segmentIDs", "body", iSegmentIdsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.SegmentIds); i++ {

		if err := validate.MinimumInt("segmentIDs"+"."+strconv.Itoa(i), "body", int64(m.SegmentIds[i]), 1, false); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PutSegmentReorderRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutSegmentReorderRequest) UnmarshalBinary(b []byte) error {
	var res PutSegmentReorderRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
