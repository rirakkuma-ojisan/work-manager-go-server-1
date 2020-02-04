package api

import (
  "github.com/labstack/echo"
  "github.com/valyala/fasthttp"
)

func FetchMonthlyAttendance() echo.HandlerFunc {
  return func(c echo.Context) error{
    return c.JSON(fasthttp.StatusOK, "GO ECHO SERVER TEST(monthly_attendance.go)")
  }
}
