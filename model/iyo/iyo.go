// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package oauthiyo

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/mattermost/mattermost-server/einterfaces"
	"github.com/mattermost/mattermost-server/model"
	"fmt"
)

type IyoProvider struct {
}

type IyoEmailAddress struct {
	EmailAddress   string  `json:"emailaddress"`
}

type IyoUser struct {
	UserName       string 				`json:"username"`
	FirstName      string 				`json:"firstname"`
	LastName       string 				`json:"lastname"`
	EmailAddresses []IyoEmailAddress 	`json:"emailaddresses"`
}

func init() {
	provider := &IyoProvider{}
	einterfaces.RegisterOauthProvider(model.USER_AUTH_SERVICE_IYO, provider)
}

func userFromIyoUser(iyoUser *IyoUser) *model.User {
	user := &model.User{}
	if len(iyoUser.EmailAddresses) > 0 {
		user.Email = iyoUser.EmailAddresses[0].EmailAddress
		user.Username = strings.Split(user.Email, "@")[0]
	}
	user.FirstName = iyoUser.FirstName
	user.LastName = iyoUser.LastName
	user.AuthData = &iyoUser.UserName
	user.AuthService = model.USER_AUTH_SERVICE_IYO
	fmt.Printf("%#v", user)
	return user
}

func iyoUserFromJson(data io.Reader) *IyoUser {
	decoder := json.NewDecoder(data)
	var iyoUser IyoUser
	err := decoder.Decode(&iyoUser)
	fmt.Printf("%#v", iyoUser)
	if err == nil {
		return &iyoUser
	} else {
		return nil
	}
}

func (iyoUser *IyoUser) ToJson() string {
	b, err := json.Marshal(iyoUser)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func (iyoUser *IyoUser) IsValid() bool {
	if len(iyoUser.UserName) == 0 {
		return false
	}

	if len(iyoUser.EmailAddresses) == 0 {
		return false
	}

	return true
}

func (iyoUser *IyoUser) getAuthData() string {
	return iyoUser.UserName
}

func (m *IyoProvider) GetIdentifier() string {
	return model.USER_AUTH_SERVICE_IYO
}

func (m *IyoProvider) GetUserFromJson(data io.Reader) *model.User {
	iyoUser := iyoUserFromJson(data)
	if iyoUser.IsValid() {
		return userFromIyoUser(iyoUser)
	}

	return &model.User{}
}

func (m *IyoProvider) GetAuthDataFromJson(data io.Reader) string {
	iyoUser := iyoUserFromJson(data)
	if iyoUser.IsValid() {
		return iyoUser.getAuthData()
	}

	return ""
}
