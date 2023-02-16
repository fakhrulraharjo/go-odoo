package odoo

import (
	"fmt"
)

// StockRequestCount represents stock.request.count model.
type StockRequestCount struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	AccountingDate *Time      `xmlrpc:"accounting_date,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	InventoryDate  *Time      `xmlrpc:"inventory_date,omptempty"`
	QuantIds       *Relation  `xmlrpc:"quant_ids,omptempty"`
	SetCount       *Selection `xmlrpc:"set_count,omptempty"`
	UserId         *Many2One  `xmlrpc:"user_id,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// StockRequestCounts represents array of stock.request.count model.
type StockRequestCounts []StockRequestCount

// StockRequestCountModel is the odoo model name.
const StockRequestCountModel = "stock.request.count"

// Many2One convert StockRequestCount to *Many2One.
func (src *StockRequestCount) Many2One() *Many2One {
	return NewMany2One(src.Id.Get(), "")
}

// CreateStockRequestCount creates a new stock.request.count model and returns its id.
func (c *Client) CreateStockRequestCount(src *StockRequestCount) (int64, error) {
	return c.Create(StockRequestCountModel, src)
}

// UpdateStockRequestCount updates an existing stock.request.count record.
func (c *Client) UpdateStockRequestCount(src *StockRequestCount) error {
	return c.UpdateStockRequestCounts([]int64{src.Id.Get()}, src)
}

// UpdateStockRequestCounts updates existing stock.request.count records.
// All records (represented by ids) will be updated by src values.
func (c *Client) UpdateStockRequestCounts(ids []int64, src *StockRequestCount) error {
	return c.Update(StockRequestCountModel, ids, src)
}

// DeleteStockRequestCount deletes an existing stock.request.count record.
func (c *Client) DeleteStockRequestCount(id int64) error {
	return c.DeleteStockRequestCounts([]int64{id})
}

// DeleteStockRequestCounts deletes existing stock.request.count records.
func (c *Client) DeleteStockRequestCounts(ids []int64) error {
	return c.Delete(StockRequestCountModel, ids)
}

// GetStockRequestCount gets stock.request.count existing record.
func (c *Client) GetStockRequestCount(id int64) (*StockRequestCount, error) {
	srcs, err := c.GetStockRequestCounts([]int64{id})
	if err != nil {
		return nil, err
	}
	if srcs != nil && len(*srcs) > 0 {
		return &((*srcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.request.count not found", id)
}

// GetStockRequestCounts gets stock.request.count existing records.
func (c *Client) GetStockRequestCounts(ids []int64) (*StockRequestCounts, error) {
	srcs := &StockRequestCounts{}
	if err := c.Read(StockRequestCountModel, ids, nil, srcs); err != nil {
		return nil, err
	}
	return srcs, nil
}

// FindStockRequestCount finds stock.request.count record by querying it with criteria.
func (c *Client) FindStockRequestCount(criteria *Criteria) (*StockRequestCount, error) {
	srcs := &StockRequestCounts{}
	if err := c.SearchRead(StockRequestCountModel, criteria, NewOptions().Limit(1), srcs); err != nil {
		return nil, err
	}
	if srcs != nil && len(*srcs) > 0 {
		return &((*srcs)[0]), nil
	}
	return nil, fmt.Errorf("stock.request.count was not found with criteria %v", criteria)
}

// FindStockRequestCounts finds stock.request.count records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockRequestCounts(criteria *Criteria, options *Options) (*StockRequestCounts, error) {
	srcs := &StockRequestCounts{}
	if err := c.SearchRead(StockRequestCountModel, criteria, options, srcs); err != nil {
		return nil, err
	}
	return srcs, nil
}

// FindStockRequestCountIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockRequestCountIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockRequestCountModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockRequestCountId finds record id by querying it with criteria.
func (c *Client) FindStockRequestCountId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockRequestCountModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.request.count was not found with criteria %v and options %v", criteria, options)
}
