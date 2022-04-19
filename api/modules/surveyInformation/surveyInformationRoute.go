package surveyInformation

import (
	// "fmt"
	"net/http"
	"survayData/api/model"

	// "fmt"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"github.com/labstack/echo"
)


func Init(o *echo.Group) {
	// o.POST("/getStudentData", getStudentDataRoute)
	o.POST("/setStudentData",setStudentDataRoute)
	o.POST("/LoginData",LoginDataRoute)
}

// func getStudentDataRoute(c echo.Context) error{
// 	logginghelper.LogDebug("IN : getStudentDataRoute")

// 	var studentData model.StudentDataStructure
// 	bindErr := c.Bind(&studentData)
// 	if bindErr != nil {
// 		logginghelper.LogInfo("Error While Binding Data", bindErr)
// 		return c.JSON(http.StatusExpectationFailed, "MAHAVAN_ERRORCODE_ERROR_WHILE_BINDING_DATA")
// 	}

// 	value, err := getStudentDataService()
// 	if err != nil {
// 		logginghelper.LogInfo("Error While Binding Data", bindErr)
// 	}
// 	defer logginghelper.LogDebug("OUT : getStudentDataRoute")
// 	return c.JSON(http.StatusOK, value)
// }
func setStudentDataRoute (c echo.Context) error {
	logginghelper.LogDebug("IN : setStudentDataRoute")

	var studentData model.StudentDataStructure 
	// fmt.Println(&StudentData )
	bindErr := c.Bind(&studentData)
	if bindErr != nil {
		logginghelper.LogInfo("Error While Binding Data", bindErr)
		return c.JSON(http.StatusExpectationFailed, "MAHAVAN_ERRORCODE_ERROR_WHILE_BINDING_DATA")
	}
	value, err := setStudentDataService(studentData)
	if err != nil {
		logginghelper.LogInfo("Error While Binding Data", bindErr)
	}
	defer logginghelper.LogDebug("OUT :setStudentDataRoute")
	return c.JSON(http.StatusOK, value)
}
//....
func LoginDataRoute (c echo.Context) error {
	logginghelper.LogDebug("IN : LoginDataRoute")

	var studentData model.StudentDataStructure 
	// fmt.Println(&StudentData )
	bindErr := c.Bind(&studentData)
	if bindErr != nil {
		logginghelper.LogInfo("Error While Binding Data", bindErr)
		return c.JSON(http.StatusExpectationFailed, "MAHAVAN_ERRORCODE_ERROR_WHILE_BINDING_DATA")
	}
	value, err := LoginDataService(studentData)
	if err != nil {
		logginghelper.LogInfo("Error While Binding Data", bindErr)
	}
	defer logginghelper.LogDebug("OUT :LoginDataRoute")
	return c.JSON(http.StatusOK, value)
}

