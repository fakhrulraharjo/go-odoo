package odoo

import (
	"fmt"
)

// StockBackorderConfirmationLine represents stock.backorder.confirmation.line model.
type StockBackorderConfirmationLine struct {
	LastUpdate              *Time     `xmlrpc:"__last_update,omptempty"`
	BackorderConfirmationId *Many2One `xmlrpc:"backorder_confirmation_id,omptempty"`
	CreateDate              *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid               *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName             *String   `xmlrpc:"display_name,omptempty"`
	Id                      *Int      `xmlrpc:"id,omptempty"`
	PickingId               *Many2One `xmlrpc:"picking_id,omptempty"`
	ToBackorder             *Bool     `xmlrpc:"to_backorder,omptempty"`
	WriteDate               *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockBackorderConfirmationLines represents array of stock.backorder.confirmation.line model.
type StockBackorderConfirmationLines []StockBackorderConfirmationLine

// StockBackorderConfirmationLineModel is the odoo model name.
const StockBackorderConfirmationLineModel = "stock.backorder.confirmation.line"

// Many2One convert StockBackorderConfirmationLine to *Many2One.
func (sbcl *StockBackorderConfirmationLine) Many2One() *Many2One {
	return NewMany2One(sbcl.Id.Get(), "")
}

// CreateStockBackorderConfirmationLine creates a new stock.backorder.confirmation.line model and returns its id.
func (c *Client) CreateStockBackorderConfirmationLine(sbcl *StockBackorderConfirmationLine) (int64, error) {
	return c.Create(StockBackorderConfirmationLineModel, sbcl)
}

// UpdateStockBackorderConfirmationLine updates an existing stock.backorder.confirmation.line record.
func (c *Client) UpdateStockBackorderConfirmationLine(sbcl *StockBackorderConfirmationLine) error {
	return c.UpdateStockBackorderConfirmationLines([]int64{sbcl.Id.Get()}, sbcl)
}

// UpdateStockBackorderConfirmationLines updates existing stock.backorder.confirmation.line records.
// All records (represented by ids) will be updated by sbcl values.
func (c *Client) UpdateStockBackorderConfirmationLines(ids []int64, sbcl *StockBackorderConfirmationLine) error {
	return c.Update(StockBackorderConfirmationLineModel, ids, sbcl)
}

// DeleteStockBackorderConfirmationLine deletes an existing stock.backorder.confirmation.line record.
func (c *Client) DeleteStockBackorderConfirmationLine(id int64) error {
	return c.DeleteStockBackorderConfirmationLines([]int64{id})
}

// DeleteStockBackorderConfirmationLines deletes existing stock.backorder.confirmation.line records.
func (c *Client) DeleteStockBackorderConfirmationLines(ids []int64) error {
	return c.Delete(StockBackorderConfirmationLineModel, ids)
}

// GetStockBackorderConfirmationLine gets stock.backorder.confirmation.line existing record.
func (c *Client) GetStockBackorderConfirmationLine(id int64) (*StockBackorderConfirmationLine, error) {
	sbcls, err := c.GetStockBackorderConfirmationLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if sbcls != nil && len(*sbcls) > 0 {
		return &((*sbcls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.backorder.confirmation.line not found", id)
}

// GetStockBackorderConfirmationLines gets stock.backorder.confirmation.line existing records.
func (c *Client) GetStockBackorderConfirmationLines(ids []int64) (*StockBackorderConfirmationLines, error) {
	sbcls := &StockBackorderConfirmationLines{}
	if err := c.Read(StockBackorderConfirmationLineModel, ids, nil, sbcls); err != nil {
		return nil, err
	}
	return sbcls, nil
}

// FindStockBackorderConfirmationLine finds stock.backorder.confirmation.line record by querying it with criteria.
func (c *Client) FindStockBackorderConfirmationLine(criteria *Criteria) (*StockBackorderConfirmationLine, error) {
	sbcls := &StockBackorderConfirmationLines{}
	if err := c.SearchRead(StockBackorderConfirmationLineModel, criteria, NewOptions().Limit(1), sbcls); err != nil {
		return nil, err
	}
	if sbcls != nil && len(*sbcls) > 0 {
		return &((*sbcls)[0]), nil
	}
	return nil, fmt.Errorf("stock.backorder.confirmation.line was not found with criteria %v", criteria)
}

// FindStockBackorderConfirmationLines finds stock.backorder.confirmation.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockBackorderConfirmationLines(criteria *Criteria, options *Options) (*StockBackorderConfirmationLines, error) {
	sbcls := &StockBackorderConfirmationLines{}
	if err := c.SearchRead(StockBackorderConfirmationLineModel, criteria, options, sbcls); err != nil {
		return nil, err
	}
	return sbcls, nil
}

// FindStockBackorderConfirmationLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockBackorderConfirmationLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockBackorderConfirmationLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockBackorderConfirmationLineId finds record id by querying it with criteria.
func (c *Client) FindStockBackorderConfirmationLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockBackorderConfirmationLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.backorder.confirmation.line was not found with criteria %v and options %v", criteria, options)
}
