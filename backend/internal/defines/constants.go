package defines

import (
	"time"
)

const (
	// Errors
	CannotParseJSONRequest            = "cannot parse JSON request"
	CannotHashPassword                = "cannot hash password"
	CannotGenerateSlugID              = "cannot generate slug id"
	CannotSaveUserOnDatabase          = "cannot save user on database"
	CannotSendEmail                   = "cannot send email"
	CannotGetUserByEmail              = "cannot get user by email"
	CannotGetUserByUsername           = "cannot get user by username"
	CannotGetUserBySlugID             = "cannot get user by slug id"
	CannotGetProfileByUserID          = "cannot get profile by user id"
	CannotUpdateUser                  = "cannot update user"
	CannotUpdateUserPassword          = "cannot update user password"
	CannotUpdateProfile               = "cannot update profile"
	CannotValidateLogin               = "cannot validate login"
	CannotValidateJwt                 = "cannot validate jwt"
	CannotCheckBrokenAccessControl    = "cannot check broken access control"
	CannotValidateRequest             = "cannot validate request"
	InvalidRequestBody                = "invalid request body"
	AccountAlreadyActive              = "account already active"
	WrongEmailOrPassword              = "wrong email or password"
	UsernameMustToBeProvided          = "username must to be provided"
	AccountNotActiveYet               = "account not active yet"
	AuthorizationMustToBeProvided     = "authorization must to be provided"
	UnauthorizedRequest               = "unauthorized request"
	CannotCreateProfileOnDatabase     = "cannot create profile on database"
	CannotGenerateJwt                 = "cannot generate jwt"
	CannotGetGender                   = "cannot get gender"
	CannotGetGenderID                 = "cannot get gender id"
	CannotGetSexualPreference         = "cannot get sexual preference"
	CannotGetSexualPreferenceID       = "cannot get sexual preference id"
	CannotDeleteUser                  = "cannot delete user"
	CannotEstablishDatabaseConnection = "cannot establish database connection"
	CannotDeleteProfile               = "cannot delete profile"
	WeakPassword                      = "password must be at least 8 characters long and contain uppercase, lowercase, number, special character, and cannot be a common English word"

	// Regex
	EmailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	PendingStatus                    = "pending"
	ActiveStatus                     = "active"
	EmptyString                      = ""
	EmptyJson                        = "{}"
	JwtExpirationTime                = time.Hour * 1
	RecoverPasswordJwtExpirationTime = time.Minute * 15
	SexualPreferenceSeparator        = "_"

	ActiveAccountEmailSubject   = "Welcome to Matcha!"
	ActiveAccountEmailContent   = "Welcome aboard!\n\nClick here to activate your account: "
	RecoverPasswordEmailSubject = "Recover your Matcha password"
	RecoverPasswordEmailContent = "Click here to recover your password: "
)

