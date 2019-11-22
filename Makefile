PHONY: build

build:
	go run utils/parser.go --source https://nvd.nist.gov/static/feeds/xml/sp80053/Rev4/800-53-controls.xml nist-800-53-latest.yaml
	sed 's/^name: NIST-800-53/name: NIST-800-53 rev4/g' nist-800-53-latest.yaml > nist-800-53-rev4.yaml
