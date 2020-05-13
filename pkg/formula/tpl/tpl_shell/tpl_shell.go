package tpl_shell

const (
	TemplateMain = `#!/bin/sh

. ./{{bin-name}}/{{bin-name}}.sh --source-only

run $SAMPLE_TEXT $SAMPLE_LIST $SAMPLE_BOOL`

	TemplateMakefile = `# SH
BINARY_NAME={{bin-name}}.sh
DIST=../dist
DIST_DIR=$(DIST)/commons/bin
build:
	mkdir -p $(DIST_DIR)
	cp main.sh $(DIST_DIR)/$(BINARY_NAME) && cp -r {{bin-name}} $(DIST_DIR) && cp Dockerfile $(DIST_DIR)
	chmod +x $(DIST_DIR)/$(BINARY_NAME)`

	TemplateDockerfile = `
FROM alpine:3.7

WORKDIR /usr/src/app

COPY . .

CMD chmod +x main.sh && ./main.sh`

	TemplateFileShell = `#!/bin/sh
run() {
  echo "Hello World! "
  echo "You receive $SAMPLE_TEXT in text. "
  echo "You receive $SAMPLE_LIST in list. " 
  echo "You receive $SAMPLE_BOOL in boolean. "  
}
`
)
