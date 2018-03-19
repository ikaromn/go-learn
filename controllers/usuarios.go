package controllers

import (
	"net/http"
	"github.com/ikaro/Usuarios/models"
	"github.com/labstack/echo"
)

func Home(c echo.Context) error {
	var usuarios []models.Usuarios

	if err := models.UsuarioModel.Find().All(&usuarios); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Erro ao tentar recuperar os dados",
		})
	}

	data := map[string]interface{}{
		"title" : "Listagem de Usuários",
		"usuarios": usuarios,
	}

	return c.Render(http.StatusOK, "index.html", data)
}

func Create(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	var usuario models.Usuarios

	usuario.Name = name
	usuario.Email = email

	if name != "" && email != "" {
		if _, err := models.UsuarioModel.Insert(usuario); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"mensagem" : "Não foi possivel adicionar os dados no banco, tente novamente!",
			})
		}

		return c.JSON(http.StatusCreated, map[string]string{
			"mensagem": "Os dados foram salvos com sucesso",
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem" : "Informe os dados nos campos",
	})
}