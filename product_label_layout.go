package odoo

import (
	"fmt"
)

// ProductLabelLayout represents product.label.layout model.
type ProductLabelLayout struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	Columns         *Int       `xmlrpc:"columns,omptempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omptempty"`
	CustomQuantity  *Int       `xmlrpc:"custom_quantity,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	ExtraHtml       *String    `xmlrpc:"extra_html,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	MoveLineIds     *Relation  `xmlrpc:"move_line_ids,omptempty"`
	PickingQuantity *Selection `xmlrpc:"picking_quantity,omptempty"`
	PrintFormat     *Selection `xmlrpc:"print_format,omptempty"`
	ProductIds      *Relation  `xmlrpc:"product_ids,omptempty"`
	ProductTmplIds  *Relation  `xmlrpc:"product_tmpl_ids,omptempty"`
	Rows            *Int       `xmlrpc:"rows,omptempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ProductLabelLayouts represents array of product.label.layout model.
type ProductLabelLayouts []ProductLabelLayout

// ProductLabelLayoutModel is the odoo model name.
const ProductLabelLayoutModel = "product.label.layout"

// Many2One convert ProductLabelLayout to *Many2One.
func (pll *ProductLabelLayout) Many2One() *Many2One {
	return NewMany2One(pll.Id.Get(), "")
}

// CreateProductLabelLayout creates a new product.label.layout model and returns its id.
func (c *Client) CreateProductLabelLayout(pll *ProductLabelLayout) (int64, error) {
	return c.Create(ProductLabelLayoutModel, pll)
}

// UpdateProductLabelLayout updates an existing product.label.layout record.
func (c *Client) UpdateProductLabelLayout(pll *ProductLabelLayout) error {
	return c.UpdateProductLabelLayouts([]int64{pll.Id.Get()}, pll)
}

// UpdateProductLabelLayouts updates existing product.label.layout records.
// All records (represented by ids) will be updated by pll values.
func (c *Client) UpdateProductLabelLayouts(ids []int64, pll *ProductLabelLayout) error {
	return c.Update(ProductLabelLayoutModel, ids, pll)
}

// DeleteProductLabelLayout deletes an existing product.label.layout record.
func (c *Client) DeleteProductLabelLayout(id int64) error {
	return c.DeleteProductLabelLayouts([]int64{id})
}

// DeleteProductLabelLayouts deletes existing product.label.layout records.
func (c *Client) DeleteProductLabelLayouts(ids []int64) error {
	return c.Delete(ProductLabelLayoutModel, ids)
}

// GetProductLabelLayout gets product.label.layout existing record.
func (c *Client) GetProductLabelLayout(id int64) (*ProductLabelLayout, error) {
	plls, err := c.GetProductLabelLayouts([]int64{id})
	if err != nil {
		return nil, err
	}
	if plls != nil && len(*plls) > 0 {
		return &((*plls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.label.layout not found", id)
}

// GetProductLabelLayouts gets product.label.layout existing records.
func (c *Client) GetProductLabelLayouts(ids []int64) (*ProductLabelLayouts, error) {
	plls := &ProductLabelLayouts{}
	if err := c.Read(ProductLabelLayoutModel, ids, nil, plls); err != nil {
		return nil, err
	}
	return plls, nil
}

// FindProductLabelLayout finds product.label.layout record by querying it with criteria.
func (c *Client) FindProductLabelLayout(criteria *Criteria) (*ProductLabelLayout, error) {
	plls := &ProductLabelLayouts{}
	if err := c.SearchRead(ProductLabelLayoutModel, criteria, NewOptions().Limit(1), plls); err != nil {
		return nil, err
	}
	if plls != nil && len(*plls) > 0 {
		return &((*plls)[0]), nil
	}
	return nil, fmt.Errorf("product.label.layout was not found with criteria %v", criteria)
}

// FindProductLabelLayouts finds product.label.layout records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductLabelLayouts(criteria *Criteria, options *Options) (*ProductLabelLayouts, error) {
	plls := &ProductLabelLayouts{}
	if err := c.SearchRead(ProductLabelLayoutModel, criteria, options, plls); err != nil {
		return nil, err
	}
	return plls, nil
}

// FindProductLabelLayoutIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductLabelLayoutIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductLabelLayoutModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductLabelLayoutId finds record id by querying it with criteria.
func (c *Client) FindProductLabelLayoutId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductLabelLayoutModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.label.layout was not found with criteria %v and options %v", criteria, options)
}
