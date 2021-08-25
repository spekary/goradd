package strings

import "strings"

// ShiftPath removes the first path item from the given path string and returns it
// and the remaining path. Paths are delimited by forward slashes (/).
//
// If the given path is the root path (/), then a "/" is returned and an
// empty string is returned for the remaining path. If the given path has any
// other text in it, that text is returned.
//
// The returned remainingPath will begin with a slash if it has content. Otherwise, it will return
// the empty string. See the tests for examples.
func ShiftPath(path string) (topName, remainingPath string){
	if len(path) == 0 {
		return "",""
	}
	var offset = 0;
	if path[0] == '/' {
		if len(path) == 1 {
			return "/", ""
		}
		offset = 1;
	}
	offset2 := strings.IndexRune(path[offset:], '/')
	if offset2 == -1 {
		return path, ""
	}
	return path[:offset2 + offset], path[offset2 + offset:]

}
