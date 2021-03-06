package dirlock

import (
	"fmt"
	"os"
	"syscall"
)

// 定义一个DirLock 的struct
type DirLock struct {
	dir string // 目录路径，例如 /home/xxx/go/src
	f *os.File
}

// 新建一个DirLock
func New(dir string) *DirLock {
	return &DirLock{
		dir: dir,
	}
}

// 加锁操作
func (l *DirLock) Lock() error {
	f, err := os.Open(l.dir) // 获取文件描述符
	if err != nil {
		return err
	}
	l.f = f
	err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB) // 加上排它锁，当遇到文件加锁的情况直接返回Error
	if err != nil {
		return fmt.Errorf("cannot flock directory %s - %s", l.dir, err)
	}
	return nil
}

// 解锁操作
func (l *DirLock) Unlock() error {
	defer l.f.Close()
	return syscall.Flock(int(l.f.Fd()), syscall.LOCK_UN) // 释放 Flock 文件锁
}

/*
1, Flock 是建议性的锁，使用的时候需要制定 how 参数, 否则容易出现多个 goroutine 共用文件的问题
2, how 参数制定 LOCK_NB 之后，goroutine 遇到已加锁的Flock，不会阻塞，而是直接返回错误
*/