/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.5.8 (3.5)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"time"
)

// checks if the Device type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Device{}

// Device Adds support for custom fields and tags.
type Device struct {
	Id         int32                  `json:"id"`
	Url        string                 `json:"url"`
	Display    string                 `json:"display"`
	Name       NullableString         `json:"name,omitempty"`
	DeviceType NestedDeviceType       `json:"device_type"`
	DeviceRole NestedDeviceRole       `json:"device_role"`
	Tenant     NullableNestedTenant   `json:"tenant,omitempty"`
	Platform   NullableNestedPlatform `json:"platform,omitempty"`
	// Chassis serial number, assigned by the manufacturer
	Serial *string `json:"serial,omitempty"`
	// A unique tag used to identify this device
	AssetTag       NullableString               `json:"asset_tag,omitempty"`
	Site           NestedSite                   `json:"site"`
	Location       NullableNestedLocation       `json:"location,omitempty"`
	Rack           NullableNestedRack           `json:"rack,omitempty"`
	Position       NullableFloat64              `json:"position,omitempty"`
	Face           *DeviceFace                  `json:"face,omitempty"`
	ParentDevice   NestedDevice                 `json:"parent_device"`
	Status         *DeviceStatus                `json:"status,omitempty"`
	Airflow        *DeviceAirflow               `json:"airflow,omitempty"`
	PrimaryIp      NestedIPAddress              `json:"primary_ip"`
	PrimaryIp4     NullableNestedIPAddress      `json:"primary_ip4,omitempty"`
	PrimaryIp6     NullableNestedIPAddress      `json:"primary_ip6,omitempty"`
	Cluster        NullableNestedCluster        `json:"cluster,omitempty"`
	VirtualChassis NullableNestedVirtualChassis `json:"virtual_chassis,omitempty"`
	VcPosition     NullableInt32                `json:"vc_position,omitempty"`
	// Virtual chassis master election priority
	VcPriority     NullableInt32                `json:"vc_priority,omitempty"`
	Description    *string                      `json:"description,omitempty"`
	Comments       *string                      `json:"comments,omitempty"`
	ConfigTemplate NullableNestedConfigTemplate `json:"config_template,omitempty"`
	// Local config context data takes precedence over source contexts in the final rendered config context
	LocalContextData     map[string]interface{} `json:"local_context_data,omitempty"`
	Tags                 []NestedTag            `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	Created              NullableTime           `json:"created"`
	LastUpdated          NullableTime           `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _Device Device

// NewDevice instantiates a new Device object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevice(id int32, url string, display string, deviceType NestedDeviceType, deviceRole NestedDeviceRole, site NestedSite, parentDevice NestedDevice, primaryIp NestedIPAddress, created NullableTime, lastUpdated NullableTime) *Device {
	this := Device{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.DeviceType = deviceType
	this.DeviceRole = deviceRole
	this.Site = site
	this.ParentDevice = parentDevice
	this.PrimaryIp = primaryIp
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewDeviceWithDefaults instantiates a new Device object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceWithDefaults() *Device {
	this := Device{}
	return &this
}

// GetId returns the Id field value
func (o *Device) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Device) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Device) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Device) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Device) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Device) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *Device) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Device) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Device) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Device) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Device) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *Device) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Device) UnsetName() {
	o.Name.Unset()
}

