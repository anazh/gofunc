package net

import (
	"sync"
	"time"
)

var isFirstConn bool = true

var WaitConn chan bool

func init() {
	WaitConn = make(chan bool)
}

type NetStatus struct {
	IsLinkNet bool
	lock      sync.RWMutex
}

var Nets NetStatus

func (n *NetStatus) NetStatus() bool {
	n.lock.RLock()
	defer n.lock.RUnlock()
	return n.IsLinkNet
}

func (n *NetStatus) netCheck() {
	for {
		n.IsLinkNet = isLinkNet()
		n.lock.Lock()
		if isFirstConn && n.IsLinkNet {
			firstLinkAction()
			isFirstConn = false
		}
		n.lock.Unlock()
		time.Sleep(10 * time.Second)
	}
}

func isLinkNet() bool {
	return IsLinkNet()
}

// 启机后第一次同步的操作
func firstLinkAction() {
	time.Sleep(1 * time.Second)
	close(WaitConn) //所有等待的都会释放
	// middle_control.RecordErr("net conn success version:" + version.VERSION)
	// middle_control.Upversion(version.VERSION)
}
