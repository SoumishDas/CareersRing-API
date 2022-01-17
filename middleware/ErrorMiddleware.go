package middleware

import (
	"go-gin-api/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

//
// Middleware Error Handler in server package
//
func JSONAppErrorReporter() gin.HandlerFunc {
    return jsonAppErrorReporterT(gin.ErrorTypeAny)
}

func jsonAppErrorReporterT(errType gin.ErrorType) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        detectedErrors := c.Errors.ByType(errType)

        
        if len(detectedErrors) > 0 {
            err := detectedErrors[0].Err
            var parsedError common.AppError
            switch err.(type) {
            case common.AppError:
                parsedError,_ = err.(common.AppError )
				
            default:
                parsedError = common.AppError{ 
                  Code: http.StatusInternalServerError,
                  Message: "Internal Server Error",
                }
				
            }
            // Put the error into response
            c.AbortWithStatusJSON(parsedError.Code, parsedError)
            return
        }

    }
}