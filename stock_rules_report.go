package odoo

import (
	"fmt"
)

// StockRulesReport represents stock.rules.report model.
type StockRulesReport struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName        *String   `xmlrpc:"display_name,omptempty"`
	Id                 *Int      `xmlrpc:"id,omptempty"`
	ProductHasVariants *Bool     `xmlrpc:"product_has_variants,omptempty"`
	ProductId          *Many2One `xmlrpc:"product_id,omptempty"`
	ProductTmplId      *Many2One `xmlrpc:"product_tmpl_id,omptempty"`
	WarehouseIds       *Relation `xmlrpc:"warehouse_ids,omptempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockRulesReports represents array of stock.rules.report model.
type StockRulesReports []StockRulesReport

// StockRulesReportModel is the odoo model name.
const StockRulesReportModel = "stock.rules.report"

// Many2One convert StockRulesReport to *Many2One.
func (srr *StockRulesReport) Many2One() *Many2One {
	return NewMany2One(srr.Id.Get(), "")
}

// CreateStockRulesReport creates a new stock.rules.report model and returns its id.
func (c *Client) CreateStockRulesReport(srr *StockRulesReport) (int64, error) {
	return c.Create(StockRulesReportModel, srr)
}

// UpdateStockRulesReport updates an existing stock.rules.report record.
func (c *Client) UpdateStockRulesReport(srr *StockRulesReport) error {
	return c.UpdateStockRulesReports([]int64{srr.Id.Get()}, srr)
}

// UpdateStockRulesReports updates existing stock.rules.report records.
// All records (represented by ids) will be updated by srr values.
func (c *Client) UpdateStockRulesReports(ids []int64, srr *StockRulesReport) error {
	return c.Update(StockRulesReportModel, ids, srr)
}

// DeleteStockRulesReport deletes an existing stock.rules.report record.
func (c *Client) DeleteStockRulesReport(id int64) error {
	return c.DeleteStockRulesReports([]int64{id})
}

// DeleteStockRulesReports deletes existing stock.rules.report records.
func (c *Client) DeleteStockRulesReports(ids []int64) error {
	return c.Delete(StockRulesReportModel, ids)
}

// GetStockRulesReport gets stock.rules.report existing record.
func (c *Client) GetStockRulesReport(id int64) (*StockRulesReport, error) {
	srrs, err := c.GetStockRulesReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if srrs != nil && len(*srrs) > 0 {
		return &((*srrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.rules.report not found", id)
}

// GetStockRulesReports gets stock.rules.report existing records.
func (c *Client) GetStockRulesReports(ids []int64) (*StockRulesReports, error) {
	srrs := &StockRulesReports{}
	if err := c.Read(StockRulesReportModel, ids, nil, srrs); err != nil {
		return nil, err
	}
	return srrs, nil
}

// FindStockRulesReport finds stock.rules.report record by querying it with criteria.
func (c *Client) FindStockRulesReport(criteria *Criteria) (*StockRulesReport, error) {
	srrs := &StockRulesReports{}
	if err := c.SearchRead(StockRulesReportModel, criteria, NewOptions().Limit(1), srrs); err != nil {
		return nil, err
	}
	if srrs != nil && len(*srrs) > 0 {
		return &((*srrs)[0]), nil
	}
	return nil, fmt.Errorf("stock.rules.report was not found with criteria %v", criteria)
}

// FindStockRulesReports finds stock.rules.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockRulesReports(criteria *Criteria, options *Options) (*StockRulesReports, error) {
	srrs := &StockRulesReports{}
	if err := c.SearchRead(StockRulesReportModel, criteria, options, srrs); err != nil {
		return nil, err
	}
	return srrs, nil
}

// FindStockRulesReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockRulesReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockRulesReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockRulesReportId finds record id by querying it with criteria.
func (c *Client) FindStockRulesReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockRulesReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.rules.report was not found with criteria %v and options %v", criteria, options)
}
