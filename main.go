package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"revisiweb/connection"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct{
	Id          int
	ProjectName string
	StartDate   string
	EndDate     string
	Duration    string
	Description string
	Javascript  bool
	Golang 			bool
	ReactJs 		bool
	Java 				bool
}

var Blogs = []Project{
	// { Id: 0,
	// 	ProjectName: "project 1",
	// 	StartDate: "2023-07-20",
	// 	EndDate: "2023-08-21",
	// 	Duration: "1 Bulan",
	// 	Description: "test 1",
	// 	Javascript: true,
	// 	Golang: true,
	// 	ReactJs: false,
	// 	Java: true,

	// },
	// {
	// 	Id: 1,
	// 	ProjectName: "project 2",
	// 	StartDate: "2023-07-20",
	// 	EndDate: "2023-08-21",
	// 	Duration: "1 Bulan",
	// 	Description: "test 2",
	// 	Javascript: true,
	// 	Golang: true,
	// 	ReactJs: true,
	// 	Java: true,
	// },
	// { Id: 2,
	// 	ProjectName: "project 3",
	// 	StartDate: "2023-07-20",
	// 	EndDate: "2023-08-21",
	// 	Duration: "1 Bulan",
	// 	Description: "test 3",
	// 	Javascript: true,
	// 	Golang: true,
	// 	ReactJs: false,
	// 	Java: true,
	// },
}

func main(){
	e := echo.New()
 connection.DatabaseConnect()
	e.Static("/public", "public")
	
	
	e.GET("/home", home)
	e.GET("/project", project)
	e.GET("/blog", blog)
	e.GET("/testimonials", testimonials)
	e.GET("/contact", contact)
	e.GET("/blog-detail/:id", blogDetail)
	e.GET("/editProject/:id", editBlog)
	
	// ...
	e.POST("/addblog", addBlog)
	e.POST("/delete-blog/:id", deleteBlog)
	e.POST("/edit-blog/:id", editBlog)
// ...

	
	e.Logger.Fatal(e.Start("localhost:5000"))
}

//FUNCTION HOME
func home(c echo.Context)error{
	tmpl, err := template.ParseFiles("views/index.html")

	data,_ := connection.Conn.Query(context.Background(),"SELECT id, project_name, start_date, end_date, duration, description, javascript, reactjs, golang, java FROM tb_project")
	 
	var result []Project
	for data.Next() { 
	var each = Project{}
	err := data.Scan(&each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Duration, &each.Description, &each.Javascript, &each.ReactJs, &each.Golang, &each.Java)
		
	if err != nil{
		fmt.Println((err.Error()))
		return c.JSON(http.StatusInternalServerError, err.Error())
	 
	}
	result = append(result, each)
}

if err != nil {
	return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
}
dataIndex := map[string]interface{}{
	"Projects": result,
}
fmt.Println("ini data index", dataIndex)
return tmpl.Execute(c.Response(), dataIndex)
}



//FUNCTION PROJECT
func project(c echo.Context)error{
	tmpl, err := template.ParseFiles("views/project.html")
	 if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	 
	}
return tmpl.Execute(c.Response(), nil)
}


//FUNCTION BLOG
func blog(c echo.Context)error{
tmpl, err := template.ParseFiles("views/blog.html")
if err!= nil{
	return c.JSON(http.StatusInternalServerError, err.Error())
	}
	data := map[string]interface{}{
		"Blogs" : Blogs,
	}
	return tmpl.Execute(c.Response(), data)
}



