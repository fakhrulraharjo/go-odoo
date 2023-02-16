package odoo

import (
	"fmt"
)

// StockPackageType represents stock.package.type model.
type StockPackageType struct {
	LastUpdate                 *Time     `xmlrpc:"__last_update,omptempty"`
	Barcode                    *String   `xmlrpc:"barcode,omptempty"`
	BaseWeight                 *Float    `xmlrpc:"base_weight,omptempty"`
	CompanyId                  *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate                 *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                  *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName                *String   `xmlrpc:"display_name,omptempty"`
	Height                     *Int      `xmlrpc:"height,omptempty"`
	Id                         *Int      `xmlrpc:"id,omptempty"`
	LengthUomName              *String   `xmlrpc:"length_uom_name,omptempty"`
	MaxWeight                  *Float    `xmlrpc:"max_weight,omptempty"`
	Name                       *String   `xmlrpc:"name,omptempty"`
	PackagingLength            *Int      `xmlrpc:"packaging_length,omptempty"`
	Sequence                   *Int      `xmlrpc:"sequence,omptempty"`
	StorageCategoryCapacityIds *Relation `xmlrpc:"storage_category_capacity_ids,omptempty"`
	WeightUomName              *String   `xmlrpc:"weight_uom_name,omptempty"`
	Width                      *Int      `xmlrpc:"width,omptempty"`
	WriteDate                  *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                   *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockPackageTypes represents array of stock.package.type model.
type StockPackageTypes []StockPackageType

// StockPackageTypeModel is the odoo model name.
const StockPackageTypeModel = "stock.package.type"

// Many2One convert StockPackageType to *Many2One.
func (spt *StockPackageType) Many2One() *Many2One {
	return NewMany2One(spt.Id.Get(), "")
}

// CreateStockPackageType creates a new stock.package.type model and returns its id.
func (c *Client) CreateStockPackageType(spt *StockPackageType) (int64, error) {
	return c.Create(StockPackageTypeModel, spt)
}

// UpdateStockPackageType updates an existing stock.package.type record.
func (c *Client) UpdateStockPackageType(spt *StockPackageType) error {
	return c.UpdateStockPackageTypes([]int64{spt.Id.Get()}, spt)
}

// UpdateStockPackageTypes updates existing stock.package.type records.
// All records (represented by ids) will be updated by spt values.
func (c *Client) UpdateStockPackageTypes(ids []int64, spt *StockPackageType) error {
	return c.Update(StockPackageTypeModel, ids, spt)
}

// DeleteStockPackageType deletes an existing stock.package.type record.
func (c *Client) DeleteStockPackageType(id int64) error {
	return c.DeleteStockPackageTypes([]int64{id})
}

// DeleteStockPackageTypes deletes existing stock.package.type records.
func (c *Client) DeleteStockPackageTypes(ids []int64) error {
	return c.Delete(StockPackageTypeModel, ids)
}

// GetStockPackageType gets stock.package.type existing record.
func (c *Client) GetStockPackageType(id int64) (*StockPackageType, error) {
	spts, err := c.GetStockPackageTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if spts != nil && len(*spts) > 0 {
		return &((*spts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.package.type not found", id)
}

// GetStockPackageTypes gets stock.package.type existing records.
func (c *Client) GetStockPackageTypes(ids []int64) (*StockPackageTypes, error) {
	spts := &StockPackageTypes{}
	if err := c.Read(StockPackageTypeModel, ids, nil, spts); err != nil {
		return nil, err
	}
	return spts, nil
}

// FindStockPackageType finds stock.package.type record by querying it with criteria.
func (c *Client) FindStockPackageType(criteria *Criteria) (*StockPackageType, error) {
	spts := &StockPackageTypes{}
	if err := c.SearchRead(StockPackageTypeModel, criteria, NewOptions().Limit(1), spts); err != nil {
		return nil, err
	}
	if spts != nil && len(*spts) > 0 {
		return &((*spts)[0]), nil
	}
	return nil, fmt.Errorf("stock.package.type was not found with criteria %v", criteria)
}

// FindStockPackageTypes finds stock.package.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPackageTypes(criteria *Criteria, options *Options) (*StockPackageTypes, error) {
	spts := &StockPackageTypes{}
	if err := c.SearchRead(StockPackageTypeModel, criteria, options, spts); err != nil {
		return nil, err
	}
	return spts, nil
}

// FindStockPackageTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPackageTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockPackageTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockPackageTypeId finds record id by querying it with criteria.
func (c *Client) FindStockPackageTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockPackageTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.package.type was not found with criteria %v and options %v", criteria, options)
}
