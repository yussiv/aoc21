package mission

import (
	"github.com/yussiv/aoc21/util"
)

type Day16 struct {
	input        []string
	binaryInput  []uint
	iWord        int
	iArray       int
	wordLength   int
	iTotal       int
	transmission *packet
}

type packet struct {
	version    uint
	id         uint
	value      int
	subPackets []*packet
}

func (d *Day16) SetInput(input []string) {
	d.input = input
	d.parseInput()
	d.transmission = d.getPacket()
}

func (d *Day16) Task1() int {
	return d.getVersionSum(d.transmission)
}

func (d *Day16) Task2() int {
	return d.transmission.value
}

func (d Day16) getVersionSum(p *packet) int {
	sum := int(p.version)
	if p.subPackets != nil {
		for _, sub := range p.subPackets {
			sum += d.getVersionSum(sub)
		}
	}
	return sum
}

func (d *Day16) getPacket() *packet {
	p := new(packet)
	p.version = d.getNextBits(3)
	p.id = d.getNextBits(3)
	if p.id == 4 {
		p.value = d.getLiteral()
	} else {
		lengthBit := d.getNextBits(1)
		var amount int
		if lengthBit == 0 {
			amount = int(d.getNextBits(15))
			p.subPackets = d.getSubPacketsByLength(amount)
		} else {
			amount = int(d.getNextBits(11))
			p.subPackets = d.getSubPacketsByCount(amount)
		}
		switch p.id {
		case 0:
			for _, sub := range p.subPackets {
				p.value += sub.value
			}
		case 1:
			p.value = 1
			for _, sub := range p.subPackets {
				p.value *= sub.value
			}
		case 2:
			p.value = int(^uint(0) >> 1)
			for _, sub := range p.subPackets {
				if sub.value < p.value {
					p.value = sub.value
				}
			}
		case 3:
			for _, sub := range p.subPackets {
				if sub.value > p.value {
					p.value = sub.value
				}
			}
		case 5:
			if p.subPackets[0].value > p.subPackets[1].value {
				p.value = 1
			}
		case 6:
			if p.subPackets[0].value < p.subPackets[1].value {
				p.value = 1
			}
		case 7:
			if p.subPackets[0].value == p.subPackets[1].value {
				p.value = 1
			}
		}
	}
	return p
}

func (d *Day16) getSubPacketsByCount(count int) []*packet {
	packets := make([]*packet, count)
	for i := 0; i < count; i++ {
		packets[i] = d.getPacket()
	}
	return packets
}

func (d *Day16) getSubPacketsByLength(length int) []*packet {
	iEnd := d.iTotal + length
	packets := make([]*packet, 0)
	for d.iTotal < iEnd {
		packets = append(packets, d.getPacket())
	}
	return packets
}

func (d *Day16) getLiteral() int {
	mask := uint(15)
	result := 0
	for {
		segment := d.getNextBits(5)
		result = (result << 4) + int(segment&mask)
		if (segment & ^mask) == 0 {
			break
		}
	}
	return result
}

func (d *Day16) getNextBits(n int) uint {
	result := uint(0)
	mask := ^uint(0) << (d.wordLength - n)
	word := d.binaryInput[d.iArray]
	end := d.iWord + n
	if diff := end - d.wordLength; diff > 0 {
		result += (word & (mask >> d.iWord) << diff)
		d.iArray++
		word = d.binaryInput[d.iArray]
		d.iWord = 0
		result += (word & (mask << (n - diff))) >> (d.wordLength - diff)
		d.iWord = diff
	} else {
		result += (word & (mask >> d.iWord)) >> (d.wordLength - d.iWord - n)
		d.iWord += n
	}
	d.iTotal = d.iArray*d.wordLength + d.iWord
	return result
}

func (d *Day16) parseInput() {
	d.wordLength = 64
	hexPerWord := d.wordLength / 4
	hexInput := d.input[0]
	d.binaryInput = make([]uint, len(hexInput)/hexPerWord+1) // zero padding doesn't hurt
	n := len(hexInput)
	k := 0
	for i := 0; i < n; i += hexPerWord {
		section := hexInput[i:util.Min(i+hexPerWord, n)]
		word := uint(0)
		for j, r := range section {
			var value uint
			if r >= 'A' {
				value = uint(10 + r - 'A')
			} else {
				value = uint(r - '0')
			}
			word += value << (4 * (hexPerWord - 1 - j))
		}
		d.binaryInput[k] = word
		k++
	}
}