//FUNCTION BLOG DETAIL
func blogDetail(c echo.Context)error{
	id, _ := strconv.Atoi(c.Param("id"))

	var blogDetail = Project{}

	for i, data := range Blogs {
		if id == i {
			blogDetail = Project{
				ProjectName: data.ProjectName,
				StartDate: data.StartDate,
				EndDate: data.EndDate,
				Duration: data.Duration,
				Description: data.Description,
				Javascript: data.Javascript,
				ReactJs: data.ReactJs,
				Golang: data.Golang,
				Java: data.Java,
			}
		}
	}
	data := map[string]interface{}{
		"Project":   blogDetail,
	}
	var tmpl, err = template.ParseFiles("views/projectdetail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), data)
}


//FUNCTION ADD BLOG
 func addBlog(c echo.Context) error {
	title := c.FormValue("input-project")
	content := c.FormValue("content")
	
	startDateStr := c.FormValue("startDate")
	endDateStr := c.FormValue("endDate")

	javascript := c.FormValue("javascript")
	reactJs := c.FormValue("reactJs") 
	golang := c.FormValue("golang") 
	java := c.FormValue("java") 

	// Hitung durasi menggunakan fungsi calculateDuration
	duration := calculateDuration(startDateStr, endDateStr)
	
	fmt.Println("title: ", title)
	fmt.Println("content: ", content)
	fmt.Println("startDate: ", startDateStr)
	fmt.Println("endDate: ", endDateStr)
	fmt.Println("Durasi: ", duration)
	fmt.Println("javascript: ", javascript)
	fmt.Println("reactJs: ", reactJs)
	fmt.Println("golang: ", golang)
	fmt.Println("java: ", java)
	
	var newProject = Project{
		ProjectName: title,
		StartDate:  startDateStr,
		EndDate:   endDateStr,
		Duration:    duration,
		Description:    content,
		Javascript: javascript == javascript,
		ReactJs: reactJs == reactJs,
		Golang: golang == golang,
		Java: java == java,
	}
	
	Blogs = append(Blogs, newProject)
	
	fmt.Println(Blogs)
	
	return c.Redirect(http.StatusMovedPermanently, "/blog")
}



// FUNCTION TO CALCULATE DURATION
func calculateDuration(startDateStr, endDateStr string) string {
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, startDateStr)
	if err != nil {
		return "Invalid Start Date"
	}

	endDate, err := time.Parse(layout, endDateStr)
	if err != nil {
		return "Invalid End Date"
	}

	duration := endDate.Sub(startDate)
	years := int(duration.Hours() / 24 / 365)
	months := int(duration.Hours() / 24 / 30) % 12
	days := int(duration.Hours() / 24) % 30

	result := ""
	if years > 0 {
		result += fmt.Sprintf("%d Tahun ", years)
	}
	if months > 0 {
		result += fmt.Sprintf("%d Bulan ", months)
	}
	if days > 0 {
		result += fmt.Sprintf("%d Hari", days)
	}

	return result
}

//function show blog
// func showBlog(c echo.Context) error {
// 	tmpl := template.Must(template.ParseFiles("blog.html"))
// 	data := map[string]interface{}{
// 		"Projects": Blogs,
// 	}
// 	return tmpl.Execute(c.Response(), data)
// }


// Perbaiki fungsi deleteBlog
func deleteBlog(c echo.Context) error {
	id := c.FormValue("id")
	idToInt, err := strconv.Atoi(id)
	if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid blog ID")
	}

	if idToInt < 0 || idToInt >= len(Blogs) {
			return c.JSON(http.StatusBadRequest, "Invalid blog ID")
	}

	Blogs = append(Blogs[:idToInt], Blogs[idToInt+1:]...)
	return c.Redirect(http.StatusFound, "/blog")
}

// Perbaiki fungsi editBlog
func editBlog(c echo.Context) error {
	id := c.FormValue("id")
	idToInt, err := strconv.Atoi(id)
	if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid blog ID")
	}

	if idToInt < 0 || idToInt >= len(Blogs) {
			return c.JSON(http.StatusBadRequest, "Invalid blog ID")
	}

	// Mendapatkan data dari form request untuk mengupdate blog
	blogToUpdate := &Blogs[idToInt]
	blogToUpdate.ProjectName = c.FormValue("input-project")
	blogToUpdate.StartDate = c.FormValue("startDate")
	blogToUpdate.EndDate = c.FormValue("endDate")
	blogToUpdate.Description = c.FormValue("content")
	blogToUpdate.Javascript = c.FormValue("javascript") == "javascript"
	blogToUpdate.ReactJs = c.FormValue("reactJs") == "reactJs"
	blogToUpdate.Golang = c.FormValue("golang") == "golang"
	blogToUpdate.Java = c.FormValue("java") == "java"

	// Calculate duration using calculateDuration function
	blogToUpdate.Duration = calculateDuration(blogToUpdate.StartDate, blogToUpdate.EndDate)

	return c.Redirect(http.StatusFound, "/blog")
}



//FUNCTION TESTIMONIALS
func testimonials(c echo.Context)error{
	tmpl, err := template.ParseFiles("views/testimonials.html")
	 if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	 
	}
return tmpl.Execute(c.Response(), nil)
}


//FUNCTION CONTACT
func contact(c echo.Context)error{
	tmpl, err := template.ParseFiles("views/contact.html")
	 if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	 
	}
return tmpl.Execute(c.Response(), nil)
}