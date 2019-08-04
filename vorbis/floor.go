package vorbis

import (
	"fmt"
	"sort"
	"unsafe"
)

var floor1InverseDB = [256]float32{
	1.0649863e-07, 1.1341951e-07, 1.2079015e-07, 1.2863978e-07,
	1.3699951e-07, 1.4590251e-07, 1.5538408e-07, 1.6548181e-07,
	1.7623575e-07, 1.8768855e-07, 1.9988561e-07, 2.1287530e-07,
	2.2670913e-07, 2.4144197e-07, 2.5713223e-07, 2.7384213e-07,
	2.9163793e-07, 3.1059021e-07, 3.3077411e-07, 3.5226968e-07,
	3.7516214e-07, 3.9954229e-07, 4.2550680e-07, 4.5315863e-07,
	4.8260743e-07, 5.1396998e-07, 5.4737065e-07, 5.8294187e-07,
	6.2082472e-07, 6.6116941e-07, 7.0413592e-07, 7.4989464e-07,
	7.9862701e-07, 8.5052630e-07, 9.0579828e-07, 9.6466216e-07,
	1.0273513e-06, 1.0941144e-06, 1.1652161e-06, 1.2409384e-06,
	1.3215816e-06, 1.4074654e-06, 1.4989305e-06, 1.5963394e-06,
	1.7000785e-06, 1.8105592e-06, 1.9282195e-06, 2.0535261e-06,
	2.1869758e-06, 2.3290978e-06, 2.4804557e-06, 2.6416497e-06,
	2.8133190e-06, 2.9961443e-06, 3.1908506e-06, 3.3982101e-06,
	3.6190449e-06, 3.8542308e-06, 4.1047004e-06, 4.3714470e-06,
	4.6555282e-06, 4.9580707e-06, 5.2802740e-06, 5.6234160e-06,
	5.9888572e-06, 6.3780469e-06, 6.7925283e-06, 7.2339451e-06,
	7.7040476e-06, 8.2047000e-06, 8.7378876e-06, 9.3057248e-06,
	9.9104632e-06, 1.0554501e-05, 1.1240392e-05, 1.1970856e-05,
	1.2748789e-05, 1.3577278e-05, 1.4459606e-05, 1.5399272e-05,
	1.6400004e-05, 1.7465768e-05, 1.8600792e-05, 1.9809576e-05,
	2.1096914e-05, 2.2467911e-05, 2.3928002e-05, 2.5482978e-05,
	2.7139006e-05, 2.8902651e-05, 3.0780908e-05, 3.2781225e-05,
	3.4911534e-05, 3.7180282e-05, 3.9596466e-05, 4.2169667e-05,
	4.4910090e-05, 4.7828601e-05, 5.0936773e-05, 5.4246931e-05,
	5.7772202e-05, 6.1526565e-05, 6.5524908e-05, 6.9783085e-05,
	7.4317983e-05, 7.9147585e-05, 8.4291040e-05, 8.9768747e-05,
	9.5602426e-05, 0.00010181521, 0.00010843174, 0.00011547824,
	0.00012298267, 0.00013097477, 0.00013948625, 0.00014855085,
	0.00015820453, 0.00016848555, 0.00017943469, 0.00019109536,
	0.00020351382, 0.00021673929, 0.00023082423, 0.00024582449,
	0.00026179955, 0.00027881276, 0.00029693158, 0.00031622787,
	0.00033677814, 0.00035866388, 0.00038197188, 0.00040679456,
	0.00043323036, 0.00046138411, 0.00049136745, 0.00052329927,
	0.00055730621, 0.00059352311, 0.00063209358, 0.00067317058,
	0.00071691700, 0.00076350630, 0.00081312324, 0.00086596457,
	0.00092223983, 0.00098217216, 0.0010459992, 0.0011139742,
	0.0011863665, 0.0012634633, 0.0013455702, 0.0014330129,
	0.0015261382, 0.0016253153, 0.0017309374, 0.0018434235,
	0.0019632195, 0.0020908006, 0.0022266726, 0.0023713743,
	0.0025254795, 0.0026895994, 0.0028643847, 0.0030505286,
	0.0032487691, 0.0034598925, 0.0036847358, 0.0039241906,
	0.0041792066, 0.0044507950, 0.0047400328, 0.0050480668,
	0.0053761186, 0.0057254891, 0.0060975636, 0.0064938176,
	0.0069158225, 0.0073652516, 0.0078438871, 0.0083536271,
	0.0088964928, 0.009474637, 0.010090352, 0.010746080,
	0.011444421, 0.012188144, 0.012980198, 0.013823725,
	0.014722068, 0.015678791, 0.016697687, 0.017782797,
	0.018938423, 0.020169149, 0.021479854, 0.022875735,
	0.024362330, 0.025945531, 0.027631618, 0.029427276,
	0.031339626, 0.033376252, 0.035545228, 0.037855157,
	0.040315199, 0.042935108, 0.045725273, 0.048696758,
	0.051861348, 0.055231591, 0.058820850, 0.062643361,
	0.066714279, 0.071049749, 0.075666962, 0.080584227,
	0.085821044, 0.091398179, 0.097337747, 0.10366330,
	0.11039993, 0.11757434, 0.12521498, 0.13335215,
	0.14201813, 0.15124727, 0.16107617, 0.17154380,
	0.18269168, 0.19456402, 0.20720788, 0.22067342,
	0.23501402, 0.25028656, 0.26655159, 0.28387361,
	0.30232132, 0.32196786, 0.34289114, 0.36517414,
	0.38890521, 0.41417847, 0.44109412, 0.46975890,
	0.50028648, 0.53279791, 0.56742212, 0.60429640,
	0.64356699, 0.68538959, 0.72993007, 0.77736504,
	0.82788260, 0.88168307, 0.9389798, 1.0}

