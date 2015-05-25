package front


import (
    "fmt"
	"html/template"
	"net/http"
	"github.com/gin-gonic/gin"
)

type blogItem struct {
    url string
    title string
}

func PingCtr(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
func HomeCtr(c *gin.Context) {
    var bl [2]blogItem
    bl[0] =  blogItem{"dsfdf", "dsfdsf"}
    bl[1] = blogItem{"sdfdsf", "34234324"}
    fmt.Println(bl);
	blogList := `
	<li>
            <a href="view.php?id=3833">How To Manually Install Oracle Java on a Debian or Ubuntu VPS</a></li>
            <li>
                <a href="view.php?id=3832">Linux 4.0 kernel released</a></li>
            <li>
                <a href="view.php?id=3824">CentOS内网升级</a></li>
            <li>
                <a href="view.php?id=3823">DDOS attacks on github improved the experience for us</a></li>
            <li>
                <a href="view.php?id=3822">Erlang what i know as a beginner</a></li>
            <li>
                <a href="view.php?id=3821">Configure proxy for android repo sync</a></li>
            <li>
                <a href="view.php?id=3819">PHP going to death, No one care some one leaved</a></li>
            <li>
                <a href="view.php?id=3818">Install ghc 7.8.4 in linux by youself</a></li>
            <li>
                <a href="view.php?id=3816">Ansible sync deploy files to remote server with sudo</a></li>
            <li>
                <a href="view.php?id=3815">CoreOS run docker container using systemd boot up</a></li>
            <li>
                <a href="view.php?id=3814">CoreOS sshd security configure guide</a></li>
            <li>
                <a href="view.php?id=3813">SSH connect server over gateway server via ProxyCommand</a></li>
            <li>
                <a href="view.php?id=3812">ping icmp open socket operation not permitted</a></li>
            <li>
                <a href="view.php?id=3811">Getting start Ansible for fun, Ansible quickstart</a></li>
            <li>
                <a href="view.php?id=3810">Today, i follow Watanabe Mayu, See after 10, 20, 30, 40, 50, 60, 70 year , are you still love mayuyu</a></li>
            <li>
                <a href="view.php?id=3809">Play GNU Assembler on Fedora 21 x64</a></li>
            <li>
                <a href="view.php?id=3808">How to use $routeParams in the AngularJS controller</a></li>
            <li>
                <a href="view.php?id=3807">AngularJS dynamic add event bind with data repeat </a></li>
            <li>
                <a href="view.php?id=3806">AngularJS route jump</a></li>
            <li>
                <a href="view.php?id=3805">AngularJS $http.post problem solved</a></li>`
	c.HTML(http.StatusOK, "index.html", gin.H{"bloglist":template.HTML(blogList)})
}