package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

type User struct {
	Id       int    `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Sex      string `json:"sex" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var db = make(map[string]string)

var secret = os.Getenv("SECRET")
var port = os.Getenv("PORT")

// LoadPublicKey 从PEM编码的公钥文件中加载公钥
func LoadPublicKey(filename string) (*rsa.PublicKey, error) {
	pemData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), nil
}

// VerifySignature 验证签名
func VerifySignature(pubKey *rsa.PublicKey, data []byte, signature string) (bool, error) {
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}
	hash := md5.Sum(data)
	hashedData := hash[:]
	//使用MD5WithSHA算法验签
	if err := rsa.VerifyPKCS1v15(pubKey, crypto.MD5, hashedData, signatureBytes); err != nil {
		return false, err
	}
	return true, nil
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	publicKey, err := LoadPublicKey("public.pem")
	if err != nil {
		panic(err)
	}
	// Ping test
	r.GET("/user/list", func(c *gin.Context) {
		signature := c.GetHeader("X-Sign")
		timestamp := c.GetHeader("X-Timestamp")
		//bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		verifyString := "secret=" + secret + "&timestamp=" + timestamp
		valid, err := VerifySignature(publicKey, []byte(verifyString), signature)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Verification error"})
			return
		}
		if valid {
			//创建User列表并填入三个模拟数据
			users := []User{
				{1, "zhangsan", "zhangsan", "M", 20, "", "12345"},
				{2, "lisi", "lisi", "F", 34, "", "12345"},
				{3, "wangwu", "wangwu", "M", 19, "", "12345"},
			}
			for i := range users {
				users[i].Password = ""
			}
			c.JSON(http.StatusOK, gin.H{"code": "0", "message": "success", "data": users})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid signature"})
		}
	})
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:1538
	r.Run(":" + port)
}
