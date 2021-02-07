package tendo

import "github.com/andrewlader/go-tendo/tendo/internal"

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
		root:         root,
		logger:       logger,
		libChan:      make(chan *internal.LibraryInfo),
		classChan:    make(chan *internal.ClassInfo),
		methodChan:   make(chan *internal.MethodInfo),
		functionChan: make(chan *internal.FunctionInfo),
		quitChan:     make(chan bool),
	}
}

func (listener *listener) Listen() {
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

	listener.logger.printfln(LogInfo, "Shutting down the listener...")
	listener.logger.printfln(LogInfo, "Shutdown of dispatcher is complete...")
}
