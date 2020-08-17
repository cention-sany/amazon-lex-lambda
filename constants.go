package lexlambda

const (
	E_CREATE_ERRAND = iota + 1
	E_SPEAK_TO_AGENT
)

const base26Encoder = "abcdefghijklmnopqrstuvwxyz"

// prefix for integer (PI).
const (
	PI_BOTNAME      = "c"
	PI_INTENT_EVENT = "e"
	PI_INTENT_BY_ID = "l"
)

// intent name prefixes (INP).
const (
	INP_EVENT  = "e"
	INP_NORMAL = "n"
	INP_BY_ID  = "l"
	INP_GLOBAL = "g"
)
