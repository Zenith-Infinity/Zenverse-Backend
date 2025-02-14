package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rayfanaqbil/Zenverse-BP/config"
	"github.com/rayfanaqbil/Zenverse-BP/helper"
	"github.com/rayfanaqbil/Zenverse-BP/model"
	"go.mongodb.org/mongo-driver/bson"
)

//Img Post
func PostUploadGithub(c *fiber.Ctx) error {
	fmt.Println("Starting file upload process")

	header, err := c.FormFile("img")
	if err != nil {
		fmt.Println("Error parsing form file:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"response": "Error parsing file",
			"info":     err.Error(),
		})
	}

	if header.Filename == "" {
		fmt.Println("Error: Filename is empty")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"response": "Filename cannot be empty",
		})
	}

	folder := helper.GetParam(c)
	pathFile := header.Filename
	if folder != "" {
		pathFile = folder + "/" + header.Filename
	}

	gh, err := helper.GetOneDoc[model.Ghcreates](config.Ulbimongoconn, "github", bson.M{})
	if err != nil {
		fmt.Println("Error fetching GitHub credentials:", err)
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"response": "Failed to get GitHub credentials",
			"info":     helper.GetSecretFromHeader(c),
		})
	}

	content, _, err := helper.GithubUpload(
		gh.GitHubAccessToken, gh.GitHubAuthorName, gh.GitHubAuthorEmail, header,
		"zenith-infinity", "img-repository", pathFile, false,
	)

	if err != nil {
		fmt.Println("Error uploading file to GitHub:", err)
		return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
			"response": "Failed to upload to GitHub",
			"info":     err.Error(),
		})
	}

	if content == nil || content.Content == nil {
		fmt.Println("Error: content or content.Content is nil")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"response": "Error uploading file",
		})
	}

	fmt.Println("File upload process completed successfully")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"response": *content.Content.Path,
		"info":     *content.Content.Name,
	})
}
