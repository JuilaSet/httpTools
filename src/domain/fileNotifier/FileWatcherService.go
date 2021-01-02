package fileNotifier

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

type FileWatcherService struct {
	fileWatcher IWatcher
	watcher     *fsnotify.Watcher
	waitCh      chan bool
}

func NewFileWatcherService(fileWatcher IWatcher) *FileWatcherService {
	// 创建一个监控对象
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	// 添加要监控的对象，文件或文件夹
	for _, file := range fileWatcher.FileNames() {
		if err := watcher.Add(file); err != nil {
			log.Fatal(err)
		}
	}
	return &FileWatcherService{
		fileWatcher: fileWatcher,
		watcher:     watcher,
		waitCh:      make(chan bool),
	}
}

func (w *FileWatcherService) Quit() {
	w.waitCh <- true
}

func (w *FileWatcherService) Run() error {
	defer w.watcher.Close()

	// 我们另启一个goroutine来处理监控对象的事件
	go func() {
		log.Println("Starting watch server....")
		for {
			select {
			case ev := <-w.watcher.Events:
				{
					// 判断事件发生的类型，如下5种
					// Create 创建
					// Write 写入
					// Remove 删除
					// Rename 重命名
					// Chmod 修改权限
					switch {
					case ev.Op&fsnotify.Create == fsnotify.Create:
						w.fileWatcher.Create(ev)
					case ev.Op&fsnotify.Write == fsnotify.Write:
						w.fileWatcher.Write(ev)
					case ev.Op&fsnotify.Remove == fsnotify.Remove:
						w.fileWatcher.Remove(ev)
					case ev.Op&fsnotify.Rename == fsnotify.Rename:
						w.fileWatcher.Rename(ev)
					case ev.Op&fsnotify.Chmod == fsnotify.Chmod:
						w.fileWatcher.Chmod(ev)
					}
				}
			case err := <-w.watcher.Errors:
				{
					log.Fatal(err)
				}
			}
		}
	}()

	// 等待
	<-w.waitCh
	return nil
}
