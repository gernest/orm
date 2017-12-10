
package b0x


import (
  "bytes"
  "compress/gzip"
  "io"
  "log"
  "net/http"
  "os"
  "path"

  "golang.org/x/net/webdav"
  "golang.org/x/net/context"


)

var ( 
  // CTX is a context for webdav vfs
  CTX = context.Background()

  
  // FS is a virtual memory file system
  FS = webdav.NewMemFS()
  

  // Handler is used to server files through a http handler
  Handler *webdav.Handler

  // HTTP is the http file system
  HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct {}



// File00ImportGoTmpl is "00_import.go.tmpl"
var File00ImportGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x7c\x8e\xc1\xca\xc2\x30\x0c\x80\xcf\xeb\x53\x84\xb2\xc3\xff\x83\xae\x77\xc1\x07\xf0\xee\x5d\xba\x36\xce\xc2\xd2\xcc\xb4\x8a\x52\xfa\xee\xb2\xb9\x93\xa0\xc7\x2f\xdf\x97\x90\x40\x13\x4b\x86\x3f\xd5\x68\xc7\x31\xe3\x23\x6b\xd5\x68\x6f\xb3\xed\x6d\x42\x93\xae\xe3\x27\x1b\x2f\xe1\x8e\x32\x8f\xcf\xb4\xd4\x43\xc8\x97\x5b\xdf\x39\x26\x33\x71\xc2\x88\x62\x58\xe8\xbb\x31\x8e\x89\x38\xfe\x08\x7c\xb0\x23\xba\xf9\x78\x29\x20\x36\x0e\x08\xed\x69\x03\xed\xfa\xed\x6e\x0f\xdd\xf1\x39\x61\x77\x58\x38\xc1\xb6\x56\xd5\xe8\x52\xd6\xa0\xd6\xf7\x26\x46\xbf\xa8\x7f\xf5\x0a\x00\x00\xff\xff\xbb\xf7\x97\x25\xe7\x00\x00\x00")

// File10OrmGoTmpl is "10_orm.go.tmpl"
var File10OrmGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x59\x5f\x6f\xdb\x48\x0e\x7f\x8e\x3f\x05\x2f\x08\x02\xfb\xe0\xaa\xef\x06\xfc\x90\x26\x2e\x2e\xb8\x5c\x9a\x8b\x53\x74\x81\xa2\x58\x8c\x25\xca\x99\x5d\x79\xe4\xce\x8c\xf3\x07\x82\xbe\xfb\x62\xc8\x91\x2c\xc9\x92\xac\xa6\xc5\x3e\x6d\x5f\x6a\xcd\x50\x1c\xf2\x47\xf2\x47\x8e\xf2\xfe\x3d\x64\xd9\x59\x70\xa7\xe5\x93\xb0\x98\xe7\x0f\x62\x95\x20\x48\x03\xcb\xff\xdf\x80\xa5\x07\x25\x36\x38\x0a\x53\x65\x6c\x9b\xe8\x1c\x4e\xb3\x2c\x78\x78\xdd\x62\x40\x0b\x79\x7e\x3a\x1a\x35\xb5\x5e\x6a\x14\x16\x2f\xd3\x64\xb7\x51\x66\x69\x85\xc5\x0d\x2a\x6b\x40\x68\x84\x90\x57\x21\xc2\x58\x2a\x69\x65\xaa\x0c\x48\x05\x91\x8c\x63\xd4\xa8\x2c\x44\x52\x24\x18\x5a\x33\x7a\x12\x7a\x98\xda\x39\x6c\xc4\xf6\xab\xb1\x5a\xaa\xf5\x37\xfe\x2f\x1b\x01\x00\x64\x19\x68\xa1\xd6\x08\x67\xbf\x4f\xe1\x2c\x82\xd9\x1c\x82\x2b\xaf\x1e\xde\xe5\x39\x09\x9d\x66\xd9\x59\x14\xdc\x8a\x8d\xf3\x65\xe6\x1f\x9b\x87\xc0\x19\xf9\x9c\xe7\xa7\xd3\x42\x33\xaa\x88\x74\xe4\xa5\xff\x84\x0a\x2b\xfa\x74\xff\x3f\x87\xaa\x7d\x44\x90\xca\xa2\x8e\x45\x88\x90\xc6\xb4\xe0\xf6\xd2\xd5\x1f\x18\xda\x91\x7d\xdd\x62\xeb\xab\xe5\x3b\xec\xc8\x65\x92\x1a\x1c\x4f\x00\xb5\x4e\x35\xaf\x10\x16\xe3\x09\xfc\x9b\x20\xda\xad\x12\x19\x16\x08\x7d\xd8\xc9\x24\x42\x96\x5b\xa2\xf3\xb6\x29\xc7\xab\x55\xb9\x6b\x65\x50\x1f\xc8\xf1\x6a\x55\xee\xf3\x36\x6a\x39\x97\x57\xab\x72\x57\x98\xe0\xa1\x1c\xaf\x56\xe5\xbe\x3c\xa2\x3e\x10\xa3\xc5\xaa\xd4\x4d\xba\x5e\xa3\x1e\xa7\x7a\x13\xf0\xcf\x89\x87\xfd\xd3\x16\x55\x0b\x7e\xe9\x16\x5d\x8a\x09\x2b\x56\xc2\xb8\x9c\x53\x0a\x43\x97\x6b\xa3\x78\xa7\xc2\x8e\xb7\xc6\x91\x96\x4f\xa8\xdd\xf3\x94\xde\x5d\xa6\x3b\x1d\xa2\x7b\x06\xce\xaa\x09\x8c\x0f\x5f\x9b\x72\x54\x26\x90\x8d\x4e\xa2\x15\x3d\xb9\x44\x33\xdf\x93\xc0\x9d\xd3\xa3\x75\x32\x3a\x91\x31\xc9\xff\x6b\x0e\x4a\x26\x4e\xc3\x89\x46\xbb\xd3\xca\x3d\x92\xaa\xd1\x49\x3e\x2a\xd6\x6e\xf1\xf9\x98\xd5\xab\x02\x98\x56\x59\x60\x45\x06\x84\x22\x4c\x7c\x1a\x42\xac\xd3\x0d\x08\x88\x56\x20\x95\xb1\x42\x85\xc8\x38\x1d\x3b\xd0\xc3\xe2\xce\x05\x17\x9b\xab\x0f\xc7\x11\x2a\x01\xf2\x75\x1e\xdc\xe2\x73\x45\xe5\x50\x4c\x5c\x56\xf8\xe5\xf3\x2c\xab\x70\x44\xaa\x54\xe6\x55\xcf\x20\x72\xa6\xcd\x20\x5a\xe5\x53\xf7\x76\x59\xab\x75\x79\xd0\xb8\xd5\x68\x98\xa2\xe0\xea\x43\x25\x5f\x20\x4e\x35\x6c\x84\x92\xdb\x5d\x22\xac\x54\x6b\x10\xb0\x96\x4f\xa8\x9c\xeb\xbb\xd0\x06\x4e\xdf\x45\x92\x80\xc3\x8b\xc9\x4c\x3c\x09\x99\x10\x57\xda\x94\x0b\x59\x84\x16\x9e\xa5\x7d\x74\xb0\xef\x79\xd6\x3e\x0a\xeb\x38\x42\x63\x22\x2c\x46\x4e\x91\x4d\xc1\x3e\x4a\xe3\x75\x4f\x89\x2f\xa3\x54\x21\xac\x5e\xdd\xbb\x45\x70\x98\x48\xa4\xf1\xe1\x23\x1b\x1e\x52\x58\xa3\x6d\x4a\xa5\x7a\x03\x3b\x83\x94\xf0\x90\x6a\x17\xd0\xbd\xa5\x41\x41\x3f\x0d\x34\xf8\x74\x0a\x16\xe3\x58\x86\xca\x73\xa7\xcb\x73\xa0\x7f\x1c\xf4\xd1\x49\x42\x45\xc9\xcf\x5c\xa0\x0e\x6a\xca\xa1\x71\xe8\x6a\xbb\x7e\xc4\xa4\x4e\x68\xee\x28\x1f\xcb\x30\x88\x56\x81\xdf\xf4\xd1\x62\x7d\x60\x90\xa2\xe3\x4f\x22\xa4\xb8\xb0\x61\x2b\xc2\x3f\xc5\x1a\xfb\x8e\xf3\xfc\xe1\x5f\xae\xd0\x88\x3b\x3a\x0c\xfc\xfa\xdc\x6b\xf7\x07\x33\x99\xee\x6b\x06\x56\xcc\x48\x0e\x58\x1f\xc9\xcb\xfb\xc5\xc5\xc3\x02\x4c\xd1\x26\x7a\x5d\x3e\xce\xd8\x15\x20\xce\x3b\x85\x5c\x41\x6c\x85\x16\x1b\x33\x83\x30\xdd\x6c\x52\x15\xf0\xfe\x1d\x2d\xba\x6d\x17\x1a\xea\xce\xb3\x96\x16\x3e\xf5\x02\xcd\x16\x37\x1b\xd4\x6c\xbf\x86\x41\x59\xb8\x62\x83\xe3\xc9\x37\x6e\x88\xee\x5f\x3e\x1d\x91\x66\x17\x95\x19\x84\xbc\x51\x34\x48\x6e\x39\xbd\x68\x2e\x17\x37\x8b\xcb\x87\x61\x68\x0e\xe8\x6b\x0e\x4d\xe3\x98\xe6\xbc\x53\xa4\x05\x4b\xde\xf7\x58\x76\x82\xe8\x5c\x3d\x29\xfd\xf4\x7c\x64\x02\xd6\x55\x4c\x0f\x30\x87\x73\x13\x18\x52\xe8\xdb\xb6\x8f\xae\xf1\xa0\x70\x7f\xed\x05\xe5\xfa\x76\xb9\xb8\x1f\x08\xca\x80\x26\xde\x95\x62\x35\xa1\x16\x58\x78\xff\x87\x61\x61\x37\x79\x3c\xe8\x75\xf3\xf3\xdd\xd5\xe0\x4a\x1a\x30\x83\x74\xb9\x59\x13\x6a\x71\x93\xf7\x7f\xcc\xcd\x6a\x96\xf3\x80\xd3\xeb\xe9\xd5\xe2\x66\x31\xd4\xd3\x01\x53\x54\x97\xa7\x35\xa1\x16\x4f\x79\xff\xed\x9e\xd2\x8c\xd6\xeb\xe8\x97\xff\x2c\xee\x07\xfa\x79\x74\x0a\xec\xf2\xb2\x2a\x93\x15\xa6\x25\xe9\x1a\x64\xbc\x6f\x16\xcf\xc2\xb8\xf6\xd1\x67\x40\x92\xae\xc7\xa6\x1c\x67\x84\x5e\x1b\x08\x82\xa0\x1c\xc3\xb3\x9c\x3a\x85\x8c\x61\xdf\x2c\x9a\xe3\x09\x4d\x6a\xc5\xf6\xd8\xb0\x96\x20\x08\x8a\x3e\x36\x9f\xcf\xa1\xce\xf6\xf3\xf9\x7c\x7f\x77\x6a\xed\x07\x04\xaa\xe9\xea\x36\x40\x21\x45\x8b\xda\xec\x6f\x12\xed\x8a\xf6\x3d\x9d\xd3\xa0\xad\x73\x8c\x28\xd0\x00\x87\xf8\x14\x6c\x15\xdf\xa6\x76\xf1\x22\x8d\x35\xdc\x8e\xaf\x3f\xc2\xed\xa7\x07\x58\xfc\x76\xbd\x7c\x58\xd2\xa8\xe4\x9a\xb2\x37\xd3\x59\xdc\x0c\xfe\xaa\xa7\xf9\x4d\xaa\xfa\x8f\xb6\xc9\x55\xc1\xb5\x55\xa3\xe6\x60\xf5\x0e\xcb\x4c\x59\x15\x8d\x3c\x55\x16\x5f\x2c\xdb\xec\xc7\x06\x5a\x28\x2c\x76\xa6\x7e\xdf\xa1\x7e\x1d\x64\xa6\x57\x37\x0e\xed\x4b\xa1\x29\xf0\x6b\x83\xad\xbe\xb4\x2f\x30\x87\xd0\xbe\x1c\x18\xeb\xd2\xa4\xce\xd8\x6d\x69\x52\x97\xd8\xa7\x49\xb3\x5b\xf4\xa6\x48\x5d\x49\x57\x8a\x54\x99\xff\x68\x8a\xfc\x2c\xd4\x35\x93\x7e\x04\xea\x83\x26\x37\x00\x6a\x7e\xa7\xf8\x9c\xc1\xd7\x95\xea\x15\xa9\xad\xfd\x96\x3c\x17\xcb\x24\xc1\x88\xe7\xfa\x27\x91\xec\xd0\x10\xf5\xf9\xdb\x81\xbf\xdf\x0f\xf1\xb1\xcd\x8a\xf1\x96\x10\xa6\x95\xc5\x8b\xa5\xbb\x96\xbf\x55\xdd\xf1\xb4\x9b\xe7\xc7\x00\xa8\x7f\xfc\x88\xe9\xe3\x07\x69\xf8\x28\xd1\x65\xcb\xbb\x3c\x27\x21\x19\xc3\x59\x1c\x5c\x9b\x25\x5a\xbe\x98\x54\x36\x54\x6a\x79\xf3\x1e\xe9\xeb\x4c\xe8\x77\x4b\x70\x2f\x8c\x91\x6b\x45\x93\x61\x70\x11\x45\xe3\xd3\x2c\x3b\x8b\xfd\xfc\x93\xe7\xa7\x53\xd8\x06\xb4\xc2\x6e\x4d\x48\x2f\x26\x06\x1b\xa7\xb3\x67\x29\xf1\x2d\x6f\xc9\xb8\xfe\x6a\xe5\x3a\x58\xf9\xf6\xf2\x56\x43\xf8\x37\x9f\xaa\xe5\x46\xe8\xd7\xff\xe2\x6b\xcd\xc8\x2e\xbb\xf2\xfa\xf9\x43\x7e\xb7\x14\x78\x7d\x56\x69\x2b\xf0\xba\x84\x2f\xf0\xd6\x84\xec\xab\xf0\xba\x96\xae\x0a\xaf\x0e\x3d\x47\x2b\x9c\xfb\x7e\x59\xdf\x8d\x26\x5f\x5c\xcc\x8e\x94\x79\xcd\xae\xa2\xfd\x3f\x93\x66\x6f\x13\x2d\x1d\x9b\xf0\xca\xe8\xb3\x51\x73\x20\x15\xbf\x9c\xff\x1b\xd6\x0e\x27\xa5\x6e\x83\xbb\x49\x89\xdf\x69\x90\xd2\x8e\x07\x68\xcf\x33\xce\x6c\x91\x24\x45\x3c\x63\xaa\xe7\x41\xa6\xb7\x29\x7f\x03\xd7\x1c\xf8\x35\x84\x6b\xfc\x67\xd3\x7f\xe8\xe6\xef\xa7\x9b\xfa\x85\xa1\x8d\x6e\xea\x12\x15\xba\x69\xde\x56\x7a\xe9\xa6\xae\xa5\x8b\x6e\xaa\x37\x8f\x81\x74\x23\xb6\xdb\x44\xa2\x81\x82\x22\x54\xe4\xff\x72\x90\xaa\x21\x05\x5c\xb3\xab\xa4\x9b\x3e\xaa\x39\xb8\x62\x1d\x52\xcd\x2f\xa7\x99\x86\x95\xc3\x69\xa6\xdb\xd8\xfe\x31\x93\xcb\xf8\xbd\xff\x08\xc1\x34\xc2\xd9\x91\x65\xef\xfa\x0b\x3a\xcf\x9d\x50\x7b\x31\xd3\xc7\x1e\x5b\x2d\x26\x42\x84\xc8\x8b\xb0\xe0\x3f\x01\x41\xad\x64\x40\x72\x24\x3b\x3e\x77\xf4\x8f\x50\xf5\xe3\xc6\x7c\x12\xad\x2c\xd1\xfe\xd4\x10\x35\xbc\xe0\xe9\xd0\xc9\x01\xce\x6f\x47\xa2\xe3\x8b\x48\x3f\xc1\xff\x32\x24\xba\x5b\xd7\x9b\x90\xa8\x10\x54\xe5\xe7\x5f\x01\x00\x00\xff\xff\x34\x75\xa1\xff\x8e\x1c\x00\x00")

// File30ExecGoTmpl is "30_exec.go.tmpl"
var File30ExecGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x58\x5d\x6f\xdb\x36\x17\xbe\x96\x7e\xc5\x69\xe0\xb6\x52\xe0\xd2\xed\x6d\x0a\x5d\xbc\xcd\x9b\x62\xc5\xb6\xb6\x6b\xb3\xdd\x04\xc1\x40\x53\x47\x36\x11\x8a\x74\x49\x2a\xb1\x21\xe8\xbf\x0f\xfc\x90\xa3\xfa\x23\xf1\xb0\x74\xd9\x45\x72\x13\x89\x87\x3c\x1f\xcf\x79\xf8\x1c\x25\x93\x09\x9c\x2d\x91\x01\xd3\x48\x2d\x1a\xa0\x60\xe9\x54\x20\x54\x4a\x83\x9d\x23\xcc\xf8\x35\x4a\x30\x56\x37\xcc\xa6\x55\x23\x19\x64\x53\x38\x6e\xdb\x11\xf9\xdc\x4c\x05\x67\x5d\x77\xea\x4f\xbe\x6b\xb8\x28\x51\xe7\xde\x5b\x96\x43\x66\xbe\x09\xf2\x05\x4d\x23\xec\x18\x50\x6b\xa5\x73\x68\xd3\xc4\xd8\xda\x8e\x81\xea\x99\x81\x93\x02\xa6\x84\x29\x29\x49\xc9\xa9\x40\x66\x49\xf0\x94\xbd\x98\x92\x05\xd5\xb4\x36\x79\x9a\xc4\x1d\x42\xcd\xb2\xa3\x60\x3e\x81\x97\xcf\xaf\x5f\xc2\xf3\xeb\xa3\x31\xdc\x7a\xcb\xd3\x44\xa3\x6d\xb4\x5c\xfb\x9c\x12\x97\xc9\xa9\x92\x16\x97\x36\x63\xaa\xae\x95\x24\xf1\xf5\x93\x7e\x47\xd9\xd5\x4c\xab\x46\x96\x59\x1f\x8d\x9c\xda\x65\x3e\xf4\x49\x08\xc9\xd3\x2e\x4d\x27\x13\xf8\xd6\xa0\x5e\x01\x37\xd0\x18\x2c\x61\xba\xf2\xd0\x7c\x45\x9f\xf5\x6f\xde\x46\x65\xd9\x2f\xfc\xc2\x6b\x6e\xc1\x61\x65\xb9\x92\x66\x37\x6a\x61\xef\x1a\x35\x1f\x20\x63\x76\x09\x2c\xa4\xd8\xa7\x9a\x43\x76\xec\xa1\x54\x37\xe6\x60\x20\x83\xf3\xbd\x40\xfa\x8c\x0f\xc7\xd1\x6f\x5f\x03\x69\x97\x7b\x30\xf2\x34\xe2\xd2\xa0\xb6\xc6\xe3\x53\x52\x4b\xc1\xaa\x01\x8d\xdc\xca\x94\x1a\xdc\x0d\xc9\x07\x7f\xf6\x40\x22\xf1\x0a\x04\xca\xdb\xe6\xfd\xcf\x18\x3e\x93\x35\x4a\x6b\x72\x28\x0a\x78\xed\x36\xf5\xb5\x48\x2e\xc6\x50\xd5\x96\x9c\xb9\xf3\x55\x76\x24\x95\x9d\x73\x39\x73\xd9\x85\x8c\x8f\xf2\x34\xe9\xee\x41\x35\xe4\xb7\x17\xd5\x60\x7e\x14\x7a\xfe\x63\xe8\x7f\x5f\x94\x87\xdf\xe1\x07\x83\xbe\xf1\x51\x0f\x81\x3e\xe4\xb7\x17\xfa\x60\x7e\x3c\xe8\x75\x23\x23\xee\x28\xd0\x22\x18\x4b\x2d\x3a\x40\x40\x49\xa0\x1b\x1d\x20\xbb\x5b\xf0\x7f\x7f\xf4\x21\x64\x34\x78\xda\x0b\x56\x30\x3f\x0a\x58\x41\x2a\x7b\x82\xee\xa7\xe3\x86\x38\xfa\x63\x0e\x8c\x8b\x4b\xb7\xeb\x7c\xb5\x40\x72\xb6\xb4\x1f\x69\x8d\x10\x5f\x3f\x53\x76\x45\x67\xd8\x75\x03\x94\x00\x00\x9c\xa2\x9e\x14\x70\x68\xde\xfe\x8c\xee\xa5\x36\x60\xbb\x96\xe6\xdc\x53\xdf\xad\x3f\x2b\x1c\xaf\xb7\x68\x8e\x5a\x7b\x2a\x97\x58\xa1\xf6\x6e\xc8\xa9\x50\x06\xb3\x3c\x4d\x93\x6b\xaa\x21\x4b\x13\x17\x80\x5b\xac\x0d\x1c\x50\x8c\x4f\xc7\xfd\xb4\x2d\xf0\x0a\x82\xf5\x27\x6a\x3e\x49\x3c\x57\xbf\x52\xb9\xfa\x82\x82\xba\x21\x03\xaf\x06\x9b\x27\x13\xc0\x25\x37\xd6\xb8\x81\x45\xa1\xa6\x8b\x85\xbb\x70\x95\x56\x35\x2c\x34\xaf\xa9\x5e\xc1\x15\xae\xdc\x0d\xa4\x42\x23\x2d\x57\xb0\xa0\xda\x0d\xb6\x30\xe6\xcd\xda\x53\x74\x53\x40\x4d\xaf\x30\xab\xe9\xe2\xa2\x6d\x63\x8a\xc1\xcf\xcf\xb8\xba\xbb\x82\xcb\xe3\xfe\xc4\xbe\x1d\xf9\xb0\x48\x94\xe5\xba\x94\x3c\x4d\xdc\x47\x88\x87\xf1\xa3\x23\x9f\x67\x7e\x2c\x90\xcd\x91\x5d\xf5\xb3\x12\x18\x95\x0c\x45\x80\x22\x42\x5c\xf5\x0d\x64\x76\xe9\xb4\x27\xcb\xdf\x0e\x5b\xd7\xbb\xf2\xfd\xde\x6c\xa1\x5b\xec\xd2\x24\x71\x7d\x1a\x10\xc1\x78\x56\x2a\x4d\xbe\x32\xea\xe4\xef\xbb\x4b\xe7\x6a\xcb\xf2\x71\xcf\xb4\x3f\xa8\x68\xd0\x64\xc7\x2e\xfb\xbc\x6d\xef\x6c\x9e\x27\xad\x47\xba\x6d\x51\x96\x43\x48\xb6\xf9\xb6\x4d\xb8\xa4\x4b\xff\x1e\x4f\x12\x4f\x63\x17\xd0\xd5\x15\x22\x5f\xb8\x5a\xc9\x8e\xee\xba\xba\xba\xee\xf2\x6d\x3c\x30\xc8\x24\xc6\xd3\x54\xce\x10\x46\x7f\x8e\x61\x54\x39\x7f\xc1\xc3\x17\x77\x05\x50\x32\x34\x31\xe4\x6d\x76\xa3\x2a\x6c\xf9\x2a\x38\xc3\x68\x4d\xbc\x77\xd2\xb6\xa3\x2a\x46\x84\x02\xe8\x62\x81\xb2\xcc\xb6\x4c\x63\x88\xc9\xae\x57\xbc\xc4\x24\x49\x32\x60\xd0\xe6\x5b\x07\x28\x0c\x06\x08\xc3\xfd\x5b\x07\xf0\xaf\x63\x38\x76\xbf\xf3\x75\x32\xf7\x42\x02\x05\xbc\xf0\x47\x2f\xdc\x34\xf4\x4f\xf9\xab\x37\x97\xbe\x21\x21\xb8\x0b\x18\x01\xbf\x3b\xe2\x30\xd3\x6e\x2d\xc1\x71\x97\xbf\x00\x9e\xc1\x51\x45\x4f\x55\x23\x2d\xd0\xb2\x04\x0a\xcc\x3f\x33\x25\x9a\x5a\xf6\x03\xdf\x0b\xd6\x41\xc2\xea\x3d\xf5\xc2\x1a\xea\x0c\xc5\x79\xc3\x03\x48\xe9\xe0\xd2\x84\x4c\x0b\xb0\xba\xc1\x7f\x5b\x66\xb7\x4a\x7b\x12\xd6\xff\xac\xb0\x06\x4e\x3e\xa9\xeb\x93\xba\xee\x52\x57\xb2\xfe\x66\x0a\xfb\x7e\xbc\xda\xbe\xe7\xda\xd8\x48\xe5\xf0\x91\x5f\x85\x15\x75\x03\x76\x4e\x2d\xd4\xd4\xb2\x39\x9a\x5b\xe9\x25\xee\xd8\x87\x0a\xa4\xf2\x9b\xb6\xec\x63\xa0\x12\xce\xb4\xfe\xa8\xec\x7b\xa7\x9b\x70\xc3\x85\x80\x29\xc6\x20\x58\x7a\x07\xe7\x73\x6e\x80\x51\x21\xe2\x25\x34\x40\xa5\xd3\x93\x99\xff\xe3\xc9\x05\xbe\xa1\x06\x0c\x5a\xb8\xe1\x76\xee\xbc\xbb\x53\x7b\xc5\x1e\x16\x1a\xaf\xb9\x6a\x8c\x58\x91\x83\x66\x83\xaf\xdb\xcd\x86\x7b\x85\xe6\x41\xe6\x44\x5c\xf9\x4c\x67\x18\xff\x81\x52\xc0\x9b\x1d\xb6\x4f\x55\xe5\x6a\x2e\xe0\xf5\x0f\x1c\x22\x95\x6f\xcb\x49\x31\xd4\xc8\x34\x2a\xc6\xb3\x60\x6c\xd3\x5d\x2a\xa7\x74\x4d\x06\x9d\x4d\x83\xe0\xf5\xc3\xe8\xc1\xbe\x25\x47\x77\xcb\x9d\xe4\xe2\x3b\xad\xdb\x44\x62\x9f\x3e\x47\x79\x1e\x5c\x88\x8d\xfb\xf0\x57\x00\x00\x00\xff\xff\x60\x97\xac\x6b\x2d\x14\x00\x00")

// File40SelectGoTmpl is "40_select.go.tmpl"
var File40SelectGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xdc\x57\x4d\x6f\xe3\x36\x10\x3d\xc7\xbf\x62\xb0\x30\x0a\x29\x70\x95\x4b\xd1\xc3\x02\x3e\x74\xdd\xf4\x33\xdb\xb8\x71\xd0\x1e\x16\x8b\x82\x12\x47\x36\x37\x12\xa9\x92\x54\x62\x43\xd0\x7f\x2f\x38\x22\xbd\x92\xed\x64\xe5\x14\xbd\xf4\x14\x8b\x1c\x3e\xbe\x37\x7c\x33\x64\xec\xae\x42\x68\x9a\x69\xb2\xac\xd3\x42\x64\x6d\xbb\xca\x98\x94\xa8\x41\x48\x8b\x3a\x67\x19\x42\x33\x01\x00\x58\x32\xcd\x4a\x13\xc5\x90\xa9\xb2\x54\x32\x59\x61\x81\x99\xed\x46\x29\xc0\x2d\x8c\xb8\x60\x6e\x18\x8c\xd5\x42\xae\x67\xf0\xc8\x8a\x1a\x0d\x7c\xf8\xc8\xb5\x78\x44\x9d\xfc\xe1\xbe\x9b\x46\xe4\x90\xdc\xef\x2a\x4c\x7e\x62\xe6\x56\xe2\xbd\x7a\xcf\xe4\xee\x0e\x0b\x66\x85\x92\x6d\x3b\x03\xdc\x0a\x63\x0d\x94\xac\xfa\xd0\x34\x5d\xe8\x52\x8b\x92\xe9\xdd\xaf\xb8\xeb\xbe\xaf\xb7\xf6\x37\x56\x22\x4c\xfd\x34\xcb\x1e\xd8\x1a\xdb\xf6\xe3\x65\x58\xf1\x5c\x44\xd3\xa0\xe4\x6d\x1b\x43\x74\xe9\x94\x7f\x8e\x7d\x40\x7e\x14\x3c\x03\xd4\x5a\xe9\x78\xd2\x4e\x26\x57\x57\x10\xb0\x1d\x70\xdb\x2e\x54\x2d\x2d\x08\x03\xcc\x29\xae\x33\x0b\xb9\xd2\x90\xb9\x51\x21\xd7\xa0\xd5\x93\x01\x95\x83\x4f\x72\x7f\xe5\xe4\xc4\x58\x87\xe6\x81\xba\xac\x7f\x49\x0a\x05\x79\x16\xd2\x7e\xfb\xcd\x9e\x65\xef\x40\xe9\xa0\xde\xd5\xa2\xe0\xa8\x21\x75\x7f\x0d\x30\x09\xab\xdf\x6f\x60\x75\x7d\x73\xbd\xb8\x07\x63\x99\xc5\x12\xa5\x85\xca\x9d\x27\x5a\xd4\x66\x72\xec\x8c\x01\xd0\x9e\xe6\x05\xad\x31\x27\x7d\x71\x91\x29\x29\xc1\x1d\xc8\x52\x8b\x47\x66\x49\xa4\x94\x93\x0b\x43\x51\x4a\x77\x1b\x84\xb9\x95\x1f\xed\xab\xd8\xcf\x79\x5f\x6a\xac\x34\x1a\x94\xd6\x25\x3d\xf5\x64\xec\x86\x59\xc0\x6d\xa5\x0c\x1a\x50\xb2\xd8\x79\xbb\x02\x93\x9c\x8c\x09\x25\xda\x8d\xe2\x3d\x51\x87\xb8\x83\xb4\x07\xdc\xcb\x67\xf5\x3b\x8a\x79\x2d\x33\x88\x8c\x8f\x3a\x00\x8c\x5f\xac\x18\xbf\x8f\x46\x5b\x6b\x09\x26\xf1\x1b\x26\x5d\x2e\x47\x80\xff\xcf\xab\x6d\x98\x9f\x5e\x82\x82\x71\x92\x7e\x02\x82\xf2\xd1\x52\x03\x27\x6f\xb4\xcf\xde\x72\xdb\x51\x75\xa8\xf4\x93\xcb\xac\x55\x90\x22\xd4\x06\x39\x08\x09\x0c\x3e\x29\x21\x41\x55\xa8\x09\x10\x9e\x84\xdd\xc0\x5e\x95\xaf\xec\xee\xe4\xd2\x17\xcc\x13\x87\x1d\xa3\xf8\x64\xef\x1d\x88\xff\xea\x94\x01\x1a\x9f\x8f\xb7\x90\xb6\x5e\xc4\x9f\x1b\xd4\x08\xac\xaa\x0a\x81\x06\x9e\xe8\x2b\x53\x92\x0b\xc7\xd4\x55\x05\xd8\x0d\xc2\xdf\x35\xea\xdd\x28\x8a\x84\x17\x05\x1c\x72\x30\x0d\xc5\x2f\xac\x72\xed\x20\xf5\x26\xee\xa2\x61\xde\x51\x99\x5c\x78\x39\xa9\xa7\x7b\x23\x4a\x61\xf7\x74\xa9\x53\x16\x34\xd4\x27\x0a\x1a\x4d\xa5\xa4\xc1\x51\x8c\x09\x32\xea\x50\xa8\x1b\x8e\xa6\xba\x64\x6b\x4c\x3a\x46\xf3\x8e\xc6\x11\x5f\x17\x32\xa4\xab\xf2\xdc\xa0\xa5\x16\xf3\x6f\x99\x3b\xf0\xa8\xc3\x9b\xc1\xab\x05\xdc\x76\x84\xe6\x9e\xd9\x19\xfa\x9a\x06\x34\x93\x6b\x84\xe9\x5f\x33\x98\xe6\xf0\x76\xee\xcb\xe8\x07\x81\xee\xba\xf8\xba\x6d\x5d\x8c\xc8\x41\x2a\x0b\xd3\x3c\xf9\xd9\xdc\x61\x8e\x1a\x65\x86\xd0\xb6\x54\x43\xc4\xad\x69\xa6\xb9\xaf\x03\x60\x9c\x1b\xe8\x0f\x58\x45\xf9\xe9\x4a\x18\x39\x64\xaa\xa8\x4b\xe9\x2e\x48\x76\x86\x2f\x8f\x36\x8a\xbe\x90\x25\x6a\xe8\xbd\xc6\x71\x44\x74\x0e\x56\xd7\xd8\xaf\x39\x97\x94\xa6\x01\x2c\x0c\xc9\x3b\xba\x8c\x68\x75\xaf\xe6\xf7\x8f\x26\x7a\x09\xf8\x0f\xba\x94\x38\xe6\x42\x22\x87\x74\x07\x7e\xe7\x28\x4e\xf6\xd5\xef\x80\x9d\x7c\x09\xb7\x77\xef\x43\xcf\x71\x4f\x88\x70\x4f\xe5\xa7\xde\x0c\x63\x88\xfc\xc7\xaf\xb7\xb0\xeb\xa8\x1b\x25\x04\x9f\x73\xa7\x84\x35\xe3\x6e\x95\xfc\xbc\x47\xdc\x2f\x4a\xc8\x03\xa7\x86\xde\xde\x15\x6f\xde\xbd\x4c\xf2\x73\x5a\xfa\x01\x68\x14\x6c\x30\xe6\xbc\xce\x73\xf0\x21\xfd\x79\xb0\xdc\x29\x0b\x4b\x4e\x0e\x1e\x54\x2f\x11\x58\x15\x22\x43\xaa\xec\xab\x2b\xb8\xd5\x1c\xf5\xbb\x5d\x1f\xd5\xb5\x12\xa5\xe9\x75\xa5\x86\x7d\xad\x2e\xdc\xe3\x2b\xcb\x94\xe6\xee\x95\x6b\x55\xa8\x64\x5a\xbe\xa0\xdf\x23\x93\x76\xbc\x6f\xc4\x85\x0e\x2e\xa5\xd9\xef\xc5\xb8\xf4\xf8\x56\x47\x6b\x4c\xf2\x1d\xe7\xd1\x9b\x01\x9f\x37\x33\xe0\x42\xc7\x87\x39\x72\xea\x7f\xd4\xaa\xae\x86\xea\x4b\xf6\x80\x3d\xd5\x6b\x17\xe1\x8a\xf8\xd5\x4a\x8f\xf7\x88\xce\x91\x45\xcb\x4f\xca\x8a\x9f\x39\x75\xdf\xb3\xc3\x4f\xa7\x73\xa1\xa4\xc5\xad\x75\x47\x6b\x48\x5c\xe6\x07\xa8\xe5\x6c\x90\xfe\x27\x18\xdf\x89\x3d\x5c\x94\xd9\x6d\x40\x4a\xfc\xd8\xe8\x8b\x6b\x61\xb7\x30\x87\xcc\x6e\x07\x17\xd2\x3f\x01\x00\x00\xff\xff\xb9\xa3\xa9\x1a\x96\x0e\x00\x00")

// File50SelectorGoTmpl is "50_selector.go.tmpl"
var File50SelectorGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xd4\x58\x5b\x73\xdb\xb6\x12\x7e\x8e\x7e\xc5\x46\x43\xfb\x50\x0e\x43\xbd\x9c\x39\x0f\x3a\xe3\xce\x74\x1c\x67\x9a\x5e\x12\x37\x4e\xdb\x07\x8f\xa6\x81\xc9\xa5\x05\x9b\x02\x18\x00\x94\xa2\x61\xf8\xdf\x3b\xb8\xf1\x22\x91\x96\xdd\x3e\xd5\x2f\x09\x81\xc5\x62\xf7\xdb\x6f\x2f\xd0\x7c\x0e\x55\x15\xc4\x57\x82\x6e\x88\xc2\xba\xbe\xc6\x1c\x13\xc5\x05\x48\xf3\x1f\x09\x09\xcf\xcb\x35\x93\x90\x71\x01\xd7\xbf\xfe\x0c\x5f\x4a\x14\x14\x25\x10\x96\x9a\xb5\x82\x08\x49\xd9\x9d\xd9\x13\x7c\x2b\x27\x6a\x57\xe0\x98\x4e\x25\xca\x44\x41\x35\x01\x00\xa8\x2a\x10\x84\xdd\x21\x04\x34\x82\x20\x83\xc5\x39\xc4\x9f\x76\x05\xc6\x6f\x29\xe6\xa9\x84\xd7\x75\xed\xe5\x68\x06\x8c\x2b\x08\xb2\xf8\x9d\xfc\x88\x19\x0a\x64\x09\x36\x02\x56\x7d\x55\x05\x59\xfc\x9e\xac\xb1\xae\xe1\x96\xf3\xdc\x9f\xc5\x5c\xb6\xa2\x3f\x72\xca\xba\x82\x3d\x33\xcd\x86\xb1\xc0\xee\x5e\x27\x84\x31\x14\x8d\x22\x96\x76\x6d\xea\x7e\x26\xbc\x64\xca\xdc\x0a\xf3\x39\x94\x12\x2d\x34\xf2\x4b\x0e\x17\x1f\x7e\x7b\xff\x29\x3c\x9b\x39\x1c\x27\xf5\x64\x32\x9f\xc3\x85\x03\x95\x08\x04\xb5\x42\x60\x64\x8d\x12\x78\xe6\x50\xc7\xd4\xc3\x3e\xc9\x4a\x96\x40\x28\xe1\x6c\x10\xd1\x99\xd7\x14\xce\xe0\x66\x29\x95\xd0\x91\xa8\x26\x2f\x36\x44\x68\x0d\xb2\x59\xdc\x43\xfc\xcf\x7f\x86\x38\xcd\x40\xc6\x87\xb0\xdb\xb8\x5a\x40\x72\x09\xe7\x40\x8a\x02\x59\x1a\xea\xaf\x08\xa6\x46\xd4\xda\x5b\xd7\xd3\x99\x11\xae\x8f\xa0\xfb\x42\xa0\x2a\x05\x33\x0a\x1d\x76\x3a\x86\x16\xb9\x7b\x4e\x19\xf0\x42\x51\xce\x0c\x78\x1a\x49\x4d\xcf\xdd\x51\xd4\x8c\x0e\x83\x59\xc2\xd7\x6b\xce\x62\xbd\x70\x45\x04\x59\x4b\x0f\xdf\xbd\xb9\x66\x40\xe0\x71\x28\x1b\xb0\x64\x0f\x2d\x9f\x01\x8b\x73\x90\xf1\x1e\x0b\xff\xdf\x6e\xbf\x3c\x07\x46\xf3\x0e\x90\xf3\x39\x90\x34\xd5\x51\xb5\xce\x66\xc0\x19\x82\xe2\xb0\x26\x6c\x07\x02\x73\xa2\xbd\x8f\x1c\x5d\x40\x8b\xa8\x15\x0a\x30\x49\x58\x70\xca\x94\xd4\xd2\x6a\x45\xa5\x59\x6b\x14\x5b\xf7\x9a\x10\x99\x4f\xad\x66\xcf\xd9\xd6\x12\xfd\x77\x45\xa8\x26\x93\x5c\xb4\xb8\xb8\xa5\xbe\x9c\x41\xe8\x60\x45\xff\xd9\xf0\x2f\x1c\x19\xde\x72\x81\xf4\x8e\xfd\x84\xbb\xf8\x5a\x24\x2d\x37\xa2\xc1\xb3\xda\x2a\x4c\x47\x35\xbc\x91\xea\x31\x0d\x75\x7f\x69\xef\xd3\x52\xc3\xfa\xbc\x68\xe2\x11\xdb\x85\x70\xd6\x0a\xd7\x03\xbc\xed\xf0\xd4\xe0\xd8\x24\xb9\xae\x0a\x1a\x78\x51\x22\x6c\x57\xc8\x80\xec\xd7\x03\x90\x2b\x5e\xe6\x29\xdc\xa2\x8e\x33\xa6\x36\x58\x4f\xa5\xb1\xb9\x21\x9c\xd9\xca\x63\x01\x77\x76\xc8\xd8\xd4\x24\x67\x89\xae\x64\x20\x13\xa2\xf3\x86\xf9\x4a\xad\xaf\x22\x50\x55\xdd\x92\xe7\x2a\xf4\xd1\x8b\xb5\xc2\x30\xa5\x44\x7f\x83\xad\x30\x11\x6c\x88\x29\x38\xa9\xa0\x1b\x14\xf1\xef\x24\x2f\xb1\xaa\x68\xe6\x12\xe3\x07\x22\x3f\x30\xfc\xc4\x7f\x21\x6c\xf7\xd1\xf1\xb6\xae\x23\xc0\xaf\x54\x2a\x09\x6b\x52\xdc\x78\x5b\xae\x04\x5d\x13\xb1\xd3\x41\x35\xdf\x97\x5f\x95\x36\x0f\x02\xb7\x4d\x92\x07\x72\x87\x75\xbd\x3c\xf3\x27\xc6\x24\xaa\x0a\x59\x5a\xd7\x33\x08\x8f\x8a\x46\x80\x42\x68\xdf\x2c\x8e\x54\xe1\xda\x2c\xd9\x94\xd5\x0e\x5b\xb4\x9d\xd7\xd6\xdd\x27\xfb\xe7\x0d\xf1\x05\x41\x2b\x3e\x48\x76\x17\x3b\x46\x73\x73\x73\x87\x68\x6e\xe7\x54\x5b\x15\xf7\x43\x16\x69\xf9\x4e\x9c\x2d\xeb\x9e\x12\x6c\x27\xf9\xf4\x88\xf7\x00\xf8\x97\x85\xbd\xe3\xf2\x5e\xa4\xe5\x96\xaa\x64\x05\xde\x2d\x3f\x9c\xbc\xee\x16\x78\xbf\xb9\x38\x87\x20\x7e\x63\x3f\x24\xf8\xee\x4f\x24\x9a\x72\xe4\xa4\xdc\x5d\xd3\xc5\x7e\x60\x65\xac\xc3\x72\x20\x18\x36\x4c\x0a\x9e\x4d\xa5\xbd\xce\x99\x62\x46\xca\x5c\x2d\x06\x29\x95\xad\x55\x7c\xa9\x3d\xcf\xc2\x69\xc9\x64\x59\x14\x5c\xe8\x51\xc3\x3b\x77\x22\xa7\x91\xff\xf0\x55\xae\x9e\x4c\xfa\xad\x6e\x14\x89\xf9\x1c\x06\xbd\x73\x54\x3c\x5c\xef\xb0\x32\xf8\x3b\x35\x68\xf8\xb6\x70\x8c\x8e\xc1\x33\xf8\x18\x3c\x9f\x90\xc1\x33\x18\x19\x1c\xa3\xa4\x9e\x3f\xc2\x36\x86\x7c\x0b\x43\x87\x1a\x01\x92\xe7\xa0\x8b\x94\x6d\x5b\xdf\xe7\x79\x38\x6b\xf6\x28\xd0\x8e\xa4\xe0\xdb\x4b\xeb\x68\x33\x22\xcf\xc6\xe6\x99\xa0\x37\x1b\x1e\x1b\x0d\xdb\x59\x47\x5b\xf3\xed\xdb\x91\x01\x91\x66\x40\xe1\xbb\x73\xc8\x91\x99\x88\xcd\xf6\xe6\x85\x31\xde\xea\x9b\x91\xf1\xf2\x6e\xd5\x3c\x4c\xac\x28\xa6\x0b\x38\x49\xa7\x51\xab\xb1\xc5\xa0\xee\x5e\xab\xf7\x6e\xe8\xd2\x57\xdf\xd3\x53\x78\xd9\xa2\x52\x4d\x3a\x8c\xba\xe0\x6c\x83\x42\x19\x0a\x5d\xf0\xd4\x07\x15\x82\xac\xae\x07\x74\x5b\x68\xf4\xb3\xe8\x11\xaa\x35\x99\xda\x9e\xc0\x2f\xe0\x10\x82\x43\xde\x99\xe5\xee\x99\xf9\x1c\x92\x15\x26\x0f\xfa\xe4\x16\x4d\x12\x30\x3d\x36\xe8\x01\x4f\x77\x07\x3d\x02\x16\x02\x37\x94\x97\xd2\xbc\xc7\x62\x78\x67\x24\x53\x9a\x46\x20\x51\xc1\xe7\xc6\xdd\xcf\x51\x57\xad\xe4\x6e\x76\xf4\xc0\x52\xe6\xe6\x46\x72\x9b\x23\x6c\x39\xfb\x8f\xd2\x73\x0a\x6e\x48\x5e\x12\x85\x69\x0c\x7f\x20\x70\x96\xef\x80\x21\xa6\x1a\xd8\x12\x65\x57\x63\x26\xf8\xda\xcf\xa3\x5a\x87\x8c\xbb\x81\xb0\x19\x77\x23\xf8\x36\x1e\xca\x38\x4b\x99\xe5\x61\x97\xec\xf3\xf8\xdc\x4c\x57\xc3\xe1\xe8\x56\xc6\x91\x25\xfa\xea\xd5\xde\x3c\xa7\x9f\x8b\x2d\xc5\x8d\x55\xd7\x39\x3d\x60\x78\x65\x20\x7b\xa0\x85\x7e\xe9\xe9\x11\x14\x1e\x70\xe7\x1f\x79\xe3\xda\xed\xb4\xd8\xff\x6a\x9f\x54\xf6\x1d\xd9\xba\xea\x5a\xd3\x86\xe4\x3a\x21\x1d\x73\xe3\x50\x0f\xf2\xdd\x84\x31\x0d\x88\x32\xf5\xbf\xff\x2e\xf6\x51\x8a\x6d\x93\x37\x87\xfb\xf2\x37\xcb\xdb\x9d\xc2\xf1\x03\xcd\x7c\x2f\x24\xbe\x63\x4a\xe7\x54\x9b\x51\x07\x7d\x66\x3f\x67\xdd\xe9\x4b\x21\x5c\x16\x85\x53\x3f\xf9\x4e\x23\xa0\x91\x77\x26\x82\xa9\x31\x3c\x72\xf6\x4c\x07\xb3\xb6\x81\xf1\x48\xad\xea\x3c\xbe\xda\x68\xdd\x8f\x3c\xba\xee\x0f\xa9\xa5\xd6\x45\x33\xf3\xdd\x9b\x99\x2f\x1c\xe8\xeb\xde\xf8\xc5\x72\xd6\xe3\xf3\xd0\x48\xf7\x58\x2d\x93\xe5\xad\xcd\x5f\xf3\x68\x6f\x4d\x8b\xdc\xd3\xfd\x24\x5d\x9c\xa4\xcb\x85\xed\xcb\xb4\x53\xd9\x8c\x8d\xb3\xd1\x12\xd4\xa7\x6d\x97\xed\x2e\xd7\x9a\x7a\xdc\x3c\xff\xf6\x36\x22\x30\x2d\xb3\x9b\x03\x57\xfa\x2d\x89\xa2\xae\xcf\x5c\x2b\x53\xeb\x62\xd6\xcb\xac\xee\xef\x2c\xc3\x97\xf5\x9a\xc7\xa8\xd6\xb1\x74\x1d\x4c\x1b\x3f\x18\x0b\xbe\xf5\x83\x70\x47\x42\xa7\xa8\xef\x89\x4e\xd4\xbd\xc6\x8c\x1d\xfe\x0d\xb6\x25\x12\x64\x81\x09\xcd\x68\x42\xf2\x7c\xd7\xfc\x0a\x73\x7c\x06\x69\x3b\xee\xc0\x23\xec\x09\xbf\x70\x75\x71\xee\xb7\x53\x78\x39\xd4\x41\x4f\x4f\xc1\x21\xe5\xfe\xd1\x62\xfe\xa9\xf7\x57\x00\x00\x00\xff\xff\x9b\x17\x65\x0f\xcc\x13\x00\x00")

// File60WhereGoTmpl is "60_where.go.tmpl"
var File60WhereGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xb4\x92\x41\x8b\xdb\x30\x10\x85\xcf\xd6\xaf\x18\x16\x1f\xec\x62\xb4\xf7\xc2\x5e\x52\x5c\xea\x4b\xba\x94\x40\x0e\x21\x14\xc5\x1e\xdb\x22\xb2\x64\x64\x39\x69\x10\xf3\xdf\x8b\xe5\xb8\x04\x4a\x68\x28\x9b\xa3\xde\xcc\xbc\x79\x7c\x1a\x77\xe9\x11\xbc\x8f\xf9\xfb\x78\x50\xb2\x24\xda\xb6\x68\x71\x35\x4a\x55\xa1\x85\xc1\xd9\xb1\x74\xe0\x89\x31\xef\xad\xd0\x0d\x42\xfc\x33\x83\xb8\x86\xcf\x6f\xc0\x37\x97\x1e\xf9\x57\x89\xaa\x1a\x88\xd8\xeb\xeb\xdf\x3e\xde\xc7\x35\x5f\x8b\x0e\x89\x40\x54\xd5\x00\x02\x4a\xa3\x2b\xe9\xa4\xd1\x60\x34\xdc\xd6\x9d\x01\xd7\x22\x6c\xbf\xe5\x3f\x72\x18\x9c\x70\xd8\xa1\x76\xac\x1e\x75\x09\xc9\xa7\x7b\x11\xd3\x5b\x8f\xc4\xf4\x50\x9a\xae\x33\x9a\x7f\xef\x33\x38\x09\x35\x57\x43\xd0\xfc\x97\x9b\xba\x20\x9e\x9f\xef\xa2\x3c\x8a\x06\x89\xd2\x65\x24\xd8\x82\x67\x91\x45\x37\x5a\xbd\xc8\x6b\x3c\x87\x4a\x62\xfa\x0c\x5e\xa6\x1c\x61\x7e\x23\x0e\x0a\x89\x5e\x66\xad\xe6\x5f\x8c\x1a\x3b\x1d\x84\x93\x50\x29\x23\xf6\x2f\x22\x85\xbe\x32\xd1\x50\xac\x9f\xca\xa5\xd0\xc9\x49\xa8\x01\x38\xe7\xff\xc3\x43\xd8\x66\x98\x3e\xbc\x13\x47\x4c\x76\x7b\xa9\x1d\xda\x5a\x94\xe8\x29\x03\x85\xb3\x77\x9a\xb2\xa8\x36\x16\xe4\xd4\x38\x5f\x4a\x58\xe9\x59\x14\xe6\x77\x72\x0f\x6f\x41\xda\xc9\x3d\x8b\xe8\x1e\xe5\x42\x27\x0f\x31\x9e\x3c\x39\xe7\x8f\x70\x5e\xa1\x3b\x23\x2e\xb0\x61\x95\x6f\xb6\x79\xfe\x5c\xe0\xd7\x95\x89\x32\xe7\x0c\x5a\xd9\xb4\x1f\x78\x87\x8b\xf7\x43\x98\xfe\x04\x98\x40\x79\x8f\xba\x22\x62\xbf\x03\x00\x00\xff\xff\xa5\x82\x54\xfa\xf4\x03\x00\x00")



func init() {
  if CTX.Err() != nil {
		log.Fatal(CTX.Err())
	}



var err error



  




  
  var f webdav.File
  

  
  
  var rb *bytes.Reader
  var r *gzip.Reader
  
  

  
  
  
  rb = bytes.NewReader(File00ImportGoTmpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "00_import.go.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(File10OrmGoTmpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "10_orm.go.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(File30ExecGoTmpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "30_exec.go.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(File40SelectGoTmpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "40_select.go.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(File50SelectorGoTmpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "50_selector.go.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  
  
  rb = bytes.NewReader(File60WhereGoTmpl)
  r, err = gzip.NewReader(rb)
  if err != nil {
    log.Fatal(err)
  }

  err = r.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "60_where.go.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  
  _, err = io.Copy(f, r)
  if err != nil {
    log.Fatal(err)
  }
  
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  


  Handler = &webdav.Handler{
    FileSystem: FS,
    LockSystem: webdav.NewMemLS(),
  }


}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {
  f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
  if err != nil {
    return nil, err
  }

  return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
  f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
  if err != nil {
    return nil, err
  }

  buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

  // If the buffer overflows, we will get bytes.ErrTooLarge.
  // Return that as an error. Any other panic remains.
  defer func() {
    e := recover()
    if e == nil {
      return
    }
    if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
      err = panicErr
    } else {
      panic(e)
    }
  }()
  _, err = buf.ReadFrom(f)
  return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
  f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
  if err != nil {
    return err
  }
  n, err := f.Write(data)
  if err == nil && n < len(data) {
    err = io.ErrShortWrite
  }
  if err1 := f.Close(); err == nil {
    err = err1
  }
  return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}


