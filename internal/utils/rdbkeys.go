package utils

import "fmt"

func GetAuthKey(sessionID string) string {
	return fmt.Sprintf("session_auth:%s", sessionID)
}

func GetSessionKey(username string) string {
	return fmt.Sprintf("session_id:%s", username)
}
