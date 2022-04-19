package surveyInformation

import (
	"fmt"
	// "strconv"
	"survayData/api/helpers"
	"survayData/api/model"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
)

// func getStudentDataDAO()([]model.StudentDataStructure, error){
// 	logginghelper.LogDebug("IN : getStudentDataDAO")
// 	var studentData []model.StudentDataStructure
// 	db, dberr := helpers.GetSQLConnection()
// 	if dberr != nil {
// 		logginghelper.LogError("ERROR in DB CONNECTION", dberr)
// 		return studentData, dberr
// 	}
// 	session := db.NewSession(nil)
// 	rows, readErr := session.Query("SELECT NAME,SIRNAME FROM student_info ;")
// 	if readErr != nil {
// 		// logginghelper.LogError(readErr)
// 		fmt.Println("Error while executing query", readErr)
// 		return studentData, dberr
// 	}
// 	for rows.Next() {
// 		var NAME 					string
// 		var SIRNAME 				string
// 		scanErr := rows.Scan ( &NAME, &SIRNAME)
// 		if scanErr != nil {
// 			fmt.Println("Error in Scanning the rows", scanErr)
// 			// logginghelper.LogError(scanErr)
// 			return nil, scanErr
// 		}
// 		studentData = append(studentData, model.StudentDataStructure{
// 			NAME: 					NAME,
// 			SIRNAME: 				SIRNAME,
// 		})
// 		fmt.Println("studentData: ",studentData)
// 	}
// 	logginghelper.LogDebug("OUT : getStudentDataDAO")
// 	return studentData, nil
// }
func setStudentDataDAO(studentData model.StudentDataStructure) (bool, error) {
	logginghelper.LogDebug("IN :setStudentDataDAO :")
	db, dberr := helpers.GetSQLConnection()
	if dberr != nil {

		logginghelper.LogError("ERROR in DB CONNECTION", dberr)
		return false, dberr
	}
	session := db.NewSession(nil)
	result, err := session.InsertBySql(
		`INSERT INTO registration(EMAIL,USERNAME,PASSWORD)
	   VALUES (?,?,?)`,
		studentData.EMAIL,
		studentData.USERNAME,
		studentData.PASSWORD,
	).Exec()
	if err != nil {
		// logginghelper.LogError(readErr)
		fmt.Println("Error while executing query", err)
		return false, dberr
	}
	fmt.Println(result)
	logginghelper.LogDebug("OUT :setStudentDataDAO")
	return true, nil
}

// ...
func LoginDataDAO(studentData model.StudentDataStructure) (bool, error) {
	logginghelper.LogDebug("IN :LoginDataDAO :")
	db, dberr := helpers.GetSQLConnection()
	if dberr != nil {

		logginghelper.LogError("ERROR in DB CONNECTION", dberr)
		return false, dberr
	}

	session := db.NewSession(nil)
	var result model.LoginDataStructure
	if studentData.USERNAME == "" || studentData.PASSWORD == "" {
		return false, nil
	}
	rows, readErr := session.SelectBySql(`
	SELECT USERNAME,PASSWORD
	FROM registration 
	WHERE USERNAME =? AND PASSWORD=?`,
		studentData.USERNAME,
		studentData.PASSWORD,
	).Load(&result)
	fmt.Println(result)
	if readErr != nil {
		// logginghelper.LogError(readErr)
		fmt.Println("Error while executing query", readErr)
		return false, dberr
	}
	fmt.Println(rows)
	logginghelper.LogDebug("OUT :LoginDataDAO")
	return true, nil
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
// 	  VALUES (?,?,?,?)`,
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
