package odoo

import (
	"fmt"
)

// StockAssignSerial represents stock.assign.serial model.
type StockAssignSerial struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	MoveId           *Many2One `xmlrpc:"move_id,omptempty"`
	NextSerialCount  *Int      `xmlrpc:"next_serial_count,omptempty"`
	NextSerialNumber *String   `xmlrpc:"next_serial_number,omptempty"`
	ProductId        *Many2One `xmlrpc:"product_id,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockAssignSerials represents array of stock.assign.serial model.
type StockAssignSerials []StockAssignSerial

// StockAssignSerialModel is the odoo model name.
const StockAssignSerialModel = "stock.assign.serial"

// Many2One convert StockAssignSerial to *Many2One.
func (sas *StockAssignSerial) Many2One() *Many2One {
	return NewMany2One(sas.Id.Get(), "")
}

// CreateStockAssignSerial creates a new stock.assign.serial model and returns its id.
func (c *Client) CreateStockAssignSerial(sas *StockAssignSerial) (int64, error) {
	return c.Create(StockAssignSerialModel, sas)
}

// UpdateStockAssignSerial updates an existing stock.assign.serial record.
func (c *Client) UpdateStockAssignSerial(sas *StockAssignSerial) error {
	return c.UpdateStockAssignSerials([]int64{sas.Id.Get()}, sas)
}

// UpdateStockAssignSerials updates existing stock.assign.serial records.
// All records (represented by ids) will be updated by sas values.
func (c *Client) UpdateStockAssignSerials(ids []int64, sas *StockAssignSerial) error {
	return c.Update(StockAssignSerialModel, ids, sas)
}

// DeleteStockAssignSerial deletes an existing stock.assign.serial record.
func (c *Client) DeleteStockAssignSerial(id int64) error {
	return c.DeleteStockAssignSerials([]int64{id})
}

// DeleteStockAssignSerials deletes existing stock.assign.serial records.
func (c *Client) DeleteStockAssignSerials(ids []int64) error {
	return c.Delete(StockAssignSerialModel, ids)
}

// GetStockAssignSerial gets stock.assign.serial existing record.
func (c *Client) GetStockAssignSerial(id int64) (*StockAssignSerial, error) {
	sass, err := c.GetStockAssignSerials([]int64{id})
	if err != nil {
		return nil, err
	}
	if sass != nil && len(*sass) > 0 {
		return &((*sass)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.assign.serial not found", id)
}

// GetStockAssignSerials gets stock.assign.serial existing records.
func (c *Client) GetStockAssignSerials(ids []int64) (*StockAssignSerials, error) {
	sass := &StockAssignSerials{}
	if err := c.Read(StockAssignSerialModel, ids, nil, sass); err != nil {
		return nil, err
	}
	return sass, nil
}

// FindStockAssignSerial finds stock.assign.serial record by querying it with criteria.
func (c *Client) FindStockAssignSerial(criteria *Criteria) (*StockAssignSerial, error) {
	sass := &StockAssignSerials{}
	if err := c.SearchRead(StockAssignSerialModel, criteria, NewOptions().Limit(1), sass); err != nil {
		return nil, err
	}
	if sass != nil && len(*sass) > 0 {
		return &((*sass)[0]), nil
	}
	return nil, fmt.Errorf("stock.assign.serial was not found with criteria %v", criteria)
}

// FindStockAssignSerials finds stock.assign.serial records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockAssignSerials(criteria *Criteria, options *Options) (*StockAssignSerials, error) {
	sass := &StockAssignSerials{}
	if err := c.SearchRead(StockAssignSerialModel, criteria, options, sass); err != nil {
		return nil, err
	}
	return sass, nil
}

// FindStockAssignSerialIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockAssignSerialIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockAssignSerialModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockAssignSerialId finds record id by querying it with criteria.
func (c *Client) FindStockAssignSerialId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockAssignSerialModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.assign.serial was not found with criteria %v and options %v", criteria, options)
}
