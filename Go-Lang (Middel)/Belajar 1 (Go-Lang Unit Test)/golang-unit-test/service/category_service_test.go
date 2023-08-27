package service

var categoryRepo := &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService := CategoryService{Repository: categoryRepo}

func TestCategoryService_GetNotFpund(t *tesing.T){
	//program mock
	categoryRepository.Mock.On("FindId","1").return(nil)
	category,err := categoryService.Get("1")

assert.NotNil(t,err)
assert.Nil(t,category)
}
