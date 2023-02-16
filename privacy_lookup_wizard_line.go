package odoo

import (
	"fmt"
)

// PrivacyLookupWizardLine represents privacy.lookup.wizard.line model.
type PrivacyLookupWizardLine struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	ExecutionDetails *String   `xmlrpc:"execution_details,omptempty"`
	HasActive        *Bool     `xmlrpc:"has_active,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	IsActive         *Bool     `xmlrpc:"is_active,omptempty"`
	IsUnlinked       *Bool     `xmlrpc:"is_unlinked,omptempty"`
	ResId            *Int      `xmlrpc:"res_id,omptempty"`
	ResModel         *String   `xmlrpc:"res_model,omptempty"`
	ResModelId       *Many2One `xmlrpc:"res_model_id,omptempty"`
	ResName          *String   `xmlrpc:"res_name,omptempty"`
	ResourceRef      *String   `xmlrpc:"resource_ref,omptempty"`
	WizardId         *Many2One `xmlrpc:"wizard_id,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PrivacyLookupWizardLines represents array of privacy.lookup.wizard.line model.
type PrivacyLookupWizardLines []PrivacyLookupWizardLine

// PrivacyLookupWizardLineModel is the odoo model name.
const PrivacyLookupWizardLineModel = "privacy.lookup.wizard.line"

// Many2One convert PrivacyLookupWizardLine to *Many2One.
func (plwl *PrivacyLookupWizardLine) Many2One() *Many2One {
	return NewMany2One(plwl.Id.Get(), "")
}

// CreatePrivacyLookupWizardLine creates a new privacy.lookup.wizard.line model and returns its id.
func (c *Client) CreatePrivacyLookupWizardLine(plwl *PrivacyLookupWizardLine) (int64, error) {
	return c.Create(PrivacyLookupWizardLineModel, plwl)
}

// UpdatePrivacyLookupWizardLine updates an existing privacy.lookup.wizard.line record.
func (c *Client) UpdatePrivacyLookupWizardLine(plwl *PrivacyLookupWizardLine) error {
	return c.UpdatePrivacyLookupWizardLines([]int64{plwl.Id.Get()}, plwl)
}

// UpdatePrivacyLookupWizardLines updates existing privacy.lookup.wizard.line records.
// All records (represented by ids) will be updated by plwl values.
func (c *Client) UpdatePrivacyLookupWizardLines(ids []int64, plwl *PrivacyLookupWizardLine) error {
	return c.Update(PrivacyLookupWizardLineModel, ids, plwl)
}

// DeletePrivacyLookupWizardLine deletes an existing privacy.lookup.wizard.line record.
func (c *Client) DeletePrivacyLookupWizardLine(id int64) error {
	return c.DeletePrivacyLookupWizardLines([]int64{id})
}

// DeletePrivacyLookupWizardLines deletes existing privacy.lookup.wizard.line records.
func (c *Client) DeletePrivacyLookupWizardLines(ids []int64) error {
	return c.Delete(PrivacyLookupWizardLineModel, ids)
}

// GetPrivacyLookupWizardLine gets privacy.lookup.wizard.line existing record.
func (c *Client) GetPrivacyLookupWizardLine(id int64) (*PrivacyLookupWizardLine, error) {
	plwls, err := c.GetPrivacyLookupWizardLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if plwls != nil && len(*plwls) > 0 {
		return &((*plwls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of privacy.lookup.wizard.line not found", id)
}

// GetPrivacyLookupWizardLines gets privacy.lookup.wizard.line existing records.
func (c *Client) GetPrivacyLookupWizardLines(ids []int64) (*PrivacyLookupWizardLines, error) {
	plwls := &PrivacyLookupWizardLines{}
	if err := c.Read(PrivacyLookupWizardLineModel, ids, nil, plwls); err != nil {
		return nil, err
	}
	return plwls, nil
}

// FindPrivacyLookupWizardLine finds privacy.lookup.wizard.line record by querying it with criteria.
func (c *Client) FindPrivacyLookupWizardLine(criteria *Criteria) (*PrivacyLookupWizardLine, error) {
	plwls := &PrivacyLookupWizardLines{}
	if err := c.SearchRead(PrivacyLookupWizardLineModel, criteria, NewOptions().Limit(1), plwls); err != nil {
		return nil, err
	}
	if plwls != nil && len(*plwls) > 0 {
		return &((*plwls)[0]), nil
	}
	return nil, fmt.Errorf("privacy.lookup.wizard.line was not found with criteria %v", criteria)
}

// FindPrivacyLookupWizardLines finds privacy.lookup.wizard.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPrivacyLookupWizardLines(criteria *Criteria, options *Options) (*PrivacyLookupWizardLines, error) {
	plwls := &PrivacyLookupWizardLines{}
	if err := c.SearchRead(PrivacyLookupWizardLineModel, criteria, options, plwls); err != nil {
		return nil, err
	}
	return plwls, nil
}

// FindPrivacyLookupWizardLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPrivacyLookupWizardLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PrivacyLookupWizardLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPrivacyLookupWizardLineId finds record id by querying it with criteria.
func (c *Client) FindPrivacyLookupWizardLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PrivacyLookupWizardLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("privacy.lookup.wizard.line was not found with criteria %v and options %v", criteria, options)
}
