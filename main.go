package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"revisiweb/connection"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
	Image string
}

type user struct{
	Id int
	Name string
	Email string
	HashedPassword string

}

type UserLoginSession struct {
	IsLogin bool
	Name    string
}

var userLoginSession = UserLoginSession{}

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

 e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.Static("/public", "public")
	
	
	e.GET("/home", home)
	e.GET("/project", project)
	e.GET("/blog", blog)
	e.GET("/testimonials", testimonials)
	e.GET("/contact", contact)
	e.GET("/blog-detail/:id", blogDetail)
	e.GET("/editProject/:id", editBlogForm)
	e.GET("/form-login", formLogin)
	e.GET("/form-register", formRegister)
	
	// ...
	e.POST("/addblog", addBlog)
	e.POST("/delete-blog/:id", deleteBlog)
	e.POST("/edit-blog/:id", editBlog)
	e.POST("/login",login)
	e.POST("/register", register)
	e.POST("/logout", logout)
// ...

	
	e.Logger.Fatal(e.Start("localhost:5000"))
}


var userData = user{}
//FUNCTION HOME
func home(c echo.Context)error{
	var tmpl, err = template.ParseFiles("views/index.html")

	sess, _ := session.Get("session", c)

	datas := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
		"DataSession":  userData,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"massage": err.Error()})
	}

	return tmpl.Execute(c.Response(), datas)
}
// 	var tmpl, err = template.ParseFiles("views/index.html")

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 	}
// 	dataBlogs, errBlogs := connection.Conn.Query(context.Background(), "SELECT id, project_name, start_date, end_date, duration, description, javascript, reactjs, golang, java FROM tb_project ORDER by id ASC" )
// 	if  errBlogs != nil {
// 		return c.JSON(http.StatusInternalServerError, errBlogs.Error())
// 	}
// var resultBlogs[] Project
// for dataBlogs.Next(){
// 	var  each = Project {}
// 	err := dataBlogs.Scan(&each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Duration, &each.Description, &each.Javascript, &each.ReactJs, &each.Golang, &each.Java)
// 	if err != nil{
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 		resultBlogs = append(resultBlogs, each)
// 	}

// 	projects := map[string]interface{}{
// 		"Projects": resultBlogs,
// 	}


// 	return tmpl.Execute(c.Response(), projects)
// }

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
	dataBlogs, errBlogs := connection.Conn.Query(context.Background(), "SELECT id, project_name, start_date, end_date, duration, description, javascript, reactjs, golang, java FROM tb_project ORDER by id ASC" )
	if  errBlogs != nil {
		return c.JSON(http.StatusInternalServerError, errBlogs.Error())
	}
var resultBlogs[] Project
for dataBlogs.Next(){
	var  each = Project {}
	err := dataBlogs.Scan(&each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Duration, &each.Description, &each.Javascript, &each.ReactJs, &each.Golang, &each.Java)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	

		resultBlogs = append(resultBlogs, each)
	}
	sess, _ := session.Get("session", c)

    if sess.Values["isLogin"] != true {
        userLoginSession.IsLogin = false
    } else {
        userLoginSession.IsLogin = true
        userLoginSession.Name = sess.Values["name"].(string)
    }



    data := map[string]interface{}{
        "Projects": resultBlogs,
        "UserLoginSession": userLoginSession,
    }


	return tmpl.Execute(c.Response(), data)
}





//FUNCTION BLOG DETAIL
func blogDetail(c echo.Context)error{
	id := c.Param("id")

	tmpl, err := template.ParseFiles("views/projectdetail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}


	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid project ID"})
	}
	 
	idToInt,_ := strconv.Atoi(id)
		projectdetail := Project{}

		//QUERY GET 1 DATA

		err = connection.Conn.QueryRow(context.Background(), "SELECT id, project_name, start_date, end_date, duration, description, javascript, reactjs, golang, java FROM tb_project WHERE id=$1", idToInt).Scan(&projectdetail.Id, &projectdetail.ProjectName, &projectdetail.StartDate, &projectdetail.EndDate, &projectdetail.Duration, &projectdetail.Description, &projectdetail.Javascript, &projectdetail.ReactJs, &projectdetail.Golang, &projectdetail.Java)
		
		fmt.Print("ini project detail:", err)
		if  err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
	
	}
		data := map[string]interface{}{
			"id": id,
			"Project": projectdetail,

			
		}
	 return tmpl.Execute(c.Response(), data)
	}


		//FUNCTION ADD BLOG
 func addBlog(c echo.Context) error {
	title := c.FormValue("input-project")
	description := c.FormValue("content")
	
	startDateStr := c.FormValue("startDate")
	endDateStr := c.FormValue("endDate")

	javascript := c.FormValue("javascript") == "javascript"
	reactJs := c.FormValue("reactJs") =="reactJs"
	golang := c.FormValue("golang") == "golang"
	java := c.FormValue("java") == "java"

	
	// Hitung durasi menggunakan fungsi calculateDuration
	duration := calculateDuration(startDateStr, endDateStr)


	

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_project (project_name, start_date, end_date, duration, description, javascript, reactjs, golang, java) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", title, startDateStr, endDateStr, duration, description, javascript, reactJs, golang, java,)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	fmt.Println("title", title)
	fmt.Println("description", description)
	fmt.Println("startDateStr", startDateStr)
	fmt.Println("endDateStr", endDateStr)
	fmt.Println("reactJs", reactJs)
	fmt.Println("golang", golang)
	fmt.Println("java", java)

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

// Perbaiki fungsi deleteBlog
func deleteBlog(c echo.Context) error {
	id := c.Param("id")
	idToInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	_, err = connection.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id = $1", idToInt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/home")
}



