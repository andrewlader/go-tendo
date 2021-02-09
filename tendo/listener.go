package tendo

import (
	"log"

	"github.com/andrewlader/go-tendo/tendo/internal"
)

type listener struct {
	root         *root
	libChan      chan *internal.LibraryInfo
	classChan    chan *internal.ClassInfo
	methodChan   chan *internal.MethodInfo
	functionChan chan *internal.FunctionInfo
	quitChan     chan bool
	logger       *Logger
}

func newListener(root *root, logger *Logger) *listener {
	return &listener{
		root:         newRoot(),
		logger:       logger,
		libChan:      make(chan *internal.LibraryInfo),
		classChan:    make(chan *internal.ClassInfo),
		methodChan:   make(chan *internal.MethodInfo),
		functionChan: make(chan *internal.FunctionInfo),
		quitChan:     make(chan bool, 1),
	}
}

func (listener *listener) start() {
	defer listener.handleStop()

	keepGoing := true

	for keepGoing {
		select {
		case libInfo := <-listener.libChan:
			_, ok := listener.root.libraries[libInfo.Name]
			if !ok {
				listener.root.libraries[libInfo.Name] = newLibrary(libInfo.Name)
				listener.logger.printf(LogTrace, "Added package %s", libInfo.Name)
			}
			listener.root.currentLibrary = libInfo.Name

		case classInfo := <-listener.classChan:
			lib := listener.root.libraries[listener.root.currentLibrary]
			lib.addClass(classInfo.Name, listener.logger)

		case methodInfo := <-listener.methodChan:
			lib := listener.root.libraries[listener.root.currentLibrary]
			_, ok := lib.classes[methodInfo.Class]
			if !ok {
				lib.addClass(methodInfo.Class, listener.logger)
			}
			lib.classes[methodInfo.Class].addMethod(methodInfo.Name, listener.logger)

		case functionInfo := <-listener.functionChan:
			lib := listener.root.libraries[listener.root.currentLibrary]
			lib.addFunction(functionInfo.Name, listener.logger)

		case <-listener.quitChan:
			keepGoing = false
		}
	}
}

func (listener *listener) stop() {
	// all done, so shutdown
	listener.quitChan <- true
}

func (listener *listener) cleanup() {
	// cleanup
	listener.root = nil
	close(listener.libChan)
	close(listener.classChan)
	close(listener.methodChan)
	close(listener.functionChan)
	close(listener.quitChan)
}

func (listener *listener) restart() {
	listener.cleanup()

	listener.root = newRoot()
	listener.libChan = make(chan *internal.LibraryInfo)
	listener.classChan = make(chan *internal.ClassInfo)
	listener.methodChan = make(chan *internal.MethodInfo)
	listener.functionChan = make(chan *internal.FunctionInfo)
	listener.quitChan = make(chan bool, 1)
}

func (listener *listener) handleStop() {
	listener.logger.printfln(LogInfo, "Stopping the listener...")

	recovery := recover()
	if recovery != nil {
		if listener.logger != nil {
			listener.logger.printf(LogErrors, "Panic occurred:\n\n%v", recovery)
		} else {
			log.Printf("Panic occurred:\n\n%v", recovery)
		}
	}

	listener.logger.printfln(LogInfo, "Stopped the listener...")
}
