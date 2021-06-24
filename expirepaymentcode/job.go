package expirecode

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
	j.service.FindExpiredPaymentCode()
}
