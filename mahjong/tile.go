package mahjong

import (
	"strings"
)

// TileExpression 麻将牌表达式基类
type TileExpression struct {
	s2i map[string]ITileID
	i2s map[ITileID]string
}

func NewTileExpression() *TileExpression {
	return &TileExpression{
		s2i: make(map[string]ITileID),
		i2s: make(map[ITileID]string),
	}
}

func (t *TileExpression) AddPrefix(prefix string, flag int) {
	// 添加前缀映射
}

func (t *TileExpression) IdToName(card ITileID) string {
	if name, ok := t.i2s[card]; ok {
		return name
	}
	return ""
}

func (t *TileExpression) NameToId(name string) ITileID {
	if id, ok := t.s2i[name]; ok {
		return id
	}
	return TileNull
}

func (t *TileExpression) CommaNamesToIds(commaString string) []ITileID {
	names := strings.Split(commaString, ",")
	ids := make([]ITileID, 0, len(names))
	for _, name := range names {
		if id := t.NameToId(name); id != TileNull {
			ids = append(ids, id)
		}
	}
	return ids
}

func (t *TileExpression) IdsToCommaNames(tiles []ITileID) string {
	names := make([]string, 0, len(tiles))
	for _, tile := range tiles {
		if name := t.IdToName(tile); name != "" {
			names = append(names, name)
		}
	}
	return strings.Join(names, ",")
}

func (t *TileExpression) addPair(key string, tile ITileID) {
	t.s2i[key] = tile
	t.i2s[tile] = key
}

// TileAINameCvt AI名称转换器
type TileAINameCvt struct {
	*TileExpression
}

func NewTileAINameCvt() *TileAINameCvt {
	cvt := &TileAINameCvt{
		TileExpression: NewTileExpression(),
	}
	// 初始化AI名称映射
	// 这里需要添加具体的映射关系
	return cvt
}

func (t *TileAINameCvt) IdsToNames(tiles []ITileID) string {
	return t.IdsToCommaNames(tiles)
}

func (t *TileAINameCvt) NamesToIds(str string) []ITileID {
	return t.CommaNamesToIds(str)
}

// TileNameCvt 中文名称转换器
type TileNameCvt struct {
	*TileExpression
}

func NewTileNameCvt() *TileNameCvt {
	cvt := &TileNameCvt{
		TileExpression: NewTileExpression(),
	}
	// 初始化中文名称映射
	// 这里需要添加具体的映射关系
	return cvt
}

func (t *TileNameCvt) TilesFromNames(names []string) []ITileID {
	ids := make([]ITileID, 0, len(names))
	for _, name := range names {
		if id := t.NameToId(name); id != TileNull {
			ids = append(ids, id)
		}
	}
	return ids
}

// TileAssociation 麻将牌关联关系
type TileAssociation struct {
	vvCards [][]int
}

func NewTileAssociation() *TileAssociation {
	return &TileAssociation{
		// 初始化关联关系
	}
}

func (t *TileAssociation) IsAssociationCards(tile, drawTile ITileID) bool {
	// 判断两张牌是否关联
	return false
}

func (t *TileAssociation) GetTileAssociationCards(drawTile ITileID) map[ITileID]struct{} {
	// 获取关联牌集合
	return make(map[ITileID]struct{})
}

// 全局转换器实例
var (
	tileAINameCvt   = NewTileAINameCvt()
	tileNameCvt     = NewTileNameCvt()
	tileAssociation = NewTileAssociation()
)

func GetTileAINameCvt() *TileAINameCvt {
	return tileAINameCvt
}

func GetTileNameCvt() *TileNameCvt {
	return tileNameCvt
}

func GetTileAssociation() *TileAssociation {
	return tileAssociation
}

func GetTileName(tile ITileID) string {
	return tileNameCvt.IdToName(tile)
}

func GetTileID(name string) ITileID {
	return tileNameCvt.NameToId(name)
}

func GetColorID(name string) EColor {
	// 根据颜色名称获取颜色ID
	return ColorUndefined
}

func GetColorName(color EColor) string {
	// 根据颜色ID获取颜色名称
	return ""
}
