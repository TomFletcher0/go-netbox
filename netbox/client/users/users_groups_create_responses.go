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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// UsersGroupsCreateReader is a Reader for the UsersGroupsCreate structure.
type UsersGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUsersGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersGroupsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersGroupsCreateCreated creates a UsersGroupsCreateCreated with default headers values
func NewUsersGroupsCreateCreated() *UsersGroupsCreateCreated {
	return &UsersGroupsCreateCreated{}
}

/*UsersGroupsCreateCreated handles this case with default header values.

UsersGroupsCreateCreated users groups create created
*/
type UsersGroupsCreateCreated struct {
	Payload *models.Group
}

func (o *UsersGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /users/groups/][%d] usersGroupsCreateCreated  %+v", 201, o.Payload)
}

func (o *UsersGroupsCreateCreated) GetPayload() *models.Group {
	return o.Payload
}

func (o *UsersGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersGroupsCreateDefault creates a UsersGroupsCreateDefault with default headers values
func NewUsersGroupsCreateDefault(code int) *UsersGroupsCreateDefault {
	return &UsersGroupsCreateDefault{
		_statusCode: code,
	}
}

/*UsersGroupsCreateDefault handles this case with default header values.

UsersGroupsCreateDefault users groups create default
*/
type UsersGroupsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users groups create default response
func (o *UsersGroupsCreateDefault) Code() int {
	return o._statusCode
}

func (o *UsersGroupsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /users/groups/][%d] users_groups_create default  %+v", o._statusCode, o.Payload)
}

func (o *UsersGroupsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersGroupsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}