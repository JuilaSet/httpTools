package fileWatcher

import "github.com/fsnotify/fsnotify"

type IWatcher interface {
	Create(fsnotify.Event)
	Write(fsnotify.Event)
	Remove(fsnotify.Event)
	Rename(fsnotify.Event)
	Chmod(fsnotify.Event)
	FileNames() []string
}
