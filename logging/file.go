package logging

import (
	"fmt"
	"time"
)

func GetLogFileName() string {
	return fmt.Sprintf("log%s.log" , time.Now().Format("2006-01-02"))
}
