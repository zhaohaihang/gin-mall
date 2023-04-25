package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	util "mall/pkg/utils"
	"mall/service"
	"mall/types"
)

// CreateProductHandler 创建商品
func CreateProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			form, _ := ctx.MultipartForm()
			files := form.File["file"]
			userId := ctx.Keys["user_id"].(uint)
			l := service.GetProductSrv()
			resp, err := l.ProductCreate(ctx.Request.Context(), userId, files, &req)
			if err != nil {
				util.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

// ListProducts 商品列表
func ListProductsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			if req.PageSize == 0 {
				req.PageSize = 15
			}
			l := service.GetProductSrv()
			resp, err := l.ProductList(ctx.Request.Context(), &req)
			if err != nil {
				util.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

// ShowProduct 商品详情
func ShowProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetProductSrv()
			resp, err := l.ProductShow(ctx.Request.Context(), &req)
			if err != nil {
				util.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

// 删除商品
func DeleteProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetProductSrv()
			resp, err := l.ProductDelete(ctx.Request.Context(), &req)
			if err != nil {
				util.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

// 更新商品
func UpdateProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetProductSrv()
			resp, err := l.ProductUpdate(ctx.Request.Context(), &req)
			if err != nil {
				util.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

// 搜索商品
func SearchProductsHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			if req.PageSize == 0 {
				req.PageSize = 15
			}
			l := service.GetProductSrv()
			resp, err := l.ProductSearch(ctx.Request.Context(), &req)
			if err != nil {
				util.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

func ListProductImgHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ProductServiceReq

		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验
			l := service.GetProductSrv()
			resp, err := l.ProductImgList(ctx.Request.Context(), &req)
			if err != nil {
				util.LogrusObj.Infoln(err)
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}
