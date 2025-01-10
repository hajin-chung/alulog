package main

import (
	"encoding/json"
	"hajin-chung/deps.me/internal/generate"
	"hajin-chung/deps.me/internal/upload"
	"hajin-chung/deps.me/internal/env"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type PostList struct {
	Posts []string `json:"posts"`
}

func main() {
	env.LoadEnv()
	app := fiber.New()
	app.Use(HandleAuth)
	app.Get("/list", HandleList)
	app.Get("/read", HandleRead)
	app.Post("/write", HandleWrite)
	app.Get("/delete", HandleDelete)
	app.Get("/publish", HandlePublish)
	app.Listen(":3000")
}

func HandleAuth(c *fiber.Ctx) error {
	auth_list := c.GetReqHeaders()["X-Auth"]
	if len(auth_list) == 0 {
		return c.SendStatus(500)
	}
	password := auth_list[0]
	if password == env.Secret {
		return c.Next()
	}
	return c.SendStatus(500)
}

func HandleList(c *fiber.Ctx) error {
	entries, err := os.ReadDir(env.PostPath)
	if err != nil {
		log.Printf("error on reading dir ./posts\n%s\n", err)
		return c.SendStatus(500)
	}

	posts := []string{}
	for _, entry := range entries {
		posts = append(posts, entry.Name())
	}

	post_list := PostList{Posts: posts}
	res, err := json.Marshal(post_list)
	if err != nil {
		log.Printf("error on encoding post list json\n%s\n", err)
		return c.SendStatus(500)
	}
	_, err = c.Write(res)
	if err != nil {
		log.Printf("error on Writing response\n%s\n", err)
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}

func HandleRead(c *fiber.Ctx) error {
	filename := c.AllParams()["file"]
	content, err := os.ReadFile(env.PostPath + filename)
	if err != nil {
		log.Printf("error on reading file %s\n%s\n", filename, err)
		return c.SendStatus(500)
	}
	_, err = c.Write(content)
	if err != nil {
		log.Printf("error on Writing response\n%s\n", err)
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}

func HandleWrite(c *fiber.Ctx) error {
	filename := c.AllParams()["file"]
	body := c.Body()
	err := os.WriteFile(env.PostPath + filename, body, os.ModePerm)
	if err != nil {
		log.Printf("error on writing new file %s\n%s\n", filename, err)
		return c.SendStatus(500)
	}

	Publish()
	return c.SendStatus(200)
}

func HandleDelete(c *fiber.Ctx) error {
	filename := c.AllParams()["file"]
	err := os.Remove(env.PostPath + filename)
	if err != nil {
		log.Printf("error on removing file %s\n%s\n", filename, err)
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}

func HandlePublish(c *fiber.Ctx) error {
	err := Publish()
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}

func Publish() error {
	err := generate.GenereatePosts()
	if err != nil {
		log.Printf("error on generating posts\n%s\n", err)
		return err
	}
	err = upload.UploadDirectory("deps.me", "out", "ap-northeast-2")
	if err != nil {
		log.Printf("error on uploading directory\n%s\n", err)
		return err
	}
	return nil
}
