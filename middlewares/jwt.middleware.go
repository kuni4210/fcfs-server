package middlewares

import (
	"database/sql"
	"fcfs-server/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtMiddleware(postgres *sql.DB, secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorization 헤더에서 토큰 가져오기
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "authorization header is required"})
			c.Abort()
			return
		}

		// Bearer 제거
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		// 토큰 검증
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// HMAC 서명 방식인지 확인
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secretKey), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			c.Abort()
			return
		}

		// 토큰 클레임 추출
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token claims"})
			c.Abort()
			return
		}

		// 클레임에서 사용자 ID 추출
		userID, ok := claims["sub"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token claims"})
			c.Abort()
			return
		}

		// DB에서 사용자 조회
		var user models.User
		err = postgres.QueryRow("SELECT id, username FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "user not found"})
			c.Abort()
			return
		}

		// 사용자 정보를 컨텍스트에 저장
		c.Set("user", &user)
		c.Next()
	}
}
