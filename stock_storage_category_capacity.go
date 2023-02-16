package odoo

import (
	"fmt"
)

// StockStorageCategoryCapacity represents stock.storage.category.capacity model.
type StockStorageCategoryCapacity struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	CompanyId         *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	PackageTypeId     *Many2One `xmlrpc:"package_type_id,omptempty"`
	ProductId         *Many2One `xmlrpc:"product_id,omptempty"`
	ProductUomId      *Many2One `xmlrpc:"product_uom_id,omptempty"`
	Quantity          *Float    `xmlrpc:"quantity,omptempty"`
	StorageCategoryId *Many2One `xmlrpc:"storage_category_id,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockStorageCategoryCapacitys represents array of stock.storage.category.capacity model.
type StockStorageCategoryCapacitys []StockStorageCategoryCapacity

// StockStorageCategoryCapacityModel is the odoo model name.
const StockStorageCategoryCapacityModel = "stock.storage.category.capacity"

// Many2One convert StockStorageCategoryCapacity to *Many2One.
func (sscc *StockStorageCategoryCapacity) Many2One() *Many2One {
	return NewMany2One(sscc.Id.Get(), "")
}

// CreateStockStorageCategoryCapacity creates a new stock.storage.category.capacity model and returns its id.
func (c *Client) CreateStockStorageCategoryCapacity(sscc *StockStorageCategoryCapacity) (int64, error) {
	return c.Create(StockStorageCategoryCapacityModel, sscc)
}

// UpdateStockStorageCategoryCapacity updates an existing stock.storage.category.capacity record.
func (c *Client) UpdateStockStorageCategoryCapacity(sscc *StockStorageCategoryCapacity) error {
	return c.UpdateStockStorageCategoryCapacitys([]int64{sscc.Id.Get()}, sscc)
}

// UpdateStockStorageCategoryCapacitys updates existing stock.storage.category.capacity records.
// All records (represented by ids) will be updated by sscc values.
func (c *Client) UpdateStockStorageCategoryCapacitys(ids []int64, sscc *StockStorageCategoryCapacity) error {
	return c.Update(StockStorageCategoryCapacityModel, ids, sscc)
}

// DeleteStockStorageCategoryCapacity deletes an existing stock.storage.category.capacity record.
func (c *Client) DeleteStockStorageCategoryCapacity(id int64) error {
	return c.DeleteStockStorageCategoryCapacitys([]int64{id})
}

// DeleteStockStorageCategoryCapacitys deletes existing stock.storage.category.capacity records.
func (c *Client) DeleteStockStorageCategoryCapacitys(ids []int64) error {
	return c.Delete(StockStorageCategoryCapacityModel, ids)
}

// GetStockStorageCategoryCapacity gets stock.storage.category.capacity existing record.
func (c *Client) GetStockStorageCategoryCapacity(id int64) (*StockStorageCategoryCapacity, error) {
	ssccs, err := c.GetStockStorageCategoryCapacitys([]int64{id})
	if err != nil {
		return nil, err
	}
	if ssccs != nil && len(*ssccs) > 0 {
		return &((*ssccs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.storage.category.capacity not found", id)
}

// GetStockStorageCategoryCapacitys gets stock.storage.category.capacity existing records.
func (c *Client) GetStockStorageCategoryCapacitys(ids []int64) (*StockStorageCategoryCapacitys, error) {
	ssccs := &StockStorageCategoryCapacitys{}
	if err := c.Read(StockStorageCategoryCapacityModel, ids, nil, ssccs); err != nil {
		return nil, err
	}
	return ssccs, nil
}

// FindStockStorageCategoryCapacity finds stock.storage.category.capacity record by querying it with criteria.
func (c *Client) FindStockStorageCategoryCapacity(criteria *Criteria) (*StockStorageCategoryCapacity, error) {
	ssccs := &StockStorageCategoryCapacitys{}
	if err := c.SearchRead(StockStorageCategoryCapacityModel, criteria, NewOptions().Limit(1), ssccs); err != nil {
		return nil, err
	}
	if ssccs != nil && len(*ssccs) > 0 {
		return &((*ssccs)[0]), nil
	}
	return nil, fmt.Errorf("stock.storage.category.capacity was not found with criteria %v", criteria)
}

// FindStockStorageCategoryCapacitys finds stock.storage.category.capacity records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockStorageCategoryCapacitys(criteria *Criteria, options *Options) (*StockStorageCategoryCapacitys, error) {
	ssccs := &StockStorageCategoryCapacitys{}
	if err := c.SearchRead(StockStorageCategoryCapacityModel, criteria, options, ssccs); err != nil {
		return nil, err
	}
	return ssccs, nil
}

// FindStockStorageCategoryCapacityIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockStorageCategoryCapacityIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockStorageCategoryCapacityModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockStorageCategoryCapacityId finds record id by querying it with criteria.
func (c *Client) FindStockStorageCategoryCapacityId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockStorageCategoryCapacityModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.storage.category.capacity was not found with criteria %v and options %v", criteria, options)
}
