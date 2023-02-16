package odoo

import (
	"fmt"
)

// StockInventoryAdjustmentName represents stock.inventory.adjustment.name model.
type StockInventoryAdjustmentName struct {
	LastUpdate              *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate              *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid               *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName             *String   `xmlrpc:"display_name,omptempty"`
	Id                      *Int      `xmlrpc:"id,omptempty"`
	InventoryAdjustmentName *String   `xmlrpc:"inventory_adjustment_name,omptempty"`
	QuantIds                *Relation `xmlrpc:"quant_ids,omptempty"`
	ShowInfo                *Bool     `xmlrpc:"show_info,omptempty"`
	WriteDate               *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockInventoryAdjustmentNames represents array of stock.inventory.adjustment.name model.
type StockInventoryAdjustmentNames []StockInventoryAdjustmentName

// StockInventoryAdjustmentNameModel is the odoo model name.
const StockInventoryAdjustmentNameModel = "stock.inventory.adjustment.name"

// Many2One convert StockInventoryAdjustmentName to *Many2One.
func (sian *StockInventoryAdjustmentName) Many2One() *Many2One {
	return NewMany2One(sian.Id.Get(), "")
}

// CreateStockInventoryAdjustmentName creates a new stock.inventory.adjustment.name model and returns its id.
func (c *Client) CreateStockInventoryAdjustmentName(sian *StockInventoryAdjustmentName) (int64, error) {
	return c.Create(StockInventoryAdjustmentNameModel, sian)
}

// UpdateStockInventoryAdjustmentName updates an existing stock.inventory.adjustment.name record.
func (c *Client) UpdateStockInventoryAdjustmentName(sian *StockInventoryAdjustmentName) error {
	return c.UpdateStockInventoryAdjustmentNames([]int64{sian.Id.Get()}, sian)
}

// UpdateStockInventoryAdjustmentNames updates existing stock.inventory.adjustment.name records.
// All records (represented by ids) will be updated by sian values.
func (c *Client) UpdateStockInventoryAdjustmentNames(ids []int64, sian *StockInventoryAdjustmentName) error {
	return c.Update(StockInventoryAdjustmentNameModel, ids, sian)
}

// DeleteStockInventoryAdjustmentName deletes an existing stock.inventory.adjustment.name record.
func (c *Client) DeleteStockInventoryAdjustmentName(id int64) error {
	return c.DeleteStockInventoryAdjustmentNames([]int64{id})
}

// DeleteStockInventoryAdjustmentNames deletes existing stock.inventory.adjustment.name records.
func (c *Client) DeleteStockInventoryAdjustmentNames(ids []int64) error {
	return c.Delete(StockInventoryAdjustmentNameModel, ids)
}

// GetStockInventoryAdjustmentName gets stock.inventory.adjustment.name existing record.
func (c *Client) GetStockInventoryAdjustmentName(id int64) (*StockInventoryAdjustmentName, error) {
	sians, err := c.GetStockInventoryAdjustmentNames([]int64{id})
	if err != nil {
		return nil, err
	}
	if sians != nil && len(*sians) > 0 {
		return &((*sians)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.inventory.adjustment.name not found", id)
}

// GetStockInventoryAdjustmentNames gets stock.inventory.adjustment.name existing records.
func (c *Client) GetStockInventoryAdjustmentNames(ids []int64) (*StockInventoryAdjustmentNames, error) {
	sians := &StockInventoryAdjustmentNames{}
	if err := c.Read(StockInventoryAdjustmentNameModel, ids, nil, sians); err != nil {
		return nil, err
	}
	return sians, nil
}

// FindStockInventoryAdjustmentName finds stock.inventory.adjustment.name record by querying it with criteria.
func (c *Client) FindStockInventoryAdjustmentName(criteria *Criteria) (*StockInventoryAdjustmentName, error) {
	sians := &StockInventoryAdjustmentNames{}
	if err := c.SearchRead(StockInventoryAdjustmentNameModel, criteria, NewOptions().Limit(1), sians); err != nil {
		return nil, err
	}
	if sians != nil && len(*sians) > 0 {
		return &((*sians)[0]), nil
	}
	return nil, fmt.Errorf("stock.inventory.adjustment.name was not found with criteria %v", criteria)
}

// FindStockInventoryAdjustmentNames finds stock.inventory.adjustment.name records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockInventoryAdjustmentNames(criteria *Criteria, options *Options) (*StockInventoryAdjustmentNames, error) {
	sians := &StockInventoryAdjustmentNames{}
	if err := c.SearchRead(StockInventoryAdjustmentNameModel, criteria, options, sians); err != nil {
		return nil, err
	}
	return sians, nil
}

// FindStockInventoryAdjustmentNameIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockInventoryAdjustmentNameIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockInventoryAdjustmentNameModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockInventoryAdjustmentNameId finds record id by querying it with criteria.
func (c *Client) FindStockInventoryAdjustmentNameId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockInventoryAdjustmentNameModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.inventory.adjustment.name was not found with criteria %v and options %v", criteria, options)
}
