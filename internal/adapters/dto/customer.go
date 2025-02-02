package dto

type CustomerDto struct {
	Id    uint32 `param:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
} //@name CustomerDto
