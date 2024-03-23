package auth

type Permission string

const (
	USER_CREATE Permission = "USER_CREATE"
)

func HasPermission(permission Permission, ctxPermissions []interface{}) bool {
	for _, perm := range ctxPermissions {
		if string(permission) == perm.(string) {
			return true
		}
	}
	return false
}
