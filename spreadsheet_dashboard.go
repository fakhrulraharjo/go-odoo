package odoo

import (
	"fmt"
)

// SpreadsheetDashboard represents spreadsheet.dashboard model.
type SpreadsheetDashboard struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	DashboardGroupId *Many2One `xmlrpc:"dashboard_group_id,omptempty"`
	Data             *String   `xmlrpc:"data,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	GroupIds         *Relation `xmlrpc:"group_ids,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	Name             *String   `xmlrpc:"name,omptempty"`
	Raw              *String   `xmlrpc:"raw,omptempty"`
	Sequence         *Int      `xmlrpc:"sequence,omptempty"`
	Thumbnail        *String   `xmlrpc:"thumbnail,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SpreadsheetDashboards represents array of spreadsheet.dashboard model.
type SpreadsheetDashboards []SpreadsheetDashboard

// SpreadsheetDashboardModel is the odoo model name.
const SpreadsheetDashboardModel = "spreadsheet.dashboard"

// Many2One convert SpreadsheetDashboard to *Many2One.
func (sd *SpreadsheetDashboard) Many2One() *Many2One {
	return NewMany2One(sd.Id.Get(), "")
}

// CreateSpreadsheetDashboard creates a new spreadsheet.dashboard model and returns its id.
func (c *Client) CreateSpreadsheetDashboard(sd *SpreadsheetDashboard) (int64, error) {
	return c.Create(SpreadsheetDashboardModel, sd)
}

// UpdateSpreadsheetDashboard updates an existing spreadsheet.dashboard record.
func (c *Client) UpdateSpreadsheetDashboard(sd *SpreadsheetDashboard) error {
	return c.UpdateSpreadsheetDashboards([]int64{sd.Id.Get()}, sd)
}

// UpdateSpreadsheetDashboards updates existing spreadsheet.dashboard records.
// All records (represented by ids) will be updated by sd values.
func (c *Client) UpdateSpreadsheetDashboards(ids []int64, sd *SpreadsheetDashboard) error {
	return c.Update(SpreadsheetDashboardModel, ids, sd)
}

// DeleteSpreadsheetDashboard deletes an existing spreadsheet.dashboard record.
func (c *Client) DeleteSpreadsheetDashboard(id int64) error {
	return c.DeleteSpreadsheetDashboards([]int64{id})
}

// DeleteSpreadsheetDashboards deletes existing spreadsheet.dashboard records.
func (c *Client) DeleteSpreadsheetDashboards(ids []int64) error {
	return c.Delete(SpreadsheetDashboardModel, ids)
}

// GetSpreadsheetDashboard gets spreadsheet.dashboard existing record.
func (c *Client) GetSpreadsheetDashboard(id int64) (*SpreadsheetDashboard, error) {
	sds, err := c.GetSpreadsheetDashboards([]int64{id})
	if err != nil {
		return nil, err
	}
	if sds != nil && len(*sds) > 0 {
		return &((*sds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of spreadsheet.dashboard not found", id)
}

// GetSpreadsheetDashboards gets spreadsheet.dashboard existing records.
func (c *Client) GetSpreadsheetDashboards(ids []int64) (*SpreadsheetDashboards, error) {
	sds := &SpreadsheetDashboards{}
	if err := c.Read(SpreadsheetDashboardModel, ids, nil, sds); err != nil {
		return nil, err
	}
	return sds, nil
}

// FindSpreadsheetDashboard finds spreadsheet.dashboard record by querying it with criteria.
func (c *Client) FindSpreadsheetDashboard(criteria *Criteria) (*SpreadsheetDashboard, error) {
	sds := &SpreadsheetDashboards{}
	if err := c.SearchRead(SpreadsheetDashboardModel, criteria, NewOptions().Limit(1), sds); err != nil {
		return nil, err
	}
	if sds != nil && len(*sds) > 0 {
		return &((*sds)[0]), nil
	}
	return nil, fmt.Errorf("spreadsheet.dashboard was not found with criteria %v", criteria)
}

// FindSpreadsheetDashboards finds spreadsheet.dashboard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetDashboards(criteria *Criteria, options *Options) (*SpreadsheetDashboards, error) {
	sds := &SpreadsheetDashboards{}
	if err := c.SearchRead(SpreadsheetDashboardModel, criteria, options, sds); err != nil {
		return nil, err
	}
	return sds, nil
}

// FindSpreadsheetDashboardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetDashboardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SpreadsheetDashboardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSpreadsheetDashboardId finds record id by querying it with criteria.
func (c *Client) FindSpreadsheetDashboardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SpreadsheetDashboardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("spreadsheet.dashboard was not found with criteria %v and options %v", criteria, options)
}
