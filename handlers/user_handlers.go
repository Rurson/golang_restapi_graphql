package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

//User Entity

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	TestField string
}

var users []User

func UserPost(context echo.Context) error {
	newUser := new(User)
	newUser.Id = len(users)
	err := context.Bind(newUser)

	if err != nil {
		return err
	}
	users = append(users, *newUser)
	return context.JSON(http.StatusCreated, newUser)
}

func UserGet(context echo.Context) error {
	UserToFind := new(User)
	err := context.Bind(UserToFind)
	if err != nil {
		return err
	}

	for i := range users {
		if users[i].Id == UserToFind.Id {
			return context.JSON(http.StatusOK, users[i])
		}
	}
	return context.JSON(http.StatusNotFound, fmt.Sprintf("Index %d not found in Users", UserToFind.Id))

}

func UserPut(context echo.Context) error {
	UserToFind := new(User)
	err := context.Bind(UserToFind)
	if err != nil {
		return err
	}

	for i := range users {
		if users[i].Id == UserToFind.Id {
			users[i] = *UserToFind
			return context.JSON(http.StatusAccepted, users[i])
		}
	}

	return context.JSON(http.StatusNotFound, fmt.Sprintf("Index %d not found in users", UserToFind.Id))

}

func UserDelete(context echo.Context) error {
	UserToDelete := new(User)
	err := context.Bind(UserToDelete)
	if err != nil {
		return err
	}

	if UserToDelete.Id >= len(users) {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Index %d not found in users", UserToDelete.Id))
	}

	users = append(users[:UserToDelete.Id], users[UserToDelete.Id+1:]...)
	return context.JSON(http.StatusAccepted, fmt.Sprintf("User of index %d has been deleted", UserToDelete.Id))

}
