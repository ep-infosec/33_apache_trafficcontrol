package v4

/*

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

import (
	"net/http"
	"testing"

	"github.com/apache/trafficcontrol/traffic_ops/testing/api/utils"
	client "github.com/apache/trafficcontrol/traffic_ops/v4-client"
)

func TestPing(t *testing.T) {

	methodTests := utils.TestCase[client.Session, client.RequestOptions, struct{}]{
		"GET": {
			"OK when VALID request": {
				ClientSession: TOSession,
				Expectations:  utils.CkRequest(utils.NoError(), utils.HasStatus(http.StatusOK)),
			},
			"OK when UNAUTHENTICATED": {
				ClientSession: NoAuthTOSession,
				Expectations:  utils.CkRequest(utils.NoError(), utils.HasStatus(http.StatusOK)),
			},
		},
	}

	for method, testCases := range methodTests {
		t.Run(method, func(t *testing.T) {
			for name, testCase := range testCases {
				switch method {
				case "GET":
					t.Run(name, func(t *testing.T) {
						resp, reqInf, err := testCase.ClientSession.Ping(testCase.RequestOpts)
						for _, check := range testCase.Expectations {
							check(t, reqInf, resp.Ping, resp.Alerts, err)
						}
					})
				}
			}
		})
	}
}
