package webutil

import (
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
	}
	return userName
}

//SetSession set user session
func (t *MSession) SetSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := MySession.cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func (t *MSession) ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
