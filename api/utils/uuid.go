package utils

import (
	"fmt"
	"github.com/google/uuid"
)

/*
生成UUID
 */
func NewUUID()(string,error)  {
	uuid :=uuid.New()
	uuid[8] = uuid[8]&^0xc0 | 0x80 //变体位
	uuid[6] = uuid[6]&^0xf0 | 0x40 //版本4
	return fmt.Sprintf("%x-%x-%x-%x-%x",uuid[0:4],uuid[4:6],uuid[6:8],uuid[8:10],uuid[10:]),nil
}


