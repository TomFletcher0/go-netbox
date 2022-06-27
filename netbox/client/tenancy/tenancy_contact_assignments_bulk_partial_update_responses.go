// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// TenancyContactAssignmentsBulkPartialUpdateReader is a Reader for the TenancyContactAssignmentsBulkPartialUpdate structure.
type TenancyContactAssignmentsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactAssignmentsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactAssignmentsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactAssignmentsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactAssignmentsBulkPartialUpdateOK creates a TenancyContactAssignmentsBulkPartialUpdateOK with default headers values
func NewTenancyContactAssignmentsBulkPartialUpdateOK() *TenancyContactAssignmentsBulkPartialUpdateOK {
	return &TenancyContactAssignmentsBulkPartialUpdateOK{}
}

/* TenancyContactAssignmentsBulkPartialUpdateOK describes a response with status code 200, with default header values.

TenancyContactAssignmentsBulkPartialUpdateOK tenancy contact assignments bulk partial update o k
*/
type TenancyContactAssignmentsBulkPartialUpdateOK struct {
	Payload *models.ContactAssignment
}

func (o *TenancyContactAssignmentsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/contact-assignments/][%d] tenancyContactAssignmentsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *TenancyContactAssignmentsBulkPartialUpdateOK) GetPayload() *models.ContactAssignment {
	return o.Payload
}

func (o *TenancyContactAssignmentsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactAssignmentsBulkPartialUpdateDefault creates a TenancyContactAssignmentsBulkPartialUpdateDefault with default headers values
func NewTenancyContactAssignmentsBulkPartialUpdateDefault(code int) *TenancyContactAssignmentsBulkPartialUpdateDefault {
	return &TenancyContactAssignmentsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* TenancyContactAssignmentsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

TenancyContactAssignmentsBulkPartialUpdateDefault tenancy contact assignments bulk partial update default
*/
type TenancyContactAssignmentsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy contact assignments bulk partial update default response
func (o *TenancyContactAssignmentsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *TenancyContactAssignmentsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /tenancy/contact-assignments/][%d] tenancy_contact-assignments_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *TenancyContactAssignmentsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactAssignmentsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}