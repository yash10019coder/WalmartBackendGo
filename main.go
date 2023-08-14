package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/yash10019coder/WalmartGolangBackend/database"
)

const (
	PORT = "8080"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageurl"`
	CategoryId  string `json:"category"`
	StoreId     string `json:"store"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Store struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	Phonenumber string `json:"phonenumber"`
}

type Maps struct {
	ID       string `json:"id"`
	StoreID  string `json:"storeid"`
	ImageUrl string `json:"imageurl"`
}

type Banner struct {
	ID        string `json:"id"`
	StoreID   string `json:"storeid"`
	ImageUrl  string `json:"imageurl"`
	ProductID string `json:"productid"`
}

type DealsOftheDay struct {
	ID         string `json:"id"`
	StoreID    string `json:"storeid"`
	ProuductID string `json:"productid"`
}

var db *sql.DB

func GetUser(context *gin.Context) {
	var id string = context.Param("id")
	var user User
	err := db.QueryRow("SELECT * FROM user WHERE id = %s", id).Scan(&user.ID, &user.Email, &user.Name, &user.Phone)
	if err != nil {
		panic(err)
	}
	context.JSON(200, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"phone": user.Phone,
	})
}

func CreateUser(context *gin.Context) {
	var user User
	err := context.BindJSON(&user)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO user (name, email, phone) VALUES ($1, $2, $3)", user.Name, user.Email, user.Phone)
	if err != nil {
		panic(err)
	}
	context.JSON(200, gin.H{
		"message": "success",
	})
}
func GetProduct(context *gin.Context) {
	var id string = context.Query("id")
	var storeid string = context.Query("storeid")
	var product Product
	query := `SELECT product.id, product.name, product.price, product.description, product.imageurl, category.name, store.name
FROM product,
     category,
     store
WHERE product.id = $1
  AND product.storeid = $2`
	err := db.QueryRow(query, id, storeid).Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.ImageUrl, &product.CategoryId, &product.StoreId)
	if err != nil {
		panic(err)
	}
	context.JSON(200, gin.H{
		"product": product,
	})
}

func CreateProduct(context *gin.Context) {
	var product Product
	err := context.BindJSON(&product)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO product (name, imageurl, price, category) VALUES ($1, $2, $3, $4)", product.Name, product.ImageUrl, product.Price, product.CategoryId)
	if err != nil {
		panic(err)
	}
	context.JSON(200, gin.H{
		"message": "success",
		"product": product,
	})
}
func GetProducts(context *gin.Context) {
	var products []Product
	var storeId string = context.Query("storeid")

	// Updated SQL query with placeholders and join condition
	query := `
		SELECT product.id, product.name, product.imageurl, product.price, product.categoryid
		FROM product
		INNER JOIN store ON product.storeid = store.id
		WHERE store.id = $1
	`

	rows, err := db.Query(query, storeId)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.ImageUrl, &product.Price, &product.CategoryId)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}

	context.JSON(200, gin.H{
		"products": products,
	})
}

func GetCategory(context *gin.Context) {
	var categories []Category
	rows, err := db.Query("SELECT * FROM category")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var category Category
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}
	context.JSON(200, gin.H{
		"categories": categories,
	})
}

func CreateCategory(context *gin.Context) {
	var category Category
	err := context.BindJSON(&category)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO category (name) VALUES ($1)", category.Name)
	if err != nil {
		panic(err)
	}
	context.JSON(200, gin.H{
		"message": "success",
	})
}

func GetStore(context *gin.Context) {
	var stores []Store
	rows, err := db.Query("SELECT * FROM store")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var store Store
		err = rows.Scan(&store.ID, &store.Name, &store.Location, &store.Phonenumber)
		if err != nil {
			panic(err)
		}
		stores = append(stores, store)
	}
	context.JSON(200, gin.H{
		"stores": stores,
	})
}

func CreateStore(context *gin.Context) {
	var store Store
	err := context.BindJSON(&store)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO store (name, location, phonenumber) VALUES ($1, $2, $3)", store.Name, store.Location, store.Phonenumber)
	if err != nil {
		panic(err)
	}
	context.JSON(200, gin.H{
		"message": "success",
	})
}

func GetMaps(context *gin.Context) {
	var maps []Maps
	rows, err := db.Query("SELECT * FROM maps")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var map_ Maps
		err = rows.Scan(&map_.ID, &map_.StoreID, &map_.ImageUrl)
		if err != nil {
			panic(err)
		}
		maps = append(maps, map_)
	}
	context.JSON(200, gin.H{
		"maps": maps,
	})
}

func CreateMaps(context *gin.Context) {
	var map_ Maps
	err := context.BindJSON(&map_)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO maps (storeid, imageurl) VALUES ($1, $2)", map_.StoreID, map_.ImageUrl)
	if err != nil {
		panic(err)
	}
	context.JSON(200, gin.H{
		"message": "success",
	})
}

func GetBanner(context *gin.Context) {
	var banners []Banner
	rows, err := db.Query("SELECT * FROM banner")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var banner Banner
		err = rows.Scan(&banner.ID, &banner.StoreID, &banner.ImageUrl, &banner.ProductID)
		if err != nil {
			panic(err)
		}
		banners = append(banners, banner)
	}
	context.JSON(200, gin.H{
		"banners": banners,
	})
}

func CreateBanner(context *gin.Context) {
	var banner Banner
	err := context.BindJSON(&banner)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO banner (storeid, imageurl, productid) VALUES ($1, $2, $3)", banner.StoreID, banner.ImageUrl, banner.ProductID)
	if err != nil {
		panic(err)
	}
	context.JSON(200, gin.H{
		"message": "success",
	})
}

func GetDealsOftheDay(context *gin.Context) {
	var deals []DealsOftheDay
	rows, err := db.Query("SELECT * FROM deals_of_the_day")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var deal DealsOftheDay
		err = rows.Scan(&deal.ID, &deal.StoreID, &deal.ProuductID)
		if err != nil {
			panic(err)
		}
		deals = append(deals, deal)
	}
	context.JSON(200, gin.H{
		"deals": deals,
	})
}

func CreateDealsOftheDay(context *gin.Context) {
	var deal DealsOftheDay
	err := context.BindJSON(&deal)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO deals_of_the_day (storeid, productid) VALUES ($1, $2)", deal.StoreID, deal.ProuductID)
	if err != nil {
		panic(err)
	}
	context.JSON(200, gin.H{
		"message": "success",
	})
}

func main() {
	db = database.SetupDB()
	router := gin.Default()

	router.GET("user/:id", GetUser)
	router.POST("user", CreateUser)

	router.GET("product/all", GetProducts)
	router.GET("product", GetProduct)
	router.POST("product", CreateProduct)

	router.GET("category", GetCategory)
	router.POST("category", CreateCategory)

	router.GET("store", GetStore)
	router.POST("store", CreateStore)

	router.GET("maps", GetMaps)
	router.POST("maps", CreateMaps)

	router.GET("banner", GetBanner)
	router.POST("banner", CreateBanner)

	router.GET("dealsoftheday", GetDealsOftheDay)
	router.POST("dealsoftheday", CreateDealsOftheDay)

	router.Run(":" + PORT)

}
