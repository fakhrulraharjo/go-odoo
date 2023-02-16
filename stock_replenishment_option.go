package odoo

import (
	"fmt"
)

// StockReplenishmentOption represents stock.replenishment.option model.
type StockReplenishmentOption struct {
	LastUpdate          *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName         *String   `xmlrpc:"display_name,omptempty"`
	FreeQty             *Float    `xmlrpc:"free_qty,omptempty"`
	Id                  *Int      `xmlrpc:"id,omptempty"`
	LeadTime            *String   `xmlrpc:"lead_time,omptempty"`
	LocationId          *Many2One `xmlrpc:"location_id,omptempty"`
	ProductId           *Many2One `xmlrpc:"product_id,omptempty"`
	QtyToOrder          *Float    `xmlrpc:"qty_to_order,omptempty"`
	ReplenishmentInfoId *Many2One `xmlrpc:"replenishment_info_id,omptempty"`
	RouteId             *Many2One `xmlrpc:"route_id,omptempty"`
	Uom                 *String   `xmlrpc:"uom,omptempty"`
	WarehouseId         *Many2One `xmlrpc:"warehouse_id,omptempty"`
	WarningMessage      *String   `xmlrpc:"warning_message,omptempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockReplenishmentOptions represents array of stock.replenishment.option model.
type StockReplenishmentOptions []StockReplenishmentOption

// StockReplenishmentOptionModel is the odoo model name.
const StockReplenishmentOptionModel = "stock.replenishment.option"

// Many2One convert StockReplenishmentOption to *Many2One.
func (sro *StockReplenishmentOption) Many2One() *Many2One {
	return NewMany2One(sro.Id.Get(), "")
}

// CreateStockReplenishmentOption creates a new stock.replenishment.option model and returns its id.
func (c *Client) CreateStockReplenishmentOption(sro *StockReplenishmentOption) (int64, error) {
	return c.Create(StockReplenishmentOptionModel, sro)
}

// UpdateStockReplenishmentOption updates an existing stock.replenishment.option record.
func (c *Client) UpdateStockReplenishmentOption(sro *StockReplenishmentOption) error {
	return c.UpdateStockReplenishmentOptions([]int64{sro.Id.Get()}, sro)
}

// UpdateStockReplenishmentOptions updates existing stock.replenishment.option records.
// All records (represented by ids) will be updated by sro values.
func (c *Client) UpdateStockReplenishmentOptions(ids []int64, sro *StockReplenishmentOption) error {
	return c.Update(StockReplenishmentOptionModel, ids, sro)
}

// DeleteStockReplenishmentOption deletes an existing stock.replenishment.option record.
func (c *Client) DeleteStockReplenishmentOption(id int64) error {
	return c.DeleteStockReplenishmentOptions([]int64{id})
}

// DeleteStockReplenishmentOptions deletes existing stock.replenishment.option records.
func (c *Client) DeleteStockReplenishmentOptions(ids []int64) error {
	return c.Delete(StockReplenishmentOptionModel, ids)
}

// GetStockReplenishmentOption gets stock.replenishment.option existing record.
func (c *Client) GetStockReplenishmentOption(id int64) (*StockReplenishmentOption, error) {
	sros, err := c.GetStockReplenishmentOptions([]int64{id})
	if err != nil {
		return nil, err
	}
	if sros != nil && len(*sros) > 0 {
		return &((*sros)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.replenishment.option not found", id)
}

// GetStockReplenishmentOptions gets stock.replenishment.option existing records.
func (c *Client) GetStockReplenishmentOptions(ids []int64) (*StockReplenishmentOptions, error) {
	sros := &StockReplenishmentOptions{}
	if err := c.Read(StockReplenishmentOptionModel, ids, nil, sros); err != nil {
		return nil, err
	}
	return sros, nil
}

// FindStockReplenishmentOption finds stock.replenishment.option record by querying it with criteria.
func (c *Client) FindStockReplenishmentOption(criteria *Criteria) (*StockReplenishmentOption, error) {
	sros := &StockReplenishmentOptions{}
	if err := c.SearchRead(StockReplenishmentOptionModel, criteria, NewOptions().Limit(1), sros); err != nil {
		return nil, err
	}
	if sros != nil && len(*sros) > 0 {
		return &((*sros)[0]), nil
	}
	return nil, fmt.Errorf("stock.replenishment.option was not found with criteria %v", criteria)
}

// FindStockReplenishmentOptions finds stock.replenishment.option records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReplenishmentOptions(criteria *Criteria, options *Options) (*StockReplenishmentOptions, error) {
	sros := &StockReplenishmentOptions{}
	if err := c.SearchRead(StockReplenishmentOptionModel, criteria, options, sros); err != nil {
		return nil, err
	}
	return sros, nil
}

// FindStockReplenishmentOptionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReplenishmentOptionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockReplenishmentOptionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockReplenishmentOptionId finds record id by querying it with criteria.
func (c *Client) FindStockReplenishmentOptionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockReplenishmentOptionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.replenishment.option was not found with criteria %v and options %v", criteria, options)
}
