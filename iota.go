package ranges

func Iota(cnt int) Range[int] {
	return &iotaRange{cnt: cnt}
}

type iotaRange struct {
	cnt int
}

func (this *iotaRange) Iterate() Iterator[int] {
	return &iotaIterator{cnt: this.cnt, crt: -1}
}

type iotaIterator struct {
	cnt, crt int
}

func (this *iotaIterator) Value() int {
	return this.crt
}

func (this *iotaIterator) Next() bool {
	if this.crt < this.cnt {
		this.crt += 1
	}
	return this.crt < this.cnt
}
