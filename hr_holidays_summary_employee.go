package odoo

import (
	"fmt"
)

// HrHolidaysSummaryEmployee represents hr.holidays.summary.employee model.
type HrHolidaysSummaryEmployee struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateFrom    *Time      `xmlrpc:"date_from,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Emp         *Relation  `xmlrpc:"emp,omptempty"`
	HolidayType *Selection `xmlrpc:"holiday_type,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrHolidaysSummaryEmployees represents array of hr.holidays.summary.employee model.
type HrHolidaysSummaryEmployees []HrHolidaysSummaryEmployee

// HrHolidaysSummaryEmployeeModel is the odoo model name.
const HrHolidaysSummaryEmployeeModel = "hr.holidays.summary.employee"

// Many2One convert HrHolidaysSummaryEmployee to *Many2One.
func (hhse *HrHolidaysSummaryEmployee) Many2One() *Many2One {
	return NewMany2One(hhse.Id.Get(), "")
}

// CreateHrHolidaysSummaryEmployee creates a new hr.holidays.summary.employee model and returns its id.
func (c *Client) CreateHrHolidaysSummaryEmployee(hhse *HrHolidaysSummaryEmployee) (int64, error) {
	return c.Create(HrHolidaysSummaryEmployeeModel, hhse)
}

// UpdateHrHolidaysSummaryEmployee updates an existing hr.holidays.summary.employee record.
func (c *Client) UpdateHrHolidaysSummaryEmployee(hhse *HrHolidaysSummaryEmployee) error {
	return c.UpdateHrHolidaysSummaryEmployees([]int64{hhse.Id.Get()}, hhse)
}

// UpdateHrHolidaysSummaryEmployees updates existing hr.holidays.summary.employee records.
// All records (represented by ids) will be updated by hhse values.
func (c *Client) UpdateHrHolidaysSummaryEmployees(ids []int64, hhse *HrHolidaysSummaryEmployee) error {
	return c.Update(HrHolidaysSummaryEmployeeModel, ids, hhse)
}

// DeleteHrHolidaysSummaryEmployee deletes an existing hr.holidays.summary.employee record.
func (c *Client) DeleteHrHolidaysSummaryEmployee(id int64) error {
	return c.DeleteHrHolidaysSummaryEmployees([]int64{id})
}

// DeleteHrHolidaysSummaryEmployees deletes existing hr.holidays.summary.employee records.
func (c *Client) DeleteHrHolidaysSummaryEmployees(ids []int64) error {
	return c.Delete(HrHolidaysSummaryEmployeeModel, ids)
}

// GetHrHolidaysSummaryEmployee gets hr.holidays.summary.employee existing record.
func (c *Client) GetHrHolidaysSummaryEmployee(id int64) (*HrHolidaysSummaryEmployee, error) {
	hhses, err := c.GetHrHolidaysSummaryEmployees([]int64{id})
	if err != nil {
		return nil, err
	}
	if hhses != nil && len(*hhses) > 0 {
		return &((*hhses)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.holidays.summary.employee not found", id)
}

// GetHrHolidaysSummaryEmployees gets hr.holidays.summary.employee existing records.
func (c *Client) GetHrHolidaysSummaryEmployees(ids []int64) (*HrHolidaysSummaryEmployees, error) {
	hhses := &HrHolidaysSummaryEmployees{}
	if err := c.Read(HrHolidaysSummaryEmployeeModel, ids, nil, hhses); err != nil {
		return nil, err
	}
	return hhses, nil
}

// FindHrHolidaysSummaryEmployee finds hr.holidays.summary.employee record by querying it with criteria.
func (c *Client) FindHrHolidaysSummaryEmployee(criteria *Criteria) (*HrHolidaysSummaryEmployee, error) {
	hhses := &HrHolidaysSummaryEmployees{}
	if err := c.SearchRead(HrHolidaysSummaryEmployeeModel, criteria, NewOptions().Limit(1), hhses); err != nil {
		return nil, err
	}
	if hhses != nil && len(*hhses) > 0 {
		return &((*hhses)[0]), nil
	}
	return nil, fmt.Errorf("hr.holidays.summary.employee was not found with criteria %v", criteria)
}

// FindHrHolidaysSummaryEmployees finds hr.holidays.summary.employee records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrHolidaysSummaryEmployees(criteria *Criteria, options *Options) (*HrHolidaysSummaryEmployees, error) {
	hhses := &HrHolidaysSummaryEmployees{}
	if err := c.SearchRead(HrHolidaysSummaryEmployeeModel, criteria, options, hhses); err != nil {
		return nil, err
	}
	return hhses, nil
}

// FindHrHolidaysSummaryEmployeeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrHolidaysSummaryEmployeeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrHolidaysSummaryEmployeeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrHolidaysSummaryEmployeeId finds record id by querying it with criteria.
func (c *Client) FindHrHolidaysSummaryEmployeeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrHolidaysSummaryEmployeeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.holidays.summary.employee was not found with criteria %v and options %v", criteria, options)
}
