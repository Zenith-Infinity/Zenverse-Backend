package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rayfanaqbil/Zenverse-BP/config"
	"github.com/rayfanaqbil/Zenverse-BP/helper"
	"github.com/rayfanaqbil/Zenverse-BP/model"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/whatsauth/itmodel"
)

//Img Post
func PostUploadGithub(c *fiber.Ctx) error {
	var respn itmodel.Response

	fmt.Println("Starting file upload process")

	header, err := c.FormFile("img")
	if err != nil {
		fmt.Println("Error parsing form file:", err)
		respn.Response = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(respn)
	}

	folder := helper.GetParam(c)
	var pathFile string
	if folder != "" {
		pathFile = folder + "/" + header.Filename
	} else {
		pathFile = header.Filename
	}

	gh, err := helper.GetOneDoc[model.Ghcreates](config.Ulbimongoconn, "github", bson.M{})
	if err != nil {
		fmt.Println("Error fetching GitHub credentials:", err)
		respn.Info = helper.GetSecretFromHeader(c)
		respn.Response = err.Error()
		return c.Status(fiber.StatusConflict).JSON(respn)
	}

	content, _, err := helper.GithubUpload(
		gh.GitHubAccessToken, gh.GitHubAuthorName, gh.GitHubAuthorEmail, header,
<<<<<<< HEAD
		"zenith-infinity", "img-repository", pathFile, false,
=======
		"zenith-infinitity", "img-repository", pathFile, false,
>>>>>>> 8caf87e3eccc271a9db77144f2972d885fb88905
	)

	if err != nil {
		fmt.Println("Error uploading file to GitHub:", err)
		respn.Info = "Gagal upload ke GitHub"
		respn.Response = err.Error()
		return c.Status(fiber.StatusExpectationFailed).JSON(respn)
	}

	if content == nil || content.Content == nil {
		fmt.Println("Error: content or content.Content is nil")
		respn.Response = "Error uploading file"
		return c.Status(fiber.StatusInternalServerError).JSON(respn)
	}

	respn.Info = *content.Content.Name
	respn.Response = *content.Content.Path
	fmt.Println("File upload process completed successfully")
	return c.Status(fiber.StatusOK).JSON(respn)
}
