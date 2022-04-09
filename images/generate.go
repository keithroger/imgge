// Run all effects and create create output images to arugment location
package main

import (
    "flag"

    "github.com/keithroger/imgge"
)

var (
    effectName = flag.String("name", "all", "name of effect")

    examples = struct {
        name string
        f func()
        args []interface{}
    }{

    }
)

func main() {

}
