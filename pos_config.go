package odoo

import (
	"fmt"
)

// PosConfig represents pos.config model.
type PosConfig struct {
	LastUpdate                          *Time      `xmlrpc:"__last_update,omptempty"`
	Active                              *Bool      `xmlrpc:"active,omptempty"`
	AmountAuthorizedDiff                *Float     `xmlrpc:"amount_authorized_diff,omptempty"`
	AvailablePricelistIds               *Relation  `xmlrpc:"available_pricelist_ids,omptempty"`
	CashControl                         *Bool      `xmlrpc:"cash_control,omptempty"`
	CashRounding                        *Bool      `xmlrpc:"cash_rounding,omptempty"`
	CompanyHasTemplate                  *Bool      `xmlrpc:"company_has_template,omptempty"`
	CompanyId                           *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                          *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                           *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                          *Many2One  `xmlrpc:"currency_id,omptempty"`
	CurrentSessionId                    *Many2One  `xmlrpc:"current_session_id,omptempty"`
	CurrentSessionState                 *String    `xmlrpc:"current_session_state,omptempty"`
	CurrentUserId                       *Many2One  `xmlrpc:"current_user_id,omptempty"`
	DefaultBillIds                      *Relation  `xmlrpc:"default_bill_ids,omptempty"`
	DefaultFiscalPositionId             *Many2One  `xmlrpc:"default_fiscal_position_id,omptempty"`
	DisplayName                         *String    `xmlrpc:"display_name,omptempty"`
	EpsonPrinterIp                      *String    `xmlrpc:"epson_printer_ip,omptempty"`
	FiscalPositionIds                   *Relation  `xmlrpc:"fiscal_position_ids,omptempty"`
	GroupPosManagerId                   *Many2One  `xmlrpc:"group_pos_manager_id,omptempty"`
	GroupPosUserId                      *Many2One  `xmlrpc:"group_pos_user_id,omptempty"`
	HasActiveSession                    *Bool      `xmlrpc:"has_active_session,omptempty"`
	Id                                  *Int       `xmlrpc:"id,omptempty"`
	IfaceAvailableCategIds              *Relation  `xmlrpc:"iface_available_categ_ids,omptempty"`
	IfaceBigScrollbars                  *Bool      `xmlrpc:"iface_big_scrollbars,omptempty"`
	IfaceCashdrawer                     *Bool      `xmlrpc:"iface_cashdrawer,omptempty"`
	IfaceCustomerFacingDisplay          *Bool      `xmlrpc:"iface_customer_facing_display,omptempty"`
	IfaceCustomerFacingDisplayLocal     *Bool      `xmlrpc:"iface_customer_facing_display_local,omptempty"`
	IfaceCustomerFacingDisplayViaProxy  *Bool      `xmlrpc:"iface_customer_facing_display_via_proxy,omptempty"`
	IfaceElectronicScale                *Bool      `xmlrpc:"iface_electronic_scale,omptempty"`
	IfacePrintAuto                      *Bool      `xmlrpc:"iface_print_auto,omptempty"`
	IfacePrintSkipScreen                *Bool      `xmlrpc:"iface_print_skip_screen,omptempty"`
	IfacePrintViaProxy                  *Bool      `xmlrpc:"iface_print_via_proxy,omptempty"`
	IfaceScanViaProxy                   *Bool      `xmlrpc:"iface_scan_via_proxy,omptempty"`
	IfaceStartCategId                   *Many2One  `xmlrpc:"iface_start_categ_id,omptempty"`
	IfaceTaxIncluded                    *Selection `xmlrpc:"iface_tax_included,omptempty"`
	IfaceTipproduct                     *Bool      `xmlrpc:"iface_tipproduct,omptempty"`
	InvoiceJournalId                    *Many2One  `xmlrpc:"invoice_journal_id,omptempty"`
	IsHeaderOrFooter                    *Bool      `xmlrpc:"is_header_or_footer,omptempty"`
	IsInstalledAccountAccountant        *Bool      `xmlrpc:"is_installed_account_accountant,omptempty"`
	IsMarginsCostsAccessibleToEveryUser *Bool      `xmlrpc:"is_margins_costs_accessible_to_every_user,omptempty"`
	IsPosbox                            *Bool      `xmlrpc:"is_posbox,omptempty"`
	JournalId                           *Many2One  `xmlrpc:"journal_id,omptempty"`
	LastSessionClosingCash              *Float     `xmlrpc:"last_session_closing_cash,omptempty"`
	LastSessionClosingDate              *Time      `xmlrpc:"last_session_closing_date,omptempty"`
	LimitCategories                     *Bool      `xmlrpc:"limit_categories,omptempty"`
	LimitedPartnersAmount               *Int       `xmlrpc:"limited_partners_amount,omptempty"`
	LimitedPartnersLoading              *Bool      `xmlrpc:"limited_partners_loading,omptempty"`
	LimitedProductsAmount               *Int       `xmlrpc:"limited_products_amount,omptempty"`
	LimitedProductsLoading              *Bool      `xmlrpc:"limited_products_loading,omptempty"`
	ManualDiscount                      *Bool      `xmlrpc:"manual_discount,omptempty"`
	ModulePosDiscount                   *Bool      `xmlrpc:"module_pos_discount,omptempty"`
	ModulePosHr                         *Bool      `xmlrpc:"module_pos_hr,omptempty"`
	ModulePosMercury                    *Bool      `xmlrpc:"module_pos_mercury,omptempty"`
	ModulePosRestaurant                 *Bool      `xmlrpc:"module_pos_restaurant,omptempty"`
	Name                                *String    `xmlrpc:"name,omptempty"`
	NumberOfOpenedSession               *Int       `xmlrpc:"number_of_opened_session,omptempty"`
	OnlyRoundCashMethod                 *Bool      `xmlrpc:"only_round_cash_method,omptempty"`
	OtherDevices                        *Bool      `xmlrpc:"other_devices,omptempty"`
	PartnerLoadBackground               *Bool      `xmlrpc:"partner_load_background,omptempty"`
	PaymentMethodIds                    *Relation  `xmlrpc:"payment_method_ids,omptempty"`
	PickingPolicy                       *Selection `xmlrpc:"picking_policy,omptempty"`
	PickingTypeId                       *Many2One  `xmlrpc:"picking_type_id,omptempty"`
	PosSessionDuration                  *String    `xmlrpc:"pos_session_duration,omptempty"`
	PosSessionState                     *String    `xmlrpc:"pos_session_state,omptempty"`
	PosSessionUsername                  *String    `xmlrpc:"pos_session_username,omptempty"`
	PricelistId                         *Many2One  `xmlrpc:"pricelist_id,omptempty"`
	ProductLoadBackground               *Bool      `xmlrpc:"product_load_background,omptempty"`
	ProxyIp                             *String    `xmlrpc:"proxy_ip,omptempty"`
	ReceiptFooter                       *String    `xmlrpc:"receipt_footer,omptempty"`
	ReceiptHeader                       *String    `xmlrpc:"receipt_header,omptempty"`
	RestrictPriceControl                *Bool      `xmlrpc:"restrict_price_control,omptempty"`
	RoundingMethod                      *Many2One  `xmlrpc:"rounding_method,omptempty"`
	RouteId                             *Many2One  `xmlrpc:"route_id,omptempty"`
	SequenceId                          *Many2One  `xmlrpc:"sequence_id,omptempty"`
	SequenceLineId                      *Many2One  `xmlrpc:"sequence_line_id,omptempty"`
	SessionIds                          *Relation  `xmlrpc:"session_ids,omptempty"`
	SetMaximumDifference                *Bool      `xmlrpc:"set_maximum_difference,omptempty"`
	ShipLater                           *Bool      `xmlrpc:"ship_later,omptempty"`
	StartCategory                       *Bool      `xmlrpc:"start_category,omptempty"`
	TaxRegimeSelection                  *Bool      `xmlrpc:"tax_regime_selection,omptempty"`
	TipProductId                        *Many2One  `xmlrpc:"tip_product_id,omptempty"`
	UsePricelist                        *Bool      `xmlrpc:"use_pricelist,omptempty"`
	Uuid                                *String    `xmlrpc:"uuid,omptempty"`
	WarehouseId                         *Many2One  `xmlrpc:"warehouse_id,omptempty"`
	WriteDate                           *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                            *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// PosConfigs represents array of pos.config model.
type PosConfigs []PosConfig

// PosConfigModel is the odoo model name.
const PosConfigModel = "pos.config"

// Many2One convert PosConfig to *Many2One.
func (pc *PosConfig) Many2One() *Many2One {
	return NewMany2One(pc.Id.Get(), "")
}

// CreatePosConfig creates a new pos.config model and returns its id.
func (c *Client) CreatePosConfig(pc *PosConfig) (int64, error) {
	return c.Create(PosConfigModel, pc)
}

// UpdatePosConfig updates an existing pos.config record.
func (c *Client) UpdatePosConfig(pc *PosConfig) error {
	return c.UpdatePosConfigs([]int64{pc.Id.Get()}, pc)
}

// UpdatePosConfigs updates existing pos.config records.
// All records (represented by ids) will be updated by pc values.
func (c *Client) UpdatePosConfigs(ids []int64, pc *PosConfig) error {
	return c.Update(PosConfigModel, ids, pc)
}

// DeletePosConfig deletes an existing pos.config record.
func (c *Client) DeletePosConfig(id int64) error {
	return c.DeletePosConfigs([]int64{id})
}

// DeletePosConfigs deletes existing pos.config records.
func (c *Client) DeletePosConfigs(ids []int64) error {
	return c.Delete(PosConfigModel, ids)
}

// GetPosConfig gets pos.config existing record.
func (c *Client) GetPosConfig(id int64) (*PosConfig, error) {
	pcs, err := c.GetPosConfigs([]int64{id})
	if err != nil {
		return nil, err
	}
	if pcs != nil && len(*pcs) > 0 {
		return &((*pcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of pos.config not found", id)
}

// GetPosConfigs gets pos.config existing records.
func (c *Client) GetPosConfigs(ids []int64) (*PosConfigs, error) {
	pcs := &PosConfigs{}
	if err := c.Read(PosConfigModel, ids, nil, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindPosConfig finds pos.config record by querying it with criteria.
func (c *Client) FindPosConfig(criteria *Criteria) (*PosConfig, error) {
	pcs := &PosConfigs{}
	if err := c.SearchRead(PosConfigModel, criteria, NewOptions().Limit(1), pcs); err != nil {
		return nil, err
	}
	if pcs != nil && len(*pcs) > 0 {
		return &((*pcs)[0]), nil
	}
	return nil, fmt.Errorf("pos.config was not found with criteria %v", criteria)
}

// FindPosConfigs finds pos.config records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosConfigs(criteria *Criteria, options *Options) (*PosConfigs, error) {
	pcs := &PosConfigs{}
	if err := c.SearchRead(PosConfigModel, criteria, options, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindPosConfigIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosConfigIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PosConfigModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPosConfigId finds record id by querying it with criteria.
func (c *Client) FindPosConfigId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosConfigModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("pos.config was not found with criteria %v and options %v", criteria, options)
}
