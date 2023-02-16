package odoo

import (
	"fmt"
)

// StockRoute represents stock.route model.
type StockRoute struct {
	LastUpdate             *Time     `xmlrpc:"__last_update,omptempty"`
	Active                 *Bool     `xmlrpc:"active,omptempty"`
	CategIds               *Relation `xmlrpc:"categ_ids,omptempty"`
	CompanyId              *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName            *String   `xmlrpc:"display_name,omptempty"`
	Id                     *Int      `xmlrpc:"id,omptempty"`
	Name                   *String   `xmlrpc:"name,omptempty"`
	PackagingIds           *Relation `xmlrpc:"packaging_ids,omptempty"`
	PackagingSelectable    *Bool     `xmlrpc:"packaging_selectable,omptempty"`
	ProductCategSelectable *Bool     `xmlrpc:"product_categ_selectable,omptempty"`
	ProductIds             *Relation `xmlrpc:"product_ids,omptempty"`
	ProductSelectable      *Bool     `xmlrpc:"product_selectable,omptempty"`
	RuleIds                *Relation `xmlrpc:"rule_ids,omptempty"`
	Sequence               *Int      `xmlrpc:"sequence,omptempty"`
	SuppliedWhId           *Many2One `xmlrpc:"supplied_wh_id,omptempty"`
	SupplierWhId           *Many2One `xmlrpc:"supplier_wh_id,omptempty"`
	WarehouseDomainIds     *Relation `xmlrpc:"warehouse_domain_ids,omptempty"`
	WarehouseIds           *Relation `xmlrpc:"warehouse_ids,omptempty"`
	WarehouseSelectable    *Bool     `xmlrpc:"warehouse_selectable,omptempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockRoutes represents array of stock.route model.
type StockRoutes []StockRoute

// StockRouteModel is the odoo model name.
const StockRouteModel = "stock.route"

// Many2One convert StockRoute to *Many2One.
func (sr *StockRoute) Many2One() *Many2One {
	return NewMany2One(sr.Id.Get(), "")
}

// CreateStockRoute creates a new stock.route model and returns its id.
func (c *Client) CreateStockRoute(sr *StockRoute) (int64, error) {
	return c.Create(StockRouteModel, sr)
}

// UpdateStockRoute updates an existing stock.route record.
func (c *Client) UpdateStockRoute(sr *StockRoute) error {
	return c.UpdateStockRoutes([]int64{sr.Id.Get()}, sr)
}

// UpdateStockRoutes updates existing stock.route records.
// All records (represented by ids) will be updated by sr values.
func (c *Client) UpdateStockRoutes(ids []int64, sr *StockRoute) error {
	return c.Update(StockRouteModel, ids, sr)
}

// DeleteStockRoute deletes an existing stock.route record.
func (c *Client) DeleteStockRoute(id int64) error {
	return c.DeleteStockRoutes([]int64{id})
}

// DeleteStockRoutes deletes existing stock.route records.
func (c *Client) DeleteStockRoutes(ids []int64) error {
	return c.Delete(StockRouteModel, ids)
}

// GetStockRoute gets stock.route existing record.
func (c *Client) GetStockRoute(id int64) (*StockRoute, error) {
	srs, err := c.GetStockRoutes([]int64{id})
	if err != nil {
		return nil, err
	}
	if srs != nil && len(*srs) > 0 {
		return &((*srs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.route not found", id)
}

// GetStockRoutes gets stock.route existing records.
func (c *Client) GetStockRoutes(ids []int64) (*StockRoutes, error) {
	srs := &StockRoutes{}
	if err := c.Read(StockRouteModel, ids, nil, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindStockRoute finds stock.route record by querying it with criteria.
func (c *Client) FindStockRoute(criteria *Criteria) (*StockRoute, error) {
	srs := &StockRoutes{}
	if err := c.SearchRead(StockRouteModel, criteria, NewOptions().Limit(1), srs); err != nil {
		return nil, err
	}
	if srs != nil && len(*srs) > 0 {
		return &((*srs)[0]), nil
	}
	return nil, fmt.Errorf("stock.route was not found with criteria %v", criteria)
}

// FindStockRoutes finds stock.route records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockRoutes(criteria *Criteria, options *Options) (*StockRoutes, error) {
	srs := &StockRoutes{}
	if err := c.SearchRead(StockRouteModel, criteria, options, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindStockRouteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockRouteIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockRouteModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockRouteId finds record id by querying it with criteria.
func (c *Client) FindStockRouteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockRouteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.route was not found with criteria %v and options %v", criteria, options)
}
