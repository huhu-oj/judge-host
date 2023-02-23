package api

import "github.com/google/wire"

// APISet Injection wire
var APISet = wire.NewSet(
	JudgeSet,
	CommonSet,
) // end
