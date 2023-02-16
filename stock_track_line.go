package odoo

import (
	"fmt"
)

// StockTrackLine represents stock.track.line model.
type StockTrackLine struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName        *String    `xmlrpc:"display_name,omptempty"`
	Id                 *Int       `xmlrpc:"id,omptempty"`
	ProductDisplayName *String    `xmlrpc:"product_display_name,omptempty"`
	ProductId          *Many2One  `xmlrpc:"product_id,omptempty"`
	Tracking           *Selection `xmlrpc:"tracking,omptempty"`
	WizardId           *Many2One  `xmlrpc:"wizard_id,omptempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// StockTrackLines represents array of stock.track.line model.
type StockTrackLines []StockTrackLine

// StockTrackLineModel is the odoo model name.
const StockTrackLineModel = "stock.track.line"

// Many2One convert StockTrackLine to *Many2One.
func (stl *StockTrackLine) Many2One() *Many2One {
	return NewMany2One(stl.Id.Get(), "")
}

// CreateStockTrackLine creates a new stock.track.line model and returns its id.
func (c *Client) CreateStockTrackLine(stl *StockTrackLine) (int64, error) {
	return c.Create(StockTrackLineModel, stl)
}

// UpdateStockTrackLine updates an existing stock.track.line record.
func (c *Client) UpdateStockTrackLine(stl *StockTrackLine) error {
	return c.UpdateStockTrackLines([]int64{stl.Id.Get()}, stl)
}

// UpdateStockTrackLines updates existing stock.track.line records.
// All records (represented by ids) will be updated by stl values.
func (c *Client) UpdateStockTrackLines(ids []int64, stl *StockTrackLine) error {
	return c.Update(StockTrackLineModel, ids, stl)
}

// DeleteStockTrackLine deletes an existing stock.track.line record.
func (c *Client) DeleteStockTrackLine(id int64) error {
	return c.DeleteStockTrackLines([]int64{id})
}

// DeleteStockTrackLines deletes existing stock.track.line records.
func (c *Client) DeleteStockTrackLines(ids []int64) error {
	return c.Delete(StockTrackLineModel, ids)
}

// GetStockTrackLine gets stock.track.line existing record.
func (c *Client) GetStockTrackLine(id int64) (*StockTrackLine, error) {
	stls, err := c.GetStockTrackLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if stls != nil && len(*stls) > 0 {
		return &((*stls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.track.line not found", id)
}

// GetStockTrackLines gets stock.track.line existing records.
func (c *Client) GetStockTrackLines(ids []int64) (*StockTrackLines, error) {
	stls := &StockTrackLines{}
	if err := c.Read(StockTrackLineModel, ids, nil, stls); err != nil {
		return nil, err
	}
	return stls, nil
}

// FindStockTrackLine finds stock.track.line record by querying it with criteria.
func (c *Client) FindStockTrackLine(criteria *Criteria) (*StockTrackLine, error) {
	stls := &StockTrackLines{}
	if err := c.SearchRead(StockTrackLineModel, criteria, NewOptions().Limit(1), stls); err != nil {
		return nil, err
	}
	if stls != nil && len(*stls) > 0 {
		return &((*stls)[0]), nil
	}
	return nil, fmt.Errorf("stock.track.line was not found with criteria %v", criteria)
}

// FindStockTrackLines finds stock.track.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockTrackLines(criteria *Criteria, options *Options) (*StockTrackLines, error) {
	stls := &StockTrackLines{}
	if err := c.SearchRead(StockTrackLineModel, criteria, options, stls); err != nil {
		return nil, err
	}
	return stls, nil
}

// FindStockTrackLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockTrackLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockTrackLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockTrackLineId finds record id by querying it with criteria.
func (c *Client) FindStockTrackLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockTrackLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.track.line was not found with criteria %v and options %v", criteria, options)
}
