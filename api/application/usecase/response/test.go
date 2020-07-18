package response

type TestGet struct {
	Test string `json:"test"`
}

func NewTestGet(s string) *TestGet {
	return &TestGet{
		Test: s,
	}
}
