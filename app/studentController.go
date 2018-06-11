package app

import (
	"database/sql"
	"fmt"
	"github.com/VividCortex/mysqlerr"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

var db = db_init()

type Student struct {
	Id              sql.NullInt64
	Number          sql.NullString
	Name            sql.NullString
	Department_id   sql.NullInt64
	Department_name sql.NullString
	Create_date     time.Time
	Last_updated    time.Time
}

func FetchAllStudents(context *gin.Context) {
	var (
		student  Student
		students []Student
	)

	rows, err := db.Query("select student.id, student.number, student.name, student.department, department.name,  student.create_date, student.last_updated from student, department where student.department = department.id")
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&student.Id, &student.Number, &student.Name, &student.Department_id, &student.Department_name, &student.Create_date, &student.Last_updated)
		checkError(err)
		students = append(students, student)
	}

	context.JSON(http.StatusOK, gin.H{
		"data":   students,
		"count":  len(students),
		"status": "ok",
	})
}

func CreateStudent(context *gin.Context) {
	number := context.PostForm("number")
	name := context.PostForm("name")
	department := context.PostForm("department_id")

	if number == "" || name == "" || department == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "Validation failed",
			"new_id":  nil,
		})
		return
	}

	stmt, err := db.Prepare("insert into student(number, name, department) values (?, ?, ?)")
	checkError(err)
	defer stmt.Close()

	status, err := stmt.Exec(number, name, department)
	if driverErr, ok := err.(*mysql.MySQLError); ok {
		if driverErr.Number == mysqlerr.ER_NO_REFERENCED_ROW_2 {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"new_id":  nil,
				"message": fmt.Sprintf("Invalid department id: %s", department),
			})
			return
		} else {
			log.Fatalln(err.Error())
			return
		}
	}

	insertId, err := status.LastInsertId()
	checkError(err)

	context.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"new_id": insertId,
	})

}

func FetchSingleStudent(context *gin.Context) {
	var (
		student Student
		result  gin.H
	)

	id := context.Params.ByName("id")
	row := db.QueryRow("select student.id, student.number, student.name, student.department, department.name,  student.create_date, student.last_updated from student, department where student.id=? and student.department = department.id", id)
	err := row.Scan(&student.Id, &student.Number, &student.Name, &student.Department_id, &student.Department_name, &student.Create_date, &student.Last_updated)
	if checkError(err) {
		result = gin.H{
			"data":   nil,
			"status": "failed",
		}
	} else {
		result = gin.H{
			"data":   student,
			"status": "ok",
		}
	}

	context.JSON(http.StatusOK, result)
}

func UpdateStudent(context *gin.Context) {
	id := context.Params.ByName("id")

	number := context.PostForm("number")
	name := context.PostForm("name")
	departmentId := context.PostForm("department_id")

	if number == "" || name == "" || departmentId == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":     "failed",
			"message":    "Validation failed",
			"updated_id": nil,
		})
		return
	}

	stmt, err := db.Prepare("update student set number = ?, name = ?, department = ? where id = ?")
	checkError(err)
	defer stmt.Close()

	_, err = stmt.Exec(number, name, departmentId, id)
	if driverErr, ok := err.(*mysql.MySQLError); ok { // How can it be this ?
		//fmt.Println(driverErr.Number, ok)
		if driverErr.Number == mysqlerr.ER_NO_REFERENCED_ROW_2 {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":     "failed",
				"updated_id": nil,
				"message":    fmt.Sprintf("Invalid department id: %s", departmentId),
			})
			return
		} else {
			log.Fatalln(err.Error())
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"status":     "ok",
		"updated_id": id,
	})

}

func DeleteStudent(context *gin.Context) {
	id := context.Params.ByName("id")

	stmt, err := db.Prepare("delete from student where id = ?")
	checkError(err)
	defer stmt.Close()

	status, err := stmt.Exec(id)
	checkError(err)

	count, err := status.RowsAffected()
	checkError(err)

	if count == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"status":        "failed",
			"rows_affected": nil,
			"message":       fmt.Sprintf("Cannot find this student by id: %s", id),
		})
	} else {
		context.JSON(http.StatusNoContent, gin.H{})
	}

}
