package odoo

import (
	"fmt"
)

// QueueJob represents queue.job model.
type QueueJob struct {
	LastUpdate                  *Time       `xmlrpc:"__last_update,omptempty"`
	ActivityDateDeadline        *Time       `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection  `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String     `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation   `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection  `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String     `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon            *String     `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId              *Many2One   `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One   `xmlrpc:"activity_user_id,omptempty"`
	Args                        interface{} `xmlrpc:"args,omptempty"`
	Channel                     *String     `xmlrpc:"channel,omptempty"`
	ChannelMethodName           *String     `xmlrpc:"channel_method_name,omptempty"`
	CompanyId                   *Many2One   `xmlrpc:"company_id,omptempty"`
	DateCancelled               *Time       `xmlrpc:"date_cancelled,omptempty"`
	DateCreated                 *Time       `xmlrpc:"date_created,omptempty"`
	DateDone                    *Time       `xmlrpc:"date_done,omptempty"`
	DateEnqueued                *Time       `xmlrpc:"date_enqueued,omptempty"`
	DateStarted                 *Time       `xmlrpc:"date_started,omptempty"`
	Dependencies                interface{} `xmlrpc:"dependencies,omptempty"`
	DependencyGraph             interface{} `xmlrpc:"dependency_graph,omptempty"`
	DisplayName                 *String     `xmlrpc:"display_name,omptempty"`
	Eta                         *Time       `xmlrpc:"eta,omptempty"`
	ExcInfo                     *String     `xmlrpc:"exc_info,omptempty"`
	ExcMessage                  *String     `xmlrpc:"exc_message,omptempty"`
	ExcName                     *String     `xmlrpc:"exc_name,omptempty"`
	ExecTime                    *Float      `xmlrpc:"exec_time,omptempty"`
	FuncString                  *String     `xmlrpc:"func_string,omptempty"`
	GraphJobsCount              *Int        `xmlrpc:"graph_jobs_count,omptempty"`
	GraphUuid                   *String     `xmlrpc:"graph_uuid,omptempty"`
	HasMessage                  *Bool       `xmlrpc:"has_message,omptempty"`
	Id                          *Int        `xmlrpc:"id,omptempty"`
	IdentityKey                 *String     `xmlrpc:"identity_key,omptempty"`
	JobFunctionId               *Many2One   `xmlrpc:"job_function_id,omptempty"`
	Kwargs                      interface{} `xmlrpc:"kwargs,omptempty"`
	MaxRetries                  *Int        `xmlrpc:"max_retries,omptempty"`
	MessageAttachmentCount      *Int        `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds          *Relation   `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool       `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int        `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool       `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation   `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool       `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One   `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool       `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int        `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation   `xmlrpc:"message_partner_ids,omptempty"`
	MethodName                  *String     `xmlrpc:"method_name,omptempty"`
	ModelName                   *String     `xmlrpc:"model_name,omptempty"`
	MyActivityDateDeadline      *Time       `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                        *String     `xmlrpc:"name,omptempty"`
	Priority                    *Int        `xmlrpc:"priority,omptempty"`
	RecordIds                   interface{} `xmlrpc:"record_ids,omptempty"`
	Records                     interface{} `xmlrpc:"records,omptempty"`
	Result                      *String     `xmlrpc:"result,omptempty"`
	Retry                       *Int        `xmlrpc:"retry,omptempty"`
	State                       *Selection  `xmlrpc:"state,omptempty"`
	UserId                      *Many2One   `xmlrpc:"user_id,omptempty"`
	Uuid                        *String     `xmlrpc:"uuid,omptempty"`
	WebsiteMessageIds           *Relation   `xmlrpc:"website_message_ids,omptempty"`
	WorkerPid                   *Int        `xmlrpc:"worker_pid,omptempty"`
}

// QueueJobs represents array of queue.job model.
type QueueJobs []QueueJob

// QueueJobModel is the odoo model name.
const QueueJobModel = "queue.job"

// Many2One convert QueueJob to *Many2One.
func (qj *QueueJob) Many2One() *Many2One {
	return NewMany2One(qj.Id.Get(), "")
}

// CreateQueueJob creates a new queue.job model and returns its id.
func (c *Client) CreateQueueJob(qj *QueueJob) (int64, error) {
	return c.Create(QueueJobModel, qj)
}

// UpdateQueueJob updates an existing queue.job record.
func (c *Client) UpdateQueueJob(qj *QueueJob) error {
	return c.UpdateQueueJobs([]int64{qj.Id.Get()}, qj)
}

// UpdateQueueJobs updates existing queue.job records.
// All records (represented by ids) will be updated by qj values.
func (c *Client) UpdateQueueJobs(ids []int64, qj *QueueJob) error {
	return c.Update(QueueJobModel, ids, qj)
}

// DeleteQueueJob deletes an existing queue.job record.
func (c *Client) DeleteQueueJob(id int64) error {
	return c.DeleteQueueJobs([]int64{id})
}

// DeleteQueueJobs deletes existing queue.job records.
func (c *Client) DeleteQueueJobs(ids []int64) error {
	return c.Delete(QueueJobModel, ids)
}

// GetQueueJob gets queue.job existing record.
func (c *Client) GetQueueJob(id int64) (*QueueJob, error) {
	qjs, err := c.GetQueueJobs([]int64{id})
	if err != nil {
		return nil, err
	}
	if qjs != nil && len(*qjs) > 0 {
		return &((*qjs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of queue.job not found", id)
}

// GetQueueJobs gets queue.job existing records.
func (c *Client) GetQueueJobs(ids []int64) (*QueueJobs, error) {
	qjs := &QueueJobs{}
	if err := c.Read(QueueJobModel, ids, nil, qjs); err != nil {
		return nil, err
	}
	return qjs, nil
}

// FindQueueJob finds queue.job record by querying it with criteria.
func (c *Client) FindQueueJob(criteria *Criteria) (*QueueJob, error) {
	qjs := &QueueJobs{}
	if err := c.SearchRead(QueueJobModel, criteria, NewOptions().Limit(1), qjs); err != nil {
		return nil, err
	}
	if qjs != nil && len(*qjs) > 0 {
		return &((*qjs)[0]), nil
	}
	return nil, fmt.Errorf("queue.job was not found with criteria %v", criteria)
}

// FindQueueJobs finds queue.job records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueJobs(criteria *Criteria, options *Options) (*QueueJobs, error) {
	qjs := &QueueJobs{}
	if err := c.SearchRead(QueueJobModel, criteria, options, qjs); err != nil {
		return nil, err
	}
	return qjs, nil
}

// FindQueueJobIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueJobIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(QueueJobModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindQueueJobId finds record id by querying it with criteria.
func (c *Client) FindQueueJobId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QueueJobModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("queue.job was not found with criteria %v and options %v", criteria, options)
}
