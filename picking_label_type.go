package odoo

import (
	"fmt"
)

// PickingLabelType represents picking.label.type model.
type PickingLabelType struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	LabelType   *Selection `xmlrpc:"label_type,omptempty"`
	PickingIds  *Relation  `xmlrpc:"picking_ids,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// PickingLabelTypes represents array of picking.label.type model.
type PickingLabelTypes []PickingLabelType

// PickingLabelTypeModel is the odoo model name.
const PickingLabelTypeModel = "picking.label.type"

// Many2One convert PickingLabelType to *Many2One.
func (plt *PickingLabelType) Many2One() *Many2One {
	return NewMany2One(plt.Id.Get(), "")
}

// CreatePickingLabelType creates a new picking.label.type model and returns its id.
func (c *Client) CreatePickingLabelType(plt *PickingLabelType) (int64, error) {
	return c.Create(PickingLabelTypeModel, plt)
}

// UpdatePickingLabelType updates an existing picking.label.type record.
func (c *Client) UpdatePickingLabelType(plt *PickingLabelType) error {
	return c.UpdatePickingLabelTypes([]int64{plt.Id.Get()}, plt)
}

// UpdatePickingLabelTypes updates existing picking.label.type records.
// All records (represented by ids) will be updated by plt values.
func (c *Client) UpdatePickingLabelTypes(ids []int64, plt *PickingLabelType) error {
	return c.Update(PickingLabelTypeModel, ids, plt)
}

// DeletePickingLabelType deletes an existing picking.label.type record.
func (c *Client) DeletePickingLabelType(id int64) error {
	return c.DeletePickingLabelTypes([]int64{id})
}

// DeletePickingLabelTypes deletes existing picking.label.type records.
func (c *Client) DeletePickingLabelTypes(ids []int64) error {
	return c.Delete(PickingLabelTypeModel, ids)
}

// GetPickingLabelType gets picking.label.type existing record.
func (c *Client) GetPickingLabelType(id int64) (*PickingLabelType, error) {
	plts, err := c.GetPickingLabelTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if plts != nil && len(*plts) > 0 {
		return &((*plts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of picking.label.type not found", id)
}

// GetPickingLabelTypes gets picking.label.type existing records.
func (c *Client) GetPickingLabelTypes(ids []int64) (*PickingLabelTypes, error) {
	plts := &PickingLabelTypes{}
	if err := c.Read(PickingLabelTypeModel, ids, nil, plts); err != nil {
		return nil, err
	}
	return plts, nil
}

// FindPickingLabelType finds picking.label.type record by querying it with criteria.
func (c *Client) FindPickingLabelType(criteria *Criteria) (*PickingLabelType, error) {
	plts := &PickingLabelTypes{}
	if err := c.SearchRead(PickingLabelTypeModel, criteria, NewOptions().Limit(1), plts); err != nil {
		return nil, err
	}
	if plts != nil && len(*plts) > 0 {
		return &((*plts)[0]), nil
	}
	return nil, fmt.Errorf("picking.label.type was not found with criteria %v", criteria)
}

// FindPickingLabelTypes finds picking.label.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPickingLabelTypes(criteria *Criteria, options *Options) (*PickingLabelTypes, error) {
	plts := &PickingLabelTypes{}
	if err := c.SearchRead(PickingLabelTypeModel, criteria, options, plts); err != nil {
		return nil, err
	}
	return plts, nil
}

// FindPickingLabelTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPickingLabelTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PickingLabelTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPickingLabelTypeId finds record id by querying it with criteria.
func (c *Client) FindPickingLabelTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PickingLabelTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("picking.label.type was not found with criteria %v and options %v", criteria, options)
}
