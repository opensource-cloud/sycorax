package domain

type (
	NestedStruct struct {
		OtherProp string `json:"other_prop" binding:"required"`
	}
	TestStruct struct {
		Property string       `json:"property" binding:"required"`
		Nested   NestedStruct `json:"nested" binding:"required"`
	}
	CreateQueueDTO struct {
		Name string      `json:"name" binding:"required"`
		Test *TestStruct `json:"test" binding:"required"`
	}
)
