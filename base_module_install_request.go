package odoo

import (
	"fmt"
)

// BaseModuleInstallRequest represents base.module.install.request model.
type BaseModuleInstallRequest struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	BodyHtml    *String   `xmlrpc:"body_html,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	ModuleId    *Many2One `xmlrpc:"module_id,omptempty"`
	UserId      *Many2One `xmlrpc:"user_id,omptempty"`
	UserIds     *Relation `xmlrpc:"user_ids,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// BaseModuleInstallRequests represents array of base.module.install.request model.
type BaseModuleInstallRequests []BaseModuleInstallRequest

// BaseModuleInstallRequestModel is the odoo model name.
const BaseModuleInstallRequestModel = "base.module.install.request"

// Many2One convert BaseModuleInstallRequest to *Many2One.
func (bmir *BaseModuleInstallRequest) Many2One() *Many2One {
	return NewMany2One(bmir.Id.Get(), "")
}

// CreateBaseModuleInstallRequest creates a new base.module.install.request model and returns its id.
func (c *Client) CreateBaseModuleInstallRequest(bmir *BaseModuleInstallRequest) (int64, error) {
	return c.Create(BaseModuleInstallRequestModel, bmir)
}

// UpdateBaseModuleInstallRequest updates an existing base.module.install.request record.
func (c *Client) UpdateBaseModuleInstallRequest(bmir *BaseModuleInstallRequest) error {
	return c.UpdateBaseModuleInstallRequests([]int64{bmir.Id.Get()}, bmir)
}

// UpdateBaseModuleInstallRequests updates existing base.module.install.request records.
// All records (represented by ids) will be updated by bmir values.
func (c *Client) UpdateBaseModuleInstallRequests(ids []int64, bmir *BaseModuleInstallRequest) error {
	return c.Update(BaseModuleInstallRequestModel, ids, bmir)
}

// DeleteBaseModuleInstallRequest deletes an existing base.module.install.request record.
func (c *Client) DeleteBaseModuleInstallRequest(id int64) error {
	return c.DeleteBaseModuleInstallRequests([]int64{id})
}

// DeleteBaseModuleInstallRequests deletes existing base.module.install.request records.
func (c *Client) DeleteBaseModuleInstallRequests(ids []int64) error {
	return c.Delete(BaseModuleInstallRequestModel, ids)
}

// GetBaseModuleInstallRequest gets base.module.install.request existing record.
func (c *Client) GetBaseModuleInstallRequest(id int64) (*BaseModuleInstallRequest, error) {
	bmirs, err := c.GetBaseModuleInstallRequests([]int64{id})
	if err != nil {
		return nil, err
	}
	if bmirs != nil && len(*bmirs) > 0 {
		return &((*bmirs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base.module.install.request not found", id)
}

// GetBaseModuleInstallRequests gets base.module.install.request existing records.
func (c *Client) GetBaseModuleInstallRequests(ids []int64) (*BaseModuleInstallRequests, error) {
	bmirs := &BaseModuleInstallRequests{}
	if err := c.Read(BaseModuleInstallRequestModel, ids, nil, bmirs); err != nil {
		return nil, err
	}
	return bmirs, nil
}

// FindBaseModuleInstallRequest finds base.module.install.request record by querying it with criteria.
func (c *Client) FindBaseModuleInstallRequest(criteria *Criteria) (*BaseModuleInstallRequest, error) {
	bmirs := &BaseModuleInstallRequests{}
	if err := c.SearchRead(BaseModuleInstallRequestModel, criteria, NewOptions().Limit(1), bmirs); err != nil {
		return nil, err
	}
	if bmirs != nil && len(*bmirs) > 0 {
		return &((*bmirs)[0]), nil
	}
	return nil, fmt.Errorf("base.module.install.request was not found with criteria %v", criteria)
}

// FindBaseModuleInstallRequests finds base.module.install.request records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseModuleInstallRequests(criteria *Criteria, options *Options) (*BaseModuleInstallRequests, error) {
	bmirs := &BaseModuleInstallRequests{}
	if err := c.SearchRead(BaseModuleInstallRequestModel, criteria, options, bmirs); err != nil {
		return nil, err
	}
	return bmirs, nil
}

// FindBaseModuleInstallRequestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseModuleInstallRequestIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseModuleInstallRequestModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseModuleInstallRequestId finds record id by querying it with criteria.
func (c *Client) FindBaseModuleInstallRequestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseModuleInstallRequestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base.module.install.request was not found with criteria %v and options %v", criteria, options)
}
