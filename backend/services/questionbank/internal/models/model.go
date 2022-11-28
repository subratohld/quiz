package models

import cmnErr "github.com/subratohld/quiz/cmnlib/errors"

type Quiz struct {
	tableName struct{} `pg:"quiz,discard_unknown_columns"`

	ID          string `pg:",pk"`
	CreatedBy   string `sql:"created_by,notnull" validate:"required"`
	CreatedOn   int64  `sql:"created_on"`
	UpdatedBy   string `sql:"updated_by"`
	UpdatedOn   int64  `sql:"updated_on"`
	Description string `sql:"description,notnull"`
}

type Auth struct {
	UserID   string
	OrgID    string
	Role     string
	IsSystem bool
	Error    cmnErr.SystemError
}

type Question struct {
	tableName struct{} `pg:"questions,discard_unknown_columns"`

	ID          string           `pg:",pk"`
	Description string           `sql:"description"`
	CreatedBy   string           `sql:"created_by,notnull" validate:"required"`
	CreatedOn   int64            `sql:"created_on"`
	UpdatedBy   string           `sql:"updated_by"`
	UpdatedOn   int64            `sql:"updated_on"`
	Options     []*AnswerOptions `sql:"answer_options"`
}

type AnswerOptions struct {
	OptionId          string
	OptionDescription string
}

type Answer struct {
	tableName struct{} `pg:"answer,discard_unknown_columns"`

	ID          string `pg:",pk"`
	Description string `sql:"description"`
	CreatedBy   string `sql:"created_by,notnull" validate:"required"`
	CreatedOn   int64  `sql:"created_on"`
	UpdatedBy   string `sql:"updated_by"`
	UpdatedOn   int64  `sql:"updated_on"`
	LinkedQnId  string `sql:"linked_qn_id,notnull"`
}
