package engine

/*
* @
* @Author:
* @Date: 2020/3/23 19:24
 */
// 调度器接口
type AllScheduler interface {
	Submit(*Request)
	ConfigureMasterWorkerChan(chan *Request)
}
