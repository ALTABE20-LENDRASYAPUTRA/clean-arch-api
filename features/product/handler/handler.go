package handler

import (
	"clean-arch-api/app/middlewares"
	"clean-arch-api/features/product"
	"clean-arch-api/utils/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService product.ProductServiceInterface
}

func New(ps product.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{
		productService: ps,
	}
}

func (handler *ProductHandler) CreateProduct(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	newProduct := ProductRequest{}
	errBind := c.Bind(&newProduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	productCore := RequestToCore(newProduct)
	productCore.UserID = userID

	errInsert := handler.productService.Create(productCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success insert data", nil))
}

func (handler *ProductHandler) GetAllProduct(c echo.Context) error {
	results, errSelect := handler.productService.GetAll()
	if errSelect != nil {

		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error read data. "+errSelect.Error(), nil))
	}
	// proses mapping dari core ke response
	productsResult := CoreToResponseList(results)

	return c.JSON(http.StatusOK, responses.WebResponse("success read data.", productsResult))
}

func (handler *ProductHandler) UpdateProduct(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error parsing product id", nil))
	}

	updateProduct := ProductRequest{}
	errBind := c.Bind(&updateProduct)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	productCore := RequestToCore(updateProduct)
	productCore.ID = uint(productID)
	productCore.UserID = userID

	errUpdate := handler.productService.Update(userID, productCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error update data", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success update data", nil))
}

func (handler *ProductHandler) DeleteProduct(c echo.Context) error {
	userID := middlewares.ExtractTokenUserId(c)
    if userID == 0 {
        return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
    }

	productID, err := strconv.Atoi(c.Param("product_id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, responses.WebResponse("error parsing product id", nil))
    }

	errDelete := handler.productService.Delete(userID, productID)
    if errDelete != nil {
        return c.JSON(http.StatusInternalServerError, responses.WebResponse("error delete data", nil))
    }

    return c.JSON(http.StatusOK, responses.WebResponse("success delete data", nil))
}

func (handler *ProductHandler) GetProductByIdUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error parsing user id", nil))
	}

	products, err := handler.productService.GettProductUser(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error get data", nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success get data", CoreToResponseUserList(products)))
}