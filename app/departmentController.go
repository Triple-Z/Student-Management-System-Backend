package app

import (
	"database/sql"
	"time"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

type Department struct {
	Id sql.NullInt64
	Name sql.NullString
	Create_date time.Time
	Last_updated time.Time
}


func FetchAllDepartments(c *gin.Context) {
	var (
		department Department
		departments []Department
	)

	rows, err := db.Query("select id, name, create_date, last_updated from department")
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&department.Id, &department.Name, &department.Create_date, &department.Last_updated)
		checkError(err)
		departments = append(departments, department)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": departments,
		"count": len(departments),
		"status": "ok",
	})
}


func CreateDepartment(c *gin.Context) {
	name := c.PostForm("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "Validation failed",
			"new_id": nil,
		})
		return
	}

	stmt, err := db.Prepare("insert into department(name) values (?)")
	checkError(err)
	defer stmt.Close()

	status, err := stmt.Exec(name)
	checkError(err)

	insert_id, err := status.LastInsertId()
	checkError(err)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"new_id": insert_id,
	})
}


func FetchSingleDepartment(c *gin.Context) {
	var result gin.H

	id := c.Params.ByName("id")

	var department Department

	row := db.QueryRow("select id, name, create_date, last_updated from department where id = ?", id)

	err := row.Scan(&department.Id, &department.Name, &department.Create_date, &department.Last_updated)

	if checkError(err) {
		result = gin.H{
			"status": "failed",
			"data": nil,
		}
	} else {
		result = gin.H{
			"status": "ok",
			"data": department,
		}
	}

	c.JSON(http.StatusOK, result)
}


func UpdateDeprtment(c *gin.Context) {
	id := c.Params.ByName("id")

	name := c.PostForm("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"message": "Validation failed",
			"new_id": nil,
		})
		return
	}

	stmt, err := db.Prepare("update department set name = ? where id = ?")
	checkError(err)
	defer stmt.Close()

	_, err = stmt.Exec(name, id)
	checkError(err)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"updated_id": id,
	})

}


func DeleteDepartment(c *gin.Context) {

	id := c.Params.ByName("id")

	stmt, err := db.Prepare("delete from department where id = ?")
	checkError(err)
	defer stmt.Close()

	status, err := stmt.Exec(id)
	checkError(err)

	count, err := status.RowsAffected()
	checkError(err)

	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "failed",
			"rows_affected": nil,
			"message": fmt.Sprintf("Cannot find this department by id: %s", id),
		})
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}

}
