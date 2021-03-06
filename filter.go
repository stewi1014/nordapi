package nordapi

import "fmt"

// Filter is an interface to an object that provides a filter for NordVPN servers.
type Filter interface {
	// GetFilter returns the filter. When applying filters, they're added to the url as follows;
	// <url>?<filter1>&<filter2>
	GetFilter() string

	// Satisfies returns true if the given server satisfies the filter.
	Satisfies(*Server) bool

	// Human readable name
	fmt.Stringer
}

// FilterList is a list of filters.
type FilterList []Filter

// GetFilter implements Filter.
func (f FilterList) GetFilter() (out string) {
	for i := range f {
		filter := f[i].GetFilter()
		if filter == "" {
			continue
		}
		if i > 0 {
			out += "&"
		}
		out += filter
	}
	return
}

// Satisfies implements Filter
func (f FilterList) Satisfies(s *Server) bool {
	for i := range f {
		if !f[i].Satisfies(s) {
			return false
		}
	}
	return true
}

// String implements fmt.Stringer
func (f FilterList) String() (str string) {
	for i := range f {
		if i > 0 {
			str += " "
		}
		str += f[i].String()
	}
	return
}