type sF1AS struct {
	flag bool // floor1_step2_flag
	x    int  // floor1_x_list'
	y    int  // floor1_final_y'
}

type sFloor struct {
	typ uint32

	partitions    uint32
	numClass      uint8
	listPartClass [32]uint8
	classDims     [17]uint8
	classSubs     [17]uint8
	classMasters  [17]uint8
	subBooks      [17][8]int
	multiplier    int
	rangebits     uint8
	sizeXList     uint32
	xList         [34]int
	values        int
}

type f1asSlice []sF1AS

func (p f1asSlice) Len() int           { return len(p) }
func (p f1asSlice) Less(i, j int) bool { return p[i].x < p[j].x }
func (p f1asSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (fl *sFloor) readConfig(vb *Vorbis) bool {
	if debug {
		fmt.Println("  read floor config.")
	}
	fl.typ = vb.pr.ReadBits(16)
	if fl.typ != 1 {
		fmt.Println(fmt.Sprintf("unsupported floor type: %d", fl.typ))
		return false
	}
	fl.partitions = vb.pr.ReadBits(5)
	fl.numClass = 0
	for i := uint32(0); i < fl.partitions; i++ {
		fl.listPartClass[i] = uint8(vb.pr.ReadBits(4))
		if fl.listPartClass[i]+1 > fl.numClass {
			fl.numClass = fl.listPartClass[i] + 1
		}
	}

	for i := uint8(0); i < fl.numClass; i++ {
		fl.classDims[i] = uint8(vb.pr.ReadBits(3) + 1)
		fl.classSubs[i] = uint8(vb.pr.ReadBits(2))
		if fl.classSubs[i] != 0 {
			fl.classMasters[i] = uint8(vb.pr.ReadBits(8))
			if uint32(fl.classMasters[i]) >= vb.numCodebooks {
				fmt.Println("corrupted floor config: master books")
				return false
			}
		}
		jj := 0x00000001 << fl.classSubs[i]
		assert(jj < 8)
		for j := 0; j < jj; j++ {
			fl.subBooks[i][j] = int(vb.pr.ReadBits(8)) - 1
			if fl.subBooks[i][j] >= int(vb.numCodebooks) {
				fmt.Println("corrupted floor config: subclass books", fl.subBooks[i][j], vb.numCodebooks)
				return false
			}
		}
	}

	fl.multiplier = int(vb.pr.ReadBits(2) + 1)
	fl.rangebits = uint8(vb.pr.ReadBits(4))
	fl.sizeXList = fl.partitions + 2
	fl.xList[0] = 0
	fl.xList[1] = 0x0000001 << fl.rangebits
	fl.values = 2
	for i := uint32(0); i < fl.partitions; i++ {
		numCurClass := fl.listPartClass[i]
		for j := uint8(0); j < fl.classDims[numCurClass]; j++ {
			if fl.values >= int(fl.sizeXList) {
				fl.sizeXList += fl.partitions
				assert(fl.values < (int)(unsafe.Sizeof(fl.xList)))
			}
			fl.xList[fl.values] = int(vb.pr.ReadBits(uint32(fl.rangebits)))
			fl.values++
		}
	}
	return true
}

func lowNeighbor(_v []int, _x int) int {
	assert(_x > 0)
	ret := -1
	m := 0
	limit := _v[_x]
	for i := 0; i < _x; i++ {
		if _v[i] >= m && _v[i] < limit {
			ret = i
			m = _v[i]
		}
	}
	return ret
}

func highNeighbor(_v []int, _x int) int {
	assert(_x > 0)
	ret := -1
	m := 100000
	limit := _v[_x]
	for i := 0; i < _x; i++ {
		if _v[i] <= m && _v[i] > limit {
			ret = i
			m = _v[i]
		}
	}
	return ret
}

func renderPoint(_x0 int, _y0 int, _x1 int, _y1 int, _x int) int {
	dy := _y1 - _y0
	adx := _x1 - _x0
	var y int
	if dy < 0 {
		//ady :=  -dy;
		err := -dy * (_x - _x0)
		off := err / adx
		y = _y0 - off
	} else {
		//ady :=  dy;
		err := dy * (_x - _x0)
		off := err / adx
		y = _y0 + off
	}
	return y
}

func renderLine(_x0 int, _y0 int, _x1 int, _y1 int, _v []int) {
	dy := _y1 - _y0
	adx := _x1 - _x0
	//ady :=  dy < 0 ? -dy : dy;
	var ady int
	if dy < 0 {
		ady = -dy
	} else {
		ady = dy
	}
	base := dy / adx
	x := _x0
	y := _y0
	err := 0

	var sy int
	if dy < 0 {
		sy = base - 1
	} else {
		sy = base + 1
	}

	var abase int
	if base < 0 {
		abase = -base
	} else {
		abase = base
	}
	ady = ady - abase*adx
	_v[x] = y

	for x = _x0 + 1; x <= _x1-1; x++ {
		err = err + ady
		if err >= adx {
			err = err - adx
			y = y + sy
		} else {
			y = y + base
		}
		_v[x] = y
	}
}

func (fl *sFloor) decode(vb *Vorbis, _buf *sChannelBuf, halfBlockSize int) bool {
	// 7.2.2 floor1 packet decode
	assert(fl.typ == 1)
	_buf.floorUnused = vb.pr.ReadBits(1) == 0
	if _buf.floorUnused {
		f := _buf.floor[:halfBlockSize]
		for i := range f {
			f[i] = 0
		}
		return true
	}

	v1 := [4]uint32{256, 128, 86, 64}
	rnge := v1[fl.multiplier-1]
	bits := uint32(ilog(rnge - 1))
	// cdim <= 9
	// partitions <= 32
	_buf.floor1Y[0] = int(vb.pr.ReadBits(bits))
	_buf.floor1Y[1] = int(vb.pr.ReadBits(bits))
	offset := 2
	for i := uint32(0); i < fl.partitions; i++ {
		cls := fl.listPartClass[i]
		cdim := fl.classDims[cls]
		cbits := fl.classSubs[cls]
		csub := (0x00000001 << cbits) - 1
		// 10
		cval := 0
		// 11
		if cbits != 0 {
			cb := &vb.codebooks[fl.classMasters[cls]]
			cval = int(cb.decode(vb))
		}

		// 13
		for j := 0; j < int(cdim); j++ {
			// 14
			book := fl.subBooks[cls][cval&csub]
			// 15
			cval = cval >> cbits
			// 16
			if book >= 0 {
				// 17
				_buf.floor1Y[j+offset] = int(vb.codebooks[book].decode(vb))
			} else {
				// 18
				_buf.floor1Y[j+offset] = 0
			}
		}
		// 19
		offset += int(cdim)
	}
	_buf.sizeFloor1Y = offset

	// 7.2.2 curve computation
	// step 1: amplitude value synthesis
	var f1as [34]sF1AS
	f1as[0].flag = true
	f1as[1].flag = true
	f1as[0].y = _buf.floor1Y[0]
	f1as[1].y = _buf.floor1Y[1]
	for i := 0; i < fl.values; i++ {
		f1as[i].x = int(fl.xList[i])
	}

	for i := 2; i < fl.values; i++ {
		lowOff := lowNeighbor(fl.xList[:], i)
		highOff := highNeighbor(fl.xList[:], i)
		assert(lowOff != i && highOff != i && highOff != lowOff)
		predicted := renderPoint(
			f1as[lowOff].x, f1as[lowOff].y,
			f1as[highOff].x, f1as[highOff].y,
			f1as[i].x)
		val := _buf.floor1Y[i]
		highroom := int(rnge) - predicted
		lowroom := predicted

		var room int
		if highroom < lowroom {
			room = highroom * 2
		} else {
			room = lowroom * 2
		}
		if val != 0 {
			f1as[lowOff].flag = true
			f1as[highOff].flag = true
			f1as[i].flag = true
			// 20
			if val >= room {
				//21
				if highroom > lowroom {
					//22
					f1as[i].y = val - lowroom + predicted
				} else {
					//23
					f1as[i].y = predicted - val + highroom - 1
				}
			} else {
				// 24
				if val&0x01 != 0 {
					// 25
					f1as[i].y = predicted - ((val + 1) / 2)
				} else {
					// 26
					f1as[i].y = predicted + (val / 2)
				}
			}
		} else {
			// 27
			f1as[i].flag = false
			// 28
			f1as[i].y = predicted
		}
	}
	//29

	// step 2, curve synthesis
	var floorBuf [8192 + 1024]int // ?

	sort.Sort(f1asSlice(f1as[:fl.values]))
	hx := 0
	hy := 0
	lx := 0
	ly := f1as[0].y * fl.multiplier
	// 4
	for i := 1; i < fl.values; i++ {
		// 5
		if f1as[i].flag {
			// 6
			hy = f1as[i].y * fl.multiplier
			// 7
			hx = f1as[i].x
			// 8
			renderLine(lx, ly, hx, hy, floorBuf[:])
			// 9
			lx = hx
			// 10
			ly = hy
		}
	}
	if hx < halfBlockSize {
		renderLine(hx, hy, halfBlockSize, hy, floorBuf[:])
	} // else if hx > halfBlockSize {
	// 	// truncate
	// }
	for i := 0; i < halfBlockSize; i++ {
		assert(floorBuf[i] >= 0 && floorBuf[i] < 256)
		_buf.floor[i] = floor1InverseDB[floorBuf[i]]
	}
	return true
}
