package delivery

import (
	"net/http"
	"strconv"
	"warehouse/features/favorite"
	"warehouse/middlewares"
	"warehouse/utils/helper"

	"github.com/labstack/echo/v4"
)

type FavoriteDelivery struct {
	favoriteUsecase favorite.UsecaseInterface
}

func New(e *echo.Echo, usecase favorite.UsecaseInterface) {
	handler := &FavoriteDelivery{
		favoriteUsecase: usecase,
	}
	e.POST("/favorite", handler.PostFavorite, middlewares.JWTMiddleware())
	e.GET("/favorite", handler.GetFavorite, middlewares.JWTMiddleware())
	e.DELETE("/favorite/:id", handler.DeleteFavorite, middlewares.JWTMiddleware())
}

func (delivery *FavoriteDelivery) PostFavorite(c echo.Context) error {
	token, role, errToken := middlewares.ExtractToken(c)

	if role != "penitip" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Unautorized"))
	}
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	}

	var dataFav FavRequest

	errBind := c.Bind(&dataFav)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}
	dataFav.UserID = token

	row, err := delivery.favoriteUsecase.PostFavorite(toCore(dataFav))
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}

	return c.JSON(201, helper.SuccessResponseHelper("success insert data"))
}

func (delivery *FavoriteDelivery) GetFavorite(c echo.Context) error {
	token, role, errToken := middlewares.ExtractToken(c)

	if role != "penitip" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Unautorized"))
	}
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	}

	data, err := delivery.favoriteUsecase.GetFavorite(token)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get data"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get data", fromCoreList(data)))
}

func (delivery *FavoriteDelivery) DeleteFavorite(c echo.Context) error {
	token, role, errToken := middlewares.ExtractToken(c)

	if role != "penitip" {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Unautorized"))
	}
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("Invalid token"))
	}

	id := c.Param("id")
	idfav, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	row, err := delivery.favoriteUsecase.DeleteFavorite(token, idfav)
	if err != nil || row != 1 {
		return c.JSON(400, helper.FailedResponseHelper("failed delete"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("success delete"))
}
