package handler

import (
	"github.com/gofiber/fiber/v2"
	"hp-hotel-rest/internal/model"
	"hp-hotel-rest/internal/service"
	"hp-hotel-rest/pkg/utils"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register @Summary Register a new user
// @Description Register a new user with the provided credentials
// @Tags auth
// @Accept json
// @Produce json
// @Param input body model.LoginRequest true "User credentials"
// @Success 201 {object} utils.APIResponse{data=nil} "User registered successfully"
// @Failure 400 {object} utils.APIResponse{} "Bad Request"
// @Failure 500 {object} utils.APIResponse{} "Internal Server Error"
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var userRegisterRequest model.LoginRequest
	if err := c.BodyParser(&userRegisterRequest); err != nil {
		return utils.RespondJSON(c, fiber.StatusBadRequest, "Bad Request", nil)
	}

	err := h.authService.Register(userRegisterRequest)
	if err != nil {
		return utils.RespondJSON(c, fiber.StatusInternalServerError, "Internal Server Error", nil)
	}

	return utils.RespondJSON(c, fiber.StatusCreated, "User registered successfully", nil)
}

// Login @Summary Authenticate user
// @Description Authenticate user with provided credentials and generate JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param input body model.LoginRequest true "User credentials"
// @Success 200 {object} utils.APIResponse{data=string} "User logged in successfully"
// @Failure 400 {object} utils.APIResponse{} "Bad Request"
// @Failure 401 {object} utils.APIResponse{} "Invalid credentials"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var loginRequest model.LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return utils.RespondJSON(c, fiber.StatusBadRequest, "Bad Request", nil)
	}

	token, err := h.authService.Login(loginRequest)
	if err != nil {
		return utils.RespondJSON(c, fiber.StatusUnauthorized, "Invalid credentials", nil)
	}

	return utils.RespondJSON(c, fiber.StatusOK, "User logged in successfully", fiber.Map{"token": token})
}

// ProtectedRoute @Summary Protected route
// @Description A protected route that requires authentication
// @Tags auth
// @Security Bearer
// @Accept json
// @Produce json
// @Success 200 {object} utils.APIResponse{data=nil} "Protected Route"
// @Router /auth/protected [get]
func (h *AuthHandler) ProtectedRoute(c *fiber.Ctx) error {
	return utils.RespondJSON(c, fiber.StatusOK, "Protected Route", nil)
}
