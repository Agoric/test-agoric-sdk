/*---
description: https://github.com/web-platform-tests/wpt/url/resources/urltestdata.json
flags: [module]
---*/

import { runTests } from "./url_FIXTURE.js";

const tests = [
	{
		"input": "#β",
		"base": "http://example.org/foo/bar",
		"href": "http://example.org/foo/bar#%CE%B2",
		"origin": "http://example.org",
		"protocol": "http:",
		"username": "",
		"password": "",
		"host": "example.org",
		"hostname": "example.org",
		"port": "",
		"pathname": "/foo/bar",
		"search": "",
		"hash": "#%CE%B2"
	},
	{
		"input": "data:text/html,test#test",
		"base": "http://example.org/foo/bar",
		"href": "data:text/html,test#test",
		"origin": "null",
		"protocol": "data:",
		"username": "",
		"password": "",
		"host": "",
		"hostname": "",
		"port": "",
		"pathname": "text/html,test",
		"search": "",
		"hash": "#test"
	},
	{
		"input": "tel:1234567890",
		"base": "http://example.org/foo/bar",
		"href": "tel:1234567890",
		"origin": "null",
		"protocol": "tel:",
		"username": "",
		"password": "",
		"host": "",
		"hostname": "",
		"port": "",
		"pathname": "1234567890",
		"search": "",
		"hash": ""
	},
	"# Based on https://felixfbecker.github.io/whatwg-url-custom-host-repro/",
	{
		"input": "ssh://example.com/foo/bar.git",
		"base": "http://example.org/",
		"href": "ssh://example.com/foo/bar.git",
		"origin": "null",
		"protocol": "ssh:",
		"username": "",
		"password": "",
		"host": "example.com",
		"hostname": "example.com",
		"port": "",
		"pathname": "/foo/bar.git",
		"search": "",
		"hash": ""
	},
	"# Based on http://trac.webkit.org/browser/trunk/LayoutTests/fast/url/file.html",
	{
		"input": "file:c:\\foo\\bar.html",
		"base": "file:///tmp/mock/path",
		"href": "file:///c:/foo/bar.html",
		"protocol": "file:",
		"username": "",
		"password": "",
		"host": "",
		"hostname": "",
		"port": "",
		"pathname": "/c:/foo/bar.html",
		"search": "",
		"hash": ""
	},
	{
		"input": "  File:c|////foo\\bar.html",
		"base": "file:///tmp/mock/path",
		"href": "file:///c:////foo/bar.html",
		"protocol": "file:",
		"username": "",
		"password": "",
		"host": "",
		"hostname": "",
		"port": "",
		"pathname": "/c:////foo/bar.html",
		"search": "",
		"hash": ""
	},
	{
		"input": "C|/foo/bar",
		"base": "file:///tmp/mock/path",
		"href": "file:///C:/foo/bar",
		"protocol": "file:",
		"username": "",
		"password": "",
		"host": "",
		"hostname": "",
		"port": "",
		"pathname": "/C:/foo/bar",
		"search": "",
		"hash": ""
	},
	{
		"input": "/C|\\foo\\bar",
		"base": "file:///tmp/mock/path",
		"href": "file:///C:/foo/bar",
		"protocol": "file:",
		"username": "",
		"password": "",
		"host": "",
		"hostname": "",
		"port": "",
		"pathname": "/C:/foo/bar",
		"search": "",
		"hash": ""
	},
	{
		"input": "//C|/foo/bar",
		"base": "file:///tmp/mock/path",
		"href": "file:///C:/foo/bar",
		"protocol": "file:",
		"username": "",
		"password": "",
		"host": "",
		"hostname": "",
		"port": "",
		"pathname": "/C:/foo/bar",
		"search": "",
		"hash": ""
	},
	{
		"input": "//server/file",
		"base": "file:///tmp/mock/path",
		"href": "file://server/file",
		"protocol": "file:",
		"username": "",
		"password": "",
		"host": "server",
		"hostname": "server",
		"port": "",
		"pathname": "/file",
		"search": "",
		"hash": ""
	},
	{
		"input": "\\\\server\\file",
		"base": "file:///tmp/mock/path",
		"href": "file://server/file",
		"protocol": "file:",
		"username": "",
		"password": "",
		"host": "server",
		"hostname": "server",
		"port": "",
		"pathname": "/file",
		"search": "",
		"hash": ""
	},
	{
		"input": "/\\server/file",
		"base": "file:///tmp/mock/path",
		"href": "file://server/file",
		"protocol": "file:",
		"username": "",
		"password": "",
		"host": "server",
		"hostname": "server",
		"port": "",
		"pathname": "/file",
		"search": "",
		"hash": ""
	},
	{
		"input": "file:///foo/bar.txt",
		"base": "file:///tmp/mock/path",
		"href": "file:///foo/bar.txt",
		"protocol": "file:",
		"username": "",
		"password": "",
		"host": "",
		"hostname": "",
		"port": "",
		"pathname": "/foo/bar.txt",
		"search": "",
		"hash": ""
	}
];

runTests(tests);
