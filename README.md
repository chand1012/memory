# Memory

Memory is a Go library for indexing and searching text content. It utilizes the [Bleve](https://github.com/blevesearch/bleve) search library for indexing and searching purposes. The library also provides helper functions for sorting search results by relevance.

## Features

* Indexing and searching text content
* In-memory or on-disk index storage
* Helper functions for sorting search results

## Installation

```bash
go get github.com/chand1012/memory
```

## Usage

### Creating an index

To create a new index, use the `New` function as shown below:

```go
import "github.com/chand1012/memory"

mem, isNew, err := memory.New("path/to/your/index")
if err != nil {
    // handle the error
}
defer mem.Close()

if isNew {
    // a new index has been created
}
```

Replace "path/to/your/index" with an actual path or use ":memory:" to create an in-memory index.

### Adding content to the index

To add content to the index, call the `Add` method:

```go
err := mem.Add("documentID", "The quick brown fox jumps over the lazy dog.")
if err != nil {
	// handle the error
}
```

### Searching the index

To search the index, use the `Search` method:

```go
results, err := mem.Search("quick brown")
if err != nil {
	// handle the error
}

for _, result := range results {
	fmt.Printf("ID: %s, Score: %f, Avg: %f\n", result.ID, result.Score, result.Avg)
}
```

### Sorting search results

To sort search results by score or average score, use the `SortByScore` or `SortByAverage` functions:

```go
sortedByScore := memory.SortByScore(results)
sortedByAverage := memory.SortByAverage(results)
```

## Contributing

We welcome contributions to the Memory library. Please create a pull request or open an issue on the [project's GitHub repository](https://github.com/chand1012/memory) to suggest improvements, report bugs or request features.

## License

This project is licensed under the MIT License.
