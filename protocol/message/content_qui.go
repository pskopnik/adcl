package message

import (
	"github.com/seoester/adcl/protocol/encoding"
	"github.com/seoester/adcl/protocol/maybe"
)

type QUIFlag string

const (
	QUIFlagID QUIFlag = "ID"
	QUIFlagTL         = "TL"
	QUIFlagMS         = "MS"
	QUIFlagRD         = "RD"
	QUIFlagDI         = "DI"
)

var _ ParamAccessor = &QUIContent{}

type QUIContent struct {
	SID    *encoding.Base32Value
	sidStr string

	ID    maybe.Base32Value
	idStr string
	TL    maybe.Int
	tlStr string
	MS    maybe.String
	msStr string
	RD    maybe.String
	rdStr string
	DI    maybe.String
	diStr string

	Flags map[string]string

	// Known additional flags
	// RX, PT; EXT § 3.32 RDEX - Redirects Extended (EXT v1.0.8)
}

func (q *QUIContent) Positional() []string {
	return []string{q.sidStr}
}

func (q *QUIContent) PosLen() int {
	return 1
}

func (q *QUIContent) PosAt(i int) string {
	switch i {
	case 0:
		return q.sidStr
	default:
		panic("index out of range")
	}
}

func (q *QUIContent) Named() map[string]string {
	ma := make(map[string]string)

	for k, v := range q.Flags {
		ma[k] = v
	}

	if q.ID.IsSet {
		ma[q.idStr[:2]] = q.idStr[2:len(q.idStr)]
	}
	if q.TL.IsSet {
		ma[q.tlStr[:2]] = q.tlStr[2:len(q.tlStr)]
	}
	if q.MS.IsSet {
		ma[q.msStr[:2]] = q.msStr[2:len(q.msStr)]
	}
	if q.RD.IsSet {
		ma[q.rdStr[:2]] = q.rdStr[2:len(q.rdStr)]
	}
	if q.DI.IsSet {
		ma[q.diStr[:2]] = q.diStr[2:len(q.diStr)]
	}

	return ma
}

func (q *QUIContent) NamedGet(key string) (string, bool) {
	if len(key) == 2 {
		switch QUIFlag(key) {
		case QUIFlagID:
			return q.idStr, q.ID.IsSet
		case QUIFlagTL:
			return q.tlStr, q.TL.IsSet
		case QUIFlagMS:
			return q.msStr, q.MS.IsSet
		case QUIFlagRD:
			return q.rdStr, q.RD.IsSet
		case QUIFlagDI:
			return q.diStr, q.DI.IsSet
		}
	}

	val, ok := q.Flags[key]
	return val, ok
}
