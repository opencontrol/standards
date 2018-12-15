# Utilities

The parser.go script will sync content in this repo against upstream NIST control
selections.

Depending on the Golang packages installed on your machine, [gopkg.in/yamlv2](https://gopkg.in/yaml.v2) may need to be installed. The ``gopkg.in/yaml.v2`` package implements YAML support for the Go language.
``
$ go get gopkg.in/yaml.v2
``

To run the ``parser.go`` script:
``
$ go run utils/parser.go <output>
``

For example, when NIST SP 800-53 Rev. 5 comes out:
``
$ go run utils/parser.go nist-800-53-rev5.yaml
``
