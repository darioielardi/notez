package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
	
	"notez/database"
	"notez/models"
	"notez/utils"
	"notez/utils/enums"
	"notez/utils/response"
)

func CheckAuth(next http.Handler, required bool, roles enums.Roles) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		if required && roles == nil {
			log.Fatalln("Auth middleware called with auth required but no roles set")
		}
		
		if !required && roles != nil {
			log.Fatalln("Auth middleware called with auth not required but roles set")
		}
		
		if !required {
			next.ServeHTTP(w, r)
			return
		}
		
		bearer := r.Header.Get("Authorization")
		
		if bearer == "" {
			response.SendError(w, "Unauthorized", 401)
			return
		}
		
		split := strings.Split(bearer, "Bearer ")
		if len(split) < 2 {
			response.SendError(w, "Unauthorized", 401)
			return
		}
		
		idToken := split[1]
		
		fb := utils.GetFb()
		ctx := context.Background()
		
		client, err := fb.Auth(ctx)
		if err != nil {
			response.SendError(w, "Internal server error", 500)
			return
		}
		
		token, err := client.VerifyIDToken(ctx, idToken)
		if err != nil {
			response.SendError(w, "Invalid token", 401)
			return
		}
		
		user := &models.User{}
		
		db := database.GetDB()
		
		db.Where("auth_id = ?", token.UID).First(user)
		
		if db.RecordNotFound() {
			response.SendError(w, "This user does not exists", 401)
			return
		}
		
		if !roles.Contains(&user.Role) {
			response.SendError(w, "Forbidden", 403)
			return
		}
		
		newCtx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(newCtx))
	}
}
