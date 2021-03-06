package controllers

import (
	"net/http"
	"strconv"

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
		"title":    "Listagem de Usuários",
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
				"mensagem": "Não foi possivel adicionar os dados no banco, tente novamente!",
			})
		}

		return c.Redirect(http.StatusFound, "/")
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"mensagem": "Informe os dados nos campos",
	})
}

func Delete(c echo.Context) error {
	usuarioID, _ := strconv.Atoi(c.Param("id"))

	result := models.UsuarioModel.Find("id=?", usuarioID)

	if count, _ := result.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensagem": "Não foi possivel encontrar o usuario",
		})
	}

	if err := result.Delete(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Não foi possivel deletar o usuario",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]string{
		"mensagem": "Usuário deletado com sucesso!",
	})
}

func Update(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))
	nome := c.FormValue("name")
	email := c.FormValue("email")

	user := models.Usuarios{
		ID:    userId,
		Name:  nome,
		Email: email,
	}

	result := models.UsuarioModel.Find("id=?", userId)

	if count, _ := result.Count(); count < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Usuário não existe",
		})
	}

	if err := result.Update(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Erro ao tentar atualizar o usuário",
		})
	}

	return c.JSON(http.StatusAccepted, user)
}

func Add(c echo.Context) error {
	return c.Render(http.StatusOK, "add.html", nil)
}

func UpdateController(c echo.Context) error {

	userId, _ := strconv.Atoi(c.Param("id"))

	result := models.UsuarioModel.Find("id=?", userId)

	if count, _ := result.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensagem": "Usuário não existe",
		})
	}

	var user models.Usuarios

	if err := result.One(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Usuário não encontrado",
		})
	}

	data := map[string]interface{}{
		"usuario": user,
	}

	return c.Render(http.StatusOK, "update.html", data)
}
