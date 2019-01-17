package main

import (
	"fmt"
)

type UserRepository struct {
	Description string
}

func NewUserRepository(description string) *UserRepository {
	return &UserRepository{
		Description: description,
	}
}

type UserService struct {
	UserRepository *UserRepository
}

func NewUserService(userRepo *UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepo,
	}
}

type Builder func(con *Container) interface{}

// definition class
type Definition struct {
	Name    string
	Builder Builder
}

// DI Container class
type Container struct {
	store map[string]Builder
	share map[string]interface{}
}

func NewContainer() *Container {
	return &Container{
		store: map[string]Builder{},
		share: map[string]interface{}{},
	}
}

// DI Container method: register the object creation method
func (c *Container) MethodRegister(d *Definition) {
	c.store[d.Name] = d.Builder
}

// DI Container method: register object itself
func (c *Container) ObjectRegister(d *Definition) {
	c.share[d.Name] = d.Builder(c)
}

// DI Container method: get object from container
func (c *Container) Get(key string) interface{} {
	instance, _ := c.share[key]
	if instance == nil {
		builder, _ := c.store[key]
		instance = builder(c)
	}
	return instance
}

func main() {
	container := NewContainer()

	rd := &Definition{
		Name: "UserRepo",
		Builder: func(con *Container) interface{} {
			return NewUserRepository("Instance of user repo class")
		},
	}
	container.ObjectRegister(rd)

	sd := &Definition{
		Name: "UserService",
		Builder: func(con *Container) interface{} {
			userRepo := con.Get("UserRepo").(*UserRepository)
			return NewUserService(userRepo)
		},
	}
	container.MethodRegister(sd)

	userServiceInstance1 := container.Get("UserService").(*UserService)
	userServiceInstance2 := container.Get("UserService").(*UserService)

	fmt.Println("Objects list:")
	// expect 2 difference object ID
	fmt.Printf("UserService1: %p\n", userServiceInstance1)
	fmt.Printf("UserService2: %p\n", userServiceInstance2)
	// expect the same object ID
	fmt.Printf("UserRepo for UserService1: %p\n", userServiceInstance1.UserRepository)
	fmt.Printf("UserRepo for UserService2: %p\n", userServiceInstance2.UserRepository)

	fmt.Println("Shared objects: ", container.share)
	fmt.Println("Stored objects definition method (not share): ", container.store)
}
