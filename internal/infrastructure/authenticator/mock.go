package authenticator

import (
	"encoding/json"
	"fmt"

	"github.com/mechatron-x/atehere/internal/config"
)

const (
	authenticatorFilename = "/mock/authenticator.json"
)

type (
	mockUser struct {
		ID       string `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	mockStorageData struct {
		Users []mockUser `json:"users"`
	}

	MockAuthenticator struct {
		users       []mockUser
		path        string
		fileManager FileManager
	}
)

type (
	FileManager interface {
		Save(savePath string, data []byte) error
		Read(readPath string) ([]byte, error)
		Delete(deletePath string) error
	}
)

func NewMock(apiConf config.Api, fileManager FileManager) (*MockAuthenticator, error) {
	path := fmt.Sprintf("%s/%s", apiConf.StaticRoot, authenticatorFilename)

	mockAuthenticator := &MockAuthenticator{
		users:       make([]mockUser, 0),
		path:        path,
		fileManager: fileManager,
	}

	err := mockAuthenticator.restore()
	if err != nil {
		return nil, err
	}

	return mockAuthenticator, nil
}

func (ma *MockAuthenticator) CreateUser(id, email, password string) error {
	user := mockUser{
		ID:       id,
		Email:    email,
		Password: password,
	}

	for _, u := range ma.users {
		if u.ID == id {
			return fmt.Errorf("user with id: %s already exists", id)
		}

		if u.Email == email {
			return fmt.Errorf("user with email: %s already exists", email)
		}
	}

	ma.users = append(ma.users, user)

	return ma.save()
}

func (ma *MockAuthenticator) GetUserID(idToken string) (string, error) {
	for _, u := range ma.users {
		if u.ID != idToken {
			continue
		}

		return u.ID, nil
	}

	return "", fmt.Errorf("user with idToken: %s not found", idToken)
}

func (ma *MockAuthenticator) GetUserEmail(idToken string) (string, error) {
	for _, u := range ma.users {
		if u.ID != idToken {
			continue
		}

		return u.Email, nil
	}

	return "", fmt.Errorf("user with idToken: %s not found", idToken)
}

func (ma *MockAuthenticator) RevokeRefreshTokens(idToken string) error {
	users := ma.users
	for i, u := range users {
		if u.ID != idToken {
			continue
		}

		users[i] = users[len(users)-1]
		ma.users = users[:len(users)-1]

		return nil
	}

	return fmt.Errorf("user with idToken: %s not found", idToken)
}

func (ma *MockAuthenticator) save() error {
	saveData := &mockStorageData{
		Users: ma.users,
	}

	bytes, err := json.Marshal(saveData)
	if err != nil {
		return err
	}

	return ma.fileManager.Save(ma.path, bytes)
}

func (ma *MockAuthenticator) restore() error {
	readData := &mockStorageData{}

	bytes, err := ma.fileManager.Read(ma.path)
	if err != nil {
		return err
	}

	if len(bytes) == 0 {
		return nil
	}

	err = json.Unmarshal(bytes, readData)
	if err != nil {
		return err
	}

	ma.users = readData.Users
	return nil
}
