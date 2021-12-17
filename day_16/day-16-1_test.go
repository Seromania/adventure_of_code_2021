package day_16

import "testing"

func Test_convertHexStringToBinaryString(t *testing.T) {
	binaryStream := convertHexStringToBinaryString("D2FE28")
	if binaryStream != "110100101111111000101000" {
		t.Errorf("Should be '110100101111111000101000' but was '%s'", binaryStream)
	}
}

func Test_get(t *testing.T) {
	binaryStream := &BinaryStream{
		binaries:      "110100101111111000101000",
		stringPointer: 0,
	}

	a := binaryStream.get(3)
	b := binaryStream.get(3)
	c := binaryStream.get(3)

	if a != "110" {
		t.Errorf("Should be '110' but was '%s'", a)
	}
	if b != "100" {
		t.Errorf("Should be '100' but was '%s'", b)
	}
	if c != "101" {
		t.Errorf("Should be '101' but was '%s'", c)
	}

	if binaryStream.stringPointer != 9 {
		t.Errorf("Pointer should be '9' but was '%d'", binaryStream.stringPointer)
	}

	d := binaryStream.get(6)

	if d != "111111" {
		t.Errorf("Should be '111111' but was '%s'", d)
	}

	e := binaryStream.get(9)

	if e != "000101000" {
		t.Errorf("Should be '000101000' but was '%s'", e)
	}
}

func Test_parseBinaryStream(t *testing.T) {
	binaryStream := &BinaryStream{
		binaries:      "110100101111111000101000",
		stringPointer: 0,
	}

	packets, _ := parseBinaryStream(binaryStream)

	if len(packets) != 1 {
		t.Errorf("Got more than one packet")
	}

	if packets[0].version != 6 {
		t.Errorf("Version is not 6, was %d", packets[0].version)
	}

	if packets[0].typeId != 4 {
		t.Errorf("TypeId is not 4, was %d", packets[0].typeId)
	}

	if packets[0].literalValueString != "011111100101" {
		t.Errorf("literalValueString is not 011111100101, was %s", packets[0].literalValueString)
	}
}

func Test_parseBinaryStream2(t *testing.T) {
	binaryStream := &BinaryStream{
		binaries:      "11010001010",
		stringPointer: 0,
	}

	packets, _ := parseBinaryStream(binaryStream)

	if len(packets) != 1 {
		t.Errorf("Got more than one packet")
	}

	if packets[0].literalValueString != "1010" {
		t.Errorf("Was awaiting '1010' got '%s'", packets[0].literalValueString)
	}
}

func Test_parseBinaryStream3(t *testing.T) {
	binaryStream := &BinaryStream{
		binaries:      "01010010001001000000000",
		stringPointer: 0,
	}

	packets, _ := parseBinaryStream(binaryStream)

	if len(packets) != 1 {
		t.Errorf("Got more than one packet")
	}

	if packets[0].literalValueString != "00010100" {
		t.Errorf("Was awaiting '00010100' got '%s'", packets[0].literalValueString)
	}
}

func Test_doDay(t *testing.T) {
	binaryStream := &BinaryStream{
		binaries:      convertHexStringToBinaryString("8A004A801A8002F478"),
		stringPointer: 0,
	}

	versionSum, _ := doDay(binaryStream)

	if versionSum != 16 {
		t.Errorf("Was awaiting '16' got '%d'", versionSum)
	}
}

func Test_doDay2(t *testing.T) {
	binaryStream := &BinaryStream{
		binaries:      convertHexStringToBinaryString("620080001611562C8802118E34"),
		stringPointer: 0,
	}

	versionSum, _ := doDay(binaryStream)

	if versionSum != 12 {
		t.Errorf("Was awaiting '12' got '%d'", versionSum)
	}
}

