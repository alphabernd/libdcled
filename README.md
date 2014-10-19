libdcled
========

GO library for Dream Cheeky USB Message Board

### Usage

```go
package main

import (
  "github.com/alphabernd/libdcled"
  "os"
  "sync"
)

func main() {
  led, err := libdcled.NewDcLed()
  if err != nil {
    os.Exit(1)
  }

  led.ScrollText(libdcled.NewText("Hello World", libdcled.STANDARD_FONT))

  wg := &sync.WaitGroup{}
  wg.Add(1)
  wg.Wait()
}
```
