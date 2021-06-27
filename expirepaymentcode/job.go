package expirecode

import (
	"fmt"
	"time"
)

type Job struct {
	service Service
}

func NewJob(service Service) Job {
	return Job{
		service: service,
	}
}

func NewDefaultJob(svc Service) Job {
	return NewJob(svc)
}

func (j Job) Execute() {
	// Imitating a CRON Execution
	for i := 0; i < 5; i++ {
		affected, err := j.service.ExpiringPaymentCode()
		if err != nil {
			fmt.Println(fmt.Errorf("Error: %v", err))
		}
		fmt.Printf("Job Done, total expired payment code %v\n", affected)

		time.Sleep(1 * time.Second)
	}
}
