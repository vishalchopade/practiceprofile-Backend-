package surveyInformation

import (
	"survayData/api/model"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
)

// func getStudentDataService()([]model.StudentDataStructure,error){
// 	logginghelper.LogDebug("IN : getStudentDataService")

// 	value,err := setStudentDataDAO()
// 	if err != nil {
// 		return false, err
// 	}

// 	logginghelper.LogDebug("OUT : getStudentDataService")
// 	return value, err
// }
func setStudentDataService(studentData model.StudentDataStructure) (bool, error){
	logginghelper.LogDebug("IN : setStudentDataService")

	value,err := setStudentDataDAO(studentData)
	if err != nil {
		return false, err
	}
	logginghelper.LogDebug("OUT : setStudentDataService")
	return value, err
}
// ...
func LoginDataService(studentData model.StudentDataStructure) (bool, error){
	logginghelper.LogDebug("IN : LoginDataService")

	value,err := LoginDataDAO(studentData)
	if err != nil {
		return false, err
	}
	logginghelper.LogDebug("OUT : LoginDataService")
	return value, err
}
// // ....
// func setStudentDataDAO(studentData model.StudentDataStructure ) (bool, error) {
// 	logginghelper.LogDebug("IN :setStudentDataDAO :")
// 	db, dberr := helpers.GetSQLConnection()
// 	if dberr != nil {

// 		logginghelper.LogError("ERROR in DB CONNECTION", dberr)
// 		return false, dberr
// 	}
// 	session := db.NewSession(nil)
// 	result, err := session.InsertBySql(
// 		`INSERT INTO registration(Id,EMAIL,USERNAME,PASSWORD)
// 	  VALUES (?,?,?)`,
// 	  studentData.Id,
// 	  studentData.EMAIL,
// 	  studentData.USERNAME,
// 	  studentData.PASSWORD,
// 	).Exec()
// 	if err != nil {
// 		// logginghelper.LogError(readErr)
// 		fmt.Println("Error while executing query", err)
// 		return false, dberr
// 	}
// 	fmt.Println(result)
// 	logginghelper.LogDebug("OUT :setStudentDataDAO")
// 	return true, nil
// }