var (
	SqlKeywords       = []string{"select", "insert", "update", "delete", "drop", "union", "or", "and"}
	CommonSqlPatterns = []string{
		"(?i)\\bselect\\b", "(?i)\\binsert\\b", "(?i)\\bupdate\\b", "(?i)\\bdelete\\b",
		"(?i)\\bdrop\\b", "(?i)\\bunion\\b", "(?i)\\b--\\b", "(?i)\\b;\\b",
		"(?i)\\b'\\b", "(?i)\\b\"\\b", "(?i)\\b#\\b", "(?i)\\b--\\b",
	}

	// CommonEnglishWords lista de palavras comuns em inglês que não devem ser aceitas como senhas
	CommonEnglishWords = []string{
		"password", "123456", "123456789", "qwerty", "abc123", "password123",
		"admin", "user", "login", "welcome", "monkey", "dragon", "master",
		"hello", "freedom", "letmein", "trustno1", "god", "sex", "secret",
		"aaron", "abby", "abc", "abcd", "abc123", "about", "access", "account",
		"action", "active", "activity", "add", "address", "admin", "administrator",
		"advanced", "again", "against", "age", "agent", "ago", "agree", "agreement",
		"ahead", "air", "all", "allow", "almost", "alone", "along", "already",
		"also", "always", "among", "amount", "and", "another", "answer", "any",
		"anyone", "anything", "appear", "apple", "area", "arm", "army", "around",
		"art", "article", "artist", "as", "ask", "assume", "at", "attack", "attention",
		"author", "authority", "available", "avoid", "away", "baby", "back", "bad",
		"bag", "ball", "bank", "base", "be", "beat", "beauty", "because", "become",
		"bed", "before", "begin", "behind", "believe", "best", "better", "between",
		"big", "bill", "billion", "bit", "black", "blood", "blue", "board", "body",
		"book", "born", "both", "box", "boy", "break", "bring", "brother", "build",
		"business", "buy", "call", "camera", "can", "car", "card", "care", "carry",
		"case", "catch", "cause", "center", "century", "certain", "chair", "chance",
		"change", "character", "charge", "check", "child", "choice", "choose",
		"church", "city", "civil", "claim", "class", "clear", "close", "coach",
		"cold", "collection", "college", "color", "come", "commercial", "common",
		"company", "compare", "computer", "concern", "condition", "conference",
		"congress", "consider", "consumer", "contain", "continue", "control",
		"cost", "could", "country", "couple", "course", "court", "cover", "create",
		"crime", "cultural", "culture", "cup", "current", "customer", "cut", "dark",
		"data", "daughter", "day", "dead", "deal", "death", "debate", "decade",
		"decide", "decision", "deep", "defense", "degree", "democrat", "democratic",
		"describe", "design", "despite", "detail", "determine", "develop", "development",
		"die", "difference", "different", "difficult", "dinner", "direction", "director",
		"discover", "discuss", "discussion", "disease", "do", "doctor", "dog", "door",
		"down", "draw", "dream", "drive", "drop", "drug", "during", "each", "early",
		"east", "easy", "eat", "economic", "economy", "edge", "education", "effect",
		"effort", "eight", "either", "election", "else", "employee", "end", "energy",
		"enjoy", "enough", "enter", "entire", "environment", "environmental", "especially",
		"establish", "even", "evening", "event", "ever", "every", "everybody", "everyone",
		"everything", "evidence", "exactly", "example", "executive", "exist", "expect",
		"experience", "expert", "explain", "eye", "face", "fact", "factor", "fail",
		"fall", "family", "far", "fast", "father", "fear", "federal", "feel", "feeling",
		"few", "field", "fight", "figure", "fill", "film", "final", "finally", "financial",
		"find", "fine", "finger", "finish", "fire", "firm", "first", "fish", "five",
		"floor", "fly", "focus", "follow", "food", "foot", "for", "force", "foreign",
		"forget", "form", "former", "forward", "four", "free", "friend", "from", "front",
		"full", "fund", "future", "game", "garden", "gas", "general", "generation",
		"get", "girl", "give", "glass", "go", "goal", "good", "government", "great",
		"green", "ground", "group", "grow", "growth", "guess", "gun", "guy", "hair",
		"half", "hand", "hang", "happen", "happy", "hard", "have", "he", "head",
		"health", "hear", "heart", "heat", "heavy", "help", "her", "here", "herself",
		"high", "him", "himself", "his", "history", "hit", "hold", "home", "hope",
		"hospital", "hot", "hotel", "hour", "house", "how", "however", "huge", "human",
		"hundred", "husband", "I", "idea", "identify", "if", "image", "imagine",
		"impact", "important", "improve", "in", "include", "including", "increase",
		"indeed", "indicate", "individual", "industry", "information", "inside",
		"instead", "institution", "interest", "interesting", "international", "interview",
		"into", "investment", "involve", "issue", "it", "item", "its", "itself",
		"job", "join", "just", "keep", "key", "kid", "kill", "kind", "kitchen",
		"know", "knowledge", "land", "language", "large", "last", "late", "later",
		"laugh", "law", "lawyer", "lay", "lead", "leader", "learn", "least", "leave",
		"left", "leg", "legal", "less", "let", "letter", "level", "lie", "life",
		"light", "like", "likely", "line", "list", "listen", "little", "live",
		"local", "long", "look", "lose", "loss", "lot", "love", "low", "machine",
		"magazine", "main", "maintain", "major", "majority", "make", "man", "manage",
		"management", "manager", "many", "market", "marriage", "material", "matter",
		"may", "maybe", "me", "mean", "measure", "media", "medical", "meet", "meeting",
		"member", "memory", "mention", "message", "method", "middle", "might", "military",
		"million", "mind", "minute", "miss", "model", "modern", "moment", "money",
		"month", "more", "morning", "most", "mother", "mouth", "move", "movement",
		"movie", "Mr", "Mrs", "much", "music", "must", "my", "myself", "name",
		"nation", "national", "natural", "nature", "near", "nearly", "necessary",
		"need", "network", "never", "new", "news", "newspaper", "next", "nice",
		"night", "nine", "no", "none", "nor", "north", "not", "note", "nothing",
		"notice", "now", "number", "occur", "of", "off", "offer", "office", "officer",
		"official", "often", "oh", "oil", "ok", "old", "on", "once", "one", "only",
		"onto", "open", "operation", "opportunity", "option", "or", "order", "organization",
		"other", "others", "our", "out", "outside", "over", "own", "owner", "page",
		"pain", "painting", "paper", "parent", "part", "participant", "particular",
		"particularly", "partner", "party", "pass", "past", "patient", "pattern",
		"pay", "peace", "people", "per", "perform", "performance", "perhaps",
		"period", "person", "personal", "phone", "physical", "pick", "picture",
		"piece", "place", "plan", "plant", "play", "player", "PM", "point", "police",
		"policy", "political", "politics", "poor", "popular", "population", "position",
		"positive", "possible", "power", "practice", "prepare", "present", "president",
		"pressure", "pretty", "prevent", "price", "private", "probably", "problem",
		"process", "produce", "product", "production", "professional", "professor",
		"program", "project", "property", "protect", "prove", "provide", "public",
		"pull", "purpose", "push", "put", "quality", "question", "quickly", "quite",
		"race", "radio", "raise", "range", "rate", "rather", "reach", "read",
		"ready", "real", "reality", "realize", "really", "reason", "receive",
		"recent", "recently", "recognize", "record", "red", "reduce", "reflect",
		"region", "relate", "relationship", "religious", "remain", "remember",
		"remove", "report", "represent", "Republican", "require", "research",
		"resource", "respond", "response", "responsibility", "rest", "result",
		"return", "reveal", "rich", "right", "rise", "risk", "road", "rock",
		"role", "room", "rule", "run", "safe", "same", "save", "say", "scene",
		"school", "science", "scientist", "score", "sea", "season", "seat",
		"second", "section", "security", "see", "seek", "seem", "sell", "send",
		"senior", "sense", "series", "serious", "serve", "service", "set", "seven",
		"several", "sex", "sexual", "shake", "share", "she", "shoot", "short",
		"shot", "should", "shoulder", "show", "side", "sign", "significant",
		"similar", "simple", "simply", "since", "sing", "single", "sister",
		"sit", "site", "situation", "six", "size", "skill", "skin", "small",
		"so", "social", "society", "soldier", "some", "somebody", "someone",
		"something", "sometimes", "son", "song", "soon", "sort", "sound",
		"south", "space", "speak", "special", "specific", "speech", "spend",
		"sport", "spring", "staff", "stage", "stand", "standard", "star",
		"start", "state", "statement", "station", "stay", "step", "still",
		"stock", "stop", "store", "story", "strategy", "street", "strong",
		"structure", "student", "study", "stuff", "style", "subject", "success",
		"successful", "such", "suddenly", "suffer", "suggest", "summer", "support",
		"sure", "surface", "system", "table", "take", "talk", "task", "tax",
		"teach", "teacher", "team", "technology", "television", "tell", "ten",
		"tend", "term", "test", "than", "thank", "that", "the", "their", "them",
		"themselves", "then", "theory", "there", "these", "they", "thing",
		"think", "third", "this", "those", "though", "thought", "thousand",
		"threat", "three", "through", "throughout", "throw", "thus", "time",
		"today", "together", "tonight", "too", "top", "total", "tough",
		"toward", "town", "trade", "traditional", "training", "travel",
		"treat", "treatment", "tree", "trial", "trip", "trouble", "true",
		"truth", "try", "turn", "TV", "two", "type", "under", "understand",
		"unit", "until", "up", "upon", "us", "use", "usually", "value",
		"various", "very", "victim", "view", "violence", "visit", "voice",
		"vote", "wait", "walk", "wall", "want", "war", "watch", "water",
		"way", "we", "weapon", "wear", "week", "weight", "well", "west",
		"western", "what", "whatever", "when", "where", "whether", "which",
		"while", "white", "who", "whole", "whom", "whose", "why", "wide",
		"wife", "will", "win", "wind", "window", "wish", "with", "within",
		"without", "woman", "wonder", "word", "work", "worker", "world",
		"worry", "would", "write", "writer", "wrong", "yard", "yeah",
		"year", "yes", "yet", "you", "young", "your", "yourself",
	}

	userAccountStatus = []string{PendingStatus, ActiveStatus}
)
