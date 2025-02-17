package service

import (
	"errors"
	"fmt"
	
	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/storage"
)

type IService interface {
	Auth(auth *models.SAuth)(int32,error)
	UpdateUser(idGenerate int32, update *models.SUser)(bool,error)
	GetUser(idGenerate int32) (*models.SUser, error)
	GetUsers()
	DeleteUser(idGenerate int32) error
	GetStatistics()
	GetMap()(map[int32]*models.SUser)
	Add(user *models.SUser)(bool,error)
}

type service struct {
	DeleteUsersCount int32
	UpdateCount int32
	GetUserCount int32
	GetUsersCount int32
	GetAuthClick int32
	GetAddClick int32
	repo storage.IStorage
}

func NewIService(repo storage.IStorage) IService{
	return &service{repo: repo, DeleteUsersCount: 0,UpdateCount: 0,GetUserCount: 0,GetUsersCount: 0,GetAuthClick: 0,GetAddClick: 0}
}

func(s *service)GetMap()(map[int32]*models.SUser){
	return s.repo.GetMap()
}

func (stat *service) GetStatistics(){
	fmt.Println("kolichestvo klick po Delete",stat.DeleteUsersCount)
	fmt.Println("kolichestvo klick po AuthClick",stat.GetAuthClick)
	fmt.Println("kolichestvo klick po GetUser",stat.GetUserCount)
	fmt.Println("kolichestvo klick po GetUsers",stat.GetUsersCount)
	fmt.Println("kolichestvo klick po Update",stat.UpdateCount)
	fmt.Println("kolichestvo klick po Add",stat.GetAddClick)
}

func (s *service) DeleteUser(idGenerate int32) error{
	s.DeleteUsersCount++

	if idGenerate < 0 {
		return errors.New("ID cant be < 0")
	}
	err := s.repo.Delete(idGenerate)
	if err != nil{
		return err
	}

	return nil
}

func (s *service) GetUsers(){
	s.GetUsersCount++
	s.repo.GetAll()
}

func (s *service) GetUser(idGenerate int32) (*models.SUser,error){
	s.GetUserCount++

	if idGenerate < 0 {
		return nil, errors.New("ID cant be < 0")
	}

	value, err := s.repo.Get(idGenerate)
	if err != nil {
		return nil, errors.New("Error cant get user")
	}

	return value,nil
}

func (s *service) UpdateUser(idGenerate int32, update *models.SUser)(bool, error){
	s.UpdateCount++

	if idGenerate < 0 {
		return false, errors.New("ID cant be < 0")
	}

	err := s.repo.Update(idGenerate,update)
	if err != nil {
		return false, errors.New(" cant update error") 
	}

	return true, nil
}

func(s *service) Add(user *models.SUser)(bool,error){
	s.GetAddClick++

	if _,err := s.repo.Add(user); err != nil{
		return false, errors.New("user is nill")
	}

		return true, nil
}

func(s *service) Auth(auth *models.SAuth)(int32,error){
	s.GetAuthClick++

	for key, value := range s.repo.GetMap(){
		if auth.Login == value.Login {
			if auth.PasswordHash == value.PasswordHash {
				return key, nil
			}

				return 0, errors.New("Error Password dont match")
		}
	}
	
	return 0, errors.New("Error cant find user")
}
