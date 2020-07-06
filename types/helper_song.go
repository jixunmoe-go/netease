package types

func (s *SongURLInfo) IsSample() bool {
	return s.FreeTrialInfo == nil
}
