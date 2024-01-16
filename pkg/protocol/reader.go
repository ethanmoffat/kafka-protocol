package protocol

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/google/uuid"
)

type MessageReader struct {
	reader io.Reader // underlying reader
	buf    [16]byte  // temporary buffer used for reading values from reader
}

func NewMessageReader(reader io.Reader) *MessageReader {
	return &MessageReader{reader, [16]byte{}}
}

func (r *MessageReader) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}

func (r *MessageReader) ReadInt8() (res int8, err error) {
	err = readShared(r.reader, r.buf[:1], &res)
	return
}

func (r *MessageReader) ReadInt16() (res int16, err error) {
	err = readShared(r.reader, r.buf[:2], &res)
	return
}

func (r *MessageReader) ReadInt32() (res int32, err error) {
	err = readShared(r.reader, r.buf[:4], &res)
	return
}

func (r *MessageReader) ReadInt64() (res int64, err error) {
	err = readShared(r.reader, r.buf[:8], &res)
	return
}

func (r *MessageReader) ReadUint32() (res uint32, err error) {
	err = readShared(r.reader, r.buf[:4], &res)
	return
}

func (r *MessageReader) ReadVarInt() (int32, error) {
	ret, err := readVarInt(r.reader, 5)
	return int32(ret), err
}

func (r *MessageReader) ReadVarLong() (int64, error) {
	return readVarInt(r.reader, 10)
}

func (r *MessageReader) ReadUuid() (uuid.UUID, error) {
	_, err := r.reader.Read(r.buf[:])
	if err != nil {
		return uuid.UUID{}, err
	}

	return uuid.FromBytes(r.buf[:])
}

func (r *MessageReader) ReadFloat64() (res float64, err error) {
	err = readShared(r.reader, r.buf[:8], &res)
	return
}

func (r *MessageReader) ReadString() (str string, err error) {
	length, err := r.ReadInt16()
	if err != nil {
		return
	}

	if length <= 0 {
		err = fmt.Errorf("non-nullable string encoded with invalid length %d", length)
		return
	}

	buf := make([]byte, length)
	_, err = r.reader.Read(buf)
	if err != nil {
		return
	}

	return string(buf), nil
}

func (r *MessageReader) ReadCompactString() (str string, err error) {
	l, err := r.ReadVarInt()
	if err != nil {
		return
	}

	length := l - 1
	if length-1 <= 0 {
		err = fmt.Errorf("non-nullable compact string encoded with invalid length %d", length)
		return
	}

	buf := make([]byte, length)
	_, err = r.reader.Read(buf)
	if err != nil {
		return
	}

	return string(buf), nil
}

func (r *MessageReader) ReadNullableString() (str string, err error) {
	length, err := r.ReadInt16()
	if err != nil || length <= 0 {
		return
	}

	buf := make([]byte, length)
	_, err = r.reader.Read(buf)
	if err != nil {
		return
	}

	return string(buf), nil
}

func (r *MessageReader) ReadCompactNullableString() (str string, err error) {
	l, err := r.ReadVarInt()
	if err != nil {
		return
	}

	length := l - 1
	if length <= 0 {
		return
	}

	buf := make([]byte, length)
	_, err = r.reader.Read(buf)
	if err != nil {
		return
	}

	return string(buf), nil
}

func (r *MessageReader) ReadBytes() (buf []byte, err error) {
	length, err := r.ReadInt32()
	if err != nil {
		return
	}

	if length <= 0 {
		err = fmt.Errorf("non-nullable string encoded with invalid length %d", length)
		return
	}

	buf = make([]byte, length)
	_, err = r.reader.Read(buf)
	if err != nil {
		return
	}

	return
}

func (r *MessageReader) ReadCompactBytes() (buf []byte, err error) {
	l, err := r.ReadVarInt()
	if err != nil {
		return
	}

	length := l - 1
	if length-1 <= 0 {
		err = fmt.Errorf("non-nullable compact string encoded with invalid length %d", length)
		return
	}

	buf = make([]byte, length)
	_, err = r.reader.Read(buf)
	if err != nil {
		return
	}

	return
}

func (r *MessageReader) ReadNullableBytes() (buf []byte, err error) {
	length, err := r.ReadInt32()
	if err != nil || length <= 0 {
		return
	}

	buf = make([]byte, length)
	_, err = r.reader.Read(buf)
	if err != nil {
		return
	}

	return
}

func (r *MessageReader) ReadCompactNullableBytes() (buf []byte, err error) {
	l, err := r.ReadVarInt()
	if err != nil {
		return
	}

	length := l - 1
	if length <= 0 {
		return
	}

	buf = make([]byte, length)
	_, err = r.reader.Read(buf)
	if err != nil {
		return
	}

	return
}

func (r *MessageReader) ReadRecords() ([]byte, error) {
	return r.ReadNullableBytes()
}

func (r *MessageReader) ReadCompactRecords() ([]byte, error) {
	return r.ReadCompactNullableBytes()
}

func readShared(reader io.Reader, target []byte, val any) error {
	if _, err := reader.Read(target); err != nil {
		return err
	}

	switch v := val.(type) {
	case *int8:
		*v = int8(target[0])
	case *int16:
		*v = int16(binary.BigEndian.Uint16(target))
	case *int32:
		*v = int32(binary.BigEndian.Uint32(target))
	case *uint32:
		*v = binary.BigEndian.Uint32(target)
	case *int64:
		*v = int64(binary.BigEndian.Uint64(target))
	case *float64:
		*v = float64(binary.BigEndian.Uint64(target))
	default:
		return fmt.Errorf("can't read object of type %T", v)
	}

	return nil
}

func readVarInt(reader io.Reader, sz int) (res int64, err error) {
	bytes := make([]byte, 0, sz)

	var b [1]byte
	for _, err = reader.Read(b[:]); err != nil && b[0]&0x80 > 0 && len(bytes) < cap(bytes); _, err = reader.Read(b[:]) {
		bytes = append([]byte{(b[0] << 1) >> 1}, bytes...)
	}

	if err != nil {
		return
	}

	bytes = append([]byte{(b[0] << 1) >> 1}, bytes...)

	for _, b[0] = range bytes {
		res <<= 7
		res |= int64(b[0])
	}

	return
}
