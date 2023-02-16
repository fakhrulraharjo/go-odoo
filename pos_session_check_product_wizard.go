package odoo

import (
	"fmt"
)

// PosSessionCheckProductWizard represents pos.session.check_product_wizard model.
type PosSessionCheckProductWizard struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PosSessionCheckProductWizards represents array of pos.session.check_product_wizard model.
type PosSessionCheckProductWizards []PosSessionCheckProductWizard

// PosSessionCheckProductWizardModel is the odoo model name.
const PosSessionCheckProductWizardModel = "pos.session.check_product_wizard"

// Many2One convert PosSessionCheckProductWizard to *Many2One.
func (psc *PosSessionCheckProductWizard) Many2One() *Many2One {
	return NewMany2One(psc.Id.Get(), "")
}

// CreatePosSessionCheckProductWizard creates a new pos.session.check_product_wizard model and returns its id.
func (c *Client) CreatePosSessionCheckProductWizard(psc *PosSessionCheckProductWizard) (int64, error) {
	return c.Create(PosSessionCheckProductWizardModel, psc)
}

// UpdatePosSessionCheckProductWizard updates an existing pos.session.check_product_wizard record.
func (c *Client) UpdatePosSessionCheckProductWizard(psc *PosSessionCheckProductWizard) error {
	return c.UpdatePosSessionCheckProductWizards([]int64{psc.Id.Get()}, psc)
}

// UpdatePosSessionCheckProductWizards updates existing pos.session.check_product_wizard records.
// All records (represented by ids) will be updated by psc values.
func (c *Client) UpdatePosSessionCheckProductWizards(ids []int64, psc *PosSessionCheckProductWizard) error {
	return c.Update(PosSessionCheckProductWizardModel, ids, psc)
}

// DeletePosSessionCheckProductWizard deletes an existing pos.session.check_product_wizard record.
func (c *Client) DeletePosSessionCheckProductWizard(id int64) error {
	return c.DeletePosSessionCheckProductWizards([]int64{id})
}

// DeletePosSessionCheckProductWizards deletes existing pos.session.check_product_wizard records.
func (c *Client) DeletePosSessionCheckProductWizards(ids []int64) error {
	return c.Delete(PosSessionCheckProductWizardModel, ids)
}

// GetPosSessionCheckProductWizard gets pos.session.check_product_wizard existing record.
func (c *Client) GetPosSessionCheckProductWizard(id int64) (*PosSessionCheckProductWizard, error) {
	pscs, err := c.GetPosSessionCheckProductWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if pscs != nil && len(*pscs) > 0 {
		return &((*pscs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of pos.session.check_product_wizard not found", id)
}

// GetPosSessionCheckProductWizards gets pos.session.check_product_wizard existing records.
func (c *Client) GetPosSessionCheckProductWizards(ids []int64) (*PosSessionCheckProductWizards, error) {
	pscs := &PosSessionCheckProductWizards{}
	if err := c.Read(PosSessionCheckProductWizardModel, ids, nil, pscs); err != nil {
		return nil, err
	}
	return pscs, nil
}

// FindPosSessionCheckProductWizard finds pos.session.check_product_wizard record by querying it with criteria.
func (c *Client) FindPosSessionCheckProductWizard(criteria *Criteria) (*PosSessionCheckProductWizard, error) {
	pscs := &PosSessionCheckProductWizards{}
	if err := c.SearchRead(PosSessionCheckProductWizardModel, criteria, NewOptions().Limit(1), pscs); err != nil {
		return nil, err
	}
	if pscs != nil && len(*pscs) > 0 {
		return &((*pscs)[0]), nil
	}
	return nil, fmt.Errorf("pos.session.check_product_wizard was not found with criteria %v", criteria)
}

// FindPosSessionCheckProductWizards finds pos.session.check_product_wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosSessionCheckProductWizards(criteria *Criteria, options *Options) (*PosSessionCheckProductWizards, error) {
	pscs := &PosSessionCheckProductWizards{}
	if err := c.SearchRead(PosSessionCheckProductWizardModel, criteria, options, pscs); err != nil {
		return nil, err
	}
	return pscs, nil
}

// FindPosSessionCheckProductWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosSessionCheckProductWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PosSessionCheckProductWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPosSessionCheckProductWizardId finds record id by querying it with criteria.
func (c *Client) FindPosSessionCheckProductWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosSessionCheckProductWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("pos.session.check_product_wizard was not found with criteria %v and options %v", criteria, options)
}