// GetDeviceType returns the DeviceType field value
func (o *Device) GetDeviceType() NestedDeviceType {
	if o == nil {
		var ret NestedDeviceType
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *Device) GetDeviceTypeOk() (*NestedDeviceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *Device) SetDeviceType(v NestedDeviceType) {
	o.DeviceType = v
}

// GetDeviceRole returns the DeviceRole field value
func (o *Device) GetDeviceRole() NestedDeviceRole {
	if o == nil {
		var ret NestedDeviceRole
		return ret
	}

	return o.DeviceRole
}

// GetDeviceRoleOk returns a tuple with the DeviceRole field value
// and a boolean to check if the value has been set.
func (o *Device) GetDeviceRoleOk() (*NestedDeviceRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceRole, true
}

// SetDeviceRole sets field value
func (o *Device) SetDeviceRole(v NestedDeviceRole) {
	o.DeviceRole = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetTenant() NestedTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret NestedTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetTenantOk() (*NestedTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *Device) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableNestedTenant and assigns it to the Tenant field.
func (o *Device) SetTenant(v NestedTenant) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *Device) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *Device) UnsetTenant() {
	o.Tenant.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetPlatform() NestedPlatform {
	if o == nil || IsNil(o.Platform.Get()) {
		var ret NestedPlatform
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetPlatformOk() (*NestedPlatform, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *Device) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableNestedPlatform and assigns it to the Platform field.
func (o *Device) SetPlatform(v NestedPlatform) {
	o.Platform.Set(&v)
}

// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *Device) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *Device) UnsetPlatform() {
	o.Platform.Unset()
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *Device) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *Device) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *Device) SetSerial(v string) {
	o.Serial = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetAssetTag() string {
	if o == nil || IsNil(o.AssetTag.Get()) {
		var ret string
		return ret
	}
	return *o.AssetTag.Get()
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetAssetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetTag.Get(), o.AssetTag.IsSet()
}

// HasAssetTag returns a boolean if a field has been set.
func (o *Device) HasAssetTag() bool {
	if o != nil && o.AssetTag.IsSet() {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given NullableString and assigns it to the AssetTag field.
func (o *Device) SetAssetTag(v string) {
	o.AssetTag.Set(&v)
}

// SetAssetTagNil sets the value for AssetTag to be an explicit nil
func (o *Device) SetAssetTagNil() {
	o.AssetTag.Set(nil)
}

// UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
func (o *Device) UnsetAssetTag() {
	o.AssetTag.Unset()
}

// GetSite returns the Site field value
func (o *Device) GetSite() NestedSite {
	if o == nil {
		var ret NestedSite
		return ret
	}

	return o.Site
}

// GetSiteOk returns a tuple with the Site field value
// and a boolean to check if the value has been set.
func (o *Device) GetSiteOk() (*NestedSite, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Site, true
}

// SetSite sets field value
func (o *Device) SetSite(v NestedSite) {
	o.Site = v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetLocation() NestedLocation {
	if o == nil || IsNil(o.Location.Get()) {
		var ret NestedLocation
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetLocationOk() (*NestedLocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *Device) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableNestedLocation and assigns it to the Location field.
func (o *Device) SetLocation(v NestedLocation) {
	o.Location.Set(&v)
}

// SetLocationNil sets the value for Location to be an explicit nil
func (o *Device) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *Device) UnsetLocation() {
	o.Location.Unset()
}

// GetRack returns the Rack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetRack() NestedRack {
	if o == nil || IsNil(o.Rack.Get()) {
		var ret NestedRack
		return ret
	}
	return *o.Rack.Get()
}

// GetRackOk returns a tuple with the Rack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetRackOk() (*NestedRack, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rack.Get(), o.Rack.IsSet()
}

// HasRack returns a boolean if a field has been set.
func (o *Device) HasRack() bool {
	if o != nil && o.Rack.IsSet() {
		return true
	}

	return false
}

// SetRack gets a reference to the given NullableNestedRack and assigns it to the Rack field.
func (o *Device) SetRack(v NestedRack) {
	o.Rack.Set(&v)
}

// SetRackNil sets the value for Rack to be an explicit nil
func (o *Device) SetRackNil() {
	o.Rack.Set(nil)
}

// UnsetRack ensures that no value is present for Rack, not even an explicit nil
func (o *Device) UnsetRack() {
	o.Rack.Unset()
}

// GetPosition returns the Position field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetPosition() float64 {
	if o == nil || IsNil(o.Position.Get()) {
		var ret float64
		return ret
	}
	return *o.Position.Get()
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetPositionOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Position.Get(), o.Position.IsSet()
}

// HasPosition returns a boolean if a field has been set.
func (o *Device) HasPosition() bool {
	if o != nil && o.Position.IsSet() {
		return true
	}

	return false
}

// SetPosition gets a reference to the given NullableFloat64 and assigns it to the Position field.
func (o *Device) SetPosition(v float64) {
	o.Position.Set(&v)
}

// SetPositionNil sets the value for Position to be an explicit nil
func (o *Device) SetPositionNil() {
	o.Position.Set(nil)
}

// UnsetPosition ensures that no value is present for Position, not even an explicit nil
func (o *Device) UnsetPosition() {
	o.Position.Unset()
}

// GetFace returns the Face field value if set, zero value otherwise.
func (o *Device) GetFace() DeviceFace {
	if o == nil || IsNil(o.Face) {
		var ret DeviceFace
		return ret
	}
	return *o.Face
}

// GetFaceOk returns a tuple with the Face field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetFaceOk() (*DeviceFace, bool) {
	if o == nil || IsNil(o.Face) {
		return nil, false
	}
	return o.Face, true
}

// HasFace returns a boolean if a field has been set.
func (o *Device) HasFace() bool {
	if o != nil && !IsNil(o.Face) {
		return true
	}

	return false
}

// SetFace gets a reference to the given DeviceFace and assigns it to the Face field.
func (o *Device) SetFace(v DeviceFace) {
	o.Face = &v
}

// GetParentDevice returns the ParentDevice field value
func (o *Device) GetParentDevice() NestedDevice {
	if o == nil {
		var ret NestedDevice
		return ret
	}

	return o.ParentDevice
}

// GetParentDeviceOk returns a tuple with the ParentDevice field value
// and a boolean to check if the value has been set.
func (o *Device) GetParentDeviceOk() (*NestedDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentDevice, true
}

// SetParentDevice sets field value
func (o *Device) SetParentDevice(v NestedDevice) {
	o.ParentDevice = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Device) GetStatus() DeviceStatus {
	if o == nil || IsNil(o.Status) {
		var ret DeviceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetStatusOk() (*DeviceStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Device) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DeviceStatus and assigns it to the Status field.
func (o *Device) SetStatus(v DeviceStatus) {
	o.Status = &v
}

// GetAirflow returns the Airflow field value if set, zero value otherwise.
func (o *Device) GetAirflow() DeviceAirflow {
	if o == nil || IsNil(o.Airflow) {
		var ret DeviceAirflow
		return ret
	}
	return *o.Airflow
}

// GetAirflowOk returns a tuple with the Airflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetAirflowOk() (*DeviceAirflow, bool) {
	if o == nil || IsNil(o.Airflow) {
		return nil, false
	}
	return o.Airflow, true
}

// HasAirflow returns a boolean if a field has been set.
func (o *Device) HasAirflow() bool {
	if o != nil && !IsNil(o.Airflow) {
		return true
	}

	return false
}

// SetAirflow gets a reference to the given DeviceAirflow and assigns it to the Airflow field.
func (o *Device) SetAirflow(v DeviceAirflow) {
	o.Airflow = &v
}

// GetPrimaryIp returns the PrimaryIp field value
func (o *Device) GetPrimaryIp() NestedIPAddress {
	if o == nil {
		var ret NestedIPAddress
		return ret
	}

	return o.PrimaryIp
}

// GetPrimaryIpOk returns a tuple with the PrimaryIp field value
// and a boolean to check if the value has been set.
func (o *Device) GetPrimaryIpOk() (*NestedIPAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryIp, true
}

// SetPrimaryIp sets field value
func (o *Device) SetPrimaryIp(v NestedIPAddress) {
	o.PrimaryIp = v
}

// GetPrimaryIp4 returns the PrimaryIp4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetPrimaryIp4() NestedIPAddress {
	if o == nil || IsNil(o.PrimaryIp4.Get()) {
		var ret NestedIPAddress
		return ret
	}
	return *o.PrimaryIp4.Get()
}

// GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetPrimaryIp4Ok() (*NestedIPAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp4.Get(), o.PrimaryIp4.IsSet()
}

