package odoo

import (
	"fmt"
)

// IrProfile represents ir.profile model.
type IrProfile struct {
	LastUpdate     *Time   `xmlrpc:"__last_update,omptempty"`
	CreateDate     *Time   `xmlrpc:"create_date,omptempty"`
	DisplayName    *String `xmlrpc:"display_name,omptempty"`
	Duration       *Float  `xmlrpc:"duration,omptempty"`
	EntryCount     *Int    `xmlrpc:"entry_count,omptempty"`
	Id             *Int    `xmlrpc:"id,omptempty"`
	InitStackTrace *String `xmlrpc:"init_stack_trace,omptempty"`
	Name           *String `xmlrpc:"name,omptempty"`
	Qweb           *String `xmlrpc:"qweb,omptempty"`
	Session        *String `xmlrpc:"session,omptempty"`
	Speedscope     *String `xmlrpc:"speedscope,omptempty"`
	SpeedscopeUrl  *String `xmlrpc:"speedscope_url,omptempty"`
	Sql            *String `xmlrpc:"sql,omptempty"`
	SqlCount       *Int    `xmlrpc:"sql_count,omptempty"`
	TracesAsync    *String `xmlrpc:"traces_async,omptempty"`
	TracesSync     *String `xmlrpc:"traces_sync,omptempty"`
}

// IrProfiles represents array of ir.profile model.
type IrProfiles []IrProfile

// IrProfileModel is the odoo model name.
const IrProfileModel = "ir.profile"

// Many2One convert IrProfile to *Many2One.
func (ip *IrProfile) Many2One() *Many2One {
	return NewMany2One(ip.Id.Get(), "")
}

// CreateIrProfile creates a new ir.profile model and returns its id.
func (c *Client) CreateIrProfile(ip *IrProfile) (int64, error) {
	return c.Create(IrProfileModel, ip)
}

// UpdateIrProfile updates an existing ir.profile record.
func (c *Client) UpdateIrProfile(ip *IrProfile) error {
	return c.UpdateIrProfiles([]int64{ip.Id.Get()}, ip)
}

// UpdateIrProfiles updates existing ir.profile records.
// All records (represented by ids) will be updated by ip values.
func (c *Client) UpdateIrProfiles(ids []int64, ip *IrProfile) error {
	return c.Update(IrProfileModel, ids, ip)
}

// DeleteIrProfile deletes an existing ir.profile record.
func (c *Client) DeleteIrProfile(id int64) error {
	return c.DeleteIrProfiles([]int64{id})
}

// DeleteIrProfiles deletes existing ir.profile records.
func (c *Client) DeleteIrProfiles(ids []int64) error {
	return c.Delete(IrProfileModel, ids)
}

// GetIrProfile gets ir.profile existing record.
func (c *Client) GetIrProfile(id int64) (*IrProfile, error) {
	ips, err := c.GetIrProfiles([]int64{id})
	if err != nil {
		return nil, err
	}
	if ips != nil && len(*ips) > 0 {
		return &((*ips)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.profile not found", id)
}

// GetIrProfiles gets ir.profile existing records.
func (c *Client) GetIrProfiles(ids []int64) (*IrProfiles, error) {
	ips := &IrProfiles{}
	if err := c.Read(IrProfileModel, ids, nil, ips); err != nil {
		return nil, err
	}
	return ips, nil
}

// FindIrProfile finds ir.profile record by querying it with criteria.
func (c *Client) FindIrProfile(criteria *Criteria) (*IrProfile, error) {
	ips := &IrProfiles{}
	if err := c.SearchRead(IrProfileModel, criteria, NewOptions().Limit(1), ips); err != nil {
		return nil, err
	}
	if ips != nil && len(*ips) > 0 {
		return &((*ips)[0]), nil
	}
	return nil, fmt.Errorf("ir.profile was not found with criteria %v", criteria)
}

// FindIrProfiles finds ir.profile records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrProfiles(criteria *Criteria, options *Options) (*IrProfiles, error) {
	ips := &IrProfiles{}
	if err := c.SearchRead(IrProfileModel, criteria, options, ips); err != nil {
		return nil, err
	}
	return ips, nil
}

// FindIrProfileIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrProfileIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrProfileModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrProfileId finds record id by querying it with criteria.
func (c *Client) FindIrProfileId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrProfileModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.profile was not found with criteria %v and options %v", criteria, options)
}
