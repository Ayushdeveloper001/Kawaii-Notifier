package bot

import (
	"github.com/enescakir/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

const (
	creatorChatID int64 = 6276277461
)

var (
	isBotRunning        bool
	lastUserMessageTime time.Time
	autoOff             *time.Timer
	supportMsg          tgbotapi.MessageConfig
)

var (
	infoEmoji       = emoji.Sprintf("%v", emoji.Information)
	infinityEmoji   = emoji.Sprintf("%v", emoji.Infinity)
	stopEmoji       = emoji.Sprintf("%v", emoji.RedCircle)
	okEmoji         = emoji.Sprintf("%v", emoji.GreenCircle)
	cactusEmoji     = emoji.Sprintf("%v", emoji.Cactus)
	idkEmoji        = emoji.Sprintf("%v", emoji.OpenHands)
	GreenHeartEmoji = emoji.Sprintf("%v", emoji.GreenHeart)
)

var (
	alreadyRunningMsg     = okEmoji + " Kawaii Notifier is already running\n[ ? ] /stop - to stop Kawaii Notifier"
	startMsg              = okEmoji + " Kawaii Notifier is started\n[ ? ] /help - to get all commands"
	stopMsg               = stopEmoji + " Kawaii Notifier is stopped\n[ ? ] /start - to start Kawaii Notifier"
	helpMsg               = infoEmoji + " Kawaii Notifier hints\n\n+ /help - to get all commands\n+ /start - to start Kawaii Notifier\n+ /stop - to stop Kawaii Notifier\n+ /schedule - to see schedule\n+ /bug_report - to tell about bugs you found"
	thxForBugReportMsg    = GreenHeartEmoji + " thanks for your bug report!"
	selectDayMsg          = infinityEmoji + " select day you're interested in"
	describeTheProblemMsg = cactusEmoji + " please describe the problem:"
	idkMsg                = idkEmoji + " i don't know what you mean\n[ ? ] /help - to get all commands"
	bugReportMsg          = "[ ! ] bug report from user @%s\n[ ! ] msg: %s"
	alreadyStoppedMsg     = stopEmoji + " Kawaii Notifier is already stopped\n[ ? ] /start - to start Kawaii Notifier"
	backToMenuMsg         = okEmoji + " back to menu"
	autoOffMsg            = stopEmoji + " Kawaii Notifier is stopped due to inactivity\n[ ? ] /start - to start Kawaii Notifier"
)
