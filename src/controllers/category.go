package controllers

import (
	"context"
	"crud/src/db"
	"crud/src/entity"
	"crud/src/repository"
	"crud/src/requestbody"
	"crud/src/responsebody"
	"crud/src/utility"
	"crud/src/validation"

	"github.com/gofiber/fiber/v2"
)

var AddCategory fiber.Handler = func(c *fiber.Ctx) error {

	body := requestbody.AddCategory{}
	c.BodyParser(&body)

	val := validation.Category(&body)
	if len(val) > 0{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(responsebody.Err{
			Errors: val,
		})
	}

	body.Created_at = utility.TimeNow()
	body.Info = []entity.Info{}

	ctx := context.Background()
	categoryCol := repository.NewCategoryRepo(db.GetCollection("category"))
	id, err := categoryCol.AddCategory(ctx,&body)

	if err != nil{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(responsebody.Err{
			Errors: []string{err.Error()},
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(responsebody.InsertDocument{
		Message: "Category created",
		InsertedID: id,
	})
}


var GetAllCategory fiber.Handler = func(c *fiber.Ctx) error {

	ctx := context.Background()
	categoryCol := repository.NewCategoryRepo(db.GetCollection("category"))
	result, err := categoryCol.GetCategory(ctx)
	if err != nil{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(responsebody.Err{
			Errors: []string{err.Error()},
		})
	}


	c.Status(fiber.StatusOK)
	return c.JSON(responsebody.Data{
		Message: "all category",
		Data: result,
	})
}

var UpdateCategory fiber.Handler = func(c *fiber.Ctx) error {
	id := c.Params("id")

	body := requestbody.UpdateCategory{}
	c.BodyParser(&body)

	val := validation.CategoryUp(&body)
	if len(val) > 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(responsebody.Err{
			Errors: val,
		})
	}

	ctx := context.Background()
	categoryCol := repository.NewCategoryRepo(db.GetCollection("category"))
	result, err := categoryCol.UpdateCategory(ctx,&body,id)
	if err != nil{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(responsebody.Err{
			Errors: []string{err.Error()},
		})
	}

	if result.MatchedCount == 0{
		c.Status(fiber.StatusNotFound)
		return c.JSON(responsebody.Err{
			Errors: []string{"category not found"},
		})
	}
	
	c.Status(fiber.StatusOK)
	return c.JSON(responsebody.Msg{
		Message: "Update category success",
	})
}

var DeleteCategory fiber.Handler = func(c *fiber.Ctx) error {
	id := c.Params("id")

	ctx := context.Background()
	categoryCol := repository.NewCategoryRepo(db.GetCollection("category"))
	result, err := categoryCol.DeleteCategory(ctx,id)
	if err != nil{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(responsebody.Err{
			Errors: []string{err.Error()},
		})
	}

	if result.DeletedCount == 0{
		c.Status(fiber.StatusNotFound)
		return c.JSON(responsebody.Err{
			Errors: []string{"Not found"},
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(responsebody.Msg{
		Message: "Delete success",
	})

}