// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/avfilter.h>
*/
import "C"

//Get a filter definition matching the given name.
func AvfilterGetByName(n string) *Filter {
	return (*Filter)(C.avfilter_get_by_name(C.CString(n)))
}

// Deprecated
// https://github.com/FFmpeg/FFmpeg/commit/8f1382f80e0d4184c54c14afdda6482f050fbba7
//Register a filter.
// func (f *Filter) AvfilterRegister() int {
// 	return int(C.avfilter_register((*C.struct_AVFilter)(f)))
// }

// Deprecated
// https://github.com/FFmpeg/FFmpeg/commit/8f1382f80e0d4184c54c14afdda6482f050fbba7
//Iterate over all registered filters.
// func (f *Filter) AvfilterNext() *Filter {
// 	return (*Filter)(C.avfilter_next((*C.struct_AVFilter)(f)))
// }
