package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/middlewares/middlewares"
	"github.com/srv-cashpay/pos/configs"
	h_history "github.com/srv-cashpay/pos/handlers/history"
	h_pos "github.com/srv-cashpay/pos/handlers/pos"
	r_history "github.com/srv-cashpay/pos/repositories/history"
	r_pos "github.com/srv-cashpay/pos/repositories/pos"
	s_history "github.com/srv-cashpay/pos/services/history"
	s_pos "github.com/srv-cashpay/pos/services/pos"
)

var (
	DB   = configs.InitDB()
	JWT  = middlewares.NewJWTService()
	posR = r_pos.NewPosRepository(DB)
	posS = s_pos.NewPosService(posR, JWT)
	posH = h_pos.NewPosHandler(posS)

	historyR = r_history.NewHistoryRepository(DB)
	historyS = s_history.NewHistoryService(historyR, JWT)
	historyH = h_history.NewHistoryHandler(historyS)
)

func New() *echo.Echo {
	e := echo.New()

	pos := e.Group("/api/pos", middlewares.AuthorizeJWT(JWT))
	{
		pos.POST("/create", posH.Create)
		pos.PUT("/update/:id", posH.Update)
		pos.GET("/:id", posH.GetById)
	}

	history := e.Group("/api/history", middlewares.AuthorizeJWT(JWT))
	{
		history.GET("/pagination", historyH.Get)
		history.GET("/:id", historyH.GetById)
	}

	return e
}
