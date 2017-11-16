# Utilities

The parser.go script will sync content in this repo against upstream NIST control
selections.
``
$ go run utils/parser.go <output>
``

For example, when NIST 800-53 rev5 comes out:
``
$ go run utils/parser.go nist-800-53-rev5.yaml
``
