package odoo

import (
	"fmt"
)

// StockOrderpointSnooze represents stock.orderpoint.snooze model.
type StockOrderpointSnooze struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	OrderpointIds  *Relation  `xmlrpc:"orderpoint_ids,omptempty"`
	PredefinedDate *Selection `xmlrpc:"predefined_date,omptempty"`
	SnoozedUntil   *Time      `xmlrpc:"snoozed_until,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// StockOrderpointSnoozes represents array of stock.orderpoint.snooze model.
type StockOrderpointSnoozes []StockOrderpointSnooze

// StockOrderpointSnoozeModel is the odoo model name.
const StockOrderpointSnoozeModel = "stock.orderpoint.snooze"

// Many2One convert StockOrderpointSnooze to *Many2One.
func (sos *StockOrderpointSnooze) Many2One() *Many2One {
	return NewMany2One(sos.Id.Get(), "")
}

// CreateStockOrderpointSnooze creates a new stock.orderpoint.snooze model and returns its id.
func (c *Client) CreateStockOrderpointSnooze(sos *StockOrderpointSnooze) (int64, error) {
	return c.Create(StockOrderpointSnoozeModel, sos)
}

// UpdateStockOrderpointSnooze updates an existing stock.orderpoint.snooze record.
func (c *Client) UpdateStockOrderpointSnooze(sos *StockOrderpointSnooze) error {
	return c.UpdateStockOrderpointSnoozes([]int64{sos.Id.Get()}, sos)
}

// UpdateStockOrderpointSnoozes updates existing stock.orderpoint.snooze records.
// All records (represented by ids) will be updated by sos values.
func (c *Client) UpdateStockOrderpointSnoozes(ids []int64, sos *StockOrderpointSnooze) error {
	return c.Update(StockOrderpointSnoozeModel, ids, sos)
}

// DeleteStockOrderpointSnooze deletes an existing stock.orderpoint.snooze record.
func (c *Client) DeleteStockOrderpointSnooze(id int64) error {
	return c.DeleteStockOrderpointSnoozes([]int64{id})
}

// DeleteStockOrderpointSnoozes deletes existing stock.orderpoint.snooze records.
func (c *Client) DeleteStockOrderpointSnoozes(ids []int64) error {
	return c.Delete(StockOrderpointSnoozeModel, ids)
}

// GetStockOrderpointSnooze gets stock.orderpoint.snooze existing record.
func (c *Client) GetStockOrderpointSnooze(id int64) (*StockOrderpointSnooze, error) {
	soss, err := c.GetStockOrderpointSnoozes([]int64{id})
	if err != nil {
		return nil, err
	}
	if soss != nil && len(*soss) > 0 {
		return &((*soss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.orderpoint.snooze not found", id)
}

// GetStockOrderpointSnoozes gets stock.orderpoint.snooze existing records.
func (c *Client) GetStockOrderpointSnoozes(ids []int64) (*StockOrderpointSnoozes, error) {
	soss := &StockOrderpointSnoozes{}
	if err := c.Read(StockOrderpointSnoozeModel, ids, nil, soss); err != nil {
		return nil, err
	}
	return soss, nil
}

// FindStockOrderpointSnooze finds stock.orderpoint.snooze record by querying it with criteria.
func (c *Client) FindStockOrderpointSnooze(criteria *Criteria) (*StockOrderpointSnooze, error) {
	soss := &StockOrderpointSnoozes{}
	if err := c.SearchRead(StockOrderpointSnoozeModel, criteria, NewOptions().Limit(1), soss); err != nil {
		return nil, err
	}
	if soss != nil && len(*soss) > 0 {
		return &((*soss)[0]), nil
	}
	return nil, fmt.Errorf("stock.orderpoint.snooze was not found with criteria %v", criteria)
}

// FindStockOrderpointSnoozes finds stock.orderpoint.snooze records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockOrderpointSnoozes(criteria *Criteria, options *Options) (*StockOrderpointSnoozes, error) {
	soss := &StockOrderpointSnoozes{}
	if err := c.SearchRead(StockOrderpointSnoozeModel, criteria, options, soss); err != nil {
		return nil, err
	}
	return soss, nil
}

// FindStockOrderpointSnoozeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockOrderpointSnoozeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockOrderpointSnoozeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockOrderpointSnoozeId finds record id by querying it with criteria.
func (c *Client) FindStockOrderpointSnoozeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockOrderpointSnoozeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.orderpoint.snooze was not found with criteria %v and options %v", criteria, options)
}
