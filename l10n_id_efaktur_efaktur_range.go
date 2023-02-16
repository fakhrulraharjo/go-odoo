package odoo

import (
	"fmt"
)

// L10NIdEfakturEfakturRange represents l10n_id_efaktur.efaktur.range model.
type L10NIdEfakturEfakturRange struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Available   *Int      `xmlrpc:"available,omptempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Max         *String   `xmlrpc:"max,omptempty"`
	Min         *String   `xmlrpc:"min,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// L10NIdEfakturEfakturRanges represents array of l10n_id_efaktur.efaktur.range model.
type L10NIdEfakturEfakturRanges []L10NIdEfakturEfakturRange

// L10NIdEfakturEfakturRangeModel is the odoo model name.
const L10NIdEfakturEfakturRangeModel = "l10n_id_efaktur.efaktur.range"

// Many2One convert L10NIdEfakturEfakturRange to *Many2One.
func (ler *L10NIdEfakturEfakturRange) Many2One() *Many2One {
	return NewMany2One(ler.Id.Get(), "")
}

// CreateL10NIdEfakturEfakturRange creates a new l10n_id_efaktur.efaktur.range model and returns its id.
func (c *Client) CreateL10NIdEfakturEfakturRange(ler *L10NIdEfakturEfakturRange) (int64, error) {
	return c.Create(L10NIdEfakturEfakturRangeModel, ler)
}

// UpdateL10NIdEfakturEfakturRange updates an existing l10n_id_efaktur.efaktur.range record.
func (c *Client) UpdateL10NIdEfakturEfakturRange(ler *L10NIdEfakturEfakturRange) error {
	return c.UpdateL10NIdEfakturEfakturRanges([]int64{ler.Id.Get()}, ler)
}

// UpdateL10NIdEfakturEfakturRanges updates existing l10n_id_efaktur.efaktur.range records.
// All records (represented by ids) will be updated by ler values.
func (c *Client) UpdateL10NIdEfakturEfakturRanges(ids []int64, ler *L10NIdEfakturEfakturRange) error {
	return c.Update(L10NIdEfakturEfakturRangeModel, ids, ler)
}

// DeleteL10NIdEfakturEfakturRange deletes an existing l10n_id_efaktur.efaktur.range record.
func (c *Client) DeleteL10NIdEfakturEfakturRange(id int64) error {
	return c.DeleteL10NIdEfakturEfakturRanges([]int64{id})
}

// DeleteL10NIdEfakturEfakturRanges deletes existing l10n_id_efaktur.efaktur.range records.
func (c *Client) DeleteL10NIdEfakturEfakturRanges(ids []int64) error {
	return c.Delete(L10NIdEfakturEfakturRangeModel, ids)
}

// GetL10NIdEfakturEfakturRange gets l10n_id_efaktur.efaktur.range existing record.
func (c *Client) GetL10NIdEfakturEfakturRange(id int64) (*L10NIdEfakturEfakturRange, error) {
	lers, err := c.GetL10NIdEfakturEfakturRanges([]int64{id})
	if err != nil {
		return nil, err
	}
	if lers != nil && len(*lers) > 0 {
		return &((*lers)[0]), nil
	}
	return nil, fmt.Errorf("id %v of l10n_id_efaktur.efaktur.range not found", id)
}

// GetL10NIdEfakturEfakturRanges gets l10n_id_efaktur.efaktur.range existing records.
func (c *Client) GetL10NIdEfakturEfakturRanges(ids []int64) (*L10NIdEfakturEfakturRanges, error) {
	lers := &L10NIdEfakturEfakturRanges{}
	if err := c.Read(L10NIdEfakturEfakturRangeModel, ids, nil, lers); err != nil {
		return nil, err
	}
	return lers, nil
}

// FindL10NIdEfakturEfakturRange finds l10n_id_efaktur.efaktur.range record by querying it with criteria.
func (c *Client) FindL10NIdEfakturEfakturRange(criteria *Criteria) (*L10NIdEfakturEfakturRange, error) {
	lers := &L10NIdEfakturEfakturRanges{}
	if err := c.SearchRead(L10NIdEfakturEfakturRangeModel, criteria, NewOptions().Limit(1), lers); err != nil {
		return nil, err
	}
	if lers != nil && len(*lers) > 0 {
		return &((*lers)[0]), nil
	}
	return nil, fmt.Errorf("l10n_id_efaktur.efaktur.range was not found with criteria %v", criteria)
}

// FindL10NIdEfakturEfakturRanges finds l10n_id_efaktur.efaktur.range records by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NIdEfakturEfakturRanges(criteria *Criteria, options *Options) (*L10NIdEfakturEfakturRanges, error) {
	lers := &L10NIdEfakturEfakturRanges{}
	if err := c.SearchRead(L10NIdEfakturEfakturRangeModel, criteria, options, lers); err != nil {
		return nil, err
	}
	return lers, nil
}

// FindL10NIdEfakturEfakturRangeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NIdEfakturEfakturRangeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(L10NIdEfakturEfakturRangeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindL10NIdEfakturEfakturRangeId finds record id by querying it with criteria.
func (c *Client) FindL10NIdEfakturEfakturRangeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(L10NIdEfakturEfakturRangeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("l10n_id_efaktur.efaktur.range was not found with criteria %v and options %v", criteria, options)
}
