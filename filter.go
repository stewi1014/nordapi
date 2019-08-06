package nordapi

// Filter is an interface to an object that provides a filter for NordVPN servers.
type Filter interface {
	// GetFilter returns the filter. When applying filters, they're added to the url as follows;
	// <url>?<filter1>&<filter2>
	GetFilter() string
}

// FilterList is a list of filters.
type FilterList []Filter

// GetFilter implements Filter.
func (f FilterList) GetFilter() (out string) {
	for i := range f {
		if i > 0 {
			out += "&"
		}
		out += f[i].GetFilter()
	}
	return
}
