![go-check](https://github.frapsoft.com/top/go.png?v=001)

# go-check

_Easy error check with colorful output_

![go-check](https://github.frapsoft.com/top/go-check.png)

## Install

`go get github.com/ellerbrock/go-check`

## Example

- `check.Error(err)` exit the program on error with colorful output

```
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ellerbrock/go-check"
)

func main() {
	content, err := ioutil.ReadFile("x.txt")
	check.Error(err)
	
	// we will only get here if the file exists ...
	fmt.Printf("file content:\n\n%v", content)
}
```

##  Contact

[![Github](https://github.frapsoft.com/social/github.png)](https://github.com/ellerbrock/)[![Docker](https://github.frapsoft.com/social/docker.png)](https://hub.docker.com/u/ellerbrock/)[![npm](https://github.frapsoft.com/social/npm.png)](https://www.npmjs.com/~ellerbrock)[![Twitter](https://github.frapsoft.com/social/twitter.png)](https://twitter.com/frapsoft/)[![Facebook](https://github.frapsoft.com/social/facebook.png)](https://www.facebook.com/frapsoft/)[![Google+](https://github.frapsoft.com/social/google-plus.png)](https://plus.google.com/116540931335841862774)[![Gitter](https://github.frapsoft.com/social/gitter.png)](https://gitter.im/frapsoft/frapsoft/)