package odoo

import (
	"fmt"
)

// AnalyticMixin represents analytic.mixin model.
type AnalyticMixin struct {
	AnalyticDistribution interface{} `xmlrpc:"analytic_distribution,omptempty"`
	AnalyticPrecision    *Int        `xmlrpc:"analytic_precision,omptempty"`
}

// AnalyticMixins represents array of analytic.mixin model.
type AnalyticMixins []AnalyticMixin

// AnalyticMixinModel is the odoo model name.
const AnalyticMixinModel = "analytic.mixin"

// Many2One convert AnalyticMixin to *Many2One.
func (am *AnalyticMixin) Many2One() *Many2One {
	return NewMany2One(am.Id.Get(), "")
}

// CreateAnalyticMixin creates a new analytic.mixin model and returns its id.
func (c *Client) CreateAnalyticMixin(am *AnalyticMixin) (int64, error) {
	return c.Create(AnalyticMixinModel, am)
}

// UpdateAnalyticMixin updates an existing analytic.mixin record.
func (c *Client) UpdateAnalyticMixin(am *AnalyticMixin) error {
	return c.UpdateAnalyticMixins([]int64{am.Id.Get()}, am)
}

// UpdateAnalyticMixins updates existing analytic.mixin records.
// All records (represented by ids) will be updated by am values.
func (c *Client) UpdateAnalyticMixins(ids []int64, am *AnalyticMixin) error {
	return c.Update(AnalyticMixinModel, ids, am)
}

// DeleteAnalyticMixin deletes an existing analytic.mixin record.
func (c *Client) DeleteAnalyticMixin(id int64) error {
	return c.DeleteAnalyticMixins([]int64{id})
}

// DeleteAnalyticMixins deletes existing analytic.mixin records.
func (c *Client) DeleteAnalyticMixins(ids []int64) error {
	return c.Delete(AnalyticMixinModel, ids)
}

// GetAnalyticMixin gets analytic.mixin existing record.
func (c *Client) GetAnalyticMixin(id int64) (*AnalyticMixin, error) {
	ams, err := c.GetAnalyticMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if ams != nil && len(*ams) > 0 {
		return &((*ams)[0]), nil
	}
	return nil, fmt.Errorf("id %v of analytic.mixin not found", id)
}

// GetAnalyticMixins gets analytic.mixin existing records.
func (c *Client) GetAnalyticMixins(ids []int64) (*AnalyticMixins, error) {
	ams := &AnalyticMixins{}
	if err := c.Read(AnalyticMixinModel, ids, nil, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAnalyticMixin finds analytic.mixin record by querying it with criteria.
func (c *Client) FindAnalyticMixin(criteria *Criteria) (*AnalyticMixin, error) {
	ams := &AnalyticMixins{}
	if err := c.SearchRead(AnalyticMixinModel, criteria, NewOptions().Limit(1), ams); err != nil {
		return nil, err
	}
	if ams != nil && len(*ams) > 0 {
		return &((*ams)[0]), nil
	}
	return nil, fmt.Errorf("analytic.mixin was not found with criteria %v", criteria)
}

// FindAnalyticMixins finds analytic.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAnalyticMixins(criteria *Criteria, options *Options) (*AnalyticMixins, error) {
	ams := &AnalyticMixins{}
	if err := c.SearchRead(AnalyticMixinModel, criteria, options, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAnalyticMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAnalyticMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AnalyticMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAnalyticMixinId finds record id by querying it with criteria.
func (c *Client) FindAnalyticMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AnalyticMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("analytic.mixin was not found with criteria %v and options %v", criteria, options)
}
