package odoo

import (
	"fmt"
)

// ReportStockQuantity represents report.stock.quantity model.
type ReportStockQuantity struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omptempty"`
	CompanyId     *Many2One  `xmlrpc:"company_id,omptempty"`
	Date          *Time      `xmlrpc:"date,omptempty"`
	DisplayName   *String    `xmlrpc:"display_name,omptempty"`
	Id            *Int       `xmlrpc:"id,omptempty"`
	MoveIds       *Relation  `xmlrpc:"move_ids,omptempty"`
	ProductId     *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductQty    *Float     `xmlrpc:"product_qty,omptempty"`
	ProductTmplId *Many2One  `xmlrpc:"product_tmpl_id,omptempty"`
	State         *Selection `xmlrpc:"state,omptempty"`
	WarehouseId   *Many2One  `xmlrpc:"warehouse_id,omptempty"`
}

// ReportStockQuantitys represents array of report.stock.quantity model.
type ReportStockQuantitys []ReportStockQuantity

// ReportStockQuantityModel is the odoo model name.
const ReportStockQuantityModel = "report.stock.quantity"

// Many2One convert ReportStockQuantity to *Many2One.
func (rsq *ReportStockQuantity) Many2One() *Many2One {
	return NewMany2One(rsq.Id.Get(), "")
}

// CreateReportStockQuantity creates a new report.stock.quantity model and returns its id.
func (c *Client) CreateReportStockQuantity(rsq *ReportStockQuantity) (int64, error) {
	return c.Create(ReportStockQuantityModel, rsq)
}

// UpdateReportStockQuantity updates an existing report.stock.quantity record.
func (c *Client) UpdateReportStockQuantity(rsq *ReportStockQuantity) error {
	return c.UpdateReportStockQuantitys([]int64{rsq.Id.Get()}, rsq)
}

// UpdateReportStockQuantitys updates existing report.stock.quantity records.
// All records (represented by ids) will be updated by rsq values.
func (c *Client) UpdateReportStockQuantitys(ids []int64, rsq *ReportStockQuantity) error {
	return c.Update(ReportStockQuantityModel, ids, rsq)
}

// DeleteReportStockQuantity deletes an existing report.stock.quantity record.
func (c *Client) DeleteReportStockQuantity(id int64) error {
	return c.DeleteReportStockQuantitys([]int64{id})
}

// DeleteReportStockQuantitys deletes existing report.stock.quantity records.
func (c *Client) DeleteReportStockQuantitys(ids []int64) error {
	return c.Delete(ReportStockQuantityModel, ids)
}

// GetReportStockQuantity gets report.stock.quantity existing record.
func (c *Client) GetReportStockQuantity(id int64) (*ReportStockQuantity, error) {
	rsqs, err := c.GetReportStockQuantitys([]int64{id})
	if err != nil {
		return nil, err
	}
	if rsqs != nil && len(*rsqs) > 0 {
		return &((*rsqs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of report.stock.quantity not found", id)
}

// GetReportStockQuantitys gets report.stock.quantity existing records.
func (c *Client) GetReportStockQuantitys(ids []int64) (*ReportStockQuantitys, error) {
	rsqs := &ReportStockQuantitys{}
	if err := c.Read(ReportStockQuantityModel, ids, nil, rsqs); err != nil {
		return nil, err
	}
	return rsqs, nil
}

// FindReportStockQuantity finds report.stock.quantity record by querying it with criteria.
func (c *Client) FindReportStockQuantity(criteria *Criteria) (*ReportStockQuantity, error) {
	rsqs := &ReportStockQuantitys{}
	if err := c.SearchRead(ReportStockQuantityModel, criteria, NewOptions().Limit(1), rsqs); err != nil {
		return nil, err
	}
	if rsqs != nil && len(*rsqs) > 0 {
		return &((*rsqs)[0]), nil
	}
	return nil, fmt.Errorf("report.stock.quantity was not found with criteria %v", criteria)
}

// FindReportStockQuantitys finds report.stock.quantity records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportStockQuantitys(criteria *Criteria, options *Options) (*ReportStockQuantitys, error) {
	rsqs := &ReportStockQuantitys{}
	if err := c.SearchRead(ReportStockQuantityModel, criteria, options, rsqs); err != nil {
		return nil, err
	}
	return rsqs, nil
}

// FindReportStockQuantityIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportStockQuantityIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ReportStockQuantityModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindReportStockQuantityId finds record id by querying it with criteria.
func (c *Client) FindReportStockQuantityId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportStockQuantityModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("report.stock.quantity was not found with criteria %v and options %v", criteria, options)
}
