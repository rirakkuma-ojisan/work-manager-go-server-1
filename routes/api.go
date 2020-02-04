package routes

import(
  "work-manager-go-server-1/web/api"
  "github.com/labstack/echo"
)

func Init(e*echo.Echo) {
  g := e.Group("/api")
  {
    g.GET("/monthly_attendance", api.FetchMonthlyAttendance())
  }
}
