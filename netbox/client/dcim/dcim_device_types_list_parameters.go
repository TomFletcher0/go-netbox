// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDcimDeviceTypesListParams creates a new DcimDeviceTypesListParams object
// with the default values initialized.
func NewDcimDeviceTypesListParams() *DcimDeviceTypesListParams {
	var ()
	return &DcimDeviceTypesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceTypesListParamsWithTimeout creates a new DcimDeviceTypesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDeviceTypesListParamsWithTimeout(timeout time.Duration) *DcimDeviceTypesListParams {
	var ()
	return &DcimDeviceTypesListParams{

		timeout: timeout,
	}
}

// NewDcimDeviceTypesListParamsWithContext creates a new DcimDeviceTypesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDeviceTypesListParamsWithContext(ctx context.Context) *DcimDeviceTypesListParams {
	var ()
	return &DcimDeviceTypesListParams{

		Context: ctx,
	}
}

// NewDcimDeviceTypesListParamsWithHTTPClient creates a new DcimDeviceTypesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDeviceTypesListParamsWithHTTPClient(client *http.Client) *DcimDeviceTypesListParams {
	var ()
	return &DcimDeviceTypesListParams{
		HTTPClient: client,
	}
}

/*DcimDeviceTypesListParams contains all the parameters to send to the API endpoint
for the dcim device types list operation typically these are written to a http.Request
*/
type DcimDeviceTypesListParams struct {

	/*IDIn
	  Multiple values may be separated by commas.

	*/
	IDIn *string
	/*IsConsoleServer*/
	IsConsoleServer *string
	/*IsFullDepth*/
	IsFullDepth *string
	/*IsNetworkDevice*/
	IsNetworkDevice *string
	/*IsPdu*/
	IsPdu *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Manufacturer*/
	Manufacturer *string
	/*ManufacturerID*/
	ManufacturerID *string
	/*Model*/
	Model *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*PartNumber*/
	PartNumber *string
	/*Q*/
	Q *string
	/*Slug*/
	Slug *string
	/*SubdeviceRole*/
	SubdeviceRole *string
	/*UHeight*/
	UHeight *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithTimeout(timeout time.Duration) *DcimDeviceTypesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithContext(ctx context.Context) *DcimDeviceTypesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithHTTPClient(client *http.Client) *DcimDeviceTypesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIDIn adds the iDIn to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithIDIn(iDIn *string) *DcimDeviceTypesListParams {
	o.SetIDIn(iDIn)
	return o
}

// SetIDIn adds the idIn to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetIDIn(iDIn *string) {
	o.IDIn = iDIn
}

// WithIsConsoleServer adds the isConsoleServer to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithIsConsoleServer(isConsoleServer *string) *DcimDeviceTypesListParams {
	o.SetIsConsoleServer(isConsoleServer)
	return o
}

// SetIsConsoleServer adds the isConsoleServer to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetIsConsoleServer(isConsoleServer *string) {
	o.IsConsoleServer = isConsoleServer
}

// WithIsFullDepth adds the isFullDepth to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithIsFullDepth(isFullDepth *string) *DcimDeviceTypesListParams {
	o.SetIsFullDepth(isFullDepth)
	return o
}

// SetIsFullDepth adds the isFullDepth to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetIsFullDepth(isFullDepth *string) {
	o.IsFullDepth = isFullDepth
}

// WithIsNetworkDevice adds the isNetworkDevice to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithIsNetworkDevice(isNetworkDevice *string) *DcimDeviceTypesListParams {
	o.SetIsNetworkDevice(isNetworkDevice)
	return o
}

// SetIsNetworkDevice adds the isNetworkDevice to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetIsNetworkDevice(isNetworkDevice *string) {
	o.IsNetworkDevice = isNetworkDevice
}

// WithIsPdu adds the isPdu to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithIsPdu(isPdu *string) *DcimDeviceTypesListParams {
	o.SetIsPdu(isPdu)
	return o
}

// SetIsPdu adds the isPdu to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetIsPdu(isPdu *string) {
	o.IsPdu = isPdu
}

// WithLimit adds the limit to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithLimit(limit *int64) *DcimDeviceTypesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithManufacturer adds the manufacturer to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithManufacturer(manufacturer *string) *DcimDeviceTypesListParams {
	o.SetManufacturer(manufacturer)
	return o
}

// SetManufacturer adds the manufacturer to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetManufacturer(manufacturer *string) {
	o.Manufacturer = manufacturer
}

// WithManufacturerID adds the manufacturerID to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithManufacturerID(manufacturerID *string) *DcimDeviceTypesListParams {
	o.SetManufacturerID(manufacturerID)
	return o
}

// SetManufacturerID adds the manufacturerId to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetManufacturerID(manufacturerID *string) {
	o.ManufacturerID = manufacturerID
}

// WithModel adds the model to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithModel(model *string) *DcimDeviceTypesListParams {
	o.SetModel(model)
	return o
}

// SetModel adds the model to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetModel(model *string) {
	o.Model = model
}

// WithOffset adds the offset to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithOffset(offset *int64) *DcimDeviceTypesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPartNumber adds the partNumber to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithPartNumber(partNumber *string) *DcimDeviceTypesListParams {
	o.SetPartNumber(partNumber)
	return o
}

