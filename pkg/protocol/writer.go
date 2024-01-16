package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"

	"github.com/google/uuid"
)

type MessageWriter struct {
	writer io.Writer // underlying writer
	buf    [16]byte
}

func NewMessageWriter(writer io.Writer) *MessageWriter {
	return &MessageWriter{writer, [16]byte{}}
}

func (w *MessageWriter) Write(p []byte) (nn int, err error) {
	return w.writer.Write(p)
}

func (w *MessageWriter) WriteInt8(v int8) error {
	_, err := w.writer.Write([]byte{byte(v)})
	return err
}

func (w *MessageWriter) WriteInt16(v int16) error {
	binary.BigEndian.PutUint16(w.buf[:2], uint16(v))
	_, err := w.writer.Write(w.buf[:2])
	return err
}

func (w *MessageWriter) WriteInt32(v int32) error {
	binary.BigEndian.PutUint32(w.buf[:4], uint32(v))
	_, err := w.writer.Write(w.buf[:4])
	return err
}

func (w *MessageWriter) WriteInt64(v int64) error {
	binary.BigEndian.PutUint64(w.buf[:8], uint64(v))
	_, err := w.writer.Write(w.buf[:8])
	return err
}

func (w *MessageWriter) WriteUint32(v uint32) error {
	binary.BigEndian.PutUint32(w.buf[:4], v)
	_, err := w.writer.Write(w.buf[:4])
	return err
}

func (w *MessageWriter) WriteVarInt(v int32) error {
	_, err := w.writer.Write(writevarint(int64(v)))
	return err
}

func (w *MessageWriter) WriteVarLong(v int64) error {
	_, err := w.writer.Write(writevarint(v))
	return err
}

func (w *MessageWriter) WriteUuid(v uuid.UUID) error {
	_, err := w.writer.Write(v[:])
	return err
}

func (w *MessageWriter) WriteFloat64(v float64) error {
	binary.BigEndian.PutUint64(w.buf[:8], uint64(v))
	_, err := w.writer.Write(w.buf[:8])
	return err
}

func (w *MessageWriter) WriteString(s string) error {
	l := len(s)
	if l > math.MaxInt16 {
		return fmt.Errorf("unable to write string with invalid length %d", l)
	}

	err := w.WriteInt16(int16(l))
	if err != nil {
		return err
	}

	_, err = w.writer.Write([]byte(s))
	return err
}

func (w *MessageWriter) WriteCompactString(s string) error {
	var l int

	if len(s) > 0 {
		l = len(s) + 1
		if l > math.MaxUint32 {
			return fmt.Errorf("unable to write string with invalid length %d", l)
		}
	}

	err := w.WriteVarInt(int32(l))
	if err != nil {
		return err
	}

	_, err = w.writer.Write([]byte(s))
	return err
}

func (w *MessageWriter) WriteNullableString(s *string) error {
	var l int
	if s == nil {
		l = -1
	} else {
		l = len(*s)
	}

	if l > math.MaxInt16 {
		return fmt.Errorf("unable to write string with invalid length %d", l)
	}

	err := w.WriteInt16(int16(l))
	if err != nil {
		return err
	}

	if s != nil {
		_, err = w.writer.Write([]byte(*s))
	}

	return err
}

func (w *MessageWriter) WriteCompactNullableString(s *string) error {
	var l int
	if s == nil || len(*s) == 0 {
		l = 0
	} else {
		l = len(*s) + 1
	}

	if l > math.MaxUint32 {
		return fmt.Errorf("unable to write string with invalid length %d", l)
	}

	err := w.WriteVarInt(int32(l))
	if err != nil {
		return err
	}

	if s != nil {
		_, err = w.writer.Write([]byte(*s))
	}

	return err
}

func (w *MessageWriter) WriteBytes(b []byte) error {
	l := len(b)
	if l > math.MaxInt16 {
		return fmt.Errorf("unable to write string with invalid length %d", l)
	}

	err := w.WriteInt16(int16(l))
	if err != nil {
		return err
	}

	_, err = w.writer.Write(b)
	return err
}

func (w *MessageWriter) WriteCompactBytes(b []byte) error {
	l := len(b) + 1
	if l > math.MaxUint32 {
		return fmt.Errorf("unable to write string with invalid length %d", l)
	}

	err := w.WriteVarInt(int32(l))
	if err != nil {
		return err
	}

	_, err = w.writer.Write(b)
	return err
}

func (w *MessageWriter) WriteNullableBytes(b []byte) error {
	var l int
	if b == nil {
		l = -1
	} else {
		l = len(b)
	}

	if l > math.MaxInt16 {
		return fmt.Errorf("unable to write string with invalid length %d", l)
	}

	err := w.WriteInt16(int16(l))
	if err != nil {
		return err
	}

	if b != nil {
		_, err = w.writer.Write(b)
	}

	return err
}

func (w *MessageWriter) WriteCompactNullableBytes(b []byte) error {
	var l int
	if b == nil {
		l = 0
	} else {
		l = len(b) + 1
	}

	if l > math.MaxUint32 {
		return fmt.Errorf("unable to write string with invalid length %d", l)
	}

	err := w.WriteVarInt(int32(l))
	if err != nil {
		return err
	}

	if b != nil {
		_, err = w.writer.Write(b)
	}

	return err
}

func (w *MessageWriter) WriteRecords(b []byte) error {
	return w.WriteNullableBytes(b)
}

func (w *MessageWriter) WriteCompactRecords(b []byte) error {
	return w.WriteCompactNullableBytes(b)
}

func writevarint(i int64) (ret []byte) {
	for i > 0 {
		ret = append(ret, byte(0x7F&i))
		i >>= 7
	}

	if len(ret) == 0 {
		ret = []byte{0}
	}

	for front, back := 0, len(ret)-1; front < back; front, back = front+1, back-1 {
		ret[front], ret[back] = ret[back], ret[front]
	}

	for i := 0; i < len(ret)-1; i++ {
		i |= 0x80
	}

	return
}
