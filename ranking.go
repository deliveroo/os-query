package esquery

func RankFeature(field string, rankFeatureType RankFeatureTypeInterface) *RankFeatureQuery {
	return &RankFeatureQuery{
		field:           field,
		rankFeatureType: rankFeatureType,
	}
}

type RankFeatureQuery struct {
	rankFeatureType RankFeatureTypeInterface
	boost           *float64
	field           string
}

func (r *RankFeatureQuery) Map() map[string]any {
	rankFeature := map[string]any{
		"field": r.field,
		"log":   r.rankFeatureType.Map(),
		"boost": 1,
	}
	m := map[string]any{
		"rank_feature": rankFeature,
	}
	if r.boost != nil {
		rankFeature["boost"] = *r.boost
	}
	return m
}

func (r *RankFeatureQuery) Boost(b float64) *RankFeatureQuery {
	r.boost = &b
	return r
}

type RankFeatureTypeInterface interface {
	RankFeatureTypeInterface()
	Map() map[string]any
}

type LogQuery struct {
	scalingFactor float64
}

func (l *LogQuery) RankFeatureTypeInterface() {}

func (l *LogQuery) Map() map[string]interface{} {
	return map[string]any{
		"scaling_factor": l.scalingFactor,
	}
}

func Log(scalingFactor float64) *LogQuery {
	return &LogQuery{
		scalingFactor: scalingFactor,
	}
}
