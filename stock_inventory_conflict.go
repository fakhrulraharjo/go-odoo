package odoo

import (
	"fmt"
)

// StockInventoryConflict represents stock.inventory.conflict model.
type StockInventoryConflict struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	QuantIds      *Relation `xmlrpc:"quant_ids,omptempty"`
	QuantToFixIds *Relation `xmlrpc:"quant_to_fix_ids,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockInventoryConflicts represents array of stock.inventory.conflict model.
type StockInventoryConflicts []StockInventoryConflict

// StockInventoryConflictModel is the odoo model name.
const StockInventoryConflictModel = "stock.inventory.conflict"

// Many2One convert StockInventoryConflict to *Many2One.
func (sic *StockInventoryConflict) Many2One() *Many2One {
	return NewMany2One(sic.Id.Get(), "")
}

// CreateStockInventoryConflict creates a new stock.inventory.conflict model and returns its id.
func (c *Client) CreateStockInventoryConflict(sic *StockInventoryConflict) (int64, error) {
	return c.Create(StockInventoryConflictModel, sic)
}

// UpdateStockInventoryConflict updates an existing stock.inventory.conflict record.
func (c *Client) UpdateStockInventoryConflict(sic *StockInventoryConflict) error {
	return c.UpdateStockInventoryConflicts([]int64{sic.Id.Get()}, sic)
}

// UpdateStockInventoryConflicts updates existing stock.inventory.conflict records.
// All records (represented by ids) will be updated by sic values.
func (c *Client) UpdateStockInventoryConflicts(ids []int64, sic *StockInventoryConflict) error {
	return c.Update(StockInventoryConflictModel, ids, sic)
}

// DeleteStockInventoryConflict deletes an existing stock.inventory.conflict record.
func (c *Client) DeleteStockInventoryConflict(id int64) error {
	return c.DeleteStockInventoryConflicts([]int64{id})
}

// DeleteStockInventoryConflicts deletes existing stock.inventory.conflict records.
func (c *Client) DeleteStockInventoryConflicts(ids []int64) error {
	return c.Delete(StockInventoryConflictModel, ids)
}

// GetStockInventoryConflict gets stock.inventory.conflict existing record.
func (c *Client) GetStockInventoryConflict(id int64) (*StockInventoryConflict, error) {
	sics, err := c.GetStockInventoryConflicts([]int64{id})
	if err != nil {
		return nil, err
	}
	if sics != nil && len(*sics) > 0 {
		return &((*sics)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.inventory.conflict not found", id)
}

// GetStockInventoryConflicts gets stock.inventory.conflict existing records.
func (c *Client) GetStockInventoryConflicts(ids []int64) (*StockInventoryConflicts, error) {
	sics := &StockInventoryConflicts{}
	if err := c.Read(StockInventoryConflictModel, ids, nil, sics); err != nil {
		return nil, err
	}
	return sics, nil
}

// FindStockInventoryConflict finds stock.inventory.conflict record by querying it with criteria.
func (c *Client) FindStockInventoryConflict(criteria *Criteria) (*StockInventoryConflict, error) {
	sics := &StockInventoryConflicts{}
	if err := c.SearchRead(StockInventoryConflictModel, criteria, NewOptions().Limit(1), sics); err != nil {
		return nil, err
	}
	if sics != nil && len(*sics) > 0 {
		return &((*sics)[0]), nil
	}
	return nil, fmt.Errorf("stock.inventory.conflict was not found with criteria %v", criteria)
}

// FindStockInventoryConflicts finds stock.inventory.conflict records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockInventoryConflicts(criteria *Criteria, options *Options) (*StockInventoryConflicts, error) {
	sics := &StockInventoryConflicts{}
	if err := c.SearchRead(StockInventoryConflictModel, criteria, options, sics); err != nil {
		return nil, err
	}
	return sics, nil
}

// FindStockInventoryConflictIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockInventoryConflictIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockInventoryConflictModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockInventoryConflictId finds record id by querying it with criteria.
func (c *Client) FindStockInventoryConflictId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockInventoryConflictModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.inventory.conflict was not found with criteria %v and options %v", criteria, options)
}
