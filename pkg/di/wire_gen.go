// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/okekefrancis112/ecommerce-gin/pkg/api"
	"github.com/okekefrancis112/ecommerce-gin/pkg/api/handler"
	"github.com/okekefrancis112/ecommerce-gin/pkg/api/middleware"
	"github.com/okekefrancis112/ecommerce-gin/pkg/config"
	"github.com/okekefrancis112/ecommerce-gin/pkg/db"
	"github.com/okekefrancis112/ecommerce-gin/pkg/repository"
	"github.com/okekefrancis112/ecommerce-gin/pkg/service/cloud"
	"github.com/okekefrancis112/ecommerce-gin/pkg/service/otp"
	"github.com/okekefrancis112/ecommerce-gin/pkg/service/token"
	"github.com/okekefrancis112/ecommerce-gin/pkg/usecase"
)

// Injectors from wire.go:

func InitializeApi(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	authRepository := repository.NewAuthRepository(gormDB)
	tokenService := token.NewTokenService(cfg)
	userRepository := repository.NewUserRepository(gormDB)
	adminRepository := repository.NewAdminRepository(gormDB)
	otpAuth := otp.NewOtpAuth(cfg)
	authUseCase := usecase.NewAuthUseCase(authRepository, tokenService, userRepository, adminRepository, otpAuth)
	authHandler := handler.NewAuthHandler(authUseCase, cfg)
	middlewareMiddleware := middleware.NewMiddleware(tokenService)
	adminUseCase := usecase.NewAdminUseCase(adminRepository, userRepository)
	adminHandler := handler.NewAdminHandler(adminUseCase)
	cartRepository := repository.NewCartRepository(gormDB)
	productRepository := repository.NewProductRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository, cartRepository, productRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	cartUseCase := usecase.NewCartUseCase(cartRepository, productRepository)
	cartHandler := handler.NewCartHandler(cartUseCase)
	paymentRepository := repository.NewPaymentRepository(gormDB)
	orderRepository := repository.NewOrderRepository(gormDB)
	couponRepository := repository.NewCouponRepository(gormDB)
	paymentUseCase := usecase.NewPaymentUseCase(paymentRepository, orderRepository, userRepository, cartRepository, couponRepository, cfg)
	paymentHandler := handler.NewPaymentHandler(paymentUseCase)
	cloudService, err := cloud.NewAWSCloudService(cfg)
	if err != nil {
		return nil, err
	}
	productUseCase := usecase.NewProductUseCase(productRepository, cloudService)
	productHandler := handler.NewProductHandler(productUseCase)
	orderUseCase := usecase.NewOrderUseCase(orderRepository, cartRepository, userRepository, paymentRepository)
	orderHandler := handler.NewOrderHandler(orderUseCase)
	couponUseCase := usecase.NewCouponUseCase(couponRepository, cartRepository)
	couponHandler := handler.NewCouponHandler(couponUseCase)
	offerRepository := repository.NewOfferRepository(gormDB)
	offerUseCase := usecase.NewOfferUseCase(offerRepository)
	offerHandler := handler.NewOfferHandler(offerUseCase)
	stockRepository := repository.NewStockRepository(gormDB)
	stockUseCase := usecase.NewStockUseCase(stockRepository)
	stockHandler := handler.NewStockHandler(stockUseCase)
	brandRepository := repository.NewBrandDatabaseRepository(gormDB)
	brandUseCase := usecase.NewBrandUseCase(brandRepository)
	brandHandler := handler.NewBrandHandler(brandUseCase)
	serverHTTP := http.NewServerHTTP(authHandler, middlewareMiddleware, adminHandler, userHandler, cartHandler, paymentHandler, productHandler, orderHandler, couponHandler, offerHandler, stockHandler, brandHandler)
	return serverHTTP, nil
}
