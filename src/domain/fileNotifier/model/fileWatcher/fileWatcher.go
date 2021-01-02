package fileWatcher

import (
	"github.com/fsnotify/fsnotify"
	"httpTools/src/domain/fileNotifier/vo"
)

type FileStatusWatcher struct {
	vo.WatchedFiles

	CreateHandlers []func(fsnotify.Event)
	WriteHandlers  []func(fsnotify.Event)
	RemoveHandlers []func(fsnotify.Event)
	RenameHandlers []func(fsnotify.Event)
	ChmodHandlers  []func(fsnotify.Event)
}

func NewFileStatusWatcher(builders ...FileStatusWatcherBuilder) *FileStatusWatcher {
	w := &FileStatusWatcher{
		WatchedFiles: vo.WatchedFiles{},
	}
	FileStatusWatcherBuilders(builders).Apply(w)
	return w
}

// methods
func (w *FileStatusWatcher) Create(ev fsnotify.Event) {
	for _, h := range w.CreateHandlers {
		h(ev)
	}
}

func (w *FileStatusWatcher) Write(ev fsnotify.Event) {
	for _, h := range w.WriteHandlers {
		h(ev)
	}
}

func (w *FileStatusWatcher) Remove(ev fsnotify.Event) {
	for _, h := range w.RemoveHandlers {
		h(ev)
	}
}

func (w *FileStatusWatcher) Rename(ev fsnotify.Event) {
	for _, h := range w.RenameHandlers {
		h(ev)
	}
}

func (w *FileStatusWatcher) Chmod(ev fsnotify.Event) {
	for _, h := range w.ChmodHandlers {
		h(ev)
	}
}

func (w *FileStatusWatcher) FileNames() (arr []string) {
	h := func(f *vo.VWatchedFile) string {
		return f.Filename
	}
	for _, v := range w.WatchedFiles {
		arr = append(arr, h(v))
	}
	return arr
}
