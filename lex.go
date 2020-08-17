// lexlambda provides helper functions to generate and parse bot and intent name
// for Amaazon Lex.

package lexlambda

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// BotPrefix is helper to create a string base on integer. As long as 'id' not
// repeat, then the string can be used as 'unique' for BotName.
func BotUniqueById(space string, id int) string {
	// this prefix need to be protected to avoid broken encoded base26 when use
	// in BotName.
	return limitBotName(botSpacedName(space, id))
}

func botSpacedName(space string, id int) string {
	return spaceBoundName(space, prefixedBotName(id))
}

// BotName generates name that is suitable for Amazon-Lex bot name and trying
// its best to fit in displayName for easier identification.
func BotName(unique, displayName string) string {
	return limitBotName(fmt.Sprint(Base26.Encode([]byte(unique)), "_",
		Base26Only(displayName)))
}

func ParseBotId(botName string) (int, error) {
	name := strings.Split(botName, "_")[0]
	b, err := Base26.Decode(name)
	if err != nil {
		return 0, err
	}
	name = strings.Split(string(b), "/")[0][1:]
	return strconv.Atoi(name)
}

func IntentNameByEventId(space string, botId, eventId int) string {
	return fmt.Sprint(INP_EVENT,
		spaceBoundLocalIntentName(botId, space, strconv.Itoa(eventId)))
}

// ParseEventId try to get event id from intent name.
func ParseEventId(intentName string) (int, error) {
	if intentName[:1] != INP_EVENT {
		return 0, errors.New("not an event intent name")
	}
	return EventIdFromEncoded(intentName[1:])
}

func EventIdFromEncoded(encoded string) (int, error) {
	b, err := Base26.Decode(encoded)
	if err != nil {
		return 0, err
	}
	nameWithPrefix := strings.Split(string(b), "/")[0]
	if nameWithPrefix[:1] != PI_INTENT_EVENT {
		return 0, errors.New("not an event type intent")
	}
	return strconv.Atoi(nameWithPrefix[1:])
}

func spaceBoundLocalIntentName(botId int, space, name string) string {
	return Base26.Encode([]byte(LimitIntentName(spaceBoundName(space,
		fmt.Sprint(PI_INTENT_EVENT, name, "/", prefixedBotName(botId))))))
}

func prefixedBotName(id int) string {
	return fmt.Sprint(PI_BOTNAME, id)
}

func limitBotName(name string) string {
	return limitNameSize(name, botNameRawLimit)
}

const botNameRawLimit = 28 // floor(<limit>/(8/log2(<baseX>))-1), eg: floor(50/(8/log2(26))-1)

func LimitIntentName(name string) string {
	return limitNameSize(name, intentNameRawLimit)
}

func limitNameSize(name string, limitSize int) string {
	if len(name) > limitSize {
		return name[:limitSize]
	}
	return name
}

const intentNameRawLimit = 57 // floor(<limit>/(8/log2(<baseX>))-1), eg: floor(100/(8/log2(26))-1)

func spaceBoundName(space, name string) string {
	return fmt.Sprint(name, "/", space)
}
