package fileWatcher

import (
	"github.com/fsnotify/fsnotify"
	"httpProxyDDD/src/domain/fileNotifier/vo"
)

// builder
type FileStatusWatcherBuilder func(w *FileStatusWatcher)
type FileStatusWatcherBuilders []FileStatusWatcherBuilder

func (bs FileStatusWatcherBuilders) Apply(w *FileStatusWatcher) {
	for _, b := range bs {
		b(w)
	}
}

// build methods
func WithFilename(filename string) FileStatusWatcherBuilder {
	return func(w *FileStatusWatcher) {
		w.WatchedFiles = append(w.WatchedFiles, vo.NewWatchedFile(filename))
	}
}

func WithCreateHandlers(handlers ...func(fsnotify.Event)) FileStatusWatcherBuilder {
	return func(w *FileStatusWatcher) {
		for _, h := range handlers {
			w.CreateHandlers = append(w.CreateHandlers, h)
		}
	}
}

func WithWriteHandlers(handlers ...func(fsnotify.Event)) FileStatusWatcherBuilder {
	return func(w *FileStatusWatcher) {
		for _, h := range handlers {
			w.WriteHandlers = append(w.WriteHandlers, h)
		}
	}
}

func WithRemoveHandlers(handlers ...func(fsnotify.Event)) FileStatusWatcherBuilder {
	return func(w *FileStatusWatcher) {
		for _, h := range handlers {
			w.RemoveHandlers = append(w.RemoveHandlers, h)
		}
	}
}

func WithRenameHandlers(handlers ...func(fsnotify.Event)) FileStatusWatcherBuilder {
	return func(w *FileStatusWatcher) {
		for _, h := range handlers {
			w.RenameHandlers = append(w.RenameHandlers, h)
		}
	}
}

func WithChmodHandlers(handlers ...func(fsnotify.Event)) FileStatusWatcherBuilder {
	return func(w *FileStatusWatcher) {
		for _, h := range handlers {
			w.ChmodHandlers = append(w.ChmodHandlers, h)
		}
	}
}

