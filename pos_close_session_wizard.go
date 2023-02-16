package odoo

import (
	"fmt"
)

// PosCloseSessionWizard represents pos.close.session.wizard model.
type PosCloseSessionWizard struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	AccountId       *Many2One `xmlrpc:"account_id,omptempty"`
	AccountReadonly *Bool     `xmlrpc:"account_readonly,omptempty"`
	AmountToBalance *Float    `xmlrpc:"amount_to_balance,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	Message         *String   `xmlrpc:"message,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PosCloseSessionWizards represents array of pos.close.session.wizard model.
type PosCloseSessionWizards []PosCloseSessionWizard

// PosCloseSessionWizardModel is the odoo model name.
const PosCloseSessionWizardModel = "pos.close.session.wizard"

// Many2One convert PosCloseSessionWizard to *Many2One.
func (pcsw *PosCloseSessionWizard) Many2One() *Many2One {
	return NewMany2One(pcsw.Id.Get(), "")
}

// CreatePosCloseSessionWizard creates a new pos.close.session.wizard model and returns its id.
func (c *Client) CreatePosCloseSessionWizard(pcsw *PosCloseSessionWizard) (int64, error) {
	return c.Create(PosCloseSessionWizardModel, pcsw)
}

// UpdatePosCloseSessionWizard updates an existing pos.close.session.wizard record.
func (c *Client) UpdatePosCloseSessionWizard(pcsw *PosCloseSessionWizard) error {
	return c.UpdatePosCloseSessionWizards([]int64{pcsw.Id.Get()}, pcsw)
}

// UpdatePosCloseSessionWizards updates existing pos.close.session.wizard records.
// All records (represented by ids) will be updated by pcsw values.
func (c *Client) UpdatePosCloseSessionWizards(ids []int64, pcsw *PosCloseSessionWizard) error {
	return c.Update(PosCloseSessionWizardModel, ids, pcsw)
}

// DeletePosCloseSessionWizard deletes an existing pos.close.session.wizard record.
func (c *Client) DeletePosCloseSessionWizard(id int64) error {
	return c.DeletePosCloseSessionWizards([]int64{id})
}

// DeletePosCloseSessionWizards deletes existing pos.close.session.wizard records.
func (c *Client) DeletePosCloseSessionWizards(ids []int64) error {
	return c.Delete(PosCloseSessionWizardModel, ids)
}

// GetPosCloseSessionWizard gets pos.close.session.wizard existing record.
func (c *Client) GetPosCloseSessionWizard(id int64) (*PosCloseSessionWizard, error) {
	pcsws, err := c.GetPosCloseSessionWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if pcsws != nil && len(*pcsws) > 0 {
		return &((*pcsws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of pos.close.session.wizard not found", id)
}

// GetPosCloseSessionWizards gets pos.close.session.wizard existing records.
func (c *Client) GetPosCloseSessionWizards(ids []int64) (*PosCloseSessionWizards, error) {
	pcsws := &PosCloseSessionWizards{}
	if err := c.Read(PosCloseSessionWizardModel, ids, nil, pcsws); err != nil {
		return nil, err
	}
	return pcsws, nil
}

// FindPosCloseSessionWizard finds pos.close.session.wizard record by querying it with criteria.
func (c *Client) FindPosCloseSessionWizard(criteria *Criteria) (*PosCloseSessionWizard, error) {
	pcsws := &PosCloseSessionWizards{}
	if err := c.SearchRead(PosCloseSessionWizardModel, criteria, NewOptions().Limit(1), pcsws); err != nil {
		return nil, err
	}
	if pcsws != nil && len(*pcsws) > 0 {
		return &((*pcsws)[0]), nil
	}
	return nil, fmt.Errorf("pos.close.session.wizard was not found with criteria %v", criteria)
}

// FindPosCloseSessionWizards finds pos.close.session.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosCloseSessionWizards(criteria *Criteria, options *Options) (*PosCloseSessionWizards, error) {
	pcsws := &PosCloseSessionWizards{}
	if err := c.SearchRead(PosCloseSessionWizardModel, criteria, options, pcsws); err != nil {
		return nil, err
	}
	return pcsws, nil
}

// FindPosCloseSessionWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosCloseSessionWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PosCloseSessionWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPosCloseSessionWizardId finds record id by querying it with criteria.
func (c *Client) FindPosCloseSessionWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosCloseSessionWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("pos.close.session.wizard was not found with criteria %v and options %v", criteria, options)
}
