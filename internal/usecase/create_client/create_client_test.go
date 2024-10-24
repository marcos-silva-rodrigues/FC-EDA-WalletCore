package createclient

import (
	"testing"

	"github.com/marcos-silva-rodrigues/wallet-ms/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateClientUseCase(t *testing.T) {
	m := &mocks.ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)
	uc := NewCreateClientUseCase(m)
	output, err := uc.Execute(CreateClientInputDTO{
		Name:  "john Doe",
		Email: "j@j",
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)

	assert.NotEmpty(t, "john Doe", output.Name)
	assert.NotEmpty(t, "j@j", output.Email)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
