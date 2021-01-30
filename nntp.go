/*
   Copyright 2021 MediaExchange.io

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

package nntp

import "net/textproto"

func DownloadArticles(hostname string, username string, password string, articleIds []string) (err error){
	// Connect to server
	conn, err := textproto.Dial("tcp", hostname + ":119")
	if err != nil {
		return
	}

	// 200 Welcome to XS Usenet
	if _, _, err = conn.ReadCodeLine(200); err != nil {
		return
	}

	// Send username
	_, err = conn.Cmd("AUTHINFO USER", username)
	if err != nil {
		return
	}

	// 381 Need more.
	if _, _, err = conn.ReadCodeLine(381); err != nil {
		return
	}

	// Send password
	_, err = conn.Cmd("AUTHINFO PASS", password)
	if err != nil {
		return
	}

	// 281 Authentication accepted.
	if _, _, err = conn.ReadCodeLine(281); err != nil {
		return
	}

	_, err = conn.Cmd("QUIT")
	if err != nil {
		return
	}

	// 205 Bye. 8386623 bytes written, 22 accounted.
	if _, _, err = conn.ReadCodeLine(205); err != nil {
		return
	}

	return
}
