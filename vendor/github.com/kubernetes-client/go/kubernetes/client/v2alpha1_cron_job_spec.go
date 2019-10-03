/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// CronJobSpec describes how the job execution will look like and when it will actually run.
type V2alpha1CronJobSpec struct {

	// Specifies how to treat concurrent executions of a Job. Valid values are: - \"Allow\" (default): allows CronJobs to run concurrently; - \"Forbid\": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - \"Replace\": cancels currently running job and replaces it with a new one
	ConcurrencyPolicy string `json:"concurrencyPolicy,omitempty"`

	// The number of failed finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
	FailedJobsHistoryLimit int32 `json:"failedJobsHistoryLimit,omitempty"`

	// Specifies the job that will be created when executing a CronJob.
	JobTemplate *V2alpha1JobTemplateSpec `json:"jobTemplate"`

	// The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
	Schedule string `json:"schedule"`

	// Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.
	StartingDeadlineSeconds int64 `json:"startingDeadlineSeconds,omitempty"`

	// The number of successful finished jobs to retain. This is a pointer to distinguish between explicit zero and not specified.
	SuccessfulJobsHistoryLimit int32 `json:"successfulJobsHistoryLimit,omitempty"`

	// This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.  Defaults to false.
	Suspend bool `json:"suspend,omitempty"`
}