// HasPrimaryIp4 returns a boolean if a field has been set.
func (o *Device) HasPrimaryIp4() bool {
	if o != nil && o.PrimaryIp4.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp4 gets a reference to the given NullableNestedIPAddress and assigns it to the PrimaryIp4 field.
func (o *Device) SetPrimaryIp4(v NestedIPAddress) {
	o.PrimaryIp4.Set(&v)
}

// SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil
func (o *Device) SetPrimaryIp4Nil() {
	o.PrimaryIp4.Set(nil)
}

// UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
func (o *Device) UnsetPrimaryIp4() {
	o.PrimaryIp4.Unset()
}

// GetPrimaryIp6 returns the PrimaryIp6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetPrimaryIp6() NestedIPAddress {
	if o == nil || IsNil(o.PrimaryIp6.Get()) {
		var ret NestedIPAddress
		return ret
	}
	return *o.PrimaryIp6.Get()
}

// GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetPrimaryIp6Ok() (*NestedIPAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp6.Get(), o.PrimaryIp6.IsSet()
}

// HasPrimaryIp6 returns a boolean if a field has been set.
func (o *Device) HasPrimaryIp6() bool {
	if o != nil && o.PrimaryIp6.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp6 gets a reference to the given NullableNestedIPAddress and assigns it to the PrimaryIp6 field.
func (o *Device) SetPrimaryIp6(v NestedIPAddress) {
	o.PrimaryIp6.Set(&v)
}

// SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil
func (o *Device) SetPrimaryIp6Nil() {
	o.PrimaryIp6.Set(nil)
}

// UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
func (o *Device) UnsetPrimaryIp6() {
	o.PrimaryIp6.Unset()
}

// GetCluster returns the Cluster field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetCluster() NestedCluster {
	if o == nil || IsNil(o.Cluster.Get()) {
		var ret NestedCluster
		return ret
	}
	return *o.Cluster.Get()
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetClusterOk() (*NestedCluster, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cluster.Get(), o.Cluster.IsSet()
}

// HasCluster returns a boolean if a field has been set.
func (o *Device) HasCluster() bool {
	if o != nil && o.Cluster.IsSet() {
		return true
	}

	return false
}

// SetCluster gets a reference to the given NullableNestedCluster and assigns it to the Cluster field.
func (o *Device) SetCluster(v NestedCluster) {
	o.Cluster.Set(&v)
}

// SetClusterNil sets the value for Cluster to be an explicit nil
func (o *Device) SetClusterNil() {
	o.Cluster.Set(nil)
}

// UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
func (o *Device) UnsetCluster() {
	o.Cluster.Unset()
}

// GetVirtualChassis returns the VirtualChassis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetVirtualChassis() NestedVirtualChassis {
	if o == nil || IsNil(o.VirtualChassis.Get()) {
		var ret NestedVirtualChassis
		return ret
	}
	return *o.VirtualChassis.Get()
}

