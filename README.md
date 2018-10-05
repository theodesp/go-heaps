go-heaps
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors)
---
<a href="https://godoc.org/github.com/theodesp/go-heaps">
<img src="https://godoc.org/github.com/theodesp/go-heaps?status.svg" alt="GoDoc">
</a>

<a href="https://opensource.org/licenses/MIT" rel="nofollow">
<img src="https://img.shields.io/github/license/mashape/apistatus.svg" alt="License"/>
</a>

<a href="https://travis-ci.org/theodesp/go-heaps" rel="nofollow">
<img src="https://travis-ci.org/theodesp/go-heaps.svg?branch=master" />
</a>

<a href="https://codecov.io/gh/theodesp/go-heaps">
  <img src="https://codecov.io/gh/theodesp/go-heaps/branch/master/graph/badge.svg" />
</a>

Reference implementations of heap data structures in Go

## Installation
```bash
$ go get -u github.com/theodesp/go-heaps
```

## Contents

**Heaps**

* [Pairing Heap](https://en.wikipedia.org/wiki/Pairing_heap): A pairing heap is a type of heap data structure with relatively simple implementation and excellent practical amortized performance,


## Usage

```go
package main

import (
	"github.com/theodesp/go-heaps"
	pairingHeap "github.com/theodesp/go-heaps/pairing"
	"fmt"
)

func main()  {
	heap := pairingHeap.New()
	heap.Insert(go_heaps.Integer(4))
	heap.Insert(go_heaps.Integer(3))
	heap.Insert(go_heaps.Integer(2))
	heap.Insert(go_heaps.Integer(5))

	fmt.Println(heap.DeleteMin()) // 2
	fmt.Println(heap.DeleteMin()) // 3
	fmt.Println(heap.DeleteMin()) // 4
	fmt.Println(heap.DeleteMin()) // 5
}

```

## Complexity
| Operation     | Pairing       |
| ------------- |:-------------:|
| FindMin       | Θ(1)          |
| DeleteMin     | O(log n)      |
| Insert        | Θ(1)          |
| Find          | O(n)          |
| Delete        | O(log n)          |
| Adjust        | O(log n)          |


## Contributors

Thanks goes to these wonderful people ([emoji key](https://github.com/kentcdodds/all-contributors#emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore -->
| [<img src="https://avatars1.githubusercontent.com/u/1137632?v=4" width="100px;"/><br /><sub><b>Miroojin Bakshi</b></sub>](http://mb-14.github.io)<br />[💻](https://github.com/theodesp/go-heaps/commits?author=mb-14 "Code") |
| :---: |
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/kentcdodds/all-contributors) specification. Contributions of any kind welcome!


## LICENCE
Copyright © 2017 Theo Despoudis MIT license