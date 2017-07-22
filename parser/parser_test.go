package parser

import (
	"strings"
	"testing"
	"github.com/Tlakatlekutl/tkproxy/log"
)

func TestChangeSourceUrl(t *testing.T) {
	log.SetDebugLevel("debug")
	testString :=
		`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8"/>
    <title>Proxy</title>
    <link rel="stylesheet" href="//static/css/index.css"/>
</head>
<body>
    <div class="navigation-panel">
        <div class="inner-panel">
            <input type="text" class="url-input" id="url-input"/>
            <button class="submit" id="submit">Go!</button>
        </div>
    </div>
    <div class="frame" id="frame">
        <h1>Welcome </h1>
<script>alert("ololo")</script>
    </div>

    <script src="static/js/index.js"></script>
</body>
</html>
`
	expectedString :=
		`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8"/>
    <title>Proxy</title>
    <link rel="stylesheet" href="/source?from=http://test.com/static/css/index.css"/>
</head>
<body>
    <div class="navigation-panel">
        <div class="inner-panel">
            <input type="text" class="url-input" id="url-input"/>
            <button class="submit" id="submit">Go!</button>
        </div>
    </div>
    <div class="frame" id="frame">
        <h1>Welcome </h1>
<script>alert("ololo")</script>
    </div>

    <script src="static/js/index.js"></script>
</body>
</html>
`
	reader := strings.NewReader(testString)

	// test function
	out, err := ChangeSourceUrl(reader, "/source?from=http://", "test.com/")

	if err != nil {
		t.Fatalf("Error after func call: %s", err.Error())
	}

	// remove spaces and '\n'
	mapf := func(r rune) rune{
		switch r {
		case ' ', '\n':
			return -1
		}
		return r
	}

	ex := strings.Map(mapf, expectedString)
	o := strings.Map(mapf, out)
	if ex != o {
		t.Fatalf("Error result compare. Got\n\"%+v\" \nexpected\n\"%+v\"", o, ex)
	}
}
