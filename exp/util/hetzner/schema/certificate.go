package schema

import "time"

//

type CertificateType string

const (
	CertificateTypeUploaded CertificateType = "uploaded"
	CertificateTypeManaged  CertificateType = "managed"
)

//

type CertificateStatusType string

const (
	CertificateStatusTypePending     CertificateStatusType = "pending"
	CertificateStatusTypeFailed      CertificateStatusType = "failed"
	CertificateStatusTypeCompleted   CertificateStatusType = "completed"
	CertificateStatusTypeScheduled   CertificateStatusType = "scheduled"
	CertificateStatusTypeUnavailable CertificateStatusType = "unavailable"
)

//

type CertificateUsedByRefType string

const (
	CertificateUsedByRefTypeLoadBalancer CertificateUsedByRefType = "load_balancer"
)

//

type CertificateUsedByRef struct {
	Id int `json:"id"`

	Type CertificateUsedByRefType `json:"type"`
}

//

type CertificateStatus struct {
	Issuance CertificateStatusType `json:"issuance"`
	Renewal  CertificateStatusType `json:"renewal"`

	Error *Error `json:"error"`
}

//

type Certificate struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	Created time.Time         `json:"created"`
	Labels  map[string]string `json:"labels"`

	Type CertificateType `json:"type"`

	Certificate string `json:"certificate"`
	Fingerprint string `json:"fingerprint"`

	NotValidBefore time.Time `json:"not_valid_before"`
	NotValidAfter  time.Time `json:"not_valid_after"`

	DomainNames []string `json:"domain_names"`

	Status *CertificateStatus     `json:"status"`
	UsedBy []CertificateUsedByRef `json:"used_by"`
}
