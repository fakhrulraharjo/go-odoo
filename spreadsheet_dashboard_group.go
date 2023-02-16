package odoo

import (
	"fmt"
)

// SpreadsheetDashboardGroup represents spreadsheet.dashboard.group model.
type SpreadsheetDashboardGroup struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DashboardIds *Relation `xmlrpc:"dashboard_ids,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	Name         *String   `xmlrpc:"name,omptempty"`
	Sequence     *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SpreadsheetDashboardGroups represents array of spreadsheet.dashboard.group model.
type SpreadsheetDashboardGroups []SpreadsheetDashboardGroup

// SpreadsheetDashboardGroupModel is the odoo model name.
const SpreadsheetDashboardGroupModel = "spreadsheet.dashboard.group"

// Many2One convert SpreadsheetDashboardGroup to *Many2One.
func (sdg *SpreadsheetDashboardGroup) Many2One() *Many2One {
	return NewMany2One(sdg.Id.Get(), "")
}

// CreateSpreadsheetDashboardGroup creates a new spreadsheet.dashboard.group model and returns its id.
func (c *Client) CreateSpreadsheetDashboardGroup(sdg *SpreadsheetDashboardGroup) (int64, error) {
	return c.Create(SpreadsheetDashboardGroupModel, sdg)
}

// UpdateSpreadsheetDashboardGroup updates an existing spreadsheet.dashboard.group record.
func (c *Client) UpdateSpreadsheetDashboardGroup(sdg *SpreadsheetDashboardGroup) error {
	return c.UpdateSpreadsheetDashboardGroups([]int64{sdg.Id.Get()}, sdg)
}

// UpdateSpreadsheetDashboardGroups updates existing spreadsheet.dashboard.group records.
// All records (represented by ids) will be updated by sdg values.
func (c *Client) UpdateSpreadsheetDashboardGroups(ids []int64, sdg *SpreadsheetDashboardGroup) error {
	return c.Update(SpreadsheetDashboardGroupModel, ids, sdg)
}

// DeleteSpreadsheetDashboardGroup deletes an existing spreadsheet.dashboard.group record.
func (c *Client) DeleteSpreadsheetDashboardGroup(id int64) error {
	return c.DeleteSpreadsheetDashboardGroups([]int64{id})
}

// DeleteSpreadsheetDashboardGroups deletes existing spreadsheet.dashboard.group records.
func (c *Client) DeleteSpreadsheetDashboardGroups(ids []int64) error {
	return c.Delete(SpreadsheetDashboardGroupModel, ids)
}

// GetSpreadsheetDashboardGroup gets spreadsheet.dashboard.group existing record.
func (c *Client) GetSpreadsheetDashboardGroup(id int64) (*SpreadsheetDashboardGroup, error) {
	sdgs, err := c.GetSpreadsheetDashboardGroups([]int64{id})
	if err != nil {
		return nil, err
	}
	if sdgs != nil && len(*sdgs) > 0 {
		return &((*sdgs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of spreadsheet.dashboard.group not found", id)
}

// GetSpreadsheetDashboardGroups gets spreadsheet.dashboard.group existing records.
func (c *Client) GetSpreadsheetDashboardGroups(ids []int64) (*SpreadsheetDashboardGroups, error) {
	sdgs := &SpreadsheetDashboardGroups{}
	if err := c.Read(SpreadsheetDashboardGroupModel, ids, nil, sdgs); err != nil {
		return nil, err
	}
	return sdgs, nil
}

// FindSpreadsheetDashboardGroup finds spreadsheet.dashboard.group record by querying it with criteria.
func (c *Client) FindSpreadsheetDashboardGroup(criteria *Criteria) (*SpreadsheetDashboardGroup, error) {
	sdgs := &SpreadsheetDashboardGroups{}
	if err := c.SearchRead(SpreadsheetDashboardGroupModel, criteria, NewOptions().Limit(1), sdgs); err != nil {
		return nil, err
	}
	if sdgs != nil && len(*sdgs) > 0 {
		return &((*sdgs)[0]), nil
	}
	return nil, fmt.Errorf("spreadsheet.dashboard.group was not found with criteria %v", criteria)
}

// FindSpreadsheetDashboardGroups finds spreadsheet.dashboard.group records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetDashboardGroups(criteria *Criteria, options *Options) (*SpreadsheetDashboardGroups, error) {
	sdgs := &SpreadsheetDashboardGroups{}
	if err := c.SearchRead(SpreadsheetDashboardGroupModel, criteria, options, sdgs); err != nil {
		return nil, err
	}
	return sdgs, nil
}

// FindSpreadsheetDashboardGroupIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetDashboardGroupIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SpreadsheetDashboardGroupModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSpreadsheetDashboardGroupId finds record id by querying it with criteria.
func (c *Client) FindSpreadsheetDashboardGroupId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SpreadsheetDashboardGroupModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("spreadsheet.dashboard.group was not found with criteria %v and options %v", criteria, options)
}
