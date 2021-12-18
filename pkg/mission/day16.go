package mission

import (
	"math/bits"
)

type Day16 struct {
	input        []string
	bitsRead     int
	transmission *packet
	hexInput     []rune
	hexIndex     int
	bitBuffer    uint
	bufferSpace  int
}

type packet struct {
	version    uint
	id         uint
	value      int
	subPackets []*packet
}

func (d *Day16) SetInput(input []string) {
	d.input = input
	d.hexInput = []rune(d.input[0])
	d.bufferSpace = bits.UintSize
	d.fillBuffer()
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
	lastBit := d.bitsRead + length
	packets := make([]*packet, 0)
	for d.bitsRead < lastBit {
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
	mask := ^uint(0) << (bits.UintSize - n)
	result := (mask & d.bitBuffer) >> (bits.UintSize - n)
	d.bitBuffer = d.bitBuffer << n
	d.bufferSpace += n
	d.bitsRead += n
	d.fillBuffer()
	return result
}

func (d *Day16) fillBuffer() {
	for d.bufferSpace >= 4 {
		nextHex := d.decodeNextHex()
		d.bitBuffer += nextHex << (d.bufferSpace - 4)
		d.bufferSpace -= 4
	}
}

func (d *Day16) decodeNextHex() uint {
	if d.hexIndex >= len(d.hexInput) {
		return uint(0)
	}
	hex := d.hexInput[d.hexIndex]
	d.hexIndex++
	if hex >= 'A' {
		return uint(10 + hex - 'A')
	} else {
		return uint(hex - '0')
	}
}
