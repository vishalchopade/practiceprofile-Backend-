package middleware

import (
	// "Survey/SurveyServer/api/Utils/jwtService"
	// "strings"

	// "corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/confighelper"
	// "corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init(e *echo.Echo, o *echo.Group, r *echo.Group, c *echo.Group, ec *echo.Group) {
	e.Static("/", "./dist")

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:  []string{"*", "localhost"},
		AllowMethods:  []string{echo.GET, echo.PUT, echo.HEAD, echo.PATCH, echo.POST, echo.DELETE},
		ExposeHeaders: []string{echo.HeaderAuthorization},
		// AllowCredentials: true,
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${status} ${method} ${uri} ${latency_human} ${header:all} ${header:*} ${header} ${form} \n",
	}))
	//NOTE: STATTIC IMAGES
	// e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
	// 	// path := getPath()
	// 	Root:   staticServePath,
	// 	Browse: false,
	// }))

	// r.Use(JwtMiddleware())
	// c.Use(JwtMiddleware())
	// ec.Use(JwtMiddlewareForExternalClient())
}



// ******************************************
// // This middleware will be called for every restricted URL
// func JwtMiddleware() echo.MiddlewareFunc {
// 	logginghelper.LogInfo("JwtMiddleware called.............. ")
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			tokenFromRequest := c.Request().Header.Get("Authorization")

// 			if strings.TrimSpace(tokenFromRequest) == "" {
// 				tokenFromRequest = c.QueryParam("qwerty")
// 			}

// 			if tokenFromRequest == "" {
// 				return echo.ErrUnauthorized
// 			}

// 			tokenArray := strings.Split(tokenFromRequest, "Bearer")

// 			if len(tokenArray) <= 1 {
// 				return echo.ErrUnauthorized
// 			}

// 			sessionID, tokenError := jwtService.GetSessionIdFromToken(c)

// 			if tokenError != nil {
// 				logginghelper.LogError("error occured while calling GetSessionIdFromToken ", tokenError)
// 				return echo.ErrUnauthorized
// 			}

// 			maxIdleDuration := confighelper.GetConfig("maxIdleDuration")
// 			serr := mongoSessionManager.SlideSession(sessionID, maxIdleDuration)
// 			if nil != serr {
// 				if serr.Error() == "ERR_SESSION_NOT_FOUND" {
// 					logginghelper.LogError("ERR_SESSION_EXPIRED", serr)
// 					return c.JSON(http.StatusUnauthorized, "ERR_SESSION_EXPIRED")
// 				}
// 				logginghelper.LogError("error occured while fetching value from gcache ", serr)
// 				return c.JSON(http.StatusExpectationFailed, serr.Error())
// 			}
// 			return next(c)
// 		}
// 	}
// }

// //JwtMiddlewareForExternalClient intercepts every request made to the REST APIs that belong to the 'ec' (external client) group
// func JwtMiddlewareForExternalClient() echo.MiddlewareFunc {

// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {

// 			jwtToken := c.QueryParam("qwerty")

// 			if strings.TrimSpace(jwtToken) == "" {
// 				return echo.ErrUnauthorized
// 			}

// 			scope, jwtSecret, sessionSlidingWindowDuration, err := externalClient.FindScopeJWTSecretAndSessionSlidingWindowDurationUsingJWTTokenService(jwtToken)

// 			if err != nil {
// 				logginghelper.LogError(err)
// 				return echo.ErrUnauthorized
// 			}

// 			decodedToken, err := DecodeExternalClientToken(jwtToken, jwtSecret)

// 			if err != nil {
// 				logginghelper.LogError("unable_to_decode_jwt_token_for_external_client", err)
// 				return echo.ErrUnauthorized
// 			}

// 			sessionID := decodedToken["sessionId"].(string)
// 			err = mongoSessionManager.SlideSessionForExternalClient(sessionID, sessionSlidingWindowDuration)

// 			if err != nil {

// 				if err.Error() == "unable_to_find_session_for_external_client" {
// 					logginghelper.LogError("ERR_SESSION_EXPIRED", err)
// 					return c.JSON(http.StatusUnauthorized, "ERR_SESSION_EXPIRED")
// 				}

// 				return c.JSON(http.StatusExpectationFailed, err.Error())
// 			}

// 			requestedURI := getRequestedURI(c)

// 			if externalClientIsAllowedToAccessTheResource(scope, requestedURI) {

// 				//Allow this method to fail silently (in case it does) as logging access statistics isn't as crucial as serving the client's request
// 				externalClient.LogRequestStatisticsForExternalClientService(decodedToken["username"].(string))

// 				return next(c)
// 			}
// 			return echo.ErrForbidden

// 		}
// 	}
// }

// func getRequestedURI(c echo.Context) string {
// 	request := c.Request()
// 	uriArr := strings.Split(request.RequestURI, "ec/")
// 	return strings.Split(uriArr[1], "?")[0]
// }

// func externalClientIsAllowedToAccessTheResource(scope model.Scope, requestedResource string) bool {
// 	for _, resource := range scope.Resources {
// 		if string(resource) == requestedResource {
// 			return true
// 		}
// 	}
// 	return false
// }

// //DecodeExternalClientToken decodes the JWT token sent by an 'external client'. This function was written separately as the token doesn't contain the 'Bearer' substring.
// func DecodeExternalClientToken(tokenFromRequest, jwtKey string) (jwt.MapClaims, error) {

// 	if strings.TrimSpace(tokenFromRequest) == "" {
// 		return nil, errors.New("Provided JWT token is nil or invalid ")
// 	}

// 	token, err := jwt.Parse(tokenFromRequest, func(token *jwt.Token) (interface{}, error) {
// 		_, ok := token.Method.(*jwt.SigningMethodHMAC)

// 		if !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return []byte(jwtKey), nil
// 	})

// 	if err != nil {
// 		logginghelper.LogError("Error while parsing JWT Token: ", err)
// 		return nil, err
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)

// 	if !ok {
// 		return nil, errors.New("Error while getting claims")
// 	}

// 	return claims, nil
// }
// *************************************************