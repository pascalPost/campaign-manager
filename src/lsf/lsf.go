package lsf

import "time"

type Job struct {
	Command    string
	ExHosts    []string
	FromHost   string
	JobId      string
	JobName    string
	JobStatus  string
	Queue      string
	SubmitTime time.Time
	User       string
}

func Jobs() []Job {
	return []Job{
		{
			Command: "sleep 20000",
			ExHosts: []string{
				"hostA",
			},
			FromHost:  "hostA",
			JobId:     "528",
			JobName:   "job_myjob",
			JobStatus: "DONE",
			Queue:     "normal",
			//SubmitTime: time.Now(),
			User: "lsfadmin",
		},
		{
			Command: "sleep 10000",
			ExHosts: []string{
				"hostB",
			},
			FromHost:   "hostA",
			JobId:      "529",
			JobName:    "jobX34578",
			JobStatus:  "FAILED",
			Queue:      "normal",
			SubmitTime: time.Now(),
			User:       "user1",
		},
	}
}
