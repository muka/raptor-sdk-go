package raptor

import (
	"fmt"

	"github.com/raptorbox/raptor-sdk-go/models"
)

//CreateUser instantiate a new API client
func CreateUser(r *Raptor) *User {
	return &User{
		Raptor: r,
	}
}

//User API client
type User struct {
	Raptor *Raptor
}

//GetConfig return the configuration
func (s *User) GetConfig() models.Config {
	return s.Raptor.GetConfig()
}

//GetClient return a client instance
func (s *User) GetClient() models.Client {
	return s.Raptor.GetClient()
}

//List the available users
func (s *User) List() ([]models.User, error) {
	raw, err := s.GetClient().Get(USER_LIST, nil)
	if err != nil {
		return nil, err
	}

	res := make([]models.User, 0)
	err = s.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//Read an user
func (s *User) Read(id string) (*models.User, error) {
	raw, err := s.GetClient().Get(fmt.Sprintf(USER_GET, id), nil)
	if err != nil {
		return nil, err
	}

	res := &models.User{}
	err = s.GetClient().FromJSON(raw, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//Create an user
func (s *User) Create(user *models.User) (*models.User, error) {
	raw, err := s.GetClient().Post(USER_CREATE, user, nil)
	if err != nil {
		return nil, err
	}

	err = s.GetClient().FromJSON(raw, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

//Update an user
func (s *User) Update(user *models.User) (*models.User, error) {
	raw, err := s.GetClient().Put(fmt.Sprintf(USER_UPDATE, user.ID), user, nil)
	if err != nil {
		return nil, err
	}

	err = s.GetClient().FromJSON(raw, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

//Delete an user
func (s *User) Delete(user *models.User) error {
	err := s.GetClient().Delete(fmt.Sprintf(USER_DELETE, user.ID), nil)
	return err
}
