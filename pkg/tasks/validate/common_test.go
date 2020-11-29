package validate

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

var testExtractGUIDItems = []struct {
	comment        string
	urlPath        string
	expectedResult string
	expectedError  string
}{
	{
		comment:        "valid guid in urlPath",
		urlPath:        "/identities/rotate/18896661-e861-47a2-b724-629a07a4c67d",
		expectedResult: "18896661-e861-47a2-b724-629a07a4c67d",
		expectedError:  "",
	},
	{
		comment:        "invalid guid in urlPath",
		urlPath:        "/identities/rotate/18896661-e861-47a2-b724-629",
		expectedResult: "",
		expectedError:  "invalid UUID length: 27",
	},
	{
		comment:        "missing guid in urlPath",
		urlPath:        "/identities/rotate/",
		expectedResult: "",
		expectedError:  "invalid UUID length: 0",
	},
}

func TestExtractGUID(t *testing.T) {

	for _, item := range testExtractGUIDItems {
		result, err := extractGUID(item.urlPath)
		if err != nil {
			assert.Equal(t, item.expectedError, err.Error(), item.comment)
		}
		assert.Equal(t, item.expectedResult, result, item.comment)
	}
}

func TestTimeWithinMinutes(t *testing.T) {

	timestamp := time.Now().UnixNano() / 1e6         // convert to milliseconds
	timestampStr := strconv.FormatInt(timestamp, 10) // convert to string type

	assert.Equal(t, true, timeWithinMinutes(timestampStr, 5), "time is now")
	assert.Equal(t, false, timeWithinMinutes("1604573826351", 5), "time is old")
	assert.Equal(t, false, timeWithinMinutes("", 5), "time is empty")
	assert.Equal(t, false, timeWithinMinutes("16045738263511604573826351", 5), "time is beyond int64")
}
