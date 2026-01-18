package main

import (
	"fmt"
	"hello/config"
	"hello/database"
	"hello/models"
	"hello/repositories"
	"log"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type TestData struct {
	Name     string
	Email    string
	Password string
	Phone    string
	Age      int
	Status   int
}

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default values")
	}

	fmt.Println("========================================")
	fmt.Println("用户数据初始化程序")
	fmt.Println("========================================")
	fmt.Println()

	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	fmt.Println("✓ 数据库连接成功")

	db := database.GetDB()
	userRepo := repositories.NewUserRepository(db)

	// Delete all existing users
	fmt.Println()
	fmt.Println("正在清理现有用户数据...")
	if err := db.Exec("DELETE FROM users").Error; err != nil {
		log.Fatalf("删除用户数据失败: %v", err)
	}
	fmt.Println("✓ 已清理所有现有用户数据")

	// Reset auto increment
	db.Exec("ALTER TABLE users AUTO_INCREMENT = 1")

	// Hash passwords
	adminPassword := hashPassword("admin123")
	userPassword := hashPassword("password123")

	// Prepare test data
	now := time.Now()
	testUsers := []TestData{
		{"系统管理员", "admin@example.com", adminPassword, "13800138000", 35, 1},
		{"张三", "zhangsan@example.com", userPassword, "13800138001", 28, 1},
		{"李四", "lisi@example.com", userPassword, "13800138002", 32, 1},
		{"王五", "wangwu@example.com", userPassword, "13800138003", 25, 1},
		{"赵六", "zhaoliu@example.com", userPassword, "13800138004", 30, 1},
		{"钱七", "qianqi@example.com", userPassword, "13800138005", 27, 1},
		{"孙八", "sunba@example.com", userPassword, "13800138006", 29, 1},
		{"周九", "zhoujiu@example.com", userPassword, "13800138007", 31, 1},
		{"吴十", "wushi@example.com", userPassword, "13800138008", 26, 1},
		{"郑十一", "zhengshiyi@example.com", userPassword, "13800138009", 33, 1},
		{"王十二", "wangshier@example.com", userPassword, "13800138010", 24, 1},
	}

	// Insert test users
	fmt.Println()
	fmt.Println("正在插入测试数据...")
	fmt.Println()

	for i, data := range testUsers {
		user := &models.User{
			Name:      data.Name,
			Email:     data.Email,
			Password:  data.Password,
			Phone:     data.Phone,
			Age:       data.Age,
			Status:    data.Status,
			CreatedAt: now,
			UpdatedAt: now,
		}

		if err := userRepo.Create(user); err != nil {
			log.Printf("插入用户 %s 失败: %v", data.Name, err)
		} else {
			fmt.Printf("  [%d] ✓ %s (%s)\n", i+1, data.Name, data.Email)
		}
	}

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("数据初始化完成!")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("测试账号信息:")
	fmt.Println()
	fmt.Println("管理员账号:")
	fmt.Println("  邮箱: admin@example.com")
	fmt.Println("  密码: admin123")
	fmt.Println()
	fmt.Println("普通用户账号 (所有用户密码都是 password123):")
	for i, data := range testUsers[1:] {
		fmt.Printf("  [%d] %s - %s\n", i+1, data.Name, data.Email)
	}
	fmt.Println()
	fmt.Println("共创建了", len(testUsers), "个测试用户")
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("密码加密失败: " + err.Error())
	}
	return string(hashedPassword)
}