// SetPartNumber adds the partNumber to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetPartNumber(partNumber *string) {
	o.PartNumber = partNumber
}

// WithQ adds the q to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithQ(q *string) *DcimDeviceTypesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithSlug(slug *string) *DcimDeviceTypesListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSubdeviceRole adds the subdeviceRole to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithSubdeviceRole(subdeviceRole *string) *DcimDeviceTypesListParams {
	o.SetSubdeviceRole(subdeviceRole)
	return o
}

// SetSubdeviceRole adds the subdeviceRole to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetSubdeviceRole(subdeviceRole *string) {
	o.SubdeviceRole = subdeviceRole
}

// WithUHeight adds the uHeight to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithUHeight(uHeight *float64) *DcimDeviceTypesListParams {
	o.SetUHeight(uHeight)
	return o
}

// SetUHeight adds the uHeight to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetUHeight(uHeight *float64) {
	o.UHeight = uHeight
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceTypesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IDIn != nil {

		// query param id__in
		var qrIDIn string
		if o.IDIn != nil {
			qrIDIn = *o.IDIn
		}
		qIDIn := qrIDIn
		if qIDIn != "" {
			if err := r.SetQueryParam("id__in", qIDIn); err != nil {
				return err
			}
		}

	}

	if o.IsConsoleServer != nil {

		// query param is_console_server
		var qrIsConsoleServer string
		if o.IsConsoleServer != nil {
			qrIsConsoleServer = *o.IsConsoleServer
		}
		qIsConsoleServer := qrIsConsoleServer
		if qIsConsoleServer != "" {
			if err := r.SetQueryParam("is_console_server", qIsConsoleServer); err != nil {
				return err
			}
		}

	}

	if o.IsFullDepth != nil {

		// query param is_full_depth
		var qrIsFullDepth string
		if o.IsFullDepth != nil {
			qrIsFullDepth = *o.IsFullDepth
		}
		qIsFullDepth := qrIsFullDepth
		if qIsFullDepth != "" {
			if err := r.SetQueryParam("is_full_depth", qIsFullDepth); err != nil {
				return err
			}
		}

	}

	if o.IsNetworkDevice != nil {

		// query param is_network_device
		var qrIsNetworkDevice string
		if o.IsNetworkDevice != nil {
			qrIsNetworkDevice = *o.IsNetworkDevice
		}
		qIsNetworkDevice := qrIsNetworkDevice
		if qIsNetworkDevice != "" {
			if err := r.SetQueryParam("is_network_device", qIsNetworkDevice); err != nil {
				return err
			}
		}

	}

	if o.IsPdu != nil {

		// query param is_pdu
		var qrIsPdu string
		if o.IsPdu != nil {
			qrIsPdu = *o.IsPdu
		}
		qIsPdu := qrIsPdu
		if qIsPdu != "" {
			if err := r.SetQueryParam("is_pdu", qIsPdu); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Manufacturer != nil {

		// query param manufacturer
		var qrManufacturer string
		if o.Manufacturer != nil {
			qrManufacturer = *o.Manufacturer
		}
		qManufacturer := qrManufacturer
		if qManufacturer != "" {
			if err := r.SetQueryParam("manufacturer", qManufacturer); err != nil {
				return err
			}
		}

	}

	if o.ManufacturerID != nil {

		// query param manufacturer_id
		var qrManufacturerID string
		if o.ManufacturerID != nil {
			qrManufacturerID = *o.ManufacturerID
		}
		qManufacturerID := qrManufacturerID
		if qManufacturerID != "" {
			if err := r.SetQueryParam("manufacturer_id", qManufacturerID); err != nil {
				return err
			}
		}

	}

	if o.Model != nil {

		// query param model
		var qrModel string
		if o.Model != nil {
			qrModel = *o.Model
		}
		qModel := qrModel
		if qModel != "" {
			if err := r.SetQueryParam("model", qModel); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.PartNumber != nil {

		// query param part_number
		var qrPartNumber string
		if o.PartNumber != nil {
			qrPartNumber = *o.PartNumber
		}
		qPartNumber := qrPartNumber
		if qPartNumber != "" {
			if err := r.SetQueryParam("part_number", qPartNumber); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Slug != nil {

		// query param slug
		var qrSlug string
		if o.Slug != nil {
			qrSlug = *o.Slug
		}
		qSlug := qrSlug
		if qSlug != "" {
			if err := r.SetQueryParam("slug", qSlug); err != nil {
				return err
			}
		}

	}

	if o.SubdeviceRole != nil {

		// query param subdevice_role
		var qrSubdeviceRole string
		if o.SubdeviceRole != nil {
			qrSubdeviceRole = *o.SubdeviceRole
		}
		qSubdeviceRole := qrSubdeviceRole
		if qSubdeviceRole != "" {
			if err := r.SetQueryParam("subdevice_role", qSubdeviceRole); err != nil {
				return err
			}
		}

	}

	if o.UHeight != nil {

		// query param u_height
		var qrUHeight float64
		if o.UHeight != nil {
			qrUHeight = *o.UHeight
		}
		qUHeight := swag.FormatFloat64(qrUHeight)
		if qUHeight != "" {
			if err := r.SetQueryParam("u_height", qUHeight); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
