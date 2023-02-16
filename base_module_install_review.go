package odoo

import (
	"fmt"
)

// BaseModuleInstallReview represents base.module.install.review model.
type BaseModuleInstallReview struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName        *String   `xmlrpc:"display_name,omptempty"`
	Id                 *Int      `xmlrpc:"id,omptempty"`
	ModuleId           *Many2One `xmlrpc:"module_id,omptempty"`
	ModuleIds          *Relation `xmlrpc:"module_ids,omptempty"`
	ModulesDescription *String   `xmlrpc:"modules_description,omptempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omptempty"`
}

// BaseModuleInstallReviews represents array of base.module.install.review model.
type BaseModuleInstallReviews []BaseModuleInstallReview

// BaseModuleInstallReviewModel is the odoo model name.
const BaseModuleInstallReviewModel = "base.module.install.review"

// Many2One convert BaseModuleInstallReview to *Many2One.
func (bmir *BaseModuleInstallReview) Many2One() *Many2One {
	return NewMany2One(bmir.Id.Get(), "")
}

// CreateBaseModuleInstallReview creates a new base.module.install.review model and returns its id.
func (c *Client) CreateBaseModuleInstallReview(bmir *BaseModuleInstallReview) (int64, error) {
	return c.Create(BaseModuleInstallReviewModel, bmir)
}

// UpdateBaseModuleInstallReview updates an existing base.module.install.review record.
func (c *Client) UpdateBaseModuleInstallReview(bmir *BaseModuleInstallReview) error {
	return c.UpdateBaseModuleInstallReviews([]int64{bmir.Id.Get()}, bmir)
}

// UpdateBaseModuleInstallReviews updates existing base.module.install.review records.
// All records (represented by ids) will be updated by bmir values.
func (c *Client) UpdateBaseModuleInstallReviews(ids []int64, bmir *BaseModuleInstallReview) error {
	return c.Update(BaseModuleInstallReviewModel, ids, bmir)
}

// DeleteBaseModuleInstallReview deletes an existing base.module.install.review record.
func (c *Client) DeleteBaseModuleInstallReview(id int64) error {
	return c.DeleteBaseModuleInstallReviews([]int64{id})
}

// DeleteBaseModuleInstallReviews deletes existing base.module.install.review records.
func (c *Client) DeleteBaseModuleInstallReviews(ids []int64) error {
	return c.Delete(BaseModuleInstallReviewModel, ids)
}

// GetBaseModuleInstallReview gets base.module.install.review existing record.
func (c *Client) GetBaseModuleInstallReview(id int64) (*BaseModuleInstallReview, error) {
	bmirs, err := c.GetBaseModuleInstallReviews([]int64{id})
	if err != nil {
		return nil, err
	}
	if bmirs != nil && len(*bmirs) > 0 {
		return &((*bmirs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base.module.install.review not found", id)
}

// GetBaseModuleInstallReviews gets base.module.install.review existing records.
func (c *Client) GetBaseModuleInstallReviews(ids []int64) (*BaseModuleInstallReviews, error) {
	bmirs := &BaseModuleInstallReviews{}
	if err := c.Read(BaseModuleInstallReviewModel, ids, nil, bmirs); err != nil {
		return nil, err
	}
	return bmirs, nil
}

// FindBaseModuleInstallReview finds base.module.install.review record by querying it with criteria.
func (c *Client) FindBaseModuleInstallReview(criteria *Criteria) (*BaseModuleInstallReview, error) {
	bmirs := &BaseModuleInstallReviews{}
	if err := c.SearchRead(BaseModuleInstallReviewModel, criteria, NewOptions().Limit(1), bmirs); err != nil {
		return nil, err
	}
	if bmirs != nil && len(*bmirs) > 0 {
		return &((*bmirs)[0]), nil
	}
	return nil, fmt.Errorf("base.module.install.review was not found with criteria %v", criteria)
}

// FindBaseModuleInstallReviews finds base.module.install.review records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseModuleInstallReviews(criteria *Criteria, options *Options) (*BaseModuleInstallReviews, error) {
	bmirs := &BaseModuleInstallReviews{}
	if err := c.SearchRead(BaseModuleInstallReviewModel, criteria, options, bmirs); err != nil {
		return nil, err
	}
	return bmirs, nil
}

// FindBaseModuleInstallReviewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseModuleInstallReviewIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseModuleInstallReviewModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseModuleInstallReviewId finds record id by querying it with criteria.
func (c *Client) FindBaseModuleInstallReviewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseModuleInstallReviewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base.module.install.review was not found with criteria %v and options %v", criteria, options)
}
