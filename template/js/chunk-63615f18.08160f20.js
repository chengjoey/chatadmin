(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-63615f18"],{"0206":function(t,n,e){"use strict";e.r(n);var i=function(){var t=this,n=t.$createElement,e=t._self._c||n;return e("v-card",{staticClass:"mx-auto"},[e("v-card-title",{staticClass:"title font-weight-regular justify-space-between"},[e("span",[t._v("Login")])]),e("v-window",[e("v-window-item",[e("v-card-text",[e("v-text-field",{attrs:{label:"Nickname"},model:{value:t.nickname,callback:function(n){t.nickname=n},expression:"nickname"}}),e("span",{staticClass:"caption grey--text text--darken-1"},[t._v(" admin, adminstrator为管理员用户，其它为普通用户 ")])],1)],1)],1),e("v-divider"),e("v-card-actions",[e("v-spacer"),e("v-btn",{attrs:{color:"primary",depressed:""},on:{click:t.login}},[t._v(" login ")])],1)],1)},s=[],o=e("5530"),r=e("6c33"),a=e("2f62"),c={data:function(){return{nickname:""}},methods:{login:function(){var t=this;this.$store.dispatch(r["b"],this.nickname).then((function(){"admin"==t.user.role?t.$router.push({path:"/app/dashboard"}):t.$router.push({path:"/app/chat"})}))}},computed:Object(o["a"])({},Object(a["c"])({user:function(t){return t.auth.user}}))},u=c,h=(e("0dc7"),e("2877")),d=e("6544"),l=e.n(d),v=e("8336"),f=e("b0af"),p=e("99d9"),w=e("ce7e"),m=e("2fa4"),g=e("8654"),x=(e("99af"),e("7db0"),e("c740"),e("13b3"),e("c3f0")),b=e("afdd"),T=e("9d26"),C=e("604c"),I=C["a"].extend({name:"v-window",directives:{Touch:x["a"]},provide:function(){return{windowGroup:this}},props:{activeClass:{type:String,default:"v-window-item--active"},continuous:Boolean,mandatory:{type:Boolean,default:!0},nextIcon:{type:[Boolean,String],default:"$next"},prevIcon:{type:[Boolean,String],default:"$prev"},reverse:Boolean,showArrows:Boolean,showArrowsOnHover:Boolean,touch:Object,touchless:Boolean,value:{required:!1},vertical:Boolean},data:function(){return{changedByDelimiters:!1,internalHeight:void 0,transitionHeight:void 0,transitionCount:0,isBooted:!1,isReverse:!1}},computed:{isActive:function(){return this.transitionCount>0},classes:function(){return Object(o["a"])(Object(o["a"])({},C["a"].options.computed.classes.call(this)),{},{"v-window--show-arrows-on-hover":this.showArrowsOnHover})},computedTransition:function(){if(!this.isBooted)return"";var t=this.vertical?"y":"x",n=this.internalReverse?!this.isReverse:this.isReverse,e=n?"-reverse":"";return"v-window-".concat(t).concat(e,"-transition")},hasActiveItems:function(){return Boolean(this.items.find((function(t){return!t.disabled})))},hasNext:function(){return this.continuous||this.internalIndex<this.items.length-1},hasPrev:function(){return this.continuous||this.internalIndex>0},internalIndex:function(){var t=this;return this.items.findIndex((function(n,e){return t.internalValue===t.getValue(n,e)}))},internalReverse:function(){return this.$vuetify.rtl?!this.reverse:this.reverse}},watch:{internalIndex:function(t,n){this.isReverse=this.updateReverse(t,n)}},mounted:function(){var t=this;window.requestAnimationFrame((function(){return t.isBooted=!0}))},methods:{genContainer:function(){var t=[this.$slots.default];return this.showArrows&&t.push(this.genControlIcons()),this.$createElement("div",{staticClass:"v-window__container",class:{"v-window__container--is-active":this.isActive},style:{height:this.internalHeight||this.transitionHeight}},t)},genIcon:function(t,n,e){var i,s=this,o={click:function(t){t.stopPropagation(),s.changedByDelimiters=!0,e()}},r={"aria-label":this.$vuetify.lang.t("$vuetify.carousel.".concat(t))},a=null!=(i=null==this.$scopedSlots[t]?void 0:this.$scopedSlots[t]({on:o,attrs:r}))?i:[this.$createElement(b["a"],{props:{icon:!0},attrs:r,on:o},[this.$createElement(T["a"],{props:{large:!0}},n)])];return this.$createElement("div",{staticClass:"v-window__".concat(t)},a)},genControlIcons:function(){var t=[],n=this.$vuetify.rtl?this.nextIcon:this.prevIcon;if(this.hasPrev&&n&&"string"===typeof n){var e=this.genIcon("prev",n,this.prev);e&&t.push(e)}var i=this.$vuetify.rtl?this.prevIcon:this.nextIcon;if(this.hasNext&&i&&"string"===typeof i){var s=this.genIcon("next",i,this.next);s&&t.push(s)}return t},getNextIndex:function(t){var n=(t+1)%this.items.length,e=this.items[n];return e.disabled?this.getNextIndex(n):n},getPrevIndex:function(t){var n=(t+this.items.length-1)%this.items.length,e=this.items[n];return e.disabled?this.getPrevIndex(n):n},next:function(){if(this.hasActiveItems&&this.hasNext){var t=this.getNextIndex(this.internalIndex),n=this.items[t];this.internalValue=this.getValue(n,t)}},prev:function(){if(this.hasActiveItems&&this.hasPrev){var t=this.getPrevIndex(this.internalIndex),n=this.items[t];this.internalValue=this.getValue(n,t)}},updateReverse:function(t,n){var e=this.items.length,i=e-1;return e<=2?t<n:t===i&&0===n||(0!==t||n!==i)&&t<n}},render:function(t){var n=this,e={staticClass:"v-window",class:this.classes,directives:[]};if(!this.touchless){var i=this.touch||{left:function(){n.$vuetify.rtl?n.prev():n.next()},right:function(){n.$vuetify.rtl?n.next():n.prev()},end:function(t){t.stopPropagation()},start:function(t){t.stopPropagation()}};e.directives.push({name:"touch",value:i})}return t("div",e,[this.genContainer()])}}),y=e("9d65"),$=e("4e82"),B=e("80d2"),_=e("58df"),O=Object(_["a"])(y["a"],Object($["a"])("windowGroup","v-window-item","v-window")),j=O.extend().extend().extend({name:"v-window-item",directives:{Touch:x["a"]},props:{disabled:Boolean,reverseTransition:{type:[Boolean,String],default:void 0},transition:{type:[Boolean,String],default:void 0},value:{required:!1}},data:function(){return{isActive:!1,inTransition:!1}},computed:{classes:function(){return this.groupClasses},computedTransition:function(){return this.windowGroup.internalReverse?"undefined"!==typeof this.reverseTransition?this.reverseTransition||"":this.windowGroup.computedTransition:"undefined"!==typeof this.transition?this.transition||"":this.windowGroup.computedTransition}},methods:{genDefaultSlot:function(){return this.$slots.default},genWindowItem:function(){return this.$createElement("div",{staticClass:"v-window-item",class:this.classes,directives:[{name:"show",value:this.isActive}],on:this.$listeners},this.genDefaultSlot())},onAfterTransition:function(){this.inTransition&&(this.inTransition=!1,this.windowGroup.transitionCount>0&&(this.windowGroup.transitionCount--,0===this.windowGroup.transitionCount&&(this.windowGroup.transitionHeight=void 0)))},onBeforeTransition:function(){this.inTransition||(this.inTransition=!0,0===this.windowGroup.transitionCount&&(this.windowGroup.transitionHeight=Object(B["f"])(this.windowGroup.$el.clientHeight)),this.windowGroup.transitionCount++)},onTransitionCancelled:function(){this.onAfterTransition()},onEnter:function(t){var n=this;this.inTransition&&this.$nextTick((function(){n.computedTransition&&n.inTransition&&(n.windowGroup.transitionHeight=Object(B["f"])(t.clientHeight))}))}},render:function(t){var n=this;return t("transition",{props:{name:this.computedTransition},on:{beforeEnter:this.onBeforeTransition,afterEnter:this.onAfterTransition,enterCancelled:this.onTransitionCancelled,beforeLeave:this.onBeforeTransition,afterLeave:this.onAfterTransition,leaveCancelled:this.onTransitionCancelled,enter:this.onEnter}},this.showLazyContent((function(){return[n.genWindowItem()]})))}}),A=Object(h["a"])(u,i,s,!1,null,null,null);n["default"]=A.exports;l()(A,{VBtn:v["a"],VCard:f["a"],VCardActions:p["a"],VCardText:p["b"],VCardTitle:p["c"],VDivider:w["a"],VSpacer:m["a"],VTextField:g["a"],VWindow:I,VWindowItem:j})},"0dc7":function(t,n,e){"use strict";e("6458")},"13b3":function(t,n,e){},"2fa4":function(t,n,e){"use strict";e("20f6");var i=e("80d2");n["a"]=Object(i["g"])("spacer","div","v-spacer")},6458:function(t,n,e){},afdd:function(t,n,e){"use strict";var i=e("8336");n["a"]=i["a"]},c3f0:function(t,n,e){"use strict";e("4160"),e("159b");var i=e("80d2"),s=function(t){var n=t.touchstartX,e=t.touchendX,i=t.touchstartY,s=t.touchendY,o=.5,r=16;t.offsetX=e-n,t.offsetY=s-i,Math.abs(t.offsetY)<o*Math.abs(t.offsetX)&&(t.left&&e<n-r&&t.left(t),t.right&&e>n+r&&t.right(t)),Math.abs(t.offsetX)<o*Math.abs(t.offsetY)&&(t.up&&s<i-r&&t.up(t),t.down&&s>i+r&&t.down(t))};function o(t,n){var e=t.changedTouches[0];n.touchstartX=e.clientX,n.touchstartY=e.clientY,n.start&&n.start(Object.assign(t,n))}function r(t,n){var e=t.changedTouches[0];n.touchendX=e.clientX,n.touchendY=e.clientY,n.end&&n.end(Object.assign(t,n)),s(n)}function a(t,n){var e=t.changedTouches[0];n.touchmoveX=e.clientX,n.touchmoveY=e.clientY,n.move&&n.move(Object.assign(t,n))}function c(t){var n={touchstartX:0,touchstartY:0,touchendX:0,touchendY:0,touchmoveX:0,touchmoveY:0,offsetX:0,offsetY:0,left:t.left,right:t.right,up:t.up,down:t.down,start:t.start,move:t.move,end:t.end};return{touchstart:function(t){return o(t,n)},touchend:function(t){return r(t,n)},touchmove:function(t){return a(t,n)}}}function u(t,n,e){var s=n.value,o=s.parent?t.parentElement:t,r=s.options||{passive:!0};if(o){var a=c(n.value);o._touchHandlers=Object(o._touchHandlers),o._touchHandlers[e.context._uid]=a,Object(i["p"])(a).forEach((function(t){o.addEventListener(t,a[t],r)}))}}function h(t,n,e){var s=n.value.parent?t.parentElement:t;if(s&&s._touchHandlers){var o=s._touchHandlers[e.context._uid];Object(i["p"])(o).forEach((function(t){s.removeEventListener(t,o[t])})),delete s._touchHandlers[e.context._uid]}}var d={inserted:u,unbind:h};n["a"]=d}}]);
//# sourceMappingURL=chunk-63615f18.08160f20.js.map