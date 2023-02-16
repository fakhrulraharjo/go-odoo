package odoo

import (
	"fmt"
)

// StockPutawayRule represents stock.putaway.rule model.
type StockPutawayRule struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	Active            *Bool     `xmlrpc:"active,omptempty"`
	CategoryId        *Many2One `xmlrpc:"category_id,omptempty"`
	CompanyId         *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	LocationInId      *Many2One `xmlrpc:"location_in_id,omptempty"`
	LocationOutId     *Many2One `xmlrpc:"location_out_id,omptempty"`
	PackageTypeIds    *Relation `xmlrpc:"package_type_ids,omptempty"`
	ProductId         *Many2One `xmlrpc:"product_id,omptempty"`
	Sequence          *Int      `xmlrpc:"sequence,omptempty"`
	StorageCategoryId *Many2One `xmlrpc:"storage_category_id,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockPutawayRules represents array of stock.putaway.rule model.
type StockPutawayRules []StockPutawayRule

// StockPutawayRuleModel is the odoo model name.
const StockPutawayRuleModel = "stock.putaway.rule"

// Many2One convert StockPutawayRule to *Many2One.
func (spr *StockPutawayRule) Many2One() *Many2One {
	return NewMany2One(spr.Id.Get(), "")
}

// CreateStockPutawayRule creates a new stock.putaway.rule model and returns its id.
func (c *Client) CreateStockPutawayRule(spr *StockPutawayRule) (int64, error) {
	return c.Create(StockPutawayRuleModel, spr)
}

// UpdateStockPutawayRule updates an existing stock.putaway.rule record.
func (c *Client) UpdateStockPutawayRule(spr *StockPutawayRule) error {
	return c.UpdateStockPutawayRules([]int64{spr.Id.Get()}, spr)
}

// UpdateStockPutawayRules updates existing stock.putaway.rule records.
// All records (represented by ids) will be updated by spr values.
func (c *Client) UpdateStockPutawayRules(ids []int64, spr *StockPutawayRule) error {
	return c.Update(StockPutawayRuleModel, ids, spr)
}

// DeleteStockPutawayRule deletes an existing stock.putaway.rule record.
func (c *Client) DeleteStockPutawayRule(id int64) error {
	return c.DeleteStockPutawayRules([]int64{id})
}

// DeleteStockPutawayRules deletes existing stock.putaway.rule records.
func (c *Client) DeleteStockPutawayRules(ids []int64) error {
	return c.Delete(StockPutawayRuleModel, ids)
}

// GetStockPutawayRule gets stock.putaway.rule existing record.
func (c *Client) GetStockPutawayRule(id int64) (*StockPutawayRule, error) {
	sprs, err := c.GetStockPutawayRules([]int64{id})
	if err != nil {
		return nil, err
	}
	if sprs != nil && len(*sprs) > 0 {
		return &((*sprs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.putaway.rule not found", id)
}

// GetStockPutawayRules gets stock.putaway.rule existing records.
func (c *Client) GetStockPutawayRules(ids []int64) (*StockPutawayRules, error) {
	sprs := &StockPutawayRules{}
	if err := c.Read(StockPutawayRuleModel, ids, nil, sprs); err != nil {
		return nil, err
	}
	return sprs, nil
}

// FindStockPutawayRule finds stock.putaway.rule record by querying it with criteria.
func (c *Client) FindStockPutawayRule(criteria *Criteria) (*StockPutawayRule, error) {
	sprs := &StockPutawayRules{}
	if err := c.SearchRead(StockPutawayRuleModel, criteria, NewOptions().Limit(1), sprs); err != nil {
		return nil, err
	}
	if sprs != nil && len(*sprs) > 0 {
		return &((*sprs)[0]), nil
	}
	return nil, fmt.Errorf("stock.putaway.rule was not found with criteria %v", criteria)
}

// FindStockPutawayRules finds stock.putaway.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPutawayRules(criteria *Criteria, options *Options) (*StockPutawayRules, error) {
	sprs := &StockPutawayRules{}
	if err := c.SearchRead(StockPutawayRuleModel, criteria, options, sprs); err != nil {
		return nil, err
	}
	return sprs, nil
}

// FindStockPutawayRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPutawayRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockPutawayRuleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockPutawayRuleId finds record id by querying it with criteria.
func (c *Client) FindStockPutawayRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockPutawayRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.putaway.rule was not found with criteria %v and options %v", criteria, options)
}
