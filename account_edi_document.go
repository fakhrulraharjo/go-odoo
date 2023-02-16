package odoo

import (
	"fmt"
)

// AccountEdiDocument represents account.edi.document model.
type AccountEdiDocument struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omptempty"`
	AttachmentId  *Many2One  `xmlrpc:"attachment_id,omptempty"`
	BlockingLevel *Selection `xmlrpc:"blocking_level,omptempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String    `xmlrpc:"display_name,omptempty"`
	EdiContent    *String    `xmlrpc:"edi_content,omptempty"`
	EdiFormatId   *Many2One  `xmlrpc:"edi_format_id,omptempty"`
	EdiFormatName *String    `xmlrpc:"edi_format_name,omptempty"`
	Error         *String    `xmlrpc:"error,omptempty"`
	Id            *Int       `xmlrpc:"id,omptempty"`
	MoveId        *Many2One  `xmlrpc:"move_id,omptempty"`
	Name          *String    `xmlrpc:"name,omptempty"`
	State         *Selection `xmlrpc:"state,omptempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountEdiDocuments represents array of account.edi.document model.
type AccountEdiDocuments []AccountEdiDocument

// AccountEdiDocumentModel is the odoo model name.
const AccountEdiDocumentModel = "account.edi.document"

// Many2One convert AccountEdiDocument to *Many2One.
func (aed *AccountEdiDocument) Many2One() *Many2One {
	return NewMany2One(aed.Id.Get(), "")
}

// CreateAccountEdiDocument creates a new account.edi.document model and returns its id.
func (c *Client) CreateAccountEdiDocument(aed *AccountEdiDocument) (int64, error) {
	return c.Create(AccountEdiDocumentModel, aed)
}

// UpdateAccountEdiDocument updates an existing account.edi.document record.
func (c *Client) UpdateAccountEdiDocument(aed *AccountEdiDocument) error {
	return c.UpdateAccountEdiDocuments([]int64{aed.Id.Get()}, aed)
}

// UpdateAccountEdiDocuments updates existing account.edi.document records.
// All records (represented by ids) will be updated by aed values.
func (c *Client) UpdateAccountEdiDocuments(ids []int64, aed *AccountEdiDocument) error {
	return c.Update(AccountEdiDocumentModel, ids, aed)
}

// DeleteAccountEdiDocument deletes an existing account.edi.document record.
func (c *Client) DeleteAccountEdiDocument(id int64) error {
	return c.DeleteAccountEdiDocuments([]int64{id})
}

// DeleteAccountEdiDocuments deletes existing account.edi.document records.
func (c *Client) DeleteAccountEdiDocuments(ids []int64) error {
	return c.Delete(AccountEdiDocumentModel, ids)
}

// GetAccountEdiDocument gets account.edi.document existing record.
func (c *Client) GetAccountEdiDocument(id int64) (*AccountEdiDocument, error) {
	aeds, err := c.GetAccountEdiDocuments([]int64{id})
	if err != nil {
		return nil, err
	}
	if aeds != nil && len(*aeds) > 0 {
		return &((*aeds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.edi.document not found", id)
}

// GetAccountEdiDocuments gets account.edi.document existing records.
func (c *Client) GetAccountEdiDocuments(ids []int64) (*AccountEdiDocuments, error) {
	aeds := &AccountEdiDocuments{}
	if err := c.Read(AccountEdiDocumentModel, ids, nil, aeds); err != nil {
		return nil, err
	}
	return aeds, nil
}

// FindAccountEdiDocument finds account.edi.document record by querying it with criteria.
func (c *Client) FindAccountEdiDocument(criteria *Criteria) (*AccountEdiDocument, error) {
	aeds := &AccountEdiDocuments{}
	if err := c.SearchRead(AccountEdiDocumentModel, criteria, NewOptions().Limit(1), aeds); err != nil {
		return nil, err
	}
	if aeds != nil && len(*aeds) > 0 {
		return &((*aeds)[0]), nil
	}
	return nil, fmt.Errorf("account.edi.document was not found with criteria %v", criteria)
}

// FindAccountEdiDocuments finds account.edi.document records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiDocuments(criteria *Criteria, options *Options) (*AccountEdiDocuments, error) {
	aeds := &AccountEdiDocuments{}
	if err := c.SearchRead(AccountEdiDocumentModel, criteria, options, aeds); err != nil {
		return nil, err
	}
	return aeds, nil
}

// FindAccountEdiDocumentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountEdiDocumentIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountEdiDocumentModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountEdiDocumentId finds record id by querying it with criteria.
func (c *Client) FindAccountEdiDocumentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountEdiDocumentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.edi.document was not found with criteria %v and options %v", criteria, options)
}
