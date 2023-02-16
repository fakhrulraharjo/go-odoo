package odoo

import (
	"fmt"
)

// StockValuationLayer represents stock.valuation.layer model.
type StockValuationLayer struct {
	LastUpdate             *Time     `xmlrpc:"__last_update,omptempty"`
	AccountMoveId          *Many2One `xmlrpc:"account_move_id,omptempty"`
	AccountMoveLineId      *Many2One `xmlrpc:"account_move_line_id,omptempty"`
	CategId                *Many2One `xmlrpc:"categ_id,omptempty"`
	CompanyId              *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId             *Many2One `xmlrpc:"currency_id,omptempty"`
	Description            *String   `xmlrpc:"description,omptempty"`
	DisplayName            *String   `xmlrpc:"display_name,omptempty"`
	Id                     *Int      `xmlrpc:"id,omptempty"`
	PriceDiffValue         *Float    `xmlrpc:"price_diff_value,omptempty"`
	ProductId              *Many2One `xmlrpc:"product_id,omptempty"`
	ProductTmplId          *Many2One `xmlrpc:"product_tmpl_id,omptempty"`
	Quantity               *Float    `xmlrpc:"quantity,omptempty"`
	Reference              *String   `xmlrpc:"reference,omptempty"`
	RemainingQty           *Float    `xmlrpc:"remaining_qty,omptempty"`
	RemainingValue         *Float    `xmlrpc:"remaining_value,omptempty"`
	StockMoveId            *Many2One `xmlrpc:"stock_move_id,omptempty"`
	StockValuationLayerId  *Many2One `xmlrpc:"stock_valuation_layer_id,omptempty"`
	StockValuationLayerIds *Relation `xmlrpc:"stock_valuation_layer_ids,omptempty"`
	UnitCost               *Float    `xmlrpc:"unit_cost,omptempty"`
	UomId                  *Many2One `xmlrpc:"uom_id,omptempty"`
	Value                  *Float    `xmlrpc:"value,omptempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockValuationLayers represents array of stock.valuation.layer model.
type StockValuationLayers []StockValuationLayer

// StockValuationLayerModel is the odoo model name.
const StockValuationLayerModel = "stock.valuation.layer"

// Many2One convert StockValuationLayer to *Many2One.
func (svl *StockValuationLayer) Many2One() *Many2One {
	return NewMany2One(svl.Id.Get(), "")
}

// CreateStockValuationLayer creates a new stock.valuation.layer model and returns its id.
func (c *Client) CreateStockValuationLayer(svl *StockValuationLayer) (int64, error) {
	return c.Create(StockValuationLayerModel, svl)
}

// UpdateStockValuationLayer updates an existing stock.valuation.layer record.
func (c *Client) UpdateStockValuationLayer(svl *StockValuationLayer) error {
	return c.UpdateStockValuationLayers([]int64{svl.Id.Get()}, svl)
}

// UpdateStockValuationLayers updates existing stock.valuation.layer records.
// All records (represented by ids) will be updated by svl values.
func (c *Client) UpdateStockValuationLayers(ids []int64, svl *StockValuationLayer) error {
	return c.Update(StockValuationLayerModel, ids, svl)
}

// DeleteStockValuationLayer deletes an existing stock.valuation.layer record.
func (c *Client) DeleteStockValuationLayer(id int64) error {
	return c.DeleteStockValuationLayers([]int64{id})
}

// DeleteStockValuationLayers deletes existing stock.valuation.layer records.
func (c *Client) DeleteStockValuationLayers(ids []int64) error {
	return c.Delete(StockValuationLayerModel, ids)
}

// GetStockValuationLayer gets stock.valuation.layer existing record.
func (c *Client) GetStockValuationLayer(id int64) (*StockValuationLayer, error) {
	svls, err := c.GetStockValuationLayers([]int64{id})
	if err != nil {
		return nil, err
	}
	if svls != nil && len(*svls) > 0 {
		return &((*svls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.valuation.layer not found", id)
}

// GetStockValuationLayers gets stock.valuation.layer existing records.
func (c *Client) GetStockValuationLayers(ids []int64) (*StockValuationLayers, error) {
	svls := &StockValuationLayers{}
	if err := c.Read(StockValuationLayerModel, ids, nil, svls); err != nil {
		return nil, err
	}
	return svls, nil
}

// FindStockValuationLayer finds stock.valuation.layer record by querying it with criteria.
func (c *Client) FindStockValuationLayer(criteria *Criteria) (*StockValuationLayer, error) {
	svls := &StockValuationLayers{}
	if err := c.SearchRead(StockValuationLayerModel, criteria, NewOptions().Limit(1), svls); err != nil {
		return nil, err
	}
	if svls != nil && len(*svls) > 0 {
		return &((*svls)[0]), nil
	}
	return nil, fmt.Errorf("stock.valuation.layer was not found with criteria %v", criteria)
}

// FindStockValuationLayers finds stock.valuation.layer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockValuationLayers(criteria *Criteria, options *Options) (*StockValuationLayers, error) {
	svls := &StockValuationLayers{}
	if err := c.SearchRead(StockValuationLayerModel, criteria, options, svls); err != nil {
		return nil, err
	}
	return svls, nil
}

// FindStockValuationLayerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockValuationLayerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockValuationLayerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockValuationLayerId finds record id by querying it with criteria.
func (c *Client) FindStockValuationLayerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockValuationLayerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.valuation.layer was not found with criteria %v and options %v", criteria, options)
}
