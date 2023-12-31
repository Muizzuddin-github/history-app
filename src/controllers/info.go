package controllers

import (
	"context"
	"crud/src/db"
	"crud/src/repository"
	"crud/src/requestbody"
	"crud/src/responsebody"
	"crud/src/utility"
	"path/filepath"
	"slices"

	"github.com/gofiber/fiber/v2"
)

var AddInfo fiber.Handler = func(c *fiber.Ctx) error {
	body := requestbody.Info{}
	c.BodyParser(&body)
	body.Created_at = utility.TimeNow()

	file, err := c.FormFile("image")
	if err == nil{
		ext := filepath.Ext(file.Filename)
		accExt := []string{".jpg",".png",".jpeg"}

		if !slices.Contains(accExt,ext){
			c.Status(fiber.StatusBadRequest)
			return c.JSON(responsebody.Err{
				Errors: []string{"ext file not allowed"},
			})
		}

		imageByte, err := utility.ReadByte(file)
		if err != nil{
			c.Status(fiber.StatusBadRequest)
			return c.JSON(responsebody.Err{
				Errors : []string{err.Error()},
			})
		}

		resJson, err := utility.UploadImageApi(imageByte,file.Filename)
		if err != nil{
			c.Status(fiber.StatusBadRequest)
			return c.JSON(responsebody.Err{
				Errors: []string{err.Error()},
			})
		}

		url := resJson.Image.File.Resource.Chain.Image
		body.Image = url

	}

	ctx := context.Background()
	categoryCol := repository.NewInfoRepo(db.GetCollection("category"))
	result,id, err := categoryCol.AddInfo(ctx,&body,c.Params("idCategory"))
	if err != nil{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(responsebody.Err{
			Errors: []string{err.Error()},
		})
	}

	if result.MatchedCount == 0{
		c.Status(fiber.StatusNotFound)
		return c.JSON(responsebody.Err{
			Errors: []string{"Category not found"},
		})
	}


	c.Status(fiber.StatusCreated)
	return c.JSON(responsebody.InsertDocument{
		Message: "Info created success",
		InsertedID: id,
	})
}


var DeleteInfo fiber.Handler = func(c *fiber.Ctx) error {

	idCategory := c.Params("idCategory")
	idInfo := c.Params("idInfo")

	ctx := context.Background()
	categoryCol := repository.NewInfoRepo(db.GetCollection("category"))
	result, err := categoryCol.DeleteInfo(ctx,idCategory,idInfo)
	if err != nil{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(responsebody.Err{
			Errors: []string{err.Error()},
		})
	}

	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(responsebody.Err{
			Errors: []string{err.Error()},
		})
	}


	c.Status(fiber.StatusOK)
	return c.JSON(responsebody.Msg{
		Message: "delete info success",
	})
}


var UpdateInfo fiber.Handler = func(c *fiber.Ctx) error {


	idCategory := c.Params("idCategory")
	idInfo := c.Params("idInfo")


	file, err := c.FormFile("image")
	if err != nil{

		body := requestbody.InfoUpdateNoImage{}
		c.BodyParser(&body)
		ctx := context.Background()
		categoryCol := repository.NewInfoRepo(db.GetCollection("category"))
		result, err := categoryCol.UpdateNoImage(ctx, idCategory,idInfo,&body)
		if err != nil{
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(responsebody.Err{
				Errors: []string{err.Error()},
			})
		}

		if result.MatchedCount == 0{
			c.Status(fiber.StatusNotFound)
			return c.JSON(responsebody.Err{
				Errors: []string{"not found"},
			})
		}

		c.Status(fiber.StatusOK)
		return c.JSON(responsebody.Msg{
			Message: "update info success",
		})
	}

	body := requestbody.InfoUpdateWithImage{}
	c.BodyParser(&body)

	file, err = c.FormFile("image")
	if err != nil{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(responsebody.Err{
			Errors: []string{err.Error()},
		})
	}

	ext := filepath.Ext(file.Filename)
	extAcc := []string{".jpg",".jpeg","png"}
	if !slices.Contains(extAcc,ext){
		c.Status(fiber.StatusBadRequest)
		return c.JSON(responsebody.Err{
			Errors: []string{"ext file not allowed"},
		})
	}


	imageByte, err := utility.ReadByte(file)
	if err != nil{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(responsebody.Err{
			Errors: []string{err.Error()},
		})
	}

	resJson, err := utility.UploadImageApi(imageByte,file.Filename)
	if err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(responsebody.Err{
			Errors: []string{err.Error()},
		})
	}

	url := resJson.Image.File.Resource.Chain.Image
	body.ImageUrl = url

	ctx := context.Background()
	categoryCol := repository.NewInfoRepo(db.GetCollection("category"))
	result, err := categoryCol.UpdateWithImage(ctx,idCategory,idInfo,&body)
	if err != nil{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(responsebody.Err{
			Errors: []string{err.Error()},
		})
	}

	if result.MatchedCount == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(responsebody.Err{
			Errors: []string{"not found"},
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(responsebody.Msg{
		Message: "update info success",
	})
}