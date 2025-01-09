package gscXml

import (
	"fmt"
	"strconv"
	"strings"
)

type XmlName struct {
	Space, Local string
}

func (n XmlName) String() string {
	arr := strings.Split(n.Space, "/")
	space := arr[len(arr)-1]
	if len(space) == 0 {
		return fmt.Sprintf("%s", n.Local)
	} else {
		return fmt.Sprintf("%s:%s", space, n.Local)
	}
}

type XmlAttr struct {
	Name  XmlName
	Value string
}

func (a XmlAttr) String() string {
	return fmt.Sprintf("%s=\"%s\"", a.Name, a.Value)
}

type XmlNode struct {
	Val  any
	Attr []XmlAttr
}

type XmlDocument struct {
	Head string
	Body map[XmlName][]*XmlNode
}

func NodeValue(key XmlName, val any) string {
	switch t := val.(type) {
	case nil:
		return ""
	case string:
		return t
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return strconv.Itoa(t.(int))
	case uintptr:
		return fmt.Sprintf("0x%s", strconv.FormatUint(uint64(t), 16))
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	case map[XmlName][]*XmlNode:
		r := ""
		for k, v := range t {
			r += NodeValue(k, v)
		}
		return r
	case []XmlNode:
		r := ""
		for _, v := range t {
			r += NodeValue(key, v)
		}
		return r
	case []*XmlNode:
		r := ""
		for _, v := range t {
			r += NodeValue(key, v)
		}
		return r
	case XmlNode:
		vKey := key.String()
		r := NodeValue(key, t.Val)
		if len(t.Attr) > 0 {
			var attrArr []string
			for _, v := range t.Attr {
				attrArr = append(attrArr, v.String())
			}
			a := strings.Join(attrArr, " ")
			vKey = fmt.Sprintf("%s %s", key, a)
		}
		if len(r) == 0 {
			return fmt.Sprintf("<%s />", vKey)
		}
		return fmt.Sprintf("<%s>%s</%s>", vKey, r, key)
	case *XmlNode:
		vKey := key.String()
		r := NodeValue(key, t.Val)
		if len(t.Attr) > 0 {
			var attrArr []string
			for _, v := range t.Attr {
				attrArr = append(attrArr, v.String())
			}
			a := strings.Join(attrArr, " ")
			vKey = fmt.Sprintf("%s %s", key, a)
		}
		if len(r) == 0 {
			return fmt.Sprintf("<%s />", vKey)
		}
		return fmt.Sprintf("<%s>%s</%s>", vKey, r, key)
	default:
		panic("unknown xml type:" + t.(string))
	}
}

func (d XmlDocument) String() string {
	bodyStr := ""
	for elName, elValue := range d.Body {
		bodyStr += NodeValue(elName, elValue)
	}
	return fmt.Sprintf("%s%s", d.Head, bodyStr)
}
