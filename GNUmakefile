#!/usr/bin/env make -f

binassets_develgo         = src/share/assets/bindata.devel.go
binassets_productiongo    = src/share/assets/bindata.production.go
bintemplates_develgo      = src/share/templates.html/bindata.devel.go
bintemplates_productiongo = src/share/templates.html/bindata.production.go
templates_html=$(shell echo src/share/templates.html/{index,usepercent,tooltipable}.html)
bindir=bin/$(shell uname -sm | awk '{ sub(/x86_64/, "amd64", $$2); print tolower($$1) "_" $$2; }')

.PHONY: all bootstrap bootstrap_develgo
all: $(bindir)/ostent
bootstrap:
#	go get -v ostent ostent/boot
	go get -v ./... github.com/jteeuwen/go-bindata/go-bindata github.com/skelterjohn/rerun
#	go get -v -tags production ostent
	go get -v -tags production ./...
	$(MAKE) $(MFLAGS) bootstrap_develgo
bootstrap_develgo: $(binassets_develgo) $(bintemplates_develgo)

%: %.sh # clear the implicit *.sh rule covering ./ostent.sh

$(bindir)/%:
	@echo '* Sources:' $^
	go build -o $@ $(patsubst src////%,%,$|)

$(bindir)/amberpp: | src////amberp/amberpp
$(bindir)/ostent:  | src////ostent

$(bindir)/amberpp: $(shell go list -f '\
{{$$dir := .Dir}}\
{{range .GoFiles }}{{$$dir}}/{{.}}{{"\n"}}{{end}}' amberp/amberpp | \
sed -n "s,^ *,,g; s,$(PWD)/,,p" | sort) # | tee /dev/stderr

$(bindir)/ostent: $(shell \
go list -tags production -f '{{.ImportPath}}{{"\n"}}{{join .Deps "\n"}}' ostent | xargs \
go list -tags production -f '{{if and (not .Standard) (not .Goroot)}}\
{{$$dir := .Dir}}\
{{range .GoFiles     }}{{$$dir}}/{{.}}{{"\n"}}{{end}}\
{{range .CgoFiles    }}{{$$dir}}/{{.}}{{"\n"}}{{end}}{{end}}' | \
sed -n "s,^ *,,g; s,$(PWD)/,,p" | sort) # | tee /dev/stderr
#	@echo '* Sources:' $^
	go build -tags production -o $@ ostent

$(bindir)/jsmakerule: $(binassets_develgo) $(shell \
go list -f '{{.ImportPath}}{{"\n"}}{{join .Deps "\n"}}' share/assets/jsmakerule | xargs \
go list -f '{{if and (not .Standard) (not .Goroot)}}\
{{$$dir := .Dir}}\
{{range .GoFiles     }}{{$$dir}}/{{.}}{{"\n"}}{{end}}\
{{range .CgoFiles    }}{{$$dir}}/{{.}}{{"\n"}}{{end}}{{end}}' | \
sed -n "s,^ *,,g; s,$(PWD)/,,p" | sort) # | tee /dev/stderr
#	@echo '* Sources:' $^
	@echo '* Prerequisite: bin-jsmakerule'
	go build -o $@ share/assets/jsmakerule

src/share/tmp/jsassets.d: # $(bindir)/jsmakerule
	@echo '* Prerequisite: src/share/tmp/jsassets.d'
#	$(MAKE) $(MFLAGS) $(bindir)/jsmakerule
	$(bindir)/jsmakerule src/share/assets/js/production/ugly/index.js >$@
#	$^ src/share/assets/js/production/ugly/index.js >$@
ifneq ($(MAKECMDGOALS), clean)
include src/share/tmp/jsassets.d
endif
src/share/assets/js/production/ugly/index.js:
	@echo    @uglifyjs -c -o $@ [devel-jsassets]
	@cat $^ | uglifyjs -c -o $@ -
#	uglifyjs -c -o $@ $^

src/share/assets/css/index.css: src/share/style/index.scss
	sass $< $@

src/share/templates.html/%.html: src/share/amber.templates/%.amber src/share/amber.templates/defines.amber $(bindir)/amberpp
	$(bindir)/amberpp -defines src/share/amber.templates/defines.amber -output $@ $<
src/share/tmp/jscript.jsx: src/share/amber.templates/jscript.amber src/share/amber.templates/defines.amber $(bindir)/amberpp
	$(bindir)/amberpp -defines src/share/amber.templates/defines.amber -j -output $@ $<

src/share/assets/js/devel/milk/index.js: src/share/coffee/index.coffee
	coffee -p $^ >/dev/null && coffee -o $(@D)/ $^

src/share/assets/js/devel/gen/jscript.js: src/share/tmp/jscript.jsx
	jsx <$^ >/dev/null && jsx <$^ 2>/dev/null >$@

$(bintemplates_productiongo): $(templates_html)
	cd $(<D) && go-bindata -ignore '.*\.go' -pkg view -tags production -o $(@F) $(^F)
$(bintemplates_develgo): $(templates_html)
	cd $(<D) && go-bindata -ignore '.*\.go' -pkg view -tags '!production' -debug -o $(@F) $(^F)
# 	cd $(dir $(word 1, $(templates_html))) && go-bindata -pkg view -tags '!production' -debug -o ../$(bintemplates_develgo) $(notdir $(templates_html))

$(binassets_productiongo):
	go-bindata -ignore '.*\.go' -ignore jsmakerule -pkg assets -o $@ -tags production -prefix src/share/assets -ignore src/share/assets/js/devel/ src/share/assets/...
$(binassets_develgo):
	go-bindata -ignore '.*\.go' -ignore jsmakerule -pkg assets -o $@ -tags '!production' -debug -prefix src/share/assets -ignore src/share/assets/js/production/ src/share/assets/...

$(binassets_productiongo): $(shell find src/share/assets -type f \! -name '*.go' \! -path src/share/assets/js/devel/)
$(binassets_productiongo): src/share/assets/css/index.css
$(binassets_productiongo): src/share/assets/js/production/ugly/index.js

ifeq (, $(findstring bootstrap, $(MAKECMDGOALS)))
$(binassets_develgo): $(shell find src/share/assets -type f \! -name '*.go' \! -path src/share/assets/js/production/)
$(binassets_develgo): src/share/assets/css/index.css
$(binassets_develgo): src/share/assets/js/devel/gen/jscript.js
endif