// GetVirtualChassisOk returns a tuple with the VirtualChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetVirtualChassisOk() (*NestedVirtualChassis, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualChassis.Get(), o.VirtualChassis.IsSet()
}

// HasVirtualChassis returns a boolean if a field has been set.
func (o *Device) HasVirtualChassis() bool {
	if o != nil && o.VirtualChassis.IsSet() {
		return true
	}

	return false
}

// SetVirtualChassis gets a reference to the given NullableNestedVirtualChassis and assigns it to the VirtualChassis field.
func (o *Device) SetVirtualChassis(v NestedVirtualChassis) {
	o.VirtualChassis.Set(&v)
}

// SetVirtualChassisNil sets the value for VirtualChassis to be an explicit nil
func (o *Device) SetVirtualChassisNil() {
	o.VirtualChassis.Set(nil)
}

// UnsetVirtualChassis ensures that no value is present for VirtualChassis, not even an explicit nil
func (o *Device) UnsetVirtualChassis() {
	o.VirtualChassis.Unset()
}

// GetVcPosition returns the VcPosition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetVcPosition() int32 {
	if o == nil || IsNil(o.VcPosition.Get()) {
		var ret int32
		return ret
	}
	return *o.VcPosition.Get()
}

// GetVcPositionOk returns a tuple with the VcPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetVcPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VcPosition.Get(), o.VcPosition.IsSet()
}

