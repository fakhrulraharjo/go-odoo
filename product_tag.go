package odoo

import (
	"fmt"
)

// ProductTag represents product.tag model.
type ProductTag struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omptempty"`
	Color              *Int      `xmlrpc:"color,omptempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName        *String   `xmlrpc:"display_name,omptempty"`
	Id                 *Int      `xmlrpc:"id,omptempty"`
	Name               *String   `xmlrpc:"name,omptempty"`
	ProductIds         *Relation `xmlrpc:"product_ids,omptempty"`
	ProductProductIds  *Relation `xmlrpc:"product_product_ids,omptempty"`
	ProductTemplateIds *Relation `xmlrpc:"product_template_ids,omptempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProductTags represents array of product.tag model.
type ProductTags []ProductTag

// ProductTagModel is the odoo model name.
const ProductTagModel = "product.tag"

// Many2One convert ProductTag to *Many2One.
func (pt *ProductTag) Many2One() *Many2One {
	return NewMany2One(pt.Id.Get(), "")
}

// CreateProductTag creates a new product.tag model and returns its id.
func (c *Client) CreateProductTag(pt *ProductTag) (int64, error) {
	return c.Create(ProductTagModel, pt)
}

// UpdateProductTag updates an existing product.tag record.
func (c *Client) UpdateProductTag(pt *ProductTag) error {
	return c.UpdateProductTags([]int64{pt.Id.Get()}, pt)
}

// UpdateProductTags updates existing product.tag records.
// All records (represented by ids) will be updated by pt values.
func (c *Client) UpdateProductTags(ids []int64, pt *ProductTag) error {
	return c.Update(ProductTagModel, ids, pt)
}

// DeleteProductTag deletes an existing product.tag record.
func (c *Client) DeleteProductTag(id int64) error {
	return c.DeleteProductTags([]int64{id})
}

// DeleteProductTags deletes existing product.tag records.
func (c *Client) DeleteProductTags(ids []int64) error {
	return c.Delete(ProductTagModel, ids)
}

// GetProductTag gets product.tag existing record.
func (c *Client) GetProductTag(id int64) (*ProductTag, error) {
	pts, err := c.GetProductTags([]int64{id})
	if err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.tag not found", id)
}

// GetProductTags gets product.tag existing records.
func (c *Client) GetProductTags(ids []int64) (*ProductTags, error) {
	pts := &ProductTags{}
	if err := c.Read(ProductTagModel, ids, nil, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProductTag finds product.tag record by querying it with criteria.
func (c *Client) FindProductTag(criteria *Criteria) (*ProductTag, error) {
	pts := &ProductTags{}
	if err := c.SearchRead(ProductTagModel, criteria, NewOptions().Limit(1), pts); err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("product.tag was not found with criteria %v", criteria)
}

// FindProductTags finds product.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductTags(criteria *Criteria, options *Options) (*ProductTags, error) {
	pts := &ProductTags{}
	if err := c.SearchRead(ProductTagModel, criteria, options, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProductTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductTagModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductTagId finds record id by querying it with criteria.
func (c *Client) FindProductTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.tag was not found with criteria %v and options %v", criteria, options)
}
