package checkout_handlers

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : warung-pintar/checkout-service
 */

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mockDTOitem "warung-pintar/checkout-service/mocks/app/dtos/checkout"
	mockUc "warung-pintar/checkout-service/mocks/app/use_cases"
	mockResp "warung-pintar/checkout-service/mocks/interface/rest/response"
	dtoItem "warung-pintar/checkout-service/src/app/dtos/checkout"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type MockCheckoutHandler struct {
	mock.Mock
}

type HandlerTest struct {
	suite.Suite
	mockUseCase *mockUc.MockCheckoutUseCase

	mockResponse *mockResp.MockResponse
	h            CheckoutHandlerInterface
	w            *httptest.ResponseRecorder
	dtoTest      *dtoItem.CheckoutReqDTO
	dtoTestFail  *dtoItem.CheckoutReqDTO
	mockDTO      *mockDTOitem.MockCheckoutDTO
}

func (suite *HandlerTest) SetupTest() {
	suite.mockUseCase = new(mockUc.MockCheckoutUseCase)
	suite.mockResponse = new(mockResp.MockResponse)
	suite.mockDTO = new(mockDTOitem.MockCheckoutDTO)
	suite.h = NewCheckoutHandler(suite.mockResponse, suite.mockUseCase)
	suite.w = httptest.NewRecorder()
	suite.dtoTest = &dtoItem.CheckoutReqDTO{
		Data: []*dtoItem.SkuReqDTO{
			&dtoItem.SkuReqDTO{
				Sku: "Test",
				Qty: 3,
			},
		},
	}
	suite.dtoTestFail = nil
}

func (s *HandlerTest) TestCheckoutSuccess() {
	bodyBytes, _ := json.Marshal(s.dtoTest)

	r := httptest.NewRequest("POST", "/checkout/", bytes.NewBuffer(bodyBytes))
	s.mockDTO.Mock.On("Validate").Return(nil)
	s.mockUseCase.Mock.On("Checkout", s.dtoTest).Return(1000, nil)
	s.mockResponse.Mock.On("JSON", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	http.HandlerFunc(s.h.Checkout).ServeHTTP(s.w, r)

	s.Equal(200, s.w.Result().StatusCode)

}

func (s *HandlerTest) TestRelicFailDTO() {
	r := httptest.NewRequest("POST", "/checkout/", nil)
	s.mockDTO.Mock.On("Validate").Return(errors.New(mock.Anything))
	s.mockResponse.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.Checkout).ServeHTTP(s.w, r)

	s.Equal(400, s.w.Result().StatusCode)

}

func (s *HandlerTest) TestRelicFailValidateDTO() {
	bodyBytes, _ := json.Marshal(s.dtoTestFail)
	r := httptest.NewRequest("POST", "/checkout/", bytes.NewBuffer(bodyBytes))
	s.mockDTO.Mock.On("Validate").Return(errors.New(mock.Anything))
	s.mockResponse.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.Checkout).ServeHTTP(s.w, r)

	s.Equal(400, s.w.Result().StatusCode)

}

func (s *HandlerTest) TestCheckoutFailNotFound() {
	bodyBytes, _ := json.Marshal(s.dtoTest)

	r := httptest.NewRequest("POST", "/checkout/", bytes.NewBuffer(bodyBytes))
	s.mockDTO.Mock.On("Validate").Return(nil)
	s.mockUseCase.Mock.On("Checkout", s.dtoTest).Return(0, errors.New(gorm.ErrRecordNotFound.Error()))
	s.mockResponse.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.Checkout).ServeHTTP(s.w, r)

	s.Equal(404, s.w.Result().StatusCode)

}

func (s *HandlerTest) TestCheckoutFailRetrieve() {
	bodyBytes, _ := json.Marshal(s.dtoTest)

	r := httptest.NewRequest("POST", "/checkout/", bytes.NewBuffer(bodyBytes))
	s.mockDTO.Mock.On("Validate").Return(nil)
	s.mockUseCase.Mock.On("Checkout", s.dtoTest).Return(0, errors.New(mock.Anything))
	s.mockResponse.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.Checkout).ServeHTTP(s.w, r)

	s.Equal(500, s.w.Result().StatusCode)

}

func TestHandler(t *testing.T) {
	suite.Run(t, new(HandlerTest))
}
