package fbi

var _ FBI = (*Fugitives)(nil)
var _ Downloader = (*Images)(nil)
var _ Downloader = (*Files)(nil)

// FBI contains the methods available
// to the Individual type
type FBI interface {
	List() error
	Find(opt *Options) error
}

// Downloader allows locally downloading the remote
// files and images from their URLs
type Downloader interface {
	Download(filename string) error
}
