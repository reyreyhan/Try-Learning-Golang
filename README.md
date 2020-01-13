# Try-Learning-Golang Using Echo Framework

<div class="Box-body">
        <article class="markdown-body entry-content p-5" itemprop="text"><p><a href="https://echo.labstack.com" rel="nofollow"><img src="https://camo.githubusercontent.com/5e68fe65866b89b9e6dc9975ddef1339c92b7534/68747470733a2f2f63646e2e6c6162737461636b2e636f6d2f696d616765732f6563686f2d6c6f676f2e737667" data-canonical-src="https://cdn.labstack.com/images/echo-logo.svg" style="max-width:100%;" height="80"></a></p>
<p><a href="https://sourcegraph.com/github.com/labstack/echo?badge" rel="nofollow"><img src="https://camo.githubusercontent.com/2a6f00d3dbd0296966b78c41abec23a17ed84624/68747470733a2f2f736f7572636567726170682e636f6d2f6769746875622e636f6d2f6c6162737461636b2f6563686f2f2d2f62616467652e7376673f7374796c653d666c61742d737175617265" alt="Sourcegraph" data-canonical-src="https://sourcegraph.com/github.com/labstack/echo/-/badge.svg?style=flat-square" style="max-width:100%;"></a>
<a href="http://godoc.org/github.com/labstack/echo" rel="nofollow"><img src="https://camo.githubusercontent.com/6609a4ae07da7917f7a7b79c45ccc24fcc01540f/687474703a2f2f696d672e736869656c64732e696f2f62616467652f676f2d646f63756d656e746174696f6e2d626c75652e7376673f7374796c653d666c61742d737175617265" alt="GoDoc" data-canonical-src="http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square" style="max-width:100%;"></a>
<a href="https://goreportcard.com/report/github.com/labstack/echo" rel="nofollow"><img src="https://camo.githubusercontent.com/a5371e54028b3d5820942fb15ebde12a3c3a73c0/68747470733a2f2f676f7265706f7274636172642e636f6d2f62616467652f6769746875622e636f6d2f6c6162737461636b2f6563686f3f7374796c653d666c61742d737175617265" alt="Go Report Card" data-canonical-src="https://goreportcard.com/badge/github.com/labstack/echo?style=flat-square" style="max-width:100%;"></a>
<a href="https://travis-ci.org/labstack/echo" rel="nofollow"><img src="https://camo.githubusercontent.com/0fff04c890e22b9a3c4cc531c6c3f78de36d3b6d/687474703a2f2f696d672e736869656c64732e696f2f7472617669732f6c6162737461636b2f6563686f2e7376673f7374796c653d666c61742d737175617265" alt="Build Status" data-canonical-src="http://img.shields.io/travis/labstack/echo.svg?style=flat-square" style="max-width:100%;"></a>
<a href="https://codecov.io/gh/labstack/echo" rel="nofollow"><img src="https://camo.githubusercontent.com/b3a3d6663158a32c760f84d270e702ce9ee2bf39/68747470733a2f2f696d672e736869656c64732e696f2f636f6465636f762f632f6769746875622f6c6162737461636b2f6563686f2e7376673f7374796c653d666c61742d737175617265" alt="Codecov" data-canonical-src="https://img.shields.io/codecov/c/github/labstack/echo.svg?style=flat-square" style="max-width:100%;"></a>
<a href="https://gitter.im/labstack/echo" rel="nofollow"><img src="https://camo.githubusercontent.com/8f09b2633c69ba6c3297c774156511bcaca270fe/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f6769747465722d6a6f696e253230636861742d627269676874677265656e2e7376673f7374796c653d666c61742d737175617265" alt="Join the chat at https://gitter.im/labstack/echo" data-canonical-src="https://img.shields.io/badge/gitter-join%20chat-brightgreen.svg?style=flat-square" style="max-width:100%;"></a>
<a href="https://forum.labstack.com" rel="nofollow"><img src="https://camo.githubusercontent.com/aed33c8881cd07a27f62ee4f1379d180b2edf2f8/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f636f6d6d756e6974792d666f72756d2d3030616664312e7376673f7374796c653d666c61742d737175617265" alt="Forum" data-canonical-src="https://img.shields.io/badge/community-forum-00afd1.svg?style=flat-square" style="max-width:100%;"></a>
<a href="https://twitter.com/labstack" rel="nofollow"><img src="https://camo.githubusercontent.com/d5b1f89f574d6c84c6a4d39ba073101ad06a65aa/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f747769747465722d406c6162737461636b2d3535616365652e7376673f7374796c653d666c61742d737175617265" alt="Twitter" data-canonical-src="https://img.shields.io/badge/twitter-@labstack-55acee.svg?style=flat-square" style="max-width:100%;"></a>
<a href="https://raw.githubusercontent.com/labstack/echo/master/LICENSE" rel="nofollow"><img src="https://camo.githubusercontent.com/c0e9b3d3e996ee6c01364035633ead82feb4a729/687474703a2f2f696d672e736869656c64732e696f2f62616467652f6c6963656e73652d6d69742d626c75652e7376673f7374796c653d666c61742d737175617265" alt="License" data-canonical-src="http://img.shields.io/badge/license-mit-blue.svg?style=flat-square" style="max-width:100%;"></a></p>
<h2><a id="user-content-supported-go-versions" class="anchor" aria-hidden="true" href="#supported-go-versions"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>Supported Go versions</h2>
<p>As of version 4.0.0, Echo is available as a <a href="https://github.com/golang/go/wiki/Modules">Go module</a>.
Therefore a Go version capable of understanding /vN suffixed imports is required:</p>
<ul>
<li>1.9.7+</li>
<li>1.10.3+</li>
<li>1.11+</li>
</ul>
<p>Any of these versions will allow you to import Echo as <code>github.com/labstack/echo/v4</code> which is the recommended
way of using Echo going forward.</p>
<p>For older versions, please use the latest v3 tag.</p>
<h2><a id="user-content-feature-overview" class="anchor" aria-hidden="true" href="#feature-overview"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>Feature Overview</h2>
<ul>
<li>Optimized HTTP router which smartly prioritize routes</li>
<li>Build robust and scalable RESTful APIs</li>
<li>Group APIs</li>
<li>Extensible middleware framework</li>
<li>Define middleware at root, group or route level</li>
<li>Data binding for JSON, XML and form payload</li>
<li>Handy functions to send variety of HTTP responses</li>
<li>Centralized HTTP error handling</li>
<li>Template rendering with any template engine</li>
<li>Define your format for the logger</li>
<li>Highly customizable</li>
<li>Automatic TLS via Let’s Encrypt</li>
<li>HTTP/2 support</li>
</ul>
<h2><a id="user-content-benchmarks" class="anchor" aria-hidden="true" href="#benchmarks"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>Benchmarks</h2>
<p>Date: 2018/03/15<br>
Source: <a href="https://github.com/vishr/web-framework-benchmark">https://github.com/vishr/web-framework-benchmark</a><br>
Lower is better!</p>
<p><a target="_blank" rel="noopener noreferrer" href="https://camo.githubusercontent.com/d8800e2ee37115207efc1f3e937a28fb49d90e22/68747470733a2f2f692e696d6775722e636f6d2f49333256644d4a2e706e67"><img src="https://camo.githubusercontent.com/d8800e2ee37115207efc1f3e937a28fb49d90e22/68747470733a2f2f692e696d6775722e636f6d2f49333256644d4a2e706e67" data-canonical-src="https://i.imgur.com/I32VdMJ.png" style="max-width:100%;"></a></p>
<h2><a id="user-content-guide" class="anchor" aria-hidden="true" href="#guide"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a><a href="https://echo.labstack.com/guide" rel="nofollow">Guide</a></h2>
<h3><a id="user-content-example" class="anchor" aria-hidden="true" href="#example"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>Example</h3>
<div class="highlight highlight-source-go"><pre><span class="pl-k">package</span> main

