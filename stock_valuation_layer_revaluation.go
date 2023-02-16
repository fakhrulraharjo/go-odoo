package odoo

import (
	"fmt"
)

// StockValuationLayerRevaluation represents stock.valuation.layer.revaluation model.
type StockValuationLayerRevaluation struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omptempty"`
	AccountId          *Many2One  `xmlrpc:"account_id,omptempty"`
	AccountJournalId   *Many2One  `xmlrpc:"account_journal_id,omptempty"`
	AddedValue         *Float     `xmlrpc:"added_value,omptempty"`
	CompanyId          *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId         *Many2One  `xmlrpc:"currency_id,omptempty"`
	CurrentQuantitySvl *Float     `xmlrpc:"current_quantity_svl,omptempty"`
	CurrentValueSvl    *Float     `xmlrpc:"current_value_svl,omptempty"`
	Date               *Time      `xmlrpc:"date,omptempty"`
	DisplayName        *String    `xmlrpc:"display_name,omptempty"`
	Id                 *Int       `xmlrpc:"id,omptempty"`
	NewValue           *Float     `xmlrpc:"new_value,omptempty"`
	NewValueByQty      *Float     `xmlrpc:"new_value_by_qty,omptempty"`
	ProductId          *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductUomName     *String    `xmlrpc:"product_uom_name,omptempty"`
	PropertyValuation  *Selection `xmlrpc:"property_valuation,omptempty"`
	Reason             *String    `xmlrpc:"reason,omptempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// StockValuationLayerRevaluations represents array of stock.valuation.layer.revaluation model.
type StockValuationLayerRevaluations []StockValuationLayerRevaluation

// StockValuationLayerRevaluationModel is the odoo model name.
const StockValuationLayerRevaluationModel = "stock.valuation.layer.revaluation"

// Many2One convert StockValuationLayerRevaluation to *Many2One.
func (svlr *StockValuationLayerRevaluation) Many2One() *Many2One {
	return NewMany2One(svlr.Id.Get(), "")
}

// CreateStockValuationLayerRevaluation creates a new stock.valuation.layer.revaluation model and returns its id.
func (c *Client) CreateStockValuationLayerRevaluation(svlr *StockValuationLayerRevaluation) (int64, error) {
	return c.Create(StockValuationLayerRevaluationModel, svlr)
}

// UpdateStockValuationLayerRevaluation updates an existing stock.valuation.layer.revaluation record.
func (c *Client) UpdateStockValuationLayerRevaluation(svlr *StockValuationLayerRevaluation) error {
	return c.UpdateStockValuationLayerRevaluations([]int64{svlr.Id.Get()}, svlr)
}

// UpdateStockValuationLayerRevaluations updates existing stock.valuation.layer.revaluation records.
// All records (represented by ids) will be updated by svlr values.
func (c *Client) UpdateStockValuationLayerRevaluations(ids []int64, svlr *StockValuationLayerRevaluation) error {
	return c.Update(StockValuationLayerRevaluationModel, ids, svlr)
}

// DeleteStockValuationLayerRevaluation deletes an existing stock.valuation.layer.revaluation record.
func (c *Client) DeleteStockValuationLayerRevaluation(id int64) error {
	return c.DeleteStockValuationLayerRevaluations([]int64{id})
}

// DeleteStockValuationLayerRevaluations deletes existing stock.valuation.layer.revaluation records.
func (c *Client) DeleteStockValuationLayerRevaluations(ids []int64) error {
	return c.Delete(StockValuationLayerRevaluationModel, ids)
}

// GetStockValuationLayerRevaluation gets stock.valuation.layer.revaluation existing record.
func (c *Client) GetStockValuationLayerRevaluation(id int64) (*StockValuationLayerRevaluation, error) {
	svlrs, err := c.GetStockValuationLayerRevaluations([]int64{id})
	if err != nil {
		return nil, err
	}
	if svlrs != nil && len(*svlrs) > 0 {
		return &((*svlrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.valuation.layer.revaluation not found", id)
}

// GetStockValuationLayerRevaluations gets stock.valuation.layer.revaluation existing records.
func (c *Client) GetStockValuationLayerRevaluations(ids []int64) (*StockValuationLayerRevaluations, error) {
	svlrs := &StockValuationLayerRevaluations{}
	if err := c.Read(StockValuationLayerRevaluationModel, ids, nil, svlrs); err != nil {
		return nil, err
	}
	return svlrs, nil
}

// FindStockValuationLayerRevaluation finds stock.valuation.layer.revaluation record by querying it with criteria.
func (c *Client) FindStockValuationLayerRevaluation(criteria *Criteria) (*StockValuationLayerRevaluation, error) {
	svlrs := &StockValuationLayerRevaluations{}
	if err := c.SearchRead(StockValuationLayerRevaluationModel, criteria, NewOptions().Limit(1), svlrs); err != nil {
		return nil, err
	}
	if svlrs != nil && len(*svlrs) > 0 {
		return &((*svlrs)[0]), nil
	}
	return nil, fmt.Errorf("stock.valuation.layer.revaluation was not found with criteria %v", criteria)
}

// FindStockValuationLayerRevaluations finds stock.valuation.layer.revaluation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockValuationLayerRevaluations(criteria *Criteria, options *Options) (*StockValuationLayerRevaluations, error) {
	svlrs := &StockValuationLayerRevaluations{}
	if err := c.SearchRead(StockValuationLayerRevaluationModel, criteria, options, svlrs); err != nil {
		return nil, err
	}
	return svlrs, nil
}

// FindStockValuationLayerRevaluationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockValuationLayerRevaluationIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockValuationLayerRevaluationModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockValuationLayerRevaluationId finds record id by querying it with criteria.
func (c *Client) FindStockValuationLayerRevaluationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockValuationLayerRevaluationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.valuation.layer.revaluation was not found with criteria %v and options %v", criteria, options)
}
