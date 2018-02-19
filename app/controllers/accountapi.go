package controllers

import (
	//"encoding/json"
	//"io/ioutil"
	"log"
	"net/mail"
	"strconv"

	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"github.com/ytn86/scoreServer/app/models"
)

type APIv1AccountController struct {
	App
}

func (c APIv1AccountController) DoRegister() revel.Result {

	var user *models.User
	jsonResponse := make(map[string]interface{})

	field := struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Password2 string `json:"password2"`
		IsITF     string `json:"is_itf"`
	}{}

	//content, err := ioutil.ReadAll(c.Request.Body)
	//defer c.Request.Body.Close()
	//
	//err = json.Unmarshal(content, &field)
    
	err := c.Params.BindJSON(&field)

	if err != nil {
		log.Println(err)
		c.Response.Status = 400
		jsonResponse["status"] = 400
		jsonResponse["msg"] = "bad json"

		return c.RenderJSON(jsonResponse)
	}

	if field.Username == "" || field.Password == "" || field.Password2 == "" || field.Email == "" {
		c.Response.Status = 400
		jsonResponse["status"] = 400
		jsonResponse["msg"] = "Please fulfill all requied fields"
		return c.RenderJSON(jsonResponse)
	}

	_, err = mail.ParseAddress(field.Email)
	if err != nil {
		log.Println(err)

		c.Response.Status = 400
		jsonResponse["status"] = 400
		jsonResponse["msg"] = "invalid email address"

		return c.RenderJSON(jsonResponse)

	}

	if field.Password != field.Password2 {
		c.Response.Status = 400
		jsonResponse["status"] = 400
		jsonResponse["msg"] = "password and password2 are not matched"
		return c.RenderJSON(jsonResponse)
	}

	c.Begin()
	exists := models.CheckIfUserExistsByName(c.Tx, field.Username)
	c.Commit()

	if exists == true {
		c.Response.Status = 400
		jsonResponse["status"] = 400
		jsonResponse["msg"] = "the username has already taken"
		return c.RenderJSON(jsonResponse)
	}

	if field.IsITF == "0" {
		user = models.NewUser(field.Username, field.Email, field.Password, true)
	} else {
		user = models.NewUser(field.Username, field.Email, field.Password, false)
	}

	c.Begin()
	suc := models.RegisterUser(c.Tx, user)
	c.Commit()

	if suc != true {
		c.Response.Status = 500
		jsonResponse["status"] = 500
		jsonResponse["msg"] = "database error"

		return c.RenderJSON(jsonResponse)
	}

	jsonResponse["status"] = 200
	jsonResponse["msg"] = "success"

	return c.RenderJSON(jsonResponse)
}

func (c APIv1AccountController) DoLogin() revel.Result {

	/*
		username := c.Params.Get("username")
		password := c.Params.Get("password")
		//csrfToken := c.Params.Get("token")
	*/
	field := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	jsonResponse := make(map[string]interface{})
	
	//content, err := ioutil.ReadAll(c.Request.Body)
	//defer c.Request.Body.Close()

	/*
	if err != nil {
		log.Fatal(err)
		
		jsonResponse["status"] = 400
		jsonResponse["msg"] = "bad request"
		
		return c.RenderJSON(jsonResponse)
	}

	log.Println(content)
	err = json.Unmarshal(content, &field)
*/
	err := c.Params.BindJSON(&field)
	
	

	if err != nil {
		log.Println(err)
		c.Response.Status = 400
		jsonResponse["status"] = 400
		jsonResponse["msg"] = "bad json"

		return c.RenderJSON(jsonResponse)
	}

	if field.Username == "" || field.Password == "" {
		jsonResponse["status"] = 400
		jsonResponse["msg"] = "username and password are required"

		return c.RenderJSON(jsonResponse)
	}

	//user check
	c.Begin()
	uid := models.AuthUserByCred(c.Tx, field.Username, field.Password)
	c.Commit()

	if uid < 0 {
		jsonResponse["status"] = 404
		jsonResponse["msg"] = "invalid username or password"

		return c.RenderJSON(jsonResponse)
	}

	c.Begin()
	isAdmin := models.IsAdmin(c.Tx, field.Username)
	c.Commit()

	jsonResponse["status"] = 200
	jsonResponse["msg"] = "success"

	c.Session["user"] = field.Username
	c.Session["userid"] = strconv.Itoa(uid)

	if isAdmin == true {
		c.Session["isAdmin"] = "1"
	}

	return c.RenderJSON(jsonResponse)
}

