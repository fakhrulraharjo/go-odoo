package odoo

import (
	"fmt"
)

// ResUsersDeletion represents res.users.deletion model.
type ResUsersDeletion struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	State       *Selection `xmlrpc:"state,omptempty"`
	UserId      *Many2One  `xmlrpc:"user_id,omptempty"`
	UserIdInt   *Int       `xmlrpc:"user_id_int,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ResUsersDeletions represents array of res.users.deletion model.
type ResUsersDeletions []ResUsersDeletion

// ResUsersDeletionModel is the odoo model name.
const ResUsersDeletionModel = "res.users.deletion"

// Many2One convert ResUsersDeletion to *Many2One.
func (rud *ResUsersDeletion) Many2One() *Many2One {
	return NewMany2One(rud.Id.Get(), "")
}

// CreateResUsersDeletion creates a new res.users.deletion model and returns its id.
func (c *Client) CreateResUsersDeletion(rud *ResUsersDeletion) (int64, error) {
	return c.Create(ResUsersDeletionModel, rud)
}

// UpdateResUsersDeletion updates an existing res.users.deletion record.
func (c *Client) UpdateResUsersDeletion(rud *ResUsersDeletion) error {
	return c.UpdateResUsersDeletions([]int64{rud.Id.Get()}, rud)
}

// UpdateResUsersDeletions updates existing res.users.deletion records.
// All records (represented by ids) will be updated by rud values.
func (c *Client) UpdateResUsersDeletions(ids []int64, rud *ResUsersDeletion) error {
	return c.Update(ResUsersDeletionModel, ids, rud)
}

// DeleteResUsersDeletion deletes an existing res.users.deletion record.
func (c *Client) DeleteResUsersDeletion(id int64) error {
	return c.DeleteResUsersDeletions([]int64{id})
}

// DeleteResUsersDeletions deletes existing res.users.deletion records.
func (c *Client) DeleteResUsersDeletions(ids []int64) error {
	return c.Delete(ResUsersDeletionModel, ids)
}

// GetResUsersDeletion gets res.users.deletion existing record.
func (c *Client) GetResUsersDeletion(id int64) (*ResUsersDeletion, error) {
	ruds, err := c.GetResUsersDeletions([]int64{id})
	if err != nil {
		return nil, err
	}
	if ruds != nil && len(*ruds) > 0 {
		return &((*ruds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.users.deletion not found", id)
}

// GetResUsersDeletions gets res.users.deletion existing records.
func (c *Client) GetResUsersDeletions(ids []int64) (*ResUsersDeletions, error) {
	ruds := &ResUsersDeletions{}
	if err := c.Read(ResUsersDeletionModel, ids, nil, ruds); err != nil {
		return nil, err
	}
	return ruds, nil
}

// FindResUsersDeletion finds res.users.deletion record by querying it with criteria.
func (c *Client) FindResUsersDeletion(criteria *Criteria) (*ResUsersDeletion, error) {
	ruds := &ResUsersDeletions{}
	if err := c.SearchRead(ResUsersDeletionModel, criteria, NewOptions().Limit(1), ruds); err != nil {
		return nil, err
	}
	if ruds != nil && len(*ruds) > 0 {
		return &((*ruds)[0]), nil
	}
	return nil, fmt.Errorf("res.users.deletion was not found with criteria %v", criteria)
}

// FindResUsersDeletions finds res.users.deletion records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersDeletions(criteria *Criteria, options *Options) (*ResUsersDeletions, error) {
	ruds := &ResUsersDeletions{}
	if err := c.SearchRead(ResUsersDeletionModel, criteria, options, ruds); err != nil {
		return nil, err
	}
	return ruds, nil
}

// FindResUsersDeletionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersDeletionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResUsersDeletionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResUsersDeletionId finds record id by querying it with criteria.
func (c *Client) FindResUsersDeletionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersDeletionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.users.deletion was not found with criteria %v and options %v", criteria, options)
}
