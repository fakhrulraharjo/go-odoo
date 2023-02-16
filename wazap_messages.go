package odoo

import (
	"fmt"
)

// WazapMessages represents wazap.messages model.
type WazapMessages struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	ContactId   *Many2One  `xmlrpc:"contact_id,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DeviceId    *Many2One  `xmlrpc:"device_id,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	Message     *String    `xmlrpc:"message,omptempty"`
	Status      *Selection `xmlrpc:"status,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// WazapMessagess represents array of wazap.messages model.
type WazapMessagess []WazapMessages

// WazapMessagesModel is the odoo model name.
const WazapMessagesModel = "wazap.messages"

// Many2One convert WazapMessages to *Many2One.
func (wm *WazapMessages) Many2One() *Many2One {
	return NewMany2One(wm.Id.Get(), "")
}

// CreateWazapMessages creates a new wazap.messages model and returns its id.
func (c *Client) CreateWazapMessages(wm *WazapMessages) (int64, error) {
	return c.Create(WazapMessagesModel, wm)
}

// UpdateWazapMessages updates an existing wazap.messages record.
func (c *Client) UpdateWazapMessages(wm *WazapMessages) error {
	return c.UpdateWazapMessagess([]int64{wm.Id.Get()}, wm)
}

// UpdateWazapMessagess updates existing wazap.messages records.
// All records (represented by ids) will be updated by wm values.
func (c *Client) UpdateWazapMessagess(ids []int64, wm *WazapMessages) error {
	return c.Update(WazapMessagesModel, ids, wm)
}

// DeleteWazapMessages deletes an existing wazap.messages record.
func (c *Client) DeleteWazapMessages(id int64) error {
	return c.DeleteWazapMessagess([]int64{id})
}

// DeleteWazapMessagess deletes existing wazap.messages records.
func (c *Client) DeleteWazapMessagess(ids []int64) error {
	return c.Delete(WazapMessagesModel, ids)
}

// GetWazapMessages gets wazap.messages existing record.
func (c *Client) GetWazapMessages(id int64) (*WazapMessages, error) {
	wms, err := c.GetWazapMessagess([]int64{id})
	if err != nil {
		return nil, err
	}
	if wms != nil && len(*wms) > 0 {
		return &((*wms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of wazap.messages not found", id)
}

// GetWazapMessagess gets wazap.messages existing records.
func (c *Client) GetWazapMessagess(ids []int64) (*WazapMessagess, error) {
	wms := &WazapMessagess{}
	if err := c.Read(WazapMessagesModel, ids, nil, wms); err != nil {
		return nil, err
	}
	return wms, nil
}

// FindWazapMessages finds wazap.messages record by querying it with criteria.
func (c *Client) FindWazapMessages(criteria *Criteria) (*WazapMessages, error) {
	wms := &WazapMessagess{}
	if err := c.SearchRead(WazapMessagesModel, criteria, NewOptions().Limit(1), wms); err != nil {
		return nil, err
	}
	if wms != nil && len(*wms) > 0 {
		return &((*wms)[0]), nil
	}
	return nil, fmt.Errorf("wazap.messages was not found with criteria %v", criteria)
}

// FindWazapMessagess finds wazap.messages records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWazapMessagess(criteria *Criteria, options *Options) (*WazapMessagess, error) {
	wms := &WazapMessagess{}
	if err := c.SearchRead(WazapMessagesModel, criteria, options, wms); err != nil {
		return nil, err
	}
	return wms, nil
}

// FindWazapMessagesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWazapMessagesIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WazapMessagesModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWazapMessagesId finds record id by querying it with criteria.
func (c *Client) FindWazapMessagesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WazapMessagesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("wazap.messages was not found with criteria %v and options %v", criteria, options)
}
