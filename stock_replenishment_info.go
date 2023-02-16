package odoo

import (
	"fmt"
)

// StockReplenishmentInfo represents stock.replenishment.info model.
type StockReplenishmentInfo struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	JsonLeadDays             *String   `xmlrpc:"json_lead_days,omptempty"`
	JsonReplenishmentHistory *String   `xmlrpc:"json_replenishment_history,omptempty"`
	OrderpointId             *Many2One `xmlrpc:"orderpoint_id,omptempty"`
	ProductId                *Many2One `xmlrpc:"product_id,omptempty"`
	QtyToOrder               *Float    `xmlrpc:"qty_to_order,omptempty"`
	WarehouseinfoIds         *Relation `xmlrpc:"warehouseinfo_ids,omptempty"`
	WhReplenishmentOptionIds *Relation `xmlrpc:"wh_replenishment_option_ids,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockReplenishmentInfos represents array of stock.replenishment.info model.
type StockReplenishmentInfos []StockReplenishmentInfo

// StockReplenishmentInfoModel is the odoo model name.
const StockReplenishmentInfoModel = "stock.replenishment.info"

// Many2One convert StockReplenishmentInfo to *Many2One.
func (sri *StockReplenishmentInfo) Many2One() *Many2One {
	return NewMany2One(sri.Id.Get(), "")
}

// CreateStockReplenishmentInfo creates a new stock.replenishment.info model and returns its id.
func (c *Client) CreateStockReplenishmentInfo(sri *StockReplenishmentInfo) (int64, error) {
	return c.Create(StockReplenishmentInfoModel, sri)
}

// UpdateStockReplenishmentInfo updates an existing stock.replenishment.info record.
func (c *Client) UpdateStockReplenishmentInfo(sri *StockReplenishmentInfo) error {
	return c.UpdateStockReplenishmentInfos([]int64{sri.Id.Get()}, sri)
}

// UpdateStockReplenishmentInfos updates existing stock.replenishment.info records.
// All records (represented by ids) will be updated by sri values.
func (c *Client) UpdateStockReplenishmentInfos(ids []int64, sri *StockReplenishmentInfo) error {
	return c.Update(StockReplenishmentInfoModel, ids, sri)
}

// DeleteStockReplenishmentInfo deletes an existing stock.replenishment.info record.
func (c *Client) DeleteStockReplenishmentInfo(id int64) error {
	return c.DeleteStockReplenishmentInfos([]int64{id})
}

// DeleteStockReplenishmentInfos deletes existing stock.replenishment.info records.
func (c *Client) DeleteStockReplenishmentInfos(ids []int64) error {
	return c.Delete(StockReplenishmentInfoModel, ids)
}

// GetStockReplenishmentInfo gets stock.replenishment.info existing record.
func (c *Client) GetStockReplenishmentInfo(id int64) (*StockReplenishmentInfo, error) {
	sris, err := c.GetStockReplenishmentInfos([]int64{id})
	if err != nil {
		return nil, err
	}
	if sris != nil && len(*sris) > 0 {
		return &((*sris)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.replenishment.info not found", id)
}

// GetStockReplenishmentInfos gets stock.replenishment.info existing records.
func (c *Client) GetStockReplenishmentInfos(ids []int64) (*StockReplenishmentInfos, error) {
	sris := &StockReplenishmentInfos{}
	if err := c.Read(StockReplenishmentInfoModel, ids, nil, sris); err != nil {
		return nil, err
	}
	return sris, nil
}

// FindStockReplenishmentInfo finds stock.replenishment.info record by querying it with criteria.
func (c *Client) FindStockReplenishmentInfo(criteria *Criteria) (*StockReplenishmentInfo, error) {
	sris := &StockReplenishmentInfos{}
	if err := c.SearchRead(StockReplenishmentInfoModel, criteria, NewOptions().Limit(1), sris); err != nil {
		return nil, err
	}
	if sris != nil && len(*sris) > 0 {
		return &((*sris)[0]), nil
	}
	return nil, fmt.Errorf("stock.replenishment.info was not found with criteria %v", criteria)
}

// FindStockReplenishmentInfos finds stock.replenishment.info records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReplenishmentInfos(criteria *Criteria, options *Options) (*StockReplenishmentInfos, error) {
	sris := &StockReplenishmentInfos{}
	if err := c.SearchRead(StockReplenishmentInfoModel, criteria, options, sris); err != nil {
		return nil, err
	}
	return sris, nil
}

// FindStockReplenishmentInfoIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReplenishmentInfoIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockReplenishmentInfoModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockReplenishmentInfoId finds record id by querying it with criteria.
func (c *Client) FindStockReplenishmentInfoId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockReplenishmentInfoModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.replenishment.info was not found with criteria %v and options %v", criteria, options)
}
