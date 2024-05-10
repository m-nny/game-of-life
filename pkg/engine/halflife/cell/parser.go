package cell

import (
	"minmax.uk/game-of-life/pkg/utils"
)

func FromString(left, right, up, down int, r_lvl int, str [][]bool) *MacroCell {
	utils.Assert(left <= right && up <= down, "IDK")
	if left == right {
		utils.Assert(up == down, "should collapse at the same time")
		return createLeaf(str[up][left])
	}
	m_col := (left + right) / 2
	m_row := (up + down) / 2

	return createCell(
		FromString(left, m_col, up, m_row, r_lvl+1, str), FromString(m_col+1, right, up, m_row, r_lvl+1, str),
		FromString(left, m_col, m_row+1, down, r_lvl+1, str), FromString(m_col+1, right, m_row+1, down, r_lvl+1, str),
	)
}
