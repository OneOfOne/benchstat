// generated by stringer -type=Kernel; DO NOT EDIT

package stats

import "fmt"

const _Kernel_name = "GaussianKernelDeltaKernel"

var _Kernel_index = [...]uint8{0, 14, 25}

func (i Kernel) String() string {
	if i < 0 || i+1 >= Kernel(len(_Kernel_index)) {
		return fmt.Sprintf("Kernel(%d)", i)
	}
	return _Kernel_name[_Kernel_index[i]:_Kernel_index[i+1]]
}
