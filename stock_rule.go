package odoo

import (
	"fmt"
)

// StockRule represents stock.rule model.
type StockRule struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omptempty"`
	Action                 *Selection `xmlrpc:"action,omptempty"`
	Active                 *Bool      `xmlrpc:"active,omptempty"`
	Auto                   *Selection `xmlrpc:"auto,omptempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omptempty"`
	Delay                  *Int       `xmlrpc:"delay,omptempty"`
	DisplayName            *String    `xmlrpc:"display_name,omptempty"`
	GroupId                *Many2One  `xmlrpc:"group_id,omptempty"`
	GroupPropagationOption *Selection `xmlrpc:"group_propagation_option,omptempty"`
	Id                     *Int       `xmlrpc:"id,omptempty"`
	LocationDestId         *Many2One  `xmlrpc:"location_dest_id,omptempty"`
	LocationSrcId          *Many2One  `xmlrpc:"location_src_id,omptempty"`
	Name                   *String    `xmlrpc:"name,omptempty"`
	PartnerAddressId       *Many2One  `xmlrpc:"partner_address_id,omptempty"`
	PickingTypeCodeDomain  *String    `xmlrpc:"picking_type_code_domain,omptempty"`
	PickingTypeId          *Many2One  `xmlrpc:"picking_type_id,omptempty"`
	ProcureMethod          *Selection `xmlrpc:"procure_method,omptempty"`
	PropagateCancel        *Bool      `xmlrpc:"propagate_cancel,omptempty"`
	PropagateCarrier       *Bool      `xmlrpc:"propagate_carrier,omptempty"`
	PropagateWarehouseId   *Many2One  `xmlrpc:"propagate_warehouse_id,omptempty"`
	RouteCompanyId         *Many2One  `xmlrpc:"route_company_id,omptempty"`
	RouteId                *Many2One  `xmlrpc:"route_id,omptempty"`
	RouteSequence          *Int       `xmlrpc:"route_sequence,omptempty"`
	RuleMessage            *String    `xmlrpc:"rule_message,omptempty"`
	Sequence               *Int       `xmlrpc:"sequence,omptempty"`
	WarehouseId            *Many2One  `xmlrpc:"warehouse_id,omptempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// StockRules represents array of stock.rule model.
type StockRules []StockRule

// StockRuleModel is the odoo model name.
const StockRuleModel = "stock.rule"

// Many2One convert StockRule to *Many2One.
func (sr *StockRule) Many2One() *Many2One {
	return NewMany2One(sr.Id.Get(), "")
}

// CreateStockRule creates a new stock.rule model and returns its id.
func (c *Client) CreateStockRule(sr *StockRule) (int64, error) {
	return c.Create(StockRuleModel, sr)
}

// UpdateStockRule updates an existing stock.rule record.
func (c *Client) UpdateStockRule(sr *StockRule) error {
	return c.UpdateStockRules([]int64{sr.Id.Get()}, sr)
}

// UpdateStockRules updates existing stock.rule records.
// All records (represented by ids) will be updated by sr values.
func (c *Client) UpdateStockRules(ids []int64, sr *StockRule) error {
	return c.Update(StockRuleModel, ids, sr)
}

// DeleteStockRule deletes an existing stock.rule record.
func (c *Client) DeleteStockRule(id int64) error {
	return c.DeleteStockRules([]int64{id})
}

// DeleteStockRules deletes existing stock.rule records.
func (c *Client) DeleteStockRules(ids []int64) error {
	return c.Delete(StockRuleModel, ids)
}

// GetStockRule gets stock.rule existing record.
func (c *Client) GetStockRule(id int64) (*StockRule, error) {
	srs, err := c.GetStockRules([]int64{id})
	if err != nil {
		return nil, err
	}
	if srs != nil && len(*srs) > 0 {
		return &((*srs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.rule not found", id)
}

// GetStockRules gets stock.rule existing records.
func (c *Client) GetStockRules(ids []int64) (*StockRules, error) {
	srs := &StockRules{}
	if err := c.Read(StockRuleModel, ids, nil, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindStockRule finds stock.rule record by querying it with criteria.
func (c *Client) FindStockRule(criteria *Criteria) (*StockRule, error) {
	srs := &StockRules{}
	if err := c.SearchRead(StockRuleModel, criteria, NewOptions().Limit(1), srs); err != nil {
		return nil, err
	}
	if srs != nil && len(*srs) > 0 {
		return &((*srs)[0]), nil
	}
	return nil, fmt.Errorf("stock.rule was not found with criteria %v", criteria)
}

// FindStockRules finds stock.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockRules(criteria *Criteria, options *Options) (*StockRules, error) {
	srs := &StockRules{}
	if err := c.SearchRead(StockRuleModel, criteria, options, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindStockRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockRuleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockRuleId finds record id by querying it with criteria.
func (c *Client) FindStockRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.rule was not found with criteria %v and options %v", criteria, options)
}