// HasVcPosition returns a boolean if a field has been set.
func (o *Device) HasVcPosition() bool {
	if o != nil && o.VcPosition.IsSet() {
		return true
	}

	return false
}

// SetVcPosition gets a reference to the given NullableInt32 and assigns it to the VcPosition field.
func (o *Device) SetVcPosition(v int32) {
	o.VcPosition.Set(&v)
}

// SetVcPositionNil sets the value for VcPosition to be an explicit nil
func (o *Device) SetVcPositionNil() {
	o.VcPosition.Set(nil)
}

// UnsetVcPosition ensures that no value is present for VcPosition, not even an explicit nil
func (o *Device) UnsetVcPosition() {
	o.VcPosition.Unset()
}

// GetVcPriority returns the VcPriority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetVcPriority() int32 {
	if o == nil || IsNil(o.VcPriority.Get()) {
		var ret int32
		return ret
	}
	return *o.VcPriority.Get()
}

// GetVcPriorityOk returns a tuple with the VcPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetVcPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VcPriority.Get(), o.VcPriority.IsSet()
}

// HasVcPriority returns a boolean if a field has been set.
func (o *Device) HasVcPriority() bool {
	if o != nil && o.VcPriority.IsSet() {
		return true
	}

	return false
}

// SetVcPriority gets a reference to the given NullableInt32 and assigns it to the VcPriority field.
func (o *Device) SetVcPriority(v int32) {
	o.VcPriority.Set(&v)
}

// SetVcPriorityNil sets the value for VcPriority to be an explicit nil
func (o *Device) SetVcPriorityNil() {
	o.VcPriority.Set(nil)
}

// UnsetVcPriority ensures that no value is present for VcPriority, not even an explicit nil
func (o *Device) UnsetVcPriority() {
	o.VcPriority.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Device) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Device) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Device) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Device) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Device) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *Device) SetComments(v string) {
	o.Comments = &v
}

// GetConfigTemplate returns the ConfigTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetConfigTemplate() NestedConfigTemplate {
	if o == nil || IsNil(o.ConfigTemplate.Get()) {
		var ret NestedConfigTemplate
		return ret
	}
	return *o.ConfigTemplate.Get()
}

// GetConfigTemplateOk returns a tuple with the ConfigTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetConfigTemplateOk() (*NestedConfigTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigTemplate.Get(), o.ConfigTemplate.IsSet()
}

// HasConfigTemplate returns a boolean if a field has been set.
func (o *Device) HasConfigTemplate() bool {
	if o != nil && o.ConfigTemplate.IsSet() {
		return true
	}

	return false
}

// SetConfigTemplate gets a reference to the given NullableNestedConfigTemplate and assigns it to the ConfigTemplate field.
func (o *Device) SetConfigTemplate(v NestedConfigTemplate) {
	o.ConfigTemplate.Set(&v)
}

// SetConfigTemplateNil sets the value for ConfigTemplate to be an explicit nil
func (o *Device) SetConfigTemplateNil() {
	o.ConfigTemplate.Set(nil)
}

// UnsetConfigTemplate ensures that no value is present for ConfigTemplate, not even an explicit nil
func (o *Device) UnsetConfigTemplate() {
	o.ConfigTemplate.Unset()
}

