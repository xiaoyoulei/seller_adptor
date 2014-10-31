package connpool

//package main

import (
	"container/list"
	"reflect"
	"sync"
	"utils"
)

type ConnPool struct {
	Dial    func() (interface{}, error)
	Alive   func(c interface{}) bool
	Close   func(c interface{}) error
	MaxIdle int // 最大连接数

	mu sync.Mutex // 对于connpool的锁

	active int // 当前活跃的连接数个数

	idle list.List
}

func (this *ConnPool) New(dial func() (interface{}, error), alive func(c interface{}) bool, closefc func(c interface{}) error, maxidle int) *ConnPool {
	return &ConnPool{Dial: dial, Alive: alive, Close: closefc, MaxIdle: maxidle}
}

func (this *ConnPool) Get() (c interface{}, err error) {
	this.mu.Lock()

	if this.active == 0 {
		utils.DebugLog.Write("create a new conn, now active[%d]", this.active)
		c, err = this.Dial()
		if err != nil {
			this.mu.Unlock()
			return
		}
		this.active += 1
	} else {
		findflag := false
		for this.idle.Len() > 0 {
			c = this.idle.Front().Value
			this.idle.Remove(this.idle.Back())
			utils.DebugLog.Write("get a interface GET type[%s]", reflect.TypeOf(c))
			this.mu.Unlock()
			isconn := this.Alive(c)
			this.mu.Lock()
			if isconn == false {
				this.active -= 1
				utils.DebugLog.Write("delete a conn, now active[%d]", this.active)
				continue
			} else {
				findflag = true
				this.active -= 1
				break
			}
		}
		if findflag == false {
			c, err = this.Dial()
			if err != nil {
				return
			} else {
				utils.DebugLog.Write("create a new conn. active[%d]", this.active)
			}
		}
	}

	this.mu.Unlock()
	return
}

func (this *ConnPool) Put(c interface{}) (err error) {
	this.mu.Lock()
	defer this.mu.Unlock()
	utils.DebugLog.Write("get a interface type[%s]", reflect.TypeOf(c))
	if this.active >= this.MaxIdle {
		err = this.Close(c)
		utils.DebugLog.Write("del a conn . active[%d]", this.active)
	} else {
		this.active += 1
		this.idle.PushBack(c)
		utils.DebugLog.Write("put a conn . active[%d]", this.active)
	}
	return
}

/*
func main() {
	pool := &ConnPool{
		Dial: func() (interface{}, error) {
			var c interface{}
			var err error
			return c, err
		},
		Alive: func(c interface{}) bool {
			return true
		},
		MaxIdle: 10,
	}
	fmt.Println(pool)
}
*/
