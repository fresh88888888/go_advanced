package domain

type Issue struct {
	IssueId          int64    `db:"issue_id"`
	IssueTitle       string   `db:"issue_title"`
	IssueDescription string   `db:"issue_description"`
	IssueProjectId   int64    `db:"issue_projectId"`
	IssueOwnerId     int64    `db:"issue_ownerId"`
	IssueStatus      Status   `db:"issue_status"`
	IssuePriority    Priority `db:"issue_priority"`
}

type IssueService interface {
	Issue(id int64) (*Issue, error)
	Issues() ([]*Issue, error)
	Create(issue *Issue) error
	Delete(id int64) error
}

type IssueRepository interface {
	GetById(id int64) (*Issue, error)
	All() ([]*Issue, error)
	Create(issue *Issue) error
	Delete(id int64) error
}
