package export

import "fmt"

var Exported = "I'm going to be available outside this package"
var notExported = "I will only be available within this package"

func ExportedFunction() {
	fmt.Println(Exported)
}

func notExportedFunction() {
	fmt.Println(notExported)
}
