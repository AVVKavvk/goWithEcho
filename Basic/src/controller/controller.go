package controller

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)
 var e=echo.New();
 var v=validator.New()

var product = map[int]string{1: "apple", 2: "banana", 3: "cherry", 4: "date",5: "ederberry"}

func ServeHome(c echo.Context) error {
	return c.JSON(200, "<h1>vipin</h1>")

}
func GetProductById(c echo.Context) error {
	id,err:=strconv.Atoi(c.Param("id"));
	if err!=nil{
		return c.JSON(http.StatusNotFound,"Id not found");
	}
	str:=product[id];
	if len(str)==0{
		return c.JSON(http.StatusNotFound,"Product not found");
	}
	return c.JSON(http.StatusOK,str)

}

func GetAllProduct(c echo.Context) error {
	return c.JSON(200,product)
}

func CreateProduct(c echo.Context) error {
	type body struct{
		Name string `json:"proname" validate:"required,min=4"`;
		// Vendor string 	`json:"vendor" validate:"min=5,max=10"`
		// Email string `json:"email" validate:"required_with=vendor,email"`
		// URL string 	 `json:"url" validate:"url"`
		// Con int `json:"con" validate:"len=2"`
	}
	// type body struct{
	// 	Name string `json:"proname" validate:"required,min=4"`;
	// }
	var b body;
	if err:=c.Bind(&b); err!=nil{
		return err;
	}
	if err:=v.Struct(b); err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if len(b.Name)==0{
		return c.JSON(http.StatusBadRequest,"Invalid request body")
	} 
	index:=len(product)+1;
	product[index]=b.Name;
	return c.JSON(http.StatusOK,product[index])
}

func QueryParamDemo(c echo.Context) error {
	return c.JSON(http.StatusOK,c.QueryParam("older"))
}

func UpdateProductById(c echo.Context) error {
	type body struct{
		Name string `json:"proname" validate:"required,min=4"`;
	}
	var b body;
	if err:=c.Bind(&b); err!=nil{
		return err;
	}
	if err:=v.Struct(b); err!=nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if len(b.Name)==0{
		return c.JSON(http.StatusBadRequest,"Invalid request body")
	} 

	id:=c.Param("id");
	index,err:=strconv.Atoi(id);

	if  err!=nil {
		return c.JSON(http.StatusBadRequest,"Invalid id")
	}
	if index>len(product) || index<0{
		return c.JSON(http.StatusBadRequest,"Invalid id outOfBoundIndex")
	}
	product[index]=b.Name;
	return c.JSON(http.StatusOK,product[index])

}
func DeleteProductById(c echo.Context) error {
	id:=c.Param("id");
	index,err:=strconv.Atoi(id);
	if err!=nil {
		return c.JSON(http.StatusBadRequest,"Invalid id")
	}
	if index>len(product) || index<0{
		return c.JSON(http.StatusBadRequest,"Invalid id outOfBoundIndex")
	}
	delete(product,index);
	return c.JSON(http.StatusOK,product)
}