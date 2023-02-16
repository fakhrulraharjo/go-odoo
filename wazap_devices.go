package odoo

import (
	"fmt"
)

// WazapDevices represents wazap.devices model.
type WazapDevices struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	MessageIds  *Relation  `xmlrpc:"message_ids,omptempty"`
	Number      *String    `xmlrpc:"number,omptempty"`
	Status      *Selection `xmlrpc:"status,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// WazapDevicess represents array of wazap.devices model.
type WazapDevicess []WazapDevices

// WazapDevicesModel is the odoo model name.
const WazapDevicesModel = "wazap.devices"

// Many2One convert WazapDevices to *Many2One.
func (wd *WazapDevices) Many2One() *Many2One {
	return NewMany2One(wd.Id.Get(), "")
}

// CreateWazapDevices creates a new wazap.devices model and returns its id.
func (c *Client) CreateWazapDevices(wd *WazapDevices) (int64, error) {
	return c.Create(WazapDevicesModel, wd)
}

// UpdateWazapDevices updates an existing wazap.devices record.
func (c *Client) UpdateWazapDevices(wd *WazapDevices) error {
	return c.UpdateWazapDevicess([]int64{wd.Id.Get()}, wd)
}

// UpdateWazapDevicess updates existing wazap.devices records.
// All records (represented by ids) will be updated by wd values.
func (c *Client) UpdateWazapDevicess(ids []int64, wd *WazapDevices) error {
	return c.Update(WazapDevicesModel, ids, wd)
}

// DeleteWazapDevices deletes an existing wazap.devices record.
func (c *Client) DeleteWazapDevices(id int64) error {
	return c.DeleteWazapDevicess([]int64{id})
}

// DeleteWazapDevicess deletes existing wazap.devices records.
func (c *Client) DeleteWazapDevicess(ids []int64) error {
	return c.Delete(WazapDevicesModel, ids)
}

// GetWazapDevices gets wazap.devices existing record.
func (c *Client) GetWazapDevices(id int64) (*WazapDevices, error) {
	wds, err := c.GetWazapDevicess([]int64{id})
	if err != nil {
		return nil, err
	}
	if wds != nil && len(*wds) > 0 {
		return &((*wds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of wazap.devices not found", id)
}

// GetWazapDevicess gets wazap.devices existing records.
func (c *Client) GetWazapDevicess(ids []int64) (*WazapDevicess, error) {
	wds := &WazapDevicess{}
	if err := c.Read(WazapDevicesModel, ids, nil, wds); err != nil {
		return nil, err
	}
	return wds, nil
}

// FindWazapDevices finds wazap.devices record by querying it with criteria.
func (c *Client) FindWazapDevices(criteria *Criteria) (*WazapDevices, error) {
	wds := &WazapDevicess{}
	if err := c.SearchRead(WazapDevicesModel, criteria, NewOptions().Limit(1), wds); err != nil {
		return nil, err
	}
	if wds != nil && len(*wds) > 0 {
		return &((*wds)[0]), nil
	}
	return nil, fmt.Errorf("wazap.devices was not found with criteria %v", criteria)
}

// FindWazapDevicess finds wazap.devices records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWazapDevicess(criteria *Criteria, options *Options) (*WazapDevicess, error) {
	wds := &WazapDevicess{}
	if err := c.SearchRead(WazapDevicesModel, criteria, options, wds); err != nil {
		return nil, err
	}
	return wds, nil
}

// FindWazapDevicesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWazapDevicesIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WazapDevicesModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWazapDevicesId finds record id by querying it with criteria.
func (c *Client) FindWazapDevicesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WazapDevicesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("wazap.devices was not found with criteria %v and options %v", criteria, options)
}
