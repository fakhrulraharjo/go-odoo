package odoo

import (
	"fmt"
)

// IrTranslation represents ir.translation model.
type IrTranslation struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	Comments    *String    `xmlrpc:"comments,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	Lang        *Selection `xmlrpc:"lang,omptempty"`
	Module      *String    `xmlrpc:"module,omptempty"`
	Name        *String    `xmlrpc:"name,omptempty"`
	ResId       *Int       `xmlrpc:"res_id,omptempty"`
	Source      *String    `xmlrpc:"source,omptempty"`
	Src         *String    `xmlrpc:"src,omptempty"`
	State       *Selection `xmlrpc:"state,omptempty"`
	Type        *Selection `xmlrpc:"type,omptempty"`
	Value       *String    `xmlrpc:"value,omptempty"`
}

// IrTranslations represents array of ir.translation model.
type IrTranslations []IrTranslation

// IrTranslationModel is the odoo model name.
const IrTranslationModel = "ir.translation"

// Many2One convert IrTranslation to *Many2One.
func (it *IrTranslation) Many2One() *Many2One {
	return NewMany2One(it.Id.Get(), "")
}

// CreateIrTranslation creates a new ir.translation model and returns its id.
func (c *Client) CreateIrTranslation(it *IrTranslation) (int64, error) {
	return c.Create(IrTranslationModel, it)
}

// UpdateIrTranslation updates an existing ir.translation record.
func (c *Client) UpdateIrTranslation(it *IrTranslation) error {
	return c.UpdateIrTranslations([]int64{it.Id.Get()}, it)
}

// UpdateIrTranslations updates existing ir.translation records.
// All records (represented by ids) will be updated by it values.
func (c *Client) UpdateIrTranslations(ids []int64, it *IrTranslation) error {
	return c.Update(IrTranslationModel, ids, it)
}

// DeleteIrTranslation deletes an existing ir.translation record.
func (c *Client) DeleteIrTranslation(id int64) error {
	return c.DeleteIrTranslations([]int64{id})
}

// DeleteIrTranslations deletes existing ir.translation records.
func (c *Client) DeleteIrTranslations(ids []int64) error {
	return c.Delete(IrTranslationModel, ids)
}

// GetIrTranslation gets ir.translation existing record.
func (c *Client) GetIrTranslation(id int64) (*IrTranslation, error) {
	its, err := c.GetIrTranslations([]int64{id})
	if err != nil {
		return nil, err
	}
	if its != nil && len(*its) > 0 {
		return &((*its)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.translation not found", id)
}

// GetIrTranslations gets ir.translation existing records.
func (c *Client) GetIrTranslations(ids []int64) (*IrTranslations, error) {
	its := &IrTranslations{}
	if err := c.Read(IrTranslationModel, ids, nil, its); err != nil {
		return nil, err
	}
	return its, nil
}

// FindIrTranslation finds ir.translation record by querying it with criteria.
func (c *Client) FindIrTranslation(criteria *Criteria) (*IrTranslation, error) {
	its := &IrTranslations{}
	if err := c.SearchRead(IrTranslationModel, criteria, NewOptions().Limit(1), its); err != nil {
		return nil, err
	}
	if its != nil && len(*its) > 0 {
		return &((*its)[0]), nil
	}
	return nil, fmt.Errorf("ir.translation was not found with criteria %v", criteria)
}

// FindIrTranslations finds ir.translation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrTranslations(criteria *Criteria, options *Options) (*IrTranslations, error) {
	its := &IrTranslations{}
	if err := c.SearchRead(IrTranslationModel, criteria, options, its); err != nil {
		return nil, err
	}
	return its, nil
}

// FindIrTranslationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrTranslationIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrTranslationModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrTranslationId finds record id by querying it with criteria.
func (c *Client) FindIrTranslationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrTranslationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.translation was not found with criteria %v and options %v", criteria, options)
}
