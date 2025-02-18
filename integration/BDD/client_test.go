package main_test

import (
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Clientes", func() {
	Context("Dado que o usuário consulta seu acesso pelo CPF", func() {
		When("quando o usuário existe", func() {
			BeforeEach(func() {
				req, _ := http.NewRequest(http.MethodPost, "http://localhost:8000/v1/customer", strings.NewReader(`{"name": "Teste nome","email": "teste@email.com","cpf": "12345678901"}`))
				req.Header.Set("Content-Type", "application/json")
				_, err := http.DefaultClient.Do(req)
				assert.NoError(GinkgoT(), err)
			})

			It("deve retornar os dados do cliente corretamente", func() {
				req, _ := http.NewRequest(http.MethodGet, "http://localhost:8000/v1/customer/cpf/12345678901", nil)
				req.Header.Set("Content-Type", "application/json")
				res, err := http.DefaultClient.Do(req)

				assert.NoError(GinkgoT(), err)
				assert.Equal(GinkgoT(), http.StatusOK, res.StatusCode)
			})
		})
	})
})
