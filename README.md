# RBAC

`model.go`

```go
var (
	Viewer = "GET"
	Editor = "^(POST|PUT|PATCH)$"
	Admin  = "*"
)

type User struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Roles []Role `gorm:"many2many:user_roles;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Role struct {
	ID          uint         `gorm:"primaryKey;autoIncrement"`
	Name        string       `gorm:"type:varchar(255);not null;uniqueIndex"`
	Permissions []Permission `gorm:"many2many:role_permissions;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Permission struct {
	ID uint `gorm:"primaryKey;autoIncrement"`

	// GET, POST
	Action string `gorm:"type:varchar(255);not null;index:idx_action_route;check:action_check"`

	// e.g. /api/v1/user, /api/v1/user/*
	Route string `gorm:"type:varchar(255);not null;index:idx_action_route;check:route_check"`

	// Allow or Deny
	Allowed bool `gorm:"not null"`
}

func (Permission) TableName() string {
	return "permissions"
}

func (Permission) CheckConstraints(db *gorm.DB) {
	db.Exec("ALTER TABLE permissions ADD CONSTRAINT action_check CHECK (Action IN ('GET', 'POST', 'PUT', 'PATCH', 'DELETE'))")
	db.Exec("ALTER TABLE permissions ADD CONSTRAINT route_check CHECK (Route LIKE '/%')")
}
```

當用戶登入後，他將會收到一個token。這個token應該在之後的請求中都包含在 HTTP 的 Authorization header 中，以進行驗證。

以下是 Gin 在收到請求後如何實作檢查權限：

1. **驗證Token**:
   在 `middleware.AuthMiddleware` 中，首先從請求的 `Authorization` header 中取出 token。之後會使用 `parseToken` 函數來解析這個 token 並取得用戶的 ID。
   若 token 不存在、不合法或過期，返回401 Unauthorized錯誤。否則，將 userId 設置在 Gin 的 context 中，以便於後續的 middleware 或處理函數中使用。

2. **檢查權限**:
   `middleware.CheckPermission`中間件會使用之前在context中設置的userId來從數據庫中取得用戶的資訊及其角色和權限
   每一個角色都有其對應的一組權限，權限確定了用戶可以訪問的 HTTP method 和 route

   透過兩個for迴圈，檢查用戶所屬的所有角色和其相對應的權限，看是否匹配當前 request 的 HTTP method 和 route

   如果找到匹配的權限且該權限被允許，將繼續執行後續的處理函數；否則，返回 403 Permission Denied

```go
func CheckPermission(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		matchedAny := false
		userId := c.MustGet("userId").(uint) // get the userId which is this after token authentication
		var user models.User
		db.Preload("Roles").Preload("Roles.Permissions").Where("ID = ?", userId).First(&user)

		for _, role := range user.Roles {
			for _, permission := range role.Permissions {
				if matches(c.Request.Method, permission.Action) && matches(c.Request.URL.Path, permission.Route) {
					matchedAny = true
					if !permission.Allowed {
						c.AbortWithStatusJSON(403, gin.H{"error": "Permission denied"})
						return
					}
				}
			}
		}
		if !matchedAny {
			c.AbortWithStatusJSON(403, gin.H{"error": "Permission denied"})
			return
		}
		c.Next()
	}
}

// case insensitive regex matching
func matches(requestValue, patternValue string) (matched bool) {
	// case insensitive: https://stackoverflow.com/a/9655186
	matched, _ = regexp.MatchString("(?i)"+patternValue, requestValue)
	return
}
```
