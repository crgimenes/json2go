json2go
=======

[![GoDoc](https://godoc.org/crg.eti.br/go/json2go?status.svg)](https://godoc.org/crg.eti.br/go/json2go)

Generate Go structs from JSON with struct tags: can specify type name, package name, and additional field tag keys with the CLI app.

The type may be one of the following:

    * struct
    * map[string]T
    * map[string][]T

By default, a struct will be generated.  

If the source JSON is an array of objects, the first element in the array will be used to generate the definition(s).  Any objects within the JSON will result in additional embedded struct types.

The generated Go code will be part of package main unless another package name is set.  Optionally, the import statement for `encoding/json` can be added to the Go source code.

The source JSON can also be written to a provided writer.

Keys with underscores, `_`, are converted to MixedCase.  If any part of a key with underscores matches the list of common initialisms, that element is uppercased, e.g "person_id" becomes "personID".  Keys starting with characters that are invalid for Go variable names have those characters discarded, unless they are a number, `0-9`, which are converted to their word equivalents. All fields are exported and the JSON field tag for the field is generated using the original JSON key value.

If a field's value is null, the field's type will be `interface{}`, as that field's type is not determinable.

There is also a [json2go CLI app](https://github.com/crgimenes/json2go/tree/master/cmd/json2go).  See that [README](https://github.com/crgimenes/json2go/tree/master/cmd/json2go) for more info and examples; including how to install it.

## Examples

__map[string][]T__

Source JSON:

```
{
  "Blackhawks": [
    {
      "name": "Tony Esposito",
      "number": 35,
      "position": "Goal Tender"
    },
    {
      "name": "Stan Mikita",
      "number": 21,
      "position": "Center"
    }
  ]
}
```

Using _Team_ as the parent type name and _Player_ as the struct name:

```
package main

type Team map[string][]Player

type Player struct {
 Name     string `json:"name"`
 Number   int    `json:"number"`
 Position string `json:"position"`
}
```

__struct__

Source JSON from <http://json.org/example.html>:

```
{
 "widget": {
  "debug": "on",
  "window": {
    "title": "Sample Konfabulator Widget",
           "name": "main_window",
           "width": 500,
           "height": 500
   },
   "image": {
           "src": "Images/Sun.png",
           "name": "sun1",
           "hOffset": 250,
           "vOffset": 250,
           "alignment": "center"
   },
   "text": {
    "data": "Click Here",
           "size": 36,
           "style": "bold",
           "name": "text1",
           "hOffset": 250,
           "vOffset": 100,
           "alignment": "center",
           "onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"
   }
 }
}
```

Using _Thing_ as the parent type name results in:

```
package main

type Thing struct {
  Widget `json:"widget"
}
type Widget struct {
 Debug  string `json:"debug"`
 Image  `json:"image"`
 Text   `json:"text"`
 Window `json:"window"`
}

type Image struct {
 Alignment string `json:"alignment"`
 HOffset   int    `json:"hOffset"`
 Name      string `json:"name"`
 Src       string `json:"src"`
 VOffset   int    `json:"vOffset"`
}

type Text struct {
 Alignment string `json:"alignment"`
 Data      string `json:"data"`
 HOffset   int    `json:"hOffset"`
 Name      string `json:"name"`
 OnMouseUp string `json:"onMouseUp"`
 Size      int    `json:"size"`
 Style     string `json:"style"`
 VOffset   int    `json:"vOffset"`
}

type Window struct {
 Height int    `json:"height"`
 Name   string `json:"name"`
 Title  string `json:"title"`
 Width  int    `json:"width"`
}
```

## About this fork

This is a heavily modified fork of the original json2go tool by mohae. https://github.com/mohae/json2go
