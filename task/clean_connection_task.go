package task

import (
	"fmt"
	"l3n641/customer_service_system_backend/ws"
	"runtime/debug"
	"time"
)

func Init() {
	Timer(3*time.Second, 30*time.Second, cleanTimeoutConnection, "", nil, nil)

}

// 清理超时连接
func cleanTimeoutConnection(param interface{}) (result bool) {
	result = true

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ClearTimeoutConnections stop", r, string(debug.Stack()))
		}
	}()

	fmt.Println("定时任务，清理超时连接", param)
	ws.ClearUserTimeoutConnections()
	ws.ClearVisitorTimeoutConnections()

	return
}