<span class="pl-k">import</span> (
  <span class="pl-s"><span class="pl-pds">"</span>net/http<span class="pl-pds">"</span></span>
  <span class="pl-s"><span class="pl-pds">"</span>github.com/labstack/echo/v4<span class="pl-pds">"</span></span>
  <span class="pl-s"><span class="pl-pds">"</span>github.com/labstack/echo/v4/middleware<span class="pl-pds">"</span></span>
)

<span class="pl-k">func</span> <span class="pl-en">main</span>() {
  <span class="pl-c"><span class="pl-c">//</span> Echo instance</span>
  <span class="pl-smi">e</span> <span class="pl-k">:=</span> echo.<span class="pl-c1">New</span>()

  <span class="pl-c"><span class="pl-c">//</span> Middleware</span>
  e.<span class="pl-c1">Use</span>(middleware.<span class="pl-c1">Logger</span>())
  e.<span class="pl-c1">Use</span>(middleware.<span class="pl-c1">Recover</span>())

  <span class="pl-c"><span class="pl-c">//</span> Routes</span>
  e.<span class="pl-c1">GET</span>(<span class="pl-s"><span class="pl-pds">"</span>/<span class="pl-pds">"</span></span>, hello)

  <span class="pl-c"><span class="pl-c">//</span> Start server</span>
  e.<span class="pl-smi">Logger</span>.<span class="pl-c1">Fatal</span>(e.<span class="pl-c1">Start</span>(<span class="pl-s"><span class="pl-pds">"</span>:1323<span class="pl-pds">"</span></span>))
}

