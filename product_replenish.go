package odoo

import (
	"fmt"
)

// ProductReplenish represents product.replenish model.
type ProductReplenish struct {
	LastUpdate           *Time     `xmlrpc:"__last_update,omptempty"`
	CompanyId            *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omptempty"`
	DatePlanned          *Time     `xmlrpc:"date_planned,omptempty"`
	DisplayName          *String   `xmlrpc:"display_name,omptempty"`
	Id                   *Int      `xmlrpc:"id,omptempty"`
	ProductHasVariants   *Bool     `xmlrpc:"product_has_variants,omptempty"`
	ProductId            *Many2One `xmlrpc:"product_id,omptempty"`
	ProductTmplId        *Many2One `xmlrpc:"product_tmpl_id,omptempty"`
	ProductUomCategoryId *Many2One `xmlrpc:"product_uom_category_id,omptempty"`
	ProductUomId         *Many2One `xmlrpc:"product_uom_id,omptempty"`
	Quantity             *Float    `xmlrpc:"quantity,omptempty"`
	RouteIds             *Relation `xmlrpc:"route_ids,omptempty"`
	WarehouseId          *Many2One `xmlrpc:"warehouse_id,omptempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProductReplenishs represents array of product.replenish model.
type ProductReplenishs []ProductReplenish

// ProductReplenishModel is the odoo model name.
const ProductReplenishModel = "product.replenish"

// Many2One convert ProductReplenish to *Many2One.
func (pr *ProductReplenish) Many2One() *Many2One {
	return NewMany2One(pr.Id.Get(), "")
}

// CreateProductReplenish creates a new product.replenish model and returns its id.
func (c *Client) CreateProductReplenish(pr *ProductReplenish) (int64, error) {
	return c.Create(ProductReplenishModel, pr)
}

// UpdateProductReplenish updates an existing product.replenish record.
func (c *Client) UpdateProductReplenish(pr *ProductReplenish) error {
	return c.UpdateProductReplenishs([]int64{pr.Id.Get()}, pr)
}

// UpdateProductReplenishs updates existing product.replenish records.
// All records (represented by ids) will be updated by pr values.
func (c *Client) UpdateProductReplenishs(ids []int64, pr *ProductReplenish) error {
	return c.Update(ProductReplenishModel, ids, pr)
}

// DeleteProductReplenish deletes an existing product.replenish record.
func (c *Client) DeleteProductReplenish(id int64) error {
	return c.DeleteProductReplenishs([]int64{id})
}

// DeleteProductReplenishs deletes existing product.replenish records.
func (c *Client) DeleteProductReplenishs(ids []int64) error {
	return c.Delete(ProductReplenishModel, ids)
}

// GetProductReplenish gets product.replenish existing record.
func (c *Client) GetProductReplenish(id int64) (*ProductReplenish, error) {
	prs, err := c.GetProductReplenishs([]int64{id})
	if err != nil {
		return nil, err
	}
	if prs != nil && len(*prs) > 0 {
		return &((*prs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.replenish not found", id)
}

// GetProductReplenishs gets product.replenish existing records.
func (c *Client) GetProductReplenishs(ids []int64) (*ProductReplenishs, error) {
	prs := &ProductReplenishs{}
	if err := c.Read(ProductReplenishModel, ids, nil, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindProductReplenish finds product.replenish record by querying it with criteria.
func (c *Client) FindProductReplenish(criteria *Criteria) (*ProductReplenish, error) {
	prs := &ProductReplenishs{}
	if err := c.SearchRead(ProductReplenishModel, criteria, NewOptions().Limit(1), prs); err != nil {
		return nil, err
	}
	if prs != nil && len(*prs) > 0 {
		return &((*prs)[0]), nil
	}
	return nil, fmt.Errorf("product.replenish was not found with criteria %v", criteria)
}

// FindProductReplenishs finds product.replenish records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductReplenishs(criteria *Criteria, options *Options) (*ProductReplenishs, error) {
	prs := &ProductReplenishs{}
	if err := c.SearchRead(ProductReplenishModel, criteria, options, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindProductReplenishIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductReplenishIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductReplenishModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductReplenishId finds record id by querying it with criteria.
func (c *Client) FindProductReplenishId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductReplenishModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.replenish was not found with criteria %v and options %v", criteria, options)
}
