// ==UserScript==
// @name         freeflix IMDB userscript
// @namespace    http://tampermonkey.net/
// @version      0.1
// @description  freeflix
// @author       You
// @match        https://www.imdb.com/title/*
// @icon         https://www.google.com/s2/favicons?sz=64&domain=tampermonkey.net
// @require      http://code.jquery.com/jquery-3.4.1.min.js
// @grant        none
// ==/UserScript==

(function($) {
  'use strict';
  const host = "example.com"
  const name = $('h1 span').html()
  $('h1').append($(`<a href="${host}/?query=${name}">+</a>`))
})($);