/*
ToDo: Improve readability
*/
func (c APIv1AccountController) ModifyUserProfile() revel.Result {

	jsonResponse := make(map[string]interface{})
	profile := struct {
		Username     string `json:"username"`
		Email        string `json:"email"`
		Password     string `json:"password"`
		NewPassword  string `json:"newpassword"`
		NewPassword2 string `json:"newpassword2"`
		IsITF        bool   `json:"is_itf"`
		Comment      string `json:"comment"`
	}{}

	if c.Session["user"] == "" {
		c.Response.Status = 403
		jsonResponse["status"] = 403
		jsonResponse["msg"] = "login first"

		return c.RenderJSON(jsonResponse)

	}
	/*
	content, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()

	if err != nil {
		log.Println(err)
		c.Response.Status = 400
		jsonResponse["status"] = 400
		jsonResponse["msg"] = "bad request"

		return c.RenderJSON(jsonResponse)
	}

	err = json.Unmarshal(content, &profile)
*/

	err := c.Params.BindJSON(&profile)
	
	
	if err != nil {
		log.Println(err)
		c.Response.Status = 400
		jsonResponse["status"] = 400
		jsonResponse["msg"] = "bad json"

		return c.RenderJSON(jsonResponse)
	}

	if profile.Username == "" || profile.Email == "" {
		c.Response.Status = 400
		jsonResponse["status"] = 400
		jsonResponse["msg"] = "please fulfill all requied fields"

		return c.RenderJSON(jsonResponse)
	}

	_, err = mail.ParseAddress(profile.Email)
	if err != nil {
		log.Println(err)

		c.Response.Status = 400
		jsonResponse["status"] = 400
		jsonResponse["msg"] = "invalid email address"

		return c.RenderJSON(jsonResponse)

	}

	c.Begin()
	user := models.GetUserByName(c.Tx, c.Session["user"])
	c.Commit()

	if user.ID == -1 {
		c.Response.Status = 500
		jsonResponse["status"] = 500
		jsonResponse["msg"] = "something wrong..."

		return c.RenderJSON(jsonResponse)
	}

	if user.ID == 0 {
		c.Response.Status = 404
		jsonResponse["status"] = 404
		jsonResponse["msg"] = "user not found"

		return c.RenderJSON(jsonResponse)
	}

	user.Name = profile.Username
	user.Email = profile.Email
	user.IsITF = profile.IsITF
	user.Comment = profile.Comment

	if profile.Password == "" && profile.NewPassword == "" && profile.NewPassword2 == "" {

		c.Begin()
		suc := models.UpdateUserProfile(c.Tx, &user)
		c.Commit()

		if suc == false {
			c.Response.Status = 500
			jsonResponse["status"] = 500
			jsonResponse["msg"] = "something wrong..."

			return c.RenderJSON(jsonResponse)
		}

		go cache.Delete("ranks")
		
		c.Session["user"] = profile.Username

		jsonResponse["status"] = 200
		jsonResponse["msg"] = "success"

		return c.RenderJSON(jsonResponse)
	}

	if profile.Password != "" {

		c.Begin()
		uid := models.AuthUserByCred(c.Tx, c.Session["user"], profile.Password)
		c.Commit()

		if uid < 0 {
			jsonResponse["status"] = 400
			jsonResponse["msg"] = "current password is not correct"

			return c.RenderJSON(jsonResponse)
		}

		if profile.NewPassword != profile.NewPassword2 {
			jsonResponse["status"] = 400
			jsonResponse["msg"] = "new passwords do not match"

			return c.RenderJSON(jsonResponse)
		}

		user.Password = models.GenPasswordHash(profile.NewPassword)

		c.Begin()
		suc := models.UpdateUserProfile(c.Tx, &user)
		c.Commit()

		if suc == false {
			c.Response.Status = 500
			jsonResponse["status"] = 500
			jsonResponse["msg"] = "something wrong..."

			return c.RenderJSON(jsonResponse)
		}

		c.Session["user"] = profile.Username

		go cache.Delete("ranks")
		
		jsonResponse["status"] = 200
		jsonResponse["msg"] = "success"

		return c.RenderJSON(jsonResponse)

	}

	jsonResponse["status"] = 500
	jsonResponse["msg"] = "something wrong"
	return c.RenderJSON(jsonResponse)
}

