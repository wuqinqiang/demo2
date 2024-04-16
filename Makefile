.PHONY: checkbuf
checkbuf:
ifeq (,$(shell which buf))
	@echo 'error:bufbuild/buf is not installed ,please exec `brew install bufbuild/buf/buf`'
	@echo 1
endif

.PHONY: generate
generate: checkbuf
	buf generate --path proto/demo2.proto