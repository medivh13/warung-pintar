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

	mockDTOitem "warung-pintar/promo-service/mocks/app/dtos/items"
	mockUc "warung-pintar/promo-service/mocks/app/use_cases"
	mockResp "warung-pintar/promo-service/mocks/interface/rest/response"
	dtoItem "warung-pintar/promo-service/src/app/dtos/items"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type MockCheckoutHandler struct {
	mock.Mock
}

type HandlerTest struct {
	suite.Suite
	mockUseCase *mockUc.MockItemUseCase

	mockResponse *mockResp.MockResponse
	h            ItemHandlerInterface
	w            *httptest.ResponseRecorder
	dtoTest      *dtoItem.ItemReqDTO
	dtoTestFail  *dtoItem.ItemReqDTO
	mockDTO      *mockDTOitem.MockItemDTO
}

func (suite *HandlerTest) SetupTest() {
	suite.mockUseCase = new(mockUc.MockItemUseCase)
	suite.mockResponse = new(mockResp.MockResponse)
	suite.mockDTO = new(mockDTOitem.MockItemDTO)
	suite.h = NewItemHandler(suite.mockResponse, suite.mockUseCase)
	suite.w = httptest.NewRecorder()
	suite.dtoTest = &dtoItem.ItemReqDTO{
		Data: []*dtoItem.DataItemReqDTO{
			&dtoItem.DataItemReqDTO{
				Sku: "Test",
			},
		},
	}
	suite.dtoTestFail = nil
}

func (s *HandlerTest) TestCheckoutSuccess() {
	bodyBytes, _ := json.Marshal(s.dtoTest)

	r := httptest.NewRequest("POST", "/promos/", bytes.NewBuffer(bodyBytes))
	s.mockDTO.Mock.On("Validate").Return(nil)
	s.mockUseCase.Mock.On("ListItems", s.dtoTest).Return(1000, nil)
	s.mockResponse.Mock.On("JSON", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	http.HandlerFunc(s.h.GetItems).ServeHTTP(s.w, r)

	s.Equal(200, s.w.Result().StatusCode)

}

func (s *HandlerTest) TestRelicFailDTO() {
	r := httptest.NewRequest("POST", "/promos/", nil)
	s.mockDTO.Mock.On("Validate").Return(errors.New(mock.Anything))
	s.mockResponse.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.GetItems).ServeHTTP(s.w, r)

	s.Equal(400, s.w.Result().StatusCode)

}

func (s *HandlerTest) TestRelicFailValidateDTO() {
	bodyBytes, _ := json.Marshal(s.dtoTestFail)
	r := httptest.NewRequest("POST", "/promos/", bytes.NewBuffer(bodyBytes))
	s.mockDTO.Mock.On("Validate").Return(errors.New(mock.Anything))
	s.mockResponse.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.GetItems).ServeHTTP(s.w, r)

	s.Equal(400, s.w.Result().StatusCode)

}

func (s *HandlerTest) TestCheckoutFailNotFound() {
	bodyBytes, _ := json.Marshal(s.dtoTest)

	r := httptest.NewRequest("POST", "/promos/", bytes.NewBuffer(bodyBytes))
	s.mockDTO.Mock.On("Validate").Return(nil)
	s.mockUseCase.Mock.On("ListItems", s.dtoTest).Return(0, errors.New(gorm.ErrRecordNotFound.Error()))
	s.mockResponse.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.GetItems).ServeHTTP(s.w, r)

	s.Equal(404, s.w.Result().StatusCode)

}

func (s *HandlerTest) TestCheckoutFailRetrieve() {
	bodyBytes, _ := json.Marshal(s.dtoTest)

	r := httptest.NewRequest("POST", "/promos/", bytes.NewBuffer(bodyBytes))
	s.mockDTO.Mock.On("Validate").Return(nil)
	s.mockUseCase.Mock.On("ListItems", s.dtoTest).Return(0, errors.New(mock.Anything))
	s.mockResponse.Mock.On("HttpError", mock.Anything, mock.Anything).Return(mock.Anything)

	http.HandlerFunc(s.h.GetItems).ServeHTTP(s.w, r)

	s.Equal(500, s.w.Result().StatusCode)

}

func TestHandler(t *testing.T) {
	suite.Run(t, new(HandlerTest))
}