func (c APIv1AccountController) GetUserSolves() revel.Result {

	jsonResponse := make(map[string]interface{})

	if c.Session["user"] == "" {
		c.Response.Status = 403
		jsonResponse["status"] = 403
		jsonResponse["msg"] = "login first"

		return c.RenderJSON(jsonResponse)

	}

	c.Begin()
	id := models.GetUserIDByName(c.Tx, c.Session["user"])
	c.Commit()

	if id == -1 {
		c.Response.Status = 500
		jsonResponse["status"] = 500
		jsonResponse["msg"] = "server error"

		return c.RenderJSON(jsonResponse)
	}

	if id == 0 {
		jsonResponse["status"] = 404
		jsonResponse["msg"] = "user not found"

		return c.RenderJSON(jsonResponse)
	}

	c.Begin()
	solved := models.GetSolvedTasksByUserID(c.Tx, id)
	c.Commit()
	
	jsonResponse["status"] = 200
	jsonResponse["msg"] = "success"
	jsonResponse["data"] = solved
	return c.RenderJSON(jsonResponse)

}

func (c APIv1AccountController) GetUserSolvesByID(userID int) revel.Result {

	jsonResponse := make(map[string]interface{})

	c.Begin()
	solved := models.GetSolvedTasksByUserID(c.Tx, userID)
	c.Commit()

	if solved == nil {
		jsonResponse["status"] = 500
		jsonResponse["msg"] = "server error"
		return c.RenderJSON(jsonResponse)
	}

	jsonResponse["status"] = 200
	jsonResponse["msg"] = "success"
	jsonResponse["data"] = solved
	return c.RenderJSON(jsonResponse)

}

func (c APIv1AccountController) GetUserProfile() revel.Result {

	jsonResponse := make(map[string]interface{})

	if c.Session["user"] == "" {
		jsonResponse["msg"] = "login first"

		return c.RenderJSON(jsonResponse)

	}

	userID, err := strconv.Atoi(c.Session["userid"])

	if err != nil {
		log.Println(err)

		c.Response.Status = 500
		jsonResponse["status"] = 500
		jsonResponse["msg"] = "server error"

		return c.RenderJSON(jsonResponse)

	}

	c.Begin()
	user := models.GetOwnProfileByID(c.Tx, userID)
	c.Commit()

	if user.ID == -1 {
		c.Response.Status = 500
		jsonResponse["status"] = 500
		jsonResponse["msg"] = "server error"

		return c.RenderJSON(jsonResponse)
	}

	if user.ID == 0 {
		jsonResponse["status"] = 404
		jsonResponse["msg"] = "user not found"

		return c.RenderJSON(jsonResponse)
	}

	c.Begin()
	userScore := models.GetUserScore(c.Tx, user.ID)
	c.Commit()

	if userScore == -1 {
		c.Response.Status = 500
		jsonResponse["status"] = 500
		jsonResponse["msg"] = "something wrong..."

		return c.RenderJSON(jsonResponse)
	}

	user.Score = userScore

	jsonResponse["status"] = 200
	jsonResponse["msg"] = "success"
	jsonResponse["data"] = user
	return c.RenderJSON(jsonResponse)

}

func (c APIv1AccountController) GetUserProfileByID(userID int) revel.Result {

	jsonResponse := make(map[string]interface{})

	if c.Session["user"] == "" {
		jsonResponse["msg"] = "login first"

		return c.RenderJSON(jsonResponse)

	}

	c.Begin()
	user := models.GetUserProfileByID(c.Tx, userID)
	c.Commit()

	if user.ID == -1 {
		c.Response.Status = 500
		jsonResponse["status"] = 500
		jsonResponse["msg"] = "server error"

		return c.RenderJSON(jsonResponse)
	}

	if user.ID == 0 {
		jsonResponse["status"] = 404
		jsonResponse["msg"] = "user not found"

		return c.RenderJSON(jsonResponse)
	}

	c.Begin()
	userScore := models.GetUserScore(c.Tx, user.ID)
	c.Commit()

	if userScore == -1 {
		c.Response.Status = 500
		jsonResponse["status"] = 500
		jsonResponse["msg"] = "something wrong..."

		return c.RenderJSON(jsonResponse)
	}

	user.Score = userScore

	jsonResponse["status"] = 200
	jsonResponse["msg"] = "success"
	jsonResponse["data"] = user
	return c.RenderJSON(jsonResponse)

}
