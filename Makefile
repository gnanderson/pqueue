include $(GOROOT)/src/Make.inc

TARG=github.com/gnanderson/pqueue
GOFILES=\
	pqueue.go\

CLEANFILES+=example

include $(GOROOT)/src/Make.pkg

EXT=
ifeq ($(GOOS),windows)
EXT=.exe
endif

example: install example.go
	$(GC) example.go
	$(LD) -o $@$(EXT) example.$O