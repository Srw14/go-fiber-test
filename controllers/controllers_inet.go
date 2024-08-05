package controllers

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"go-fiber-test/database"
	m "go-fiber-test/models"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func HelloTestV2(c *fiber.Ctx) error {
	return c.SendString("Hello, World! V2")
}

func BodyParserTest(c *fiber.Ctx) error {
	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Name) // john
	log.Println(p.Pass) // doe
	str := p.Name + p.Pass
	return c.JSON(str)
}

func ParamsTest(c *fiber.Ctx) error {

	str := "hello ==> " + c.Params("name")
	return c.JSON(str)
}

func QueryTest(c *fiber.Ctx) error {
	a := c.Query("search")
	str := "my search is  " + a
	return c.JSON(str)
}

func ValidTest(c *fiber.Ctx) error {
	//Connect to database
	user := new(m.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}
	return c.JSON(user)
}

//CRUD

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //10ตัว

	var dataResults []m.DogsRes
	var redCount, greenCount, pinkCount, noColorCount int
	for _, v := range dogs { //1 inet 112 //2 inet1 113
		typeStr := ""
		if v.DogID >= 50 && v.DogID <= 100 {
			typeStr = "red"
			redCount++
		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
			greenCount++
		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
			pinkCount++
		} else {
			typeStr = "no color"
			noColorCount++
		}

		d := m.DogsRes{
			Name:  v.Name,  //inet
			DogID: v.DogID, //112
			Type:  typeStr, //no color
		}
		dataResults = append(dataResults, d)
	}

	r := m.ResultData{
		Data:        dataResults,
		Name:        "golang-test",
		Count:       len(dogs), //หาผลรวม,
		Red_Sum:     redCount,
		Green_Sum:   greenCount,
		Pink_Sum:    pinkCount,
		NoColor_Sum: noColorCount,
	}
	return c.Status(200).JSON(r)
}

func GetDelectDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&dogs)
	return c.Status(200).JSON(dogs)
}

func GetBetweenDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Where("dog_id BETWEEN ? AND ?", 50, 100).Find(&dogs)
	return c.Status(200).JSON(dogs)
}

//CRUD

// CRUD company
func GetComs(c *fiber.Ctx) error {
	db := database.DBConn
	var coms []m.Companys

	db.Find(&coms)
	return c.Status(200).JSON(coms)
}

func GetCom(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var com []m.Companys

	result := db.Find(&com, "company_id = ?", search)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&com)
}

func AddCom(c *fiber.Ctx) error {
	db := database.DBConn
	var com m.Companys

	if err := c.BodyParser(&com); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&com)
	return c.Status(201).JSON(com)
}

func UpdateCom(c *fiber.Ctx) error {
	db := database.DBConn
	var com m.Companys
	id := c.Params("id")

	if err := c.BodyParser(&com); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&com)
	return c.Status(200).JSON(com)
}

func RemoveCom(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var com m.Companys

	result := db.Delete(&com, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}

func GetDelectComs(c *fiber.Ctx) error {
	db := database.DBConn
	var coms []m.Companys

	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&coms)
	return c.Status(200).JSON(coms)
}

func GetBetweenComs(c *fiber.Ctx) error {
	db := database.DBConn
	var coms []m.Companys

	db.Where("company_id BETWEEN ? AND ?", 5, 10).Find(&coms)
	return c.Status(200).JSON(coms)
}

//CRUD company

// CRUD project
func GetUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []m.Users

	db.Find(&users)
	return c.Status(200).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var user []m.Users

	result := db.Find(&user, "company_id = ?", search)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&user)
}

func AddUser(c *fiber.Ctx) error {
	db := database.DBConn
	var user m.Users

	if err := c.BodyParser(&user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&user)
	return c.Status(201).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.DBConn
	var user m.Users
	id := c.Params("id")

	if err := c.BodyParser(&user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&user)
	return c.Status(200).JSON(user)
}

func RemoveUser(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var user m.Users

	result := db.Delete(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(400)
	}
	return c.SendStatus(200)
}

func GetUsersJson(c *fiber.Ctx) error {
	db := database.DBConn
	var users []m.Users

	db.Find(&users)

	var dataResults []m.UsersRes
	var GenZCount, GenYCount, GenXCount, BabyBoomerCount, GIGenCount int
	for _, v := range users {
		typeStr := ""
		if v.Age < 24 {
			typeStr = "GenZ"
			GenZCount++
		} else if v.Age >= 24 && v.Age <= 41 {
			typeStr = "GenY"
			GenYCount++
		} else if v.Age >= 42 && v.Age <= 56 {
			typeStr = "GenX"
			GenXCount++
		} else if v.Age >= 57 && v.Age <= 75 {
			typeStr = "Baby Boomer"
			BabyBoomerCount++
		} else {
			typeStr = "G.I. Generation"
			GIGenCount++
		}

		d := m.UsersRes{
			Name:       v.Name,
			EmployeeID: v.EmployeeID,
			Age:        v.Age,
			Type:       typeStr,
		}
		dataResults = append(dataResults, d)
	}
	r := m.ResultUsersData{
		Data:            dataResults,
		GenZ_Sum:        GenZCount,
		GenY_Sum:        GenYCount,
		GenX_Sum:        GenXCount,
		Baby_Boomer_Sum: BabyBoomerCount,
		GIGen_Sum:       GIGenCount,
	}
	return c.Status(200).JSON(r)
}

func GetDeleteUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []m.Users
	db.Unscoped().Where("deleted_at IS NO NULL").Find(&users)
	return c.Status(200).JSON(users)
}

func GetBetweenUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []m.Users
	db.Where("employee_id BETWEEN ? AND ?", 50, 10).Find(&users)
	return c.Status(200).JSON(users)
}

//CRUD project

// 5.1
func Factorial(c *fiber.Ctx) error {
	n := c.Params("num")
	i, _ := strconv.Atoi(n)

	result := factorial(i)

	return c.JSON(result)
}

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

// 5.1

// 5.2
func Ascii(c *fiber.Ctx) error {
	str := c.Query("tax_id")

	runes := []rune(str)

	result := []string{}

	for i := 0; i < len(runes); i++ {
		result = append(result, strconv.Itoa(int(runes[i])))
	}

	return c.JSON(fiber.Map{
		"ascii": result,
	})
}

// 5.2

//6

func Register(c *fiber.Ctx) error {
	info := new(m.Register)

	if err := c.BodyParser(info); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(info)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}

	if !isValidUsername(info.Username) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "สามารถใช่ได้แค่ A-Z a-z 0-9 และเครื่องหมาย_เท่านั้น",
		})
	}

	if !isValidWebname(info.Webname) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "สามารถใช่ได้แค่ A-Z a-z 0-9 และเครื่องหมาย-เท่านั้น",
		})
	}

	return c.JSON(info)

}

func isValidUsername(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return re.MatchString(username)
}

func isValidWebname(webname string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
	return re.MatchString(webname)
}

//6
