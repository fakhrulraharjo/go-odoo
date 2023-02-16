package odoo

import (
	"fmt"
)

// MailChannel represents mail.channel model.
type MailChannel struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	Active                   *Bool      `xmlrpc:"active,omptempty"`
	Avatar128                *String    `xmlrpc:"avatar_128,omptempty"`
	ChannelMemberIds         *Relation  `xmlrpc:"channel_member_ids,omptempty"`
	ChannelPartnerIds        *Relation  `xmlrpc:"channel_partner_ids,omptempty"`
	ChannelType              *Selection `xmlrpc:"channel_type,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	DefaultDisplayMode       *Selection `xmlrpc:"default_display_mode,omptempty"`
	Description              *String    `xmlrpc:"description,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	GroupIds                 *Relation  `xmlrpc:"group_ids,omptempty"`
	GroupPublicId            *Many2One  `xmlrpc:"group_public_id,omptempty"`
	HasMessage               *Bool      `xmlrpc:"has_message,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	Image128                 *String    `xmlrpc:"image_128,omptempty"`
	InvitationUrl            *String    `xmlrpc:"invitation_url,omptempty"`
	IsChat                   *Bool      `xmlrpc:"is_chat,omptempty"`
	IsMember                 *Bool      `xmlrpc:"is_member,omptempty"`
	IsPinned                 *Bool      `xmlrpc:"is_pinned,omptempty"`
	MemberCount              *Int       `xmlrpc:"member_count,omptempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId  *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	RtcSessionIds            *Relation  `xmlrpc:"rtc_session_ids,omptempty"`
	Uuid                     *String    `xmlrpc:"uuid,omptempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailChannels represents array of mail.channel model.
type MailChannels []MailChannel

// MailChannelModel is the odoo model name.
const MailChannelModel = "mail.channel"

// Many2One convert MailChannel to *Many2One.
func (mc *MailChannel) Many2One() *Many2One {
	return NewMany2One(mc.Id.Get(), "")
}

// CreateMailChannel creates a new mail.channel model and returns its id.
func (c *Client) CreateMailChannel(mc *MailChannel) (int64, error) {
	return c.Create(MailChannelModel, mc)
}

// UpdateMailChannel updates an existing mail.channel record.
func (c *Client) UpdateMailChannel(mc *MailChannel) error {
	return c.UpdateMailChannels([]int64{mc.Id.Get()}, mc)
}

// UpdateMailChannels updates existing mail.channel records.
// All records (represented by ids) will be updated by mc values.
func (c *Client) UpdateMailChannels(ids []int64, mc *MailChannel) error {
	return c.Update(MailChannelModel, ids, mc)
}

// DeleteMailChannel deletes an existing mail.channel record.
func (c *Client) DeleteMailChannel(id int64) error {
	return c.DeleteMailChannels([]int64{id})
}

// DeleteMailChannels deletes existing mail.channel records.
func (c *Client) DeleteMailChannels(ids []int64) error {
	return c.Delete(MailChannelModel, ids)
}

// GetMailChannel gets mail.channel existing record.
func (c *Client) GetMailChannel(id int64) (*MailChannel, error) {
	mcs, err := c.GetMailChannels([]int64{id})
	if err != nil {
		return nil, err
	}
	if mcs != nil && len(*mcs) > 0 {
		return &((*mcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.channel not found", id)
}

// GetMailChannels gets mail.channel existing records.
func (c *Client) GetMailChannels(ids []int64) (*MailChannels, error) {
	mcs := &MailChannels{}
	if err := c.Read(MailChannelModel, ids, nil, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMailChannel finds mail.channel record by querying it with criteria.
func (c *Client) FindMailChannel(criteria *Criteria) (*MailChannel, error) {
	mcs := &MailChannels{}
	if err := c.SearchRead(MailChannelModel, criteria, NewOptions().Limit(1), mcs); err != nil {
		return nil, err
	}
	if mcs != nil && len(*mcs) > 0 {
		return &((*mcs)[0]), nil
	}
	return nil, fmt.Errorf("mail.channel was not found with criteria %v", criteria)
}

// FindMailChannels finds mail.channel records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailChannels(criteria *Criteria, options *Options) (*MailChannels, error) {
	mcs := &MailChannels{}
	if err := c.SearchRead(MailChannelModel, criteria, options, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMailChannelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailChannelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailChannelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailChannelId finds record id by querying it with criteria.
func (c *Client) FindMailChannelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailChannelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.channel was not found with criteria %v and options %v", criteria, options)
}
