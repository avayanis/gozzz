########################
# Makefile Definitions #
########################
GO ?= go
MAKE ?= make
DEMOPID = tmp/demo.pid

#####################
# Color Definitions #
#####################
NO_COLOR    = \x1b[0m
OK_COLOR    = \x1b[32;01m
WARN_COLOR  = \x1b[33;01m
ERROR_COLOR = \x1b[31;01m

all: build

build:
	$(GO) build

test:
	$(GO) test

dev-build: dev-clean
	@echo "$(OK_COLOR)Building Demo Server...$(NO_COLOR)"
	mkdir tmp
	$(GO) build -o tmp/demo demo/example.go
	@echo ""

dev-clean:
	@echo "$(OK_COLOR)Cleaning Demo...$(NO_COLOR)"

	@if [ -e $(DEMOPID) ]; \
	then \
		echo "$(OK_COLOR)Cleaning Up Demo Server...$(NO_COLOR)"; \
		kill `cat $(DEMOPID)` || true; \
	fi;

	rm -rf tmp
	@echo ""

dev-run: dev-build
	@echo "$(OK_COLOR)Starting Demo Server...$(NO_COLOR)"

	@./tmp/demo & echo $$! > $(DEMOPID)
	@sleep 1
	@echo ""
	@read -p "Press any key to exit..."
	@echo "$(OK_COLOR)Shutting Down Demo Server...$(NO_COLOR)"
	@kill `cat $(DEMOPID)`
	@rm $(DEMOPID)
