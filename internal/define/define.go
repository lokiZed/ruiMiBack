package define

type GameId int64

const (
	GameIdShuErTe GameId = iota + 1
)

type GameLevel int64

const (
	GameLevel1 GameLevel = iota + 1
	GameLevel2
	GameLevel3
	GameLevel4
	GameLevel5
)

type ResponseStatus int64

const (
	ResponseStatusOk   = 1
	ResponseStatusFail = 2
)

const (
	JwtUserId      = "userId"
	JwtChallengeId = "challengeId"
	JwtExpireAt    = "expireAt"
)

const (
	FileServerUrl  = "http://8.137.119.65/ruiMi/statics"
	StaticFilePath = "/usr/local/project/static"
)
