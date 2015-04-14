# gohadoopxml: Simple Go Hadoop Configuration Parser #

This simple library allows users to read and extract Hadoop configuration
properties.

## Example Usage ##

The following is an example of how to use this library:

    package main

    import (
        "fmt"
        "github.com/kennyballou/gohadoopxml"
        "log"
        "os"
    )

    func main() {
        configuration, err := gohadoopxml.ParseXML(os.Args[1])
        if err != nil {
            log.Panic("Error occurred while parsing xml")
        }

        for _, p := range configuration.Properties {
            fmt.Printf("%s -> %s\n", p.Name, p.Value)
        }
    }

## License ##

This code is licensed and distributed, as-is without warranty and in the hopes
it will be useful under the terms and conditions of the MIT license. Please see
the LICENSE file for more information. If the LICENSE file was not distributed
with your copy, you may also find it [on the web][1].

[1]: http://opensource.org/licenses/MIT
