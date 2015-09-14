package doubanfm

import (
	"bytes"
	"mime/multipart"
)

type User struct {
	Id     string `json:"user_id"`
	Name   string `json:"user_name"`
	Email  string
	Token  string
	Expire string
}

// POST http://www.douban.com/j/app/login
//	{
//	  "user_id":"35586494",
//	  "err":"ok",
//	  "token":"8e81d12345",
//	  "expire":"1456109051",
//	  "r":0,
//	  "user_name":"Gerry",
//	  "email":"ginuerzh@gmail.com"
//	}
func Login(id, password string) (*User, error) {
	formdata := &bytes.Buffer{}

	w := multipart.NewWriter(formdata)
	w.WriteField("app_name", AppName)
	w.WriteField("version", AppVersion)
	w.WriteField("email", id)
	w.WriteField("password", password)
	defer w.Close()

	resp, err := post(LoginUrl, w.FormDataContentType(), formdata)
	if err != nil {
		return nil, err
	}

	var r struct {
		User
		dbError
	}

	if err = decode(resp, &r); err != nil {
		return nil, err
	}

	if r.R != 0 {
		return nil, &r.dbError
	}
	return &r.User, nil
}
