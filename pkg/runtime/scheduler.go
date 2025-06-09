package runtime

import (
	"context"
	"hal/pkg/state"
	"time"
)

// Scheduler manages scheduled graph executions
type Scheduler struct {
	executor *Executor
	jobs     map[string]*ScheduledJob
}

// ScheduledJob represents a scheduled graph execution
type ScheduledJob struct {
	ID        string
	GraphName string
	Schedule  Schedule
	State     *state.State
}

// Schedule defines when a job should run
type Schedule interface {
	Next(time.Time) time.Time
}

// IntervalSchedule runs at fixed intervals
type IntervalSchedule struct {
	Interval time.Duration
}

// Next returns the next execution time
func (s *IntervalSchedule) Next(t time.Time) time.Time {
	return t.Add(s.Interval)
}

// NewScheduler creates a new scheduler
func NewScheduler(executor *Executor) *Scheduler {
	return &Scheduler{
		executor: executor,
		jobs:     make(map[string]*ScheduledJob),
	}
}

// Schedule adds a job to the scheduler
func (s *Scheduler) Schedule(job *ScheduledJob) {
	s.jobs[job.ID] = job
}

// Start begins the scheduler
func (s *Scheduler) Start(ctx context.Context) {
	// Implementation would include timer management and job execution
}
