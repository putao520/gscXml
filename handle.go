package gscXml

import (
	"encoding/xml"
	"fmt"
	"strings"
)

func appendNode(m map[XmlName][]*XmlNode, k XmlName, v *XmlNode) map[XmlName][]*XmlNode {
	cp, ok := m[k]
	if ok { // 已经存在
		m[k] = append(cp, v)
	} else {
		m[k] = []*XmlNode{v}
	}
	return m
}

func NewXmlDocument(data string) XmlDocument {
	var result XmlDocument
	decoder := xml.NewDecoder(strings.NewReader(data))
	cursor := make(map[XmlName][]*XmlNode)
	var stack Stack[map[XmlName][]*XmlNode]

	// 处理 XML 声明
	firstToken, err := decoder.Token()
	if err == nil {
		if procInst, ok := firstToken.(xml.ProcInst); ok {
			result.Head = fmt.Sprintf("<?%s %s?>", procInst.Target, procInst.Inst)
		}
	}

	// 继续解析 XML 内容
	var node *XmlNode
	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}
		switch tok := token.(type) {
		case xml.StartElement:
			var attrs []XmlAttr
			for _, attr := range tok.Attr {
				attrs = append(attrs, XmlAttr{
					Name:  toGscXmdName(attr.Name),
					Value: attr.Value,
				})
			}
			name := toGscXmdName(tok.Name)
			node = new(XmlNode)
			node.Attr = attrs
			// node.Val = nil
			appendNode(cursor, name, node)
			stack.Push(cursor)
			cursor = make(map[XmlName][]*XmlNode)
			node.Val = cursor
		case xml.EndElement:
			cursor, _ = stack.Pop()
		case xml.CharData:
			content := strings.TrimSpace(string(tok))
			if content != "" {
				node.Val = content
			}
		}
	}
	result.Body = cursor
	return result
}
