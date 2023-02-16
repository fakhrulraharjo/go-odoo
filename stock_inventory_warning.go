package odoo

import (
	"fmt"
)

// StockInventoryWarning represents stock.inventory.warning model.
type StockInventoryWarning struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	QuantIds    *Relation `xmlrpc:"quant_ids,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockInventoryWarnings represents array of stock.inventory.warning model.
type StockInventoryWarnings []StockInventoryWarning

// StockInventoryWarningModel is the odoo model name.
const StockInventoryWarningModel = "stock.inventory.warning"

// Many2One convert StockInventoryWarning to *Many2One.
func (siw *StockInventoryWarning) Many2One() *Many2One {
	return NewMany2One(siw.Id.Get(), "")
}

// CreateStockInventoryWarning creates a new stock.inventory.warning model and returns its id.
func (c *Client) CreateStockInventoryWarning(siw *StockInventoryWarning) (int64, error) {
	return c.Create(StockInventoryWarningModel, siw)
}

// UpdateStockInventoryWarning updates an existing stock.inventory.warning record.
func (c *Client) UpdateStockInventoryWarning(siw *StockInventoryWarning) error {
	return c.UpdateStockInventoryWarnings([]int64{siw.Id.Get()}, siw)
}

// UpdateStockInventoryWarnings updates existing stock.inventory.warning records.
// All records (represented by ids) will be updated by siw values.
func (c *Client) UpdateStockInventoryWarnings(ids []int64, siw *StockInventoryWarning) error {
	return c.Update(StockInventoryWarningModel, ids, siw)
}

// DeleteStockInventoryWarning deletes an existing stock.inventory.warning record.
func (c *Client) DeleteStockInventoryWarning(id int64) error {
	return c.DeleteStockInventoryWarnings([]int64{id})
}

// DeleteStockInventoryWarnings deletes existing stock.inventory.warning records.
func (c *Client) DeleteStockInventoryWarnings(ids []int64) error {
	return c.Delete(StockInventoryWarningModel, ids)
}

// GetStockInventoryWarning gets stock.inventory.warning existing record.
func (c *Client) GetStockInventoryWarning(id int64) (*StockInventoryWarning, error) {
	siws, err := c.GetStockInventoryWarnings([]int64{id})
	if err != nil {
		return nil, err
	}
	if siws != nil && len(*siws) > 0 {
		return &((*siws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.inventory.warning not found", id)
}

// GetStockInventoryWarnings gets stock.inventory.warning existing records.
func (c *Client) GetStockInventoryWarnings(ids []int64) (*StockInventoryWarnings, error) {
	siws := &StockInventoryWarnings{}
	if err := c.Read(StockInventoryWarningModel, ids, nil, siws); err != nil {
		return nil, err
	}
	return siws, nil
}

// FindStockInventoryWarning finds stock.inventory.warning record by querying it with criteria.
func (c *Client) FindStockInventoryWarning(criteria *Criteria) (*StockInventoryWarning, error) {
	siws := &StockInventoryWarnings{}
	if err := c.SearchRead(StockInventoryWarningModel, criteria, NewOptions().Limit(1), siws); err != nil {
		return nil, err
	}
	if siws != nil && len(*siws) > 0 {
		return &((*siws)[0]), nil
	}
	return nil, fmt.Errorf("stock.inventory.warning was not found with criteria %v", criteria)
}

// FindStockInventoryWarnings finds stock.inventory.warning records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockInventoryWarnings(criteria *Criteria, options *Options) (*StockInventoryWarnings, error) {
	siws := &StockInventoryWarnings{}
	if err := c.SearchRead(StockInventoryWarningModel, criteria, options, siws); err != nil {
		return nil, err
	}
	return siws, nil
}

// FindStockInventoryWarningIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockInventoryWarningIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockInventoryWarningModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockInventoryWarningId finds record id by querying it with criteria.
func (c *Client) FindStockInventoryWarningId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockInventoryWarningModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.inventory.warning was not found with criteria %v and options %v", criteria, options)
}
