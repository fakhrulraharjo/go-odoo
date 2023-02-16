package odoo

import (
	"fmt"
)

// ProductTemplate represents product.template model.
type ProductTemplate struct {
	LastUpdate                           *Time      `xmlrpc:"__last_update,omptempty"`
	AccountTagIds                        *Relation  `xmlrpc:"account_tag_ids,omptempty"`
	Active                               *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline                 *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration          *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon                *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                          *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                        *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                      *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon                     *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId                       *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                       *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AttributeLineIds                     *Relation  `xmlrpc:"attribute_line_ids,omptempty"`
	AvailableInPos                       *Bool      `xmlrpc:"available_in_pos,omptempty"`
	Barcode                              *String    `xmlrpc:"barcode,omptempty"`
	CanImage1024BeZoomed                 *Bool      `xmlrpc:"can_image_1024_be_zoomed,omptempty"`
	CategId                              *Many2One  `xmlrpc:"categ_id,omptempty"`
	Color                                *Int       `xmlrpc:"color,omptempty"`
	CompanyId                            *Many2One  `xmlrpc:"company_id,omptempty"`
	CostCurrencyId                       *Many2One  `xmlrpc:"cost_currency_id,omptempty"`
	CostMethod                           *Selection `xmlrpc:"cost_method,omptempty"`
	CreateDate                           *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                            *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                           *Many2One  `xmlrpc:"currency_id,omptempty"`
	DefaultCode                          *String    `xmlrpc:"default_code,omptempty"`
	Description                          *String    `xmlrpc:"description,omptempty"`
	DescriptionPicking                   *String    `xmlrpc:"description_picking,omptempty"`
	DescriptionPickingin                 *String    `xmlrpc:"description_pickingin,omptempty"`
	DescriptionPickingout                *String    `xmlrpc:"description_pickingout,omptempty"`
	DescriptionPurchase                  *String    `xmlrpc:"description_purchase,omptempty"`
	DescriptionSale                      *String    `xmlrpc:"description_sale,omptempty"`
	DetailedType                         *Selection `xmlrpc:"detailed_type,omptempty"`
	DisplayName                          *String    `xmlrpc:"display_name,omptempty"`
	HasAvailableRouteIds                 *Bool      `xmlrpc:"has_available_route_ids,omptempty"`
	HasConfigurableAttributes            *Bool      `xmlrpc:"has_configurable_attributes,omptempty"`
	HasMessage                           *Bool      `xmlrpc:"has_message,omptempty"`
	Id                                   *Int       `xmlrpc:"id,omptempty"`
	Image1024                            *String    `xmlrpc:"image_1024,omptempty"`
	Image128                             *String    `xmlrpc:"image_128,omptempty"`
	Image1920                            *String    `xmlrpc:"image_1920,omptempty"`
	Image256                             *String    `xmlrpc:"image_256,omptempty"`
	Image512                             *String    `xmlrpc:"image_512,omptempty"`
	IncomingQty                          *Float     `xmlrpc:"incoming_qty,omptempty"`
	IsProductVariant                     *Bool      `xmlrpc:"is_product_variant,omptempty"`
	ListPrice                            *Float     `xmlrpc:"list_price,omptempty"`
	LocationId                           *Many2One  `xmlrpc:"location_id,omptempty"`
	MessageAttachmentCount               *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds                   *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                      *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter               *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError                   *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                           *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                    *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId              *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction                    *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter             *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds                    *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MyActivityDateDeadline               *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                                 *String    `xmlrpc:"name,omptempty"`
	NbrMovesIn                           *Int       `xmlrpc:"nbr_moves_in,omptempty"`
	NbrMovesOut                          *Int       `xmlrpc:"nbr_moves_out,omptempty"`
	NbrReorderingRules                   *Int       `xmlrpc:"nbr_reordering_rules,omptempty"`
	OutgoingQty                          *Float     `xmlrpc:"outgoing_qty,omptempty"`
	PackagingIds                         *Relation  `xmlrpc:"packaging_ids,omptempty"`
	PosCategId                           *Many2One  `xmlrpc:"pos_categ_id,omptempty"`
	PricelistItemCount                   *Int       `xmlrpc:"pricelist_item_count,omptempty"`
	Priority                             *Selection `xmlrpc:"priority,omptempty"`
	ProductTagIds                        *Relation  `xmlrpc:"product_tag_ids,omptempty"`
	ProductTooltip                       *String    `xmlrpc:"product_tooltip,omptempty"`
	ProductVariantCount                  *Int       `xmlrpc:"product_variant_count,omptempty"`
	ProductVariantId                     *Many2One  `xmlrpc:"product_variant_id,omptempty"`
	ProductVariantIds                    *Relation  `xmlrpc:"product_variant_ids,omptempty"`
	PropertyAccountExpenseId             *Many2One  `xmlrpc:"property_account_expense_id,omptempty"`
	PropertyAccountIncomeId              *Many2One  `xmlrpc:"property_account_income_id,omptempty"`
	PropertyStockInventory               *Many2One  `xmlrpc:"property_stock_inventory,omptempty"`
	PropertyStockProduction              *Many2One  `xmlrpc:"property_stock_production,omptempty"`
	PurchaseOk                           *Bool      `xmlrpc:"purchase_ok,omptempty"`
	QtyAvailable                         *Float     `xmlrpc:"qty_available,omptempty"`
	ReorderingMaxQty                     *Float     `xmlrpc:"reordering_max_qty,omptempty"`
	ReorderingMinQty                     *Float     `xmlrpc:"reordering_min_qty,omptempty"`
	ResponsibleId                        *Many2One  `xmlrpc:"responsible_id,omptempty"`
	RouteFromCategIds                    *Relation  `xmlrpc:"route_from_categ_ids,omptempty"`
	RouteIds                             *Relation  `xmlrpc:"route_ids,omptempty"`
	SaleDelay                            *Float     `xmlrpc:"sale_delay,omptempty"`
	SaleOk                               *Bool      `xmlrpc:"sale_ok,omptempty"`
	SellerIds                            *Relation  `xmlrpc:"seller_ids,omptempty"`
	Sequence                             *Int       `xmlrpc:"sequence,omptempty"`
	ShowForecastedQtyStatusButton        *Bool      `xmlrpc:"show_forecasted_qty_status_button,omptempty"`
	ShowOnHandQtyStatusButton            *Bool      `xmlrpc:"show_on_hand_qty_status_button,omptempty"`
	StandardPrice                        *Float     `xmlrpc:"standard_price,omptempty"`
	SupplierTaxesId                      *Relation  `xmlrpc:"supplier_taxes_id,omptempty"`
	TaxString                            *String    `xmlrpc:"tax_string,omptempty"`
	TaxesId                              *Relation  `xmlrpc:"taxes_id,omptempty"`
	ToWeight                             *Bool      `xmlrpc:"to_weight,omptempty"`
	Tracking                             *Selection `xmlrpc:"tracking,omptempty"`
	Type                                 *Selection `xmlrpc:"type,omptempty"`
	UomId                                *Many2One  `xmlrpc:"uom_id,omptempty"`
	UomName                              *String    `xmlrpc:"uom_name,omptempty"`
	UomPoId                              *Many2One  `xmlrpc:"uom_po_id,omptempty"`
	ValidProductTemplateAttributeLineIds *Relation  `xmlrpc:"valid_product_template_attribute_line_ids,omptempty"`
	Valuation                            *Selection `xmlrpc:"valuation,omptempty"`
	VariantSellerIds                     *Relation  `xmlrpc:"variant_seller_ids,omptempty"`
	VirtualAvailable                     *Float     `xmlrpc:"virtual_available,omptempty"`
	Volume                               *Float     `xmlrpc:"volume,omptempty"`
	VolumeUomName                        *String    `xmlrpc:"volume_uom_name,omptempty"`
	WarehouseId                          *Many2One  `xmlrpc:"warehouse_id,omptempty"`
	WebsiteMessageIds                    *Relation  `xmlrpc:"website_message_ids,omptempty"`
	Weight                               *Float     `xmlrpc:"weight,omptempty"`
	WeightUomName                        *String    `xmlrpc:"weight_uom_name,omptempty"`
	WriteDate                            *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                             *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ProductTemplates represents array of product.template model.
type ProductTemplates []ProductTemplate

// ProductTemplateModel is the odoo model name.
const ProductTemplateModel = "product.template"

// Many2One convert ProductTemplate to *Many2One.
func (pt *ProductTemplate) Many2One() *Many2One {
	return NewMany2One(pt.Id.Get(), "")
}

// CreateProductTemplate creates a new product.template model and returns its id.
func (c *Client) CreateProductTemplate(pt *ProductTemplate) (int64, error) {
	return c.Create(ProductTemplateModel, pt)
}

// UpdateProductTemplate updates an existing product.template record.
func (c *Client) UpdateProductTemplate(pt *ProductTemplate) error {
	return c.UpdateProductTemplates([]int64{pt.Id.Get()}, pt)
}

// UpdateProductTemplates updates existing product.template records.
// All records (represented by ids) will be updated by pt values.
func (c *Client) UpdateProductTemplates(ids []int64, pt *ProductTemplate) error {
	return c.Update(ProductTemplateModel, ids, pt)
}

// DeleteProductTemplate deletes an existing product.template record.
func (c *Client) DeleteProductTemplate(id int64) error {
	return c.DeleteProductTemplates([]int64{id})
}

// DeleteProductTemplates deletes existing product.template records.
func (c *Client) DeleteProductTemplates(ids []int64) error {
	return c.Delete(ProductTemplateModel, ids)
}

// GetProductTemplate gets product.template existing record.
func (c *Client) GetProductTemplate(id int64) (*ProductTemplate, error) {
	pts, err := c.GetProductTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.template not found", id)
}

// GetProductTemplates gets product.template existing records.
func (c *Client) GetProductTemplates(ids []int64) (*ProductTemplates, error) {
	pts := &ProductTemplates{}
	if err := c.Read(ProductTemplateModel, ids, nil, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProductTemplate finds product.template record by querying it with criteria.
func (c *Client) FindProductTemplate(criteria *Criteria) (*ProductTemplate, error) {
	pts := &ProductTemplates{}
	if err := c.SearchRead(ProductTemplateModel, criteria, NewOptions().Limit(1), pts); err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("product.template was not found with criteria %v", criteria)
}

// FindProductTemplates finds product.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductTemplates(criteria *Criteria, options *Options) (*ProductTemplates, error) {
	pts := &ProductTemplates{}
	if err := c.SearchRead(ProductTemplateModel, criteria, options, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProductTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductTemplateId finds record id by querying it with criteria.
func (c *Client) FindProductTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.template was not found with criteria %v and options %v", criteria, options)
}
