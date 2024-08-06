package draw_tile

func DrawTile(piece string) [][4]int {
	switch piece {
	// TileX, TileY, OffsetX, OffsetY
	case "T":
		return [][4]int{{0, 6, 1, 0}, {0, 4, 0, 1}, {4, 0, 1, 1}, {0, 5, 2, 1}}
	}
	return [][4]int{}
}
