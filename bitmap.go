package cron

type Bitmap struct {
	data []uint64
}

func NewBitmap() *Bitmap {
	return &Bitmap{}
}

func (bitmap *Bitmap) Has(num uint) bool {
	index, bit := num/64, uint(num%64)
	return int(index) < len(bitmap.data) && (bitmap.data[index]&(1<<bit)) != 0
}

func (bitmap *Bitmap) Set(num uint) {
	index, bit := num/64, uint(num%64)
	for int(index) >= len(bitmap.data) {
		bitmap.data = append(bitmap.data, 0)
	}
	if bitmap.data[index]&(1<<bit) == 0 {
		bitmap.data[index] |= 1 << bit
	}
}
