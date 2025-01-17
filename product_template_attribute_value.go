package odoo

import (
	"fmt"
)

// ProductTemplateAttributeValue represents product.template.attribute.value model.
type ProductTemplateAttributeValue struct {
	LastUpdate              *Time      `xmlrpc:"__last_update,omptempty"`
	AttributeId             *Many2One  `xmlrpc:"attribute_id,omptempty"`
	AttributeLineId         *Many2One  `xmlrpc:"attribute_line_id,omptempty"`
	Color                   *Int       `xmlrpc:"color,omptempty"`
	CreateDate              *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid               *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId              *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName             *String    `xmlrpc:"display_name,omptempty"`
	DisplayType             *Selection `xmlrpc:"display_type,omptempty"`
	ExcludeFor              *Relation  `xmlrpc:"exclude_for,omptempty"`
	HtmlColor               *String    `xmlrpc:"html_color,omptempty"`
	Id                      *Int       `xmlrpc:"id,omptempty"`
	IsCustom                *Bool      `xmlrpc:"is_custom,omptempty"`
	Name                    *String    `xmlrpc:"name,omptempty"`
	PriceExtra              *Float     `xmlrpc:"price_extra,omptempty"`
	ProductAttributeValueId *Many2One  `xmlrpc:"product_attribute_value_id,omptempty"`
	ProductTmplId           *Many2One  `xmlrpc:"product_tmpl_id,omptempty"`
	PtavActive              *Bool      `xmlrpc:"ptav_active,omptempty"`
	PtavProductVariantIds   *Relation  `xmlrpc:"ptav_product_variant_ids,omptempty"`
	WriteDate               *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ProductTemplateAttributeValues represents array of product.template.attribute.value model.
type ProductTemplateAttributeValues []ProductTemplateAttributeValue

// ProductTemplateAttributeValueModel is the odoo model name.
const ProductTemplateAttributeValueModel = "product.template.attribute.value"

// Many2One convert ProductTemplateAttributeValue to *Many2One.
func (ptav *ProductTemplateAttributeValue) Many2One() *Many2One {
	return NewMany2One(ptav.Id.Get(), "")
}

// CreateProductTemplateAttributeValue creates a new product.template.attribute.value model and returns its id.
func (c *Client) CreateProductTemplateAttributeValue(ptav *ProductTemplateAttributeValue) (int64, error) {
	return c.Create(ProductTemplateAttributeValueModel, ptav)
}

// UpdateProductTemplateAttributeValue updates an existing product.template.attribute.value record.
func (c *Client) UpdateProductTemplateAttributeValue(ptav *ProductTemplateAttributeValue) error {
	return c.UpdateProductTemplateAttributeValues([]int64{ptav.Id.Get()}, ptav)
}

// UpdateProductTemplateAttributeValues updates existing product.template.attribute.value records.
// All records (represented by ids) will be updated by ptav values.
func (c *Client) UpdateProductTemplateAttributeValues(ids []int64, ptav *ProductTemplateAttributeValue) error {
	return c.Update(ProductTemplateAttributeValueModel, ids, ptav)
}

// DeleteProductTemplateAttributeValue deletes an existing product.template.attribute.value record.
func (c *Client) DeleteProductTemplateAttributeValue(id int64) error {
	return c.DeleteProductTemplateAttributeValues([]int64{id})
}

// DeleteProductTemplateAttributeValues deletes existing product.template.attribute.value records.
func (c *Client) DeleteProductTemplateAttributeValues(ids []int64) error {
	return c.Delete(ProductTemplateAttributeValueModel, ids)
}

// GetProductTemplateAttributeValue gets product.template.attribute.value existing record.
func (c *Client) GetProductTemplateAttributeValue(id int64) (*ProductTemplateAttributeValue, error) {
	ptavs, err := c.GetProductTemplateAttributeValues([]int64{id})
	if err != nil {
		return nil, err
	}
	if ptavs != nil && len(*ptavs) > 0 {
		return &((*ptavs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.template.attribute.value not found", id)
}

// GetProductTemplateAttributeValues gets product.template.attribute.value existing records.
func (c *Client) GetProductTemplateAttributeValues(ids []int64) (*ProductTemplateAttributeValues, error) {
	ptavs := &ProductTemplateAttributeValues{}
	if err := c.Read(ProductTemplateAttributeValueModel, ids, nil, ptavs); err != nil {
		return nil, err
	}
	return ptavs, nil
}

// FindProductTemplateAttributeValue finds product.template.attribute.value record by querying it with criteria.
func (c *Client) FindProductTemplateAttributeValue(criteria *Criteria) (*ProductTemplateAttributeValue, error) {
	ptavs := &ProductTemplateAttributeValues{}
	if err := c.SearchRead(ProductTemplateAttributeValueModel, criteria, NewOptions().Limit(1), ptavs); err != nil {
		return nil, err
	}
	if ptavs != nil && len(*ptavs) > 0 {
		return &((*ptavs)[0]), nil
	}
	return nil, fmt.Errorf("product.template.attribute.value was not found with criteria %v", criteria)
}

// FindProductTemplateAttributeValues finds product.template.attribute.value records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductTemplateAttributeValues(criteria *Criteria, options *Options) (*ProductTemplateAttributeValues, error) {
	ptavs := &ProductTemplateAttributeValues{}
	if err := c.SearchRead(ProductTemplateAttributeValueModel, criteria, options, ptavs); err != nil {
		return nil, err
	}
	return ptavs, nil
}

// FindProductTemplateAttributeValueIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductTemplateAttributeValueIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductTemplateAttributeValueModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductTemplateAttributeValueId finds record id by querying it with criteria.
func (c *Client) FindProductTemplateAttributeValueId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductTemplateAttributeValueModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.template.attribute.value was not found with criteria %v and options %v", criteria, options)
}
