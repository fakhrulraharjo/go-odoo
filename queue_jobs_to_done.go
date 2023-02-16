package odoo

import (
	"fmt"
)

// QueueJobsToDone represents queue.jobs.to.done model.
type QueueJobsToDone struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	JobIds      *Relation `xmlrpc:"job_ids,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// QueueJobsToDones represents array of queue.jobs.to.done model.
type QueueJobsToDones []QueueJobsToDone

// QueueJobsToDoneModel is the odoo model name.
const QueueJobsToDoneModel = "queue.jobs.to.done"

// Many2One convert QueueJobsToDone to *Many2One.
func (qjtd *QueueJobsToDone) Many2One() *Many2One {
	return NewMany2One(qjtd.Id.Get(), "")
}

// CreateQueueJobsToDone creates a new queue.jobs.to.done model and returns its id.
func (c *Client) CreateQueueJobsToDone(qjtd *QueueJobsToDone) (int64, error) {
	return c.Create(QueueJobsToDoneModel, qjtd)
}

// UpdateQueueJobsToDone updates an existing queue.jobs.to.done record.
func (c *Client) UpdateQueueJobsToDone(qjtd *QueueJobsToDone) error {
	return c.UpdateQueueJobsToDones([]int64{qjtd.Id.Get()}, qjtd)
}

// UpdateQueueJobsToDones updates existing queue.jobs.to.done records.
// All records (represented by ids) will be updated by qjtd values.
func (c *Client) UpdateQueueJobsToDones(ids []int64, qjtd *QueueJobsToDone) error {
	return c.Update(QueueJobsToDoneModel, ids, qjtd)
}

// DeleteQueueJobsToDone deletes an existing queue.jobs.to.done record.
func (c *Client) DeleteQueueJobsToDone(id int64) error {
	return c.DeleteQueueJobsToDones([]int64{id})
}

// DeleteQueueJobsToDones deletes existing queue.jobs.to.done records.
func (c *Client) DeleteQueueJobsToDones(ids []int64) error {
	return c.Delete(QueueJobsToDoneModel, ids)
}

// GetQueueJobsToDone gets queue.jobs.to.done existing record.
func (c *Client) GetQueueJobsToDone(id int64) (*QueueJobsToDone, error) {
	qjtds, err := c.GetQueueJobsToDones([]int64{id})
	if err != nil {
		return nil, err
	}
	if qjtds != nil && len(*qjtds) > 0 {
		return &((*qjtds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of queue.jobs.to.done not found", id)
}

// GetQueueJobsToDones gets queue.jobs.to.done existing records.
func (c *Client) GetQueueJobsToDones(ids []int64) (*QueueJobsToDones, error) {
	qjtds := &QueueJobsToDones{}
	if err := c.Read(QueueJobsToDoneModel, ids, nil, qjtds); err != nil {
		return nil, err
	}
	return qjtds, nil
}

// FindQueueJobsToDone finds queue.jobs.to.done record by querying it with criteria.
func (c *Client) FindQueueJobsToDone(criteria *Criteria) (*QueueJobsToDone, error) {
	qjtds := &QueueJobsToDones{}
	if err := c.SearchRead(QueueJobsToDoneModel, criteria, NewOptions().Limit(1), qjtds); err != nil {
		return nil, err
	}
	if qjtds != nil && len(*qjtds) > 0 {
		return &((*qjtds)[0]), nil
	}
	return nil, fmt.Errorf("queue.jobs.to.done was not found with criteria %v", criteria)
}

// FindQueueJobsToDones finds queue.jobs.to.done records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueJobsToDones(criteria *Criteria, options *Options) (*QueueJobsToDones, error) {
	qjtds := &QueueJobsToDones{}
	if err := c.SearchRead(QueueJobsToDoneModel, criteria, options, qjtds); err != nil {
		return nil, err
	}
	return qjtds, nil
}

// FindQueueJobsToDoneIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueJobsToDoneIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(QueueJobsToDoneModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindQueueJobsToDoneId finds record id by querying it with criteria.
func (c *Client) FindQueueJobsToDoneId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QueueJobsToDoneModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("queue.jobs.to.done was not found with criteria %v and options %v", criteria, options)
}