//FUNCTION EDITPROJECT
func editBlogForm(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_project WHERE id=$1", id).Scan(&ProjectDetail.Id, &ProjectDetail.ProjectName, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Duration, &ProjectDetail.Description, &ProjectDetail.Javascript, &ProjectDetail.ReactJs, &ProjectDetail.Golang, &ProjectDetail.Java)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message query": err.Error()})
	}
	data := map[string]interface{}{
		"Project": ProjectDetail,
	}

	var tmpl, errTmp = template.ParseFiles("views/editproject.html")
	if errTmp != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message rout": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)

}
	

//FUNCION EDIT BLOG
func editBlog(c echo.Context)error{
	id := c.FormValue("id")
	projectName := c.FormValue("input-project")
	startDateStr := c.FormValue("startDate")
	endDateStr := c.FormValue("endDate")
	description := c.FormValue("content")

	javascript:= c.FormValue("javascript") =="javascript"
	reactJs:= c.FormValue("reactJs") =="reactjs"
	golang:= c.FormValue("golang") =="golang"
	java:= c.FormValue("java") =="java"

	duration := calculateDuration(startDateStr, endDateStr)

	
	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_project SET project_name=$1, start_date=$2, end_date=$3, duration=$4, description=$5, javascript=$6, reactjs=$7, golang=$8, java=$9 WHERE id=$10", projectName, startDateStr, endDateStr, duration, description, javascript, reactJs, golang, java, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/blog")
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


//FUNCTION FORMLOGIN

func formLogin(c echo.Context) error {


	tmpl, err := template.ParseFiles("views/form-login.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}	

	flash := map[string]interface{}{
		"FlashMessage": sess.Values["message"], // "Register berhasil"
		"FlashStatus":  sess.Values["status"],  // true
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), flash)
}


//FUNCTION FORM REGISTER


func formRegister(c echo.Context) error {


	tmpl, err := template.ParseFiles("views/form-register.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	

	flash := map[string]interface{}{
		"FlashMessage": sess.Values["message"], // "Register berhasil"
		"FlashStatus":  sess.Values["status"],  // true
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), flash)
}


//function login

func login(c echo.Context) error {
	inputEmail := c.FormValue("inputEmail")
	inputPassword := c.FormValue("inputPassword") //

	user := user{}

	// check apakah ada emailnya di db
	err := connection.Conn.QueryRow(context.Background(), "SELECT id, name, email, password FROM tb_user WHERE email=$1", inputEmail).Scan(&user.Id, &user.Name, &user.Email, &user.HashedPassword)

	if err != nil {
		return redirectWithMessage(c, "Login gagal!", false, "/form-login")
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(inputPassword))

	if errPassword != nil {
		return redirectWithMessage(c, "Login gagal!", false, "/form-login")
	}

	// return c.JSON(http.StatusOK, "Berhasil login!")

	// set session login (berhasil login)
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 // 3 JAM -> berapa lama expired
	sess.Values["message"] = "Login success!"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.Id
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/home")
}


//function register
func register(c echo.Context) error {

	inputName := c.FormValue("inputName")
	inputEmail := c.FormValue("inputEmail") // harus valid email
	inputPassword := c.FormValue("inputPassword")
	
	// validasi (trim, validasi valid email)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(inputPassword), 10)

	if err != nil {
		// fmt.Println("masuk sini")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	fmt.Println(inputName, inputEmail, inputPassword)

	query, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_user (name, email, password) VALUES($1, $2, $3)", inputName, inputEmail, hashedPassword)

	fmt.Println("affected row : ", query.RowsAffected())

	if err != nil {
		return redirectWithMessage(c, "Register gagal!", false, "/form-register")
	}

	return redirectWithMessage(c, "Register berhasil!", true, "/form-login")
}
//function logout
func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/home")
}


//function redirect
func redirectWithMessage(c echo.Context, message string, status bool, redirectPath string) error {
	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, redirectPath)
}

	
	