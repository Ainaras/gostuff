# DataLoader Package

A Go package that provides efficient data loading functionality with automatic caching capabilities.

## Public Functions

### `LoadData(file string) string`

Loads data from the specified file path. The function follows this priority:

1. **Cache Check**: First attempts to load from local cache
2. **HTTP Fallback**: If cache miss, fetches from GitHub repository  
3. **Cache Save**: Automatically caches successful HTTP responses
4. **Error Handling**: Returns empty string if both cache and HTTP fail

**Parameters:**
- `file`: Relative path to the data file in the repository (e.g., "f1/races_with_tracks.txt")

**Returns:**
- `string`: File content as string, or empty string on failure

**Example:**
```go
content := dataloader.LoadData("tracks/tracks.yaml")
```

### `ClearCache() error`

Removes all cached files from the local filesystem.

**Returns:**
- `error`: Error if cache clearing fails, nil on success

**Example:**
```go
if err := dataloader.ClearCache(); err != nil {
    log.Printf("Failed to clear cache: %v", err)
}
```
