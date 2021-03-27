package assets

//go:generate file2byteslice -package=assets -input=./petv1.png -output=./petv1.go -var=PetV1_png
//go:generate gofmt -s -w .
import (
	// Dummy imports for go.mod for some Go files with 'ignore' tags. For example, `go mod tidy` does not
	// recognize Go files with 'ignore' build tag.
	//
	// Note that this affects only importing this package, but not 'file2byteslice' commands in //go:generate.
	_ "github.com/hajimehoshi/file2byteslice"
)
