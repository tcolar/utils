// History: May 02 14 tcolar Creation

package utils

import "os"

// ByFileInfoModTime is used to sort Files by modification time
type FileModTimeSorter []os.FileInfo

func (v FileModTimeSorter) Len() int      { return len(v) }
func (v FileModTimeSorter) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v FileModTimeSorter) Less(i, j int) bool {
  return v[i].ModTime().Unix() > v[j].ModTime().Unix()
}