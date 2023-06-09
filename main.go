package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id          int
	ProjectName string
	StartDate   string
	EndDate     string
	Duration    string
	Description string
	postingTime string
	Html        bool
	Css         bool
	Javascript  bool
	Java        bool
}

var dataProject = []Project{
	{
		ProjectName: "ex 1",
		StartDate:   "12-02",
		EndDate:     "",
		Duration:    "3 Months",
		Html:        false,
		Css:         true,
		Javascript:  true,
		Java:        false,
		postingTime: "23",
		Description: "lorem lorem",
	},
	{
		ProjectName: "ex 2",
		StartDate:   "11-01",
		EndDate:     "22-01",
		Duration:    "3 Months",
		Html:        true,
		Css:         true,
		Javascript:  false,
		Java:        true,
		postingTime: "23",
		Description: "blablabalala",
	},
}

func main() {
	e := echo.New()

	e.Static("/public", "public")
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/detailproject/:id", detailproject)
	e.GET("/myproject", myProject)
	e.GET("/testimoni", testimoni)
	e.POST("/add-Project", addProject)
	e.POST("/deleteProject/:id", deleteProject)
	e.POST("/edit-project/:id", ressEditProject)
	e.GET("/edit-project/:id", editProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	projects := map[string]interface{}{
		"Projects": dataProject,
	}

	return tmpl.Execute(c.Response(), projects)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func testimoni(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimoni.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func myProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/myproject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func detailproject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var DetailProject = Project{}

	for i, data := range dataProject {
		if id == i {
			DetailProject = Project{
				ProjectName: data.ProjectName,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Duration:    data.Duration,
				Html:        data.Html,
				Css:         data.Css,
				Javascript:  data.Javascript,
				Java:        data.Java,
				Description: data.Description,
			}
		}
	}

	data := map[string]interface{}{
		"Project": DetailProject,
	}

	var tmpl, err = template.ParseFiles("views/detailproject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func addProject(c echo.Context) error {
	projectName := c.FormValue("input-name")
	startDate := c.FormValue("input-start")
	endDate := c.FormValue("input-end")
	description := c.FormValue("input-description")
	html := c.FormValue("input-check-html")
	css := c.FormValue("input-check-css")
	javascript := c.FormValue("input-check-javascript")
	java := c.FormValue("input-check-java")
	postingTime := time.Now()
	// konversi value cekbox, string to boolean
	htmlValue := html != ""
	cssValue := css != ""
	javascriptValue := javascript != ""
	javaValue := java != ""

	var newProject = Project{
		ProjectName: projectName,
		StartDate:   startDate,
		EndDate:     endDate,
		Duration:    getDuration(startDate, endDate),
		Description: description,
		Html:        htmlValue,
		Css:         cssValue,
		Javascript:  javascriptValue,
		Java:        javaValue,
		postingTime: time.Now().String(),
	}

	dataProject = append(dataProject, newProject)

	fmt.Println(projectName, "posted at :", postingTime)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func editProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	for i, data := range dataProject {
		if id == i {
			ProjectDetail = Project{
				Id:          id,
				ProjectName: data.ProjectName,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Duration:    data.Duration,
				Description: data.Description,
				Html:        data.Html,
				Css:         data.Css,
				Javascript:  data.Javascript,
				Java:        data.Java,
			}
		}
	}

	data := map[string]interface{}{
		"Project": ProjectDetail,
	}

	var tmpl, err = template.ParseFiles("views/edit-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func ressEditProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Index :", id)

	projectName := c.FormValue("input-name")
	startDate := c.FormValue("input-start")
	endDate := c.FormValue("input-end")
	description := c.FormValue("input-description")
	html := c.FormValue("input-check-html")
	css := c.FormValue("input-check-css")
	javascript := c.FormValue("input-check-javascript")
	java := c.FormValue("input-check-java")
	postingTime := time.Now()

	// konversi cekbox string to boolean
	htmlValue := html != ""
	cssValue := css != ""
	javascriptValue := javascript != ""
	javaValue := java != ""

	var ressEditProject = Project{
		ProjectName: projectName,
		StartDate:   startDate,
		EndDate:     endDate,
		Duration:    getDuration(startDate, endDate),
		Description: description,
		Html:        htmlValue,
		Css:         cssValue,
		Javascript:  javascriptValue,
		Java:        javaValue,
		postingTime: time.Now().String(),
	}

	fmt.Println(projectName, "Update at :", postingTime)

	dataProject[id] = ressEditProject

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Index : ", id)

	dataProject = append(dataProject[:id], dataProject[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func getDuration(startDate, endDate string) string {
	startTime, _ := time.Parse("2006-01-02", startDate)
	endTime, _ := time.Parse("2006-01-02", endDate)

	durationTime := int(endTime.Sub(startTime).Hours())
	durationDays := durationTime / 24
	durationWeeks := durationDays / 7
	durationMonths := durationWeeks / 4
	durationYears := durationMonths / 12

	var duration string

	if durationYears > 1 {
		duration = strconv.Itoa(durationYears) + " years"
	} else if durationYears > 0 {
		duration = strconv.Itoa(durationYears) + " year"
	} else {
		if durationMonths > 1 {
			duration = strconv.Itoa(durationMonths) + " months"
		} else if durationMonths > 0 {
			duration = strconv.Itoa(durationMonths) + " month"
		} else {
			if durationWeeks > 1 {
				duration = strconv.Itoa(durationWeeks) + " weeks"
			} else if durationWeeks > 0 {
				duration = strconv.Itoa(durationWeeks) + " week"
			} else {
				if durationDays > 1 {
					duration = strconv.Itoa(durationDays) + " days"
				} else {
					duration = strconv.Itoa(durationDays) + " day"
				}
			}
		}
	}

	return duration

}
