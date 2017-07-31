package soupchain

import (
	"fmt"
	"testing"

	"github.com/anaskhan96/soup"
	soupchain "github.com/leeola/soup-chain"
)

const testHTML = `
<html>
  <head>
    <title>Sample "Hello, World" Application</title>
  </head>
  <body bgcolor=white>
    <table border="0" cellpadding="10">
      <tr>
        <td>
          <img src="images/springsource.png">
        </td>
        <td>
          <h1>Sample "Hello, World" Application</h1>
        </td>
      </tr>
    </table>
    <div id="0">
      <div id="1">Just two divs peacing out</div>
      <div id="2">Second div peacing out</div>
    </div>
    check
    <div id="2">One more</div>
    <p>This is the home page for the HelloWorld Web application. </p>
    <p>To prove that they work, you can execute either of the following links:
    <ul>
      <li>To a <a href="hello.jsp">JSP page</a></li>
      <li>To a <a href="hello">servlet</a></li>
    </ul>
    </p>
    <div id="3">
      <div id="4">Last one</div>
    </div>
  </body>
</html>
`

func TestFind(t *testing.T) {
	doc := soupchain.New(soup.HTMLParse(testHTML))

	vs, _ := doc.Find("div").FindAll("div")
	for _, val := range vs {
		v, err := val.Text()
		fmt.Printf("value? %#v\n", v)
		fmt.Println("err:", err)
	}

	// // Find() and Attrs()
	// actual := doc.Find("img").Attrs()["src"]
	// if !reflect.DeepEqual(actual, "images/springsource.png") {
	// 	t.Error("Instead of `images/springsource.png`, got", actual)
	// }
	// // Find(...) and Text()
	// actual = doc.Find("a", "href", "hello").Text()
	// if !reflect.DeepEqual(actual, "servlet") {
	// 	t.Error("Instead of `servlet`, got", actual)
	// }
	// // Nested Find()
	// actual = doc.Find("div").Find("div").Text()
	// if !reflect.DeepEqual(actual, "Just two divs peacing out") {
	// 	t.Error("Instead of `Just two divs peacing out`, got", actual)
	// }

	t.Fail()
}