func Test_doDay3(t *testing.T) {
	binaryStream := &BinaryStream{
		binaries:      convertHexStringToBinaryString("C0015000016115A2E0802F182340"),
		stringPointer: 0,
	}

	versionSum, _ := doDay(binaryStream)

	if versionSum != 23 {
		t.Errorf("Was awaiting '23' got '%d'", versionSum)
	}
}

func Test_doDay4(t *testing.T) {
	binaryStream := &BinaryStream{
		binaries:      convertHexStringToBinaryString("A0016C880162017C3686B18A3D4780"),
		stringPointer: 0,
	}

	versionSum, _ := doDay(binaryStream)

	if versionSum != 31 {
		t.Errorf("Was awaiting '31' got '%d'", versionSum)
	}
}

func Test_doDay_sum(t *testing.T) {
	binaryStream := &BinaryStream{
		binaries:      convertHexStringToBinaryString("C200B40A82"),
		stringPointer: 0,
	}

	_, sum := doDay(binaryStream)

	if sum != 3 {
		t.Errorf("Was awaiting '3' got '%d'", sum)
	}
}

func Test_doDay_product(t *testing.T) {
	binaryStream := &BinaryStream{
		binaries:      convertHexStringToBinaryString("04005AC33890"),
		stringPointer: 0,
	}

	_, sum := doDay(binaryStream)

	if sum != 54 {
		t.Errorf("Was awaiting '54' got '%d'", sum)
	}
}

func Test_doDay_min(t *testing.T) {
	binaryStream := &BinaryStream{
		binaries:      convertHexStringToBinaryString("880086C3E88112"),
		stringPointer: 0,
	}

	_, sum := doDay(binaryStream)

	if sum != 7 {
		t.Errorf("Was awaiting '7' got '%d'", sum)
	}
}

func Test_doDay_max(t *testing.T) {
	binaryStream := &BinaryStream{
		binaries:      convertHexStringToBinaryString("CE00C43D881120"),
		stringPointer: 0,
	}

	_, sum := doDay(binaryStream)

	if sum != 9 {
		t.Errorf("Was awaiting '9' got '%d'", sum)
	}
}

func Test_doDay_greaterThan(t *testing.T) {
	binaryStream := &BinaryStream{
		binaries:      convertHexStringToBinaryString("D8005AC2A8F0"),
		stringPointer: 0,
	}

	_, sum := doDay(binaryStream)

	if sum != 1 {
		t.Errorf("Was awaiting '1' got '%d'", sum)
	}
}

func Test_doDay_lessThan(t *testing.T) {
	binaryStream := &BinaryStream{
		binaries:      convertHexStringToBinaryString("F600BC2D8F"),
		stringPointer: 0,
	}

	_, sum := doDay(binaryStream)

	if sum != 0 {
		t.Errorf("Was awaiting '0' got '%d'", sum)
	}
}

func Test_doDay_equal(t *testing.T) {
	binaryStream := &BinaryStream{
		binaries:      convertHexStringToBinaryString("9C005AC2F8F0"),
		stringPointer: 0,
	}

	_, sum := doDay(binaryStream)

	if sum != 0 {
		t.Errorf("Was awaiting '0' got '%d'", sum)
	}
}

func Test_doDay_equal_sub(t *testing.T) {
	binaryStream := &BinaryStream{
		binaries:      convertHexStringToBinaryString("9C0141080250320F1802104A08"),
		stringPointer: 0,
	}

	_, sum := doDay(binaryStream)

	if sum != 1 {
		t.Errorf("Was awaiting '1' got '%d'", sum)
	}
}

func Test_all_doDayTests(t *testing.T) {
	Test_doDay(t)
	Test_doDay2(t)
	Test_doDay3(t)
	Test_doDay4(t)
	Test_doDay_sum(t)
	Test_doDay_product(t)
	Test_doDay_min(t)
	Test_doDay_max(t)
	Test_doDay_greaterThan(t)
	Test_doDay_lessThan(t)
	Test_doDay_equal(t)
	Test_doDay_equal_sub(t)
}
