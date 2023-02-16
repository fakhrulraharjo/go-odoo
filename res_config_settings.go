package odoo

import (
	"fmt"
)

// ResConfigSettings represents res.config.settings model.
type ResConfigSettings struct {
	LastUpdate                                  *Time      `xmlrpc:"__last_update,omptempty"`
	AccountCashBasisBaseAccountId               *Many2One  `xmlrpc:"account_cash_basis_base_account_id,omptempty"`
	AccountDefaultCreditLimit                   *Float     `xmlrpc:"account_default_credit_limit,omptempty"`
	AccountDefaultPosReceivableAccountId        *Many2One  `xmlrpc:"account_default_pos_receivable_account_id,omptempty"`
	AccountFiscalCountryId                      *Many2One  `xmlrpc:"account_fiscal_country_id,omptempty"`
	AccountJournalEarlyPayDiscountGainAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_gain_account_id,omptempty"`
	AccountJournalEarlyPayDiscountLossAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_loss_account_id,omptempty"`
	AccountJournalPaymentCreditAccountId        *Many2One  `xmlrpc:"account_journal_payment_credit_account_id,omptempty"`
	AccountJournalPaymentDebitAccountId         *Many2One  `xmlrpc:"account_journal_payment_debit_account_id,omptempty"`
	AccountJournalSuspenseAccountId             *Many2One  `xmlrpc:"account_journal_suspense_account_id,omptempty"`
	AccountStorno                               *Bool      `xmlrpc:"account_storno,omptempty"`
	AccountUseCreditLimit                       *Bool      `xmlrpc:"account_use_credit_limit,omptempty"`
	ActiveUserCount                             *Int       `xmlrpc:"active_user_count,omptempty"`
	AliasDomain                                 *String    `xmlrpc:"alias_domain,omptempty"`
	AnnualInventoryDay                          *Int       `xmlrpc:"annual_inventory_day,omptempty"`
	AnnualInventoryMonth                        *Selection `xmlrpc:"annual_inventory_month,omptempty"`
	AuthSignupResetPassword                     *Bool      `xmlrpc:"auth_signup_reset_password,omptempty"`
	AuthSignupTemplateUserId                    *Many2One  `xmlrpc:"auth_signup_template_user_id,omptempty"`
	AuthSignupUninvited                         *Selection `xmlrpc:"auth_signup_uninvited,omptempty"`
	BarcodeNomenclatureId                       *Many2One  `xmlrpc:"barcode_nomenclature_id,omptempty"`
	ChartTemplateId                             *Many2One  `xmlrpc:"chart_template_id,omptempty"`
	CompanyCount                                *Int       `xmlrpc:"company_count,omptempty"`
	CompanyId                                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CompanyInformations                         *String    `xmlrpc:"company_informations,omptempty"`
	CompanyName                                 *String    `xmlrpc:"company_name,omptempty"`
	CountryCode                                 *String    `xmlrpc:"country_code,omptempty"`
	CreateDate                                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyExchangeJournalId                   *Many2One  `xmlrpc:"currency_exchange_journal_id,omptempty"`
	CurrencyId                                  *Many2One  `xmlrpc:"currency_id,omptempty"`
	DigestEmails                                *Bool      `xmlrpc:"digest_emails,omptempty"`
	DigestId                                    *Many2One  `xmlrpc:"digest_id,omptempty"`
	DisplayName                                 *String    `xmlrpc:"display_name,omptempty"`
	EarlyPayDiscountComputation                 *Selection `xmlrpc:"early_pay_discount_computation,omptempty"`
	ExpenseCurrencyExchangeAccountId            *Many2One  `xmlrpc:"expense_currency_exchange_account_id,omptempty"`
	ExternalEmailServerDefault                  *Bool      `xmlrpc:"external_email_server_default,omptempty"`
	ExternalReportLayoutId                      *Many2One  `xmlrpc:"external_report_layout_id,omptempty"`
	FailCounter                                 *Int       `xmlrpc:"fail_counter,omptempty"`
	GoogleGmailClientIdentifier                 *String    `xmlrpc:"google_gmail_client_identifier,omptempty"`
	GoogleGmailClientSecret                     *String    `xmlrpc:"google_gmail_client_secret,omptempty"`
	GroupAnalyticAccounting                     *Bool      `xmlrpc:"group_analytic_accounting,omptempty"`
	GroupCashRounding                           *Bool      `xmlrpc:"group_cash_rounding,omptempty"`
	GroupDiscountPerSoLine                      *Bool      `xmlrpc:"group_discount_per_so_line,omptempty"`
	GroupLotOnDeliverySlip                      *Bool      `xmlrpc:"group_lot_on_delivery_slip,omptempty"`
	GroupLotOnInvoice                           *Bool      `xmlrpc:"group_lot_on_invoice,omptempty"`
	GroupMultiCurrency                          *Bool      `xmlrpc:"group_multi_currency,omptempty"`
	GroupProductPricelist                       *Bool      `xmlrpc:"group_product_pricelist,omptempty"`
	GroupProductVariant                         *Bool      `xmlrpc:"group_product_variant,omptempty"`
	GroupSaleDeliveryAddress                    *Bool      `xmlrpc:"group_sale_delivery_address,omptempty"`
	GroupSalePricelist                          *Bool      `xmlrpc:"group_sale_pricelist,omptempty"`
	GroupShowLineSubtotalsTaxExcluded           *Bool      `xmlrpc:"group_show_line_subtotals_tax_excluded,omptempty"`
	GroupShowLineSubtotalsTaxIncluded           *Bool      `xmlrpc:"group_show_line_subtotals_tax_included,omptempty"`
	GroupShowPurchaseReceipts                   *Bool      `xmlrpc:"group_show_purchase_receipts,omptempty"`
	GroupShowSaleReceipts                       *Bool      `xmlrpc:"group_show_sale_receipts,omptempty"`
	GroupStockAdvLocation                       *Bool      `xmlrpc:"group_stock_adv_location,omptempty"`
	GroupStockLotPrintGs1                       *Bool      `xmlrpc:"group_stock_lot_print_gs1,omptempty"`
	GroupStockMultiLocations                    *Bool      `xmlrpc:"group_stock_multi_locations,omptempty"`
	GroupStockPackaging                         *Bool      `xmlrpc:"group_stock_packaging,omptempty"`
	GroupStockPickingWave                       *Bool      `xmlrpc:"group_stock_picking_wave,omptempty"`
	GroupStockProductionLot                     *Bool      `xmlrpc:"group_stock_production_lot,omptempty"`
	GroupStockReceptionReport                   *Bool      `xmlrpc:"group_stock_reception_report,omptempty"`
	GroupStockSignDelivery                      *Bool      `xmlrpc:"group_stock_sign_delivery,omptempty"`
	GroupStockStorageCategories                 *Bool      `xmlrpc:"group_stock_storage_categories,omptempty"`
	GroupStockTrackingLot                       *Bool      `xmlrpc:"group_stock_tracking_lot,omptempty"`
	GroupStockTrackingOwner                     *Bool      `xmlrpc:"group_stock_tracking_owner,omptempty"`
	GroupUom                                    *Bool      `xmlrpc:"group_uom,omptempty"`
	GroupWarningAccount                         *Bool      `xmlrpc:"group_warning_account,omptempty"`
	GroupWarningStock                           *Bool      `xmlrpc:"group_warning_stock,omptempty"`
	HasAccountingEntries                        *Bool      `xmlrpc:"has_accounting_entries,omptempty"`
	HasChartOfAccounts                          *Bool      `xmlrpc:"has_chart_of_accounts,omptempty"`
	Id                                          *Int       `xmlrpc:"id,omptempty"`
	IncomeCurrencyExchangeAccountId             *Many2One  `xmlrpc:"income_currency_exchange_account_id,omptempty"`
	IncotermId                                  *Many2One  `xmlrpc:"incoterm_id,omptempty"`
	InvoiceIsEmail                              *Bool      `xmlrpc:"invoice_is_email,omptempty"`
	InvoiceIsPrint                              *Bool      `xmlrpc:"invoice_is_print,omptempty"`
	InvoiceIsSnailmail                          *Bool      `xmlrpc:"invoice_is_snailmail,omptempty"`
	InvoiceTerms                                *String    `xmlrpc:"invoice_terms,omptempty"`
	InvoiceTermsHtml                            *String    `xmlrpc:"invoice_terms_html,omptempty"`
	IsDefaultPricelistDisplayed                 *Bool      `xmlrpc:"is_default_pricelist_displayed,omptempty"`
	L10NIdTaxAddress                            *String    `xmlrpc:"l10n_id_tax_address,omptempty"`
	L10NIdTaxName                               *String    `xmlrpc:"l10n_id_tax_name,omptempty"`
	LanguageCount                               *Int       `xmlrpc:"language_count,omptempty"`
	ModuleAccountAccountant                     *Bool      `xmlrpc:"module_account_accountant,omptempty"`
	ModuleAccountBankStatementImportCamt        *Bool      `xmlrpc:"module_account_bank_statement_import_camt,omptempty"`
	ModuleAccountBankStatementImportCsv         *Bool      `xmlrpc:"module_account_bank_statement_import_csv,omptempty"`
	ModuleAccountBankStatementImportOfx         *Bool      `xmlrpc:"module_account_bank_statement_import_ofx,omptempty"`
	ModuleAccountBankStatementImportQif         *Bool      `xmlrpc:"module_account_bank_statement_import_qif,omptempty"`
	ModuleAccountBatchPayment                   *Bool      `xmlrpc:"module_account_batch_payment,omptempty"`
	ModuleAccountBudget                         *Bool      `xmlrpc:"module_account_budget,omptempty"`
	ModuleAccountCheckPrinting                  *Bool      `xmlrpc:"module_account_check_printing,omptempty"`
	ModuleAccountInterCompanyRules              *Bool      `xmlrpc:"module_account_inter_company_rules,omptempty"`
	ModuleAccountIntrastat                      *Bool      `xmlrpc:"module_account_intrastat,omptempty"`
	ModuleAccountInvoiceExtract                 *Bool      `xmlrpc:"module_account_invoice_extract,omptempty"`
	ModuleAccountPayment                        *Bool      `xmlrpc:"module_account_payment,omptempty"`
	ModuleAccountReports                        *Bool      `xmlrpc:"module_account_reports,omptempty"`
	ModuleAccountSepa                           *Bool      `xmlrpc:"module_account_sepa,omptempty"`
	ModuleAccountSepaDirectDebit                *Bool      `xmlrpc:"module_account_sepa_direct_debit,omptempty"`
	ModuleAccountTaxcloud                       *Bool      `xmlrpc:"module_account_taxcloud,omptempty"`
	ModuleAuthLdap                              *Bool      `xmlrpc:"module_auth_ldap,omptempty"`
	ModuleAuthOauth                             *Bool      `xmlrpc:"module_auth_oauth,omptempty"`
	ModuleBaseGengo                             *Bool      `xmlrpc:"module_base_gengo,omptempty"`
	ModuleBaseGeolocalize                       *Bool      `xmlrpc:"module_base_geolocalize,omptempty"`
	ModuleBaseImport                            *Bool      `xmlrpc:"module_base_import,omptempty"`
	ModuleCurrencyRateLive                      *Bool      `xmlrpc:"module_currency_rate_live,omptempty"`
	ModuleDelivery                              *Bool      `xmlrpc:"module_delivery,omptempty"`
	ModuleDeliveryBpost                         *Bool      `xmlrpc:"module_delivery_bpost,omptempty"`
	ModuleDeliveryDhl                           *Bool      `xmlrpc:"module_delivery_dhl,omptempty"`
	ModuleDeliveryEasypost                      *Bool      `xmlrpc:"module_delivery_easypost,omptempty"`
	ModuleDeliveryFedex                         *Bool      `xmlrpc:"module_delivery_fedex,omptempty"`
	ModuleDeliverySendcloud                     *Bool      `xmlrpc:"module_delivery_sendcloud,omptempty"`
	ModuleDeliveryUps                           *Bool      `xmlrpc:"module_delivery_ups,omptempty"`
	ModuleDeliveryUsps                          *Bool      `xmlrpc:"module_delivery_usps,omptempty"`
	ModuleGoogleCalendar                        *Bool      `xmlrpc:"module_google_calendar,omptempty"`
	ModuleGoogleGmail                           *Bool      `xmlrpc:"module_google_gmail,omptempty"`
	ModuleGoogleRecaptcha                       *Bool      `xmlrpc:"module_google_recaptcha,omptempty"`
	ModuleL10NEuOss                             *Bool      `xmlrpc:"module_l10n_eu_oss,omptempty"`
	ModuleLoyalty                               *Bool      `xmlrpc:"module_loyalty,omptempty"`
	ModuleMailPlugin                            *Bool      `xmlrpc:"module_mail_plugin,omptempty"`
	ModuleMicrosoftCalendar                     *Bool      `xmlrpc:"module_microsoft_calendar,omptempty"`
	ModuleMicrosoftOutlook                      *Bool      `xmlrpc:"module_microsoft_outlook,omptempty"`
	ModulePartnerAutocomplete                   *Bool      `xmlrpc:"module_partner_autocomplete,omptempty"`
	ModulePosAdyen                              *Bool      `xmlrpc:"module_pos_adyen,omptempty"`
	ModulePosMercury                            *Bool      `xmlrpc:"module_pos_mercury,omptempty"`
	ModulePosSix                                *Bool      `xmlrpc:"module_pos_six,omptempty"`
	ModulePosStripe                             *Bool      `xmlrpc:"module_pos_stripe,omptempty"`
	ModuleProductExpiry                         *Bool      `xmlrpc:"module_product_expiry,omptempty"`
	ModuleProductImages                         *Bool      `xmlrpc:"module_product_images,omptempty"`
	ModuleProductMargin                         *Bool      `xmlrpc:"module_product_margin,omptempty"`
	ModuleQualityControl                        *Bool      `xmlrpc:"module_quality_control,omptempty"`
	ModuleQualityControlWorksheet               *Bool      `xmlrpc:"module_quality_control_worksheet,omptempty"`
	ModuleSaleProductMatrix                     *Bool      `xmlrpc:"module_sale_product_matrix,omptempty"`
	ModuleSnailmailAccount                      *Bool      `xmlrpc:"module_snailmail_account,omptempty"`
	ModuleStockBarcode                          *Bool      `xmlrpc:"module_stock_barcode,omptempty"`
	ModuleStockLandedCosts                      *Bool      `xmlrpc:"module_stock_landed_costs,omptempty"`
	ModuleStockPickingBatch                     *Bool      `xmlrpc:"module_stock_picking_batch,omptempty"`
	ModuleStockSms                              *Bool      `xmlrpc:"module_stock_sms,omptempty"`
	ModuleVoip                                  *Bool      `xmlrpc:"module_voip,omptempty"`
	ModuleWebUnsplash                           *Bool      `xmlrpc:"module_web_unsplash,omptempty"`
	PartnerAutocompleteInsufficientCredit       *Bool      `xmlrpc:"partner_autocomplete_insufficient_credit,omptempty"`
	PointOfSaleUseTicketQrCode                  *Bool      `xmlrpc:"point_of_sale_use_ticket_qr_code,omptempty"`
	PortalAllowApiKeys                          *Bool      `xmlrpc:"portal_allow_api_keys,omptempty"`
	PosAllowedPricelistIds                      *Relation  `xmlrpc:"pos_allowed_pricelist_ids,omptempty"`
	PosAmountAuthorizedDiff                     *Float     `xmlrpc:"pos_amount_authorized_diff,omptempty"`
	PosAvailablePricelistIds                    *Relation  `xmlrpc:"pos_available_pricelist_ids,omptempty"`
	PosCashControl                              *Bool      `xmlrpc:"pos_cash_control,omptempty"`
	PosCashRounding                             *Bool      `xmlrpc:"pos_cash_rounding,omptempty"`
	PosCompanyHasTemplate                       *Bool      `xmlrpc:"pos_company_has_template,omptempty"`
	PosConfigId                                 *Many2One  `xmlrpc:"pos_config_id,omptempty"`
	PosDefaultBillIds                           *Relation  `xmlrpc:"pos_default_bill_ids,omptempty"`
	PosDefaultFiscalPositionId                  *Many2One  `xmlrpc:"pos_default_fiscal_position_id,omptempty"`
	PosEpsonPrinterIp                           *String    `xmlrpc:"pos_epson_printer_ip,omptempty"`
	PosFiscalPositionIds                        *Relation  `xmlrpc:"pos_fiscal_position_ids,omptempty"`
	PosHasActiveSession                         *Bool      `xmlrpc:"pos_has_active_session,omptempty"`
	PosIfaceAvailableCategIds                   *Relation  `xmlrpc:"pos_iface_available_categ_ids,omptempty"`
	PosIfaceBigScrollbars                       *Bool      `xmlrpc:"pos_iface_big_scrollbars,omptempty"`
	PosIfaceCashdrawer                          *Bool      `xmlrpc:"pos_iface_cashdrawer,omptempty"`
	PosIfaceCustomerFacingDisplayLocal          *Bool      `xmlrpc:"pos_iface_customer_facing_display_local,omptempty"`
	PosIfaceCustomerFacingDisplayViaProxy       *Bool      `xmlrpc:"pos_iface_customer_facing_display_via_proxy,omptempty"`
	PosIfaceElectronicScale                     *Bool      `xmlrpc:"pos_iface_electronic_scale,omptempty"`
	PosIfacePrintAuto                           *Bool      `xmlrpc:"pos_iface_print_auto,omptempty"`
	PosIfacePrintSkipScreen                     *Bool      `xmlrpc:"pos_iface_print_skip_screen,omptempty"`
	PosIfacePrintViaProxy                       *Bool      `xmlrpc:"pos_iface_print_via_proxy,omptempty"`
	PosIfaceScanViaProxy                        *Bool      `xmlrpc:"pos_iface_scan_via_proxy,omptempty"`
	PosIfaceStartCategId                        *Many2One  `xmlrpc:"pos_iface_start_categ_id,omptempty"`
	PosIfaceTaxIncluded                         *Selection `xmlrpc:"pos_iface_tax_included,omptempty"`
	PosIfaceTipproduct                          *Bool      `xmlrpc:"pos_iface_tipproduct,omptempty"`
	PosInvoiceJournalId                         *Many2One  `xmlrpc:"pos_invoice_journal_id,omptempty"`
	PosIsHeaderOrFooter                         *Bool      `xmlrpc:"pos_is_header_or_footer,omptempty"`
	PosIsMarginsCostsAccessibleToEveryUser      *Bool      `xmlrpc:"pos_is_margins_costs_accessible_to_every_user,omptempty"`
	PosIsPosbox                                 *Bool      `xmlrpc:"pos_is_posbox,omptempty"`
	PosJournalId                                *Many2One  `xmlrpc:"pos_journal_id,omptempty"`
	PosLimitCategories                          *Bool      `xmlrpc:"pos_limit_categories,omptempty"`
	PosLimitedPartnersAmount                    *Int       `xmlrpc:"pos_limited_partners_amount,omptempty"`
	PosLimitedPartnersLoading                   *Bool      `xmlrpc:"pos_limited_partners_loading,omptempty"`
	PosLimitedProductsAmount                    *Int       `xmlrpc:"pos_limited_products_amount,omptempty"`
	PosLimitedProductsLoading                   *Bool      `xmlrpc:"pos_limited_products_loading,omptempty"`
	PosManualDiscount                           *Bool      `xmlrpc:"pos_manual_discount,omptempty"`
	PosModulePosDiscount                        *Bool      `xmlrpc:"pos_module_pos_discount,omptempty"`
	PosModulePosHr                              *Bool      `xmlrpc:"pos_module_pos_hr,omptempty"`
	PosModulePosRestaurant                      *Bool      `xmlrpc:"pos_module_pos_restaurant,omptempty"`
	PosOnlyRoundCashMethod                      *Bool      `xmlrpc:"pos_only_round_cash_method,omptempty"`
	PosOtherDevices                             *Bool      `xmlrpc:"pos_other_devices,omptempty"`
	PosPartnerLoadBackground                    *Bool      `xmlrpc:"pos_partner_load_background,omptempty"`
	PosPaymentMethodIds                         *Relation  `xmlrpc:"pos_payment_method_ids,omptempty"`
	PosPickingPolicy                            *Selection `xmlrpc:"pos_picking_policy,omptempty"`
	PosPickingTypeId                            *Many2One  `xmlrpc:"pos_picking_type_id,omptempty"`
	PosPricelistId                              *Many2One  `xmlrpc:"pos_pricelist_id,omptempty"`
	PosProductLoadBackground                    *Bool      `xmlrpc:"pos_product_load_background,omptempty"`
	PosProxyIp                                  *String    `xmlrpc:"pos_proxy_ip,omptempty"`
	PosReceiptFooter                            *String    `xmlrpc:"pos_receipt_footer,omptempty"`
	PosReceiptHeader                            *String    `xmlrpc:"pos_receipt_header,omptempty"`
	PosRestrictPriceControl                     *Bool      `xmlrpc:"pos_restrict_price_control,omptempty"`
	PosRoundingMethod                           *Many2One  `xmlrpc:"pos_rounding_method,omptempty"`
	PosRouteId                                  *Many2One  `xmlrpc:"pos_route_id,omptempty"`
	PosSelectableCategIds                       *Relation  `xmlrpc:"pos_selectable_categ_ids,omptempty"`
	PosSequenceId                               *Many2One  `xmlrpc:"pos_sequence_id,omptempty"`
	PosSetMaximumDifference                     *Bool      `xmlrpc:"pos_set_maximum_difference,omptempty"`
	PosShipLater                                *Bool      `xmlrpc:"pos_ship_later,omptempty"`
	PosStartCategory                            *Bool      `xmlrpc:"pos_start_category,omptempty"`
	PosTaxRegimeSelection                       *Bool      `xmlrpc:"pos_tax_regime_selection,omptempty"`
	PosTipProductId                             *Many2One  `xmlrpc:"pos_tip_product_id,omptempty"`
	PosUsePricelist                             *Bool      `xmlrpc:"pos_use_pricelist,omptempty"`
	PosWarehouseId                              *Many2One  `xmlrpc:"pos_warehouse_id,omptempty"`
	PreviewReady                                *Bool      `xmlrpc:"preview_ready,omptempty"`
	PrimaryColor                                *String    `xmlrpc:"primary_color,omptempty"`
	ProductPricelistSetting                     *Selection `xmlrpc:"product_pricelist_setting,omptempty"`
	ProductVolumeVolumeInCubicFeet              *Selection `xmlrpc:"product_volume_volume_in_cubic_feet,omptempty"`
	ProductWeightInLbs                          *Selection `xmlrpc:"product_weight_in_lbs,omptempty"`
	ProfilingEnabledUntil                       *Time      `xmlrpc:"profiling_enabled_until,omptempty"`
	PurchaseTaxId                               *Many2One  `xmlrpc:"purchase_tax_id,omptempty"`
	QrCode                                      *Bool      `xmlrpc:"qr_code,omptempty"`
	QuickEditMode                               *Selection `xmlrpc:"quick_edit_mode,omptempty"`
	ReportFooter                                *String    `xmlrpc:"report_footer,omptempty"`
	RestrictTemplateRendering                   *Bool      `xmlrpc:"restrict_template_rendering,omptempty"`
	SaleTaxId                                   *Many2One  `xmlrpc:"sale_tax_id,omptempty"`
	SecondaryColor                              *String    `xmlrpc:"secondary_color,omptempty"`
	ShowEffect                                  *Bool      `xmlrpc:"show_effect,omptempty"`
	ShowLineSubtotalsTaxSelection               *Selection `xmlrpc:"show_line_subtotals_tax_selection,omptempty"`
	SnailmailColor                              *Bool      `xmlrpc:"snailmail_color,omptempty"`
	SnailmailCover                              *Bool      `xmlrpc:"snailmail_cover,omptempty"`
	SnailmailDuplex                             *Bool      `xmlrpc:"snailmail_duplex,omptempty"`
	StockMoveEmailValidation                    *Bool      `xmlrpc:"stock_move_email_validation,omptempty"`
	StockMoveSmsValidation                      *Bool      `xmlrpc:"stock_move_sms_validation,omptempty"`
	StockSmsConfirmationTemplateId              *Many2One  `xmlrpc:"stock_sms_confirmation_template_id,omptempty"`
	TaxCalculationRoundingMethod                *Selection `xmlrpc:"tax_calculation_rounding_method,omptempty"`
	TaxCashBasisJournalId                       *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omptempty"`
	TaxExigibility                              *Bool      `xmlrpc:"tax_exigibility,omptempty"`
	TermsType                                   *Selection `xmlrpc:"terms_type,omptempty"`
	TransferAccountId                           *Many2One  `xmlrpc:"transfer_account_id,omptempty"`
	TwilioAccountSid                            *String    `xmlrpc:"twilio_account_sid,omptempty"`
	TwilioAccountToken                          *String    `xmlrpc:"twilio_account_token,omptempty"`
	UnsplashAccessKey                           *String    `xmlrpc:"unsplash_access_key,omptempty"`
	UnsplashAppId                               *String    `xmlrpc:"unsplash_app_id,omptempty"`
	UpdateStockQuantities                       *Selection `xmlrpc:"update_stock_quantities,omptempty"`
	UseInvoiceTerms                             *Bool      `xmlrpc:"use_invoice_terms,omptempty"`
	UseTwilioRtcServers                         *Bool      `xmlrpc:"use_twilio_rtc_servers,omptempty"`
	UserDefaultRights                           *Bool      `xmlrpc:"user_default_rights,omptempty"`
	VatCheckVies                                *Bool      `xmlrpc:"vat_check_vies,omptempty"`
	WazapServer                                 *String    `xmlrpc:"wazap_server,omptempty"`
	WazapServerToken                            *String    `xmlrpc:"wazap_server_token,omptempty"`
	WriteDate                                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ResConfigSettingss represents array of res.config.settings model.
type ResConfigSettingss []ResConfigSettings

// ResConfigSettingsModel is the odoo model name.
const ResConfigSettingsModel = "res.config.settings"

// Many2One convert ResConfigSettings to *Many2One.
func (rcs *ResConfigSettings) Many2One() *Many2One {
	return NewMany2One(rcs.Id.Get(), "")
}

// CreateResConfigSettings creates a new res.config.settings model and returns its id.
func (c *Client) CreateResConfigSettings(rcs *ResConfigSettings) (int64, error) {
	return c.Create(ResConfigSettingsModel, rcs)
}

// UpdateResConfigSettings updates an existing res.config.settings record.
func (c *Client) UpdateResConfigSettings(rcs *ResConfigSettings) error {
	return c.UpdateResConfigSettingss([]int64{rcs.Id.Get()}, rcs)
}

// UpdateResConfigSettingss updates existing res.config.settings records.
// All records (represented by ids) will be updated by rcs values.
func (c *Client) UpdateResConfigSettingss(ids []int64, rcs *ResConfigSettings) error {
	return c.Update(ResConfigSettingsModel, ids, rcs)
}

// DeleteResConfigSettings deletes an existing res.config.settings record.
func (c *Client) DeleteResConfigSettings(id int64) error {
	return c.DeleteResConfigSettingss([]int64{id})
}

// DeleteResConfigSettingss deletes existing res.config.settings records.
func (c *Client) DeleteResConfigSettingss(ids []int64) error {
	return c.Delete(ResConfigSettingsModel, ids)
}

// GetResConfigSettings gets res.config.settings existing record.
func (c *Client) GetResConfigSettings(id int64) (*ResConfigSettings, error) {
	rcss, err := c.GetResConfigSettingss([]int64{id})
	if err != nil {
		return nil, err
	}
	if rcss != nil && len(*rcss) > 0 {
		return &((*rcss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.config.settings not found", id)
}

// GetResConfigSettingss gets res.config.settings existing records.
func (c *Client) GetResConfigSettingss(ids []int64) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.Read(ResConfigSettingsModel, ids, nil, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettings finds res.config.settings record by querying it with criteria.
func (c *Client) FindResConfigSettings(criteria *Criteria) (*ResConfigSettings, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, NewOptions().Limit(1), rcss); err != nil {
		return nil, err
	}
	if rcss != nil && len(*rcss) > 0 {
		return &((*rcss)[0]), nil
	}
	return nil, fmt.Errorf("res.config.settings was not found with criteria %v", criteria)
}

// FindResConfigSettingss finds res.config.settings records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingss(criteria *Criteria, options *Options) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, options, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettingsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResConfigSettingsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResConfigSettingsId finds record id by querying it with criteria.
func (c *Client) FindResConfigSettingsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResConfigSettingsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.config.settings was not found with criteria %v and options %v", criteria, options)
}