<span class="pl-c"><span class="pl-c">//</span> Handler</span>
<span class="pl-k">func</span> <span class="pl-en">hello</span>(<span class="pl-v">c</span> <span class="pl-v">echo</span>.<span class="pl-v">Context</span>) <span class="pl-v">error</span> {
  <span class="pl-k">return</span> c.<span class="pl-c1">String</span>(http.<span class="pl-smi">StatusOK</span>, <span class="pl-s"><span class="pl-pds">"</span>Hello, World!<span class="pl-pds">"</span></span>)
}</pre></div>
<h2><a id="user-content-help" class="anchor" aria-hidden="true" href="#help"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>Help</h2>
<ul>
<li><a href="https://forum.labstack.com" rel="nofollow">Forum</a></li>
<li><a href="https://gitter.im/labstack/echo" rel="nofollow">Chat</a></li>
</ul>
<h2><a id="user-content-contribute" class="anchor" aria-hidden="true" href="#contribute"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>Contribute</h2>
<p><strong>Use issues for everything</strong></p>
<ul>
<li>For a small change, just send a PR.</li>
<li>For bigger changes open an issue for discussion before sending a PR.</li>
<li>PR should have:
<ul>
<li>Test case</li>
<li>Documentation</li>
<li>Example (If it makes sense)</li>
</ul>
</li>
<li>You can also contribute by:
<ul>
<li>Reporting issues</li>
<li>Suggesting new features or enhancements</li>
<li>Improve/fix documentation</li>
</ul>
</li>
</ul>
<h2><a id="user-content-credits" class="anchor" aria-hidden="true" href="#credits"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>Credits</h2>
<ul>
<li><a href="https://github.com/vishr">Vishal Rana</a> - Author</li>
<li><a href="https://github.com/nr17">Nitin Rana</a> - Consultant</li>
<li><a href="https://github.com/labstack/echo/graphs/contributors">Contributors</a></li>
</ul>
<h2><a id="user-content-license" class="anchor" aria-hidden="true" href="#license"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path fill-rule="evenodd" d="M4 9h1v1H4c-1.5 0-3-1.69-3-3.5S2.55 3 4 3h4c1.45 0 3 1.69 3 3.5 0 1.41-.91 2.72-2 3.25V8.59c.58-.45 1-1.27 1-2.09C10 5.22 8.98 4 8 4H4c-.98 0-2 1.22-2 2.5S3 9 4 9zm9-3h-1v1h1c1 0 2 1.22 2 2.5S13.98 12 13 12H9c-.98 0-2-1.22-2-2.5 0-.83.42-1.64 1-2.09V6.25c-1.09.53-2 1.84-2 3.25C6 11.31 7.55 13 9 13h4c1.45 0 3-1.69 3-3.5S14.5 6 13 6z"></path></svg></a>License</h2>
<p><a href="https://github.com/labstack/echo/blob/master/LICENSE">MIT</a></p>
</article>
      </div>
