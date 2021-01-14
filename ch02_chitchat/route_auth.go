package main

import (
	"gwp/Chapter_2_Go_ChitChat/chitchat/my"
	//"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
	"net/http"
)






// GET /signup
//注册页面处理器函数
func signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "login.layout","signup")
}

// POST /signup_account
// 处理注册提交
func signupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}
	user := my.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		danger(err, "Cannot create user")
	}
	//记录一下注册的用户看看
	info(user)
	//注册成功之后，跳转到登录页面
	http.Redirect(writer, request, "/login", 302)
}

// GET /login
// 登录页面
func login(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("login.layout", "login")
	t.Execute(writer, nil)
}
// POST /authenticate
// 登录处理控制器
func authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	//先判断用户是否存在
	user, err := my.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}
	//判断密码是否正确
	if user.Password == my.Encrypt(request.PostFormValue("password")) {
		//创建服务器的session，目前session是存在sql数据库中
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		//设置cookie，让客户端保存好sessionid
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		//验证成功之后跳转到首页
		http.Redirect(writer, request, "/", 302)
	} else {
		//验证失败还是跳转到登录页
		http.Redirect(writer, request, "/login", 302)
	}
}

// GET /logout
//退出登录
func logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		warning(err, "Failed to get cookie")
		session := my.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/login", 302)
}
