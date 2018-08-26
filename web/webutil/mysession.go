package webutil

import (
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
)

type MSession struct {
	cookieHandler *securecookie.SecureCookie
}

var MySession MSession = MSession{
	cookieHandler: securecookie.New(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32)),
}

func (t *MSession) GetUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = MySession.cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
		fmt.Println("getusername cookie :", cookie)
	}

	fmt.Println(MySession)
	return userName
}

func (t *MSession) GetOrgName(request *http.Request) (orgName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = MySession.cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			orgName = cookieValue["orgname"]
		}
		fmt.Println("getusername cookie :", cookie)
	}

	fmt.Println(MySession)
	return orgName
}

//SetSession set user session
func (t *MSession) SetSession(userName string, orgName string, response http.ResponseWriter) {
	value := map[string]string{
		"name":    userName,
		"orgname": orgName,
	}
	if encoded, err := MySession.cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		fmt.Println("setsession cookie :", cookie)
		http.SetCookie(response, cookie)
	}
	fmt.Println(MySession)
}

func (t *MSession) ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	fmt.Println("clearsession cookie :", cookie)
	fmt.Println(MySession)
	http.SetCookie(response, cookie)
}
