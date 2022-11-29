package consts

const (
	ErrAddingQuizInDB          = "error while inserting a new quiz"
	ErrAddingQuestionInDB      = "error while inserting a new question"
	ErrInAuthorization         = "error authorizing the request"
	ErrInvalidReqForAddingQns  = "no questions found in the request to add with respect to quiz id"
	ErrInvalidQuizID           = "quiz id can not be empty while requesting to add questions to it"
	ErrAddingAnswerDetails     = "no answer options are passed in the request"
	ErrLinkedQuestionIDMissing = "linked question id for the answer options are missing"
	ErrAddingCorrectAnswerInDB = "error while inserting correct answer for the question"

	SuccessfulQuizCreation       = "successfully created quiz"
	SuccessfulQuestionCreation   = "successfully added question to the quiz"
	SuccessfulCorrectAnsCreation = "successfully added correct answer to the question"

	EmptyString = ""
)
