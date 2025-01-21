package structures

type User struct {
	Name   string `json:"name" validate:"required,min=3,max=100"`
	Age    int    `json:"age" validate:"required,min=0,max=99"`
	Gender string `json:"gender" validate:"required"`
}

type Admin struct {
	Inh         User
	IsSuperUser bool
}
