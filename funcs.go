package websocket_manager

func inSliceStr(in string, slice []string) bool {
	for _, s := range slice {
		if s == in {
			return true
			break
		}
	}
	return false
}
