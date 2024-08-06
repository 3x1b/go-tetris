package draw_tile

func DrawTile(piece string) [][4]int {
	switch piece {
	// TileX, TileY, OffsetX, OffsetY
	case "T":
		return [][4]int{{6, 0, 1, 0}, {3, 0, 0, 1}, {0, 4, 1, 1}, {5, 0, 2, 1}}
	}
	return [][4]int{}
}
