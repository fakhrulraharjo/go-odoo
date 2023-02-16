package odoo

import (
	"fmt"
)

// IrAsset represents ir.asset model.
type IrAsset struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	Active      *Bool      `xmlrpc:"active,omptempty"`
	Bundle      *String    `xmlrpc:"bundle,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	Directive   *Selection `xmlrpc:"directive,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	Name        *String    `xmlrpc:"name,omptempty"`
	Path        *String    `xmlrpc:"path,omptempty"`
	Sequence    *Int       `xmlrpc:"sequence,omptempty"`
	Target      *String    `xmlrpc:"target,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// IrAssets represents array of ir.asset model.
type IrAssets []IrAsset

// IrAssetModel is the odoo model name.
const IrAssetModel = "ir.asset"

// Many2One convert IrAsset to *Many2One.
func (ia *IrAsset) Many2One() *Many2One {
	return NewMany2One(ia.Id.Get(), "")
}

// CreateIrAsset creates a new ir.asset model and returns its id.
func (c *Client) CreateIrAsset(ia *IrAsset) (int64, error) {
	return c.Create(IrAssetModel, ia)
}

// UpdateIrAsset updates an existing ir.asset record.
func (c *Client) UpdateIrAsset(ia *IrAsset) error {
	return c.UpdateIrAssets([]int64{ia.Id.Get()}, ia)
}

// UpdateIrAssets updates existing ir.asset records.
// All records (represented by ids) will be updated by ia values.
func (c *Client) UpdateIrAssets(ids []int64, ia *IrAsset) error {
	return c.Update(IrAssetModel, ids, ia)
}

// DeleteIrAsset deletes an existing ir.asset record.
func (c *Client) DeleteIrAsset(id int64) error {
	return c.DeleteIrAssets([]int64{id})
}

// DeleteIrAssets deletes existing ir.asset records.
func (c *Client) DeleteIrAssets(ids []int64) error {
	return c.Delete(IrAssetModel, ids)
}

// GetIrAsset gets ir.asset existing record.
func (c *Client) GetIrAsset(id int64) (*IrAsset, error) {
	ias, err := c.GetIrAssets([]int64{id})
	if err != nil {
		return nil, err
	}
	if ias != nil && len(*ias) > 0 {
		return &((*ias)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.asset not found", id)
}

// GetIrAssets gets ir.asset existing records.
func (c *Client) GetIrAssets(ids []int64) (*IrAssets, error) {
	ias := &IrAssets{}
	if err := c.Read(IrAssetModel, ids, nil, ias); err != nil {
		return nil, err
	}
	return ias, nil
}

// FindIrAsset finds ir.asset record by querying it with criteria.
func (c *Client) FindIrAsset(criteria *Criteria) (*IrAsset, error) {
	ias := &IrAssets{}
	if err := c.SearchRead(IrAssetModel, criteria, NewOptions().Limit(1), ias); err != nil {
		return nil, err
	}
	if ias != nil && len(*ias) > 0 {
		return &((*ias)[0]), nil
	}
	return nil, fmt.Errorf("ir.asset was not found with criteria %v", criteria)
}

// FindIrAssets finds ir.asset records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrAssets(criteria *Criteria, options *Options) (*IrAssets, error) {
	ias := &IrAssets{}
	if err := c.SearchRead(IrAssetModel, criteria, options, ias); err != nil {
		return nil, err
	}
	return ias, nil
}

// FindIrAssetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrAssetIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrAssetModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrAssetId finds record id by querying it with criteria.
func (c *Client) FindIrAssetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrAssetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.asset was not found with criteria %v and options %v", criteria, options)
}