// GetLocalContextData returns the LocalContextData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Device) GetLocalContextData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.LocalContextData
}

// GetLocalContextDataOk returns a tuple with the LocalContextData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetLocalContextDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.LocalContextData) {
		return map[string]interface{}{}, false
	}
	return o.LocalContextData, true
}

// HasLocalContextData returns a boolean if a field has been set.
func (o *Device) HasLocalContextData() bool {
	if o != nil && IsNil(o.LocalContextData) {
		return true
	}

	return false
}

// SetLocalContextData gets a reference to the given map[string]interface{} and assigns it to the LocalContextData field.
func (o *Device) SetLocalContextData(v map[string]interface{}) {
	o.LocalContextData = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Device) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Device) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *Device) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Device) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Device) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Device) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Device) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *Device) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Device) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Device) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *Device) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o Device) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Device) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: url is readOnly
	// skip: display is readOnly
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	toSerialize["device_type"] = o.DeviceType
	toSerialize["device_role"] = o.DeviceRole
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Platform.IsSet() {
		toSerialize["platform"] = o.Platform.Get()
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if o.AssetTag.IsSet() {
		toSerialize["asset_tag"] = o.AssetTag.Get()
	}
	toSerialize["site"] = o.Site
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Rack.IsSet() {
		toSerialize["rack"] = o.Rack.Get()
	}
	if o.Position.IsSet() {
		toSerialize["position"] = o.Position.Get()
	}
	if !IsNil(o.Face) {
		toSerialize["face"] = o.Face
	}
	// skip: parent_device is readOnly
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Airflow) {
		toSerialize["airflow"] = o.Airflow
	}
	// skip: primary_ip is readOnly
	if o.PrimaryIp4.IsSet() {
		toSerialize["primary_ip4"] = o.PrimaryIp4.Get()
	}
	if o.PrimaryIp6.IsSet() {
		toSerialize["primary_ip6"] = o.PrimaryIp6.Get()
	}
	if o.Cluster.IsSet() {
		toSerialize["cluster"] = o.Cluster.Get()
	}
	if o.VirtualChassis.IsSet() {
		toSerialize["virtual_chassis"] = o.VirtualChassis.Get()
	}
	if o.VcPosition.IsSet() {
		toSerialize["vc_position"] = o.VcPosition.Get()
	}
	if o.VcPriority.IsSet() {
		toSerialize["vc_priority"] = o.VcPriority.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if o.ConfigTemplate.IsSet() {
		toSerialize["config_template"] = o.ConfigTemplate.Get()
	}
	if o.LocalContextData != nil {
		toSerialize["local_context_data"] = o.LocalContextData
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Device) UnmarshalJSON(bytes []byte) (err error) {
	varDevice := _Device{}

	if err = json.Unmarshal(bytes, &varDevice); err == nil {
		*o = Device(varDevice)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "device_role")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "asset_tag")
		delete(additionalProperties, "site")
		delete(additionalProperties, "location")
		delete(additionalProperties, "rack")
		delete(additionalProperties, "position")
		delete(additionalProperties, "face")
		delete(additionalProperties, "parent_device")
		delete(additionalProperties, "status")
		delete(additionalProperties, "airflow")
		delete(additionalProperties, "primary_ip")
		delete(additionalProperties, "primary_ip4")
		delete(additionalProperties, "primary_ip6")
		delete(additionalProperties, "cluster")
		delete(additionalProperties, "virtual_chassis")
		delete(additionalProperties, "vc_position")
		delete(additionalProperties, "vc_priority")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "config_template")
		delete(additionalProperties, "local_context_data")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDevice struct {
	value *Device
	isSet bool
}

func (v NullableDevice) Get() *Device {
	return v.value
}

func (v *NullableDevice) Set(val *Device) {
	v.value = val
	v.isSet = true
}

func (v NullableDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevice(val *Device) *NullableDevice {
	return &NullableDevice{value: val, isSet: true}
}

func (v NullableDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}