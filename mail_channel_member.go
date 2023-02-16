package odoo

import (
	"fmt"
)

// MailChannelMember represents mail.channel.member model.
type MailChannelMember struct {
	LastUpdate           *Time      `xmlrpc:"__last_update,omptempty"`
	ChannelId            *Many2One  `xmlrpc:"channel_id,omptempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omptempty"`
	CustomChannelName    *String    `xmlrpc:"custom_channel_name,omptempty"`
	DisplayName          *String    `xmlrpc:"display_name,omptempty"`
	FetchedMessageId     *Many2One  `xmlrpc:"fetched_message_id,omptempty"`
	FoldState            *Selection `xmlrpc:"fold_state,omptempty"`
	GuestId              *Many2One  `xmlrpc:"guest_id,omptempty"`
	Id                   *Int       `xmlrpc:"id,omptempty"`
	IsMinimized          *Bool      `xmlrpc:"is_minimized,omptempty"`
	IsPinned             *Bool      `xmlrpc:"is_pinned,omptempty"`
	LastInterestDt       *Time      `xmlrpc:"last_interest_dt,omptempty"`
	LastSeenDt           *Time      `xmlrpc:"last_seen_dt,omptempty"`
	MessageUnreadCounter *Int       `xmlrpc:"message_unread_counter,omptempty"`
	PartnerEmail         *String    `xmlrpc:"partner_email,omptempty"`
	PartnerId            *Many2One  `xmlrpc:"partner_id,omptempty"`
	RtcInvitingSessionId *Many2One  `xmlrpc:"rtc_inviting_session_id,omptempty"`
	RtcSessionIds        *Relation  `xmlrpc:"rtc_session_ids,omptempty"`
	SeenMessageId        *Many2One  `xmlrpc:"seen_message_id,omptempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailChannelMembers represents array of mail.channel.member model.
type MailChannelMembers []MailChannelMember

// MailChannelMemberModel is the odoo model name.
const MailChannelMemberModel = "mail.channel.member"

// Many2One convert MailChannelMember to *Many2One.
func (mcm *MailChannelMember) Many2One() *Many2One {
	return NewMany2One(mcm.Id.Get(), "")
}

// CreateMailChannelMember creates a new mail.channel.member model and returns its id.
func (c *Client) CreateMailChannelMember(mcm *MailChannelMember) (int64, error) {
	return c.Create(MailChannelMemberModel, mcm)
}

// UpdateMailChannelMember updates an existing mail.channel.member record.
func (c *Client) UpdateMailChannelMember(mcm *MailChannelMember) error {
	return c.UpdateMailChannelMembers([]int64{mcm.Id.Get()}, mcm)
}

// UpdateMailChannelMembers updates existing mail.channel.member records.
// All records (represented by ids) will be updated by mcm values.
func (c *Client) UpdateMailChannelMembers(ids []int64, mcm *MailChannelMember) error {
	return c.Update(MailChannelMemberModel, ids, mcm)
}

// DeleteMailChannelMember deletes an existing mail.channel.member record.
func (c *Client) DeleteMailChannelMember(id int64) error {
	return c.DeleteMailChannelMembers([]int64{id})
}

// DeleteMailChannelMembers deletes existing mail.channel.member records.
func (c *Client) DeleteMailChannelMembers(ids []int64) error {
	return c.Delete(MailChannelMemberModel, ids)
}

// GetMailChannelMember gets mail.channel.member existing record.
func (c *Client) GetMailChannelMember(id int64) (*MailChannelMember, error) {
	mcms, err := c.GetMailChannelMembers([]int64{id})
	if err != nil {
		return nil, err
	}
	if mcms != nil && len(*mcms) > 0 {
		return &((*mcms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.channel.member not found", id)
}

// GetMailChannelMembers gets mail.channel.member existing records.
func (c *Client) GetMailChannelMembers(ids []int64) (*MailChannelMembers, error) {
	mcms := &MailChannelMembers{}
	if err := c.Read(MailChannelMemberModel, ids, nil, mcms); err != nil {
		return nil, err
	}
	return mcms, nil
}

// FindMailChannelMember finds mail.channel.member record by querying it with criteria.
func (c *Client) FindMailChannelMember(criteria *Criteria) (*MailChannelMember, error) {
	mcms := &MailChannelMembers{}
	if err := c.SearchRead(MailChannelMemberModel, criteria, NewOptions().Limit(1), mcms); err != nil {
		return nil, err
	}
	if mcms != nil && len(*mcms) > 0 {
		return &((*mcms)[0]), nil
	}
	return nil, fmt.Errorf("mail.channel.member was not found with criteria %v", criteria)
}

// FindMailChannelMembers finds mail.channel.member records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailChannelMembers(criteria *Criteria, options *Options) (*MailChannelMembers, error) {
	mcms := &MailChannelMembers{}
	if err := c.SearchRead(MailChannelMemberModel, criteria, options, mcms); err != nil {
		return nil, err
	}
	return mcms, nil
}

// FindMailChannelMemberIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailChannelMemberIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailChannelMemberModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailChannelMemberId finds record id by querying it with criteria.
func (c *Client) FindMailChannelMemberId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailChannelMemberModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.channel.member was not found with criteria %v and options %v", criteria, options)
}
