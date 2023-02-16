package odoo

import (
	"fmt"
)

// StockPackageLevel represents stock.package_level model.
type StockPackageLevel struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	IsDone          *Bool      `xmlrpc:"is_done,omptempty"`
	IsFreshPackage  *Bool      `xmlrpc:"is_fresh_package,omptempty"`
	LocationDestId  *Many2One  `xmlrpc:"location_dest_id,omptempty"`
	LocationId      *Many2One  `xmlrpc:"location_id,omptempty"`
	MoveIds         *Relation  `xmlrpc:"move_ids,omptempty"`
	MoveLineIds     *Relation  `xmlrpc:"move_line_ids,omptempty"`
	PackageId       *Many2One  `xmlrpc:"package_id,omptempty"`
	PickingId       *Many2One  `xmlrpc:"picking_id,omptempty"`
	PickingTypeCode *Selection `xmlrpc:"picking_type_code,omptempty"`
	ShowLotsM2O     *Bool      `xmlrpc:"show_lots_m2o,omptempty"`
	ShowLotsText    *Bool      `xmlrpc:"show_lots_text,omptempty"`
	State           *Selection `xmlrpc:"state,omptempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// StockPackageLevels represents array of stock.package_level model.
type StockPackageLevels []StockPackageLevel

// StockPackageLevelModel is the odoo model name.
const StockPackageLevelModel = "stock.package_level"

// Many2One convert StockPackageLevel to *Many2One.
func (sp *StockPackageLevel) Many2One() *Many2One {
	return NewMany2One(sp.Id.Get(), "")
}

// CreateStockPackageLevel creates a new stock.package_level model and returns its id.
func (c *Client) CreateStockPackageLevel(sp *StockPackageLevel) (int64, error) {
	return c.Create(StockPackageLevelModel, sp)
}

// UpdateStockPackageLevel updates an existing stock.package_level record.
func (c *Client) UpdateStockPackageLevel(sp *StockPackageLevel) error {
	return c.UpdateStockPackageLevels([]int64{sp.Id.Get()}, sp)
}

// UpdateStockPackageLevels updates existing stock.package_level records.
// All records (represented by ids) will be updated by sp values.
func (c *Client) UpdateStockPackageLevels(ids []int64, sp *StockPackageLevel) error {
	return c.Update(StockPackageLevelModel, ids, sp)
}

// DeleteStockPackageLevel deletes an existing stock.package_level record.
func (c *Client) DeleteStockPackageLevel(id int64) error {
	return c.DeleteStockPackageLevels([]int64{id})
}

// DeleteStockPackageLevels deletes existing stock.package_level records.
func (c *Client) DeleteStockPackageLevels(ids []int64) error {
	return c.Delete(StockPackageLevelModel, ids)
}

// GetStockPackageLevel gets stock.package_level existing record.
func (c *Client) GetStockPackageLevel(id int64) (*StockPackageLevel, error) {
	sps, err := c.GetStockPackageLevels([]int64{id})
	if err != nil {
		return nil, err
	}
	if sps != nil && len(*sps) > 0 {
		return &((*sps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.package_level not found", id)
}

// GetStockPackageLevels gets stock.package_level existing records.
func (c *Client) GetStockPackageLevels(ids []int64) (*StockPackageLevels, error) {
	sps := &StockPackageLevels{}
	if err := c.Read(StockPackageLevelModel, ids, nil, sps); err != nil {
		return nil, err
	}
	return sps, nil
}

// FindStockPackageLevel finds stock.package_level record by querying it with criteria.
func (c *Client) FindStockPackageLevel(criteria *Criteria) (*StockPackageLevel, error) {
	sps := &StockPackageLevels{}
	if err := c.SearchRead(StockPackageLevelModel, criteria, NewOptions().Limit(1), sps); err != nil {
		return nil, err
	}
	if sps != nil && len(*sps) > 0 {
		return &((*sps)[0]), nil
	}
	return nil, fmt.Errorf("stock.package_level was not found with criteria %v", criteria)
}

// FindStockPackageLevels finds stock.package_level records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPackageLevels(criteria *Criteria, options *Options) (*StockPackageLevels, error) {
	sps := &StockPackageLevels{}
	if err := c.SearchRead(StockPackageLevelModel, criteria, options, sps); err != nil {
		return nil, err
	}
	return sps, nil
}

// FindStockPackageLevelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPackageLevelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockPackageLevelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockPackageLevelId finds record id by querying it with criteria.
func (c *Client) FindStockPackageLevelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockPackageLevelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.package_level was not found with criteria %v and options %v", criteria, options)
}
