package main

import "log"

/*
流控
防止链接和宽带被消耗掉
bucket token算法
当由新的request进来时，往bucket里写入1个；离开时再释放掉
 */


type ConnLimiter struct {
	concurrentCon int // 同时链接数量（最大链接数）
	bucket chan int
}

//初始化
func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{concurrentCon: cc,bucket: make(chan int,cc)} //buffer channel,容量是cc
}

//判断token是否已经拿到
func (cl *ConnLimiter) GetCon() bool {
	//判断bucket里面的链接数量有没有超过最大链接数
	if len(cl.bucket) >= cl.concurrentCon{
		log.Printf("Reached the rate limitation.")
		return false
	}
	cl.bucket <- 1
	return true
}
//request离开时，bucket释放掉
func (cl *ConnLimiter) ReleaseConn()  {
	c :=<- cl.bucket
	log.Printf("New connection coming:%d",c)
}