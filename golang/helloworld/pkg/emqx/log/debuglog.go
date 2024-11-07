/**
 * @Author: birney.dong
 * @Date: 2024/11/4 09:09
 * @Description: debuglog
 */

package log

import (
	"fmt"
	"log"
)

type EMQXDebugLogger struct {
	Module string
}

func NewEMQXDebugLogger(module string) EMQXDebugLogger {
	return EMQXDebugLogger{Module: module}
}

// Println is the library provided NOOPLogger's
// implementation of the required interface function()
func (e EMQXDebugLogger) Println(v ...interface{}) {
	for _, vv := range v {
		log.Println(fmt.Sprintf("[%s] - %+v", e.Module, vv))
	}
}

// Printf is the library provided NOOPLogger's
// implementation of the required interface function(){}
func (e EMQXDebugLogger) Printf(format string, v ...interface{}) {
	for _, vv := range v {
		log.Println(fmt.Sprintf("[%s] - "+format, e.Module, vv))
	}
}
