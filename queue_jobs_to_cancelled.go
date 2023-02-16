package odoo

import (
	"fmt"
)

// QueueJobsToCancelled represents queue.jobs.to.cancelled model.
type QueueJobsToCancelled struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	JobIds      *Relation `xmlrpc:"job_ids,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// QueueJobsToCancelleds represents array of queue.jobs.to.cancelled model.
type QueueJobsToCancelleds []QueueJobsToCancelled

// QueueJobsToCancelledModel is the odoo model name.
const QueueJobsToCancelledModel = "queue.jobs.to.cancelled"

// Many2One convert QueueJobsToCancelled to *Many2One.
func (qjtc *QueueJobsToCancelled) Many2One() *Many2One {
	return NewMany2One(qjtc.Id.Get(), "")
}

// CreateQueueJobsToCancelled creates a new queue.jobs.to.cancelled model and returns its id.
func (c *Client) CreateQueueJobsToCancelled(qjtc *QueueJobsToCancelled) (int64, error) {
	return c.Create(QueueJobsToCancelledModel, qjtc)
}

// UpdateQueueJobsToCancelled updates an existing queue.jobs.to.cancelled record.
func (c *Client) UpdateQueueJobsToCancelled(qjtc *QueueJobsToCancelled) error {
	return c.UpdateQueueJobsToCancelleds([]int64{qjtc.Id.Get()}, qjtc)
}

// UpdateQueueJobsToCancelleds updates existing queue.jobs.to.cancelled records.
// All records (represented by ids) will be updated by qjtc values.
func (c *Client) UpdateQueueJobsToCancelleds(ids []int64, qjtc *QueueJobsToCancelled) error {
	return c.Update(QueueJobsToCancelledModel, ids, qjtc)
}

// DeleteQueueJobsToCancelled deletes an existing queue.jobs.to.cancelled record.
func (c *Client) DeleteQueueJobsToCancelled(id int64) error {
	return c.DeleteQueueJobsToCancelleds([]int64{id})
}

// DeleteQueueJobsToCancelleds deletes existing queue.jobs.to.cancelled records.
func (c *Client) DeleteQueueJobsToCancelleds(ids []int64) error {
	return c.Delete(QueueJobsToCancelledModel, ids)
}

// GetQueueJobsToCancelled gets queue.jobs.to.cancelled existing record.
func (c *Client) GetQueueJobsToCancelled(id int64) (*QueueJobsToCancelled, error) {
	qjtcs, err := c.GetQueueJobsToCancelleds([]int64{id})
	if err != nil {
		return nil, err
	}
	if qjtcs != nil && len(*qjtcs) > 0 {
		return &((*qjtcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of queue.jobs.to.cancelled not found", id)
}

// GetQueueJobsToCancelleds gets queue.jobs.to.cancelled existing records.
func (c *Client) GetQueueJobsToCancelleds(ids []int64) (*QueueJobsToCancelleds, error) {
	qjtcs := &QueueJobsToCancelleds{}
	if err := c.Read(QueueJobsToCancelledModel, ids, nil, qjtcs); err != nil {
		return nil, err
	}
	return qjtcs, nil
}

// FindQueueJobsToCancelled finds queue.jobs.to.cancelled record by querying it with criteria.
func (c *Client) FindQueueJobsToCancelled(criteria *Criteria) (*QueueJobsToCancelled, error) {
	qjtcs := &QueueJobsToCancelleds{}
	if err := c.SearchRead(QueueJobsToCancelledModel, criteria, NewOptions().Limit(1), qjtcs); err != nil {
		return nil, err
	}
	if qjtcs != nil && len(*qjtcs) > 0 {
		return &((*qjtcs)[0]), nil
	}
	return nil, fmt.Errorf("queue.jobs.to.cancelled was not found with criteria %v", criteria)
}

// FindQueueJobsToCancelleds finds queue.jobs.to.cancelled records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueJobsToCancelleds(criteria *Criteria, options *Options) (*QueueJobsToCancelleds, error) {
	qjtcs := &QueueJobsToCancelleds{}
	if err := c.SearchRead(QueueJobsToCancelledModel, criteria, options, qjtcs); err != nil {
		return nil, err
	}
	return qjtcs, nil
}

// FindQueueJobsToCancelledIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueJobsToCancelledIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(QueueJobsToCancelledModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindQueueJobsToCancelledId finds record id by querying it with criteria.
func (c *Client) FindQueueJobsToCancelledId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QueueJobsToCancelledModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("queue.jobs.to.cancelled was not found with criteria %v and options %v", criteria, options)
}
