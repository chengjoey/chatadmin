(function(e){function t(t){for(var o,s,i=t[0],c=t[1],u=t[2],l=0,h=[];l<i.length;l++)s=i[l],Object.prototype.hasOwnProperty.call(r,s)&&r[s]&&h.push(r[s][0]),r[s]=0;for(o in c)Object.prototype.hasOwnProperty.call(c,o)&&(e[o]=c[o]);f&&f(t);while(h.length)h.shift()();return a.push.apply(a,u||[]),n()}function n(){for(var e,t=0;t<a.length;t++){for(var n=a[t],o=!0,s=1;s<n.length;s++){var i=n[s];0!==r[i]&&(o=!1)}o&&(a.splice(t--,1),e=c(c.s=n[0]))}return e}var o={},s={app:0},r={app:0},a=[];function i(e){return c.p+"js/"+({}[e]||e)+"."+{"chunk-5cbc54d8":"6fc523b9","chunk-63615f18":"08160f20"}[e]+".js"}function c(t){if(o[t])return o[t].exports;var n=o[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,c),n.l=!0,n.exports}c.e=function(e){var t=[],n={"chunk-5cbc54d8":1,"chunk-63615f18":1};s[e]?t.push(s[e]):0!==s[e]&&n[e]&&t.push(s[e]=new Promise((function(t,n){for(var o="css/"+({}[e]||e)+"."+{"chunk-5cbc54d8":"ab41a107","chunk-63615f18":"93758af8"}[e]+".css",r=c.p+o,a=document.getElementsByTagName("link"),i=0;i<a.length;i++){var u=a[i],l=u.getAttribute("data-href")||u.getAttribute("href");if("stylesheet"===u.rel&&(l===o||l===r))return t()}var h=document.getElementsByTagName("style");for(i=0;i<h.length;i++){u=h[i],l=u.getAttribute("data-href");if(l===o||l===r)return t()}var f=document.createElement("link");f.rel="stylesheet",f.type="text/css",f.onload=t,f.onerror=function(t){var o=t&&t.target&&t.target.src||r,a=new Error("Loading CSS chunk "+e+" failed.\n("+o+")");a.code="CSS_CHUNK_LOAD_FAILED",a.request=o,delete s[e],f.parentNode.removeChild(f),n(a)},f.href=r;var d=document.getElementsByTagName("head")[0];d.appendChild(f)})).then((function(){s[e]=0})));var o=r[e];if(0!==o)if(o)t.push(o[2]);else{var a=new Promise((function(t,n){o=r[e]=[t,n]}));t.push(o[2]=a);var u,l=document.createElement("script");l.charset="utf-8",l.timeout=120,c.nc&&l.setAttribute("nonce",c.nc),l.src=i(e);var h=new Error;u=function(t){l.onerror=l.onload=null,clearTimeout(f);var n=r[e];if(0!==n){if(n){var o=t&&("load"===t.type?"missing":t.type),s=t&&t.target&&t.target.src;h.message="Loading chunk "+e+" failed.\n("+o+": "+s+")",h.name="ChunkLoadError",h.type=o,h.request=s,n[1](h)}r[e]=void 0}};var f=setTimeout((function(){u({type:"timeout",target:l})}),12e4);l.onerror=l.onload=u,document.head.appendChild(l)}return Promise.all(t)},c.m=e,c.c=o,c.d=function(e,t,n){c.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},c.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},c.t=function(e,t){if(1&t&&(e=c(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(c.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var o in e)c.d(n,o,function(t){return e[t]}.bind(null,o));return n},c.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return c.d(t,"a",t),t},c.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},c.p="/",c.oe=function(e){throw console.error(e),e};var u=window["webpackJsonp"]=window["webpackJsonp"]||[],l=u.push.bind(u);u.push=t,u=u.slice();for(var h=0;h<u.length;h++)t(u[h]);var f=l;a.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("56d7")},"00c2":function(e,t,n){"use strict";var o="token",s="role",r=function(){return window.localStorage.getItem(o)},a=function(){return window.localStorage.getItem(s)},i=function(e){window.localStorage.setItem(o,e)},c=function(e){window.localStorage.setItem(s,e)},u=function(){window.localStorage.removeItem(o),window.localStorage.removeItem(s)};t["a"]={getToken:r,saveToken:i,destroyToken:u,setRole:c,getRole:a}},1:function(e,t){},2:function(e,t){},"56d7":function(e,t,n){"use strict";n.r(t);n("b0c0"),n("e260"),n("e6cf"),n("cca6"),n("a79d");var o=n("2b0e"),s=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-app",{attrs:{id:"inspire"}},[n("router-view")],1)},r=[],a=n("5530"),i=n("2f62"),c={name:"App",computed:Object(a["a"])({},Object(i["b"])(["isAuthenticated"])),created:function(){var e=this.$router.history.current.path;!1===this.isAuthenticated&&this.$router.push("/login"),"/"!==e&&"/app"!==e||this.$router.push("/app/dashboard")}},u=c,l=n("2877"),h=n("6544"),f=n.n(h),d=n("7496"),m=Object(l["a"])(u,s,r,!1,null,null,null),g=m.exports;f()(m,{VApp:d["a"]});var p=n("f309");o["default"].use(p["a"]);var v=new p["a"]({}),b=(n("d3b7"),n("8c4f")),k=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-container",[n("v-row",[n("v-col",{attrs:{cols:"12"}},[n("strong",[e._v("在线用户")])])],1),n("v-row",[e._l(e.line,(function(t){return e._l(e.users.slice(4*(t-1),4*t),(function(o,s){return n("v-col",{key:""+t+s,attrs:{cols:"6",md:"2"}},[n("v-card",{staticClass:"mx-auto",attrs:{"max-width":"344",outlined:""}},[n("v-list-item",{attrs:{"three-line":""}},[n("v-list-item-content",[n("v-list-item-title",{staticClass:"headline mb-1"},[e._v(" "+e._s(o.nickname)+" ")]),o.is_chated?n("v-list-item-subtitle",[e._v("已在聊天")]):n("v-list-item-subtitle",[e._v("待聊天")]),n("v-list-item-subtitle",[e._v("加入时间: "+e._s(o.enter_at))])],1)],1),n("v-card-actions",[n("v-btn",{attrs:{outlined:"",rounded:"",text:"",disabled:o.is_chated},on:{click:function(t){return e.chat(o.nickname)}}},[e._v(" 点击聊天 ")])],1)],1)],1)}))}))],2)],1)},w=[],y="http://"+location.hostname+":"+location.port+"/",_="ws://"+location.hostname+":"+location.port+"/v1/chat",O=n("bc3a"),S=n.n(O),j=n("2106"),x=n.n(j),T=n("00c2"),C={init:function(){o["default"].use(x.a,S.a),o["default"].axios.defaults.baseURL=y},getIDToken:function(){return T["a"].getToken()},setHeader:function(){o["default"].axios.defaults.headers.common["Authorization"]=T["a"].getToken()},post:function(e,t){return o["default"].axios.post("".concat(e),t)},get:function(e){var t={headers:{Authorization:T["a"].getToken()}};return o["default"].axios.get("".concat(e),t)},put:function(e,t){return o["default"].axios.put("{resource}",t)}},A=C,I=n("24b5"),M=n.n(I),U={data:function(){return{users:[],line:0,evtSource:null}},methods:{createSSE:function(){var e=M.a;this.evtSource=new e(y+"v1/stream",{headers:{Authorization:A.getIDToken()}}),this.evtSource.onmessage=this.sseOnMessage},getUsers:function(){var e=this;A.get("v1/users").then((function(t){e.users=t.data.value,e.line=Math.ceil(e.users.length/4)})).catch((function(e){console.log(e)}))},chat:function(e){this.$router.push({paht:"/app/chat",name:"chat",params:{client:e}})},sseOnMessage:function(e){"null"===e.data?(this.users=[],this.line=0):(this.users=JSON.parse(e.data),this.line=Math.ceil(this.users.length/4)),console.log(this.users,this.line)}},mounted:function(){this.getUsers(),this.createSSE()}},E=U,L=n("8336"),V=n("b0af"),P=n("99d9"),D=n("62ad"),$=n("a523"),N=n("da13"),R=n("5d23"),J=n("0fd9"),B=Object(l["a"])(E,k,w,!1,null,null,null),H=B.exports;f()(B,{VBtn:L["a"],VCard:V["a"],VCardActions:P["a"],VCol:D["a"],VContainer:$["a"],VListItem:N["a"],VListItemContent:R["a"],VListItemSubtitle:R["b"],VListItemTitle:R["c"],VRow:J["a"]});var z=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-container",{staticClass:"fill-height"},[n("section",{ref:"chatArea",staticClass:"chat-area"},[e._l(e.historyMsgList,(function(t,o){return n("p",{key:"hsmsg-"+o,staticClass:"message",class:{"message-out":t.user.nickname===e.currentUser.nickname,"message-in":t.user.nickname!==e.currentUser.nickname}},[e._v(" "+e._s(t.user.nickname)+" at "+e._s(e.formatDate(t.msg_time))+" "),n("br"),n("span",[e._v(e._s(t.content))])])})),n("br"),n("p",{attrs:{align:"center"}},[e._v("以上为历史消息")]),n("v-divider"),n("br"),e._l(e.msglist,(function(t,o){return n("p",{key:"msg-"+o,staticClass:"message",class:{"message-out":t.user.nickname===e.currentUser.nickname,"message-in":t.user.nickname!==e.currentUser.nickname}},[e._v(" "+e._s(t.user.nickname)+" at "+e._s(e.formatDate(t.msg_time))+" "),n("br"),n("span",[e._v(e._s(t.content))])])}))],2),n("v-col",{attrs:{cols:"12"}},[n("v-text-field",{attrs:{"append-outer-icon":"mdi-send",filled:"","clear-icon":"mdi-close-circle",clearable:"",label:"Message",type:"text"},on:{"click:append":e.toggleMarker,"click:append-outer":e.sendMessage,"click:clear":e.clearMessage},model:{value:e.content,callback:function(t){e.content=t},expression:"content"}})],1)],1)},F=[],q={data:function(){return{websock:null,msglist:[],historyMsgList:[],content:"Hey!",client:"",show:!1,message:"Hey!",marker:!0}},methods:{joinChat:function(){var e=encodeURIComponent(A.getIDToken()),t=_+"?token="+e+"&client="+this.client;console.log(t),this.websock=new WebSocket(t),this.websock.onmessage=this.websocketonmessage,this.websock.onopen=this.websocketonopen,this.websock.onerror=this.websocketonerror,this.websock.onclose=this.websocketonclose},websocketonmessage:function(e){var t=JSON.parse(e.data);t.receive_time=new Date,this.msglist.push(t)},websocketonopen:function(){console.log("websocket已建立")},websocketonerror:function(){console.log("websocket断开链接")},websocketonsend:function(){console.log(this.currentUser.nickname);var e=JSON.stringify({content:this.content});this.websock.send(e);var t={user:{nickname:this.currentUser.nickname,uid:this.currentUser.uid},type:0,content:this.content,msg_time:(new Date).getTime()};this.msglist.push(t),this.content=""},websocketonclose:function(e){console.log("断开链接",e)},closewebsocket:function(){this.websock.close()},lineFeed:function(e){console.log(e),this.content=this.content+"\n"},formatDate:function(e){var t=new Date(e);return t.toLocaleString()},toggleMarker:function(){this.marker=!this.marker},sendMessage:function(){this.websocketonsend(),this.clearMessage()},clearMessage:function(){this.message=""},resetIcon:function(){this.iconIndex=0},changeIcon:function(){this.iconIndex===this.icons.length-1?this.iconIndex=0:this.iconIndex++},getMsgs:function(){var e=this,t="v1/msgs?client="+this.client;A.get(t).then((function(t){e.historyMsgList=t.data.value})).catch((function(e){console.log(e)}))},getParams:function(){this.client=this.$route.params.client},printUser:function(){console.log(this.currentUser)}},mounted:function(){this.printUser(),this.joinChat()},computed:Object(a["a"])({},Object(i["b"])(["currentUser"])),destroyed:function(){this.websock.close()},created:function(){this.getParams(),this.getMsgs()},beforeRouteLeave:function(e,t,n){this.websock.close(),n()}},K=q,W=(n("b567"),n("ce7e")),G=n("8654"),Q=Object(l["a"])(K,z,F,!1,null,"60ad22f0",null),X=Q.exports;f()(Q,{VCol:D["a"],VContainer:$["a"],VDivider:W["a"],VTextField:G["a"]}),o["default"].use(b["a"]);var Y,Z,ee=new b["a"]({routes:[{path:"/login",component:function(){return n.e("chunk-63615f18").then(n.bind(null,"0206"))},name:"login"},{path:"/app",component:function(){return n.e("chunk-5cbc54d8").then(n.bind(null,"ebad"))},name:"layout",children:[{path:"dashboard",name:"dashboard",component:H},{path:"chat",component:X,name:"chat"}]}]}),te=n("ade3"),ne=n("6c33"),oe="setUser",se="setError",re="logOut",ae={errors:null,user:{},isAuthenticated:!!T["a"].getToken()},ie={currentUser:function(e){return e.user},isAuthenticated:function(e){return e.isAuthenticated},isAdmin:function(e){return"admin"===e.user.role}},ce=(Y={},Object(te["a"])(Y,ne["b"],(function(e,t){return new Promise((function(n){A.post("login",{nickname:t}).then((function(t){t.data.value.token=t.headers.token,e.commit(oe,t.data.value),n(t.data)})).catch((function(t){console.log(t),e.commit(se,t.data.msg)}))}))})),Object(te["a"])(Y,ne["a"],(function(e){T["a"].getToken()?console.log("already login"):e.commit(re)})),Object(te["a"])(Y,ne["c"],(function(e){e.commit(re)})),Y),ue=(Z={},Object(te["a"])(Z,se,(function(e,t){e.errors=t})),Object(te["a"])(Z,oe,(function(e,t){e.isAuthenticated=!0,e.user=t,e.errors={},T["a"].saveToken(t.token),T["a"].setRole(t.role),console.log(t.role)})),Object(te["a"])(Z,re,(function(e){e.isAuthenticated=!1,e.user={},e.errors={},T["a"].destroyToken()})),Z),le={state:ae,actions:ce,mutations:ue,getters:ie};o["default"].use(i["a"]);var he=new i["a"].Store({modules:{auth:le}}),fe=n("37b6"),de=n("91b2"),me=n.n(de),ge=n("e54b"),pe=n.n(ge);o["default"].use(pe.a),o["default"].config.productionTip=!1,o["default"].use(fe["a"]),o["default"].use(me.a),A.init(),ee.beforeEach((function(e,t,n){"login"===e.name||T["a"].getToken()?n():n({path:"/login"})})),new o["default"]({vuetify:v,router:ee,store:he,render:function(e){return e(g)}}).$mount("#app")},"6c33":function(e,t,n){"use strict";n.d(t,"b",(function(){return o})),n.d(t,"c",(function(){return s})),n.d(t,"a",(function(){return r}));var o="login",s="logout",r="checkAuth"},b567:function(e,t,n){"use strict";n("f3e8")},f3e8:function(e,t,n){}});
//# sourceMappingURL=app.ef6686f1.js.map