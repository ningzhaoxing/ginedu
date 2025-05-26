package test

import (
	"fmt"
	"gindeu/initialize"
	"gindeu/internal/modules/user/logic"
	"gindeu/internal/modules/user/models"

	//"gindeu/internal/user/logic"
	//"gindeu/internal/user/models"
	errors2 "gindeu/pkg/errors"
	"gindeu/pkg/globals"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"testing"
)

type Member struct {
	name string
	age  int
}

type Option func(member *Member)

func NewMember(option ...Option) *Member {
	member := new(Member)
	for _, o := range option {
		o(member)
	}
	return member
}

func NewDefaultMember(option ...Option) *Member {
	member := new(Member)
	for _, o := range option {
		o(member)
	}
	return member
}

func NameOption() Option {
	return func(member *Member) {
		member.name = "zs"
	}

}

func AgeOption() Option {
	return func(member *Member) {
		member.age = 3
	}
}
func (m *Member) GetName() string {
	return m.name
}

func (m *Member) SetName(v string) {
	m.name = v
}

func (m *Member) SetAge(v int) {
	m.age = v
}

func (m *Member) GetAge() int {
	return m.age
}

func TestObj(t *testing.T) {
	m := NewMember(NameOption())
	fmt.Println(m)
}

type User struct {
	gorm.Model
	Email        uint         `gorm:"index:,unique"`
	MemberNumber string       `gorm:"type:varchar;index:,unique"`
	Languages    []*Language  `gorm:"many2many:user_languages; foreignKey:email;References:id;"`
	CreditCards  []CreditCard `gorm:"foreignKey:UserNumber;references:MemberNumber"`
}

type Language struct {
	gorm.Model
	Name string
}

type CreditCard struct {
	gorm.Model
	Number     string `gorm:"type:varchar"`
	UserNumber string `gorm:"type:varchar;index:,unique"`
}

func TestMany(t *testing.T) {
	var err error
	err, globals.C = initialize.InitConfig()
	if err != nil {
		return
	}
	d := initialize.InitDb(globals.C)
	_ = d.AutoMigrate(&models.User{}, &models.Article{})
}

func TestCreate(t *testing.T) {
	var err error
	err, globals.C = initialize.InitConfig()
	if err != nil {
		t.Error(err)
		return
	}
	d := initialize.InitDb(globals.C)

	user := models.User{Name: "afa", Age: 1221, Articles: []models.Article{{Title: "明天会更好"}}}
	err = d.Create(&user).Error
	if err != nil {
		t.Error(err)
	}

	fmt.Println("sdf")
	t.Log(user.ID)

}

type ApiUser struct {
	Name     string       `json:"name"`
	Id       uint         `json:"id"`
	Articles []ApiArticle `gorm:"foreignKey:user_id" json:"articles"`
}

func (r *ApiArticle) TableName() string {
	return "articles"
}

type ApiArticle struct {
	gorm.Model
	UserId uint `json:"userId" gorm:"column:user_id"`
	//Title string `json:"title" gorm:"type:varchar(50);"`
}

func TestSelect(t *testing.T) {
	var err error
	err, globals.C = initialize.InitConfig()
	if err != nil {
		t.Error(err)
		return
	}
	d := initialize.InitDb(globals.C)

	var users []ApiUser
	d.Debug().Model(&models.User{}).Preload("Articles").Find(&users)
	//s, _ := json.Marshal(&users)
	//for _, user := range users {
	//fmt.Printf("%+v \n ", user)
	//}

	//fmt.Printf("%s", reflect.TypeOf(EventBus.New()).Kind().String())
	fmt.Errorf("%w and %w", errors.New("asfa"), errors.New("afasf"))
}

func TestEvent(t *testing.T) {

	//initialize.AppInit()

	globals.EventBus = initialize.InitEventBus()
	globals.EventBus.Publish(globals.DemoErrTopic, errors2.DemoError{Code: globals.CodeFailed, User: "xxx"})

	//globals.EventBus.Publish(globals.TestTopic, subscribes.TestTopic{A: 4, B: 4})

}

func TestInitDb(t *testing.T) {
	_, globals.C = initialize.InitConfig()
	initialize.InitDb(globals.C)
}

func TestUserChange(t *testing.T) {
	initialize.AppInit()
	var g *gin.Context
	userLogin := logic.NewUserLogic(globals.NewDefaultAppCtx(), g)
	userLogin.SetUserIntegral("zhang shan", 10)

}
