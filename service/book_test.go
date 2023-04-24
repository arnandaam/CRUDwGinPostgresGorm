package service

import (
	"errors"
	"showcase/model"
	"showcase/repository/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_BookService_GetBookbyID(t *testing.T) {
	type testCase struct {
		name           string
		expectedResult model.Book
		expectError    error
		wantError      bool
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookbyID(gomock.Any()).Return(model.Book{
				ID:     1,
				Name:   "yanto",
				Author: "bukasitik",
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:     1,
			Name:   "yanto",
			Author: "bukasitik",
		},
	})
	testTable = append(testTable, testCase{
		name:        "record Error",
		wantError:   true,
		expectError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookbyID(gomock.Any()).Return(model.Book{}, errors.New("record not found")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)
			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}
			service := Service{
				repo: bookRepo,
			}

			res, err := service.GetBookbyID(1)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}

}

func Test_BookService_GetAllBooks(t *testing.T) {
	type testCase struct {
		name           string
		expectedResult []model.Book
		expectError    error
		wantError      bool
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetAllBooks().Return([]model.Book{
				{
					ID:     1,
					Name:   "yanto",
					Author: "bukasitik",
				},
				{
					ID:     2,
					Name:   "pandu",
					Author: "bukasitik",
				},
			}, nil).Times(1)
		},
		expectedResult: []model.Book{
			{
				ID:     1,
				Name:   "yanto",
				Author: "bukasitik",
			},
			{
				ID:     2,
				Name:   "pandu",
				Author: "bukasitik",
			},
		},
	})
	testTable = append(testTable, testCase{
		name:        "record Error",
		wantError:   true,
		expectError: errors.New("no records found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetAllBooks().Return(nil, errors.New("no records found")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)
			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}
			service := Service{
				repo: bookRepo,
			}

			res, err := service.GetAllBooks()

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func Test_BookService_CreateBook(t *testing.T) {
	type testCase struct {
		name           string
		input          model.Book
		expectedResult model.Book
		expectError    error
		wantError      bool
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name: "success",
		input: model.Book{
			Name:   "yanto",
			Author: "bukasitik",
		},
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Book{
				ID:     1,
				Name:   "yanto",
				Author: "bukasitik",
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:     1,
			Name:   "yanto",
			Author: "bukasitik",
		},
	})

	testTable = append(testTable, testCase{
		name: "record Error",
		input: model.Book{
			Name:   "yanto",
			Author: "bukasitik",
		},
		wantError:   true,
		expectError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Book{}, errors.New("record not found")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)
			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}
			service := Service{
				repo: bookRepo,
			}

			res, err := service.CreateBook(testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}

}

func Test_BookService_UpdateBookid(t *testing.T) {
	type testCase struct {
		name           string
		input          model.Book
		expectedResult model.Book
		expectError    error
		wantError      bool
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name: "success",
		input: model.Book{
			ID:     1,
			Name:   "yanto",
			Author: "bukasitik",
		},
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().UpdateBookid(gomock.Any()).Return(model.Book{
				ID:     1,
				Name:   "yanto",
				Author: "bukasitik",
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:     1,
			Name:   "yanto",
			Author: "bukasitik",
		},
	})

	testTable = append(testTable, testCase{
		name: "record Error",
		input: model.Book{
			ID:     1,
			Name:   "yanto",
			Author: "bukasitik",
		},
		wantError:   true,
		expectError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().UpdateBookid(gomock.Any()).Return(model.Book{}, errors.New("record not found")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)
			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}
			service := Service{
				repo: bookRepo,
			}

			res, err := service.UpdateBookid(testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}

}

func Test_BookService_DeleteBookid(t *testing.T) {
	type testCase struct {
		name        string
		expectedErr error
		wantError   bool
		onBookRepo  func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().DeleteBookid(gomock.Any()).Return(nil).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:        "error",
		wantError:   true,
		expectedErr: errors.New("failed to delete book"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().DeleteBookid(gomock.Any()).Return(errors.New("failed to delete book")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)
			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}
			service := Service{
				repo: bookRepo,
			}

			err := service.DeleteBookid(1)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
