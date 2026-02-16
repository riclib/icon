SKILL_DIR := $(HOME)/.claude/skills/icon-search
BIN_DIR := $(SKILL_DIR)/bin
BINARY := $(BIN_DIR)/icon-search

.PHONY: build install-skill clean

build:
	go build -o icon-search ../cmd/icon-search

install-skill: build
	mkdir -p $(BIN_DIR)
	cp icon-search $(BINARY)
	cp SKILL.md $(SKILL_DIR)/SKILL.md
	@echo "Installed to $(SKILL_DIR)"

clean:
	rm -f icon-search
